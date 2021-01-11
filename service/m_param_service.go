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

// RoutesParam ...
func RoutesParam(rg *gin.RouterGroup) {
	param := rg.Group("/param")

	param.GET("/by-id/:id", util.TokenAuthMiddleware(), getParamByID)
	param.GET("/", util.TokenAuthMiddleware(), getParams)
	param.POST("/", util.TokenAuthMiddleware(), createParam)
	param.PUT("/", util.TokenAuthMiddleware(), updateParam)
	param.DELETE("/:id", util.TokenAuthMiddleware(), deleteParamByID)
}

func getParamByID(c *gin.Context) {
	var param model.MParam
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	param, err = repository.GetMParamByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.DeepEqual(model.MParam{}, param) {
		c.JSON(http.StatusNotFound, param)
	} else {
		c.JSON(http.StatusOK, param)
	}
}

func getParams(c *gin.Context) {

	var err error

	var param []model.MParam
	param, err = repository.GetMParamAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, param)
}

func createParam(c *gin.Context) {

	var param model.MParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	param, err := repository.CreateMParam(param)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, param)
}

func updateParam(c *gin.Context) {

	var param model.MParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateMParam(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteParamByID(c *gin.Context) {

	var param model.MParam

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteMParamByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, param)
}
