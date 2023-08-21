package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDB() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln("cannnot get timezone " + ": " + err.Error())
	}
	conn := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", conn.FormatDSN())
	if err != nil {
		log.Fatalln()
	}
	Db = db
}

func CloseDb() {
	Db.Close()
}
