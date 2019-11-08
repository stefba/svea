package main

import (
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v3"
	"svea/hyph"
)

type info struct {
	Title string
	Job   string
	Desc  string
	Lead  string
}

type link struct {
	Name, Href string
	IsActive bool
}

func langNav(lang string) []*link {
	if lang == "de" {
		return []*link {
			&link{
				Name: "Deutsch",
				Href: "/de/",
				IsActive: true,
			},
			&link{
				Name: "English",
				Href: "/",
			},
		}
	}
	return []*link {
		&link{
			Name: "Deutsch",
			Href: "/de/",
		},
		&link{
			Name: "English",
			Href: "/",
			IsActive: true,
		},
	}
}

func readInfo(path string) (map[string]*info, error) {
	split, err := splitFile(path)
	if err != nil {
		return nil, err
	}

	m := map[string]*info{
		"en": &info{},
		"de": &info{},
	}

	err = yaml.Unmarshal(split["en"], m["en"])
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(split["de"], m["de"])
	if err != nil {
		return nil, err
	}

	return m, nil
}

func readAbout(path string) (map[string]string, error) {
	texts, err := splitFile(path)
	if err != nil {
		return nil, err
	}

	de, err := hyph.HyphenateText(string(texts["de"]), "de")
	if err != nil {
		return nil, err
	}

	en, err := hyph.HyphenateText(string(texts["en"]), "en")
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"de": string(blackfriday.MarkdownCommon([]byte(de))),
		"en": string(blackfriday.MarkdownCommon([]byte(en))),
	}, nil
}
