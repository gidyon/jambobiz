// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Page struct {
	TemplateName  string
	Title         string
	DateEffective string
}

var dirsWithTemplates = []string{
	"./templates",
	"./data/company/privacy",
	"./data/company/terms",
}
var templates = template.Must(parseTemplates(dirsWithTemplates...))

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		TemplateName: "about",
		Title:        "About Jambobiz",
	}
	renderTemplate(w, page, "index")
}
func termsHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		TemplateName:  "terms-of-service",
		Title:         "Terms of Services",
		DateEffective: "Last Edited on: Feb 6, 2018",
	}
	renderTemplate(w, page, "index")
}
func privacyHandler(w http.ResponseWriter, r *http.Request) {
	page := &Page{
		TemplateName:  "privacy-policy",
		Title:         "Privacy Policy",
		DateEffective: "Effective as of July 6, 2018",
	}
	renderTemplate(w, page, "index")
}

func parseTemplates(directory ...string) (*template.Template, error) {
	var allFiles []string
	for _, dir := range directory {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range files {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				allFiles = append(allFiles, dir+"/"+filename)
			}
		}
	}
	return template.ParseFiles(allFiles...)
}

func renderTemplate(w http.ResponseWriter, p *Page, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/privacy-policy", privacyHandler)
	http.HandleFunc("/terms-of-service", termsHandler)
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
