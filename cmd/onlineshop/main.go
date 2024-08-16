package main

import (
	initDb "github.com/OnlineShop/internal/app/database"
	v1 "github.com/OnlineShop/internal/app/http/v1"
	"github.com/OnlineShop/internal/pkg/logger"

	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
)



func main() {
	
	Log := logger.New()
	defer Log.Sync()

	db := initDb.MysqlDatabaseConnection(Log)
	RepositoryFactory := repositories.NewRepositoryFactory(db, Log)

	appRouter := router.New(Log)
	middles := middlewares.New(Log)
	appRouter.Use(middles.SanitizeURLParamsMiddleware)
	appRouter.Use(middles.SanitizeBodyMiddleware)

	// register routes
	v1.RegisterRoutes(appRouter, RepositoryFactory, Log)

	appRouter.Serve(":5000")

}
