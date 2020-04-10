package tm

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/encoding"
)

// TM ...
type TM interface {
	Encoding() encoding.Encoding
}

// Encode ...
func Encode(v TM) ([]byte, error) {
	var b bytes.Buffer
	if err := EncodeTo(&b, v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// EncodeTo ...
func EncodeTo(w io.Writer, v TM) error {
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
