# bitstring
extend Go asn1.BitString

Extract *at* position *nmemb* bits.

~~~ go
reader := BitString([]byte{0x1, 0x2, 0x7}, 64)

//get bit
reader.At(0, 1)
//get first byte
reader.By(0, 8);
//or get 8 bytes
reader.ByUint64(0)
//or get 4 bytes
reader.ByUint32(0)
//or get 2 bytes
reader.ByUint16(0)
//or get 1 byte
reader.ByUint8(0)
~~~