package databaseConfig

import (
	"fmt"
	"os"

	"github.com/OnlineShop/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection() *gorm.DB {

	//dsn := "root:S@eed1372144@tcp(127.0.0.1)/onlineshop?parseTime=true"

	dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

	dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

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
