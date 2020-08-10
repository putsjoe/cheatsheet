package main

import (
	"cheatsheet/load"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type data struct {
	data   []load.CheatSheet
	titles []string
}

func find(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (data *data) menu(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("menu.html")
	var err = t.Execute(w, data.titles)
	if err != nil {
		panic(err)
	}
}

func (data *data) editSheet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data.add(
			r.FormValue("sheet"), r.FormValue("title"), r.FormValue("text"))
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
		return
	}
	sheet := strings.TrimPrefix(r.URL.Path, "/edit/")
	found := find(sheet, data.titles)
	if !found {
		http.NotFound(w, r)
	}
	for _, c := range data.data {
		if c.Title == sheet {
			t, _ := template.ParseFiles("edit.html")
			var err = t.Execute(w, c)
			if err != nil {
				panic(err)
			}
		}
	}

}

func (data *data) showCheatsheet(w http.ResponseWriter, r *http.Request) {
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

func (data *data) saveData() {
	file, _ := json.MarshalIndent(data.data, "", " ")

	if err := ioutil.WriteFile("data.json", file, 0644); err != nil {
		panic(err)
	}
}

func (data *data) add(sheet, title, text string) {
	if len(text) < 5 || len(title) < 1 {
		return
	}

	newitem := load.Item{Name: title, Text: text}
	for i, c := range data.data {
		if c.Title == sheet {
			for _, item := range c.Items {
				if item.Name == newitem.Name {
					return
				}
			}
			data.data[i].Items = append(c.Items, newitem)
			data.saveData()
			break
		}
	}
}

func main() {
	dat := load.PrepData()
	menu := make([]string, len(dat))
	for _, d := range dat {
		menu = append(menu, d.Title)
	}
	data := &data{dat, menu}

	http.HandleFunc("/", data.menu)
	http.HandleFunc("/c/", data.showCheatsheet)
	http.HandleFunc("/edit/", data.editSheet)
	http.HandleFunc("/menu", data.menu)

	fmt.Println("Serving on port 8000...")
	http.ListenAndServe(":8000", nil)
}
