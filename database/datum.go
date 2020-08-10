package database

import (
	"cheatsheet/load"
	"database/sql"
	"fmt"
)

type Datum struct {
	Db *sql.DB
}

func (d *Datum) GetMenu() {
	rows, _ := d.Db.Query("SELECT name FROM cheatsheet ORDER BY NAME")
	var name string
	for rows.Next() {
		rows.Scan(&name)
		fmt.Println(name)
	}
}

func (d *Datum) GetCheatsheet(sheet string) (load.CheatSheet, bool) {
	var res int
	row := d.Db.QueryRow("SELECT id FROM cheatsheet WHERE name=?", sheet)
	err := row.Scan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			return load.CheatSheet{}, false
		}
		panic(err)
	}
	rows, _ := d.Db.Query("SELECT name, text FROM data WHERE sid=?", res)
	var itm load.Item
	itms := make([]load.Item, 0)
	for rows.Next() {
		rows.Scan(&itm.Name, &itm.Text)
		itms = append(itms, itm)
	}
	return load.CheatSheet{Title: sheet, Items: itms}, true

}
