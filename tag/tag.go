package tag

import (
	"sync"

	"golang.org/x/text/language"
)

var (
	langsMu sync.RWMutex
	langs   = make([]Language, 0)
)

// Language ...
type Language struct {
	Tag         language.Tag
	TagName     string
	DisplayName string
}

// Languages ...
func Languages() []Language {
	return langs
}

// Register makes a language tag available by the provided name.
// If Register is called twice with the same name or if tag is empty,
// it panics.
func Register(tagName string, displayName string) {
	langsMu.Lock()
	defer langsMu.Unlock()

	if tagName == "" {
		panic("tag: Register tag name is empty")
	}
	if displayName == "" {
		panic("tag: Register display name is empty")
	}

	for _, lang := range langs {
		if lang.TagName == tagName {
			panic("tag: Register called twice for tag name " + tagName)
		}
		if lang.DisplayName == displayName {
			panic("tag: Register called twice for display name " + displayName)
		}
	}

	langs = append(langs, Language{
		Tag:         mustParse(tagName),
		TagName:     tagName,
		DisplayName: displayName,
	})
}
