package util_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
	//"net/http"
	//"os"
	//"path/filepath"
)

type dbStore struct {
	db *sql.DB
}

func GenerateDBConnectionString(username string, password string, protocal string, ipaddr string, port string, dbname string) string {
	connString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, protocal, ipaddr, port, dbname)
	return connString
}

func Connect2MySql(connectionString string) (dbNew sql.DB, err error) {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return *db, err
}
