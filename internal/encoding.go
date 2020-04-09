package internal

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

var (
	// UTF8Encoding is the UTF-8 encoding.
	UTF8Encoding encoding.Encoding = unicode.UTF8
	// UTF16BEEncoding is the UTF-16 Big Endian encoding.
	UTF16BEEncoding encoding.Encoding = unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	// UTF16LEEncoding is the UTF-16 Little Endian encoding.
	UTF16LEEncoding encoding.Encoding = unicode.UTF16(unicode.LittleEndian, unicode.UseBOM)
)
