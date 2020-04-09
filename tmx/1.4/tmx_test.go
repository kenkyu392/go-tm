package tmx

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestNew(t *testing.T) {
	jaJP := language.MustParse("ja-JP")
	enUS := language.MustParse("en-US")
	frFR := language.MustParse("fr-FR")
	tmx := New(
		SourceLangOption(jaJP),
		UseUTF8XMLEncodingOption(),
	)
	tmx.AddTU(
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

	t.Run("case=WriteTo", func(t *testing.T) {
		buf := &bytes.Buffer{}
		if _, err := tmx.WriteTo(buf); err != nil {
			t.Errorf("tmx: %+v", err)
		}
		fmt.Printf("%s\n", buf.Bytes())
	})
}
