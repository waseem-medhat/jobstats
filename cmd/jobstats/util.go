package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(dbURL string, local bool) (*sql.DB, error) {
	if local {
		dbConn, err := sql.Open("sqlite3", dbURL)
		if err != nil {
			return nil, fmt.Errorf("couldn't open sql connection %v", err)
		}

		err = dbConn.Ping()
		if err != nil {
			return nil, fmt.Errorf("couldn't ping sql db %v", err)
		}

		return dbConn, nil
	}

	return nil, fmt.Errorf("unimplemented prod DB")
}
