package tbx

import (
	"bytes"
	"testing"

	"github.com/kenkyu392/go-tm"
)

func TestNew(t *testing.T) {
	tbx, err := New(
		UseUTF8XMLEncodingOption(),
		LanguageOption(tm.Tag_enUS),
		FileDescriptionOption(
			"Japanese Books",
			"Japanese Books",
		),
	)
	if err != nil {
		t.Errorf("tbx: %+v", err)
	}
	tbx.AddTermEntry(
		"1905",
		[]*LangSet{
			{
				XMLLang: tm.Tag_enUS.String(),
				Descrip: Descrip{
					XMLText: `"I Am a Cat" is a satirical novel written in 1905–1906 by Natsume Sōseki about Japanese society during the Meiji period (1868–1912).`,
					Type:    DescripTypeDefinition,
				},
				TermGroup: TermGroup{
					Term:     Term{XMLText: "I Am a Cat", ID: "10"},
					TermNote: TermNote{XMLText: "Noun", Type: "partOfSpeech"},
				},
			},
			{
				XMLLang: tm.Tag_jaJP.String(),
				Descrip: Descrip{
					XMLText: `「私は猫です」は、1905〜1906年に夏目漱石が書いた明治時代（1868〜1912）の日本社会についての風刺小説です。`,
					Type:    DescripTypeDefinition,
				},
				TermGroup: TermGroup{
					Term:     Term{XMLText: "吾輩は猫である", ID: "20"},
					TermNote: TermNote{XMLText: "Noun", Type: "partOfSpeech"},
				},
			},
		}...,
	)

	t.Run("case=Encode", func(t *testing.T) {
		raw, err := tm.Encode(tbx)
		if err != nil {
			t.Errorf("tbx: %+v", err)
		}
		t.Logf("%s\n", raw)
	})
	t.Run("case=EncodeTo", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := tm.EncodeTo(buf, tbx); err != nil {
			t.Errorf("tbx: %+v", err)
		}
		t.Logf("%s\n", buf.Bytes())
	})
}
