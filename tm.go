package tm

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/encoding"
)

// Message ...
type Message interface {
	Encoding() encoding.Encoding
	SetEncoding(enc encoding.Encoding)
}

// TM ..
type TM struct {
	encoding encoding.Encoding
}

// Encoding is Message interface implements.
func (t *TM) Encoding() encoding.Encoding {
	return t.encoding
}

// SetEncoding Message interface implements.
func (t *TM) SetEncoding(e encoding.Encoding) {
	t.encoding = e
}

// Encode ...
func Encode(v Message) ([]byte, error) {
	var b bytes.Buffer
	if err := EncodeTo(&b, v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// EncodeTo ...
func EncodeTo(w io.Writer, v Message) error {
	var (
		e = v.Encoding()
		h = internal.XMLHeaderByEncoding(e)
	)
	ew := e.NewEncoder().Writer(w)
	if _, err := ew.Write([]byte(h)); err != nil {
		return err
	}
	enc := xml.NewEncoder(ew)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		return err
	}
	return nil
}
