package config

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/transaction_service")
    if err != nil {
        log.Fatal(err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatal(err)
    }
}
