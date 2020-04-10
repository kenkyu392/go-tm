package tmx

import (
	"bytes"
	"testing"
	"time"

	"github.com/kenkyu392/go-tm"
)

func TestNew(t *testing.T) {
	tmx, err := New(
		SourceLangOption(tm.Tag_jaJP),
		UseUTF8XMLEncodingOption(),
	)
	if err != nil {
		t.Errorf("tmx: %+v", err)
	}
	now := time.Now()

	tmx.AddTU(
		"ID001",
		NewTUV(
			CreationTUVOption(now, DefaultUser),
			ChangeTUVOption(now, DefaultUser),
			XMLLangTUVOption(tm.Tag_jaJP),
			SegmentTUVOption("吾輩は猫である"),
		),
		NewTUV(
			CreationTUVOption(now, DefaultUser),
			ChangeTUVOption(now, DefaultUser),
			XMLLangTUVOption(tm.Tag_enUS),
			SegmentTUVOption("I Am a Cat"),
		),
	)
	tmx.AddTU(
		"ID001",
		NewTUV(
			CreationTUVOption(now, DefaultUser),
			ChangeTUVOption(now, DefaultUser),
			XMLLangTUVOption(tm.Tag_jaJP),
			SegmentTUVOption("人間失格"),
		),
		NewTUV(
			CreationTUVOption(now, DefaultUser),
			ChangeTUVOption(now, DefaultUser),
			XMLLangTUVOption(tm.Tag_enUS),
			SegmentTUVOption("No Longer Human"),
		),
		NewTUV(
			CreationTUVOption(now, DefaultUser),
			ChangeTUVOption(now, DefaultUser),
			XMLLangTUVOption(tm.Tag_frFR),
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
