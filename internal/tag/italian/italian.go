package italian

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
it-CH,Italian (Switzerland)
it-IT,Italian (Italy)
it-SM,Italian (San Marino)
it-VA,Italian (Vatican City State)
*/

func init() {
	tag.Register("it-CH", "Italian (Switzerland)")
	tag.Register("it-IT", "Italian (Italy)")
	tag.Register("it-SM", "Italian (San Marino)")
	tag.Register("it-VA", "Italian (Vatican City State)")
}
