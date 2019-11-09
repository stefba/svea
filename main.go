package main

import (
	//"strings"
	"fmt"
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var root = "."
var debug = false

var siteEn = &site{}
var siteDe = &site{}
var tmpl = &template.Template{}

func main() {

	permaReload := flag.Bool("reload", false, "reload files on every request")
	path := flag.String("path", root+"", "set the root path of this app")
	flag.Parse()

	root = *path
	debug = *permaReload

	err := load()
	if err != nil {
		log.Println(err)
		return
	}

	http.HandleFunc("/", renderEn)
	http.HandleFunc("/de/", renderDe)
	http.HandleFunc("/rl/", reload)
	http.HandleFunc("/static/", serveStatic)
	http.HandleFunc("/robots.txt", serveRobots)
	http.HandleFunc("/googledbd0f1dfe416dbee.html", serveGoogle)

	http.ListenAndServe(":8444", nil)

}

func reload(w http.ResponseWriter, r *http.Request) {
	err := load()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, "did")
}

func renderEn(w http.ResponseWriter, r *http.Request) {
	if debug {
		err := load()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	if filepath.Ext(r.URL.Path) == ".jpg" {
		http.ServeFile(w, r, root+"/data"+r.URL.Path)
		return
	}
	err := tmpl.Execute(w, siteEn)
	if err != nil {
		log.Println(err)
	}
}

func renderDe(w http.ResponseWriter, r *http.Request) {
	if debug {
		err := load()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}
	if filepath.Ext(r.URL.Path) == ".jpg" {
		http.ServeFile(w, r, root+"/data"+r.URL.Path)
		return
	}
	err := tmpl.Execute(w, siteDe)
	if err != nil {
		log.Println(err)
	}
}

func serveGoogle(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, root+"/static/googledbd0f1dfe416dbee.html")
}

func serveRobots(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, root+"/static/robots.txt")
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, root+r.URL.Path)
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
