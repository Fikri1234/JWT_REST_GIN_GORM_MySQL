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

// RoutesUser ...
func RoutesUser(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/by-id/:id", util.TokenAuthMiddleware(), getUserByID)
	user.GET("/by-user-role-id/:userRoleId", util.TokenAuthMiddleware(), getUsersByUserRoleID)
	user.GET("/", util.TokenAuthMiddleware(), getUsers)
	user.GET("", util.TokenAuthMiddleware(), getUsersPaging)
	user.POST("/", createUser)
	user.PUT("/", util.TokenAuthMiddleware(), updateUser)
	user.DELETE("/:id", util.TokenAuthMiddleware(), deleteUserByID)
}

func getUserByID(c *gin.Context) {
	var user model.MUser
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err = repository.GetMUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.DeepEqual(model.MUser{}, user) {
		c.JSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func getUsers(c *gin.Context) {

	var err error

	var users []model.MUser
	users, err = repository.GetMUserAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func getUsersPaging(c *gin.Context) {

	// fmt.Println(c.Query("Email"))
	// fmt.Println(c.QueryArray("BranchID"))

	// fmt.Println(strings.Join((c.QueryArray("BranchID")), ", "))

	var err error

	var user model.MUser
	user.UserName = c.Query("userName")
	user.Email = c.Query("email")
	user.MBranch.Name = c.Query("branchName")
	user.MUserRole.RoleCode = c.Query("roleCode")

	var users []model.MUser
	users, err = repository.GetMUserPaging(c.Request, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {

	var user model.MUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	user, err := repository.CreateMUser(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {

	var user model.MUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateMUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteUserByID(c *gin.Context) {

	var user model.MUser

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteMUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, user)
}

// getUsersByUserRoleID ...
func getUsersByUserRoleID(c *gin.Context) {

	var err error

	paramID := c.Param("userRoleId")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var user model.MUser
	user, err = repository.GetMUserByUserRoleID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
