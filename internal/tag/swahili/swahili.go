package swahili

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
sw-KE,Swahili (Kenya)
sw-TZ,Swahili (Tanzania)
sw-UG,Swahili (Uganda)
*/

func init() {
	tag.Register("sw-KE", "Swahili (Kenya)")
	tag.Register("sw-TZ", "Swahili (Tanzania)")
	tag.Register("sw-UG", "Swahili (Uganda)")
}
