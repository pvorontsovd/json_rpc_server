package main

import (
	"json_rpc_server/api"
	"json_rpc_server/db"
	"log"
)

func main() {
	p := db.New(
		"postgres",
		"127.0.0.1",
		"pgadmin",
		"pgadmin",
		54321,
	)

	if err := p.Open(); err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	if err := p.Init(); err != nil {
		log.Fatal(err)
	}

	api := api.New("127.0.0.1:8080", p)

	if err := api.Listen(); err != nil {
		log.Fatal(err)
	}
}
