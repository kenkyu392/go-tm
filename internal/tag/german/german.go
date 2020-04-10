package german

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
de-AT,German (Austria)
de-BE,German (Belgium)
de-CH,German (Switzerland)
de-DE,German (Germany)
de-IT,German (Italy)
de-LI,German (Liechtenstein)
de-LU,German (Luxembourg)
*/

func init() {
	tag.Register("de-AT", "German (Austria)")
	tag.Register("de-BE", "German (Belgium)")
	tag.Register("de-CH", "German (Switzerland)")
	tag.Register("de-DE", "German (Germany)")
	tag.Register("de-IT", "German (Italy)")
	tag.Register("de-LI", "German (Liechtenstein)")
	tag.Register("de-LU", "German (Luxembourg)")
}
