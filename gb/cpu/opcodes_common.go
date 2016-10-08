package cpu

import "fmt"

func clearAllFlags() {
	setFlag(flagNone)
}

func toggleZeroFlagFromResult(result uint8) {
	if result == 0 {
		toggleFlag(flagZero)
	}
}

func setFlag(flag uint8) {
	af.SetLow(flag)
}

func flipFlag(flag uint8) {
	af.SetLow(af.GetLow() ^ flag)
}

func toggleFlag(flag uint8) {
	af.SetLow(af.GetLow() | flag)
}

func untoggleFlag(flag uint8) {
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
	fmt.Println("INVALID OP Code")
}

func opcodesLDFromValue(reg1 *EightBitReg, reg2 uint8) {
	reg1.SetValue(reg2)
}

func opcodesLDFromAddress(reg *EightBitReg, address uint16) {
	reg.SetValue(mem.Read(address))
}

func opcodesLDToMemory(address uint16, reg uint8) {
	mem.Write(address, reg)
}

func opcodesOR(number uint8) {
	result := af.GetHigh() | number
	af.SetHigh(result)
	clearAllFlags()
	toggleZeroFlagFromResult(result)
}

func opcodesXOR(number uint8) {
	result := af.GetHigh() ^ number
	af.SetHigh(result)
	clearAllFlags()
	toggleZeroFlagFromResult(result)
}

func opcodesAND(number uint8) {
	result := af.GetHigh() & number
	af.SetHigh(result)
	setFlag(flagHalf)
	toggleZeroFlagFromResult(result)
}

func opcodesCP(number uint8) {
	setFlag(flagSub)
	if af.GetHigh() < number {
		toggleFlag(flagCarry)
	} else if af.GetHigh() == number {
		toggleFlag(flagZero)
	}
	if ((af.GetHigh() - number) & 0xF) > (af.GetHigh() & 0xF) {
		toggleFlag(flagHalf)
	}
}

func opcodesINC(reg *EightBitReg) {
	result := reg.GetValue() + 1
	reg.SetValue(result)
	if isSetFlag(flagCarry) {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	toggleZeroFlagFromResult(result)
	if (result & 0x0F) == 0x00 {
		toggleFlag(flagHalf)
	}
}

func opcodesINCHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	result++
	mem.Write(address, result)
	if isSetFlag(flagCarry) {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	toggleZeroFlagFromResult(result)
	if (result & 0x0F) == 0x00 {
		toggleFlag(flagHalf)
	}
}

func opcodesDEC(reg *EightBitReg) {
	result := reg.GetValue()
	result--
	reg.SetValue(result)
	if isSetFlag(flagCarry) {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	toggleFlag(flagSub)
	toggleZeroFlagFromResult(result)
	if (result & 0x0F) == 0x0F {
		toggleFlag(flagHalf)
	}
}

func opcodesDECHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	result--
	mem.Write(address, result)
	if isSetFlag(flagCarry) {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	toggleFlag(flagSub)
	toggleZeroFlagFromResult(result)
	if (result & 0x0F) == 0x0F {
		toggleFlag(flagHalf)
	}
}

func opcodesADD(number uint8) {
	result := uint(af.GetHigh()) + uint(number)
	carrybits := uint(af.GetHigh()) ^ uint(number) ^ result
	af.SetHigh(uint8(result))
	clearAllFlags()
	toggleZeroFlagFromResult(uint8(result))
	if (carrybits & 0x100) != 0 {
		toggleFlag(flagCarry)
	}
	if (carrybits & 0x10) != 0 {
		toggleFlag(flagHalf)
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
	clearAllFlags()
	toggleZeroFlagFromResult(uint8(result))
	if result > 0xFF {
		toggleFlag(flagCarry)
	}
	if ((uint(af.GetHigh()) & 0x0F) + (uint(number) & 0x0F) + carry) > 0x0F {
		toggleFlag(flagHalf)
	}
	af.SetHigh(uint8(result))
}

func opcodesSUB(number uint8) {
	result := int(af.GetHigh()) - int(number)
	carrybits := int(af.GetHigh()) ^ int(number) ^ result
	af.SetHigh(uint8(result))
	setFlag(flagSub)
	toggleZeroFlagFromResult(uint8(result))
	if (carrybits & 0x100) != 0 {
		toggleFlag(flagCarry)
	}
	if (carrybits & 0x10) != 0 {
		toggleFlag(flagHalf)
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
	setFlag(flagSub)
	toggleZeroFlagFromResult(uint8(result))
	if result < 0 {
		toggleFlag(flagCarry)
	}
	if ((int(af.GetHigh()) & 0x0F) - (int(number) & 0x0F) - carry) < 0 {
		toggleFlag(flagHalf)
	}
	af.SetHigh(uint8(result))
}

func opcodesADDHL(number uint16) {
	result := uint(hl.GetValue()) + uint(number)
	if isSetFlag(flagZero) {
		setFlag(flagZero)
	} else {
		clearAllFlags()
	}
	if (result & 0x10000) != 0 {
		toggleFlag(flagCarry)
	}
	if ((uint(hl.GetValue()) ^ uint(number) ^ (result & 0xFFFF)) & 0x1000) != 0 {
		toggleFlag(flagHalf)
	}
	hl.SetValue(uint16(result))
}

func opcodesADDSP(number int8) {
	result := int(sp.GetValue()) + int(number)
	clearAllFlags()
	carrybits := int(sp.GetValue()) ^ int(number) ^ (result & 0xFFFF)
	if (carrybits & 0x100) == 0x100 {
		toggleFlag(flagCarry)
	}
	if (carrybits & 0x10) == 0x10 {
		toggleFlag(flagHalf)
	}
	sp.SetValue(uint16(result))
}

func opcodesSWAPReg(reg *EightBitReg) {
	lowHalf := reg.GetValue() & 0x0F
	highHalf := (reg.GetValue() >> 4) & 0x0F
	reg.SetValue((lowHalf << 4) + highHalf)
	clearAllFlags()
	toggleZeroFlagFromResult(reg.GetValue())
}

func opcodesSWAPHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	lowHalf := result & 0x0F
	highHalf := (result >> 4) & 0x0F
	result = (lowHalf << 4) + highHalf
	mem.Write(address, result)
	clearAllFlags()
	toggleZeroFlagFromResult(result)
}

func opcodesSLA(reg *EightBitReg) {
	if (reg.GetValue() & 0x80) != 0 {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result := reg.GetValue() << 1
	reg.SetValue(result)
	toggleZeroFlagFromResult(result)
}

func opcodesSLAHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	if (result & 0x80) != 0 {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result <<= 1
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesSRA(reg *EightBitReg) {
	value := reg.GetValue()
	if (value & 0x01) != 0 {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result := value >> 1
	if (value & 0x80) != 0 {
		result |= 0x80
	}
	reg.SetValue(result)
	toggleZeroFlagFromResult(result)
}

func opcodesSRAHL() {
	address := hl.GetValue()
	value := mem.Read(address)
	if (value & 0x01) != 0 {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result := value >> 1
	if (value & 0x80) != 0 {
		result |= 0x80
	}
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesSRL(reg *EightBitReg) {
	result := reg.GetValue()
	if (result & 0x01) != 0 {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result >>= 1
	reg.SetValue(result)
	toggleZeroFlagFromResult(result)
}

func opcodesSRLHL() {
	address := hl.GetValue()
	result := mem.Read(address)
	if (result & 0x01) != 0 {
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result >>= 1
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesRLC(reg *EightBitReg) {
	opcodesRLCA(reg)
	toggleZeroFlagFromResult(reg.GetValue())
}

func opcodesRLCA(reg *EightBitReg) {
	value := reg.GetValue()
	result := value << 1
	if (value & 0x80) != 0 {
		setFlag(flagCarry)
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
		setFlag(flagCarry)
		result |= 0x01
	} else {
		clearAllFlags()
	}
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesRL(reg *EightBitReg) {
	opcodesRLA(reg)
	toggleZeroFlagFromResult(reg.GetValue())
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
		setFlag(flagCarry)
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
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result := (value << 1) | carry
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesRRC(reg *EightBitReg) {
	opcodesRRCA(reg)
	toggleZeroFlagFromResult(reg.GetValue())
}

func opcodesRRCA(reg *EightBitReg) {
	value := reg.GetValue()
	result := value >> 1
	if (value & 0x01) != 0 {
		setFlag(flagCarry)
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
		setFlag(flagCarry)
		result |= 0x80
	} else {
		clearAllFlags()
	}
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesRR(reg *EightBitReg) {
	opcodesRRA(reg)
	toggleZeroFlagFromResult(reg.GetValue())
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
		setFlag(flagCarry)
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
		setFlag(flagCarry)
	} else {
		clearAllFlags()
	}
	result := (value >> 1) | carry
	mem.Write(address, result)
	toggleZeroFlagFromResult(result)
}

func opcodesBIT(reg *EightBitReg, bit uint) {
	if ((reg.GetValue() >> bit) & 0x01) == 0 {
		toggleFlag(flagZero)
	} else {
		untoggleFlag(flagZero)
	}
	toggleFlag(flagHalf)
	untoggleFlag(flagSub)
}

func opcodesBITHL(bit uint) {
	if ((mem.Read(hl.GetValue()) >> bit) & 0x01) == 0 {
		toggleFlag(flagZero)
	} else {
		untoggleFlag(flagZero)
	}
	toggleFlag(flagHalf)
	untoggleFlag(flagSub)
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
