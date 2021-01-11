package service

import (
	"JWT_REST_GIN_GORM_MySQL/model"
	"JWT_REST_GIN_GORM_MySQL/repository"
	"JWT_REST_GIN_GORM_MySQL/util"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoutesMenuRole ...
func RoutesMenuRole(rg *gin.RouterGroup) {
	menuRole := rg.Group("/menu-role")

	menuRole.GET("/by-id/:id", util.TokenAuthMiddleware(), getMenuRoleByID)
	menuRole.GET("/", util.TokenAuthMiddleware(), getMenuRoles)
	menuRole.POST("/", util.TokenAuthMiddleware(), createMenuRole)
	menuRole.PUT("/", util.TokenAuthMiddleware(), updateMenuRole)
	menuRole.DELETE("/:id", util.TokenAuthMiddleware(), deleteMenuRoleByID)
	menuRole.GET("/by-role-id/:roleId", util.TokenAuthMiddleware(), getMenuRolesByRoleID)
}

func getMenuRoleByID(c *gin.Context) {
	var menu model.MMenuRole
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	menu, err = repository.GetMMenuRoleByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.DeepEqual(model.MMenuRole{}, menu) {
		c.JSON(http.StatusNotFound, menu)
	} else {
		c.JSON(http.StatusOK, menu)
	}
}

func getMenuRoles(c *gin.Context) {

	var err error

	// find child menu
	var menuRoles []model.MMenuRole
	menuRoles, err = repository.GetMMenuRoleAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menuRoles)
}

func createMenuRole(c *gin.Context) {

	var menu model.MMenuRole

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	menu, err := repository.CreateMMenuRole(menu)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}

func updateMenuRole(c *gin.Context) {

	var menu model.MMenuRole

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateMMenuRole(menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteMenuRoleByID(c *gin.Context) {

	var menu model.MMenuRole

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteMMenuRoleByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, menu)
}

// getMenuRolesByRoleID ...
func getMenuRolesByRoleID(c *gin.Context) {

	var err error

	paramID := c.Param("roleId")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// find child menu
	var menuRoles []model.MMenuRole
	menuRoles, err = repository.GetMMenuRoleByUserRoleID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menuRoles)
}
