package portuguese

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
pt-BR,Portuguese (Brazil)
pt-PT,Portuguese (Portugal)
*/

func init() {
	tag.Register("pt-BR", "Portuguese (Brazil)")
	tag.Register("pt-PT", "Portuguese (Portugal)")
}
