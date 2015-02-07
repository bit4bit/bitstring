package bitstring

import (
	"encoding/asn1"
)

type bitString struct {
	asn1.BitString
}

func BitString(data []byte, bitLength int) bitString {
	return bitString{asn1.BitString{data, bitLength}}
}

func (b bitString) By(at int, block int) uint64 {
	if at < 0 || at >= b.BitLength || block > 64 {
		return 0
	}

	ret := uint64(0)
	idx := 0

	for ; block > 0; block-- {
		pos := at + idx
		x := pos / 8
		y := 7 - uint(pos%8)
		val := uint64(b.Bytes[x]>>y) & 1
		ret <<= uint64(1)
		ret |= uint64(val)
		idx++
	}

	return ret
}
func (b bitString) ByUint64(at int) uint64 {
	return b.By(at, 64)
}

func (b bitString) ByUint32(at int) uint32 {
	return uint32(b.By(at, 32))
}

func (b bitString) ByUint16(at int) uint16 {
	return uint16(b.By(at, 16))
}

func (b bitString) ByUint8(at int) uint8 {
	return uint8(b.By(at, 8))
}
