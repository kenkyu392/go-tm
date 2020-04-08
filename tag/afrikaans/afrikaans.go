package afrikaans

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
af-NA,Afrikaans (Namibia)
af-ZA,Afrikaans (South Africa)
*/

func init() {
	tag.Register("af-NA", "Afrikaans (Namibia)")
	tag.Register("af-ZA", "Afrikaans (South Africa)")
}
