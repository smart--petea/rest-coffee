package helper

import (
    "os"
    "github.com/go-pg/pg/v10"
)

func GetDb() *pg.DB {
    return pg.Connect(&pg.Options{
        Addr: "database:5432",
        User:  os.Getenv("POSTGRES_USER"),
        Password: os.Getenv("POSTGRES_PASSWORD"),
        Database: os.Getenv("POSTGRES_DB"),
    })
}
