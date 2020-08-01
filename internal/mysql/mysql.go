package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"micro_service/config"
	"time"
)

var MysqlDB *sqlx.DB

func InitMysql() {
	db, err := sqlx.Open("mysql", config.Conf.GetString("mysql.address"))
	// 打印日志
	//db.LogMode(true)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(200)
	db.SetConnMaxLifetime(time.Hour)
	MysqlDB = db
	log.Println("Mysql is Collection!!!")
}

func GetMysqlDB() *sqlx.DB {
	return MysqlDB
}
