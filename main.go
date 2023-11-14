package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/machingclee/2023-11-04-go-gin/api"
	"github.com/machingclee/2023-11-04-go-gin/internal/db"
	"github.com/machingclee/2023-11-04-go-gin/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot laod config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot Start Server:", err)
	}

}
