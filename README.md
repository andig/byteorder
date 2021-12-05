# byteorder [![GoDoc](https://godoc.org/github.com/andig/byteorder?status.svg)](http://godoc.org/github.com/andig/byteorder)

`byteorder` is a Go module for working with arbitrarily-ordered byte slices.
It is useful e.g. when dealing with Modbus wire formats.

Installation:

    go get github.com/andig/byteorder

Usage:

    import "github.com/andig/byteorder"

    // reverse byte order, i.e. little endian
    bo := New("BA")
    fmt.Printf("0x%04x", bo.Uint16([]byte{1, 2}))
    // Output: 0x0201

NOTE: `byteorder` is not optimized for performance. The conversion functions will allocate which may not be desirable in high-frequency code paths.

In addition to the Go `binary.ByteOrder` implementation, this modules provides convenience functions for the [ioBroker.modbus](https://github.com/ioBroker/ioBroker.modbus) data types.
