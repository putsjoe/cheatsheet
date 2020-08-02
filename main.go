package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type item struct {
	Name string
	Text template.HTML
}

type cheatSheet struct {
	Title string
	Items []item
}

func prepData() []cheatSheet {
	fmt.Println("Loading json data")
	file, _ := ioutil.ReadFile("data.json")

	// type Item1 struct {
	// 	Name string `json:"name"`
	// 	Text string `json:"text"`
	// }

	var dat []cheatSheet
	if err := json.Unmarshal([]byte(file), &dat); err != nil {
		panic(err)
	}
	// fmt.Printf("%T\n", dat)
	// for _, d := range dat {
	// 	fmt.Println(d.Title)
	// 	for _, i := range d.Items {
	// 		fmt.Println("  %s - %s \n", i.Name, i.Text)
	// 	}
	// }
	// fmt.Println(" -- --")
	// fmt.Println(dat)

	return dat
}

func hello(w http.ResponseWriter, r *http.Request) {
	var items = []item{
		{"Example1", "Blahsnsn stuff"},
		{"Example2", "<code>Cheeky Cheeky</code>"},
	}
	var sheet = cheatSheet{"First CH", items}

	t, _ := template.ParseFiles("stuff.html")
	var err = t.Execute(w, sheet)
	if err != nil {
		panic(err)
	}
}

type Data struct {
	data   []cheatSheet
	titles []string
}

func (data *Data) showMenu(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("menu.html")
	var err = t.Execute(w, data.titles)
	if err != nil {
		panic(err)
	}
}

func find(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (data *Data) showCheatsheet(w http.ResponseWriter, r *http.Request) {
	sheet := strings.TrimPrefix(r.URL.Path, "/c/")
	found := find(sheet, data.titles)
	if !found {
		http.NotFound(w, r)
	}
	for _, c := range data.data {
		if c.Title == sheet {
			t, _ := template.ParseFiles("stuff.html")
			var err = t.Execute(w, c)
			if err != nil {
				panic(err)
			}
		}
	}

}

func main() {
	dat := prepData()
	menu := make([]string, len(dat))
	for _, d := range dat {
		menu = append(menu, d.Title)
	}
	data := &Data{dat, menu}
	http.HandleFunc("/", hello)
	http.HandleFunc("/menu", data.showMenu)
	http.HandleFunc("/c/", data.showCheatsheet)

	fmt.Println("Serving on port 8000...")
	http.ListenAndServe(":8000", nil)
}
