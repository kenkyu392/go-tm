package turkish

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
tr-CY,Turkish (Cyprus)
tr-TR,Turkish (Turkey)
*/

func init() {
	tag.Register("tr-CY", "Turkish (Cyprus)")
	tag.Register("tr-TR", "Turkish (Turkey)")
}
