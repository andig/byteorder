package byteorder

import (
	"encoding/binary"
	"strings"
)

const networkOrder = "ABCDEFGH"

type byteOrder string

// New creates a new binary.ByteOrder with the order of bytes given as a string ("ABCD")
func New(order string) binary.ByteOrder {
	return byteOrder(strings.ToUpper(order))
}

func (bo byteOrder) Uint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(bo.get(b))
}

func (bo byteOrder) Uint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(bo.get(b))
}

func (bo byteOrder) Uint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(bo.get(b))
}

func (bo byteOrder) PutUint16(b []byte, val uint16) {
	if strings.HasPrefix(networkOrder, string(bo)) {
		binary.BigEndian.PutUint16(b, val)
		return
	}

	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, val)
	bo.put(buf, b)
}

func (bo byteOrder) PutUint32(b []byte, val uint32) {
	if strings.HasPrefix(networkOrder, string(bo)) {
		binary.BigEndian.PutUint32(b, val)
		return
	}

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, val)
	bo.put(buf, b)
}

func (bo byteOrder) PutUint64(b []byte, val uint64) {
	if strings.HasPrefix(networkOrder, string(bo)) {
		binary.BigEndian.PutUint64(b, val)
		return
	}

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	bo.put(buf, b)
}

func (bo byteOrder) String() string {
	return string(bo)
}

func (bo byteOrder) get(b []byte) []byte {
	// bounds check hint to compiler; see golang.org/issue/14808
	_ = b[len(bo)-1]

	if strings.HasPrefix(networkOrder, string(bo)) {
		return b
	}

	res := make([]byte, len(b))
	for i, c := range bo {
		pos := c - 'A'
		res[pos] = b[i]
	}

	return res
}

func (bo byteOrder) put(b, res []byte) {
	for i, c := range bo {
		pos := c - 'A'
		res[pos] = b[i]
	}
}
