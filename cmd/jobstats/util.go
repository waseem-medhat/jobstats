package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func initDB(dbURL string, local bool) (*sql.DB, error) {
	driver := ""
	if local {
		driver = "sqlite3"
	} else {
		driver = "libsql"
	}

	dbConn, err := sql.Open(driver, dbURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't open sql connection %v", err)
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("couldn't ping sql db %v", err)
	}

	return dbConn, nil
}
