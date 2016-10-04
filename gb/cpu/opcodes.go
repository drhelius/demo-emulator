package cpu

/*

func opcode0x00() {
    // NOP
}

func opcode0x01() {
    // LD BC,nn
    OPCodes_LD(bc.GetLowReg(), pc.GetValue());
    pc.Increment();
    OPCodes_LD(bc.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x02() {
    // LD (BC),A
    OPCodes_LD(bc.GetValue(), af.GetHigh());
}

func opcode0x03() {
    // INC BC
    bc.Increment();
}

func opcode0x04() {
    // INC B
    OPCodes_INC(bc.GetHighRegister());
}

func opcode0x05() {
    // DEC B
    OPCodes_DEC(bc.GetHighRegister());
}

func opcode0x06() {
    // LD B,n
    OPCodes_LD(bc.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x07() {
    // RLCA
    OPCodes_RLC(af.GetHighRegister(), true);
}

func opcode0x08() {
    // LD (nn),SP
    u8 l = memory.Read(pc.GetValue());
    pc.Increment();
    u8 h = memory.Read(pc.GetValue());
    pc.Increment();
    u16 address = ((h << 8) + l);
    memory.Write(address, sp.GetLow());
    memory.Write(address + 1, sp.GetHigh());
}

func opcode0x09() {
    // ADD HL,BC
    OPCodes_ADD_HL(bc.GetValue());
}

func opcode0x0A() {
    // LD A,(BC)
    OPCodes_LD(af.GetHighRegister(), bc.GetValue());
}

func opcode0x0B() {
    // DEC BC
    bc.Decrement();
}

func opcode0x0C() {
    // INC C
    OPCodes_INC(bc.GetLowReg());
}

func opcode0x0D() {
    // DEC C
    OPCodes_DEC(bc.GetLowReg());
}

func opcode0x0E() {
    // LD C,n
    OPCodes_LD(bc.GetLowReg(), pc.GetValue());
    pc.Increment();
}

func opcode0x0F() {
    // RRCA
    OPCodes_RRC(af.GetHighRegister(), true);
}

func opcode0x10() {
    // STOP
    pc.Increment();

    if (m_bCGB)
    {
        u8 current_key1 = memory.Read(0xFF4D);

        if (IsSetBit(current_key1, 0))
        {
            m_bCGBSpeed = !m_bCGBSpeed;

            if (m_bCGBSpeed)
            {
                m_iSpeedMultiplier = 1;
                memory.Write(0xFF4D, 0x80);
            }
            else
            {
                m_iSpeedMultiplier = 0;
                memory.Write(0xFF4D, 0x00);
            }
        }
    }
}

func opcode0x11() {
    // LD DE,nn
    OPCodes_LD(de.GetLowReg(), pc.GetValue());
    pc.Increment();
    OPCodes_LD(de.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x12() {
    // LD (DE),A
    OPCodes_LD(de.GetValue(), af.GetHigh());
}

func opcode0x13() {
    // INC DE
    de.Increment();
}

func opcode0x14() {
    // INC D
    OPCodes_INC(de.GetHighRegister());
}

func opcode0x15() {
    // DEC D
    OPCodes_DEC(de.GetHighRegister());
}

func opcode0x16() {
    // LD D,n
    OPCodes_LD(de.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x17() {
    // RLA
    OPCodes_RL(af.GetHighRegister(), true);
}

func opcode0x18() {
    // JR n
    pc.SetValue(pc.GetValue() + 1 + (static_cast<s8> (memory.Read(pc.GetValue()))));
}

func opcode0x19() {
    // ADD HL,DE
    OPCodes_ADD_HL(de.GetValue());
}

func opcode0x1A() {
    // LD A,(DE)
    OPCodes_LD(af.GetHighRegister(), de.GetValue());
}

func opcode0x1B() {
    // DEC DE
    de.Decrement();
}

func opcode0x1C() {
    // INC E
    OPCodes_INC(de.GetLowReg());
}

func opcode0x1D() {
    // DEC E
    OPCodes_DEC(de.GetLowReg());
}

func opcode0x1E() {
    // LD E,n
    OPCodes_LD(de.GetLowReg(), pc.GetValue());
    pc.Increment();
}

func opcode0x1F() {
    // RRA
    OPCodes_RR(af.GetHighRegister(), true);
}

func opcode0x20() {
    // JR NZ,n
    if (!isSetFlag(flagZero))
    {
        pc.SetValue(pc.GetValue() + 1 + (static_cast<s8> (memory.Read(pc.GetValue()))));
        branchTaken = true;
    }
    else
    {
        pc.Increment();
    }
}

func opcode0x21() {
    // LD HL,nn
    OPCodes_LD(hl.GetLowReg(), pc.GetValue());
    pc.Increment();
    OPCodes_LD(hl.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x22() {
    // LD (HLI),A
    OPCodes_LD(hl.GetValue(), af.GetHigh());
    hl.Increment();
}

func opcode0x23() {
    // INC HL
    hl.Increment();
}

func opcode0x24() {
    // INC H
    OPCodes_INC(hl.GetHighRegister());
}

func opcode0x25() {
    // DEC H
    OPCodes_DEC(hl.GetHighRegister());
}

func opcode0x26() {
    // LD H,n
    OPCodes_LD(hl.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x27() {
    // DAA
    int a = af.GetHigh();

    if (!isSetFlag(flagSub))
    {
        if (isSetFlag(flagHalf) || ((a & 0xF) > 9))
            a += 0x06;

        if (isSetFlag(flagCarry) || (a > 0x9F))
            a += 0x60;
    }
    else
    {
        if (isSetFlag(flagHalf))
            a = (a - 6) & 0xFF;

        if (isSetFlag(flagCarry))
            a -= 0x60;
    }

    untoggleFlag(flagHalf);
    untoggleFlag(flagZero);

    if ((a & 0x100) == 0x100)
        toggleFlag(flagCarry);

    a &= 0xFF;

    ToggleZeroFlagFromResult(a);

    af.SetHigh(a);
}

func opcode0x28() {
    // JR Z,n
    if (isSetFlag(flagZero))
    {
        pc.SetValue(pc.GetValue() + 1 + (static_cast<s8> (memory.Read(pc.GetValue()))));
        branchTaken = true;
    }
    else
    {
        pc.Increment();
    }
}

func opcode0x29() {
    // ADD HL,HL
    OPCodes_ADD_HL(hl.GetValue());
}

func opcode0x2A() {
    // LD A,(HLI)
    OPCodes_LD(af.GetHighRegister(), hl.GetValue());
    hl.Increment();
}

func opcode0x2B() {
    // DEC HL
    hl.Decrement();
}

func opcode0x2C() {
    // INC L
    OPCodes_INC(hl.GetLowReg());
}

func opcode0x2D() {
    // DEC L
    OPCodes_DEC(hl.GetLowReg());
}

func opcode0x2E() {
    // LD L,n
    OPCodes_LD(hl.GetLowReg(), pc.GetValue());
    pc.Increment();
}

func opcode0x2F() {
    // CPL
    af.SetHigh(~af.GetHigh());
    toggleFlag(flagHalf);
    toggleFlag(flagSub);
}

func opcode0x30() {
    // JR NC,n
    if (!isSetFlag(flagCarry))
    {
        pc.SetValue(pc.GetValue() + 1 + (static_cast<s8> (memory.Read(pc.GetValue()))));
        branchTaken = true;
    }
    else
    {
        pc.Increment();
    }
}

func opcode0x31() {
    // LD SP,nn
    sp.SetLow(memory.Read(pc.GetValue()));
    pc.Increment();
    sp.SetHigh(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0x32() {
    // LD (HLD), A
    OPCodes_LD(hl.GetValue(), af.GetHigh());
    hl.Decrement();
}

func opcode0x33() {
    // INC SP
    sp.Increment();
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
    memory.Write(hl.GetValue(), memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0x37() {
    // SCF
    toggleFlag(flagCarry);
    untoggleFlag(flagHalf);
    untoggleFlag(flagSub);
}

func opcode0x38() {
    // JR C,n
    if (isSetFlag(flagCarry))
    {
        pc.SetValue(pc.GetValue() + 1 + (static_cast<s8> (memory.Read(pc.GetValue()))));
        branchTaken = true;
    }
    else
    {
        pc.Increment();
    }
}

func opcode0x39() {
    // ADD HL,SP
    OPCodes_ADD_HL(sp.GetValue());
}

func opcode0x3A() {
    // LD A,(HLD)
    OPCodes_LD(af.GetHighRegister(), hl.GetValue());
    hl.Decrement();
}

func opcode0x3B() {
    // DEC SP
    sp.Decrement();
}

func opcode0x3C() {
    // INC A
    OPCodes_INC(af.GetHighRegister());
}

func opcode0x3D() {
    // DEC A
    OPCodes_DEC(af.GetHighRegister());

}

func opcode0x3E() {
    // LD A,n
    OPCodes_LD(af.GetHighRegister(), pc.GetValue());
    pc.Increment();
}

func opcode0x3F() {
    // CCF
    flipFlag(flagCarry);
    untoggleFlag(flagHalf);
    untoggleFlag(flagSub);
}

func opcode0x40() {
    // LD B,B
    OPCodes_LD(bc.GetHighRegister(), bc.GetHigh());
}

func opcode0x41() {
    // LD B,C
    OPCodes_LD(bc.GetHighRegister(), bc.GetLow());
}

func opcode0x42() {
    // LD B,D
    OPCodes_LD(bc.GetHighRegister(), de.GetHigh());
}

func opcode0x43() {
    // LD B,E
    OPCodes_LD(bc.GetHighRegister(), de.GetLow());
}

func opcode0x44() {
    // LD B,H
    OPCodes_LD(bc.GetHighRegister(), hl.GetHigh());
}

func opcode0x45() {
    // LD B,L
    OPCodes_LD(bc.GetHighRegister(), hl.GetLow());
}

func opcode0x46() {
    // LD B,(HL)
    OPCodes_LD(bc.GetHighRegister(), hl.GetValue());
}

func opcode0x47() {
    // LD B,A
    OPCodes_LD(bc.GetHighRegister(), af.GetHigh());
}

func opcode0x48() {
    // LD C,B
    OPCodes_LD(bc.GetLowReg(), bc.GetHigh());
}

func opcode0x49() {
    // LD C,C
    OPCodes_LD(bc.GetLowReg(), bc.GetLow());
}

func opcode0x4A() {
    // LD C,D
    OPCodes_LD(bc.GetLowReg(), de.GetHigh());
}

func opcode0x4B() {
    // LD C,E
    OPCodes_LD(bc.GetLowReg(), de.GetLow());
}

func opcode0x4C() {
    // LD C,H
    OPCodes_LD(bc.GetLowReg(), hl.GetHigh());
}

func opcode0x4D() {
    // LD C,L
    OPCodes_LD(bc.GetLowReg(), hl.GetLow());
}

func opcode0x4E() {
    // LD C,(HL)
    OPCodes_LD(bc.GetLowReg(), hl.GetValue());
}

func opcode0x4F() {
    // LD C,A
    OPCodes_LD(bc.GetLowReg(), af.GetHigh());
}

func opcode0x50() {
    // LD D,B
    OPCodes_LD(de.GetHighRegister(), bc.GetHigh());
}

func opcode0x51() {
    // LD D,C
    OPCodes_LD(de.GetHighRegister(), bc.GetLow());
}

func opcode0x52() {
    // LD D,D
    OPCodes_LD(de.GetHighRegister(), de.GetHigh());
}

func opcode0x53() {
    // LD D,E
    OPCodes_LD(de.GetHighRegister(), de.GetLow());
}

func opcode0x54() {
    // LD D,H
    OPCodes_LD(de.GetHighRegister(), hl.GetHigh());
}

func opcode0x55() {
    // LD D,L
    OPCodes_LD(de.GetHighRegister(), hl.GetLow());
}

func opcode0x56() {
    // LD D,(HL)
    OPCodes_LD(de.GetHighRegister(), hl.GetValue());
}

func opcode0x57() {
    // LD D,A
    OPCodes_LD(de.GetHighRegister(), af.GetHigh());
}

func opcode0x58() {
    // LD E,B
    OPCodes_LD(de.GetLowReg(), bc.GetHigh());
}

func opcode0x59() {
    // LD E,C
    OPCodes_LD(de.GetLowReg(), bc.GetLow());
}

func opcode0x5A() {
    // LD E,D
    OPCodes_LD(de.GetLowReg(), de.GetHigh());
}

func opcode0x5B() {
    // LD E,E
    OPCodes_LD(de.GetLowReg(), de.GetLow());
}

func opcode0x5C() {
    // LD E,H
    OPCodes_LD(de.GetLowReg(), hl.GetHigh());
}

func opcode0x5D() {
    // LD E,L
    OPCodes_LD(de.GetLowReg(), hl.GetLow());
}

func opcode0x5E() {
    // LD E,(HL)
    OPCodes_LD(de.GetLowReg(), hl.GetValue());
}

func opcode0x5F() {
    // LD E,A
    OPCodes_LD(de.GetLowReg(), af.GetHigh());
}

func opcode0x60() {
    // LD H,B
    OPCodes_LD(hl.GetHighRegister(), bc.GetHigh());
}

func opcode0x61() {
    // LD H,C
    OPCodes_LD(hl.GetHighRegister(), bc.GetLow());
}

func opcode0x62() {
    // LD H,D
    OPCodes_LD(hl.GetHighRegister(), de.GetHigh());
}

func opcode0x63() {
    // LD H,E
    OPCodes_LD(hl.GetHighRegister(), de.GetLow());
}

func opcode0x64() {
    // LD H,H
    OPCodes_LD(hl.GetHighRegister(), hl.GetHigh());
}

func opcode0x65() {
    // LD H,L
    OPCodes_LD(hl.GetHighRegister(), hl.GetLow());
}

func opcode0x66() {
    // LD H,(HL)
    OPCodes_LD(hl.GetHighRegister(), hl.GetValue());
}

func opcode0x67() {
    // LD H,A
    OPCodes_LD(hl.GetHighRegister(), af.GetHigh());
}

func opcode0x68() {
    // LD L,B
    OPCodes_LD(hl.GetLowReg(), bc.GetHigh());
}

func opcode0x69() {
    // LD L,C
    OPCodes_LD(hl.GetLowReg(), bc.GetLow());
}

func opcode0x6A() {
    // LD L,D
    OPCodes_LD(hl.GetLowReg(), de.GetHigh());
}

func opcode0x6B() {
    // LD L,E
    OPCodes_LD(hl.GetLowReg(), de.GetLow());
}

func opcode0x6C() {
    // LD L,H
    OPCodes_LD(hl.GetLowReg(), hl.GetHigh());
}

func opcode0x6D() {
    // LD L,L
    OPCodes_LD(hl.GetLowReg(), hl.GetLow());
}

func opcode0x6E() {
    // LD L,(HL)
    OPCodes_LD(hl.GetLowReg(), hl.GetValue());
}

func opcode0x6F() {
    // LD L,A
    OPCodes_LD(hl.GetLowReg(), af.GetHigh());
}

func opcode0x70() {
    // LD (HL),B
    OPCodes_LD(hl.GetValue(), bc.GetHigh());
}

func opcode0x71() {
    // LD (HL),C
    OPCodes_LD(hl.GetValue(), bc.GetLow());
}

func opcode0x72() {
    // LD (HL),D
    OPCodes_LD(hl.GetValue(), de.GetHigh());
}

func opcode0x73() {
    // LD (HL),E
    OPCodes_LD(hl.GetValue(), de.GetLow());
}

func opcode0x74() {
    // LD (HL),H
    OPCodes_LD(hl.GetValue(), hl.GetHigh());
}

func opcode0x75() {
    // LD (HL),L
    OPCodes_LD(hl.GetValue(), hl.GetLow());
}

func opcode0x76() {
    // HALT
    halt = true
}

func opcode0x77() {
    // LD (HL),A
    OPCodes_LD(hl.GetValue(), af.GetHigh());
}

func opcode0x78() {
    // LD A,B
    OPCodes_LD(af.GetHighRegister(), bc.GetHigh());
}

func opcode0x79() {
    // LD A,C
    OPCodes_LD(af.GetHighRegister(), bc.GetLow());
}

func opcode0x7A() {
    // LD A,D
    OPCodes_LD(af.GetHighRegister(), de.GetHigh());
}

func opcode0x7B() {
    // LD A,E
    OPCodes_LD(af.GetHighRegister(), de.GetLow());
}

func opcode0x7C() {
    // LD A,H
    OPCodes_LD(af.GetHighRegister(), hl.GetHigh());
}

func opcode0x7D() {
    // LD A,L
    OPCodes_LD(af.GetHighRegister(), hl.GetLow());
}

func opcode0x7E() {
    // LD A,(HL)
    OPCodes_LD(af.GetHighRegister(), hl.GetValue());
}

func opcode0x7F() {
    // LD A,A
    OPCodes_LD(af.GetHighRegister(), af.GetHigh());
}

func opcode0x80() {
    // ADD A,B
    OPCodes_ADD(bc.GetHigh());
}

func opcode0x81() {
    // ADD A,C
    OPCodes_ADD(bc.GetLow());
}

func opcode0x82() {
    // ADD A,D
    OPCodes_ADD(de.GetHigh());
}

func opcode0x83() {
    // ADD A,E
    OPCodes_ADD(de.GetLow());
}

func opcode0x84() {
    // ADD A,H
    OPCodes_ADD(hl.GetHigh());
}

func opcode0x85() {
    // ADD A,L
    OPCodes_ADD(hl.GetLow());
}

func opcode0x86() {
    // ADD A,(HL)
    OPCodes_ADD(memory.Read(hl.GetValue()));
}

func opcode0x87() {
    // ADD A,A
    OPCodes_ADD(af.GetHigh());
}

func opcode0x88() {
    // ADC A,B
    OPCodes_ADC(bc.GetHigh());
}

func opcode0x89() {
    // ADC A,C
    OPCodes_ADC(bc.GetLow());
}

func opcode0x8A() {
    // ADC A,D
    OPCodes_ADC(de.GetHigh());
}

func opcode0x8B() {
    // ADC A,E
    OPCodes_ADC(de.GetLow());
}

func opcode0x8C() {
    // ADC A,H
    OPCodes_ADC(hl.GetHigh());
}

func opcode0x8D() {
    // ADC A,L
    OPCodes_ADC(hl.GetLow());
}

func opcode0x8E() {
    // ADC A,(HL)
    OPCodes_ADC(memory.Read(hl.GetValue()));
}

func opcode0x8F() {
    // ADC A,A
    OPCodes_ADC(af.GetHigh());
}

func opcode0x90() {
    // SUB B
    OPCodes_SUB(bc.GetHigh());
}

func opcode0x91() {
    // SUB C
    OPCodes_SUB(bc.GetLow());
}

func opcode0x92() {
    // SUB D
    OPCodes_SUB(de.GetHigh());
}

func opcode0x93() {
    // SUB E
    OPCodes_SUB(de.GetLow());
}

func opcode0x94() {
    // SUB H
    OPCodes_SUB(hl.GetHigh());
}

func opcode0x95() {
    // SUB L
    OPCodes_SUB(hl.GetLow());
}

func opcode0x96() {
    // SUB (HL)
    OPCodes_SUB(memory.Read(hl.GetValue()));
}

func opcode0x97() {
    // SUB A
    OPCodes_SUB(af.GetHigh());
}

func opcode0x98() {
    // SBC B
    OPCodes_SBC(bc.GetHigh());
}

func opcode0x99() {
    // SBC C
    OPCodes_SBC(bc.GetLow());
}

func opcode0x9A() {
    // SBC D
    OPCodes_SBC(de.GetHigh());
}

func opcode0x9B() {
    // SBC E
    OPCodes_SBC(de.GetLow());
}

func opcode0x9C() {
    // SBC H
    OPCodes_SBC(hl.GetHigh());
}

func opcode0x9D() {
    // SBC L
    OPCodes_SBC(hl.GetLow());
}

func opcode0x9E() {
    // SBC (HL)
    OPCodes_SBC(memory.Read(hl.GetValue()));
}

func opcode0x9F() {
    // SBC A
    OPCodes_SBC(af.GetHigh());
}

func opcode0xA0() {
    // AND B
    OPCodes_AND(bc.GetHigh());
}

func opcode0xA1() {
    // AND C
    OPCodes_AND(bc.GetLow());
}

func opcode0xA2() {
    // AND D
    OPCodes_AND(de.GetHigh());
}

func opcode0xA3() {
    // AND E
    OPCodes_AND(de.GetLow());
}

func opcode0xA4() {
    // AND H
    OPCodes_AND(hl.GetHigh());
}

func opcode0xA5() {
    // AND L
    OPCodes_AND(hl.GetLow());
}

func opcode0xA6() {
    // AND (HL)
    OPCodes_AND(memory.Read(hl.GetValue()));
}

func opcode0xA7() {
    // AND A
    OPCodes_AND(af.GetHigh());
}

func opcode0xA8() {
    // XOR B
    OPCodes_XOR(bc.GetHigh());
}

func opcode0xA9() {
    // XOR C
    OPCodes_XOR(bc.GetLow());
}

func opcode0xAA() {
    // XOR D
    OPCodes_XOR(de.GetHigh());
}

func opcode0xAB() {
    // XOR E
    OPCodes_XOR(de.GetLow());
}

func opcode0xAC() {
    // XOR H
    OPCodes_XOR(hl.GetHigh());
}

func opcode0xAD() {
    // XOR L
    OPCodes_XOR(hl.GetLow());
}

func opcode0xAE() {
    // XOR (HL)
    OPCodes_XOR(memory.Read(hl.GetValue()));
}

func opcode0xAF() {
    // XOR A
    OPCodes_XOR(af.GetHigh());
}

func opcode0xB0() {
    // OR B
    OPCodes_OR(bc.GetHigh());
}

func opcode0xB1() {
    // OR C
    OPCodes_OR(bc.GetLow());
}

func opcode0xB2() {
    // OR D
    OPCodes_OR(de.GetHigh());
}

func opcode0xB3() {
    // OR E
    OPCodes_OR(de.GetLow());
}

func opcode0xB4() {
    // OR H
    OPCodes_OR(hl.GetHigh());
}

func opcode0xB5() {
    // OR L
    OPCodes_OR(hl.GetLow());
}

func opcode0xB6() {
    // OR (HL)
    OPCodes_OR(memory.Read(hl.GetValue()));
}

func opcode0xB7() {
    // OR A
    OPCodes_OR(af.GetHigh());
}

func opcode0xB8() {
    // CP B
    OPCodes_CP(bc.GetHigh());
}

func opcode0xB9() {
    // CP C
    OPCodes_CP(bc.GetLow());
}

func opcode0xBA() {
    // CP D
    OPCodes_CP(de.GetHigh());
}

func opcode0xBB() {
    // CP E
    OPCodes_CP(de.GetLow());
}

func opcode0xBC() {
    // CP H
    OPCodes_CP(hl.GetHigh());
}

func opcode0xBD() {
    // CP L
    OPCodes_CP(hl.GetLow());
}

func opcode0xBE() {
    // CP (HL)
    OPCodes_CP(memory.Read(hl.GetValue()));
}

func opcode0xBF() {
    // CP A
    OPCodes_CP(af.GetHigh());
}

func opcode0xC0() {
    // RET NZ
    if (!isSetFlag(flagZero))
    {
        stackPop(&PC);
        branchTaken = true;
    }
}

func opcode0xC1() {
    // POP BC
    stackPop(&BC);
}

func opcode0xC2() {
    // JP NZ,nn
    if (!isSetFlag(flagZero))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xC3() {
    // JP nn
    u8 l = memory.Read(pc.GetValue());
    pc.Increment();
    u8 h = memory.Read(pc.GetValue());
    pc.SetHigh(h);
    pc.SetLow(l);
}

func opcode0xC4() {
    // CALL NZ,nn
    if (!isSetFlag(flagZero))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.Increment();
        stackPush(&PC);
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xC5() {
    // PUSH BC
    stackPush(&BC);
}

func opcode0xC6() {
    // ADD A,n
    OPCodes_ADD(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xC7() {
    // RST 00H
    stackPush(&PC);
    pc.SetValue(0x0000);
}

func opcode0xC8() {
    // RET Z
    if (isSetFlag(flagZero))
    {
        stackPop(&PC);
        branchTaken = true;
    }
}

func opcode0xC9() {
    // RET
    stackPop(&PC);
}

func opcode0xCA() {
    // JP Z,nn
    if (isSetFlag(flagZero))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xCB() {
    // CB prefixed instruction
}

func opcode0xCC() {
    // CALL Z,nn
    if (isSetFlag(flagZero))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.Increment();
        stackPush(&PC);
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xCD() {
    // CALL nn
    u8 l = memory.Read(pc.GetValue());
    pc.Increment();
    u8 h = memory.Read(pc.GetValue());
    pc.Increment();
    stackPush(&PC);
    pc.SetHigh(h);
    pc.SetLow(l);
}

func opcode0xCE() {
    // ADC A,n
    OPCodes_ADC(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xCF() {
    // RST 08H
    stackPush(&PC);
    pc.SetValue(0x0008);
}

func opcode0xD0() {
    // RET NC
    if (!isSetFlag(flagCarry))
    {
        stackPop(&PC);
        branchTaken = true;
    }
}

func opcode0xD1() {
    // POP DE
    stackPop(&DE);
}

func opcode0xD2() {
    // JP NC,nn
    if (!isSetFlag(flagCarry))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xD3() {
    invalidOPCode();
}

func opcode0xD4() {
    // CALL NC,nn
    if (!isSetFlag(flagCarry))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.Increment();
        stackPush(&PC);
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xD5() {
    // PUSH DE
    stackPush(&DE);
}

func opcode0xD6() {
    // SUB n
    OPCodes_SUB(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xD7() {
    // RST 10H
    stackPush(&PC);
    pc.SetValue(0x0010);
}

func opcode0xD8() {
    // RET C
    if (isSetFlag(flagCarry))
    {
        stackPop(&PC);
        branchTaken = true;
    }
}

func opcode0xD9() {
    // RETI
    stackPop(&PC);
    ime = true;
}

func opcode0xDA() {
    // JP C,nn
    if (isSetFlag(flagCarry))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xDB() {
    invalidOPCode();
}

func opcode0xDC() {
    // CALL C,nn
    if (isSetFlag(flagCarry))
    {
        u8 l = memory.Read(pc.GetValue());
        pc.Increment();
        u8 h = memory.Read(pc.GetValue());
        pc.Increment();
        stackPush(&PC);
        pc.SetHigh(h);
        pc.SetLow(l);
        branchTaken = true;
    }
    else
    {
        pc.Increment();
        pc.Increment();
    }
}

func opcode0xDD() {
    invalidOPCode();
}

func opcode0xDE() {
    // SBC n
    OPCodes_SBC(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xDF() {
    // RST 18H
    stackPush(&PC);
    pc.SetValue(0x0018);
}

func opcode0xE0() {
    // LD (0xFF00+n),A
    OPCodes_LD(static_cast<u16> (0xFF00 + memory.Read(pc.GetValue())), af.GetHigh());
    pc.Increment();
}

func opcode0xE1() {
    // POP HL
    stackPop(&HL);
}

func opcode0xE2() {
    // LD (0xFF00+C),A
    OPCodes_LD(static_cast<u16> (0xFF00 + bc.GetLow()), af.GetHigh());
}

func opcode0xE3() {
    invalidOPCode();
}

func opcode0xE4() {
    invalidOPCode();
}

func opcode0xE5() {
    // PUSH HL
    stackPush(&HL);
}

func opcode0xE6() {
    // AND n
    OPCodes_AND(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xE7() {
    // RST 20H
    stackPush(&PC);
    pc.SetValue(0x0020);
}

func opcode0xE8() {
    // ADD SP,n
    OPCodes_ADD_SP(static_cast<u8> (memory.Read(pc.GetValue())));
    pc.Increment();
}

func opcode0xE9() {
    // JP (HL)
    pc.SetValue(hl.GetValue());
}

func opcode0xEA() {
    // LD (nn),A
    SixteenBitRegister tmp;
    tmp.SetLow(memory.Read(pc.GetValue()));
    pc.Increment();
    tmp.SetHigh(memory.Read(pc.GetValue()));
    pc.Increment();
    OPCodes_LD(tmp.GetValue(), af.GetHigh());
}

func opcode0xEB() {
    invalidOPCode();
}

func opcode0xEC() {
    invalidOPCode();
}

func opcode0xED() {
    invalidOPCode();
}

func opcode0xEE() {
    // XOR n
    OPCodes_XOR(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xEF() {
    // RST 28H
    stackPush(&PC);
    pc.SetValue(0x28);
}

func opcode0xF0() {
    // LD A,(0xFF00+n)
    OPCodes_LD(af.GetHighRegister(),
            static_cast<u16> (0xFF00 + memory.Read(pc.GetValue())));
    pc.Increment();
}

func opcode0xF1() {
    // POP AF
    stackPop(&AF);
    af.SetLow(af.GetLow() & 0xF0);
}

func opcode0xF2() {
    // LD A,(C)
    OPCodes_LD(af.GetHighRegister(), static_cast<u16> (0xFF00 + bc.GetLow()));
}

func opcode0xF3() {
    // DI
    ime = false;
    m_iIMECycles = 0;
}

func opcode0xF4() {
    invalidOPCode();
}

func opcode0xF5() {
    // PUSH AF
    stackPush(&AF);
}

func opcode0xF6() {
    // OR n
    OPCodes_OR(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xF7() {
    // RST 30H
    stackPush(&PC);
    pc.SetValue(0x0030);
}

func opcode0xF8() {
    // LD HL,SP+n
    s8 n = memory.Read(pc.GetValue());
    u16 result = sp.GetValue() + n;
    clearAllFlags();
    if (((sp.GetValue() ^ n ^ result) & 0x100) == 0x100)
        toggleFlag(flagCarry);
    if (((sp.GetValue() ^ n ^ result) & 0x10) == 0x10)
        toggleFlag(flagHalf);
    hl.SetValue(result);
    pc.Increment();
}

func opcode0xF9() {
    // LD SP,HL
    sp.SetValue(hl.GetValue());
}

func opcode0xFA() {
    // LD A,(nn)
    SixteenBitRegister tmp;
    tmp.SetLow(memory.Read(pc.GetValue()));
    pc.Increment();
    tmp.SetHigh(memory.Read(pc.GetValue()));
    pc.Increment();
    OPCodes_LD(af.GetHighRegister(), tmp.GetValue());
}

func opcode0xFB() {
    // EI
    int ei_cycles = kOPCodeMachineCycles[0xFB] * AdjustedCycles(4);
    m_iIMECycles = ei_cycles + 1;
}

func opcode0xFC() {
    invalidOPCode();
}

func opcode0xFD() {
    invalidOPCode();
}

func opcode0xFE() {
    // CP n
    OPCodes_CP(memory.Read(pc.GetValue()));
    pc.Increment();
}

func opcode0xFF() {
    // RST 38H
    stackPush(&PC);
    pc.SetValue(0x0038);
}
*/
