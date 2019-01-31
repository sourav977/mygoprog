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
	stmt, err := db.Prepare("INSERT INTO family(id,firstname,lastname,address) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(5, "santosh", "patnaik", "bangelore")
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
