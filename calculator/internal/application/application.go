package application

import (
	"Calculator/db/sqlite"
	"Calculator/internal/calculator"
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	app := &Application{}

	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	if err = sqlite.InitDB(ctx, db); err != nil {
		panic(err)
	}
	app.DB = db

	calcServer := calculator.New(app.DB)
	calcServer.Start(ctx)
	app.Calculator = calcServer

}
