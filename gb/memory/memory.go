package memory

import (
	"fmt"

	"github.com/drhelius/demo-emulator/gb/input"
)

var (
	memoryMap = make([]uint8, 0x10000)
	rom       []uint8
)

// SetupROM Receives the rom data
func SetupROM(r []uint8) {
	rom = r
}

// Read returns the 8 bit value at the 16 bit address of the memory
func Read(addr uint16) uint8 {
	if addr < 0x8000 {
		// ROM
		return rom[addr]
	}

	if addr >= 0xFF00 {
		switch addr {
		case 0xFF00:
			// P1
			return input.Read()
		case 0xFF07:
			// TAC
			return memoryMap[addr] | 0xF8
		case 0xFF0F:
			// IF
			return memoryMap[addr] | 0xE0
		case 0xFF41:
			// STAT
			return memoryMap[addr] | 0x80
		//case 0xFF44:
		//    return (m_pVideo->IsScreenEnabled() ? m_pMemory->Retrieve(0xFF44) : 0x00);
		case 0xFF4F:
			// VBK
			return memoryMap[addr] | 0xFE
		}
	}

	return memoryMap[addr]
}

// Write stores the 8 bit value at the 16 bit address of the memory
func Write(addr uint16, value uint8) {

	if addr < 0x8000 {
		// ROM
		fmt.Printf("** Attempting to write on ROM address %X %X\n", addr, value)
		return
	}

	if addr >= 0xFF00 {
		switch addr {
		case 0xFF00:
			// P1
			input.Write(value)
			break
		case 0xFF04:
			// DIV
			//m_pProcessor->ResetDIVCycles();
			break
		case 0xFF07:
			// TAC
			/*
			   value &= 0x07;
			   u8 current_tac = m_pMemory->Retrieve(0xFF07);
			   if ((current_tac & 0x03) != (value & 0x03))
			   {
			       m_pProcessor->ResetTIMACycles();
			   }
			   m_pMemory->Load(address, value);
			*/
		case 0xFF0F:
			// IF
			memoryMap[addr] = value & 0x1F
		case 0xFF40:
			// LCDC
			/*
			   u8 current_lcdc = m_pMemory->Retrieve(0xFF40);
			   u8 new_lcdc = value;
			   m_pMemory->Load(address, new_lcdc);
			   if (!IsSetBit(current_lcdc, 5) && IsSetBit(new_lcdc, 5))
			       m_pVideo->ResetWindowLine();
			   if (IsSetBit(new_lcdc, 7))
			       m_pVideo->EnableScreen();
			   else
			       m_pVideo->DisableScreen();
			*/
		case 0xFF41:
			// STAT
			/*
			   u8 current_stat = m_pMemory->Retrieve(0xFF41) & 0x07;
			   u8 new_stat = (value & 0x78) | (current_stat & 0x07);
			   m_pMemory->Load(address, new_stat);
			   u8 lcdc = m_pMemory->Retrieve(0xFF40);
			   u8 signal = m_pVideo->GetIRQ48Signal();
			   int mode = m_pVideo->GetCurrentStatusMode();
			   signal &= ((new_stat >> 3) & 0x0F);
			   m_pVideo->SetIRQ48Signal(signal);

			   if (IsSetBit(lcdc, 7))
			   {
			       if (IsSetBit(new_stat, 3) && (mode == 0))
			       {
			           if (signal == 0)
			           {
			               m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
			           }
			           signal = SetBit(signal, 0);
			       }
			       if (IsSetBit(new_stat, 4) && (mode == 1))
			       {
			           if (signal == 0)
			           {
			               m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
			           }
			           signal = SetBit(signal, 1);
			       }
			       if (IsSetBit(new_stat, 5) && (mode == 2))
			       {
			           if (signal == 0)
			           {
			               m_pProcessor->RequestInterrupt(Processor::LCDSTAT_Interrupt);
			           }
			           //signal = SetBit(signal, 2);
			       }
			       m_pVideo->CompareLYToLYC();
			   }
			*/
		case 0xFF44:
			// LY
			/*
			   u8 current_ly = m_pMemory->Retrieve(0xFF44);
			   if (IsSetBit(current_ly, 7) && !IsSetBit(value, 7))
			   {
			       m_pVideo->DisableScreen();
			   }
			*/
		case 0xFF45:
		/*
		   // LYC
		   u8 current_lyc = m_pMemory->Retrieve(0xFF45);
		   if (current_lyc != value)
		   {
		       m_pMemory->Load(0xFF45, value);
		       u8 lcdc = m_pMemory->Retrieve(0xFF40);
		       if (IsSetBit(lcdc, 7))
		       {
		           m_pVideo->CompareLYToLYC();
		       }
		   }
		*/
		case 0xFF46:
			// DMA
			//m_pMemory->Load(address, value);
			//m_pMemory->PerformDMA(value);
		case 0xFFFF:
			// IE
			memoryMap[addr] = value & 0x1F
			break
		default:
			memoryMap[addr] = value
		}

	}

}
