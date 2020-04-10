package xliff

import (
	"encoding/xml"
	"errors"
	"time"

	"github.com/kenkyu392/go-tm"
	"github.com/kenkyu392/go-tm/internal"
)

// const ...
const (
	Version                  = "1.2"
	FileExtension            = "xlf"
	MIMEType                 = "application/x-xliff+xml"
	DefaultXMLNS             = "urn:oasis:names:tc:xliff:document:1.2"
	DefaultFileOriginal      = "original.xlf"
	DefaultFileDataType      = "plaintext"
	DefaultHeaderToolID      = "go-tmx"
	DefaultHeaderToolName    = "go-tmx"
	DefaultHeaderToolVersion = "1.0.0"
	DefaultXMLSpace          = "preserve"
)

// New : [XML Localisation Interchange File Format](https://en.wikipedia.org/wiki/XLIFF)
// http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html
// http://docs.oasis-open.org/xliff/xliff-core/xliff-core.html
func New(opts ...Option) (*XLIFF, error) {
	xlf := &XLIFF{
		XMLNS:    DefaultXMLNS,
		XMLSpace: DefaultXMLSpace,
		Version:  Version,
		File: File{
			Date:           time.Now().UTC().Format(time.RFC3339),
			DataType:       DefaultFileDataType,
			Original:       DefaultFileOriginal,
			SourceLanguage: "",
			TargetLanguage: "",
			Header: Header{
				Tool: Tool{
					ID:      DefaultHeaderToolID,
					Name:    DefaultHeaderToolName,
					Version: DefaultHeaderToolVersion,
				},
			},
		},
	}
	xlf.SetEncoding(internal.UTF16BEEncoding)
	for _, opt := range opts {
		opt(xlf)
	}
	if xlf.File.SourceLanguage == "" {
		return nil, errors.New("xliff: source language is empty")
	}
	if xlf.File.TargetLanguage == "" {
		return nil, errors.New("xliff: target language is empty")
	}
	return xlf, nil
}

// XLIFF ...
type XLIFF struct {
	XMLName  xml.Name `xml:"xliff"`
	XMLNS    string   `xml:"xmlns,attr"`
	XMLSpace string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
	Version  string   `xml:"version,attr"`
	File     File

	tm.TM
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
	Version string   `xml:"tool-version,attr,omitempty"`
	Company string   `xml:"tool-company,attr,omitempty"`
}

// Body : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#body
type Body struct {
	XMLName       xml.Name `xml:"body"`
	TransUnitList []*TransUnit
}

// TransUnit : http://docs.oasis-open.org/xliff/v1.2/os/xliff-core.html#trans-unit
type TransUnit struct {
	XMLName  xml.Name `xml:"trans-unit"`
	XMLSpace string   `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
	ID       string   `xml:"id,attr,omitempty"`
	// Source text - The <source> element is used to delimit a unit of text that could be a paragraph,
	// a title, a menu item, a caption, etc. The content of the <source> is generally the translatable text,
	// depending upon the translate attribute of the parent <trans-unit>.
	// The optional xml:lang attribute is used to specify the content language of the <source>;
	// this should always match source-language as a child of <trans-unit> but can vary as a child of <alt-trans>.
	// The optional ts attribute was DEPRECATED in XLIFF 1.1.
	Source *Segment `xml:"source"`
	// Target - The <target> element contains the translation of the content of the sibling <source> element.
	// The optional state and state-qualifier attributes indicate in which state the <target> is.
	// The optional phase-name attribute references the <phase> in which the <target> was last modified.
	// The optional xml:lang attribute is used to specify the content language of the <target>;
	// this should always match target-language as a child of <trans-unit> but can vary as a child of <alt-trans>.
	// The optional coord, font, css-style, style, and exstyle attributes describe the resource contained within the <target>;
	// these are the modifiable attributes for the <trans-unit> depending upon the reformat attribute of the parent <trans-unit>.
	// The optional equiv-trans describes if the target language translation is a direct equivalent of the source text.
	// The optional ts attribute was DEPRECATED in XLIFF 1.1.
	// The restype attribute is DEPRECATED in XLIFF 1.2, since <target> will always be of the same restype as its parent <trans-unit> or <alt-trans>.
	// A list of preferred values for the restype, state, and state-qualifier attributes are provided by this specification.
	Target *Segment `xml:"target"`
	// Note - The <note> element is used to add localization-related comments to the XLIFF document.
	// The content of <note> may be instructions from developers about how to handle the <source>,
	// comments from the translator about the translation, or any comment from anyone involved in processing the XLIFF file.
	// The optional xml:lang attribute specifies the language of the note content.
	// The optional from attribute indicates who entered the note.
	// The optional priority attribute allows a priority from 1 (high) to 10 (low) to be assigned to the note.
	// The optional annotates attribute indicates if the note is a general note or, in the case of a <trans-unit>,
	// pertains specifically to the <source> or the <target> element.
	Note []*Segment `xml:"note"`
	// Translate (yes or no) - The translate attribute indicates whether or not the text referred to should be translated.
	// Approved (yes or no) - Indicates whether a translation is final or has passed its final review.
}

// Segment ...
type Segment struct {
	Lang    string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Content string `xml:",innerxml"`
}

// Unit ...
type Unit struct {
	Text  string
	Notes []string
}

// AddTransUnit ...
func (x *XLIFF) AddTransUnit(id string, source, target *Unit) error {
	if source == nil {
		return errors.New("xliff: source is empty")
	}
	if target == nil {
		return errors.New("xliff: target is empty")
	}

	tu := &TransUnit{
		XMLSpace: DefaultXMLSpace,
		ID:       id,
		Source: &Segment{
			Lang:    x.File.SourceLanguage,
			Content: source.Text,
		},
		Target: &Segment{
			Lang:    x.File.TargetLanguage,
			Content: target.Text,
		},
		Note: make([]*Segment, 0),
	}

	for _, note := range source.Notes {
		tu.Note = append(tu.Note, &Segment{
			Lang:    x.File.SourceLanguage,
			Content: note,
		})
	}

	for _, note := range target.Notes {
		tu.Note = append(tu.Note, &Segment{
			Lang:    x.File.TargetLanguage,
			Content: note,
		})
	}

	x.File.Body.TransUnitList = append(x.File.Body.TransUnitList, tu)
	return nil
}
