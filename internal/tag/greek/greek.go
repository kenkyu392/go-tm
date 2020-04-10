package greek

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
el-CY,Greek (Cyprus)
el-GR,Greek (Greece)
grc-GR,Ancient Greek (Greece)
gre-GR,Modern Greek (Greece)
*/

func init() {
	tag.Register("el-CY", "Greek (Cyprus)")
	tag.Register("el-GR", "Greek (Greece)")
	tag.Register("grc-GR", "Ancient Greek (Greece)")
	tag.Register("gre-GR", "Modern Greek (Greece)")
}
