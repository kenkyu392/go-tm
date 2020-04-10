package swedish

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
sv-FI,Swedish (Finland)
sv-SE,Swedish (Sweden)
*/

func init() {
	tag.Register("sv-FI", "Swedish (Finland)")
	tag.Register("sv-SE", "Swedish (Sweden)")
}
