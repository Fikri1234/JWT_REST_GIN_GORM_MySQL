package repository

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/model"
)

// GetMUserRoleByID ...
func GetMUserRoleByID(id int64) (model.MUserRole, error) {
	db := configuration.DB

	var mUserRole model.MUserRole
	err := db.First(&mUserRole, id).Error

	if err != nil {
		return mUserRole, err
	}

	return mUserRole, nil
}

// GetMUserRoleByRoleCode ...
func GetMUserRoleByRoleCode(roleCode string) (model.MUserRole, error) {
	db := configuration.DB

	var mUserRole model.MUserRole
	err := db.Where("role_code = ?", roleCode).Find(&mUserRole).Error

	if err != nil {
		return mUserRole, err
	}

	return mUserRole, nil
}

// GetMUserRoleAll ...
func GetMUserRoleAll() ([]model.MUserRole, error) {
	db := configuration.DB

	var mUserRoles []model.MUserRole

	err := db.Debug().Find(&mUserRoles).Error

	if err != nil {
		return mUserRoles, err
	}

	return mUserRoles, nil
}

// CreateMUserRole ...
func CreateMUserRole(mUserRole model.MUserRole) (model.MUserRole, error) {
	db := configuration.DB

	var err error

	err = db.Create(&mUserRole).Error
	if err != nil {
		return mUserRole, err
	}

	return mUserRole, nil
}

// UpdateMUserRole ...
func UpdateMUserRole(mUserRole model.MUserRole) (model.MUserRole, error) {
	db := configuration.DB

	var err error

	err = db.Save(&mUserRole).Error
	if err != nil {
		return mUserRole, err
	}

	// find mUserRole by id
	res, err := GetMUserRoleByID(mUserRole.ID)
	if err != nil {
		return mUserRole, err
	}

	return res, nil
}

// DeleteMUserRoleByID ...
func DeleteMUserRoleByID(id int64) error {
	db := configuration.DB

	res, err := GetMUserRoleByID(id)
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
