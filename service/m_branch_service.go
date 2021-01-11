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

// RoutesBranch ...
func RoutesBranch(rg *gin.RouterGroup) {
	branch := rg.Group("/branch")

	branch.GET("/by-id/:id", util.TokenAuthMiddleware(), getBranchByID)
	branch.GET("/", util.TokenAuthMiddleware(), getBranchs)
	branch.POST("/", util.TokenAuthMiddleware(), createBranch)
	branch.PUT("/", util.TokenAuthMiddleware(), updateBranch)
	branch.DELETE("/:id", util.TokenAuthMiddleware(), deleteBranchByID)
}

func getBranchByID(c *gin.Context) {
	var branch model.MBranch
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	branch, err = repository.GetMBranchByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.DeepEqual(model.MBranch{}, branch) {
		c.JSON(http.StatusNotFound, branch)
	} else {
		c.JSON(http.StatusOK, branch)
	}
}

func getBranchs(c *gin.Context) {

	var err error

	var branch []model.MBranch
	branch, err = repository.GetMBranchAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, branch)
}

func createBranch(c *gin.Context) {

	var branch model.MBranch

	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	branch, err := repository.CreateMBranch(branch)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, branch)
}

func updateBranch(c *gin.Context) {

	var branch model.MBranch

	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateMBranch(branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteBranchByID(c *gin.Context) {

	var branch model.MBranch

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteMBranchByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, branch)
}
