package persian

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
fa-AF,Persian (Afghanistan)
fa-IR,Persian (Islamic Republic of Iran)
*/

func init() {
	tag.Register("fa-AF", "Persian (Afghanistan)")
	tag.Register("fa-IR", "Persian (Islamic Republic of Iran)")
}
