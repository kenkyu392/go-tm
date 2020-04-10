package tibetan

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
bo-CN,Tibetan (People's Republic of China)
bo-IN,Tibetan (India)
bo-NP,Tibetan (Federal Democratic Republic of Nepal)
*/

func init() {
	tag.Register("bo-CN", "Tibetan (People's Republic of China)")
	tag.Register("bo-IN", "Tibetan (India)")
	tag.Register("bo-NP", "Tibetan (Federal Democratic Republic of Nepal)")
}
