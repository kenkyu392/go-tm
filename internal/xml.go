package internal

import "golang.org/x/text/encoding"

const (
	// UTF8XMLHeader is the UTF-8 encoding header.
	UTF8XMLHeader = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	// UTF16XMLHeader is the UTF-16 encoding header.
	UTF16XMLHeader = `<?xml version="1.0" encoding="UTF-16"?>` + "\n"
)

// XMLHeaderByEncoding ...
func XMLHeaderByEncoding(e encoding.Encoding) string {
	switch e {
	case UTF16BEEncoding:
		fallthrough
	case UTF16LEEncoding:
		return UTF16XMLHeader
	case UTF8Encoding:
		fallthrough
	default:
		return UTF8XMLHeader
	}
}
