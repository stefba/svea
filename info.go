package main

import (
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v3"
)

type info struct {
	Title string
	Jobs  []string
	Desc  string
	Lead  string
}

func (i *info) PrintJobs() string {
	str := ""
	for i, v := range i.Jobs {
		if i > 0 {
			str += ", "
		}
		str += v
	}
	return str
}

type link struct {
	Name, Href string
	IsActive   bool
}

func langNav(lang string) []*link {
	if lang == "de" {
		return []*link{
			{
				Name:     "Deutsch",
				Href:     "/de/",
				IsActive: true,
			},
			{
				Name: "English",
				Href: "/",
			},
		}
	}
	return []*link{
		{
			Name: "Deutsch",
			Href: "/de/",
		},
		{
			Name:     "English",
			Href:     "/",
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
		"en": {},
		"de": {},
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

func readMarkdownSplit(path string) (map[string]string, error) {
	texts, err := splitFile(path)
	if err != nil {
		return nil, err
	}

	en := string(blackfriday.MarkdownCommon(texts["en"]))
	de := string(blackfriday.MarkdownCommon(texts["de"]))

	return map[string]string{
		"de": de,
		"en": en,
	}, nil
}
