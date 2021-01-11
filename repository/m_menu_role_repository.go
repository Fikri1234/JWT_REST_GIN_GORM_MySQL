package repository

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/model"
)

// GetMMenuRoleByID ...
func GetMMenuRoleByID(id int64) (model.MMenuRole, error) {
	db := configuration.DB

	var mMenuRole model.MMenuRole
	err := db.Joins("MUserRole").Joins("MMenu").First(&mMenuRole, id).Error

	if err != nil {
		return mMenuRole, err
	}

	return mMenuRole, nil
}

// GetMMenuRoleByUserRoleID ...
func GetMMenuRoleByUserRoleID(userRoleID int64) ([]model.MMenuRole, error) {
	db := configuration.DB

	var mMenus []model.MMenuRole
	err := db.Where("m_user_role_id = ?", userRoleID).Find(&mMenus).Error

	if err != nil {
		return mMenus, err
	}

	return mMenus, nil
}

// GetMMenuRoleByUserRoleIDAndPid ...
func GetMMenuRoleByUserRoleIDAndPid(userRoleID int64, menuID int64) ([]model.MMenuRole, error) {
	db := configuration.DB

	var mMenus []model.MMenuRole
	err := db.Joins("MUserRole").Joins("MMenu").Where("MUserRole.id = ? AND MMenu.parent_menu_id = ?", userRoleID, menuID).Find(&mMenus).Error

	if err != nil {
		return mMenus, err
	}

	return mMenus, nil
}

// GetMMenuRoleByUserRoleIDAndPidNOT ...
func GetMMenuRoleByUserRoleIDAndPidNOT(userRoleID int64, menuID int64) ([]model.MMenuRole, error) {
	db := configuration.DB

	var mMenus []model.MMenuRole
	err := db.Joins("MUserRole").Joins("MMenu").Where("MUserRole.id = ? AND MMenu.parent_menu_id != ?", userRoleID, menuID).Find(&mMenus).Error

	if err != nil {
		return mMenus, err
	}

	return mMenus, nil
}

// GetMMenuRoleAll ...
func GetMMenuRoleAll() ([]model.MMenuRole, error) {
	db := configuration.DB

	var mTreeMenus []model.MMenuRole

	err := db.Debug().Find(&mTreeMenus).Error

	if err != nil {
		return mTreeMenus, err
	}

	return mTreeMenus, nil
}

// CreateMMenuRole ...
func CreateMMenuRole(mMenuRole model.MMenuRole) (model.MMenuRole, error) {
	db := configuration.DB

	var err error

	err = db.Create(&mMenuRole).Error
	if err != nil {
		return mMenuRole, err
	}

	return mMenuRole, nil
}

// UpdateMMenuRole ...
func UpdateMMenuRole(mMenuRole model.MMenuRole) (model.MMenuRole, error) {
	db := configuration.DB

	var err error

	err = db.Save(&mMenuRole).Error
	if err != nil {
		return mMenuRole, err
	}

	// find mMenuRole by id
	res, err := GetMMenuRoleByID(mMenuRole.ID)
	if err != nil {
		return mMenuRole, err
	}

	return res, nil
}

// DeleteMMenuRoleByID ...
func DeleteMMenuRoleByID(id int64) error {
	db := configuration.DB

	res, err := GetMMenuRoleByID(id)
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
