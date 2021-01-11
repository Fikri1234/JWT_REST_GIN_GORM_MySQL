package repository

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/model"
	// Use prefix blank identifier _ when importing driver for its side
	// effect and not use it explicity anywhere in our code.
	// When a package is imported prefixed with a blank identifier,the init
	// function of the package will be called. Also, the GO compiler will
	// not complain if the package is not used anywhere in the code
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetMMenuByID ...
func GetMMenuByID(id int64) (model.MMenu, error) {
	db := configuration.DB

	var mMenu model.MMenu
	err := db.Debug().First(&mMenu, id).Error

	if err != nil {
		return mMenu, err
	}

	return mMenu, nil
}

// GetMMenuByMenuID ...
func GetMMenuByMenuID(menuID int64) ([]model.MMenu, error) {
	db := configuration.DB

	var mTreeMenus []model.MMenu
	err := db.Debug().Where("parent_menu_id = ?", menuID).Find(&mTreeMenus).Error

	if err != nil {
		return mTreeMenus, err
	}

	return mTreeMenus, nil
}

// GetMMenuByNotMenuID ...
func GetMMenuByNotMenuID(menuID int64) ([]model.MMenu, error) {
	db := configuration.DB

	var mTreeMenus []model.MMenu
	err := db.Debug().Where("parent_menu_id != ?", menuID).Find(&mTreeMenus).Error

	if err != nil {
		return mTreeMenus, err
	}

	return mTreeMenus, nil
}

// GetMMenuAll ...
func GetMMenuAll() ([]model.MMenu, error) {
	db := configuration.DB

	var mTreeMenus []model.MMenu

	err := db.Find(&mTreeMenus).Error

	if err != nil {
		return mTreeMenus, err
	}

	return mTreeMenus, nil
}

// CreateMMenu ...
func CreateMMenu(mMenu model.MMenu) (model.MMenu, error) {
	db := configuration.DB

	var err error

	err = db.Create(&mMenu).Error
	if err != nil {
		return mMenu, err
	}

	return mMenu, nil
}

// UpdateMMenu ...
func UpdateMMenu(mMenu model.MMenu) (model.MMenu, error) {
	db := configuration.DB

	var err error

	err = db.Save(&mMenu).Error
	if err != nil {
		return mMenu, err
	}

	// find mMenu by id
	res, err := GetMMenuByID(mMenu.ID)
	if err != nil {
		return mMenu, err
	}

	return res, nil
}

// DeleteMMenuByID ...
func DeleteMMenuByID(id int64) error {
	db := configuration.DB

	res, err := GetMMenuByID(id)
	if err != nil {
		return err
	}

	var errDel error

	errDel = db.Delete(&res).Error
	if errDel != nil {
		return errDel
	}

	return nil
}
