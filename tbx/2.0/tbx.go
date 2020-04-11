package tbx

import (
	"encoding/xml"
	"errors"

	"github.com/kenkyu392/go-tm"
	"github.com/kenkyu392/go-tm/internal"
)

// const ...
const (
	Version               = "2.0"
	FileExtension         = "tbx"
	MIMEType              = "application/x-tbx+xml"
	Type                  = "TBX"
	DescripTypeDefinition = "definition"
)

// New ...
func New(opts ...Option) (*TBX, error) {
	tbx := &TBX{
		Type:    Type,
		XMLLang: "",
		Header: Header{
			FileDescription: FileDescription{
				TitleStatement:    "",
				SourceDescription: "",
			},
		},
		Body: Body{
			TermEntryList: make([]*TermEntry, 0),
		},
	}
	tbx.SetEncoding(internal.UTF16BEEncoding)
	for _, opt := range opts {
		opt(tbx)
	}
	return tbx, nil
}

// TBX ...
type TBX struct {
	XMLName xml.Name `xml:"martif"`
	XMLLang string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Type    string   `xml:"type,attr"`
	Header  Header   `xml:"martifHeader"`
	Body    Body     `xml:"text>body"`

	tm.TM
}

// Header is a grouping element that contains child elements which describe the TBX document instance.
type Header struct {
	XMLName         xml.Name        `xml:"martifHeader"`
	FileDescription FileDescription `xml:"fileDesc"`
}

// FileDescription is a nesting element containing child elements that describe the TBX document instance.
type FileDescription struct {
	XMLName xml.Name `xml:"fileDesc"`
	// TitleStatement is a nesting element containing the title and any notes about the TBX document instance.
	TitleStatement string `xml:"titleStmt>title"`
	// sourceDesc is a information about the source of the TBX document instance.
	SourceDescription string `xml:"sourceDesc>p"`
}

// Body is a nesting element that contains terminological entries (<termEntry>).
type Body struct {
	XMLName       xml.Name     `xml:"body"`
	TermEntryList []*TermEntry `xml:"termEntry"`
}

// TermEntry is a root element of a terminological entry. It shall contain at least one language section.
type TermEntry struct {
	XMLName     xml.Name   `xml:"termEntry"`
	ID          string     `xml:"id,attr"`
	LangSetList []*LangSet `xml:"langSet"`
}

// LangSet is a nesting element that contains all the information in a terminological entry pertaining to one language,
// including all the <tig> or <ntig>elements (terms and associated information) for that language.
type LangSet struct {
	XMLName xml.Name `xml:"langSet"`
	XMLLang string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	// <descripGrp>: Contains one <descrip> element as well as additional child elements for associated administrative information.
	Descrip Descrip `xml:"descripGrp>descrip"`
	// <ntig>: A nesting term information group. A grouping element that contains child elements describing a term.
	// It is equivalent to a <tig> except that it allows the description of term components.
	TermGroup TermGroup `xml:"ntig>termGrp"`
}

// Descrip ia an element that contains descriptive information about a concept, or relations to other concepts.
// The type of information that the element contains, and any restrictions on the permissible values of the element,
// are determined by the value of the type attribute.
type Descrip struct {
	XMLName xml.Name `xml:"descrip"`
	XMLText string   `xml:",innerxml"`
	Type    string   `xml:"type,attr"`
}

// TermGroup is a grouping element that contains one <termNote> element,
// and auxiliary information such as administrative information,
// transaction information, notes and cross references.
// It does not allow any <descrip> elements.
// It can only appear at the term (tig or ntig) level and below.
type TermGroup struct {
	XMLName  xml.Name `xml:"termGrp"`
	Term     Term     `xml:"term"`
	TermNote TermNote `xml:"termNote"`
}

// Term that is being described in a <tig> or <ntig>. This element,
// as well as other term-like elements such as those mentioned in the table Types of terms,
// relations to terms in section 9, can contain a <hi> element to allow a limited amount of inline markup.
// This is intended to handle markup requirements in special cases such as may be required for terms that represent scientific concepts.
// However, it is strongly recommended to use inline markup only when necessary to represent the term in its base form.
// Do not use the <hi> element for presentational styles chosen for esthetic purposes.
type Term struct {
	XMLName xml.Name `xml:"term"`
	XMLText string   `xml:",innerxml"`
	ID      string   `xml:"id,attr"`
}

// TermNote is a meta data-category used for describing terms.
// A type attribute specifies what kind of information is included in a particular instance of this element.
type TermNote struct {
	XMLName xml.Name `xml:"termNote"`
	XMLText string   `xml:",innerxml"`
	Type    string   `xml:"type,attr"`
}

// AddTermEntry ...
func (t *TBX) AddTermEntry(id string, list ...*LangSet) error {
	for _, item := range list {
		if item == nil {
			return errors.New("tbx: langSet is empty")
		}
		if item.XMLLang == "" {
			item.XMLLang = t.XMLLang
		}
	}
	t.Body.TermEntryList = append(t.Body.TermEntryList, &TermEntry{
		ID: id, LangSetList: list,
	})
	return nil
}
