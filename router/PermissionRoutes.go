package router

import (
	"github.com/OnlineShop/services"
)

func RegisterPermissionRoutes(r IRouter, s *services.IPermissionService) {
	PermissionRoutes := r.AddPrefix("/permission")
	PermissionRoutes.Post("", (*s).Create)
	PermissionRoutes.Get("", (*s).FindAll)
	PermissionRoutes.Get("/{id}", (*s).FindById)
	PermissionRoutes.Delete("/{id}", (*s).Delete)
	PermissionRoutes.Put("", (*s).Delete)
}
