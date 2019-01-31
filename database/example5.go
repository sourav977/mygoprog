package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "rythmos:Rythm0$@1234@tcp(127.0.0.1:3306)/familydb")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established")
	}
	newdb, err := db.Exec("select * from family where id=? ", 4)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	lastID, err := newdb.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := newdb.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
}
