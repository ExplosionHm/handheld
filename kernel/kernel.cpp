#include "kernel.h"
#include <circle/gpiopin.h>
#include <circle/timer.h>

CKernel::CKernel (void)
{
}

CKernel::~CKernel (void)
{
}

boolean CKernel::Initialize (void)
{
	return TRUE;
}

TShutdownMode CKernel::Run (void)
{
	int i = 0;
	// flash the Act LED on and off
	while (1)
	{
		m_ActLED.On();
		CTimer::SimpleMsDelay(200);

		m_ActLED.Off();
		CTimer::SimpleMsDelay(500);
	}

	return ShutdownReboot;
}
