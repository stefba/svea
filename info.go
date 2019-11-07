package main

import (
	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v3"
)

type info struct {
	Title string
	Job   string
	Desc  string
	Lead  string
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

	return map[string]string{
		"de": string(blackfriday.MarkdownCommon(texts["de"])),
		"en": string(blackfriday.MarkdownCommon(texts["en"])),
	}, nil
}
