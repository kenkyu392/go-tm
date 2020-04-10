package xliff

import (
	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/language"
)

// Option ...
type Option func(x *XLIFF)

// UseUTF8XMLEncodingOption ...
func UseUTF8XMLEncodingOption() Option {
	return func(x *XLIFF) {
		x.SetEncoding(internal.UTF8Encoding)
	}
}

// UseUTF16BEXMLEncodingOption ...
func UseUTF16BEXMLEncodingOption() Option {
	return func(x *XLIFF) {
		x.SetEncoding(internal.UTF16BEEncoding)
	}
}

// UseUTF16LEXMLEncodingOption ...
func UseUTF16LEXMLEncodingOption() Option {
	return func(x *XLIFF) {
		x.SetEncoding(internal.UTF16LEEncoding)
	}
}

// SourceLanguageOption ...
func SourceLanguageOption(tag language.Tag) Option {
	return func(x *XLIFF) {
		x.File.SourceLanguage = tag.String()
	}
}

// TargetLanguageOption ...
func TargetLanguageOption(tag language.Tag) Option {
	return func(x *XLIFF) {
		x.File.TargetLanguage = tag.String()
	}
}

// DataTypeOption ...
func DataTypeOption(dataType string) Option {
	return func(x *XLIFF) {
		x.File.DataType = dataType
	}
}

// ToolOption ...
func ToolOption(id, name, version, company string) Option {
	return func(x *XLIFF) {
		x.File.Header.Tool.ID = id
		x.File.Header.Tool.Name = name
		x.File.Header.Tool.Version = version
		x.File.Header.Tool.Company = company
	}
}
