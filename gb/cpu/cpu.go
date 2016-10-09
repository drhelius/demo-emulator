package cpu

import (
	"github.com/drhelius/demo-emulator/gb/memory"
	"github.com/drhelius/demo-emulator/gb/util"
)

// Interrupt types
const (
	InterruptNone    uint8 = 0x00
	InterruptVBlank  uint8 = 0x01
	InterruptLCDSTAT uint8 = 0x02
	InterruptTimer   uint8 = 0x04
	InterruptSerial  uint8 = 0x08
	InterruptJoypad  uint8 = 0x10
)

const (
	flagZero  uint8 = 0x80
	flagSub   uint8 = 0x40
	flagHalf  uint8 = 0x20
	flagCarry uint8 = 0x10
	flagNone  uint8 = 0x00
)

var (
	af          SixteenBitReg
	bc          SixteenBitReg
	de          SixteenBitReg
	hl          SixteenBitReg
	sp          SixteenBitReg
	pc          SixteenBitReg
	mem         memory.IMemory
	ime         bool
	halt        bool
	branchTaken bool
	clockCycles uint32
	divCycles   uint32
	timaCycles  uint32
	imeCycles   int
	skipPCBug   bool
)

func init() {
	pc.SetValue(0x0100)
	sp.SetValue(0xFFFE)
	af.SetValue(0x01B0)
	bc.SetValue(0x0013)
	de.SetValue(0x00D8)
	hl.SetValue(0x014D)
}

// SetMem injects the memory impl
func SetMem(m memory.IMemory) {
	mem = m
}

// Tick runs a single instruction of the processor
// Then returns the number of cycles used
func Tick() uint32 {
	clockCycles = 0

	if halt {
		if interruptPending() != InterruptNone {
			halt = false
		} else {
			clockCycles += 4
		}
	}

	if !halt {
		//fmt.Printf("-> PC: 0x%X  OP: 0x%X\n", pc.GetValue(), mem.Read(pc.GetValue()))
		serveInterrupt(interruptPending())
		runOpcode(fetchOpcode())
	}

	updateTimers()

	if imeCycles > 0 {
		imeCycles -= int(clockCycles)
		if imeCycles <= 0 {
			imeCycles = 0
			ime = true
		}
	}

	return clockCycles
}

// RequestInterrupt is used to raise a new interrupt
func RequestInterrupt(interrupt uint8) {
	mem.Write(0xFF0F, mem.Read(0xFF0F)|interrupt)
}

// ResetDivCycles sets divCycles to 0
func ResetDivCycles() {
	divCycles = 0
}

// ResetTimaCycles sets timaCycles to 0
func ResetTimaCycles() {
	timaCycles = 0
}

func fetchOpcode() uint8 {
	opcode := mem.Read(pc.GetValue())
	if skipPCBug {
		skipPCBug = false
	} else {
		pc.Increment()
	}
	return opcode
}

func runOpcode(opcode uint8) {
	if opcode == 0xCB {
		opcode = fetchOpcode()
		opcodeCBArray[opcode]()
		clockCycles += machineCyclesCB[opcode] * 4
	} else {
		opcodeArray[opcode]()
		if branchTaken {
			branchTaken = false
			clockCycles += machineCyclesBranched[opcode] * 4
		} else {
			clockCycles += machineCycles[opcode] * 4
		}
	}
}

func interruptIsAboutToRaise() bool {
	ieReg := mem.Read(0xFFFF)
	ifReg := mem.Read(0xFF0F)
	return (ifReg & ieReg & 0x1F) != 0
}

func interruptPending() uint8 {
	ieReg := mem.Read(0xFFFF)
	ifReg := mem.Read(0xFF0F)
	ieIf := ieReg & ifReg

	switch {
	case (ieIf & 0x01) != 0:
		return InterruptVBlank
	case (ieIf & 0x02) != 0:
		return InterruptLCDSTAT
	case (ieIf & 0x04) != 0:
		return InterruptTimer
	case (ieIf & 0x08) != 0:
		return InterruptSerial
	case (ieIf & 0x10) != 0:
		return InterruptJoypad
	}

	return InterruptNone
}

func serveInterrupt(interrupt uint8) {
	if ime {
		ifReg := mem.Read(0xFF0F)
		switch interrupt {
		case InterruptVBlank:
			mem.Write(0xFF0F, ifReg&0xFE)
			ime = false
			stackPush(&pc)
			pc.SetValue(0x0040)
			clockCycles += 20
		case InterruptLCDSTAT:
			mem.Write(0xFF0F, ifReg&0xFD)
			ime = false
			stackPush(&pc)
			pc.SetValue(0x0048)
			clockCycles += 20
		case InterruptTimer:
			mem.Write(0xFF0F, ifReg&0xFB)
			ime = false
			stackPush(&pc)
			pc.SetValue(0x0050)
			clockCycles += 20
		case InterruptSerial:
			mem.Write(0xFF0F, ifReg&0xF7)
			ime = false
			stackPush(&pc)
			pc.SetValue(0x0058)
			clockCycles += 20
		case InterruptJoypad:
			mem.Write(0xFF0F, ifReg&0xEF)
			ime = false
			stackPush(&pc)
			pc.SetValue(0x0060)
			clockCycles += 20
		}
	}
}

func updateTimers() {
	divCycles += clockCycles

	for divCycles >= 256 {
		divCycles -= 256
		div := mem.Read(0xFF04)
		div++
		mem.Write(0xFF04, div)
	}

	tac := mem.Read(0xFF07)

	// if tima is running
	if util.IsSetBit(tac, 2) {
		timaCycles += clockCycles

		var freq uint32

		switch tac & 0x03 {
		case 0:
			freq = 1024
		case 1:
			freq = 16
		case 2:
			freq = 64
		case 3:
			freq = 256
		}

		for timaCycles >= freq {
			timaCycles -= freq
			tima := mem.Read(0xFF05)

			if tima == 0xFF {
				tima = mem.Read(0xFF06)
				RequestInterrupt(InterruptTimer)
			} else {
				tima++
			}

			mem.Write(0xFF05, tima)
		}
	}
}
