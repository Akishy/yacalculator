package sqlite

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func tryConnectToDB() {
	ctx := context.TODO()
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err)
	}
}
