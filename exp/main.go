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
	password = "your-password"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var id int
	var name, email string
	row := db.QueryRow(`
  SELECT id, name, email
  FROM users
  WHERE id=$1`, 1)
	err = row.Scan(&id, &name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println("ID:", id, "Name:", name, "Email:", email)

	db.Close()
}
