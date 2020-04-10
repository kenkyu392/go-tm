package romanian

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
ro-MD,Romanian (Moldova)
ro-RO,Romanian (Romania)
*/

func init() {
	tag.Register("ro-MD", "Romanian (Moldova)")
	tag.Register("ro-RO", "Romanian (Romania)")
}
