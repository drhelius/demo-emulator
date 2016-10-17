package cpu

import "github.com/drhelius/demo-emulator/gb/mapper"

// Interrupt types
const (
	InterruptNone    uint8 = 0x00
	InterruptVBlank  uint8 = 0x01
	InterruptLCDSTAT uint8 = 0x02
	InterruptTimer   uint8 = 0x04
	InterruptSerial  uint8 = 0x08
	InterruptJoypad  uint8 = 0x10
)

// All the flags in the F register
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
	mem         mapper.Mapper
	ime         bool
	halted      bool
	branchTaken bool
	clockCycles uint
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

// SetMapper injects the memory impl
func SetMapper(m mapper.Mapper) {
	mem = m
}

// Tick runs a single instruction of the processor
// Then returns the number of cycles used
func Tick() uint {
	clockCycles = 0

	if halted {
		// if an interrupt is pending leave halt
		if interruptPending() != InterruptNone {
			halted = false
		} else {
			clockCycles += 4
		}
	}

	if !halted {
		// acknowledge any pending interrupt s
		serveInterrupt(interruptPending())

		// fetch the next opcode and execute it
		runOpcode(fetchOpcode())
	}

	updateTimers()
	updateSerial()

	// this is in order to delay the activation
	// of ima one instruction
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

func fetchOpcode() uint8 {
	opcode := mem.Read(pc.GetValue())

	// if there is an interrupt pending and
	// the cpu is halted it fails to advance the PC register
	// once the cpu resumes operation
	// this bug is present in all the original DMGs
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
