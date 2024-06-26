package application

import (
	"Calculator/db/sqlite"
	"Calculator/grpc/agentgrpc"
	"Calculator/internal/calculator"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func Init() {
	app := &Application{}

	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "./database/store.db")
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
	go calcServer.Start(ctx)
	app.Calculator = calcServer

	grpcServer := grpc.NewServer()
	agentgrpc.Register(grpcServer, app.Calculator)
	app.gRPCServer = grpcServer
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}

	if err := app.gRPCServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
