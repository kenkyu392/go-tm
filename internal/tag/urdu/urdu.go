package urdu

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
ur-IN,Urdu (India)
ur-PK,Urdu (Islamic Republic of Pakistan)
*/

func init() {
	tag.Register("ur-IN", "Urdu (India)")
	tag.Register("ur-PK", "Urdu (Islamic Republic of Pakistan)")
}
