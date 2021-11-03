package models

import (
	"database/sql"
	"fmt"
	"go-lesson-webapp_211103/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	// idは自動で増分するよう設定。uuidはNULLを禁止してユニークに設定。
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		uuid STRING NOT NULL UNIQUE, 
		name STRING, 
		email STRING,
		password STRING, 
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

}
