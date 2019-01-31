package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Family struct to store retrive data from database
type Family struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

func main() {
	db, err := sql.Open("mysql", "rythmos:Rythm0$@1234@tcp(127.0.0.1:3306)/familydb")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established")
	}

	rows, err := db.Query("select * from family")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var f Family
	for rows.Next() {
		err := rows.Scan(&f.ID, &f.Firstname, &f.Lastname, &f.Address)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, firstname, lastname, address)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
