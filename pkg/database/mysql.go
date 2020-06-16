package database

import (
	"database/sql"
	"fmt"

	"github.com/898anil/gotest/pkg/common/logger"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ExecuteMysql(query string) (*sql.Rows, error) {
	logger.Log.Print(query)
	return db.Query(query)
}

func init() {
	connection, err := sql.Open("mysql", "root:evolution@tcp(127.0.0.1:3306)/navi")
	db = connection
	if err != nil {
		fmt.Println(err)
	}
}
