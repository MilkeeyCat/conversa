package database

import (
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var Db *sql.DB

func InitDB() error {
	db, err := sql.Open("libsql", "file:///tmp/test.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
            created_at TEXT NOT NULL DEFAULT CURRENT_DATE
        )`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	    CREATE TABLE IF NOT EXISTS messages (
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
	        user_id INTEGER NOT NULL,
	        message TEXT NOT NULL,
	        created_at TEXT NOT NULL DEFAULT CURRENT_DATE
	    )`)
	if err != nil {
		return err
	}

	Db = db

	return nil
}
