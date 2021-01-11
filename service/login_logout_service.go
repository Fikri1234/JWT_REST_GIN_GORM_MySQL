package service

import (
	"JWT_REST_GIN_GORM_MySQL/model"
	"JWT_REST_GIN_GORM_MySQL/repository"
	"JWT_REST_GIN_GORM_MySQL/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CredentialsLogin Create a struct to read the username and password from the request body
type CredentialsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RoutesLoginLogout ...
func RoutesLoginLogout(rg *gin.RouterGroup) {
	cred := rg.Group("/")

	cred.POST("login", getUserLogin)
	cred.GET("logout", getUserLogout)
}

func getUserLogin(c *gin.Context) {

	var creds CredentialsLogin
	var user model.MUser
	var err error

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	user, err = repository.GetUserLogin(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	jwt, err := util.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   jwt.AccessToken,
		Expires: time.Unix(jwt.AtExpires, 0),
	})

	c.JSON(http.StatusOK, jwt)
}

func getUserLogout(c *gin.Context) {

	cookies := http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-7 * 24 * time.Hour),
		MaxAge:  -1,
	}
	http.SetCookie(c.Writer, &cookies)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out!"})
}
