package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/machingclee/2023-11-04-go-gin/api"
	"github.com/machingclee/2023-11-04-go-gin/internal/db"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://pguser:pguser@127.0.0.1:5432/pgdb?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot Start Server:", err)
	}

}
