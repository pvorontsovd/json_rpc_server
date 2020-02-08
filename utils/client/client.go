package main

import (
	"fmt"
	"json_rpc_server/api"
	"json_rpc_server/db"
	"log"

	"github.com/powerman/rpc-codec/jsonrpc2"
)

func main() {
	clientTCP, err := jsonrpc2.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer clientTCP.Close()

	var res int

	err = clientTCP.Call("Users.Add", api.UserArgs{Login: "TEST"}, &res)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("return code %d\n", res)

	user := db.User{}

	err = clientTCP.Call("Users.Get", api.UserArgs{Login: "TEST"}, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", user)

	err = clientTCP.Call("Users.Set", api.UserArgs{Login: "TEST"}, &res)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("return code %d\n", res)

	err = clientTCP.Call("Users.Get", api.UserArgs{Login: "TEST"}, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", user)
}
