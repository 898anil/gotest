package jsonrpcserver

import (
	"fmt"

	"github.com/898anil/gotest/pkg/database"
	"github.com/898anil/gotest/pkg/rpcserver"
)

type GetUserArgs struct {
	Id int
}

type User struct{}

func (t *User) Get(args *GetUserArgs, reply *rpcserver.GetUserResponse) error {
	query := "SELECT id, email from auth_user where id=%d;"
	query = fmt.Sprintf(query, args.Id)
	rows, err := database.ExecuteMysql(query)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var username string
		err := rows.Scan(&id, &username)
		if err != nil {
			panic(err)
		}
		*reply = rpcserver.GetUserResponse{Id: id, Email: username}
	}
	return nil
}
