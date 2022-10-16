package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tamarelhe/lets_game/api"
	db "github.com/tamarelhe/lets_game/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://lg:lg2022@localhost:5432/lets_game?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api server:", err)
		return
	}
}
