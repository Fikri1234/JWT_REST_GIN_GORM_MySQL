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

// RoutesUserRole ...
func RoutesUserRole(rg *gin.RouterGroup) {
	userRole := rg.Group("/user-role")

	userRole.GET("/by-id/:id", util.TokenAuthMiddleware(), getUserRoleByID)
	userRole.GET("/by-role-code/:roleCode", util.TokenAuthMiddleware(), getUserRolesByRoleCode)
	userRole.GET("/", util.TokenAuthMiddleware(), getUserRoles)
	userRole.POST("/", util.TokenAuthMiddleware(), createUserRole)
	userRole.PUT("/", util.TokenAuthMiddleware(), updateUserRole)
	userRole.DELETE("/:id", util.TokenAuthMiddleware(), deleteUserRoleByID)
}

func getUserRoleByID(c *gin.Context) {
	var userRole model.MUserRole
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userRole, err = repository.GetMUserRoleByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.DeepEqual(model.MUserRole{}, userRole) {
		c.JSON(http.StatusNotFound, userRole)
	} else {
		c.JSON(http.StatusOK, userRole)
	}
}

func getUserRoles(c *gin.Context) {

	var err error

	var userRole []model.MUserRole
	userRole, err = repository.GetMUserRoleAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRole)
}

func createUserRole(c *gin.Context) {

	var userRole model.MUserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	userRole, err := repository.CreateMUserRole(userRole)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userRole)
}

func updateUserRole(c *gin.Context) {

	var userRole model.MUserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateMUserRole(userRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteUserRoleByID(c *gin.Context) {

	var userRole model.MUserRole

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteMUserRoleByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, userRole)
}

// getUserRolesByRoleCode ...
func getUserRolesByRoleCode(c *gin.Context) {

	var err error

	paramID := c.Param("roleCode")

	var userRole model.MUserRole
	userRole, err = repository.GetMUserRoleByRoleCode(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userRole)
}
