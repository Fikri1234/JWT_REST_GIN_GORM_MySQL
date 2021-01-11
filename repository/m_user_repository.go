package repository

import (
	"JWT_REST_GIN_GORM_MySQL/configuration"
	"JWT_REST_GIN_GORM_MySQL/model"
	"JWT_REST_GIN_GORM_MySQL/scopes"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"golang.org/x/crypto/bcrypt"
)

var queryPaging = "(user_name IS NULL OR user_name = '' OR LOWER(user_name) like LOWER(?)) AND " +
	"(email IS NULL OR email = '' OR LOWER(email) like LOWER(?)) AND " +
	"(MBranch.name IS NULL OR MBranch.name = '' OR LOWER(MBranch.name) like LOWER(?)) AND " +
	"(MUserRole.role_code IS NULL OR MUserRole.role_code = '' OR LOWER(MUserRole.role_code) like LOWER(?)) AND " +
	"enabled = ? "

// GetMUserByID ...
func GetMUserByID(id int64) (model.MUser, error) {
	db := configuration.DB

	var mUser model.MUser
	err := db.Joins("MBranch").Joins("MUserRole").First(&mUser, id).Error

	if err != nil {
		return mUser, err
	}

	return mUser, nil
}

// GetUserLogin ...
func GetUserLogin(username string, password string) (model.MUser, error) {
	var user model.MUser
	var err error

	user, err = GetMUserByUserName(username)
	if err != nil {
		return user, err
	}

	if reflect.DeepEqual(model.MUser{}, user) {
		return user, errors.New("Bad credential")
	}

	retVal := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if retVal != nil {
		return user, errors.New("Wrong password")
	}

	return user, nil
}

// GetMUserByUserName ...
func GetMUserByUserName(userName string) (model.MUser, error) {
	db := configuration.DB

	var mUser model.MUser
	err := db.Joins("MBranch").Joins("MUserRole").Where("enabled = ? AND user_name = ?", true, userName).Find(&mUser).Error

	if err != nil {
		return mUser, err
	}

	if mUser.AccountNonExpired == false {
		mUser = model.MUser{}
		return mUser, errors.New("Account Expired")
	}

	if mUser.AccountNonLocked == false {
		mUser = model.MUser{}
		return mUser, errors.New("Account Locked")
	}

	if mUser.CredentialsNonExpired == false {
		mUser = model.MUser{}
		return mUser, errors.New("Credential Expired")
	}

	return mUser, nil
}

// GetMUserByUserRoleID ...
func GetMUserByUserRoleID(roleCode int64) (model.MUser, error) {
	db := configuration.DB

	var mUser model.MUser
	err := db.Joins("MBranch").Joins("MUserRole").Where("user_role_id = ?", roleCode).Find(&mUser).Error

	if err != nil {
		return mUser, err
	}

	return mUser, nil
}

// GetMUserAll ...
func GetMUserAll() ([]model.MUser, error) {
	db := configuration.DB

	var mUsers []model.MUser

	err := db.Joins("MBranch").Joins("MUserRole").Where("enabled = ? ", true).Find(&mUsers).Error

	if err != nil {
		return mUsers, err
	}

	return mUsers, nil
}

// GetMUserPaging ...
func GetMUserPaging(r *http.Request, mUser model.MUser) ([]model.MUser, error) {
	db := configuration.DB

	var mUsers []model.MUser

	fmt.Println("nan: ")

	err := db.Scopes(scopes.ScopeDBPaginate(r)).Joins("MBranch").Joins("MUserRole").Where(queryPaging,
		"%"+mUser.UserName+"%", "%"+mUser.Email+"%", "%"+mUser.MBranch.Name+"%",
		"%"+mUser.MUserRole.RoleCode+"%", true).Find(&mUsers).Error

	if err != nil {
		return mUsers, err
	}

	return mUsers, nil
}

// CreateMUser ...
func CreateMUser(mUser model.MUser) (model.MUser, error) {
	db := configuration.DB

	var err error

	// hash, _ := util.HashPassword(mUser.Password, bcrypt.DefaultCost)
	// mUser.Password = hash

	bytes, err := bcrypt.GenerateFromPassword([]byte(mUser.Password), bcrypt.DefaultCost)
	mUser.Password = string(bytes)

	err = db.Create(&mUser).Error
	if err != nil {
		return mUser, err
	}

	return mUser, nil
}

// UpdateMUser ...
func UpdateMUser(mUser model.MUser) (model.MUser, error) {
	db := configuration.DB

	var err error

	var user model.MUser
	err = db.First(&user, mUser.ID).Error

	if err != nil {
		return mUser, err
	}

	// check update password
	retVal := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(mUser.Password))
	if retVal != nil {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(mUser.Password), bcrypt.DefaultCost)
		mUser.Password = string(bytes)
	} else {
		mUser.Password = user.Password
	}

	err = db.Save(&mUser).Error
	if err != nil {
		return mUser, err
	}

	// find mUser by id
	res, err := GetMUserByID(mUser.ID)
	if err != nil {
		return mUser, err
	}

	return res, nil
}

// DeleteMUserByID ...
func DeleteMUserByID(id int64) error {
	db := configuration.DB

	res, err := GetMUserByID(id)
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
