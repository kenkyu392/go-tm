package russian

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
ru-BY,Russian (Belarus)
ru-KG,Russian (Kyrgyzstan)
ru-KZ,Russian (Kazakhstan)
ru-MD,Russian (Moldova)
ru-RU,Russian (Russia)
ru-UA,Russian (Ukraine)
*/

func init() {
	tag.Register("ru-BY", "Russian (Belarus)")
	tag.Register("ru-KG", "Russian (Kyrgyzstan)")
	tag.Register("ru-KZ", "Russian (Kazakhstan)")
	tag.Register("ru-MD", "Russian (Moldova)")
	tag.Register("ru-RU", "Russian (Russia)")
	tag.Register("ru-UA", "Russian (Ukraine)")
}
