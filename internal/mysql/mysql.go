package mysql

import "github.com/jinzhu/gorm"

// 获取mysql实例
func getMysqlDB() *gorm.DB {
	return mysqlDB
}

type mysqlStruct struct {
	db *gorm.DB
}

func MCurd() *mysqlStruct {
	return &mysqlStruct{
		db: getMysqlDB(),
	}

}

//暂时定下来
func (db *mysqlStruct) Find()   {}
func (db *mysqlStruct) Insert() {}
func (db *mysqlStruct) Update() {}
func (db *mysqlStruct) Delete() {}
