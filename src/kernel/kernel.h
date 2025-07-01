#ifndef _kernel_h
#define _kernel_h

#include <circle/actled.h>
#include <circle/koptions.h>
#include <circle/2dgraphics.h>
#include <circle/serial.h>
#include <circle/types.h>

#include <circle/devicenameservice.h>
#include <circle/fs/fat/fatfs.h>
#include <circle/fs/fat/fat.h>
#include <circle/fs/fat/fatinfo.h>
#include <circle/fs/fat/fatcache.h>
#include <circle/interrupt.h>
#include <circle/timer.h>
#include <SDCard/emmc.h>

enum TShutdownMode
{
	ShutdownNone,
	ShutdownHalt,
	ShutdownReboot
};

class CKernel
{
public:
	CKernel(void);
	~CKernel(void);

	boolean Initialize(void);

	TShutdownMode Run(void);

private:
	CFATCache m_fatcache;
	CFATInfo m_fatinfo;
	CFAT m_fat;
	CFATFileSystem m_fatfs;

	CKernelOptions m_options;
	CEMMCDevice m_emmc;

	CInterruptSystem m_interupt;
	CTimer m_timer;
	CActLED m_ActLED;
	C2DGraphics m_2DGraphics;

	CDeviceNameService m_DeviceNameService;
};
#endif