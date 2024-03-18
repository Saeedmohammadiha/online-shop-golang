package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var lock = &sync.Mutex{}

type Database struct{}

var db *Database

func GetInstance() *Database {

	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating single instance now.")
			dsn := "root:S@eed1372144@tcp(127.0.0.1)/onlineshop"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			db = db
			if err != nil {
				fmt.Println("error", err)
			}
			db = &Database{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

}
