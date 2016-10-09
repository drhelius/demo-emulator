package mapper

// Mapper is the interface to implement MBCs
type Mapper interface {
	Setup([]uint8)
	Read(uint16) uint8
	Write(uint16, uint8)
}
