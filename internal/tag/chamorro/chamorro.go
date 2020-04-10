package chamorro

import (
	"github.com/kenkyu392/go-tm/internal/tag"
)

/*
ch-GU,Chamorro (Guam)
ch-US,Chamorro (United States of America)
*/

func init() {
	tag.Register("ch-GU", "Chamorro (Guam)")
	tag.Register("ch-US", "Chamorro (United States of America)")
}
