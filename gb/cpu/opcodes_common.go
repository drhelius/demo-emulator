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

/*
func OPCodes_LD(EightBitRegister* reg1, uint8 reg2)
{
    reg1->SetValue(reg2);
}

func OPCodes_LD(EightBitRegister* reg, u16 address)
{
    reg->SetValue(m_pMemory->Read(address));
}

func OPCodes_LD(u16 address, uint8 reg)
{
    m_pMemory->Write(address, reg);
}

func OPCodes_OR(uint8 number)
{
    uint8 result = af.GetHigh() | number;
    af.SetHigh(result);
    ClearAllFlags();
    ToggleZeroFlagFromResult(result);
}

func OPCodes_XOR(uint8 number)
{
    uint8 result = af.GetHigh() ^ number;
    af.SetHigh(result);
    ClearAllFlags();
    ToggleZeroFlagFromResult(result);
}

func OPCodes_AND(uint8 number)
{
    uint8 result = af.GetHigh() & number;
    af.SetHigh(result);
    SetFlag(FLAG_HALF);
    ToggleZeroFlagFromResult(result);
}

func OPCodes_CP(uint8 number)
{
    SetFlag(FLAG_SUB);
    if (af.GetHigh() < number)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if (af.GetHigh() == number)
    {
        ToggleFlag(FLAG_ZERO);
    }
    if (((af.GetHigh() - number) & 0xF) > (af.GetHigh() & 0xF))
    {
        ToggleFlag(FLAG_HALF);
    }
}

func OPCodes_INC(EightBitRegister* reg)
{
    uint8 result = reg->GetValue() + 1;
    reg->SetValue(result);
    IsSetFlag(FLAG_CARRY) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    ToggleZeroFlagFromResult(result);
    if ((result & 0x0F) == 0x00)
    {
        ToggleFlag(FLAG_HALF);
    }
}

func OPCodes_INC_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue()) + 1;
        return;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    IsSetFlag(FLAG_CARRY) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    ToggleZeroFlagFromResult(m_iReadCache);
    if ((m_iReadCache & 0x0F) == 0x00)
    {
        ToggleFlag(FLAG_HALF);
    }
}

func OPCodes_DEC(EightBitRegister* reg)
{
    uint8 result = reg->GetValue() - 1;
    reg->SetValue(result);
    IsSetFlag(FLAG_CARRY) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    ToggleFlag(FLAG_SUB);
    ToggleZeroFlagFromResult(result);
    if ((result & 0x0F) == 0x0F)
    {
        ToggleFlag(FLAG_HALF);
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
    IsSetFlag(FLAG_CARRY) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    ToggleFlag(FLAG_SUB);
    ToggleZeroFlagFromResult(m_iReadCache);
    if ((m_iReadCache & 0x0F) == 0x0F)
    {
        ToggleFlag(FLAG_HALF);
    }
}

func OPCodes_ADD(uint8 number)
{
    int result = af.GetHigh() + number;
    int carrybits = af.GetHigh() ^ number ^ result;
    af.SetHigh(static_cast<uint8> (result));
    ClearAllFlags();
    ToggleZeroFlagFromResult(static_cast<uint8> (result));
    if ((carrybits & 0x100) != 0)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if ((carrybits & 0x10) != 0)
    {
        ToggleFlag(FLAG_HALF);
    }
}

func OPCodes_ADC(uint8 number)
{
    int carry = IsSetFlag(FLAG_CARRY) ? 1 : 0;
    int result = af.GetHigh() + number + carry;
    ClearAllFlags();
    ToggleZeroFlagFromResult(static_cast<uint8> (result));
    if (result > 0xFF)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if (((af.GetHigh()& 0x0F) + (number & 0x0F) + carry) > 0x0F)
    {
        ToggleFlag(FLAG_HALF);
    }
    af.SetHigh(static_cast<uint8> (result));
}

func OPCodes_SUB(uint8 number)
{
    int result = af.GetHigh() - number;
    int carrybits = af.GetHigh() ^ number ^ result;
    af.SetHigh(static_cast<uint8> (result));
    SetFlag(FLAG_SUB);
    ToggleZeroFlagFromResult(static_cast<uint8> (result));
    if ((carrybits & 0x100) != 0)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if ((carrybits & 0x10) != 0)
    {
        ToggleFlag(FLAG_HALF);
    }
}

func OPCodes_SBC(uint8 number)
{
    int carry = IsSetFlag(FLAG_CARRY) ? 1 : 0;
    int result = af.GetHigh() - number - carry;
    SetFlag(FLAG_SUB);
    ToggleZeroFlagFromResult(static_cast<uint8> (result));
    if (result < 0)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if (((af.GetHigh() & 0x0F) - (number & 0x0F) - carry) < 0)
    {
        ToggleFlag(FLAG_HALF);
    }
    af.SetHigh(static_cast<uint8> (result));
}

func OPCodes_ADD_HL(u16 number)
{
    int result = HL.GetValue() + number;
    IsSetFlag(FLAG_ZERO) ? SetFlag(FLAG_ZERO) : ClearAllFlags();
    if (result & 0x10000)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if ((HL.GetValue() ^ number ^ (result & 0xFFFF)) & 0x1000)
    {
        ToggleFlag(FLAG_HALF);
    }
    HL.SetValue(static_cast<u16> (result));
}

func OPCodes_ADD_sp(s8 number)
{
    int result = sp.GetValue() + number;
    ClearAllFlags();
    if (((sp.GetValue() ^ number ^ (result & 0xFFFF)) & 0x100) == 0x100)
    {
        ToggleFlag(FLAG_CARRY);
    }
    if (((sp.GetValue() ^ number ^ (result & 0xFFFF)) & 0x10) == 0x10)
    {
        ToggleFlag(FLAG_HALF);
    }
    sp.SetValue(static_cast<u16> (result));
}

func OPCodes_SWAP_Register(EightBitRegister* reg)
{
    uint8 low_half = reg->GetValue() & 0x0F;
    uint8 high_half = (reg->GetValue() >> 4) & 0x0F;
    reg->SetValue((low_half << 4) + high_half);
    ClearAllFlags();
    ToggleZeroFlagFromResult(reg->GetValue());
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
    ClearAllFlags();
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_SLA(EightBitRegister* reg)
{
    (reg->GetValue() & 0x80) != 0 ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    uint8 result = reg->GetValue() << 1;
    reg->SetValue(result);
    ToggleZeroFlagFromResult(result);
}

func OPCodes_SLA_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    (m_iReadCache & 0x80) != 0 ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    m_iReadCache <<= 1;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_SRA(EightBitRegister* reg)
{
    uint8 result = reg->GetValue();
    (result & 0x01) != 0 ? SetFlag(FLAG_CARRY) : ClearAllFlags();
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
    ToggleZeroFlagFromResult(result);
}

func OPCodes_SRA_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    (m_iReadCache & 0x01) != 0 ? SetFlag(FLAG_CARRY) : ClearAllFlags();
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
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_SRL(EightBitRegister* reg)
{
    uint8 result = reg->GetValue();
    (result & 0x01) != 0 ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    result >>= 1;
    reg->SetValue(result);
    ToggleZeroFlagFromResult(result);
}

func OPCodes_SRL_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    (m_iReadCache & 0x01) != 0 ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    m_iReadCache >>= 1;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RLC(EightBitRegister* reg, bool isRegisterA)
{
    uint8 result = reg->GetValue();
    if ((result & 0x80) != 0)
    {
        SetFlag(FLAG_CARRY);
        result <<= 1;
        result |= 0x1;
    }
    else
    {
        ClearAllFlags();
        result <<= 1;
    }
    reg->SetValue(result);
    if (!isRegisterA)
    {
        ToggleZeroFlagFromResult(result);
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
        SetFlag(FLAG_CARRY);
        m_iReadCache <<= 1;
        m_iReadCache |= 0x1;
    }
    else
    {
        ClearAllFlags();
        m_iReadCache <<= 1;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RL(EightBitRegister* reg, bool isRegisterA)
{
    uint8 carry = IsSetFlag(FLAG_CARRY) ? 1 : 0;
    uint8 result = reg->GetValue();
    ((result & 0x80) != 0) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    result <<= 1;
    result |= carry;
    reg->SetValue(result);
    if (!isRegisterA)
    {
        ToggleZeroFlagFromResult(result);
    }
}

func OPCodes_RL_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    uint8 carry = IsSetFlag(FLAG_CARRY) ? 1 : 0;
    ((m_iReadCache & 0x80) != 0) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    m_iReadCache <<= 1;
    m_iReadCache |= carry;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RRC(EightBitRegister* reg, bool isRegisterA)
{
    uint8 result = reg->GetValue();
    if ((result & 0x01) != 0)
    {
        SetFlag(FLAG_CARRY);
        result >>= 1;
        result |= 0x80;
    }
    else
    {
        ClearAllFlags();
        result >>= 1;
    }
    reg->SetValue(result);
    if (!isRegisterA)
    {
        ToggleZeroFlagFromResult(result);
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
        SetFlag(FLAG_CARRY);
        m_iReadCache >>= 1;
        m_iReadCache |= 0x80;
    }
    else
    {
        ClearAllFlags();
        m_iReadCache >>= 1;
    }
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_RR(EightBitRegister* reg, bool isRegisterA)
{
    uint8 carry = IsSetFlag(FLAG_CARRY) ? 0x80 : 0x00;
    uint8 result = reg->GetValue();
    ((result & 0x01) != 0) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    result >>= 1;
    result |= carry;
    reg->SetValue(result);
    if (!isRegisterA)
    {
        ToggleZeroFlagFromResult(result);
    }
}

func OPCodes_RR_HL()
{
    if (m_iAccurateOPCodeState == 1)
    {
        m_iReadCache = m_pMemory->Read(HL.GetValue());
        return;
    }
    uint8 carry = IsSetFlag(FLAG_CARRY) ? 0x80 : 0x00;
    ((m_iReadCache & 0x01) != 0) ? SetFlag(FLAG_CARRY) : ClearAllFlags();
    m_iReadCache >>= 1;
    m_iReadCache |= carry;
    m_pMemory->Write(HL.GetValue(), m_iReadCache);
    ToggleZeroFlagFromResult(m_iReadCache);
}

func OPCodes_BIT(EightBitRegister* reg, int bit)
{
    if (((reg->GetValue() >> bit) & 0x01) == 0)
    {
        ToggleFlag(FLAG_ZERO);
    }
    else
    {
        UntoggleFlag(FLAG_ZERO);
    }
    ToggleFlag(FLAG_HALF);
    UntoggleFlag(FLAG_SUB);
}

func OPCodes_BIT_HL(int bit)
{
    if (((m_pMemory->Read(HL.GetValue()) >> bit) & 0x01) == 0)
    {
        ToggleFlag(FLAG_ZERO);
    }
    else
    {
        UntoggleFlag(FLAG_ZERO);
    }
    ToggleFlag(FLAG_HALF);
    UntoggleFlag(FLAG_SUB);
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
