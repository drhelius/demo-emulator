package util

// SetBit sets the selected bit to 1 inside a byte and returns the new byte
func SetBit(value uint8, bit uint8) uint8 {
	return value | (0x01 << bit)
}

// UnsetBit clears the selected bit to 0 inside a byte and returns the new byte
func UnsetBit(value uint8, bit uint8) uint8 {
	return value & (^(0x01 << bit))
}

// IsSetBit returns true if the selected bit inside a byte is 1
func IsSetBit(value uint8, bit uint8) bool {
	return (value & (0x01 << bit)) != 0
}
