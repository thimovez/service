package main

import (
	"database/sql"
	"fmt"
	"github.com/thimovez/service/pkg/api/user"
	"github.com/thimovez/service/pkg/repository"
	"net/http"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	connectionString := "postgres://postgres:postgres@localhost/**NAME-OF-YOUR-DATABASE-HERE**?sslmode=disable"
	db, err := setupDatabase(connectionString)
	if err != nil {
		return err
	}
	storage := repository.NewStorage(db)
	http.HandleFunc("/login", user.Login)

	http.ListenAndServe(":8080", nil)
}

func setupDatabase(connString string) (*sql.DB, error) {
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
