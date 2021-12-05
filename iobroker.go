package byteorder

import "math"

// Unsigned 16 bit (Big Endian): AB
func Uint16be(b []byte) uint16 {
	return New("AB").Uint16(b)
}

// Unsigned 16 bit (Little Endian): BA
func Uint16le(b []byte) uint16 {
	return New("BA").Uint16(b)
}

// Unsigned 32 bit (Big Endian): ABCD
func Uint32be(b []byte) uint32 {
	return New("ABCD").Uint32(b)
}

// Unsigned 32 bit (Little Endian): DCBA
func Uint32le(b []byte) uint32 {
	return New("DCBA").Uint32(b)
}

// Unsigned 32 bit (Big Endian Swap Word): CDAB
func Uint32sw(b []byte) uint32 {
	return New("CDAB").Uint32(b)
}

// Unsigned 32 bit (Big Endian Swap Byte): DCBA
func Uint32sb(b []byte) uint32 {
	return New("DCBA").Uint32(b)
}

// Unsigned 64 bit (Big Endian): ABCDEFGH
func Uint64be(b []byte) uint64 {
	return New("ABCDEFGH").Uint64(b)
}

// Unsigned 64 bit (Little Endian): HGFEDCBA
func Uint64le(b []byte) uint64 {
	return New("HGFEDCBA").Uint64(b)
}

// Unsigned 64 bit (Big Endian Swap Word): GHEFCDAB
func Uint64sw(b []byte) uint64 {
	return New("GHEFCDAB").Uint64(b)
}

// Unsigned 64 bit (Big Endian Swap Byte): HGFEDCBA
func Uint64sb(b []byte) uint64 {
	return New("HGFEDCBA").Uint64(b)
}

// Float (Big Endian): ABCD
func Floatbe(b []byte) float32 {
	return math.Float32frombits(New("ABCD").Uint32(b))
}

// Float (Little Endian): DCBA
func Floatle(b []byte) float32 {
	return math.Float32frombits(New("DCBA").Uint32(b))
}

// Float (Big Endian Swap Word): CDAB
func Floatsw(b []byte) float32 {
	return math.Float32frombits(New("CDAB").Uint32(b))
}

// Float (Big Endian Swap Byte): DCBA
func Floatsb(b []byte) float32 {
	return math.Float32frombits(New("DCBA").Uint32(b))
}

// Double (Big Endian): ABCDEFGH
func Doublebe(b []byte) float64 {
	return math.Float64frombits(New("ABCDEFGH").Uint64(b))
}

// Double (Little Endian): HGFEDCBA
func Doublele(b []byte) float64 {
	return math.Float64frombits(New("HGFEDCBA").Uint64(b))
}
