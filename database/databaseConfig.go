package databaseConfig

import (
	"fmt"

	"github.com/OnlineShop/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection() *gorm.DB {

	dsn := "root:S@eed1372144@tcp(127.0.0.1)/onlineshop?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("create connection to database failed")
	}
	db.AutoMigrate(&models.User{},
		&models.Product{},
		&models.Order{},
		&models.Address{},
		&models.Comment{},
		&models.Discount{},
		&models.OrderItem{},
		&models.OrderStatus{},
		&models.Permission{},
		&models.Role{},
		&models.Score{},
		&models.TransactionStatus{},
		&models.Transaction{},
	)
	return db
}
