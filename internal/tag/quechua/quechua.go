package quechua

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
qu-BO,Quechua (Bolivia)
qu-EC,Quechua (Ecuador)
qu-PE,Quechua (Peru)
*/

func init() {
	tag.Register("qu-BO", "Quechua (Bolivia)")
	tag.Register("qu-EC", "Quechua (Ecuador)")
	tag.Register("qu-PE", "Quechua (Peru)")
}
