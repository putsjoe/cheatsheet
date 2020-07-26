package main

import (
	"html/template"
	"net/http"
)

type item struct {
	Title string
	Text  template.HTML
}

type cheatSheet struct {
	Title string
	Items []item
}

func hello(w http.ResponseWriter, req *http.Request) {
	var items = []item{
		item{"Example1", "Blahsnsn more stuff"},
		item{"Example2", "<code>Cheeky Cheeky</code>"},
	}
	var sheet = cheatSheet{"First CH", items}
	t, _ := template.ParseFiles("stuff.html")
	var err = t.Execute(w, sheet)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8000", nil)
}
