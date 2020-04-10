package kurdish

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
ku-IQ,Kurdish (Iraq)
ku-IR,Kurdish (Islamic Republic of Iran)
*/

func init() {
	tag.Register("ku-IQ", "Kurdish (Iraq)")
	tag.Register("ku-IR", "Kurdish (Islamic Republic of Iran)")
}
