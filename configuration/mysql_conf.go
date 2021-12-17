package configuration

import (
	"JWT_REST_GIN_GORM_MySQL/model"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var username, password, hostname, dbname string

func initConfMySQL() {
	username = viper.GetString("DB.USER_NAME")
	password = viper.GetString("DB.PASSWORD")
	hostname = viper.GetString("DB.HOST_NAME")
	dbname = viper.GetString("DB.NAME")
}

// DB database global
var DB *gorm.DB

// SetupDB db
func SetupDB() (*gorm.DB, error) {
	initConfMySQL()

	db, err := gorm.Open(mysql.Open(confMysql(dbname)), &gorm.Config{})

	if err != nil {
		return db, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return db, err
	}

	db.AutoMigrate(&model.MBranch{}, &model.MUser{}, &model.MUserRole{},
		&model.MMenu{}, &model.MMenuRole{}, &model.MParam{})

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func confMysql(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, hostname, dbName)
}
