package tmx

import (
	"bytes"
	"testing"

	"github.com/kenkyu392/go-tm"
	"golang.org/x/text/language"
)

func TestNew(t *testing.T) {
	jaJP := language.MustParse("ja-JP")
	enUS := language.MustParse("en-US")
	frFR := language.MustParse("fr-FR")
	tmx, err := New(
		SourceLangOption(jaJP),
		UseUTF8XMLEncodingOption(),
	)
	if err != nil {
		t.Errorf("tmx: %+v", err)
	}
	tmx.AddTU(
		"ID001",
		NewTUV(
			DefaultTUVOption(),
			XMLLangTUVOption(jaJP),
			SegmentTUVOption("吾輩は猫である"),
		),
		NewTUV(
			DefaultTUVOption(),
			XMLLangTUVOption(enUS),
			SegmentTUVOption("I Am a Cat"),
		),
	)
	tmx.AddTU(
		"ID001",
		NewTUV(
			DefaultTUVOption(),
			XMLLangTUVOption(jaJP),
			SegmentTUVOption("人間失格"),
		),
		NewTUV(
			DefaultTUVOption(),
			XMLLangTUVOption(enUS),
			SegmentTUVOption("No Longer Human"),
		),
		NewTUV(
			DefaultTUVOption(),
			XMLLangTUVOption(frFR),
			SegmentTUVOption("La déchéance d'un homme"),
		),
	)

	t.Run("case=Encode", func(t *testing.T) {
		raw, err := tm.Encode(tmx)
		if err != nil {
			t.Errorf("tmx: %+v", err)
		}
		t.Logf("%s\n", raw)
	})
	t.Run("case=EncodeTo", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if err := tm.EncodeTo(buf, tmx); err != nil {
			t.Errorf("tmx: %+v", err)
		}
		t.Logf("%s\n", buf.Bytes())
	})
}
