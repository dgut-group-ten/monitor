package mydb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"monitor/core/util"
	"os"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:dgutdev#@tcp(music-01.niracler.com:3306)/music_db?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		util.Log.Println("Failed connect to mysql:" + err.Error())
		os.Exit(1)
	}
}

// 返回数据库连接对象
func DBConn() *sql.DB {
	return db
}
