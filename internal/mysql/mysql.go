package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"micro_service/config"
	"time"
)

var (
	mysqlDB *sqlx.DB
	err     error
)

func InitMysql() {
	mysqlDB, err = sqlx.Open("mysql", config.GetConf().Mysql.Address)
	// 打印日志
	if err != nil {
		panic(err)
	}
	err = mysqlDB.Ping()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(200)
	mysqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("Mysql is Collection!!!")
}

func GetMysqlDB() *sqlx.DB {
	return mysqlDB
}
