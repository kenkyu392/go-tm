package oromo

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
om-ET,Oromo (Ethiopia)
om-KE,Oromo (Kenya)
*/

func init() {
	tag.Register("om-ET", "Oromo (Ethiopia)")
	tag.Register("om-KE", "Oromo (Kenya)")
}
