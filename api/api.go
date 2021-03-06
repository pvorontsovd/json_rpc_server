package api

import (
	"context"
	"fmt"
	"json_rpc_server/db"
	"net"
	"net/rpc"
	"time"

	"github.com/google/uuid"
	"github.com/powerman/rpc-codec/jsonrpc2"
)

type Users struct {
	db db.DB
}

type UserArgs struct {
	Login string
}

type UserResp struct {
	UUID  string
	Login string
	Date  string
}

func (u *Users) Add(args *UserArgs, result *int) error {
	if args.Login == "" {
		return fmt.Errorf("need login")
	}

	if err := u.db.Add(&db.User{
		UUID:  uuid.New(),
		Login: args.Login,
		Date:  time.Now().Format("2006-01-02"),
	}); err != nil {
		return err
	}

	*result = 1

	return nil
}

func (u *Users) Get(args *UserArgs, result *UserResp) error {
	if args.Login == "" {
		return fmt.Errorf("need login")
	}

	user, err := u.db.Get(args.Login)
	if err != nil {
		return err
	}

	result.UUID = user.UUID.String()
	result.Login = user.Login
	result.Date = user.Date

	return nil
}

func (u *Users) Set(args *UserArgs, result *int) error {
	if args.Login == "" {
		return fmt.Errorf("need login")
	}

	if err := u.db.Set(&db.User{
		UUID:  uuid.New(),
		Login: args.Login,
		Date:  time.Now().Format("2006-01-02"),
	}); err != nil {
		return err
	}

	*result = 1

	return nil
}

type contextKey string

var RemoteAddrContextKey contextKey = "RemoteAddr"

func Listen(addr string, db db.DB) error {
	if err := rpc.Register(&Users{db: db}); err != nil {
		return err
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		ctx := context.WithValue(context.Background(), RemoteAddrContextKey, conn.RemoteAddr())
		go jsonrpc2.ServeConnContext(ctx, conn)
	}
}
