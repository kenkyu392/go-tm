package albanian

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
sq-AL,Albanian (Albania)
sq-MK,Albanian (North Macedonia)
*/

func init() {
	tag.Register("sq-AL", "Albanian (Albania)")
	tag.Register("sq-MK", "Albanian (North Macedonia)")
}
