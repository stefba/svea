package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sort"
)

type year struct {
	Year    string
	Entries []*experience
}

type experience struct {
	Title string
	Extra string
	Role  string
	Type  string
}

func readExperiences(path string) (map[string][]*year, error) {
	l, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	langs := map[string][]*year{
		"en": []*year{},
		"de": []*year{},
	}

	for _, file := range l {
		if file.Name() == "info.yaml" {
			continue
		}

		p := path + "/" + file.Name()

		if len(file.Name()) < 4 {
			return nil, fmt.Errorf("Filename too short: %v", p)
		}

		if file.Name()[0] == '.' || file.Name()[0] == '_' {
			continue
		}

		exps, err := readExperienceFile(p)
		if err != nil {
			return nil, err
		}

		de, en := []*experience{}, []*experience{}

		for _, exp := range exps {
			de = append(de, exp["de"])
			en = append(en, exp["en"])
		}

		date := file.Name()[:4]

		langs["en"] = append(langs["en"], &year{
			Year:    date,
			Entries: reverseEntries(en),
		})
		langs["de"] = append(langs["de"], &year{
			Year:    date,
			Entries: reverseEntries(de),
		})
	}

	sort.Sort(yearDesc(langs["de"]))
	sort.Sort(yearDesc(langs["en"]))

	return langs, nil
}

type yearDesc []*year

func (a yearDesc) Len() int           { return len(a) }
func (a yearDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a yearDesc) Less(i, j int) bool { return a[i].Year > a[j].Year }

func reverseEntries(e []*experience) []*experience {
	for i := len(e)/2 - 1; i >= 0; i-- {
		opp := len(e) - 1 - i
		e[i], e[opp] = e[opp], e[i]
	}
	return e
}

func readExperienceFile(path string) ([]map[string]*experience, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	b = makeLinks(b)

	rawExperiences := bytes.Split(bytes.Trim(b, "="), []byte("==="))

	experiences := []map[string]*experience{}

	for _, rawE := range rawExperiences {
		if len(rawE) <= 1 {
			continue
		}

		raw, err := splitText(rawE)
		if err != nil {
			return nil, err
		}

		exp := map[string]*experience{
			"en": &experience{},
			"de": &experience{},
		}

		for _, lang := range []string{"en", "de"} {
			err = yaml.Unmarshal(raw[lang], exp[lang])
			if err != nil {
				return nil, fmt.Errorf("makeLinksLine: %v\n%v", path, err)
			}
		}

		experiences = append(experiences, exp)
	}

	return experiences, nil
}

func makeLinks(bts []byte) []byte {
	b := bytes.Buffer{}
	scanner := bufio.NewScanner(bytes.NewReader(bts))

	for scanner.Scan() {
		b.Write(makeLinksLine(scanner.Bytes()))
		b.WriteString("\n")
	}
	return b.Bytes()
}
