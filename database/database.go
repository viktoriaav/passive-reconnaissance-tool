package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const createtables string = `
CREATE TABLE user_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    address VARCHAR(100),
    postal_code VARCHAR(10),
    city VARCHAR(50),
    country VARCHAR(50),
    phone_number VARCHAR(15)
);
`

func OpenDB() (*sql.DB, error) {
	dbPath := "./database/database.db"
	// Check if the database file exists
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		// Open a new database connection
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, err
		}
		// Tables
		if _, err := db.Exec(createtables); err != nil {
			return nil, err
		}
		paths := []string{"./database/fill_tables.sql"}
		for _, path := range paths {
			sqlFile, err := os.ReadFile(path)
			if err != nil {
				log.Println("Error reading file:", err)
				continue
			}
			queries := strings.Split(string(sqlFile), ";")
			for _, query := range queries {
				query = strings.TrimSpace(query)
				if query == "" {
					continue
				}
				tx, err := db.Begin()
				if err != nil {
					log.Println("Error starting transaction:", err)
					continue
				}
				_, err = tx.Exec(query)
				if err != nil {
					log.Println("Error executing query:", err)
					tx.Rollback()
					continue
				}
				err = tx.Commit()
				if err != nil {
					log.Println("Error committing transaction:", err)
					continue
				}
			}
		}
		return db, nil
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
