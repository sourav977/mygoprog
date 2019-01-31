package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//var col *sql.ColumnType

func main() {
	db, err := sql.Open("mysql", "root:root@1234@tcp(127.0.0.1:3306)/familydb")
	//db, err := sql.Open("mysql", "rythmos:Rythm0$@1234@tcp(127.0.0.1:3306)/familydb")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established")
	}
	drivers := sql.Drivers()
	fmt.Println("registered drivers:", drivers)

	//name := col.DatabaseTypeName()
	//fmt.Println("name:", name)
	var (
		id        int
		firstname string
		lastname  string
		address   string
	)
	rows, err := db.Query("select * from family where id>=? and id <? order by id desc", 2, 4)
	//rows, err := db.Query("select id,firstname from family where id=?",2)
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
		//b := rows.NextResultSet()
		//fmt.Println("b:", b)
	}
	b := rows.NextResultSet()
	fmt.Println("b:", b)
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
