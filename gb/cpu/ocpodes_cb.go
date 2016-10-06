package cpu

func opcodeCB0x00() {
	// RLC B
	opcodesRLC(bc.GetHighReg())
}

func opcodeCB0x01() {
	// RLC C
	opcodesRLC(bc.GetLowReg())
}

func opcodeCB0x02() {
	// RLC D
	opcodesRLC(de.GetHighReg())
}

func opcodeCB0x03() {
	// RLC E
	opcodesRLC(de.GetLowReg())
}

func opcodeCB0x04() {
	// RLC H
	opcodesRLC(hl.GetHighReg())
}

func opcodeCB0x05() {
	// RLC L
	opcodesRLC(hl.GetLowReg())
}

func opcodeCB0x06() {
	// RLC (HL)
	opcodesRLCHL()
}

func opcodeCB0x07() {
	// RLC A
	opcodesRLC(af.GetHighReg())
}

func opcodeCB0x08() {
	// RRC B
	opcodesRRC(bc.GetHighReg())
}

func opcodeCB0x09() {
	// RRC C
	opcodesRRC(bc.GetLowReg())
}

func opcodeCB0x0A() {
	// RRC D
	opcodesRRC(de.GetHighReg())
}

func opcodeCB0x0B() {
	// RRC E
	opcodesRRC(de.GetLowReg())
}

func opcodeCB0x0C() {
	// RRC H
	opcodesRRC(hl.GetHighReg())
}

func opcodeCB0x0D() {
	// RRC L
	opcodesRRC(hl.GetLowReg())
}

func opcodeCB0x0E() {
	// RRC (HL)
	opcodesRRCHL()
}

func opcodeCB0x0F() {
	// RRC A
	opcodesRRC(af.GetHighReg())
}

func opcodeCB0x10() {
	// RL B
	opcodesRL(bc.GetHighReg())
}

func opcodeCB0x11() {
	// RL C
	opcodesRL(bc.GetLowReg())
}

func opcodeCB0x12() {
	// RL D
	opcodesRL(de.GetHighReg())
}

func opcodeCB0x13() {
	// RL E
	opcodesRL(de.GetLowReg())
}

func opcodeCB0x14() {
	// RL H
	opcodesRL(hl.GetHighReg())
}

func opcodeCB0x15() {
	// RL L
	opcodesRL(hl.GetLowReg())
}

func opcodeCB0x16() {
	// RL (HL)
	opcodesRLHL()
}

func opcodeCB0x17() {
	// RL A
	opcodesRL(af.GetHighReg())
}

func opcodeCB0x18() {
	// RR B
	opcodesRR(bc.GetHighReg())
}

func opcodeCB0x19() {
	// RR C
	opcodesRR(bc.GetLowReg())
}

func opcodeCB0x1A() {
	// RR D
	opcodesRR(de.GetHighReg())
}

func opcodeCB0x1B() {
	// RR E
	opcodesRR(de.GetLowReg())
}

func opcodeCB0x1C() {
	// RR H
	opcodesRR(hl.GetHighReg())
}

func opcodeCB0x1D() {
	// RR L
	opcodesRR(hl.GetLowReg())
}

func opcodeCB0x1E() {
	// RR (HL)
	opcodesRRHL()
}

func opcodeCB0x1F() {
	// RR A
	opcodesRR(af.GetHighReg())
}

func opcodeCB0x20() {
	// SLA B
	opcodesSLA(bc.GetHighReg())
}

func opcodeCB0x21() {
	// SLA C
	opcodesSLA(bc.GetLowReg())
}

func opcodeCB0x22() {
	// SLA D
	opcodesSLA(de.GetHighReg())
}

func opcodeCB0x23() {
	// SLA E
	opcodesSLA(de.GetLowReg())
}

func opcodeCB0x24() {
	// SLA H
	opcodesSLA(hl.GetHighReg())
}

func opcodeCB0x25() {
	// SLA L
	opcodesSLA(hl.GetLowReg())
}

func opcodeCB0x26() {
	// SLA (HL)
	opcodesSLAHL()
}

func opcodeCB0x27() {
	// SLA A
	opcodesSLA(af.GetHighReg())
}

func opcodeCB0x28() {
	// SRA B
	opcodesSRA(bc.GetHighReg())
}

func opcodeCB0x29() {
	// SRA C
	opcodesSRA(bc.GetLowReg())
}

func opcodeCB0x2A() {
	// SRA D
	opcodesSRA(de.GetHighReg())
}

func opcodeCB0x2B() {
	// SRA E
	opcodesSRA(de.GetLowReg())
}

func opcodeCB0x2C() {
	// SRA H
	opcodesSRA(hl.GetHighReg())
}

func opcodeCB0x2D() {
	// SRA L
	opcodesSRA(hl.GetLowReg())
}

func opcodeCB0x2E() {
	// SRA (HL)
	opcodesSRAHL()
}

func opcodeCB0x2F() {
	// SRA A
	opcodesSRA(af.GetHighReg())
}

func opcodeCB0x30() {
	// SWAP B
	opcodesSWAPReg(bc.GetHighReg())
}

func opcodeCB0x31() {
	// SWAP C
	opcodesSWAPReg(bc.GetLowReg())
}

func opcodeCB0x32() {
	// SWAP D
	opcodesSWAPReg(de.GetHighReg())
}

func opcodeCB0x33() {
	// SWAP E
	opcodesSWAPReg(de.GetLowReg())
}

func opcodeCB0x34() {
	// SWAP H
	opcodesSWAPReg(hl.GetHighReg())
}

func opcodeCB0x35() {
	// SWAP L
	opcodesSWAPReg(hl.GetLowReg())
}

func opcodeCB0x36() {
	// SWAP (HL)
	opcodesSWAPHL()
}

func opcodeCB0x37() {
	// SWAP A
	opcodesSWAPReg(af.GetHighReg())
}

func opcodeCB0x38() {
	// SRL B
	opcodesSRL(bc.GetHighReg())
}

func opcodeCB0x39() {
	// SRL C
	opcodesSRL(bc.GetLowReg())
}

func opcodeCB0x3A() {
	// SRL D
	opcodesSRL(de.GetHighReg())
}

func opcodeCB0x3B() {
	// SRL E
	opcodesSRL(de.GetLowReg())
}

func opcodeCB0x3C() {
	// SRL H
	opcodesSRL(hl.GetHighReg())
}

func opcodeCB0x3D() {
	// SRL L
	opcodesSRL(hl.GetLowReg())
}

func opcodeCB0x3E() {
	// SRL (HL)
	opcodesSRLHL()
}

func opcodeCB0x3F() {
	// SRL A
	opcodesSRL(af.GetHighReg())
}

func opcodeCB0x40() {
	// BIT 0 B
	opcodesBIT(bc.GetHighReg(), 0)
}

func opcodeCB0x41() {
	// BIT 0 C
	opcodesBIT(bc.GetLowReg(), 0)
}

func opcodeCB0x42() {
	// BIT 0 D
	opcodesBIT(de.GetHighReg(), 0)
}

func opcodeCB0x43() {
	// BIT 0 E
	opcodesBIT(de.GetLowReg(), 0)
}

func opcodeCB0x44() {
	// BIT 0 H
	opcodesBIT(hl.GetHighReg(), 0)
}

func opcodeCB0x45() {
	// BIT 0 L
	opcodesBIT(hl.GetLowReg(), 0)
}

func opcodeCB0x46() {
	// BIT 0 (HL)
	opcodesBITHL(0)
}

func opcodeCB0x47() {
	// BIT 0 A
	opcodesBIT(af.GetHighReg(), 0)
}

func opcodeCB0x48() {
	// BIT 1 B
	opcodesBIT(bc.GetHighReg(), 1)
}

func opcodeCB0x49() {
	// BIT 1 C
	opcodesBIT(bc.GetLowReg(), 1)
}

func opcodeCB0x4A() {
	// BIT 1 D
	opcodesBIT(de.GetHighReg(), 1)
}

func opcodeCB0x4B() {
	// BIT 1 E
	opcodesBIT(de.GetLowReg(), 1)
}

func opcodeCB0x4C() {
	// BIT 1 H
	opcodesBIT(hl.GetHighReg(), 1)
}

func opcodeCB0x4D() {
	// BIT 1 L
	opcodesBIT(hl.GetLowReg(), 1)
}

func opcodeCB0x4E() {
	// BIT 1 (HL)
	opcodesBITHL(1)
}

func opcodeCB0x4F() {
	// BIT 1 A
	opcodesBIT(af.GetHighReg(), 1)
}

func opcodeCB0x50() {
	// BIT 2 B
	opcodesBIT(bc.GetHighReg(), 2)
}

func opcodeCB0x51() {
	// BIT 2 C
	opcodesBIT(bc.GetLowReg(), 2)
}

func opcodeCB0x52() {
	// BIT 2 D
	opcodesBIT(de.GetHighReg(), 2)
}

func opcodeCB0x53() {
	// BIT 2 E
	opcodesBIT(de.GetLowReg(), 2)
}

func opcodeCB0x54() {
	// BIT 2 H
	opcodesBIT(hl.GetHighReg(), 2)
}

func opcodeCB0x55() {
	// BIT 2 L
	opcodesBIT(hl.GetLowReg(), 2)
}

func opcodeCB0x56() {
	// BIT 2 (HL)
	opcodesBITHL(2)
}

func opcodeCB0x57() {
	// BIT 2 A
	opcodesBIT(af.GetHighReg(), 2)
}

func opcodeCB0x58() {
	// BIT 3 B
	opcodesBIT(bc.GetHighReg(), 3)
}

func opcodeCB0x59() {
	// BIT 3 C
	opcodesBIT(bc.GetLowReg(), 3)
}

func opcodeCB0x5A() {
	// BIT 3 D
	opcodesBIT(de.GetHighReg(), 3)
}

func opcodeCB0x5B() {
	// BIT 3 E
	opcodesBIT(de.GetLowReg(), 3)
}

func opcodeCB0x5C() {
	// BIT 3 H
	opcodesBIT(hl.GetHighReg(), 3)
}

func opcodeCB0x5D() {
	// BIT 3 L
	opcodesBIT(hl.GetLowReg(), 3)
}

func opcodeCB0x5E() {
	// BIT 3 (HL)
	opcodesBITHL(3)
}

func opcodeCB0x5F() {
	// BIT 3 A
	opcodesBIT(af.GetHighReg(), 3)
}

func opcodeCB0x60() {
	// BIT 4 B
	opcodesBIT(bc.GetHighReg(), 4)
}

func opcodeCB0x61() {
	// BIT 4 C
	opcodesBIT(bc.GetLowReg(), 4)
}

func opcodeCB0x62() {
	// BIT 4 D
	opcodesBIT(de.GetHighReg(), 4)
}

func opcodeCB0x63() {
	// BIT 4 E
	opcodesBIT(de.GetLowReg(), 4)
}

func opcodeCB0x64() {
	// BIT 4 H
	opcodesBIT(hl.GetHighReg(), 4)
}

func opcodeCB0x65() {
	// BIT 4 L
	opcodesBIT(hl.GetLowReg(), 4)
}

func opcodeCB0x66() {
	// BIT 4 (HL)
	opcodesBITHL(4)
}

func opcodeCB0x67() {
	// BIT 4 A
	opcodesBIT(af.GetHighReg(), 4)
}

func opcodeCB0x68() {
	// BIT 5 B
	opcodesBIT(bc.GetHighReg(), 5)
}

func opcodeCB0x69() {
	// BIT 5 C
	opcodesBIT(bc.GetLowReg(), 5)
}

func opcodeCB0x6A() {
	// BIT 5 D
	opcodesBIT(de.GetHighReg(), 5)
}

func opcodeCB0x6B() {
	// BIT 5 E
	opcodesBIT(de.GetLowReg(), 5)
}

func opcodeCB0x6C() {
	// BIT 5 H
	opcodesBIT(hl.GetHighReg(), 5)
}

func opcodeCB0x6D() {
	// BIT 5 L
	opcodesBIT(hl.GetLowReg(), 5)
}

func opcodeCB0x6E() {
	// BIT 5 (HL)
	opcodesBITHL(5)
}

func opcodeCB0x6F() {
	// BIT 5 A
	opcodesBIT(af.GetHighReg(), 5)
}

func opcodeCB0x70() {
	// BIT 6 B
	opcodesBIT(bc.GetHighReg(), 6)
}

func opcodeCB0x71() {
	// BIT 6 C
	opcodesBIT(bc.GetLowReg(), 6)
}

func opcodeCB0x72() {
	// BIT 6 D
	opcodesBIT(de.GetHighReg(), 6)
}

func opcodeCB0x73() {
	// BIT 6 E
	opcodesBIT(de.GetLowReg(), 6)
}

func opcodeCB0x74() {
	// BIT 6 H
	opcodesBIT(hl.GetHighReg(), 6)
}

func opcodeCB0x75() {
	// BIT 6 L
	opcodesBIT(hl.GetLowReg(), 6)
}

func opcodeCB0x76() {
	// BIT 6 (HL)
	opcodesBITHL(6)
}

func opcodeCB0x77() {
	// BIT 6 A
	opcodesBIT(af.GetHighReg(), 6)
}

func opcodeCB0x78() {
	// BIT 7 B
	opcodesBIT(bc.GetHighReg(), 7)
}

func opcodeCB0x79() {
	// BIT 7 C
	opcodesBIT(bc.GetLowReg(), 7)
}

func opcodeCB0x7A() {
	// BIT 7 D
	opcodesBIT(de.GetHighReg(), 7)
}

func opcodeCB0x7B() {
	// BIT 7 E
	opcodesBIT(de.GetLowReg(), 7)
}

func opcodeCB0x7C() {
	// BIT 7 H
	opcodesBIT(hl.GetHighReg(), 7)
}

func opcodeCB0x7D() {
	// BIT 7 L
	opcodesBIT(hl.GetLowReg(), 7)
}

func opcodeCB0x7E() {
	// BIT 7 (HL)
	opcodesBITHL(7)
}

func opcodeCB0x7F() {
	// BIT 7 A
	opcodesBIT(af.GetHighReg(), 7)
}

func opcodeCB0x80() {
	// RES 0 B
	opcodesRES(bc.GetHighReg(), 0)
}

func opcodeCB0x81() {
	// RES 0 C
	opcodesRES(bc.GetLowReg(), 0)
}

func opcodeCB0x82() {
	// RES 0 D
	opcodesRES(de.GetHighReg(), 0)
}

func opcodeCB0x83() {
	// RES 0 E
	opcodesRES(de.GetLowReg(), 0)
}

func opcodeCB0x84() {
	// RES 0 H
	opcodesRES(hl.GetHighReg(), 0)
}

func opcodeCB0x85() {
	// RES 0 L
	opcodesRES(hl.GetLowReg(), 0)
}

func opcodeCB0x86() {
	// RES 0 (HL)
	opcodesRESHL(0)
}

func opcodeCB0x87() {
	// RES 0 A
	opcodesRES(af.GetHighReg(), 0)
}

func opcodeCB0x88() {
	// RES 1 B
	opcodesRES(bc.GetHighReg(), 1)
}

func opcodeCB0x89() {
	// RES 1 C
	opcodesRES(bc.GetLowReg(), 1)
}

func opcodeCB0x8A() {
	// RES 1 D
	opcodesRES(de.GetHighReg(), 1)
}

func opcodeCB0x8B() {
	// RES 1 E
	opcodesRES(de.GetLowReg(), 1)
}

func opcodeCB0x8C() {
	// RES 1 H
	opcodesRES(hl.GetHighReg(), 1)
}

func opcodeCB0x8D() {
	// RES 1 L
	opcodesRES(hl.GetLowReg(), 1)
}

func opcodeCB0x8E() {
	// RES 1 (HL)
	opcodesRESHL(1)
}

func opcodeCB0x8F() {
	// RES 1 A
	opcodesRES(af.GetHighReg(), 1)
}

func opcodeCB0x90() {
	// RES 2 B
	opcodesRES(bc.GetHighReg(), 2)
}

func opcodeCB0x91() {
	// RES 2 C
	opcodesRES(bc.GetLowReg(), 2)
}

func opcodeCB0x92() {
	// RES 2 D
	opcodesRES(de.GetHighReg(), 2)
}

func opcodeCB0x93() {
	// RES 2 E
	opcodesRES(de.GetLowReg(), 2)
}

func opcodeCB0x94() {
	// RES 2 H
	opcodesRES(hl.GetHighReg(), 2)
}

func opcodeCB0x95() {
	// RES 2 L
	opcodesRES(hl.GetLowReg(), 2)
}

func opcodeCB0x96() {
	// RES 2 (HL)
	opcodesRESHL(2)
}

func opcodeCB0x97() {
	// RES 2 A
	opcodesRES(af.GetHighReg(), 2)
}

func opcodeCB0x98() {
	// RES 3 B
	opcodesRES(bc.GetHighReg(), 3)
}

func opcodeCB0x99() {
	// RES 3 C
	opcodesRES(bc.GetLowReg(), 3)
}

func opcodeCB0x9A() {
	// RES 3 D
	opcodesRES(de.GetHighReg(), 3)
}

func opcodeCB0x9B() {
	// RES 3 E
	opcodesRES(de.GetLowReg(), 3)
}

func opcodeCB0x9C() {
	// RES 3 H
	opcodesRES(hl.GetHighReg(), 3)
}

func opcodeCB0x9D() {
	// RES 3 L
	opcodesRES(hl.GetLowReg(), 3)
}

func opcodeCB0x9E() {
	// RES 3 (HL)
	opcodesRESHL(3)
}

func opcodeCB0x9F() {
	// RES 3 A
	opcodesRES(af.GetHighReg(), 3)
}

func opcodeCB0xA0() {
	// RES 4 B
	opcodesRES(bc.GetHighReg(), 4)
}

func opcodeCB0xA1() {
	// RES 4 C
	opcodesRES(bc.GetLowReg(), 4)
}

func opcodeCB0xA2() {
	// RES 4 D
	opcodesRES(de.GetHighReg(), 4)
}

func opcodeCB0xA3() {
	// RES 4 E
	opcodesRES(de.GetLowReg(), 4)
}

func opcodeCB0xA4() {
	// RES 4 H
	opcodesRES(hl.GetHighReg(), 4)
}

func opcodeCB0xA5() {
	// RES 4 L
	opcodesRES(hl.GetLowReg(), 4)
}

func opcodeCB0xA6() {
	// RES 4 (HL)
	opcodesRESHL(4)
}

func opcodeCB0xA7() {
	// RES 4 A
	opcodesRES(af.GetHighReg(), 4)
}

func opcodeCB0xA8() {
	// RES 5 B
	opcodesRES(bc.GetHighReg(), 5)
}

func opcodeCB0xA9() {
	// RES 5 C
	opcodesRES(bc.GetLowReg(), 5)
}

func opcodeCB0xAA() {
	// RES 5 D
	opcodesRES(de.GetHighReg(), 5)
}

func opcodeCB0xAB() {
	// RES 5 E
	opcodesRES(de.GetLowReg(), 5)
}

func opcodeCB0xAC() {
	// RES 5 H
	opcodesRES(hl.GetHighReg(), 5)
}

func opcodeCB0xAD() {
	// RES 5 L
	opcodesRES(hl.GetLowReg(), 5)
}

func opcodeCB0xAE() {
	// RES 5 (HL)
	opcodesRESHL(5)
}

func opcodeCB0xAF() {
	// RES 5 A
	opcodesRES(af.GetHighReg(), 5)
}

func opcodeCB0xB0() {
	// RES 6 B
	opcodesRES(bc.GetHighReg(), 6)
}

func opcodeCB0xB1() {
	// RES 6 C
	opcodesRES(bc.GetLowReg(), 6)
}

func opcodeCB0xB2() {
	// RES 6 D
	opcodesRES(de.GetHighReg(), 6)
}

func opcodeCB0xB3() {
	// RES 6 E
	opcodesRES(de.GetLowReg(), 6)
}

func opcodeCB0xB4() {
	// RES 6 H
	opcodesRES(hl.GetHighReg(), 6)
}

func opcodeCB0xB5() {
	// RES 6 L
	opcodesRES(hl.GetLowReg(), 6)
}

func opcodeCB0xB6() {
	// RES 6 (HL)
	opcodesRESHL(6)
}

func opcodeCB0xB7() {
	// RES 6 A
	opcodesRES(af.GetHighReg(), 6)
}

func opcodeCB0xB8() {
	// RES 7 B
	opcodesRES(bc.GetHighReg(), 7)
}

func opcodeCB0xB9() {
	// RES 7 C
	opcodesRES(bc.GetLowReg(), 7)
}

func opcodeCB0xBA() {
	// RES 7 D
	opcodesRES(de.GetHighReg(), 7)
}

func opcodeCB0xBB() {
	// RES 7 E
	opcodesRES(de.GetLowReg(), 7)
}

func opcodeCB0xBC() {
	// RES 7 H
	opcodesRES(hl.GetHighReg(), 7)
}

func opcodeCB0xBD() {
	// RES 7 L
	opcodesRES(hl.GetLowReg(), 7)
}

func opcodeCB0xBE() {
	// RES 7 (HL)
	opcodesRESHL(7)
}

func opcodeCB0xBF() {
	// RES 7 A
	opcodesRES(af.GetHighReg(), 7)
}

func opcodeCB0xC0() {
	// SET 0 B
	opcodesSET(bc.GetHighReg(), 0)
}

func opcodeCB0xC1() {
	// SET 0 C
	opcodesSET(bc.GetLowReg(), 0)
}

func opcodeCB0xC2() {
	// SET 0 D
	opcodesSET(de.GetHighReg(), 0)
}

func opcodeCB0xC3() {
	// SET 0 E
	opcodesSET(de.GetLowReg(), 0)
}

func opcodeCB0xC4() {
	// SET 0 H
	opcodesSET(hl.GetHighReg(), 0)
}

func opcodeCB0xC5() {
	// SET 0 L
	opcodesSET(hl.GetLowReg(), 0)
}

func opcodeCB0xC6() {
	// SET 0 (HL)
	opcodesSETHL(0)
}

func opcodeCB0xC7() {
	// SET 0 A
	opcodesSET(af.GetHighReg(), 0)
}

func opcodeCB0xC8() {
	// SET 1 B
	opcodesSET(bc.GetHighReg(), 1)
}

func opcodeCB0xC9() {
	// SET 1 C
	opcodesSET(bc.GetLowReg(), 1)
}

func opcodeCB0xCA() {
	// SET 1 D
	opcodesSET(de.GetHighReg(), 1)
}

func opcodeCB0xCB() {
	// SET 1 E
	opcodesSET(de.GetLowReg(), 1)
}

func opcodeCB0xCC() {
	// SET 1 H
	opcodesSET(hl.GetHighReg(), 1)
}

func opcodeCB0xCD() {
	// SET 1 L
	opcodesSET(hl.GetLowReg(), 1)
}

func opcodeCB0xCE() {
	// SET 1 (HL)
	opcodesSETHL(1)
}

func opcodeCB0xCF() {
	// SET 1 A
	opcodesSET(af.GetHighReg(), 1)
}

func opcodeCB0xD0() {
	// SET 2 B
	opcodesSET(bc.GetHighReg(), 2)
}

func opcodeCB0xD1() {
	// SET 2 C
	opcodesSET(bc.GetLowReg(), 2)
}

func opcodeCB0xD2() {
	// SET 2 D
	opcodesSET(de.GetHighReg(), 2)
}

func opcodeCB0xD3() {
	// SET 2 E
	opcodesSET(de.GetLowReg(), 2)
}

func opcodeCB0xD4() {
	// SET 2 H
	opcodesSET(hl.GetHighReg(), 2)
}

func opcodeCB0xD5() {
	// SET 2 L
	opcodesSET(hl.GetLowReg(), 2)
}

func opcodeCB0xD6() {
	// SET 2 (HL)
	opcodesSETHL(2)
}

func opcodeCB0xD7() {
	// SET 2 A
	opcodesSET(af.GetHighReg(), 2)
}

func opcodeCB0xD8() {
	// SET 3 B
	opcodesSET(bc.GetHighReg(), 3)
}

func opcodeCB0xD9() {
	// SET 3 C
	opcodesSET(bc.GetLowReg(), 3)
}

func opcodeCB0xDA() {
	// SET 3 D
	opcodesSET(de.GetHighReg(), 3)
}

func opcodeCB0xDB() {
	// SET 3 E
	opcodesSET(de.GetLowReg(), 3)
}

func opcodeCB0xDC() {
	// SET 3 H
	opcodesSET(hl.GetHighReg(), 3)
}

func opcodeCB0xDD() {
	// SET 3 L
	opcodesSET(hl.GetLowReg(), 3)
}

func opcodeCB0xDE() {
	// SET 3 (HL)
	opcodesSETHL(3)
}

func opcodeCB0xDF() {
	// SET 3 A
	opcodesSET(af.GetHighReg(), 3)
}

func opcodeCB0xE0() {
	// SET 4 B
	opcodesSET(bc.GetHighReg(), 4)
}

func opcodeCB0xE1() {
	// SET 4 C
	opcodesSET(bc.GetLowReg(), 4)
}

func opcodeCB0xE2() {
	// SET 4 D
	opcodesSET(de.GetHighReg(), 4)
}

func opcodeCB0xE3() {
	// SET 4 E
	opcodesSET(de.GetLowReg(), 4)
}

func opcodeCB0xE4() {
	// SET 4 H
	opcodesSET(hl.GetHighReg(), 4)
}

func opcodeCB0xE5() {
	// SET 4 L
	opcodesSET(hl.GetLowReg(), 4)
}

func opcodeCB0xE6() {
	// SET 4 (HL)
	opcodesSETHL(4)
}

func opcodeCB0xE7() {
	// SET 4 A
	opcodesSET(af.GetHighReg(), 4)

}

func opcodeCB0xE8() {
	// SET 5 B
	opcodesSET(bc.GetHighReg(), 5)
}

func opcodeCB0xE9() {
	// SET 5 C
	opcodesSET(bc.GetLowReg(), 5)
}

func opcodeCB0xEA() {
	// SET 5 D
	opcodesSET(de.GetHighReg(), 5)
}

func opcodeCB0xEB() {
	// SET 5 E
	opcodesSET(de.GetLowReg(), 5)
}

func opcodeCB0xEC() {
	// SET 5 H
	opcodesSET(hl.GetHighReg(), 5)
}

func opcodeCB0xED() {
	// SET 5 L
	opcodesSET(hl.GetLowReg(), 5)
}

func opcodeCB0xEE() {
	// SET 5 (HL)
	opcodesSETHL(5)
}

func opcodeCB0xEF() {
	// SET 5 A
	opcodesSET(af.GetHighReg(), 5)
}

func opcodeCB0xF0() {
	// SET 6 B
	opcodesSET(bc.GetHighReg(), 6)
}

func opcodeCB0xF1() {
	// SET 6 C
	opcodesSET(bc.GetLowReg(), 6)
}

func opcodeCB0xF2() {
	// SET 6 D
	opcodesSET(de.GetHighReg(), 6)
}

func opcodeCB0xF3() {
	// SET 6 E
	opcodesSET(de.GetLowReg(), 6)
}

func opcodeCB0xF4() {
	// SET 6 H
	opcodesSET(hl.GetHighReg(), 6)
}

func opcodeCB0xF5() {
	// SET 6 L
	opcodesSET(hl.GetLowReg(), 6)
}

func opcodeCB0xF6() {
	// SET 6 (HL)
	opcodesSETHL(6)
}

func opcodeCB0xF7() {
	// SET 6 A
	opcodesSET(af.GetHighReg(), 6)
}

func opcodeCB0xF8() {
	// SET 7 B
	opcodesSET(bc.GetHighReg(), 7)
}

func opcodeCB0xF9() {
	// SET 7 C
	opcodesSET(bc.GetLowReg(), 7)
}

func opcodeCB0xFA() {
	// SET 7 D
	opcodesSET(de.GetHighReg(), 7)
}

func opcodeCB0xFB() {
	// SET 7 E
	opcodesSET(de.GetLowReg(), 7)
}

func opcodeCB0xFC() {
	// SET 7 H
	opcodesSET(hl.GetHighReg(), 7)
}

func opcodeCB0xFD() {
	// SET 7 L
	opcodesSET(hl.GetLowReg(), 7)
}

func opcodeCB0xFE() {
	// SET 7 (HL)
	opcodesSETHL(7)
}

func opcodeCB0xFF() {
	// SET 7 A
	opcodesSET(af.GetHighReg(), 7)
}
