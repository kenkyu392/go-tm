package tmx

import (
	"bytes"
	"testing"

	"golang.org/x/text/language"
)

func TestNew(t *testing.T) {
	tmx := New(language.MustParse("ja-JP"), language.MustParse("en-US"))

	tmx.AddTU(
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
			SegmentTUVOption("人間失格"),
		),
		NewTUV(
			DefaultTUVOption(),
			SegmentTUVOption("No Longer Human"),
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
