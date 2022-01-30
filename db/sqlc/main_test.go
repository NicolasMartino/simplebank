package db

import (
	"database/sql"
	"log"
	"simple_bank/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

//Setup and teardown for DB tests
func setupDBTestSuite(tb testing.TB) func(tb testing.TB) {
	var err error
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Cannot load configuration", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	testQueries = New(testDB)

	// Return a function to teardown the test
	return func(tb testing.TB) {
	}
}
