package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	title string
	body  []byte
}

func (p *Page) save() error {
	filename := p.title + ".txt"
	return os.WriteFile(filename, p.body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{title: title, body: body}, nil
}

func handleView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.title, p.body)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{title: title}
	}

	fmt.Fprintf(
		w, 
		"<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
		p.title, p.title, p.body,
	)
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
