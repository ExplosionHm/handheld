#include "kernel.h"
#include <circle/gpiopin.h>
#include <circle/timer.h>
#include <circle/2dgraphics.h>
#include <circle/display.h>
#include <circle/serial.h>

bool serial = false;

CKernel::CKernel(void)
	: m_fatcache(),
	  m_fatinfo(&m_fatcache),
	  m_fat(&m_fatcache, &m_fatinfo),
	  m_fatfs(),
	  m_emmc(&m_interupt, &m_timer, &m_ActLED),
	  m_timer(&m_interupt),
	  m_2DGraphics(m_options.GetWidth(), m_options.GetHeight(), TRUE)
{
}

boolean CKernel::Initialize(void)
{
	/*if (!m_Serial.Initialize()) {
		return false;
	}
	serial = true;
	char msg[13] = "Hello world!";
	m_Serial.Write(&msg, sizeof(msg));*/

	if (!m_fatinfo.Initialize()) {
		return false;
	}

	if (!m_timer.Initialize()) {
		return false;
	}

	if (!m_emmc.Initialize())
	{
		return false;
	}

	if (!m_fatfs.Mount(&m_emmc)) {
		return false;
	}
	

	m_DeviceNameService.AddDevice("fs", &m_emmc, TRUE);
	return m_2DGraphics.Initialize();
}

TShutdownMode CKernel::Run(void)
{
	/*CString msgexts[3] = {
		".",
		"..",
		"..."};
	CString result = "";
	while (1)
	{
		m_2DGraphics.DrawRect(0, 0, m_Options.GetWidth(), m_Options.GetHeight(), COLOR2D(32, 32, 32));
		m_2DGraphics.UpdateDisplay();
		if (serial) {
			m_ActLED.Blink(5, 200, 100);
		} else {
			m_ActLED.Blink(3, 200, 100);
		}
		
		CTimer::SimpleMsDelay(1000);
	}*/

	while (1)
	{
		m_2DGraphics.ClearScreen(COLOR2D(0, 0, 0));
		m_2DGraphics.DrawRect(m_2DGraphics.GetWidth() / 4, m_2DGraphics.GetHeight()/4, m_2DGraphics.GetWidth() / 2, m_2DGraphics.GetHeight() / 2, COLOR2D(10, 40, 20));
		m_2DGraphics.UpdateDisplay();

		m_ActLED.Blink(1, 200, 200);
		CTimer::SimpleMsDelay(1000);
	}

	return ShutdownHalt;
}
