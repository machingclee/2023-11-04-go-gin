package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const dbDriver = "postgres"
const dbSource = "postgresql://pguser:pguser@127.0.0.1:5432/pgdb?sslmode=disable"

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
