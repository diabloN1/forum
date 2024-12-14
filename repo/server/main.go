package main

import (
	"database/sql"
	"fmt"
	"log"
	myDB "forum/database"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)
func init() {
	var err error
	db, err = sql.Open("sqlite3", "database/myDatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")

	// Execute migrations
	myDB.ExecuteMigrations(db, "database/migrations.sql")
}

func main() {
	
}
