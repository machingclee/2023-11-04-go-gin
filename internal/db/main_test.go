package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/machingclee/2023-11-04-go-gin/util"
)

// const dbDriver = "postgres"
// const dbSource = "postgresql://pguser:pguser@127.0.0.1:5432/pgdb?sslmode=disable"

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot open config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
