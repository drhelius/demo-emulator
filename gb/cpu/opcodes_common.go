package cpu

import (
	"fmt"

	"github.com/drhelius/demo-emulator/gb/memory"
)

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
	memory.Write(sp.GetValue(), reg.GetHigh())
	sp.Decrement()
	memory.Write(sp.GetValue(), reg.GetLow())
}

func stackPop(reg *SixteenBitReg) {
	reg.SetLow(memory.Read(sp.GetValue()))
	sp.Increment()
	reg.SetHigh(memory.Read(sp.GetValue()))
	sp.Increment()
}

func invalidOPCode() {
	fmt.Println("INVALID OP Code")
}

func opcodesLDFromValue(reg1 *EightBitReg, reg2 uint8) {
	reg1.SetValue(reg2)
}

func opcodesLDFromAddress(reg *EightBitReg, address uint16) {
	reg.SetValue(memory.Read(address))
}

func opcodesLDToMemory(address uint16, reg uint8) {
	memory.Write(address, reg)
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

/*
func OPCodes_INC_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue()) + 1;
        return;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    isSetFlag(flagCarry) ? setFlag(flagCarry) : clearAllFlags();
    toggleZeroFlagFromResult(m_iReadCache);
    if ((m_iReadCache & 0x0F) == 0x00)
    {
        toggleFlag(flagHalf);
    }
}

func OPCodes_DEC(EightBitRegister* reg)
{
    uint8 result = reg->GetValue() - 1;
    reg->SetValue(result);
    isSetFlag(flagCarry) ? setFlag(flagCarry) : clearAllFlags();
    toggleFlag(flagSub);
    toggleZeroFlagFromResult(result);
    if ((result & 0x0F) == 0x0F)
    {
        toggleFlag(flagHalf);
    }
}

func OPCodes_DEC_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue()) - 1;
        return;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    isSetFlag(flagCarry) ? setFlag(flagCarry) : clearAllFlags();
    toggleFlag(flagSub);
    toggleZeroFlagFromResult(m_iReadCache);
    if ((m_iReadCache & 0x0F) == 0x0F)
    {
        toggleFlag(flagHalf);
    }
}

func OPCodes_ADD(uint8 number)
{
    int result = af.GetHigh() + number;
    int carrybits = af.GetHigh() ^ number ^ result;
    af.SetHigh(static_cast<uint8> (result));
    clearAllFlags();
    toggleZeroFlagFromResult(static_cast<uint8> (result));
    if ((carrybits & 0x100) != 0)
    {
        toggleFlag(flagCarry);
    }
    if ((carrybits & 0x10) != 0)
    {
        toggleFlag(flagHalf);
    }
}

func OPCodes_ADC(uint8 number)
{
    int carry = isSetFlag(flagCarry) ? 1 : 0;
    int result = af.GetHigh() + number + carry;
    clearAllFlags();
    toggleZeroFlagFromResult(static_cast<uint8> (result));
    if (result > 0xFF)
    {
        toggleFlag(flagCarry);
    }
    if (((af.GetHigh()& 0x0F) + (number & 0x0F) + carry) > 0x0F)
    {
        toggleFlag(flagHalf);
    }
    af.SetHigh(static_cast<uint8> (result));
}

func OPCodes_SUB(uint8 number)
{
    int result = af.GetHigh() - number;
    int carrybits = af.GetHigh() ^ number ^ result;
    af.SetHigh(static_cast<uint8> (result));
    setFlag(flagSub);
    toggleZeroFlagFromResult(static_cast<uint8> (result));
    if ((carrybits & 0x100) != 0)
    {
        toggleFlag(flagCarry);
    }
    if ((carrybits & 0x10) != 0)
    {
        toggleFlag(flagHalf);
    }
}

func OPCodes_SBC(uint8 number)
{
    int carry = isSetFlag(flagCarry) ? 1 : 0;
    int result = af.GetHigh() - number - carry;
    setFlag(flagSub);
    toggleZeroFlagFromResult(static_cast<uint8> (result));
    if (result < 0)
    {
        toggleFlag(flagCarry);
    }
    if (((af.GetHigh() & 0x0F) - (number & 0x0F) - carry) < 0)
    {
        toggleFlag(flagHalf);
    }
    af.SetHigh(static_cast<uint8> (result));
}

func OPCodes_ADD_HL(u16 number)
{
    int result = HL.GetValue() + number;
    isSetFlag(flagZero) ? setFlag(flagZero) : clearAllFlags();
    if (result & 0x10000)
    {
        toggleFlag(flagCarry);
    }
    if ((HL.GetValue() ^ number ^ (result & 0xFFFF)) & 0x1000)
    {
        toggleFlag(flagHalf);
    }
    HL.SetValue(static_cast<u16> (result));
}

func OPCodes_ADD_sp(s8 number)
{
    int result = sp.GetValue() + number;
    clearAllFlags();
    if (((sp.GetValue() ^ number ^ (result & 0xFFFF)) & 0x100) == 0x100)
    {
        toggleFlag(flagCarry);
    }
    if (((sp.GetValue() ^ number ^ (result & 0xFFFF)) & 0x10) == 0x10)
    {
        toggleFlag(flagHalf);
    }
    sp.SetValue(static_cast<u16> (result));
}

func OPCodes_SWAP_Register(EightBitRegister* reg)
{
    uint8 low_half = reg->GetValue() & 0x0F;
    uint8 high_half = (reg->GetValue() >> 4) & 0x0F;
    reg->SetValue((low_half << 4) + high_half);
    clearAllFlags();
    toggleZeroFlagFromResult(reg->GetValue());
}

func OPCodes_SWAP_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    uint8 low_half = m_iReadCache & 0x0F;
    uint8 high_half = (m_iReadCache >> 4) & 0x0F;
    m_iReadCache = (low_half << 4) + high_half;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    clearAllFlags();
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_SLA(EightBitRegister* reg)
{
    (reg->GetValue() & 0x80) != 0 ? setFlag(flagCarry) : clearAllFlags();
    uint8 result = reg->GetValue() << 1;
    reg->SetValue(result);
    toggleZeroFlagFromResult(result);
}

func OPCodes_SLA_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    (m_iReadCache & 0x80) != 0 ? setFlag(flagCarry) : clearAllFlags();
    m_iReadCache <<= 1;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_SRA(EightBitRegister* reg)
{
    uint8 result = reg->GetValue();
    (result & 0x01) != 0 ? setFlag(flagCarry) : clearAllFlags();
    if ((result & 0x80) != 0)
    {
        result >>= 1;
        result |= 0x80;
    }
    else
    {
        result >>= 1;
    }
    reg->SetValue(result);
    toggleZeroFlagFromResult(result);
}

func OPCodes_SRA_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    (m_iReadCache & 0x01) != 0 ? setFlag(flagCarry) : clearAllFlags();
    if ((m_iReadCache & 0x80) != 0)
    {
        m_iReadCache >>= 1;
        m_iReadCache |= 0x80;
    }
    else
    {
        m_iReadCache >>= 1;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_SRL(EightBitRegister* reg)
{
    uint8 result = reg->GetValue();
    (result & 0x01) != 0 ? setFlag(flagCarry) : clearAllFlags();
    result >>= 1;
    reg->SetValue(result);
    toggleZeroFlagFromResult(result);
}

func OPCodes_SRL_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    (m_iReadCache & 0x01) != 0 ? setFlag(flagCarry) : clearAllFlags();
    m_iReadCache >>= 1;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RLC(EightBitRegister* reg, bool isRegisterA)
{
    uint8 result = reg->GetValue();
    if ((result & 0x80) != 0)
    {
        setFlag(flagCarry);
        result <<= 1;
        result |= 0x1;
    }
    else
    {
        clearAllFlags();
        result <<= 1;
    }
    reg->SetValue(result);
    if (!isRegisterA)
    {
        toggleZeroFlagFromResult(result);
    }
}

func OPCodes_RLC_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    if ((m_iReadCache & 0x80) != 0)
    {
        setFlag(flagCarry);
        m_iReadCache <<= 1;
        m_iReadCache |= 0x1;
    }
    else
    {
        clearAllFlags();
        m_iReadCache <<= 1;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RL(EightBitRegister* reg, bool isRegisterA)
{
    uint8 carry = isSetFlag(flagCarry) ? 1 : 0;
    uint8 result = reg->GetValue();
    ((result & 0x80) != 0) ? setFlag(flagCarry) : clearAllFlags();
    result <<= 1;
    result |= carry;
    reg->SetValue(result);
    if (!isRegisterA)
    {
        toggleZeroFlagFromResult(result);
    }
}

func OPCodes_RL_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    uint8 carry = isSetFlag(flagCarry) ? 1 : 0;
    ((m_iReadCache & 0x80) != 0) ? setFlag(flagCarry) : clearAllFlags();
    m_iReadCache <<= 1;
    m_iReadCache |= carry;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RRC(EightBitRegister* reg, bool isRegisterA)
{
    uint8 result = reg->GetValue();
    if ((result & 0x01) != 0)
    {
        setFlag(flagCarry);
        result >>= 1;
        result |= 0x80;
    }
    else
    {
        clearAllFlags();
        result >>= 1;
    }
    reg->SetValue(result);
    if (!isRegisterA)
    {
        toggleZeroFlagFromResult(result);
    }
}

func OPCodes_RRC_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    if ((m_iReadCache & 0x01) != 0)
    {
        setFlag(flagCarry);
        m_iReadCache >>= 1;
        m_iReadCache |= 0x80;
    }
    else
    {
        clearAllFlags();
        m_iReadCache >>= 1;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RR(EightBitRegister* reg, bool isRegisterA)
{
    uint8 carry = isSetFlag(flagCarry) ? 0x80 : 0x00;
    uint8 result = reg->GetValue();
    ((result & 0x01) != 0) ? setFlag(flagCarry) : clearAllFlags();
    result >>= 1;
    result |= carry;
    reg->SetValue(result);
    if (!isRegisterA)
    {
        toggleZeroFlagFromResult(result);
    }
}

func OPCodes_RR_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    uint8 carry = isSetFlag(flagCarry) ? 0x80 : 0x00;
    ((m_iReadCache & 0x01) != 0) ? setFlag(flagCarry) : clearAllFlags();
    m_iReadCache >>= 1;
    m_iReadCache |= carry;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    toggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_BIT(EightBitRegister* reg, int bit)
{
    if (((reg->GetValue() >> bit) & 0x01) == 0)
    {
        toggleFlag(flagZero);
    }
    else
    {
        UntoggleFlag(flagZero);
    }
    toggleFlag(flagHalf);
    UntoggleFlag(flagSub);
}

func OPCodes_BIT_HL(int bit)
{
    if (((m_pMemory->Read(HL.GetValue()) >> bit) & 0x01) == 0)
    {
        toggleFlag(flagZero);
    }
    else
    {
        UntoggleFlag(flagZero);
    }
    toggleFlag(flagHalf);
    UntoggleFlag(flagSub);
}

func OPCodes_SET(EightBitRegister* reg, int bit)
{
    reg->SetValue(reg->GetValue() | (0x1 << bit));
}

func OPCodes_SET_HL(int bit)
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    m_iReadCache |= (0x1 << bit);
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
}

func OPCodes_RES(EightBitRegister* reg, int bit)
{
    reg->SetValue(reg->GetValue() & (~(0x1 << bit)));
}

func OPCodes_RES_HL(int bit)
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    m_iReadCache &= ~(0x1 << bit);
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
}
*/
