package v1

import (
	"github.com/OnlineShop/router"
	"github.com/OnlineShop/services"
)

func RegisterPermissionRoutes(r router.IRouter, permissionService services.IPermissionService) {
	permissionRouter := r.CreateSubRouter("/permissions")
	permissionRouter.RegisterRoute("GET", "", permissionService.FindAll)
	permissionRouter.RegisterRoute("POST", "", permissionService.Create)
	permissionRouter.RegisterRoute("PUT", "", permissionService.Update)
	permissionRouter.RegisterRoute("GET", "/{id}", permissionService.FindById)
	permissionRouter.RegisterRoute("DELETE", "/{id}", permissionService.Delete)
}
