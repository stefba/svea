package main

import (
	"fmt"
	"os"
	"text/template"
)

type site struct {
	Home       string
	Langs      []*link
	Info       *info
	Experience []*year
	Work       []*work
	About      string
	Services   string
	Sections   sections
	//Contact
}

func (s *site) Lang() string {
	if s.Home == "/de/" {
		return "de"
	}
	return "en"
}

func load() error {
	err := loadData()
	if err != nil {
		return err
	}
	return loadTemplate()
}

func loadData() error {
	info, err := readInfo(root + "/data/info.yaml")
	if err != nil {
		return err
	}

	works, err := readWorks(root + "/data/work")
	if err != nil {
		return err
	}

	exp, err := readExperiences(root + "/data/experience")
	if err != nil {
		return err
	}

	about, err := readMarkdownSplit(root + "/data/about/about.md")
	if err != nil {
		return err
	}

	services, err := readMarkdownSplit(root + "/data/services/services.md")
	if err != nil {
		return err
	}

	siteDe = &site{
		Home:       "/de/",
		Langs:      langNav("de"),
		Sections:   sectionsLangs["de"],
		Info:       info["de"],
		Work:       works["de"],
		Experience: exp["de"],
		About:      about["de"],
		Services:   services["de"],
	}

	siteEn = &site{
		Home:       "/",
		Langs:      langNav("en"),
		Sections:   sectionsLangs["en"],
		Info:       info["en"],
		Work:       works["en"],
		Experience: exp["en"],
		About:      about["en"],
		Services:   services["en"],
	}

	return nil
}

func loadTemplate() error {
	t, err := template.ParseFiles(root + "/svea.html")
	if err != nil {
		return err
	}
	css, err := loadCSS()
	if err != nil {
		return err
	}
	t, err = t.Parse(fmt.Sprintf(`{{define "css"}}<style>%v</style>{{end}}`, css))
	if err != nil {
		return err
	}
	tmpl = t
	return nil
}

func loadCSS() (string, error) {
	b, err := os.ReadFile(root + "/css/dist/main.css")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
