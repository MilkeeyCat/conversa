package database

import (
	"database/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var Db *sql.DB

func InitDB() error {
	db, err := sql.Open("libsql", "file:///tmp/test.db")

	if err != nil {
		return err
	}

	db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name STRING
        )`)

	Db = db

	return nil
}
