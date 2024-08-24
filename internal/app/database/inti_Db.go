package initDb

import (
	"fmt"
	"os"
	"time"

	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection() *gorm.DB {
	// TODO: add retry logic
	log := logger.Logger()
	dsn := getDbConfig()
	log.Info("got the dsn", dsn)

	// The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections.
	// Thus, the Open function should be called just once. It is rarely necessary to close a DB.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		// after logging it calls os.Exit(1)
		log.Fatalf("Failed to connect to database: %v", err)

	}

	log.Info("opened conection to datebase")
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("could not set the configs to sqldb: %v", err)

	} else {

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		log.Info("the configuration is set to the database connection",
			"MaxIdleConns", 10,
			"MaxOpenConns", 100,
			"ConnMaxLifetime", time.Hour,
		)

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

		log.Info("database migration is done")
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
