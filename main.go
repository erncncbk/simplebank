package main

import (
	"database/sql"
	"log"

	"github.com/erncncbk/simplebank/api"
	db "github.com/erncncbk/simplebank/db/sqlc"
	"github.com/erncncbk/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load congig: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)

	}

}
