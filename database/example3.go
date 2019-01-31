package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@1234@tcp(127.0.0.1:3306)/familydb")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established")
	}
	var (
		id        int
		firstname string
		lastname  string
		address   string
	)
	rows, err := db.Query("insert into family (id,firstname,lastname,address) values(?,?,?,?)", 6, "sunaina", "patnaik", "bangelore")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &firstname, &lastname, &address)
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
