package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/MidNight91119/simplebank/db/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot find config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the Database (db): ", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
