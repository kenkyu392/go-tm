package tmx

import (
	"encoding/xml"
	"errors"
	"io"
	"time"

	"github.com/kenkyu392/go-tm/internal"
	"golang.org/x/text/language"
)

// const ...
const (
	Version                                = "1.4"
	FileExtension                          = "tmx"
	MIMEType                               = "application/x-tmx+xml"
	DefaultXMLNS                           = "http://www.lisa.org/tmx14"
	DefaultCreationTool                    = "go-tmx"
	DefaultCreationToolVersion             = "1.0.0"
	DefaultDataType                        = "PlainText"
	DefaultSegmentType                     = "sentence"
	DefaultOriginalTranslationMemoryFormat = "GO TMX"
	DefaultUser                            = "anonymous"
	TimeFormat                             = "20060102T150405Z"
)

// New : [Translation Memory eXchange](https://en.wikipedia.org/wiki/Translation_Memory_eXchange)
func New(opts ...Option) *TMX {
	tmx := &TMX{
		XMLNS:   DefaultXMLNS,
		Version: Version,
		Header: Header{
			CreationTool:                    DefaultCreationTool,
			CreationToolVersion:             DefaultCreationToolVersion,
			DataType:                        DefaultDataType,
			SegmentType:                     DefaultSegmentType,
			OriginalTranslationMemoryFormat: DefaultOriginalTranslationMemoryFormat,
			SourceLang:                      "",
			AdminLang:                       "",
			CreationDate:                    "",
			CreationID:                      "",
			ChangeDate:                      "",
			ChangeID:                        "",
		},
		encodeToXML: internal.EncodeToUTF16BEXML,
	}
	for _, opt := range opts {
		opt(tmx)
	}
	return tmx
}

// Option ...
type Option func(tmx *TMX)

// UseUTF8XMLEncodingOption ...
func UseUTF8XMLEncodingOption() Option {
	return func(tmx *TMX) {
		tmx.encodeToXML = internal.EncodeToUTF8XML
	}
}

// UseUTF16BEXMLEncodingOption ...
func UseUTF16BEXMLEncodingOption() Option {
	return func(tmx *TMX) {
		tmx.encodeToXML = internal.EncodeToUTF16BEXML
	}
}

// UseUTF16LEXMLEncodingOption ...
func UseUTF16LEXMLEncodingOption() Option {
	return func(tmx *TMX) {
		tmx.encodeToXML = internal.EncodeToUTF16LEXML
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

// TMX ...
type TMX struct {
	XMLName xml.Name `xml:"tmx"`
	XMLNS   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	Header  Header
	Body    Body

	encodeToXML internal.EncodeToXMLFunc
}

// Header : The <header> element contains zero, one or more <note> elements; zero, one or more <ude> elements; and zero, one or more <prop> elements.
type Header struct {
	XMLName xml.Name `xml:"header"`
	// The creationtool attribute identifies the tool that created the TMX document. Its possible values are not specified by the standard but each tool provider will publish the string identifier it uses.
	CreationTool string `xml:"creationtool,attr"`
	// The creationtoolversion attribute identifies the version of the tool that created the TMX document. Its possible values are not specified by the standard but each tool provider will publish the string identifier it uses.
	CreationToolVersion string `xml:"creationtoolversion,attr"`
	// The datatype attribute specifies the type of data contained in an element. Its default value is "unknown". See the recommended values section for more information.
	DataType string `xml:"datatype,attr"`
	// The segtype attribute specifies the kind of segmentation used in the <tu> element. Its value must be either "block", "paragraph", "sentence" or "phrase". If a<tu> element does not have a segtype attribute specified, it is of the type defined in the <header> element.
	// See the Implementation Notes for examples of how to use segtype.
	SegmentType string `xml:"segtype,attr"`
	// The o-tmf (Original Translation Memory Format) element specifies the format of the Translation Memory file from which the TMX document or segment thereof have been generated.
	OriginalTranslationMemoryFormat string `xml:"o-tmf,attr,omitempty"`
	// The srclang attribute specifies the language or locale of the source text. Its value must be one of the values used by a xml:lang attribute or the value "*all*" to indicate that any language combination can be used.
	SourceLang string `xml:"srclang,attr,omitempty"`
	// The adminlang attribute is used in the <header> element to specify the default language for the administrative and informative elements <note> and <prop>. Its value must be one of the values used by a xml:lang attribute.
	AdminLang string `xml:"adminlang,attr,omitempty"`
	// The creationdate attribute specifies the date of the creation of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	CreationDate string `xml:"creationdate,attr,omitempty"`
	// The creationid attribute specifies the user who created the element.
	CreationID string `xml:"creationid,attr,omitempty"`
	// The changedate attribute specifies the date of the modification of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	ChangeDate string `xml:"changedate,attr,omitempty"`
	// The changeid attribute specifies the user who modified the element.
	ChangeID string `xml:"changeid,attr,omitempty"`
}

// Body : The <body> element encloses the main data, the set of <tu> elements that are comprised within the file.
type Body struct {
	XMLName xml.Name `xml:"body"`
	TUList  []*TU
}

// TU : Each <tu> (Translation Unit) element contains zero, one or more <note> elements or <prop> elements, followed by one or more <tuv> elements. Logically, a complete translation-memory database will contain at least two <tuv> elements in each Translation Unit.
type TU struct {
	XMLName xml.Name `xml:"tu"`
	TUVList []*TUV
}

// TUV : Each <tuv> (Translation Unit Variant) element specifies text in a given language. It contains zero, one or more <note> elements or <prop> elements, followed by one <seg> element.
type TUV struct {
	XMLName xml.Name `xml:"tuv"`
	XMLLang string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Segment string   `xml:"seg"`
	// The creationdate attribute specifies the date of the creation of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	CreationDate string `xml:"creationdate,attr,omitempty"`
	// The creationid attribute specifies the user who created the element.
	CreationID string `xml:"creationid,attr,omitempty"`
	// The changedate attribute specifies the date of the modification of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	ChangeDate string `xml:"changedate,attr,omitempty"`
	// The changeid attribute specifies the user who modified the element.
	ChangeID string `xml:"changeid,attr,omitempty"`
}

// WriteTo is io.WriterTo interface implements.
func (t *TMX) WriteTo(w io.Writer) (int64, error) {
	return t.encodeToXML(w, t)
}

// AddTU ...
func (t *TMX) AddTU(list ...*TUV) error {
	for _, tuv := range list {
		if tuv == nil {
			return errors.New("tmx: tuv is empty")
		}
		if tuv.XMLLang == "" {
			return errors.New("tmx: tuv lang is empty")
		}
	}
	t.Body.TUList = append(t.Body.TUList, &TU{
		TUVList: list,
	})
	return nil
}

// NewTUV ...
func NewTUV(opts ...TUVOption) *TUV {
	tuv := new(TUV)
	for _, opt := range opts {
		opt(tuv)
	}
	return tuv
}

// TUVOption ...
type TUVOption func(tuv *TUV)

// DefaultTUVOption ...
func DefaultTUVOption() TUVOption {
	return func(tuv *TUV) {
		now := time.Now()
		CreationTUVOption(now, DefaultUser)
		ChangeTUVOption(now, DefaultUser)
	}
}

// XMLLangTUVOption ...
func XMLLangTUVOption(tag language.Tag) TUVOption {
	return func(tuv *TUV) {
		tuv.XMLLang = tag.String()
	}
}

// SegmentTUVOption ...
func SegmentTUVOption(segment string) TUVOption {
	return func(tuv *TUV) {
		tuv.Segment = segment
	}
}

// CreationTUVOption ...
func CreationTUVOption(t time.Time, id string) TUVOption {
	return func(tuv *TUV) {
		tuv.CreationDate = t.UTC().Format(TimeFormat)
		tuv.CreationID = id
	}
}

// ChangeTUVOption ...
func ChangeTUVOption(t time.Time, id string) TUVOption {
	return func(tuv *TUV) {
		tuv.ChangeDate = t.UTC().Format(TimeFormat)
		tuv.ChangeID = id
	}
}
