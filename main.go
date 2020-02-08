package main

import (
	"flag"
	"fmt"
	"json_rpc_server/api"
	"json_rpc_server/db"
	"log"
)

func main() {
	dbName := flag.String("db.name", "postgres", "database name")
	dbHost := flag.String("db.host", "127.0.0.1", "database host")
	dbPort := flag.Int("db.port", 5432, "databse port")
	dbUser := flag.String("db.user", "pgadmin", "database user")
	dbPass := flag.String("db.pass", "pgadmin", "database password")
	apiHost := flag.String("api.host", "127.0.0.1", "api host")
	apiPort := flag.Int("api.port", 8080, "api port")
	flag.Parse()

	db := db.New(*dbName, *dbHost, *dbUser, *dbPass, *dbPort)
	addr := fmt.Sprintf("%s:%d", *apiHost, *apiPort)

	if err := api.Listen(addr, db); err != nil {
		log.Fatal(err)
	}
}
