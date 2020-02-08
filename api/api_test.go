package api

import (
	"json_rpc_server/db"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

var mockDB db.MockDB

func init() {
	mockDB = db.MockDB(make(map[string]*db.User))

	go func() {
		if err := Listen("127.0.0.1:8080", mockDB); err != nil {
			panic(err)
		}
	}()
}

func TestAdd(t *testing.T) {
	client, err := jsonrpc2.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var res int

	if err := client.Call("Users.Add", UserArgs{Login: "TEST"}, &res); err != nil {
		t.Fatal(err)
	}

	if res != 1 {
		t.Fatal("result must be 1")
	}

	user, exist := mockDB["TEST"]
	if !exist {
		t.Fatal("user not exit")
	}

	if user.Login != "TEST" || user.Date == "" && user.UUID.String() == "" {
		t.Fatal("get incorrect user params")
	}
}

func TestGet(t *testing.T) {
	client, err := jsonrpc2.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	mockDB["TEST"] = &db.User{
		UUID:  uuid.New(),
		Login: "TEST",
		Date:  time.Now().Format("2006-01-02"),
	}

	dbUser := mockDB["TEST"]
	apiUser := db.User{}

	if err := client.Call("Users.Get", UserArgs{Login: "TEST"}, &apiUser); err != nil {
		log.Fatal(err)
	}

	if apiUser.UUID != dbUser.UUID || apiUser.Date != dbUser.Date {
		t.Fatal("get incorrect user params")
	}
}

func TestSet(t *testing.T) {
	client, err := jsonrpc2.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	mockDB["TEST"] = &db.User{
		UUID:  uuid.New(),
		Login: "TEST",
		Date:  time.Now().Format("2006-01-02"),
	}

	old := mockDB["TEST"].UUID

	var res int

	if err := client.Call("Users.Set", UserArgs{Login: "TEST"}, &res); err != nil {
		log.Fatal(err)
	}

	if res != 1 {
		t.Fatal("result must be 1")
	}

	new := mockDB["TEST"].UUID

	if old == new {
		t.Fatal("uuid must be changed")
	}
}
