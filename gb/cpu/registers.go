package cpu

// EightBitReg models an 8 bit processor register
type EightBitReg struct {
	value uint8
}

// SixteenBitReg models a 16 bit processor register
type SixteenBitReg struct {
	high EightBitReg
	low  EightBitReg
}

// SetValue sets the 8 bit value of the register
func (reg *EightBitReg) SetValue(value uint8) {
	reg.value = value
}

// GetValue returns the 8 bit value of the register
func (reg *EightBitReg) GetValue() uint8 {
	return reg.value
}

// Increment increments (++) the 8 bit value of the register
func (reg *EightBitReg) Increment() {
	reg.value++
}

// Decrement decrements (--) the 8 bit value of the register
func (reg *EightBitReg) Decrement() {
	reg.value--
}

// SetHigh sets the 8 bit value of the 16 bit higher part
func (reg *SixteenBitReg) SetHigh(value uint8) {
	reg.high.value = value
}

// GetHigh returns the 8 bit value of the 16 bit higher part
func (reg *SixteenBitReg) GetHigh() uint8 {
	return reg.high.value
}

// GetHighReg returns the 8 bit register of the 16 bit higher part
func (reg *SixteenBitReg) GetHighReg() *EightBitReg {
	return &reg.high
}

// SetLow sets the 8 bit value of the 16 bit lower part
func (reg *SixteenBitReg) SetLow(value uint8) {
	reg.low.value = value
}

// GetLow returns the 8 bit value of the 16 bit lower part
func (reg *SixteenBitReg) GetLow() uint8 {
	return reg.low.value
}

// GetLowReg returns the 8 bit register of the 16 bit lower part
func (reg *SixteenBitReg) GetLowReg() *EightBitReg {
	return &reg.low
}

// SetValue sets the 16 bit value of the register
func (reg *SixteenBitReg) SetValue(value uint16) {
	reg.low.SetValue(uint8(value & 0xFF))
	reg.high.SetValue(uint8((value >> 8) & 0xFF))
}

// GetValue returns the 16 bit value of the register
func (reg *SixteenBitReg) GetValue() uint16 {
	high := uint16(reg.high.GetValue())
	low := uint16(reg.low.GetValue())
	return (high << 8) | low
}

// Increment increments (++) the 16 bit value of the register
func (reg *SixteenBitReg) Increment() {
	value := reg.GetValue()
	value++
	reg.SetValue(value)
}

// Decrement decrements (--) the 16 bit value of the register
func (reg *SixteenBitReg) Decrement() {
	value := reg.GetValue()
	value--
	reg.SetValue(value)
}
