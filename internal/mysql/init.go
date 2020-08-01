package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"micro_service/config"
	"time"
)

var mysqlDB *gorm.DB

func InitMysql() {
	db, err := gorm.Open("mysql", config.Conf.GetString("mysql.address"))
	// 打印日志
	//db.LogMode(true)
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)
	db.DB().SetConnMaxLifetime(time.Hour)
	mysqlDB = db
	log.Println("Mysql is Collection!!!")
}

