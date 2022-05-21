package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	pswd := "Aya2000"
	db, err := sql.Open("mysql", "Aya150:"+pswd+"@tcp(192.168.137.254:3306)/task11")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("opened")
	fmt.Println(pswd)

	insert, err := db.Query("INSERT INTO student VALUES ( '120' ,'doaa','FCI','Level3' )")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	Delete
	delete, err := db.Query("DELETE FROM student WHERE name = 'aya'; ")
	if err != nil {
	  panic(err.Error())
	}

	defer delete.Close()
}
