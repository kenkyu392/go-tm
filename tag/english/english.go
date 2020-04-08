package english

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
en-AU,English (Australia)
en-BZ,English (Belize)
en-CA,English (Canada)
en-GB,English (United Kingdom)
en-IE,English (Ireland)
en-JM,English (Jamaica)
en-NZ,English (New Zealand)
en-TT,English (Trinidad and Tobago)
en-US,English (United States of America)
en-ZA,English (South Africa)
en-ZW,English (Zimbabwe)
*/

func init() {
	tag.Register("en-AU", "English (Australia)")
	tag.Register("en-BZ", "English (Belize)")
	tag.Register("en-CA", "English (Canada)")
	tag.Register("en-GB", "English (United Kingdom)")
	tag.Register("en-IE", "English (Ireland)")
	tag.Register("en-JM", "English (Jamaica)")
	tag.Register("en-NZ", "English (New Zealand)")
	tag.Register("en-TT", "English (Trinidad and Tobago)")
	tag.Register("en-US", "English (United States of America)")
	tag.Register("en-ZA", "English (South Africa)")
	tag.Register("en-ZW", "English (Zimbabwe)")
}
