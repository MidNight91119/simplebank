package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq" // Side-effect import registers "postgres" driver
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the Database (db): ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
