package main

import (
	"database/sql"
	"github.com/Pedroxer/Simple-todo/api"
	"github.com/Pedroxer/Simple-todo/db/sqlc"
	"github.com/Pedroxer/Simple-todo/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBAddress)
	if err != nil {
		log.Fatal("cannot connect to the db", err)
	}
	query := sqlc.New(db)
	server, err := api.NewServer(config, query)
	if err != nil {
		log.Fatal("cannot create a server")
	}
	err = server.Start(config)
	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
