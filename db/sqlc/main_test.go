package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/todo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	testQueries = New(testDB)
	
	os.Exit(m.Run())

}
