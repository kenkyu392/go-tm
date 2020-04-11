package internal

import "golang.org/x/text/encoding"

const (
	// UTF8XMLHeader is the UTF-8 encoding header.
	UTF8XMLHeader = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n"
	// UTF16XMLHeader is the UTF-16 encoding header.
	UTF16XMLHeader = `<?xml version="1.0" encoding="UTF-16" standalone="yes"?>` + "\n"
	// XMLSpacePreserve is tells the user agent to convert all newline and tab characters into spaces.
	// Then, it draws all space characters (including leading, trailing and multiple consecutive space characters).
	// Example: xml:space="preserve"
	XMLSpacePreserve = "preserve"
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
