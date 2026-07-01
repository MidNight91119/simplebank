package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/MidNight91119/simplebank/db/util"
	_ "github.com/lib/pq" // Side-effect import registers "postgres" driver
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot find config:", err)
	}
	
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the Database (db): ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
