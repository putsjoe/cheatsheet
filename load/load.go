package load

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Name string
	Text string
}

type CheatSheet struct {
	Title string
	Items []Item
}

func PrepData() []CheatSheet {
	fmt.Println("Loading json data")
	file, _ := ioutil.ReadFile("data.json")
	var dat []CheatSheet

	if err := json.Unmarshal([]byte(file), &dat); err != nil {
		panic(err)
	}

	return dat
}

func insertItem(item Item, cid int, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO data(name, text, sid) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := stmt.Exec(item.Name, item.Text, cid); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(".")
}

func insertCheatsheet(title string, db *sql.DB) int {
	stmt, err := db.Prepare("INSERT INTO cheatsheet(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(title)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(lastID)

}

func loadDb(cheatsheet []CheatSheet, db *sql.DB) {
	for _, i := range cheatsheet {
		cid := insertCheatsheet(i.Title, db)
		for _, item := range i.Items {
			insertItem(item, cid, db)
		}
	}
	fmt.Println("Finished")
}

func main() {
	db, err := sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic("Cant get DB")
	}

	defer db.Close()
	loadDb(PrepData(), db)
}
