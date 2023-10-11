package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dasotd/gocypher/api"
	db "github.com/dasotd/gocypher/db/sqlc"
	"github.com/dasotd/gocypher/event"
	"github.com/dasotd/gocypher/util"
	_ "github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Queries *db.Queries


func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config")
	}

	conpool, err := pgxpool.New(context.Background(), config.DBSource)
	cypher := db.New(conpool)
	// api.NewServer(config, cypher)
	event := event.SingleTon()

	sub := event.On("event1")
	event.Emit("event1", "Hello Hi")

	// Receive and print events
	fmt.Println("Event 1:", <-sub)
	runGinServer(config, cypher)
	
}




func runGinServer(config util.Config, cypher db.Cypher) {
	server, err := api.NewServer(config, cypher)
	if err != nil {

		log.Fatal("cannot create server")
		// log.Fatal().Err(err).Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {

	log.Fatal("cannot start server")
		// log.Fatal().Err(err).Msg("cannot start server")
	}
	
}