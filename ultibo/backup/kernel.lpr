program kernel;

{$mode objfpc}{$H+}{$inline on}

uses
  RaspberryPi3,
  GlobalConfig,
  GlobalConst,
  GlobalTypes,
  Platform,
  Threads,
  Framebuffer,
  Console,
  SysUtils,
  Serial;

var
  PageSize: Integer;
  CurrentPage: Integer;
  BufferStart: Pointer;
  FramebufferDevice: PFramebufferDevice;
  FramebufferProperties: TFramebufferProperties;
  WindowHandle: TWindowHandle;

procedure PutPixel8(X, Y: Integer; Color: Byte); inline;
var
  PixelOffset: Cardinal;
begin
  if (X >= 0) and (Y >= 0) and (X < FramebufferProperties.PhysicalWidth) and (Y < FramebufferProperties.PhysicalHeight) then
  begin
    PixelOffset := (CurrentPage * PageSize) + X + (Y * FramebufferProperties.Pitch);
    PByte(BufferStart + PixelOffset)^ := Color;
  end;
end;

procedure FillRect(X, Y, Width, Height, Color: Integer); inline;
var
  CurrentY: Integer;
  PixelOffset: Cardinal;
begin
  PixelOffset := (CurrentPage * PageSize) + X + (Y * FramebufferProperties.Pitch);
  for CurrentY := 0 to Height - 1 do
  begin
    FillChar(Pointer(BufferStart + PixelOffset)^, Width, Color);
    Inc(PixelOffset, FramebufferProperties.Pitch);
  end;
end;

procedure ClearScreen(Color: Integer); inline;
begin
  FillChar(Pointer(BufferStart + (CurrentPage * PageSize))^, PageSize, Color);
end;

procedure DrawButton(X, Y, W, H: Integer; Color: Byte);
begin
  FillRect(X, Y, W, H, Color);
  FillRect(X, Y, W, 1, 15);
  FillRect(X, Y + H - 1, W, 1, 15);
  FillRect(X, Y, 1, H, 15);
  FillRect(X + W - 1, Y, 1, H, 15);
end;

procedure Draw;
begin
  CurrentPage := (CurrentPage + 1) mod 2;
  ClearScreen(0);
  DrawButton(10, 10, 100, 30, 8);

  if (FramebufferProperties.Flags and FRAMEBUFFER_FLAG_CACHED) <> 0 then
    CleanDataCacheRange(PtrUInt(BufferStart) + (CurrentPage * PageSize), PageSize);

  FramebufferDeviceSetOffset(FramebufferDevice, 0, CurrentPage * FramebufferProperties.PhysicalHeight, True);

  if (FramebufferProperties.Flags and FRAMEBUFFER_FLAG_SYNC) <> 0 then
    FramebufferDeviceWaitSync(FramebufferDevice)
  else
    MicrosecondDelay(1000000 div 60);
end;

var
  Count: LongWord;
  Character: Char;
  Characters: String;

begin
  ThreadSetCPU(ThreadGetCurrent, CPU_ID_3);
  Sleep(0);

  FramebufferDevice := FramebufferDeviceGetDefault;
  if FramebufferDevice <> nil then
  begin
    FramebufferDeviceGetProperties(FramebufferDevice, @FramebufferProperties);
    FramebufferDeviceRelease(FramebufferDevice);
    Sleep(1000);

    FramebufferProperties.Depth := 8;
    FramebufferProperties.VirtualWidth := FramebufferProperties.PhysicalWidth;
    FramebufferProperties.VirtualHeight := FramebufferProperties.PhysicalHeight * 2;

    FramebufferDeviceAllocate(FramebufferDevice, @FramebufferProperties);
    Sleep(1000);

    FramebufferDeviceGetProperties(FramebufferDevice, @FramebufferProperties);
    BufferStart := Pointer(FramebufferProperties.Address);
    PageSize := FramebufferProperties.Pitch * FramebufferProperties.PhysicalHeight;
    CurrentPage := 0;
  end;

  WindowHandle := ConsoleWindowCreate(ConsoleDeviceGetDefault, CONSOLE_POSITION_FULLSCREEN, True);

  ConsoleWindowWriteLn(WindowHandle, 'Welcome to the UI + Serial demo');

  if SerialOpen(9600, SERIAL_DATA_8BIT, SERIAL_STOP_1BIT, SERIAL_PARITY_NONE, SERIAL_FLOW_NONE, 0, 0) = ERROR_SUCCESS then
  begin
    ConsoleWindowWriteLn(WindowHandle, 'Serial device opened. Type QUIT to exit echo loop.');
    Count := 0;
    Characters := '';

    while True do
    begin
      SerialRead(@Character, SizeOf(Character), Count);

      if Character = #13 then
      begin
        ConsoleWindowWriteLn(WindowHandle, 'Received: ' + Characters);
        if Uppercase(Characters) = 'QUIT' then
        begin
          Characters := 'Goodbye!' + Chr(13) + Chr(10);
          SerialWrite(PChar(Characters), Length(Characters), Count);
          Sleep(1000);
          Break;
        end;
        Characters := Characters + Chr(13) + Chr(10);
        SerialWrite(PChar(Characters), Length(Characters), Count);
        Characters := '';
      end
      else
      begin
        Characters := Characters + Character;
      end;
    end;

    SerialClose;
    ConsoleWindowWriteLn(WindowHandle, 'Serial device closed');
  end
  else
  begin
    ConsoleWindowWriteLn(WindowHandle, 'Failed to open serial device');
  end;

  while True do
    Draw;
end.

