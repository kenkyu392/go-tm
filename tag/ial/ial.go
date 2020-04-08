package ial

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
constructed language > International auxiliary language
eo,Esperanto (World)
io,Ido (World)
vo,Volapük (World)
*/

func init() {
	tag.Register("eo", "Esperanto (World)")
	tag.Register("io", "Ido (World)")
	tag.Register("vo", "Volapük (World)")
}
