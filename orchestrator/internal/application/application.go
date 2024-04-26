package application

import (
	"context"
	"database/sql"
	"orchestrator/db/sqlite"
	"orchestrator/internal/http"
	"orchestrator/internal/orchestrator"
)

func Init() {
	app := new(Application)

	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "store.db")
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

	initOrchestrator := orchestrator.RunOrchestrator(ctx, db)
	app.Orchestrator = initOrchestrator

	initHttpServer := http.NewOrchServer(db)
	app.HttpServer = initHttpServer
	app.HttpServer.RunHttpServer(ctx)
}
