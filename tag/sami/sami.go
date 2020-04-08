package sami

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
sma-NO,Southern Sami (Norway)
sma-SE,Southern Sami (Sweden)
sme-FI,Northern Sami (Finland)
sme-NO,Northern Sami (Norway)
sme-SE,Northern Sami (Sweden)
smj-NO,Lule Sami (Norway)
smj-SE,Lule Sami (Sweden)
smn-FI,Inari Sami (Finland)
sms-FI,Skolt Sami (Finland)
*/

func init() {
	tag.Register("sma-NO", "Southern Sami (Norway)")
	tag.Register("sma-SE", "Southern Sami (Sweden)")
	tag.Register("sme-FI", "Northern Sami (Finland)")
	tag.Register("sme-NO", "Northern Sami (Norway)")
	tag.Register("sme-SE", "Northern Sami (Sweden)")
	tag.Register("smj-NO", "Lule Sami (Norway)")
	tag.Register("smj-SE", "Lule Sami (Sweden)")
	tag.Register("smn-FI", "Inari Sami (Finland)")
	tag.Register("sms-FI", "Skolt Sami (Finland)")

}
