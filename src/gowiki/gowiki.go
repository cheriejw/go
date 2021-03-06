package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil { // if there is error
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil { // if there is an error, return no page.
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: string("ShinyPage"), Body: []byte("Some text for Shiny Page.")}
	p1.save()
	p2, _ := loadPage("ShinyPage")
	fmt.Println(p2.Title)
	fmt.Println(string(p2.Body))

	http.HandleFunc("/", handler) // Kind of like a router?
	http.HandleFunc("/view/", viewHandler)
	// http.HandleFumc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)
	
	// Listen and serve to localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}