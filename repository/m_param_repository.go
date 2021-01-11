package repository

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/model"
	"strings"
)

// GetMParamByID ...
func GetMParamByID(id int64) (model.MParam, error) {
	db := configuration.DB

	var mParam model.MParam
	err := db.First(&mParam, id).Error

	if err != nil {
		return mParam, err
	}

	return mParam, nil
}

// GetMParamByKey ...
func GetMParamByKey(key string) (model.MParam, error) {
	db := configuration.DB

	var mParam model.MParam
	err := db.Debug().Where("param_key = ?", key).Find(&mParam).Error

	if err != nil {
		return mParam, err
	}

	return mParam, nil
}

// GetMParamByLikeKey ...
func GetMParamByLikeKey(key string) ([]model.MParam, error) {
	db := configuration.DB

	var mParams []model.MParam
	err := db.Where("LOWER(key) = ?", "%"+strings.ToLower(key)+"%").Find(&mParams).Error

	if err != nil {
		return mParams, err
	}

	return mParams, nil
}

// GetMParamAll ...
func GetMParamAll() ([]model.MParam, error) {
	db := configuration.DB

	var mParams []model.MParam

	err := db.Debug().Find(&mParams).Error

	if err != nil {
		return mParams, err
	}

	return mParams, nil
}

// CreateMParam ...
func CreateMParam(mParam model.MParam) (model.MParam, error) {
	db := configuration.DB

	var err error

	err = db.Create(&mParam).Error
	if err != nil {
		return mParam, err
	}

	return mParam, nil
}

// UpdateMParam ...
func UpdateMParam(mParam model.MParam) (model.MParam, error) {
	db := configuration.DB

	var err error

	err = db.Save(&mParam).Error
	if err != nil {
		return mParam, err
	}

	// find mParam by id
	res, err := GetMParamByID(mParam.ID)
	if err != nil {
		return mParam, err
	}

	return res, nil
}

// DeleteMParamByID ...
func DeleteMParamByID(id int64) error {
	db := configuration.DB

	res, err := GetMParamByID(id)
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
