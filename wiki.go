package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	// p1 := &Page{title: "TestPage", body: []byte("This is a sample Page.")}
	// p1.save()

	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.body))

	http.HandleFunc("/view/", handleView)
	http.HandleFunc("/edit/", handleEdit)
	// http.HandleFunc("/save/", handleSave)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
