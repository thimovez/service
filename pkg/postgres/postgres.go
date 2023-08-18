package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func SetupDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// ping the DB to ensure that it is connected
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
