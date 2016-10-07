package memory

// IMemory is the interface for Memory
type IMemory interface {
	Setup([]uint8)
	Read(uint16) uint8
	Write(uint16, uint8)
}
