package service

import (
	"JWT_REST_GIN_GORM_MySQL/model"
	"JWT_REST_GIN_GORM_MySQL/repository"
	"JWT_REST_GIN_GORM_MySQL/util"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoutesMenu ...
func RoutesMenu(rg *gin.RouterGroup) {
	menu := rg.Group("/menu")

	menu.GET("/by-id/:id", util.TokenAuthMiddleware(), getMenuByID)
	menu.GET("/by-menuid/:id", util.TokenAuthMiddleware(), getMenuByMenuID)
	menu.GET("/by-rolecode/:roleCode", util.TokenAuthMiddleware(), getTreeMenu)
	menu.POST("/", util.TokenAuthMiddleware(), createMenu)
	menu.PUT("/", util.TokenAuthMiddleware(), updateMenu)
	menu.DELETE("/:id", util.TokenAuthMiddleware(), deleteMenuByID)
}

func getMenuByID(c *gin.Context) {
	var menu model.MMenu
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	menu, err = repository.GetMMenuByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.DeepEqual(model.MMenu{}, menu) {
		c.JSON(http.StatusNotFound, menu)
	} else {
		c.JSON(http.StatusOK, menu)
	}
}

func getMenuByMenuID(c *gin.Context) {
	var menus []model.MMenu
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	menus, err = repository.GetMMenuByMenuID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for index, element := range menus {
		log.Println(index, "--", element.MenuName)
		log.Println(index, "--", element.ParentMenuID)
	}

	c.JSON(http.StatusOK, menus)
}

func getTreeMenu(c *gin.Context) {

	var err error

	paramID := c.Param("roleCode")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// find base parent menu role
	var roleMenuBases []model.MMenuRole
	roleMenuBases, err = repository.GetMMenuRoleByUserRoleIDAndPid(varID, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// find child menu role
	var roleMenuNotBases []model.MMenuRole
	roleMenuNotBases, err = repository.GetMMenuRoleByUserRoleIDAndPidNOT(varID, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var menuBases []model.MMenu
	for idx := range roleMenuBases {
		var menuChild = iterateMenus(roleMenuNotBases, roleMenuBases[idx].MMenu.ID)
		menuBases = append(menuBases, roleMenuBases[idx].MMenu)
		menuBases[idx].MenuChildren = menuChild
	}

	c.JSON(http.StatusOK, menuBases)
}

func iterateMenus(menus []model.MMenuRole, pID int64) []model.MMenu {
	var result []model.MMenu

	for _, menu := range menus {
		var menuID = menu.MMenu.ID
		var parentID = menu.MMenu.ParentMenuID

		if parentID == pID {
			var menuChild = iterateMenus(menus, menuID)
			menu.MMenu.MenuChildren = menuChild
			result = append(result, menu.MMenu)
		}
	}

	return result
}

func createMenu(c *gin.Context) {

	var menu model.MMenu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	menu, err := repository.CreateMMenu(menu)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, menu)
}

func updateMenu(c *gin.Context) {

	var menu model.MMenu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid json"})
		return
	}

	usr, err := repository.UpdateMMenu(menu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func deleteMenuByID(c *gin.Context) {

	var menu model.MMenu

	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = repository.DeleteMMenuByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, menu)
}
