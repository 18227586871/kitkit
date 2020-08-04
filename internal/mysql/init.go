package mysql

import (
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
	mysqlDB.SetMaxIdleConns(config.GetConf().Mysql.MaxIdleConn)
	mysqlDB.SetMaxOpenConns(config.GetConf().Mysql.MaxOpenConn)
	mysqlDB.SetConnMaxLifetime(time.Second * 5)
	log.Println("Mysql is Collection!!!")
}
