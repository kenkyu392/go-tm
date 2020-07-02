package tm

import "golang.org/x/text/language"

//go:generate go run cmd/gentags/main.go
//go:generate goimports -w .

// Language ...
type Language struct {
	DisplayName         string
	LanguageDisplayName string
	RegionDisplayName   string
	OriginalName        string
	Tag                 language.Tag
	TagName             string
}
