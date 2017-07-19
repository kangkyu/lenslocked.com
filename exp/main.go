package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{})

	// Get the first user in the DB
	// GORM typically sorts these by ID
	var u User
	db.First(&u)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)

	// Query by ID
	id := 1
	db.First(&u, id)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)

	// Query by <= and get the first result even if there are
	// multiple
	maxId := 3
	db.Where("id <= ?", maxId).First(&u)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)

	// Query by user object
	var u2 User
	u2.Email = "jon@calhoun.io"
	db.Where(u2).First(&u2)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u2)
}
