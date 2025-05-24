package app

import (
	"database/sql"
	"fmt"
	"golang-restfull/helper"
	"time"
)

func NewDB() *sql.DB {
	fmt.Println("connecting to db")
	db, err := sql.Open("mysql", "root:basmalah@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.PanicIfError(err)

	fmt.Println("connection success")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
