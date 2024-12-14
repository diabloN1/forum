package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ExecuteMigrations(db *sql.DB, filePath string) {
	migrationsSql, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to open SQL file: %v", err)
	}

	_, err = db.Exec(string(migrationsSql))
	if err != nil {
		log.Fatalf("Failed to execute SQL: %v", err)
	}

	fmt.Println("Migrations executed successfully!")
}
