package gvariant

import (
	"encoding/binary"
	"unsafe"
)

// getHostByteOrder detects the host byte order
func getHostByteOrder() binary.ByteOrder {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	if buf[0] == 0xCD {
		return binary.LittleEndian
	}
	return binary.BigEndian
}

// guint64FromBE Converts a uint64_t value from big-endian to host byte order.
func guint64FromBE(beValue uint64) uint64 {
	hostOrder := getHostByteOrder()

	if hostOrder == binary.LittleEndian {
		// Host is little-endian, needs conversion
		beBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(beBytes, beValue)
		return binary.LittleEndian.Uint64(beBytes)
	}

	// Host is big-endian, no conversion needed
	return beValue
}
