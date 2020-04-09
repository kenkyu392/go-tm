package tmx

import (
	"encoding/xml"
	"errors"
	"io"
	"time"

	"golang.org/x/text/language"
)

// const ...
const (
	Version                                = "1.4"
	DefaultXMLNS                           = "http://www.lisa.org/tmx14"
	DefaultXMLHeader                       = "<?xml version=\"1.0\" encoding=\"UTF-16\"?>\n"
	DefaultCreationTool                    = "encoding/tmx"
	DefaultCreationToolVersion             = "1.0.0"
	DefaultDataType                        = "PlainText"
	DefaultSegmentType                     = "sentence"
	DefaultOriginalTranslationMemoryFormat = "encoding/tmx"
	DefaultUser                            = "anonymous"
	TimeFormat                             = "20060102T150405Z"
)

// New : [Translation Memory eXchange](https://en.wikipedia.org/wiki/Translation_Memory_eXchange)
func New(sourceTag, targetTag language.Tag) *TMX {
	return &TMX{
		XMLNS:   DefaultXMLNS,
		Version: Version,
		Header: Header{
			CreationTool:                    DefaultCreationTool,
			CreationToolVersion:             DefaultCreationToolVersion,
			DataType:                        DefaultDataType,
			SegmentType:                     DefaultSegmentType,
			OriginalTranslationMemoryFormat: DefaultOriginalTranslationMemoryFormat,
			SourceLang:                      sourceTag.String(),
			AdminLang:                       targetTag.String(),
		},
	}
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
	OriginalTranslationMemoryFormat string `xml:"o-tmf,attr"`
	// The srclang attribute specifies the language or locale of the source text. Its value must be one of the values used by a xml:lang attribute or the value "*all*" to indicate that any language combination can be used.
	SourceLang string `xml:"srclang,attr"`
	// The adminlang attribute is used in the <header> element to specify the default language for the administrative and informative elements <note> and <prop>. Its value must be one of the values used by a xml:lang attribute.
	AdminLang string `xml:"adminlang,attr"`
	// The creationdate attribute specifies the date of the creation of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	CreationDate string `xml:"creationdate,attr"`
	// The creationid attribute specifies the user who created the element.
	CreationID string `xml:"creationid,attr"`
	// The changedate attribute specifies the date of the modification of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	ChangeDate string `xml:"changedate,attr"`
	// The changeid attribute specifies the user who modified the element.
	ChangeID string `xml:"changeid,attr"`
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
	CreationDate string `xml:"creationdate,attr"`
	// The creationid attribute specifies the user who created the element.
	CreationID string `xml:"creationid,attr"`
	// The changedate attribute specifies the date of the modification of the element.
	// Its value must be in ASCII, in the format YYYYMMDDThhmmssZ.
	// (e.g. 19970811T133402Z for August 11th 1997 at 1:34pm 2 seconds.)
	// This is one of the options described in ISO 8601:1988. The value is always given in UTC (as indicated by the terminal Z).
	ChangeDate string `xml:"changedate,attr"`
	// The changeid attribute specifies the user who modified the element.
	ChangeID string `xml:"changeid,attr"`
}

// TMX ...
type TMX struct {
	XMLName xml.Name `xml:"tmx"`
	XMLNS   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	Header  Header
	Body    Body
}

// WriteTo is io.WriterTo interface implements.
func (t *TMX) WriteTo(w io.Writer) (int64, error) {
	raw, err := xml.MarshalIndent(t, "", "  ")
	if err != nil {
		return 0, err
	}
	n, err := w.Write(append([]byte(DefaultXMLHeader), raw...))
	return int64(n), err
}

// Bytes ...
func (t *TMX) Bytes() ([]byte, error) {
	contents, err := xml.MarshalIndent(t, "", "  ")
	if err != nil {
		return nil, err
	}
	return append([]byte(DefaultXMLHeader), contents...), nil
}

// SetToolInfo ...
func (t *TMX) SetToolInfo(name, version string) {
	t.Header.CreationTool = name
	t.Header.CreationToolVersion = version
}

// AddTU ...
func (t *TMX) AddTU(source *TUV, target *TUV) error {
	if source == nil {
		return errors.New("tmx: source is empty")
	}
	if target == nil {
		return errors.New("tmx: target is empty")
	}
	now := time.Now()
	setDefaultToTUV(source, now)
	setDefaultToTUV(target, now)
	if source.XMLLang == "" {
		source.XMLLang = t.Header.SourceLang
	}
	if target.XMLLang == "" {
		target.XMLLang = t.Header.AdminLang
	}
	t.Body.TUList = append(t.Body.TUList, &TU{
		TUVList: []*TUV{source, target},
	})
	return nil
}

func setDefaultToTUV(tuv *TUV, now time.Time) {
	now = now.UTC()
	if tuv.CreationDate == "" {
		tuv.CreationDate = now.Format(TimeFormat)
	}
	if tuv.CreationID == "" {
		tuv.CreationID = DefaultUser
	}
	if tuv.ChangeDate == "" {
		tuv.ChangeDate = now.Format(TimeFormat)
	}
	if tuv.ChangeID == "" {
		tuv.ChangeID = DefaultUser
	}
}
