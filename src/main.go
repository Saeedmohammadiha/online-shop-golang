package main

import (
	initDb "github.com/OnlineShop/src/database"
	v1 "github.com/OnlineShop/src/http/v1"
	"github.com/OnlineShop/src/logger"
	"github.com/OnlineShop/src/repository"
	"github.com/OnlineShop/src/router"
)

func main() {

	Log := logger.New()
	defer Log.Sync()
	appRouter := router.New(Log)
	db := initDb.MysqlDatabaseConnection(Log)
	RepositoryFactory := repository.NewRepositoryFactory(db, Log)

	// register routes
	v1.RegisterRoutes(appRouter, RepositoryFactory, Log)

	appRouter.Serve(":5000")

}
