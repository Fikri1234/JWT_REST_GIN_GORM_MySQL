package repository

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/model"
	"strings"
)

// GetMBranchByID ...
func GetMBranchByID(id int64) (model.MBranch, error) {
	db := configuration.DB

	var mBranch model.MBranch
	err := db.First(&mBranch, id).Error

	if err != nil {
		return mBranch, err
	}

	return mBranch, nil
}

// GetMBranchByLikeCode ...
func GetMBranchByLikeCode(code string) (model.MBranch, error) {
	db := configuration.DB

	var mBranch model.MBranch
	err := db.Where("LOWER(code) LIKE ?", "%"+strings.ToLower(code)+"%").Find(&mBranch).Error

	if err != nil {
		return mBranch, err
	}

	return mBranch, nil
}

// GetMBranchByLikeName ...
func GetMBranchByLikeName(name string) (model.MBranch, error) {
	db := configuration.DB

	var mBranch model.MBranch
	err := db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").Find(&mBranch).Error

	if err != nil {
		return mBranch, err
	}

	return mBranch, nil
}

// GetMBranchByLikeOwner ...
func GetMBranchByLikeOwner(owner string) (model.MBranch, error) {
	db := configuration.DB

	var mBranch model.MBranch
	err := db.Where("LOWER(owner) LIKE ?", "%"+strings.ToLower(owner)+"%").Find(&mBranch).Error

	if err != nil {
		return mBranch, err
	}

	return mBranch, nil
}

// GetMBranchAll ...
func GetMBranchAll() ([]model.MBranch, error) {
	db := configuration.DB

	var mBranchs []model.MBranch

	err := db.Debug().Find(&mBranchs).Error

	if err != nil {
		return mBranchs, err
	}

	return mBranchs, nil
}

// CreateMBranch ...
func CreateMBranch(mBranch model.MBranch) (model.MBranch, error) {
	db := configuration.DB

	var err error

	err = db.Create(&mBranch).Error
	if err != nil {
		return mBranch, err
	}

	return mBranch, nil
}

// UpdateMBranch ...
func UpdateMBranch(mBranch model.MBranch) (model.MBranch, error) {
	db := configuration.DB

	var err error

	err = db.Save(&mBranch).Error
	if err != nil {
		return mBranch, err
	}

	// find mBranch by id
	res, err := GetMBranchByID(mBranch.ID)
	if err != nil {
		return mBranch, err
	}

	return res, nil
}

// DeleteMBranchByID ...
func DeleteMBranchByID(id int64) error {
	db := configuration.DB

	res, err := GetMBranchByID(id)
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
