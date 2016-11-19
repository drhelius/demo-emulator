package cpu

import "fmt"

func clearAllFlags() {
	initFlagReg(flagNone)
}

func setZeroFlagFromResult(result uint8) {
	if result == 0 {
		setFlag(flagZero)
	}
}

func initFlagReg(flag uint8) {
	af.SetLow(flag)
}

func flipFlag(flag uint8) {
	af.SetLow(af.GetLow() ^ flag)
}

func setFlag(flag uint8) {
	af.SetLow(af.GetLow() | flag)
}

func resetFlag(flag uint8) {
	af.SetLow(af.GetLow() &^ flag)
}

func isSetFlag(flag uint8) bool {
	return (af.GetLow() & flag) != 0
}

func stackPush(reg *SixteenBitReg) {
	sp.Decrement()
	mem.Write(sp.GetValue(), reg.GetHigh())
	sp.Decrement()
	mem.Write(sp.GetValue(), reg.GetLow())
}

func stackPop(reg *SixteenBitReg) {
	reg.SetLow(mem.Read(sp.GetValue()))
	sp.Increment()
	reg.SetHigh(mem.Read(sp.GetValue()))
	sp.Increment()
}

func invalidOPCode() {
	fmt.Println("INVALID opcode")
}

func opcodesLDValueToReg(reg1 *EightBitReg, value uint8) {
	reg1.SetValue(value)
}

func opcodesLDAddrToReg(reg *EightBitReg, address uint16) {
	reg.SetValue(mem.Read(address))
}

func opcodesLDValueToAddr(address uint16, value uint8) {
	mem.Write(address, value)
}

func opcodesOR(number uint8) {
	result := af.GetHigh() | number
	af.SetHigh(result)
	clearAllFlags()
	setZeroFlagFromResult(result)
}

func opcodesXOR(number uint8) {
	result := af.GetHigh() ^ number
	af.SetHigh(result)
	clearAllFlags()
	setZeroFlagFromResult(result)
}

func opcodesAND(number uint8) {
	result := af.GetHigh() & number
	af.SetHigh(result)
	if result == 0 {
		setFlag(flagZero)
	} else {
		resetFlag(flagZero)
	}
	resetFlag(flagNegative)
	setFlag(flagHalf)
	resetFlag(flagCarry)
}

func opcodesCP(number uint8) {
	initFlagReg(flagNegative)
	if af.GetHigh() < number {
		setFlag(flagCarry)
	} else if af.GetHigh() == number {
		setFlag(flagZero)
	}
	if ((af.GetHigh() - number) & 0xF) > (af.GetHigh() & 0xF) {
		setFlag(flagHalf)
	}
}

func opcodesINC(reg *EightBitReg) {
	result := reg.GetValue() + 1
	reg.SetValue(result)
	if result == 0 {
		setFlag(flagZero)
	} else {
		resetFlag(flagZero)
	}
	resetFlag(flagNegative)
	if (result & 0x0F) == 0 {
		setFlag(flagHalf)
	} else {
		resetFlag(flagHalf)
	}
}

func opcodesINCHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	result++
	mem.Write(address, result)
	if isSetFlag(flagCarry) {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	setZeroFlagFromResult(result)
	if (result & 0x0F) == 0x00 {
		setFlag(flagHalf)
	}
}

func opcodesDEC(reg *EightBitReg) {
	result := reg.GetValue()
	result--
	reg.SetValue(result)
	if isSetFlag(flagCarry) {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	setFlag(flagNegative)
	setZeroFlagFromResult(result)
	if (result & 0x0F) == 0x0F {
		setFlag(flagHalf)
	}
}

func opcodesDECHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	result--
	mem.Write(address, result)
	if isSetFlag(flagCarry) {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	setFlag(flagNegative)
	setZeroFlagFromResult(result)
	if (result & 0x0F) == 0x0F {
		setFlag(flagHalf)
	}
}

func opcodesADD(number uint8) {
	result := uint(af.GetHigh()) + uint(number)
	carrybits := uint(af.GetHigh()) ^ uint(number) ^ result
	af.SetHigh(uint8(result))
	clearAllFlags()
	setZeroFlagFromResult(uint8(result))
	if (carrybits & 0x100) != 0 {
		setFlag(flagCarry)
	}
	if (carrybits & 0x10) != 0 {
		setFlag(flagHalf)
	}
}

func opcodesADC(number uint8) {
	var carry uint
	if isSetFlag(flagCarry) {
		carry = 1
	} else {
		carry = 0
	}
	result := uint(af.GetHigh()) + uint(number) + carry
	if uint8(result) == 0 {
		setFlag(flagZero)
	} else {
		resetFlag(flagZero)
	}
	resetFlag(flagNegative)
	if result > 0xFF {
		setFlag(flagCarry)
	} else {
		resetFlag(flagCarry)
	}
	if ((uint(af.GetHigh()) & 0x0F) + (uint(number) & 0x0F) + carry) > 0x0F {
		setFlag(flagHalf)
	} else {
		resetFlag(flagHalf)
	}
	af.SetHigh(uint8(result))
}

func opcodesSUB(number uint8) {
	result := int(af.GetHigh()) - int(number)
	carrybits := int(af.GetHigh()) ^ int(number) ^ result
	af.SetHigh(uint8(result))
	initFlagReg(flagNegative)
	setZeroFlagFromResult(uint8(result))
	if (carrybits & 0x100) != 0 {
		setFlag(flagCarry)
	}
	if (carrybits & 0x10) != 0 {
		setFlag(flagHalf)
	}
}

func opcodesSBC(number uint8) {
	var carry int
	if isSetFlag(flagCarry) {
		carry = 1
	} else {
		carry = 0
	}
	result := int(af.GetHigh()) - int(number) - carry
	initFlagReg(flagNegative)
	setZeroFlagFromResult(uint8(result))
	if result < 0 {
		setFlag(flagCarry)
	}
	if ((int(af.GetHigh()) & 0x0F) - (int(number) & 0x0F) - carry) < 0 {
		setFlag(flagHalf)
	}
	af.SetHigh(uint8(result))
}

func opcodesADDHL(number uint16) {
	result := uint(hl.GetValue()) + uint(number)
	if isSetFlag(flagZero) {
		initFlagReg(flagZero)
	} else {
		clearAllFlags()
	}
	if (result & 0x10000) != 0 {
		setFlag(flagCarry)
	}
	if ((uint(hl.GetValue()) ^ uint(number) ^ (result & 0xFFFF)) & 0x1000) != 0 {
		setFlag(flagHalf)
	}
	hl.SetValue(uint16(result))
}

func opcodesADDSP(number int8) {
	result := int(sp.GetValue()) + int(number)
	clearAllFlags()
	carrybits := int(sp.GetValue()) ^ int(number) ^ (result & 0xFFFF)
	if (carrybits & 0x100) == 0x100 {
		setFlag(flagCarry)
	}
	if (carrybits & 0x10) == 0x10 {
		setFlag(flagHalf)
	}
	sp.SetValue(uint16(result))
}

func opcodesSWAPReg(reg *EightBitReg) {
	lowHalf := reg.GetValue() & 0x0F
	highHalf := (reg.GetValue() >> 4) & 0x0F
	reg.SetValue((lowHalf << 4) + highHalf)
	clearAllFlags()
	setZeroFlagFromResult(reg.GetValue())
}

func opcodesSWAPHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	lowHalf := result & 0x0F
	highHalf := (result >> 4) & 0x0F
	result = (lowHalf << 4) + highHalf
	mem.Write(address, result)
	clearAllFlags()
	setZeroFlagFromResult(result)
}

func opcodesSLA(reg *EightBitReg) {
	if (reg.GetValue() & 0x80) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := reg.GetValue() << 1
	reg.SetValue(result)
	setZeroFlagFromResult(result)
}

func opcodesSLAHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	if (result & 0x80) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result <<= 1
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesSRA(reg *EightBitReg) {
	value := reg.GetValue()
	if (value & 0x01) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := value >> 1
	if (value & 0x80) != 0 {
		result |= 0x80
	}
	reg.SetValue(result)
	setZeroFlagFromResult(result)
}

func opcodesSRAHL() {
	address := hl.GetValue()
	value := mem.Read(address)
	if (value & 0x01) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := value >> 1
	if (value & 0x80) != 0 {
		result |= 0x80
	}
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesSRL(reg *EightBitReg) {
	result := reg.GetValue()
	if (result & 0x01) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result >>= 1
	reg.SetValue(result)
	setZeroFlagFromResult(result)
}

func opcodesSRLHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	if (result & 0x01) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result >>= 1
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesRLC(reg *EightBitReg) {
	opcodesRLCA(reg)
	setZeroFlagFromResult(reg.GetValue())
}

func opcodesRLCA(reg *EightBitReg) {
	value := reg.GetValue()
	result := value << 1
	if (value & 0x80) != 0 {
		initFlagReg(flagCarry)
		result |= 0x01
	} else {
		clearAllFlags()
	}
	reg.SetValue(result)
}

func opcodesRLCHL() {
	address := hl.GetValue()
	value := mem.Read(address)
	result := value << 1
	if (value & 0x80) != 0 {
		initFlagReg(flagCarry)
		result |= 0x01
	} else {
		clearAllFlags()
	}
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesRL(reg *EightBitReg) {
	opcodesRLA(reg)
	setZeroFlagFromResult(reg.GetValue())
}

func opcodesRLA(reg *EightBitReg) {
	var carry uint8
	if isSetFlag(flagCarry) {
		carry = 0x01
	} else {
		carry = 0x00
	}
	value := reg.GetValue()
	if (value & 0x80) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := (value << 1) | carry
	reg.SetValue(result)
}

func opcodesRLHL() {
	var carry uint8
	if isSetFlag(flagCarry) {
		carry = 0x01
	} else {
		carry = 0x00
	}
	address := hl.GetValue()
	value := mem.Read(address)
	if (value & 0x80) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := (value << 1) | carry
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesRRC(reg *EightBitReg) {
	opcodesRRCA(reg)
	setZeroFlagFromResult(reg.GetValue())
}

func opcodesRRCA(reg *EightBitReg) {
	value := reg.GetValue()
	result := value >> 1
	if (value & 0x01) != 0 {
		initFlagReg(flagCarry)
		result |= 0x80
	} else {
		clearAllFlags()
	}
	reg.SetValue(result)
}

func opcodesRRCHL() {
	address := hl.GetValue()
	value := mem.Read(address)
	result := value >> 1
	if (value & 0x01) != 0 {
		initFlagReg(flagCarry)
		result |= 0x80
	} else {
		clearAllFlags()
	}
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesRR(reg *EightBitReg) {
	opcodesRRA(reg)
	setZeroFlagFromResult(reg.GetValue())
}

func opcodesRRA(reg *EightBitReg) {
	var carry uint8
	if isSetFlag(flagCarry) {
		carry = 0x80
	} else {
		carry = 0x00
	}
	value := reg.GetValue()
	if (value & 0x01) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := (value >> 1) | carry
	reg.SetValue(result)
}

func opcodesRRHL() {
	var carry uint8
	if isSetFlag(flagCarry) {
		carry = 0x80
	} else {
		carry = 0x00
	}
	address := hl.GetValue()
	value := mem.Read(address)
	if (value & 0x01) != 0 {
		initFlagReg(flagCarry)
	} else {
		clearAllFlags()
	}
	result := (value >> 1) | carry
	mem.Write(address, result)
	setZeroFlagFromResult(result)
}

func opcodesBIT(reg *EightBitReg, bit uint) {
	if ((reg.GetValue() >> bit) & 0x01) == 0 {
		setFlag(flagZero)
	} else {
		resetFlag(flagZero)
	}
	setFlag(flagHalf)
	resetFlag(flagNegative)
}

func opcodesBITHL(bit uint) {
	if ((mem.Read(hl.GetValue()) >> bit) & 0x01) == 0 {
		setFlag(flagZero)
	} else {
		resetFlag(flagZero)
	}
	setFlag(flagHalf)
	resetFlag(flagNegative)
}

func opcodesSET(reg *EightBitReg, bit uint) {
	reg.SetValue(reg.GetValue() | (0x01 << bit))
}

func opcodesSETHL(bit uint) {
	address := hl.GetValue()
	result := mem.Read(address)
	result |= (0x01 << bit)
	mem.Write(address, result)
}

func opcodesRES(reg *EightBitReg, bit uint) {
	reg.SetValue(reg.GetValue() &^ (0x01 << bit))
}

func opcodesRESHL(bit uint) {
	address := hl.GetValue()
	result := mem.Read(address)
	result &= ^(0x01 << bit)
	mem.Write(address, result)
}
