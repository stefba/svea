package main

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type work struct {
	Title string
	Role  string
	Type  string
	Link  string
	Thumb string
}

func readWorks(path string) (map[string][]*work, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	works := map[string][]*work{
		"en": []*work{},
		"de": []*work{},
	}

	for _, file := range l {
		if file.IsDir() {
			m, err := readWork(path + "/" + file.Name())
			if err != nil {
				return nil, err
			}
			works["en"] = append(works["en"], m["en"])
			works["de"] = append(works["de"], m["de"])
		}
	}
	return works, nil
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
		"en": &work{},
		"de": &work{},
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
