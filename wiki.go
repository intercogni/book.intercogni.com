package main

import (
	"fmt"
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

func main() {
	p1 := &Page{title: "TestPage", body: []byte("This is a sample Page.")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.body))
}
