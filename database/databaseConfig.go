package databaseConfig

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection() *gorm.DB {

	dsn := "root:S@eed1372144@tcp(127.0.0.1)/onlineshop"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("create connection to database failed")
	}
	return db
}
