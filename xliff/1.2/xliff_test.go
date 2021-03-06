package xliff

import (
	"bytes"
	"testing"

	"github.com/kenkyu392/go-tm"
)

func TestNew(t *testing.T) {
	xlf, err := New(
		UseUTF8XMLEncodingOption(),
		SourceLanguageOption(tm.Tag_jaJP),
		TargetLanguageOption(tm.Tag_enUS),
	)
	if err != nil {
		t.Errorf("xliff: %+v", err)
	}
	xlf.AddTransUnit(
		"ID001",
		&Unit{
			Text: "吾輩は猫である",
			Notes: []string{
				"著者：夏目漱石",
				"発行日：1905年～1906年",
			},
		},
		&Unit{
			Text: "I Am a Cat",
			Notes: []string{
				"Author‎: ‎Natsume Sōseki",
				"Publication date: 1905–1906",
			},
		},
	)
	xlf.AddTransUnit(
		"ID002",
		&Unit{
			Text: "人間失格",
			Notes: []string{
				"著者：太宰治",
				"発行日：1948年",
			},
		},
		&Unit{
			Text: "No Longer Human",
			Notes: []string{
				"Author‎: Osamu Dazai",
				"Publication date: 1948",
			},
		},
	)

	opts := []tm.EncodeOption{
		tm.WithXMLHeaderEncodeOption(),
		tm.XMLEncodeOption("", "  "),
	}
	t.Run("case=Encode", func(t *testing.T) {
		raw, err := tm.Encode(xlf, opts...)
		if err != nil {
			t.Errorf("xliff: %+v", err)
		}
		t.Logf("%s\n", raw)
	})
	t.Run("case=EncodeTo", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := tm.EncodeTo(buf, xlf, opts...); err != nil {
			t.Errorf("xliff: %+v", err)
		}
		t.Logf("%s\n", buf.Bytes())
	})
}
