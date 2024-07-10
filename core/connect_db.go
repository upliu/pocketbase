package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pocketbase/dbx"
	"os"
)

func connectDB(dbPath string) (*dbx.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root@(localhost)/test?charset=utf8&parseTime=True&loc=Local"
	}
	db, err := dbx.Open("mysql", dsn)
	return db, err
}
