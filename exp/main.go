package main

import (
	"fmt"

	"lenslocked.com/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	user := models.User{
		Name: "Mike Palmer",
		Email: "mike@dundie.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	foundUser, err := us.ByEmail("mike@dundie.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}
