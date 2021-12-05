package byteorder

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func ExampleNew() {
	// reverse byte order, i.e. little endian
	bo := New("BA")
	fmt.Printf("0x%04x", bo.Uint16([]byte{1, 2}))
	// Output: 0x0201
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func Test16(t *testing.T) {
	bo := "AB"
	in := []byte{1, 2}

	test16(New(bo), in, 0x0102, t)
	test16(New(reverse(bo)), in, 0x0201, t)
}

func test16(bo binary.ByteOrder, in []byte, expect uint16, t *testing.T) {
	if u := bo.Uint16(in); u != expect {
		t.Errorf("Uint16: %x", u)

		out := make([]byte, 2)
		bo.PutUint16(out, u)
		if !bytes.Equal(in, out) {
			t.Errorf("PutUint16: %x", out)
		}
	}
}

func Test32(t *testing.T) {
	bo := "ABCD"
	in := []byte{1, 2, 3, 4}

	test16(New(bo), in, 0x0102, t)
	test16(New(reverse(bo)), in, 0x0403, t)

	test32(New(bo), in, 0x01020304, t)
	test32(New(reverse(bo)), in, 0x04030201, t)
}

func test32(bo binary.ByteOrder, in []byte, expect uint32, t *testing.T) {
	if u := bo.Uint32(in); u != expect {
		t.Errorf("Uint32: %x", u)

		out := make([]byte, 4)
		bo.PutUint32(out, u)
		if !bytes.Equal(in, out) {
			t.Errorf("PutUint32: %x", out)
		}
	}
}

func Test64(t *testing.T) {
	bo := "ABCDEFGH"
	in := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	test16(New(bo), in, 0x0102, t)
	test16(New(reverse(bo)), in, 0x0807, t)

	test32(New(bo), in, 0x01020304, t)
	test32(New(reverse(bo)), in, 0x08070605, t)

	test64(New(bo), in, 0x0102030405060708, t)
	test64(New(reverse(bo)), in, 0x0807060504030201, t)
}

func test64(bo binary.ByteOrder, in []byte, expect uint64, t *testing.T) {
	if u := bo.Uint64(in); u != expect {
		t.Errorf("Uint64: %x", u)

		out := make([]byte, 8)
		bo.PutUint64(out, u)
		if !bytes.Equal(in, out) {
			t.Errorf("PutUint32: %x", out)
		}
	}
}
