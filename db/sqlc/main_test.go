package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
// var testDB *sql.DB

var (
	DB_DRIVER="postgres"
DB_SOURCE="postgresql://root:secret@localhost:5432/cypherdb?sslmode=disable"
)

func TestMain(m *testing.M){
	testDB, err := pgxpool.New(context.Background(), DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}