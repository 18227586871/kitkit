package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"micro_service/config"
	"time"
)

var mysqlDB *sqlx.DB

func InitMysql() {
	db, err := sqlx.Open("mysql", config.GetConf().Mysql.Address)
	// 打印日志
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
	mysqlDB = db
	log.Println("Mysql is Collection!!!")
}

func GetMysqlDB() *sqlx.DB {
	return mysqlDB
}
