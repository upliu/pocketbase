package core

import (
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
	"os"
)

func connectDB(dbPath string) (*dbx.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://liu@localhost/test?sslmode=disable"
	}
	db, err := dbx.Open("postgres", dsn)
	return db, err
}
