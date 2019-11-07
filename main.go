package main

import (
	//"strings"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type site struct {
	Home       string
	Info       *info
	Experience []*year
	Work       []*work
	About      string
	Sections   []*section
	//Contact
}

var root = "."

var siteEn = &site{}
var siteDe = &site{}
var tmpl = &template.Template{}

func main() {

	err := load()
	if err != nil {
		log.Println(err)
		return
	}

	http.HandleFunc("/", renderEn)
	http.HandleFunc("/de/", renderDe)
	http.ListenAndServe(":8444", nil)

}

func renderEn(w http.ResponseWriter, r *http.Request) {
	err := load()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if filepath.Ext(r.URL.Path) == ".jpg" {
		http.ServeFile(w, r, "./data"+r.URL.Path)
		return
	}
	err = tmpl.Execute(w, siteEn)
	if err != nil {
		log.Println(err)
	}
}

func renderDe(w http.ResponseWriter, r *http.Request) {
	err := load()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if filepath.Ext(r.URL.Path) == ".jpg" {
		http.ServeFile(w, r, "./data"+r.URL.Path)
		return
	}
	err = tmpl.Execute(w, siteDe)
	if err != nil {
		log.Println(err)
	}
}

/*
func readSort() ([]string, error) {
	b, err := ioutil.ReadFile(root + "/data/sort")
	if err != nil {
		return nil, err
	}
	l := strings.Split(strings.TrimSpace(string(b)), "\n")
	return l, nil
}
*/
