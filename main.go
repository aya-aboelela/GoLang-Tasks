package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1011"
	dbname   = "first_db"
)

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	// Initialize connection .
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Initialize connection object.
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()
	CheckError(err)
	fmt.Println("Successfully created connection to database.")

	// Insert some data into table.

	insert := "INSERT INTO company (name, age, address, salary, joined_date) VALUES ($1,$2,$3,$4,$5);"
	_, err = db.Exec(insert, "Paul", 32, "California", 20000, "2001-07-13")
	CheckError(err)

	_, err = db.Exec(insert, "Allen", 25, "Texas", 35000, "2007-12-13")
	CheckError(err)

	_, err = db.Exec(insert, "Teddy", 23, "Norway", 40000, "2008-10-9")
	CheckError(err)

	_, err = db.Exec(insert, "David", 27, "South-Hall", 85000.00, "2015-08-13")
	CheckError(err)

	fmt.Println("Inserted 4 rows of data.")

	// Update some data in table.
	update := "UPDATE company SET  age= $2 WHERE name = $1;"
	_, err = db.Exec(update, "Teddy",30 )
	CheckError(err)
	fmt.Println("Updated 1 row of data.")

	// Alter (Add Column) some data in table
	alter_add := "ALTER TABLE company ADD gender char;"
	_, err = db.Exec(alter_add)
	CheckError(err)
	fmt.Println(" Added Success..!")

	// Alter (delete Column) some data in table
	alter_del := "ALTER TABLE company DROP COLUMN gender;"
	_, err = db.Exec(alter_del)
	CheckError(err)
	fmt.Println(" Deleted Column Success..!")

	// Delete some data from table.
	delete := "DELETE FROM company WHERE name = $1;"
	_, err = db.Exec(delete, "Allen")
	CheckError(err)
	fmt.Println("deleted 1 row of data.")

	// Delete The Whole Data In Table
	delete := "DELETE FROM company;"
	_, err = db.Exec(delete)
	CheckError(err)
	fmt.Println("The Whole Data Deleted ..!")

	// Select
	_select:="SELECT age FROM company;"
	_,err = db.Exec(_select)
	CheckError(err)
	fmt.Println("Showing 1 row of data.")
;}