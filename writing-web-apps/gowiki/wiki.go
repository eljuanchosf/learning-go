package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

// This method will get an HTTP request, parse it, extract the title and load the corresponding file.
// Then it will format the response with some HTML and send it to the ResponseWriter w.
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// This is the main entry point to the program.
func main() {
    // This will register the Handler for the viewHandler
    http.HandleFunc("/view/", viewHandler)
    // This will start the server in the 8080 port
    http.ListenAndServe(":8080", nil)
}
