// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	_ "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func renderImage(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl)

	t.Execute(w, p)
}

func test(w http.ResponseWriter, r *http.Request) {

	index := template.Must(template.ParseFiles("index.html"))

	http.ServeFile(w, r, "index.html")

	index.Execute(w, nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, Nice to meet you %s!", r.URL.Path[1:])
	p, _ := loadPage("welcome")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	renderTemplate(w, "welcome", p)

}

func handlerImage(w http.ResponseWriter, r *http.Request) {

	//	p, _ := loadPage("welcome")

	http.ServeFile(w, r, "cat.jpg")

}

var ImageTemplate string = `<!DOCTYPE html>
    <html lang="en"><head></head>
    <body><img src="{{staticFile "cat.jpg"}}"></img></body>`

func main() {

	http.HandleFunc("/", handler)
	//	http.HandleFunc("/img/", handlerImage)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
