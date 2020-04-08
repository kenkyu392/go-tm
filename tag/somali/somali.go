package somali

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
so-DJ,Somali (Djibouti)
so-ET,Somali (Ethiopia)
so-KE,Somali (Kenya)
so-SO,Somali (Somalia)
*/

func init() {
	tag.Register("so-DJ", "Somali (Djibouti)")
	tag.Register("so-ET", "Somali (Ethiopia)")
	tag.Register("so-KE", "Somali (Kenya)")
	tag.Register("so-SO", "Somali (Somalia)")
}
