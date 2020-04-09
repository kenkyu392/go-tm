package tmx

import (
	"bytes"
	"testing"

	"golang.org/x/text/language"
)

func TestNew(t *testing.T) {
	tmx := New(language.MustParse("ja-JP"), language.MustParse("en-US"))
	tmx.AddTU(&TUV{Segment: "吾輩は猫である"}, &TUV{Segment: "I Am a Cat"})
	tmx.AddTU(&TUV{Segment: "人間失格"}, &TUV{Segment: "No Longer Human"})

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
