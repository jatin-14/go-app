package apis

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (db *sql.DB){
	db, err := sql.Open("mysql", "root:mysqlroot@/imdb")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
