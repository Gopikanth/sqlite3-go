package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("file.db")
	fmt.Println("creating a db")
	os, err := os.Create("file.db")
	if err != nil {
		log.Fatal(err)
	}
	os.Close()
	db, err := sql.Open("sqlite3", "./file.db")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	stat, err := db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	stat.Exec("Raju", "jaggu")
	rows, _ :=
		db.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

}
