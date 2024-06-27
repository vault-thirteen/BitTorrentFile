package endianness

import (
	"encoding/binary"
	"unsafe"
)

func GetNativeEndianness() binary.ByteOrder {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		return binary.LittleEndian

	case [2]byte{0xAB, 0xCD}:
		return binary.BigEndian

	default:
		panic("native endianness is unknown")
	}
}
