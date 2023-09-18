package main

import (
	"fmt"
	"strings"

	"github.com/kennygrant/sanitize"
)

type sections []*section

func (ss sections) Title(key string) (string, error) {
	for _, s := range ss {
		if s.Key == key {
			return s.Title, nil
		}
	}
	return "", fmt.Errorf("Section key not found")
}

type section struct {
	Key   string
	Title string
}

func (s *section) Href() string {
	return "#" + strings.ToLower(sanitize.BaseName(s.Title))
}

var sectionsLangs = map[string]sections{
	"en": []*section{
		{
			Key:   "work",
			Title: "Works",
		},
		{
			Key:   "about",
			Title: "About me",
		},
		{
			Key:   "contact",
			Title: "Contact",
		},
		{
			Key:   "experience",
			Title: "Experience",
		},
	},
	"de": []*section{
		{
			Key:   "work",
			Title: "Arbeit",
		},
		{
			Key:   "about",
			Title: "Über mich",
		},
		{
			Key:   "contact",
			Title: "Kontakt",
		},
		{
			Key:   "experience",
			Title: "Erfahrung",
		},
	},
}
