package main

import (
	"text/template"
)

func load() error {
	err := loadData()
	if err != nil {
		return err
	}
	return loadTemplate()
}

func loadData() error {
	info, err := readInfo(root+"/data/info.yaml")
	if err != nil {
		return err
	}

	works, err := readWorks(root+"/data/work")
	if err != nil {
		return err
	}

	exp, err := readExperiences(root+"/data/experience")
	if err != nil {
		return err
	}

	about, err := readAbout(root+"/data/about/about.md")
	if err != nil {
		return err
	}

	siteDe = &site{
		Home:       "/de/",
		Langs:		langNav("de"),
		Sections:   sectionsLangs["de"],
		Info:       info["de"],
		Work:       works["de"],
		Experience: exp["de"],
		About:      about["de"],
	}

	siteEn = &site{
		Home:       "/",
		Langs:		langNav("en"),
		Sections:   sectionsLangs["en"],
		Info:       info["en"],
		Work:       works["en"],
		Experience: exp["en"],
		About:      about["en"],
	}

	return nil
}

func loadTemplate() error {
	t, err := template.ParseFiles(root+"/svea.html")
	if err != nil {
		return err
	}
	tmpl = t
	return nil
}
