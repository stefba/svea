package main

import (
	"fmt"
	"github.com/kennygrant/sanitize"
	"strings"
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
		&section{
			Key:   "work",
			Title: "Work",
		},
		&section{
			Key:   "about",
			Title: "About",
		},
		&section{
			Key:   "contact",
			Title: "Contact",
		},
		&section{
			Key:   "experience",
			Title: "Experience",
		},
	},
	"de": []*section{
		&section{
			Key:   "work",
			Title: "Arbeit",
		},
		&section{
			Key:   "about",
			Title: "Über",
		},
		&section{
			Key:   "contact",
			Title: "Kontakt",
		},
		&section{
			Key:   "experience",
			Title: "Erfahrung",
		},
	},
}
