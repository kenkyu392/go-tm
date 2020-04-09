package xliff

import (
	"encoding/xml"

	"golang.org/x/text/language"
)

// const ...
const (
	Version                  = "2.0"
	DefaultXMLNS             = "urn:oasis:names:tc:xliff:document:2.0"
	DefaultXMLHeader         = "<?xml version=\"1.0\" encoding=\"UTF-16\"?>\n"
	DefaultFileOriginal      = "encoding/2.0/xliff"
	DefaultFileDataType      = "plaintext"
	DefaultHeaderToolID      = "encoding/2.0/xliff"
	DefaultHeaderToolName    = "encoding/2.0/xliff"
	DefaultHeaderToolVersion = "1.0.0"
	DefaultXMLSpace          = "preserve"
)

// New : [XML Localisation Interchange File Format](https://en.wikipedia.org/wiki/XLIFF)
// http://docs.oasis-open.org/xliff/xliff-core/v2.0/xliff-core-v2.0.html
func New(id string, sourceTag, targetTag language.Tag) *XLIFF {
	return &XLIFF{
		XMLNS:    DefaultXMLNS,
		XMLSpace: DefaultXMLSpace,
		Version:  Version,
		SrcLang:  sourceTag.String(),
		TrgLang:  targetTag.String(),
		File: File{
			ID: id,
		},
	}
}

// File : http://docs.oasis-open.org/xliff/xliff-core/v2.0/xliff-core-v2.0.html#file
type File struct {
	XMLName  xml.Name `xml:"file"`
	ID       string   `xml:"id,attr"`
	UnitList []Unit
}

// Unit : http://docs.oasis-open.org/xliff/xliff-core/v2.0/xliff-core-v2.0.html#unit
type Unit struct {
	XMLName  xml.Name `xml:"unit"`
	XMLSpace string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
	ID       string   `xml:"id,attr,omitempty"`
	// segment : http://docs.oasis-open.org/xliff/xliff-core/v2.0/xliff-core-v2.0.html#segment
	Source string   `xml:"segment>source"`
	Target string   `xml:"segment>target"`
	Notes  []string `xml:"notes>note"`
}

// XLIFF ...
type XLIFF struct {
	XMLName  xml.Name `xml:"xliff"`
	XMLNS    string   `xml:"xmlns,attr"`
	XMLSpace string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
	Version  string   `xml:"version,attr"`
	SrcLang  string   `xml:"srcLang,attr"`
	TrgLang  string   `xml:"trgLang,attr"`
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

// Store ...
func (x *XLIFF) Store(id, sourceText string, targetText string, notes ...string) {
	x.File.UnitList = append(x.File.UnitList, Unit{
		XMLSpace: DefaultXMLSpace,
		ID:       id,
		Source:   sourceText,
		Target:   targetText,
		Notes:    notes,
	})
}
