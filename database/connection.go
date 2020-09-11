package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SetConnection() *sql.DB {
	connectionDB, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/go_api_rest_basic")
	if err != nil {
		fmt.Println("Error to connect", err)
	}

	return connectionDB
}
