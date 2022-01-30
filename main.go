package main

import (
	"database/sql"
	"log"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
	"simple_bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configuration", err)
	}

	var conn *sql.DB
	conn, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
