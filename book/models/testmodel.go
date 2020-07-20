package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func PrintUser() {
	db, err := sql.Open("mysql", "root:199411@tcp(192.168.31.205:3306)/test?charset=utf8")
	defer db.Close()
	if err != nil {
		log.Println(err)
		return
	}

	stmt, err := db.Prepare("select * from user limit 10")
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var id int
		var name string
		var sex int

		rows.Scan(&id, &name, &sex)
		fmt.Println(id, name, sex)
	}
}
