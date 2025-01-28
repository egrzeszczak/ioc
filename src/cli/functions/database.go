package functions

// global handler
// path to database is "./database.go"

// CreateNewDatabase()
// ListAllDatabaseContents()

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func GetDB() *sql.DB {
	once.Do(initDB)
	return db
}
