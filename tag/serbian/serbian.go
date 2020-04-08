package serbian

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
sr-BA,Serbian (Bosnia and Herzegovina)
sr-CS,Serbian (Serbia and Montenegro)
sr-ME,Serbian (Montenegro)
sr-RS,Serbian (Serbia)
*/

func init() {
	tag.Register("sr-BA", "Serbian (Bosnia and Herzegovina)")
	tag.Register("sr-CS", "Serbian (Serbia and Montenegro)")
	tag.Register("sr-ME", "Serbian (Montenegro)")
	tag.Register("sr-RS", "Serbian (Serbia)")
}
