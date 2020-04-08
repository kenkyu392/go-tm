package dutch

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
nl-AW,Dutch (Aruba)
nl-BE,Dutch (Belgium)
nl-BQ,"Dutch (Bonaire, Sint Eustatius and Saba)"
nl-CW,Dutch (Curaçao)
nl-NL,Dutch (Netherlands)
nl-SR,Dutch (Suriname)
nl-SX,Dutch (Sint Maarten)
*/

func init() {
	tag.Register("nl-AW", "Dutch (Aruba)")
	tag.Register("nl-BE", "Dutch (Belgium)")
	tag.Register("nl-BQ", "Dutch (Bonaire, Sint Eustatius and Saba)")
	tag.Register("nl-CW", "Dutch (Curaçao)")
	tag.Register("nl-NL", "Dutch (Netherlands)")
	tag.Register("nl-SR", "Dutch (Suriname)")
	tag.Register("nl-SX", "Dutch (Sint Maarten)")
}
