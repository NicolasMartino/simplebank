package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

//Setup and teardown for DB tests
func setupDBTestSuite(tb testing.TB) func(tb testing.TB) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	testQueries = New(testDB)

	// Return a function to teardown the test
	return func(tb testing.TB) {
	}
}
