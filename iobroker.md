# iobroker data types

Source: https://github.com/ioBroker/ioBroker.modbus

Uint8be - Unsigned 8 bit (Big Endian): A
Uint8le - Unsigned 8 bit (Little Endian): A
Uint16be - Unsigned 16 bit (Big Endian): AB
Uint16le - Unsigned 16 bit (Little Endian): BA
Uint32be - Unsigned 32 bit (Big Endian): ABCD
Uint32le - Unsigned 32 bit (Little Endian): DCBA
Uint32sw - Unsigned 32 bit (Big Endian Word Swap): CDAB
Uint32sb - Unsigned 32 bit (Big Endian Byte Swap): DCBA
Uint64be - Unsigned 64 bit (Big Endian): ABCDEFGH
Uint64le - Unsigned 64 bit (Little Endian): HGFEDCBA

Floatbe - Float (Big Endian): ABCD
Floatle - Float (Little Endian): DCBA
Floatsw - Float (Big Endian Word Swap): CDAB
Floatsb - Float (Big Endian Byte Swap): DCBA
Doublebe - Double (Big Endian): ABCDEFGH
Doublele - Double (Little Endian): HGFEDCBA

Int8be - Signed 8 bit (Big Endian): A
Int8le - Signed 8 bit (Little Endian): A
Int16be - Signed 16 bit (Big Endian): AB
Int16le - Signed 16 bit (Little Endian): BA
Int32be - Signed 32 bit (Big Endian): ABCD
Int32le - Signed 32 bit (Little Endian): DCBA
Int32sw - Signed 32 bit (Big Endian Word Swap): CDAB
Int32sb - Signed 32 bit (Big Endian Byte Swap): DCBA

String - String (Zero-end): ABCDEF\0
Stringle - String (Little Endian, Zero-end): ABCDEF\0
