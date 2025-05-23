//
// bcmpciehostbridge.h
//
// This driver has been ported from the Linux driver which is:
//	drivers/pci/controller/pcie-brcmstb.c
//	Copyright (C) 2009 - 2017 Broadcom
//	Licensed under GPL-2.0
//
// Circle - A C++ bare metal environment for Raspberry Pi
// Copyright (C) 2019-2025  R. Stange <rsta2@o2online.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
#ifndef _circle_bcmpciehostbridge_h
#define _circle_bcmpciehostbridge_h

#include <circle/interrupt.h>
#include <circle/machineinfo.h>
#include <circle/timer.h>
#include <circle/types.h>

struct TPCIeMemoryWindow
{
	u64		pcie_addr;
	u64		cpu_addr;
	u64		size;
};

typedef void TPCIeMSIHandler (unsigned nVector, void *pParam);

struct TPCIeMSIData
{
	uintptr		base;
	u64		target_addr;
	uintptr		intr_base;	// base of interrupt status/set/clr regs
	unsigned	rev;
	TPCIeMSIHandler *handler;
	void		*param;
};

class CBcmPCIeHostBridge	/// Driver for PCIe Host Bridge(s) of Raspberry Pi 4 and 5
{
public:
	/// \param nBus PCIE_BUS_ONBOARD or PCIE_BUS_EXTERNAL
	/// \param pInterrupt Pointer to the interrupt system object
	CBcmPCIeHostBridge (unsigned nBus, CInterruptSystem *pInterrupt);

	~CBcmPCIeHostBridge (void);

	/// \return Operation successful?
	boolean Initialize (void);

	/// \param nClassCode 24-bit PCI class code of the device to be used
	/// \param nSlot Slot number of the device to be used
	/// \param nFunc Function number of the device to be used
	/// \return Operation successful?
	boolean EnableDevice (u32 nClassCode, unsigned nSlot, unsigned nFunc);

#if 0
	boolean ConnectMSI (TPCIeMSIHandler *pHandler, void *pParam);
	void DisconnectMSI (void);
#endif

	/// \param nBus 0-based PCIe bus number
	/// \return Base address of the inbound memory window
	static u64 GetDMAAddress (unsigned nBus)
	{
		return s_nDMAAddress[nBus];
	}

#ifndef NDEBUG
	void DumpStatus (unsigned nSlot, unsigned nFunc);
#endif

private:
	int pcie_probe(void);
	int pcie_setup(void);

	int enable_bridge (void);
	int enable_device (u32 nClassCode, unsigned nSlot, unsigned nFunc);

	int pcie_set_pci_ranges(void);
	int pcie_set_dma_ranges(void);

	void pcie_set_outbound_win(unsigned win, u64 cpu_addr, u64 pcie_addr, u64 size);
	uintptr pcie_map_conf(unsigned busnr, unsigned devfn, int where);

	static uintptr find_pci_capability (uintptr nPCIConfig, u8 uchCapID);

	void pcie_bridge_sw_init_set(unsigned val);
	void pcie_perst_set(unsigned int val);

	bool pcie_link_up(void);
	bool pcie_rc_mode(void);

	int pcie_enable_msi(TPCIeMSIHandler *pHandler, void *pParam);
	static void msi_set_regs(TPCIeMSIData *msi);

	static int cfg_index(int busnr, int devfn, int reg);

	static void set_gen(uintptr base, int gen);

	static const char *link_speed_to_str(int s);
	static int encode_ibar_size(u64 size);

	static u32 rd_fld(uintptr p, u32 mask, int shift);
	static void wr_fld(uintptr p, u32 mask, int shift, u32 val);
	static void wr_fld_rb(uintptr p, u32 mask, int shift, u32 val);

	static void InterruptHandler (void *pParam);

	void usleep_range (unsigned min, unsigned max);
	void msleep (unsigned ms);
	static int ilog2 (u64 v);

#if RASPPI >= 5
	void pcie_set_tc_qos(void);
	void pcie_config_clkreq(void);
	void pcie_munge_pll(void);

	int pcie_mdio_read(u8 port, u8 regad, u32 *val);
	int pcie_mdio_write(u8 port, u8 regad, u16 wrdata);
	static u32 pcie_mdio_form_pkt(int port, int regad, int cmd);

	int reset_assert(unsigned long id);
	int reset_deassert(unsigned long id);
	int rescal_reset_deassert(void);
#endif

private:
	unsigned		 m_nBus;
	CInterruptSystem	*m_pInterrupt;

	uintptr			 m_base;		// mmio base address
	unsigned		 m_rev;			// controller revision

	TPCIeMemoryWindow	 m_out_wins[1];		// outbound window
	int			 m_num_out_wins;

	TPCIeMemoryWindow	 m_dma_ranges[1];	// inbound window
	int			 m_num_dma_ranges;
	u64			 m_scb_size[1];
	int			 m_num_scbs;

	u64			 m_msi_target_addr;
	TPCIeMSIData		*m_msi;

	static u64		 s_nDMAAddress[PCIE_BUS_NUM];
};

#endif
