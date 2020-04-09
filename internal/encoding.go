package internal

import (
	"encoding/xml"
	"io"

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

// NewUTF8Encoder ...
func NewUTF8Encoder() *encoding.Encoder {
	return UTF8Encoding.NewEncoder()
}

// NewUTF8Decoder ...
func NewUTF8Decoder() *encoding.Decoder {
	return UTF8Encoding.NewDecoder()
}

// NewUTF16BEEncoder ...
func NewUTF16BEEncoder() *encoding.Encoder {
	return UTF16BEEncoding.NewEncoder()
}

// NewUTF16BEDecoder ...
func NewUTF16BEDecoder() *encoding.Decoder {
	return UTF16BEEncoding.NewDecoder()
}

// NewUTF16LEEncoder ...
func NewUTF16LEEncoder() *encoding.Encoder {
	return UTF16LEEncoding.NewEncoder()
}

// NewUTF16LEDecoder ...
func NewUTF16LEDecoder() *encoding.Decoder {
	return UTF16LEEncoding.NewDecoder()
}

// EncodeToXMLFunc ...
type EncodeToXMLFunc func(w io.Writer, v interface{}) (int64, error)

// EncodeToUTF8XML ...
func EncodeToUTF8XML(w io.Writer, v interface{}) (int64, error) {
	return encodeToXML(NewUTF8Encoder().Writer(w), UTF8XMLHeader, v)
}

// EncodeToUTF16BEXML ...
func EncodeToUTF16BEXML(w io.Writer, v interface{}) (int64, error) {
	return encodeToXML(NewUTF16BEEncoder().Writer(w), UTF16XMLHeader, v)
}

// EncodeToUTF16LEXML ...
func EncodeToUTF16LEXML(w io.Writer, v interface{}) (int64, error) {
	return encodeToXML(NewUTF16LEEncoder().Writer(w), UTF16XMLHeader, v)
}

func encodeToXML(w io.Writer, header string, v interface{}) (int64, error) {
	raw, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return 0, err
	}
	n, err := w.Write(append([]byte(header), raw...))
	return int64(n), err
}
