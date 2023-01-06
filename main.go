package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/logando/simplebank/api"
	db "github.com/logando/simplebank/db/sqlc"
	"github.com/logando/simplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load the configurations...")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connet to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
