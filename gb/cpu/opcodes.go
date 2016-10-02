package cpu

/*

func opcode0x00() {
    // NOP
}

func opcode0x01() {
    // LD BC,nn
    OPCodes_LD(BC.GetLowRegister(), PC.GetValue());
    PC.Increment();
    OPCodes_LD(BC.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x02() {
    // LD (BC),A
    OPCodes_LD(BC.GetValue(), AF.GetHigh());
}

func opcode0x03() {
    // INC BC
    BC.Increment();
}

func opcode0x04() {
    // INC B
    OPCodes_INC(BC.GetHighRegister());
}

func opcode0x05() {
    // DEC B
    OPCodes_DEC(BC.GetHighRegister());
}

func opcode0x06() {
    // LD B,n
    OPCodes_LD(BC.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x07() {
    // RLCA
    OPCodes_RLC(AF.GetHighRegister(), true);
}

func opcode0x08() {
    // LD (nn),SP
    u8 l = m_pMemory->Read(PC.GetValue());
    PC.Increment();
    u8 h = m_pMemory->Read(PC.GetValue());
    PC.Increment();
    u16 address = ((h << 8) + l);
    m_pMemory->Write(address, SP.GetLow());
    m_pMemory->Write(address + 1, SP.GetHigh());
}

func opcode0x09() {
    // ADD HL,BC
    OPCodes_ADD_HL(BC.GetValue());
}

func opcode0x0A() {
    // LD A,(BC)
    OPCodes_LD(AF.GetHighRegister(), BC.GetValue());
}

func opcode0x0B() {
    // DEC BC
    BC.Decrement();
}

func opcode0x0C() {
    // INC C
    OPCodes_INC(BC.GetLowRegister());
}

func opcode0x0D() {
    // DEC C
    OPCodes_DEC(BC.GetLowRegister());
}

func opcode0x0E() {
    // LD C,n
    OPCodes_LD(BC.GetLowRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x0F() {
    // RRCA
    OPCodes_RRC(AF.GetHighRegister(), true);
}

func opcode0x10() {
    // STOP
    PC.Increment();

    if (m_bCGB)
    {
        u8 current_key1 = m_pMemory->Retrieve(0xFF4D);

        if (IsSetBit(current_key1, 0))
        {
            m_bCGBSpeed = !m_bCGBSpeed;

            if (m_bCGBSpeed)
            {
                m_iSpeedMultiplier = 1;
                m_pMemory->Load(0xFF4D, 0x80);
            }
            else
            {
                m_iSpeedMultiplier = 0;
                m_pMemory->Load(0xFF4D, 0x00);
            }
        }
    }
}

func opcode0x11() {
    // LD DE,nn
    OPCodes_LD(DE.GetLowRegister(), PC.GetValue());
    PC.Increment();
    OPCodes_LD(DE.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x12() {
    // LD (DE),A
    OPCodes_LD(DE.GetValue(), AF.GetHigh());
}

func opcode0x13() {
    // INC DE
    DE.Increment();
}

func opcode0x14() {
    // INC D
    OPCodes_INC(DE.GetHighRegister());
}

func opcode0x15() {
    // DEC D
    OPCodes_DEC(DE.GetHighRegister());
}

func opcode0x16() {
    // LD D,n
    OPCodes_LD(DE.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x17() {
    // RLA
    OPCodes_RL(AF.GetHighRegister(), true);
}

func opcode0x18() {
    // JR n
    PC.SetValue(PC.GetValue() + 1 + (static_cast<s8> (m_pMemory->Read(PC.GetValue()))));
}

func opcode0x19() {
    // ADD HL,DE
    OPCodes_ADD_HL(DE.GetValue());
}

func opcode0x1A() {
    // LD A,(DE)
    OPCodes_LD(AF.GetHighRegister(), DE.GetValue());
}

func opcode0x1B() {
    // DEC DE
    DE.Decrement();
}

func opcode0x1C() {
    // INC E
    OPCodes_INC(DE.GetLowRegister());
}

func opcode0x1D() {
    // DEC E
    OPCodes_DEC(DE.GetLowRegister());
}

func opcode0x1E() {
    // LD E,n
    OPCodes_LD(DE.GetLowRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x1F() {
    // RRA
    OPCodes_RR(AF.GetHighRegister(), true);
}

func opcode0x20() {
    // JR NZ,n
    if (!IsSetFlag(FLAG_ZERO))
    {
        PC.SetValue(PC.GetValue() + 1 + (static_cast<s8> (m_pMemory->Read(PC.GetValue()))));
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
    }
}

func opcode0x21() {
    // LD HL,nn
    OPCodes_LD(HL.GetLowRegister(), PC.GetValue());
    PC.Increment();
    OPCodes_LD(HL.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x22() {
    // LD (HLI),A
    OPCodes_LD(HL.GetValue(), AF.GetHigh());
    HL.Increment();
}

func opcode0x23() {
    // INC HL
    HL.Increment();
}

func opcode0x24() {
    // INC H
    OPCodes_INC(HL.GetHighRegister());
}

func opcode0x25() {
    // DEC H
    OPCodes_DEC(HL.GetHighRegister());
}

func opcode0x26() {
    // LD H,n
    OPCodes_LD(HL.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x27() {
    // DAA
    int a = AF.GetHigh();

    if (!IsSetFlag(FLAG_SUB))
    {
        if (IsSetFlag(FLAG_HALF) || ((a & 0xF) > 9))
            a += 0x06;

        if (IsSetFlag(FLAG_CARRY) || (a > 0x9F))
            a += 0x60;
    }
    else
    {
        if (IsSetFlag(FLAG_HALF))
            a = (a - 6) & 0xFF;

        if (IsSetFlag(FLAG_CARRY))
            a -= 0x60;
    }

    UntoggleFlag(FLAG_HALF);
    UntoggleFlag(FLAG_ZERO);

    if ((a & 0x100) == 0x100)
        ToggleFlag(FLAG_CARRY);

    a &= 0xFF;

    ToggleZeroFlagFromResult(a);

    AF.SetHigh(a);
}

func opcode0x28() {
    // JR Z,n
    if (IsSetFlag(FLAG_ZERO))
    {
        PC.SetValue(PC.GetValue() + 1 + (static_cast<s8> (m_pMemory->Read(PC.GetValue()))));
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
    }
}

func opcode0x29() {
    // ADD HL,HL
    OPCodes_ADD_HL(HL.GetValue());
}

func opcode0x2A() {
    // LD A,(HLI)
    OPCodes_LD(AF.GetHighRegister(), HL.GetValue());
    HL.Increment();
}

func opcode0x2B() {
    // DEC HL
    HL.Decrement();
}

func opcode0x2C() {
    // INC L
    OPCodes_INC(HL.GetLowRegister());
}

func opcode0x2D() {
    // DEC L
    OPCodes_DEC(HL.GetLowRegister());
}

func opcode0x2E() {
    // LD L,n
    OPCodes_LD(HL.GetLowRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x2F() {
    // CPL
    AF.SetHigh(~AF.GetHigh());
    ToggleFlag(FLAG_HALF);
    ToggleFlag(FLAG_SUB);
}

func opcode0x30() {
    // JR NC,n
    if (!IsSetFlag(FLAG_CARRY))
    {
        PC.SetValue(PC.GetValue() + 1 + (static_cast<s8> (m_pMemory->Read(PC.GetValue()))));
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
    }
}

func opcode0x31() {
    // LD SP,nn
    SP.SetLow(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
    SP.SetHigh(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0x32() {
    // LD (HLD), A
    OPCodes_LD(HL.GetValue(), AF.GetHigh());
    HL.Decrement();
}

func opcode0x33() {
    // INC SP
    SP.Increment();
}

func opcode0x34() {
    // INC (HL)
    OPCodes_INC_HL();
}

func opcode0x35() {
    // DEC (HL)
    OPCodes_DEC_HL();
}

func opcode0x36() {
    // LD (HL),n
    m_pMemory->Write(HL.GetValue(), m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0x37() {
    // SCF
    ToggleFlag(FLAG_CARRY);
    UntoggleFlag(FLAG_HALF);
    UntoggleFlag(FLAG_SUB);
}

func opcode0x38() {
    // JR C,n
    if (IsSetFlag(FLAG_CARRY))
    {
        PC.SetValue(PC.GetValue() + 1 + (static_cast<s8> (m_pMemory->Read(PC.GetValue()))));
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
    }
}

func opcode0x39() {
    // ADD HL,SP
    OPCodes_ADD_HL(SP.GetValue());
}

func opcode0x3A() {
    // LD A,(HLD)
    OPCodes_LD(AF.GetHighRegister(), HL.GetValue());
    HL.Decrement();
}

func opcode0x3B() {
    // DEC SP
    SP.Decrement();
}

func opcode0x3C() {
    // INC A
    OPCodes_INC(AF.GetHighRegister());
}

func opcode0x3D() {
    // DEC A
    OPCodes_DEC(AF.GetHighRegister());

}

func opcode0x3E() {
    // LD A,n
    OPCodes_LD(AF.GetHighRegister(), PC.GetValue());
    PC.Increment();
}

func opcode0x3F() {
    // CCF
    FlipFlag(FLAG_CARRY);
    UntoggleFlag(FLAG_HALF);
    UntoggleFlag(FLAG_SUB);
}

func opcode0x40() {
    // LD B,B
    OPCodes_LD(BC.GetHighRegister(), BC.GetHigh());
}

func opcode0x41() {
    // LD B,C
    OPCodes_LD(BC.GetHighRegister(), BC.GetLow());
}

func opcode0x42() {
    // LD B,D
    OPCodes_LD(BC.GetHighRegister(), DE.GetHigh());
}

func opcode0x43() {
    // LD B,E
    OPCodes_LD(BC.GetHighRegister(), DE.GetLow());
}

func opcode0x44() {
    // LD B,H
    OPCodes_LD(BC.GetHighRegister(), HL.GetHigh());
}

func opcode0x45() {
    // LD B,L
    OPCodes_LD(BC.GetHighRegister(), HL.GetLow());
}

func opcode0x46() {
    // LD B,(HL)
    OPCodes_LD(BC.GetHighRegister(), HL.GetValue());
}

func opcode0x47() {
    // LD B,A
    OPCodes_LD(BC.GetHighRegister(), AF.GetHigh());
}

func opcode0x48() {
    // LD C,B
    OPCodes_LD(BC.GetLowRegister(), BC.GetHigh());
}

func opcode0x49() {
    // LD C,C
    OPCodes_LD(BC.GetLowRegister(), BC.GetLow());
}

func opcode0x4A() {
    // LD C,D
    OPCodes_LD(BC.GetLowRegister(), DE.GetHigh());
}

func opcode0x4B() {
    // LD C,E
    OPCodes_LD(BC.GetLowRegister(), DE.GetLow());
}

func opcode0x4C() {
    // LD C,H
    OPCodes_LD(BC.GetLowRegister(), HL.GetHigh());
}

func opcode0x4D() {
    // LD C,L
    OPCodes_LD(BC.GetLowRegister(), HL.GetLow());
}

func opcode0x4E() {
    // LD C,(HL)
    OPCodes_LD(BC.GetLowRegister(), HL.GetValue());
}

func opcode0x4F() {
    // LD C,A
    OPCodes_LD(BC.GetLowRegister(), AF.GetHigh());
}

func opcode0x50() {
    // LD D,B
    OPCodes_LD(DE.GetHighRegister(), BC.GetHigh());
}

func opcode0x51() {
    // LD D,C
    OPCodes_LD(DE.GetHighRegister(), BC.GetLow());
}

func opcode0x52() {
    // LD D,D
    OPCodes_LD(DE.GetHighRegister(), DE.GetHigh());
}

func opcode0x53() {
    // LD D,E
    OPCodes_LD(DE.GetHighRegister(), DE.GetLow());
}

func opcode0x54() {
    // LD D,H
    OPCodes_LD(DE.GetHighRegister(), HL.GetHigh());
}

func opcode0x55() {
    // LD D,L
    OPCodes_LD(DE.GetHighRegister(), HL.GetLow());
}

func opcode0x56() {
    // LD D,(HL)
    OPCodes_LD(DE.GetHighRegister(), HL.GetValue());
}

func opcode0x57() {
    // LD D,A
    OPCodes_LD(DE.GetHighRegister(), AF.GetHigh());
}

func opcode0x58() {
    // LD E,B
    OPCodes_LD(DE.GetLowRegister(), BC.GetHigh());
}

func opcode0x59() {
    // LD E,C
    OPCodes_LD(DE.GetLowRegister(), BC.GetLow());
}

func opcode0x5A() {
    // LD E,D
    OPCodes_LD(DE.GetLowRegister(), DE.GetHigh());
}

func opcode0x5B() {
    // LD E,E
    OPCodes_LD(DE.GetLowRegister(), DE.GetLow());
}

func opcode0x5C() {
    // LD E,H
    OPCodes_LD(DE.GetLowRegister(), HL.GetHigh());
}

func opcode0x5D() {
    // LD E,L
    OPCodes_LD(DE.GetLowRegister(), HL.GetLow());
}

func opcode0x5E() {
    // LD E,(HL)
    OPCodes_LD(DE.GetLowRegister(), HL.GetValue());
}

func opcode0x5F() {
    // LD E,A
    OPCodes_LD(DE.GetLowRegister(), AF.GetHigh());
}

func opcode0x60() {
    // LD H,B
    OPCodes_LD(HL.GetHighRegister(), BC.GetHigh());
}

func opcode0x61() {
    // LD H,C
    OPCodes_LD(HL.GetHighRegister(), BC.GetLow());
}

func opcode0x62() {
    // LD H,D
    OPCodes_LD(HL.GetHighRegister(), DE.GetHigh());
}

func opcode0x63() {
    // LD H,E
    OPCodes_LD(HL.GetHighRegister(), DE.GetLow());
}

func opcode0x64() {
    // LD H,H
    OPCodes_LD(HL.GetHighRegister(), HL.GetHigh());
}

func opcode0x65() {
    // LD H,L
    OPCodes_LD(HL.GetHighRegister(), HL.GetLow());
}

func opcode0x66() {
    // LD H,(HL)
    OPCodes_LD(HL.GetHighRegister(), HL.GetValue());
}

func opcode0x67() {
    // LD H,A
    OPCodes_LD(HL.GetHighRegister(), AF.GetHigh());
}

func opcode0x68() {
    // LD L,B
    OPCodes_LD(HL.GetLowRegister(), BC.GetHigh());
}

func opcode0x69() {
    // LD L,C
    OPCodes_LD(HL.GetLowRegister(), BC.GetLow());
}

func opcode0x6A() {
    // LD L,D
    OPCodes_LD(HL.GetLowRegister(), DE.GetHigh());
}

func opcode0x6B() {
    // LD L,E
    OPCodes_LD(HL.GetLowRegister(), DE.GetLow());
}

func opcode0x6C() {
    // LD L,H
    OPCodes_LD(HL.GetLowRegister(), HL.GetHigh());
}

func opcode0x6D() {
    // LD L,L
    OPCodes_LD(HL.GetLowRegister(), HL.GetLow());
}

func opcode0x6E() {
    // LD L,(HL)
    OPCodes_LD(HL.GetLowRegister(), HL.GetValue());
}

func opcode0x6F() {
    // LD L,A
    OPCodes_LD(HL.GetLowRegister(), AF.GetHigh());
}

func opcode0x70() {
    // LD (HL),B
    OPCodes_LD(HL.GetValue(), BC.GetHigh());
}

func opcode0x71() {
    // LD (HL),C
    OPCodes_LD(HL.GetValue(), BC.GetLow());
}

func opcode0x72() {
    // LD (HL),D
    OPCodes_LD(HL.GetValue(), DE.GetHigh());
}

func opcode0x73() {
    // LD (HL),E
    OPCodes_LD(HL.GetValue(), DE.GetLow());
}

func opcode0x74() {
    // LD (HL),H
    OPCodes_LD(HL.GetValue(), HL.GetHigh());
}

func opcode0x75() {
    // LD (HL),L
    OPCodes_LD(HL.GetValue(), HL.GetLow());
}

func opcode0x76() {
    // HALT
    halt = true
}

func opcode0x77() {
    // LD (HL),A
    OPCodes_LD(HL.GetValue(), AF.GetHigh());
}

func opcode0x78() {
    // LD A,B
    OPCodes_LD(AF.GetHighRegister(), BC.GetHigh());
}

func opcode0x79() {
    // LD A,C
    OPCodes_LD(AF.GetHighRegister(), BC.GetLow());
}

func opcode0x7A() {
    // LD A,D
    OPCodes_LD(AF.GetHighRegister(), DE.GetHigh());
}

func opcode0x7B() {
    // LD A,E
    OPCodes_LD(AF.GetHighRegister(), DE.GetLow());
}

func opcode0x7C() {
    // LD A,H
    OPCodes_LD(AF.GetHighRegister(), HL.GetHigh());
}

func opcode0x7D() {
    // LD A,L
    OPCodes_LD(AF.GetHighRegister(), HL.GetLow());
}

func opcode0x7E() {
    // LD A,(HL)
    OPCodes_LD(AF.GetHighRegister(), HL.GetValue());
}

func opcode0x7F() {
    // LD A,A
    OPCodes_LD(AF.GetHighRegister(), AF.GetHigh());
}

func opcode0x80() {
    // ADD A,B
    OPCodes_ADD(BC.GetHigh());
}

func opcode0x81() {
    // ADD A,C
    OPCodes_ADD(BC.GetLow());
}

func opcode0x82() {
    // ADD A,D
    OPCodes_ADD(DE.GetHigh());
}

func opcode0x83() {
    // ADD A,E
    OPCodes_ADD(DE.GetLow());
}

func opcode0x84() {
    // ADD A,H
    OPCodes_ADD(HL.GetHigh());
}

func opcode0x85() {
    // ADD A,L
    OPCodes_ADD(HL.GetLow());
}

func opcode0x86() {
    // ADD A,(HL)
    OPCodes_ADD(m_pMemory->Read(HL.GetValue()));
}

func opcode0x87() {
    // ADD A,A
    OPCodes_ADD(AF.GetHigh());
}

func opcode0x88() {
    // ADC A,B
    OPCodes_ADC(BC.GetHigh());
}

func opcode0x89() {
    // ADC A,C
    OPCodes_ADC(BC.GetLow());
}

func opcode0x8A() {
    // ADC A,D
    OPCodes_ADC(DE.GetHigh());
}

func opcode0x8B() {
    // ADC A,E
    OPCodes_ADC(DE.GetLow());
}

func opcode0x8C() {
    // ADC A,H
    OPCodes_ADC(HL.GetHigh());
}

func opcode0x8D() {
    // ADC A,L
    OPCodes_ADC(HL.GetLow());
}

func opcode0x8E() {
    // ADC A,(HL)
    OPCodes_ADC(m_pMemory->Read(HL.GetValue()));
}

func opcode0x8F() {
    // ADC A,A
    OPCodes_ADC(AF.GetHigh());
}

func opcode0x90() {
    // SUB B
    OPCodes_SUB(BC.GetHigh());
}

func opcode0x91() {
    // SUB C
    OPCodes_SUB(BC.GetLow());
}

func opcode0x92() {
    // SUB D
    OPCodes_SUB(DE.GetHigh());
}

func opcode0x93() {
    // SUB E
    OPCodes_SUB(DE.GetLow());
}

func opcode0x94() {
    // SUB H
    OPCodes_SUB(HL.GetHigh());
}

func opcode0x95() {
    // SUB L
    OPCodes_SUB(HL.GetLow());
}

func opcode0x96() {
    // SUB (HL)
    OPCodes_SUB(m_pMemory->Read(HL.GetValue()));
}

func opcode0x97() {
    // SUB A
    OPCodes_SUB(AF.GetHigh());
}

func opcode0x98() {
    // SBC B
    OPCodes_SBC(BC.GetHigh());
}

func opcode0x99() {
    // SBC C
    OPCodes_SBC(BC.GetLow());
}

func opcode0x9A() {
    // SBC D
    OPCodes_SBC(DE.GetHigh());
}

func opcode0x9B() {
    // SBC E
    OPCodes_SBC(DE.GetLow());
}

func opcode0x9C() {
    // SBC H
    OPCodes_SBC(HL.GetHigh());
}

func opcode0x9D() {
    // SBC L
    OPCodes_SBC(HL.GetLow());
}

func opcode0x9E() {
    // SBC (HL)
    OPCodes_SBC(m_pMemory->Read(HL.GetValue()));
}

func opcode0x9F() {
    // SBC A
    OPCodes_SBC(AF.GetHigh());
}

func opcode0xA0() {
    // AND B
    OPCodes_AND(BC.GetHigh());
}

func opcode0xA1() {
    // AND C
    OPCodes_AND(BC.GetLow());
}

func opcode0xA2() {
    // AND D
    OPCodes_AND(DE.GetHigh());
}

func opcode0xA3() {
    // AND E
    OPCodes_AND(DE.GetLow());
}

func opcode0xA4() {
    // AND H
    OPCodes_AND(HL.GetHigh());
}

func opcode0xA5() {
    // AND L
    OPCodes_AND(HL.GetLow());
}

func opcode0xA6() {
    // AND (HL)
    OPCodes_AND(m_pMemory->Read(HL.GetValue()));
}

func opcode0xA7() {
    // AND A
    OPCodes_AND(AF.GetHigh());
}

func opcode0xA8() {
    // XOR B
    OPCodes_XOR(BC.GetHigh());
}

func opcode0xA9() {
    // XOR C
    OPCodes_XOR(BC.GetLow());
}

func opcode0xAA() {
    // XOR D
    OPCodes_XOR(DE.GetHigh());
}

func opcode0xAB() {
    // XOR E
    OPCodes_XOR(DE.GetLow());
}

func opcode0xAC() {
    // XOR H
    OPCodes_XOR(HL.GetHigh());
}

func opcode0xAD() {
    // XOR L
    OPCodes_XOR(HL.GetLow());
}

func opcode0xAE() {
    // XOR (HL)
    OPCodes_XOR(m_pMemory->Read(HL.GetValue()));
}

func opcode0xAF() {
    // XOR A
    OPCodes_XOR(AF.GetHigh());
}

func opcode0xB0() {
    // OR B
    OPCodes_OR(BC.GetHigh());
}

func opcode0xB1() {
    // OR C
    OPCodes_OR(BC.GetLow());
}

func opcode0xB2() {
    // OR D
    OPCodes_OR(DE.GetHigh());
}

func opcode0xB3() {
    // OR E
    OPCodes_OR(DE.GetLow());
}

func opcode0xB4() {
    // OR H
    OPCodes_OR(HL.GetHigh());
}

func opcode0xB5() {
    // OR L
    OPCodes_OR(HL.GetLow());
}

func opcode0xB6() {
    // OR (HL)
    OPCodes_OR(m_pMemory->Read(HL.GetValue()));
}

func opcode0xB7() {
    // OR A
    OPCodes_OR(AF.GetHigh());
}

func opcode0xB8() {
    // CP B
    OPCodes_CP(BC.GetHigh());
}

func opcode0xB9() {
    // CP C
    OPCodes_CP(BC.GetLow());
}

func opcode0xBA() {
    // CP D
    OPCodes_CP(DE.GetHigh());
}

func opcode0xBB() {
    // CP E
    OPCodes_CP(DE.GetLow());
}

func opcode0xBC() {
    // CP H
    OPCodes_CP(HL.GetHigh());
}

func opcode0xBD() {
    // CP L
    OPCodes_CP(HL.GetLow());
}

func opcode0xBE() {
    // CP (HL)
    OPCodes_CP(m_pMemory->Read(HL.GetValue()));
}

func opcode0xBF() {
    // CP A
    OPCodes_CP(AF.GetHigh());
}

func opcode0xC0() {
    // RET NZ
    if (!IsSetFlag(FLAG_ZERO))
    {
        StackPop(&PC);
        m_bBranchTaken = true;
    }
}

func opcode0xC1() {
    // POP BC
    StackPop(&BC);
}

func opcode0xC2() {
    // JP NZ,nn
    if (!IsSetFlag(FLAG_ZERO))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xC3() {
    // JP nn
    u8 l = m_pMemory->Read(PC.GetValue());
    PC.Increment();
    u8 h = m_pMemory->Read(PC.GetValue());
    PC.SetHigh(h);
    PC.SetLow(l);
}

func opcode0xC4() {
    // CALL NZ,nn
    if (!IsSetFlag(FLAG_ZERO))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        StackPush(&PC);
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xC5() {
    // PUSH BC
    StackPush(&BC);
}

func opcode0xC6() {
    // ADD A,n
    OPCodes_ADD(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xC7() {
    // RST 00H
    StackPush(&PC);
    PC.SetValue(0x0000);
}

func opcode0xC8() {
    // RET Z
    if (IsSetFlag(FLAG_ZERO))
    {
        StackPop(&PC);
        m_bBranchTaken = true;
    }
}

func opcode0xC9() {
    // RET
    StackPop(&PC);
}

func opcode0xCA() {
    // JP Z,nn
    if (IsSetFlag(FLAG_ZERO))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xCB() {
    // CB prefixed instruction
}

func opcode0xCC() {
    // CALL Z,nn
    if (IsSetFlag(FLAG_ZERO))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        StackPush(&PC);
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xCD() {
    // CALL nn
    u8 l = m_pMemory->Read(PC.GetValue());
    PC.Increment();
    u8 h = m_pMemory->Read(PC.GetValue());
    PC.Increment();
    StackPush(&PC);
    PC.SetHigh(h);
    PC.SetLow(l);
}

func opcode0xCE() {
    // ADC A,n
    OPCodes_ADC(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xCF() {
    // RST 08H
    StackPush(&PC);
    PC.SetValue(0x0008);
}

func opcode0xD0() {
    // RET NC
    if (!IsSetFlag(FLAG_CARRY))
    {
        StackPop(&PC);
        m_bBranchTaken = true;
    }
}

func opcode0xD1() {
    // POP DE
    StackPop(&DE);
}

func opcode0xD2() {
    // JP NC,nn
    if (!IsSetFlag(FLAG_CARRY))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xD3() {
    InvalidOPCode();
}

func opcode0xD4() {
    // CALL NC,nn
    if (!IsSetFlag(FLAG_CARRY))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        StackPush(&PC);
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xD5() {
    // PUSH DE
    StackPush(&DE);
}

func opcode0xD6() {
    // SUB n
    OPCodes_SUB(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xD7() {
    // RST 10H
    StackPush(&PC);
    PC.SetValue(0x0010);
}

func opcode0xD8() {
    // RET C
    if (IsSetFlag(FLAG_CARRY))
    {
        StackPop(&PC);
        m_bBranchTaken = true;
    }
}

func opcode0xD9() {
    // RETI
    StackPop(&PC);
    m_bIME = true;
}

func opcode0xDA() {
    // JP C,nn
    if (IsSetFlag(FLAG_CARRY))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xDB() {
    InvalidOPCode();
}

func opcode0xDC() {
    // CALL C,nn
    if (IsSetFlag(FLAG_CARRY))
    {
        u8 l = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        u8 h = m_pMemory->Read(PC.GetValue());
        PC.Increment();
        StackPush(&PC);
        PC.SetHigh(h);
        PC.SetLow(l);
        m_bBranchTaken = true;
    }
    else
    {
        PC.Increment();
        PC.Increment();
    }
}

func opcode0xDD() {
    InvalidOPCode();
}

func opcode0xDE() {
    // SBC n
    OPCodes_SBC(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xDF() {
    // RST 18H
    StackPush(&PC);
    PC.SetValue(0x0018);
}

func opcode0xE0() {
    // LD (0xFF00+n),A
    OPCodes_LD(static_cast<u16> (0xFF00 + m_pMemory->Read(PC.GetValue())), AF.GetHigh());
    PC.Increment();
}

func opcode0xE1() {
    // POP HL
    StackPop(&HL);
}

func opcode0xE2() {
    // LD (0xFF00+C),A
    OPCodes_LD(static_cast<u16> (0xFF00 + BC.GetLow()), AF.GetHigh());
}

func opcode0xE3() {
    InvalidOPCode();
}

func opcode0xE4() {
    InvalidOPCode();
}

func opcode0xE5() {
    // PUSH HL
    StackPush(&HL);
}

func opcode0xE6() {
    // AND n
    OPCodes_AND(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xE7() {
    // RST 20H
    StackPush(&PC);
    PC.SetValue(0x0020);
}

func opcode0xE8() {
    // ADD SP,n
    OPCodes_ADD_SP(static_cast<u8> (m_pMemory->Read(PC.GetValue())));
    PC.Increment();
}

func opcode0xE9() {
    // JP (HL)
    PC.SetValue(HL.GetValue());
}

func opcode0xEA() {
    // LD (nn),A
    SixteenBitRegister tmp;
    tmp.SetLow(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
    tmp.SetHigh(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
    OPCodes_LD(tmp.GetValue(), AF.GetHigh());
}

func opcode0xEB() {
    InvalidOPCode();
}

func opcode0xEC() {
    InvalidOPCode();
}

func opcode0xED() {
    InvalidOPCode();
}

func opcode0xEE() {
    // XOR n
    OPCodes_XOR(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xEF() {
    // RST 28H
    StackPush(&PC);
    PC.SetValue(0x28);
}

func opcode0xF0() {
    // LD A,(0xFF00+n)
    OPCodes_LD(AF.GetHighRegister(),
            static_cast<u16> (0xFF00 + m_pMemory->Read(PC.GetValue())));
    PC.Increment();
}

func opcode0xF1() {
    // POP AF
    StackPop(&AF);
    AF.SetLow(AF.GetLow() & 0xF0);
}

func opcode0xF2() {
    // LD A,(C)
    OPCodes_LD(AF.GetHighRegister(), static_cast<u16> (0xFF00 + BC.GetLow()));
}

func opcode0xF3() {
    // DI
    m_bIME = false;
    m_iIMECycles = 0;
}

func opcode0xF4() {
    InvalidOPCode();
}

func opcode0xF5() {
    // PUSH AF
    StackPush(&AF);
}

func opcode0xF6() {
    // OR n
    OPCodes_OR(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xF7() {
    // RST 30H
    StackPush(&PC);
    PC.SetValue(0x0030);
}

func opcode0xF8() {
    // LD HL,SP+n
    s8 n = m_pMemory->Read(PC.GetValue());
    u16 result = SP.GetValue() + n;
    ClearAllFlags();
    if (((SP.GetValue() ^ n ^ result) & 0x100) == 0x100)
        ToggleFlag(FLAG_CARRY);
    if (((SP.GetValue() ^ n ^ result) & 0x10) == 0x10)
        ToggleFlag(FLAG_HALF);
    HL.SetValue(result);
    PC.Increment();
}

func opcode0xF9() {
    // LD SP,HL
    SP.SetValue(HL.GetValue());
}

func opcode0xFA() {
    // LD A,(nn)
    SixteenBitRegister tmp;
    tmp.SetLow(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
    tmp.SetHigh(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
    OPCodes_LD(AF.GetHighRegister(), tmp.GetValue());
}

func opcode0xFB() {
    // EI
    int ei_cycles = kOPCodeMachineCycles[0xFB] * AdjustedCycles(4);
    m_iIMECycles = ei_cycles + 1;
}

func opcode0xFC() {
    InvalidOPCode();
}

func opcode0xFD() {
    InvalidOPCode();
}

func opcode0xFE() {
    // CP n
    OPCodes_CP(m_pMemory->Read(PC.GetValue()));
    PC.Increment();
}

func opcode0xFF() {
    // RST 38H
    StackPush(&PC);
    PC.SetValue(0x0038);
}
*/
