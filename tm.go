package tm

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/encoding"
)

// Document ...
type Document interface {
	Encoding() encoding.Encoding
	SetEncoding(enc encoding.Encoding)
}

// TM ..
type TM struct {
	encoding encoding.Encoding
}

// Encoding is Document interface implements.
func (t *TM) Encoding() encoding.Encoding {
	return t.encoding
}

// SetEncoding is Document interface implements.
func (t *TM) SetEncoding(e encoding.Encoding) {
	t.encoding = e
}

// Encode returns XML encoded v.
func Encode(v Document, opts ...EncodeOption) ([]byte, error) {
	var b bytes.Buffer
	if err := EncodeTo(&b, v, opts...); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// EncodeTo writes XML encoded v to w.
func EncodeTo(w io.Writer, v Document, opts ...EncodeOption) error {
	var (
		encOpts = new(EncodeOptions)
		eng     = v.Encoding()
		ew      = eng.NewEncoder().Writer(w)
	)
	for _, opt := range opts {
		opt(encOpts)
	}
	if encOpts.header {
		if _, err := ew.Write([]byte(internal.XMLHeaderByEncoding(eng))); err != nil {
			return err
		}
	}
	enc := xml.NewEncoder(ew)
	enc.Indent(encOpts.prefix, encOpts.indent)
	if err := enc.Encode(v); err != nil {
		return err
	}
	return nil
}

// EncodeOptions ...
type EncodeOptions struct {
	indent string
	prefix string
	header bool
}

// EncodeOption ...
type EncodeOption func(opts *EncodeOptions)

// XMLEncodeOption ...
func XMLEncodeOption(prefix, indent string) EncodeOption {
	return func(opts *EncodeOptions) {
		opts.prefix = prefix
		opts.indent = indent
	}
}

// WithXMLHeaderEncodeOption ...
func WithXMLHeaderEncodeOption() EncodeOption {
	return func(opts *EncodeOptions) {
		opts.header = true
	}
}
