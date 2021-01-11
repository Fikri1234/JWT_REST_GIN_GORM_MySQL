package router

import (
	"JWT_REST_GIN_GORM_MySQL/service"

	"github.com/gin-gonic/gin"
)

// NewRoutes router global
func NewRoutes() *gin.Engine {

	router := gin.Default()
	v1 := router.Group("/api")

	// register router from each controller service
	service.RoutesBranch(v1)
	service.RoutesUser(v1)
	service.RoutesParam(v1)
	service.RoutesMenu(v1)
	service.RoutesMenuRole(v1)
	service.RoutesUserRole(v1)
	service.RoutesLoginLogout(v1)

	return router
}
