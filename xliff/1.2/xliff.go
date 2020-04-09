package xliff

import (
	"encoding/xml"
	"time"

	"golang.org/x/text/language"
)

// const ...
const (
	Version                  = "1.2"
	FileExtension            = "xlf"
	MIMEType                 = "application/x-xliff+xml"
	DefaultXMLNS             = "urn:oasis:names:tc:xliff:document:1.2"
	DefaultXMLHeader         = "<?xml version=\"1.0\" encoding=\"UTF-16\"?>\n"
	DefaultFileOriginal      = "encoding/1.2/xliff"
	DefaultFileDataType      = "plaintext"
	DefaultHeaderToolID      = "encoding/1.2/xliff"
	DefaultHeaderToolName    = "encoding/1.2/xliff"
	DefaultHeaderToolVersion = "1.0.0"
	DefaultXMLSpace          = "preserve"
)

// New : [XML Localisation Interchange File Format](https://en.wikipedia.org/wiki/XLIFF)
// http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html
// http://docs.oasis-open.org/xliff/xliff-core/xliff-core.html
func New(sourceTag, targetTag language.Tag) *XLIFF {
	return &XLIFF{
		XMLNS:    DefaultXMLNS,
		XMLSpace: DefaultXMLSpace,
		Version:  Version,
		File: File{
			Date:           time.Now().Format(time.RFC3339),
			DataType:       DefaultFileDataType,
			Original:       DefaultFileOriginal,
			SourceLanguage: sourceTag.String(),
			TargetLanguage: targetTag.String(),
			Header: Header{
				Tool: Tool{
					ID:      DefaultHeaderToolID,
					Name:    DefaultHeaderToolName,
					Version: DefaultHeaderToolVersion,
				},
			},
		},
	}
}

// File : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#file
type File struct {
	XMLName        xml.Name `xml:"file"`
	Date           string   `xml:"date,attr"`
	Original       string   `xml:"original,attr"`
	DataType       string   `xml:"datatype,attr"`
	SourceLanguage string   `xml:"source-language,attr"`
	TargetLanguage string   `xml:"target-language,attr"`
	Header         Header
	Body           Body
}

// Header : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#header
type Header struct {
	XMLName xml.Name `xml:"header"`
	Tool    Tool
}

// Tool : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#tool_elem
type Tool struct {
	XMLName xml.Name `xml:"tool"`
	ID      string   `xml:"tool-id,attr"`
	Name    string   `xml:"tool-name,attr"`
	Version string   `xml:"tool-version,attr"`
}

// Body : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#body
type Body struct {
	XMLName       xml.Name `xml:"body"`
	TransUnitList []TransUnit
}

// TransUnit : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#trans-unit
type TransUnit struct {
	XMLName  xml.Name `xml:"trans-unit"`
	XMLSpace string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
	ID       string   `xml:"id,attr,omitempty"`
	Source   Segment  `xml:"source"`
	Target   Segment  `xml:"target"`
	Note     []string `xml:"note"`
	// Translate (yes or no) - The translate attribute indicates whether or not the text referred to should be translated.
	// Approved (yes or no) - Indicates whether a translation is final or has passed its final review.
}

// Segment ...
type Segment struct {
	Lang    string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Content string `xml:",innerxml"`
}

// XLIFF ...
type XLIFF struct {
	XMLName  xml.Name `xml:"xliff"`
	XMLNS    string   `xml:"xmlns,attr"`
	XMLSpace string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
	Version  string   `xml:"version,attr"`
	File     File
}

// Marshal ...
func (x *XLIFF) Marshal() ([]byte, error) {
	bs := []byte(DefaultXMLHeader)
	contents, err := xml.MarshalIndent(x, "", "  ")
	if err != nil {
		return nil, err
	}
	bs = append(bs, contents...)
	return bs, nil
}

// SetToolInfo ...
func (x *XLIFF) SetToolInfo(id, name, version string) {
	x.File.Header.Tool.ID = id
	x.File.Header.Tool.Name = name
	x.File.Header.Tool.Version = version
}

// Store ...
func (x *XLIFF) Store(id, sourceText string, targetText string, notes ...string) {
	x.File.Body.TransUnitList = append(x.File.Body.TransUnitList, TransUnit{
		XMLSpace: DefaultXMLSpace,
		ID:       id,
		Source: Segment{
			Lang:    x.File.SourceLanguage,
			Content: sourceText,
		},
		Target: Segment{
			Lang:    x.File.TargetLanguage,
			Content: targetText,
		},
		Note: notes,
	})
}
