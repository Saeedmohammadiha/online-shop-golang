package main

import (
	initDb "github.com/OnlineShop/internal/app/database"
	v1 "github.com/OnlineShop/internal/app/http/v1"
	"github.com/OnlineShop/internal/pkg/logger"

	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
)



func main() {
	
	Log := logger.Logger()
	defer Log.Sync()

	db := initDb.MysqlDatabaseConnection()
	RepositoryFactory := repositories.NewRepositoryFactory(db)

	appRouter := router.New()


	// register routes
	v1.RegisterRoutes(appRouter, RepositoryFactory)

	appRouter.Serve(":5000")

}
