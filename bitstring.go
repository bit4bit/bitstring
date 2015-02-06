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

func (b bitString) ByUint64(at int, block int) uint64 {
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
