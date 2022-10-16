package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tamarelhe/lets_game/api"
	db "github.com/tamarelhe/lets_game/db/sqlc"
	"github.com/tamarelhe/lets_game/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
		return
	}

	err = server.Start(config.DBAddress)
	if err != nil {
		log.Fatal("cannot start api server:", err)
		return
	}
}
