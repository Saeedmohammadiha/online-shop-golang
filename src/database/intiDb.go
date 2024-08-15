package initDb

import (
	"fmt"
	"os"
	"time"

	"github.com/OnlineShop/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection() *gorm.DB {
	// TODO: add retry logic

	dsn := getDbConfig()

	// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections.
	// Thus, the Open function should be called just once. It is rarely necessary to close a DB.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("create connection to database failed")
		// TODO: add panic
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)


	if err != nil {
		fmt.Println("failed to set config on db")
		// TODO: add panic

	} else {
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
	}
	return db
}

func getDbConfig() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
}
