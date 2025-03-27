# handheld

## Goal
I am creating a custom handheld gaming console powered by the Raspberry Pi Zero 2 W (RP3A0). This console will only run my own games, primarily focusing on 2D games.

## Operating System
To maximize performance and eliminate unnecessary overhead, I will develop my own bare-metal OS. This OS will be built using the [Circle](<https://github.com/rsta2/circle>) library, as it provides low-level functionality while avoiding the complexity of writing everything from scratch (saves my sanity).

## Game Development
Rather than relying on emulation (e.g., SNES, N64, PS1), I will develop a collection of original games inspired by classics like Super Mario World, Crash Bandicoot, and Super Mario 64. These games will be optimized for my hardware, ensuring smooth performance.

All games will be made in C++

## Hardware Components
 - Raspberry Pi Zero 2 W (RP3A0)
 - ILI9341 LCD Display Module (240x320, with touch support)
 - Speaker (for in-game audio)
 - Physical buttons (for gameplay controls)
 - Power source (likely a LiPo battery solution, such as PiSugar 2 Pro)

### To Be Determined…
 - Battery capacity and management
 - Additional components (e.g., cooling solutions, additional inputs)
 - Best approach for handling graphics and game logic