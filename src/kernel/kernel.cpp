#include "kernel.h"
#include <circle/gpiopin.h>
#include <circle/timer.h>
#include <circle/2dgraphics.h>
#include <circle/display.h>
#include <SDCard/emmc.h>

static const T2DColor c1 = COLOR2D(255, 255, 0);
static const T2DColor c2 = COLOR2D(127, 255, 255);
static const T2DColor c3 = COLOR2D(100, 50, 60);

#define PARTITION "emmc1-1"
#define FILENAME "circle.txt"

const unsigned int pixelBuffer[9] = {
	c1, c2, c1,
	c2, c3, c2,
	c1, c2, c1,
};

CKernel::CKernel(void)
	: m_2DGraphics(m_Options.GetWidth(), m_Options.GetHeight(), TRUE),
	  m_EMMC(&m_Interrupt, &m_Timer, &m_ActLED)
{
}

boolean CKernel::Initialize(void)
{
	
	return m_2DGraphics.Initialize();
}

static CString msg = "Hello world!";

TShutdownMode CKernel::Run(void)
{
	// flash the Act LED on and off
	unsigned int i = 0;
	while (1)
	{
		if ((i & 1) == 1) {
			m_2DGraphics.ClearScreen(COLOR2D(255, 0, 0));
		} else {
			m_2DGraphics.ClearScreen(COLOR2D(0, 0, 255));
		}
		m_2DGraphics.DrawImageRect(m_2DGraphics.GetWidth()/2, m_2DGraphics.GetHeight()/2, 3, 3, 0, 0, &pixelBuffer);
		m_2DGraphics.UpdateDisplay();
		m_ActLED.On();
		CTimer::SimpleMsDelay(500);

		m_ActLED.Off();
		CTimer::SimpleMsDelay(500);
		i++;
	}

	return ShutdownHalt;
}
