package tmx

import (
	"time"

	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/language"
)

// Option ...
type Option func(tmx *TMX)

// UseUTF8XMLEncodingOption ...
func UseUTF8XMLEncodingOption() Option {
	return func(tmx *TMX) {
		tmx.SetEncoding(internal.UTF8Encoding)
	}
}

// UseUTF16BEXMLEncodingOption ...
func UseUTF16BEXMLEncodingOption() Option {
	return func(tmx *TMX) {
		tmx.SetEncoding(internal.UTF16BEEncoding)
	}
}

// UseUTF16LEXMLEncodingOption ...
func UseUTF16LEXMLEncodingOption() Option {
	return func(tmx *TMX) {
		tmx.SetEncoding(internal.UTF16LEEncoding)
	}
}

// OriginalTranslationMemoryFormatOption ...
func OriginalTranslationMemoryFormatOption(otmf string) Option {
	return func(tmx *TMX) {
		tmx.Header.OriginalTranslationMemoryFormat = otmf
	}
}

// DataTypeOption ...
func DataTypeOption(dataType string) Option {
	return func(tmx *TMX) {
		tmx.Header.DataType = dataType
	}
}

// SegmentTypeOption ...
func SegmentTypeOption(segmentType string) Option {
	return func(tmx *TMX) {
		tmx.Header.SegmentType = segmentType
	}
}

// SourceLangOption ...
func SourceLangOption(tag language.Tag) Option {
	return func(tmx *TMX) {
		tmx.Header.SourceLang = tag.String()
	}
}

// AdminLangOption ...
func AdminLangOption(tag language.Tag) Option {
	return func(tmx *TMX) {
		tmx.Header.AdminLang = tag.String()
	}
}

// CreationToolOption ...
func CreationToolOption(name, version string) Option {
	return func(tmx *TMX) {
		tmx.Header.CreationTool = name
		tmx.Header.CreationToolVersion = Version
	}
}

// CreationOption ...
func CreationOption(t time.Time, id string) Option {
	return func(tmx *TMX) {
		tmx.Header.CreationDate = t.UTC().Format(TimeFormat)
		tmx.Header.CreationID = id
	}
}

// ChangeOption ...
func ChangeOption(t time.Time, id string) Option {
	return func(tmx *TMX) {
		tmx.Header.ChangeDate = t.UTC().Format(TimeFormat)
		tmx.Header.ChangeID = id
	}
}
