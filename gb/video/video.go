package video

import (
	"github.com/drhelius/demo-emulator/gb/memory"
	"github.com/drhelius/demo-emulator/gb/util"
)

var (
	// GbFrameBuffer is the internal Game Boy frame buffer
	GbFrameBuffer [util.GbWidth * util.GbHeight]uint8
	// ScreenEnabled keeps track of the screen state
	ScreenEnabled    bool
	statusMode       uint8
	statusModeCycles uint32
	lyCounter        uint8
	mem              memory.IMemory
)

func init() {
	statusMode = 1
	lyCounter = 144
	ScreenEnabled = true
}

// SetMem injects the memory impl
func SetMem(m memory.IMemory) {
	mem = m
}

// Tick runs the video eumulation n cycles
// Then updates the frameBuffer and returns true if the simulation reached the vblank
func Tick(cycles uint32) bool {
	vblank := false
	statusModeCycles += cycles

	switch statusMode {
	case 0:
		// During H-BLANK
		if statusModeCycles >= 204 {
			statusModeCycles -= 204
			statusMode = 2

			lyCounter++
			mem.Write(0xFF44, lyCounter)
			CompareLYToLYC()

			if lyCounter == 144 {
				statusMode = 1
				//m_iStatusVBlankLine = 0
			}

			//updateStatRegister()
		}
	case 1:
		// During V-BLANK
	case 2:
		// During searching OAM RAM
	case 3:
		// During transfering data to LCD driver

	}

	/*
	   statusModeCycles += clockCycles;

	           switch (statusMode)
	           {
	               // During H-BLANK
	               case 0:
	               {
	                   if (statusModeCycles >= 204)
	                   {
	                       statusModeCycles -= 204;
	                       statusMode = 2;

	                       lyCounter++;
	                       m_pMemory->Load(0xFF44, lyCounter);
	                       CompareLYToLYC();

	                       if (m_bCGB && m_pMemory->IsHDMAEnabled() && (!m_pProcessor->Halted() || m_pProcessor->InterruptIsAboutToRaise()))
	                       {
	                           unsigned int cycles = m_pMemory->PerformHDMA();
	                           statusModeCycles += cycles;
	                           clockCycles += cycles;
	                       }

	                       if (lyCounter == 144)
	                       {
	                           statusMode = 1;
	                           m_iStatusVBlankLine = 0;
	                           statusModeCyclesAux = statusModeCycles;

	                           m_pProcessor->RequestInterrupt(Processor::VBlank_Interrupt);

	                           m_IRQ48Signal &= 0x09;
	                           u8 stat = m_pMemory->Retrieve(0xFF41);
	                           if (IsSetBit(stat, 4))
	                           {
	                               if (!IsSetBit(m_IRQ48Signal, 0) && !IsSetBit(m_IRQ48Signal, 3))
	                               {
	                                   m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
	                               }
	                               m_IRQ48Signal = SetBit(m_IRQ48Signal, 1);
	                           }
	                           m_IRQ48Signal &= 0x0E;

	                           if (m_iHideFrames > 0)
	                               m_iHideFrames--;
	                           else
	                               vblank = true;

	                           m_iWindowLine = 0;
	                       }
	                       else
	                       {
	                           m_IRQ48Signal &= 0x09;
	                           u8 stat = m_pMemory->Retrieve(0xFF41);
	                           if (IsSetBit(stat, 5))
	                           {
	                               if (m_IRQ48Signal == 0)
	                               {
	                                   m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
	                               }
	                               m_IRQ48Signal = SetBit(m_IRQ48Signal, 2);
	                           }
	                           m_IRQ48Signal &= 0x0E;
	                       }

	                       UpdateStatRegister();
	                   }
	                   break;
	               }
	               // During V-BLANK
	               case 1:
	               {
	                   statusModeCyclesAux += clockCycles;

	                   if (statusModeCyclesAux >= 456)
	                   {
	                       statusModeCyclesAux -= 456;
	                       m_iStatusVBlankLine++;

	                       if (m_iStatusVBlankLine <= 9)
	                       {
	                           lyCounter++;
	                           m_pMemory->Load(0xFF44, lyCounter);
	                           CompareLYToLYC();
	                       }
	                   }

	                   if ((statusModeCycles >= 4104) && (statusModeCyclesAux >= 4) && (lyCounter == 153))
	                   {
	                       lyCounter = 0;
	                       m_pMemory->Load(0xFF44, lyCounter);
	                       CompareLYToLYC();
	                   }

	                   if (statusModeCycles >= 4560)
	                   {
	                       statusModeCycles -= 4560;
	                       statusMode = 2;
	                       UpdateStatRegister();
	                       m_IRQ48Signal &= 0x07;


	                       m_IRQ48Signal &= 0x0A;
	                       u8 stat = m_pMemory->Retrieve(0xFF41);
	                       if (IsSetBit(stat, 5))
	                       {
	                           if (m_IRQ48Signal == 0)
	                           {
	                               m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
	                           }
	                           m_IRQ48Signal = SetBit(m_IRQ48Signal, 2);
	                       }
	                       m_IRQ48Signal &= 0x0D;
	                   }
	                   break;
	               }
	               // During searching OAM RAM
	               case 2:
	               {
	                   if (statusModeCycles >= 80)
	                   {
	                       statusModeCycles -= 80;
	                       statusMode = 3;
	                       m_bScanLineTransfered = false;
	                       m_IRQ48Signal &= 0x08;
	                       UpdateStatRegister();
	                   }
	                   break;
	               }
	               // During transfering data to LCD driver
	               case 3:
	               {
	                   if (m_iPixelCounter < 160)
	                   {
	                       m_iTileCycleCounter += clockCycles;
	                       u8 lcdc = m_pMemory->Retrieve(0xFF40);

	                       if (m_bScreenEnabled && IsSetBit(lcdc, 7))
	                       {
	                           while (m_iTileCycleCounter >= 3)
	                           {
	                               RenderBG(lyCounter, m_iPixelCounter, 4);
	                               m_iPixelCounter += 4;
	                               m_iTileCycleCounter -= 3;

	                               if (m_iPixelCounter >= 160)
	                               {
	                                   break;
	                               }
	                           }
	                       }
	                   }

	                   if (statusModeCycles >= 160 && !m_bScanLineTransfered)
	                   {
	                       ScanLine(lyCounter);
	                       m_bScanLineTransfered = true;
	                   }

	                   if (statusModeCycles >= 172)
	                   {
	                       m_iPixelCounter = 0;
	                       statusModeCycles -= 172;
	                       statusMode = 0;
	                       m_iTileCycleCounter = 0;
	                       UpdateStatRegister();

	                       m_IRQ48Signal &= 0x08;
	                       u8 stat = m_pMemory->Retrieve(0xFF41);
	                       if (IsSetBit(stat, 3))
	                       {
	                           if (!IsSetBit(m_IRQ48Signal, 3))
	                           {
	                               m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
	                           }
	                           m_IRQ48Signal = SetBit(m_IRQ48Signal, 0);
	                       }
	                   }
	                   break;
	               }
	           }
	*/
	return vblank
}

// EnableScreen enables the screen
func EnableScreen() {

}

// DisableScreen disables the screen
func DisableScreen() {

}

// ResetWindowLine resets the line counter
func ResetWindowLine() {

}

// CompareLYToLYC compares LY counter with LYC register
func CompareLYToLYC() {

}
