package main

import (
	"database/sql"
	"log"
	"simplt_bank/api"
	db "simplt_bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAdress = "0.0.0.0:8090"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
