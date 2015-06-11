package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
    // Title is a regular string
    Title string
    // Body is a "byte slice". They are like arrays, but way more flexible: http://blog.golang.org/go-slices-usage-and-internals
    Body  []byte
}

// This method will save the above defined struct to disk
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// This method will return a Page (defined above) by reading the data from a file in disk.
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// This is the main entry point to the program. Shoud print "This is a sample Page."
func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}