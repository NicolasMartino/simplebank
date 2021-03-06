package main

import (
	"database/sql"
	"log"

	"github.com/NicolasMartino/simplebank/api"
	"github.com/NicolasMartino/simplebank/core"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
	"github.com/NicolasMartino/simplebank/util"

	_ "github.com/golang/mock/mockgen/model"
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
	services := core.NewServices(store)
	server := api.NewServer(services, &config)

	err = server.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
