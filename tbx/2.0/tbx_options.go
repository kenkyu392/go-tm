package tbx

import (
	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/language"
)

// Option ...
type Option func(tbx *TBX)

// UseUTF8XMLEncodingOption ...
func UseUTF8XMLEncodingOption() Option {
	return func(tbx *TBX) {
		tbx.SetEncoding(internal.UTF8Encoding)
	}
}

// UseUTF16BEXMLEncodingOption ...
func UseUTF16BEXMLEncodingOption() Option {
	return func(tbx *TBX) {
		tbx.SetEncoding(internal.UTF16BEEncoding)
	}
}

// UseUTF16LEXMLEncodingOption ...
func UseUTF16LEXMLEncodingOption() Option {
	return func(tbx *TBX) {
		tbx.SetEncoding(internal.UTF16LEEncoding)
	}
}

// LanguageOption ...
func LanguageOption(tag language.Tag) Option {
	return func(tbx *TBX) {
		tbx.XMLLang = tag.String()
	}
}

// FileDescriptionOption ...
func FileDescriptionOption(titleStmt, sourceDesc string) Option {
	return func(tbx *TBX) {
		tbx.Header.FileDescription.TitleStatement = titleStmt
		tbx.Header.FileDescription.SourceDescription = sourceDesc
	}
}
