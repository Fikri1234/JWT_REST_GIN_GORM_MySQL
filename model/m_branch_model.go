package model

import (
	"time"

	"gorm.io/gorm"
)

// MBranch ...
type MBranch struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Code      string    `gorm:"size:255" json:"code"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Owner     string    `gorm:"size:255" json:"owner"`
	Address   string    `gorm:"size:255" json:"address"`
	Npwp      string    `gorm:"size:255" json:"npwp"`
	Flag      bool      `gorm:"default:true" json:"flag"`
	CreatedBy string    `gorm:"size:255;<-:create" json:"createdBy"`
	UpdatedBy string    `gorm:"size:255" json:"updatedBy"`
	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
