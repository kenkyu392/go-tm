package tag

import (
	"golang.org/x/text/language"
)

func mustParse(s string) language.Tag {
	tag, err := language.Parse(s)
	if err != nil {
		panic(err)
	}
	return tag
}

func mustParseRegion(s string) language.Region {
	region, err := language.ParseRegion(s)
	if err != nil {
		panic(err)
	}
	return region
}
