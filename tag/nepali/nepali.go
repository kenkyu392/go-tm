package nepali

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
ne-IN,Nepali (India)
ne-NP,Nepali (Federal Democratic Republic of Nepal)
*/

func init() {
	tag.Register("ne-IN", "Nepali (India)")
	tag.Register("ne-NP", "Nepali (Federal Democratic Republic of Nepal)")
}
