package tmx

import (
	"bytes"
	"testing"

	"golang.org/x/text/language"
)

func TestNew(t *testing.T) {
	jaJP := language.MustParse("ja-JP")
	enUS := language.MustParse("en-US")
	frFR := language.MustParse("fr-FR")

	tmx := New(jaJP, enUS)
	tmx.AddTUPair(
		NewTUV(
			DefaultTUVOption(),
			SegmentTUVOption("吾輩は猫である"),
		),
		NewTUV(
			DefaultTUVOption(),
			SegmentTUVOption("I Am a Cat"),
		),
	)
	tmx.AddTU(
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

	t.Run("case=Bytes", func(t *testing.T) {
		raw, err := tmx.Bytes()
		if err != nil {
			t.Errorf("tmx: %+v", err)
		}
		t.Logf("%s\n", raw)
	})
	t.Run("case=WriteTo", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if _, err := tmx.WriteTo(buf); err != nil {
			t.Errorf("tmx: %+v", err)
		}
		t.Logf("%s\n", buf.Bytes())
	})
}
