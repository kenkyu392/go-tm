package tmx

import (
	"bytes"
	"testing"
	"time"

	"github.com/kenkyu392/go-tm"
)

func TestNew(t *testing.T) {
	now := time.Unix(1577836800, 0)

	tmx, err := New(
		SourceLangOption(tm.Tag_jaJP),
		AdminLangOption(tm.Tag_jaJP),
		UseUTF8XMLEncodingOption(),
		CreationToolOption(DefaultCreationTool, DefaultCreationToolVersion),
		CreationOption(now, DefaultUserName),
		ChangeOption(now, DefaultUserName),
		OriginalTranslationMemoryFormatOption(DefaultOriginalTranslationMemoryFormat),
	)
	if err != nil {
		t.Errorf("tmx: %+v", err)
	}

	tmx.AddTU(
		"ID001",
		NewTUV(
			XMLLangTUVOption(tm.Tag_jaJP),
			SegmentTUVOption("吾輩は猫である"),
			CreationTUVOption(now, DefaultUserName),
			ChangeTUVOption(now, DefaultUserName),
		),
		NewTUV(
			XMLLangTUVOption(tm.Tag_enUS),
			SegmentTUVOption("I Am a Cat"),
			CreationTUVOption(now, DefaultUserName),
			ChangeTUVOption(now, DefaultUserName),
		),
	)
	tmx.AddTU(
		"ID001",
		NewTUV(
			XMLLangTUVOption(tm.Tag_jaJP),
			SegmentTUVOption("人間失格"),
			CreationTUVOption(now, DefaultUserName),
			ChangeTUVOption(now, DefaultUserName),
		),
		NewTUV(
			XMLLangTUVOption(tm.Tag_enUS),
			SegmentTUVOption("No Longer Human"),
			CreationTUVOption(now, DefaultUserName),
			ChangeTUVOption(now, DefaultUserName),
		),
		NewTUV(
			XMLLangTUVOption(tm.Tag_frFR),
			SegmentTUVOption("La déchéance d'un homme"),
			CreationTUVOption(now, DefaultUserName),
			ChangeTUVOption(now, DefaultUserName),
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
