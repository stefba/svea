package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type work struct {
	Title string
	Role  string
	Type  string
	Link  string
	Thumb string
}

func readWorks(path string) (map[string][]*work, error) {
	works := map[string][]*work{
		"en": {},
		"de": {},
	}

	dirs, err := getWorks(path)
	if err != nil {
		return nil, err
	}

	for _, d := range dirs {
		m, err := readWork(path + "/" + d)
		if err != nil {
			return nil, err
		}
		works["en"] = append(works["en"], m["en"])
		works["de"] = append(works["de"], m["de"])
	}

	return works, nil
}

func getWorks(path string) ([]string, error) {
	dirs := []string{}
	dirList, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	sortfile, err := ioutil.ReadFile(path + "/sort")
	if err != nil {
		return nil, err
	}

	sort := splitToLines(string(sortfile))

	for _, file := range dirList {
		if len(file.Name()) > 0 && file.Name()[0] == '_' {
			continue
		}
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	return applySort(dirs, sort), nil
}

func applySort(dirs, sort []string) []string {
	for _, sortElement := range invert(sort) {
		for i, d := range dirs {
			if d == sortElement {
				cut := dirs[i]
				dirs = append([]string{cut}, append(dirs[:i], dirs[i+1:]...)...)
			}
		}
	}
	return dirs
}

func readWork(path string) (map[string]*work, error) {
	m, err := readWorkInfo(path + "/info.yaml")
	if err != nil {
		return nil, err
	}

	return addImagePath(path, m)
}

func addImagePath(path string, m map[string]*work) (map[string]*work, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range l {
		if filepath.Ext(file.Name()) == ".jpg" {
			p := path + "/" + file.Name()
			i := strings.Index(p, "/work")
			if i >= 0 {
				p = p[i:]
			}
			m["de"].Thumb = p
			m["en"].Thumb = p
		}
	}
	return m, nil
}

func readWorkInfo(path string) (map[string]*work, error) {
	split, err := splitFile(path)
	if err != nil {
		return nil, err
	}

	m := map[string]*work{
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

func splitFile(path string) (map[string][]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%v:\n%v", path, err.Error())
	}

	return splitText(b)
}

func splitText(b []byte) (map[string][]byte, error) {
	texts := bytes.Split(b, []byte("---"))
	if len(texts) != 2 {
		return nil, fmt.Errorf("Cannot split text. Check for ---.\n\n%s", b)
	}
	return map[string][]byte{
		"en": texts[0],
		"de": texts[1],
	}, nil
}

func splitToLines(file string) []string {
	return strings.Split(strings.TrimSpace(file), "\n")
}

func invert(ss []string) []string {
	ns := []string{}
	for i := len(ss) - 1; i >= 0; i-- {
		ns = append(ns, ss[i])
	}
	return ns
}
