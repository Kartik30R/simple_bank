package main

import (
	"database/sql"
	"log"

	"github.com/kartik30R/simple_bank/api"
	db "github.com/kartik30R/simple_bank/db/sqlc"
	"github.com/kartik30R/simple_bank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("config cant be loaded", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config,store)
	

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error starting a server", err)
	}

}
