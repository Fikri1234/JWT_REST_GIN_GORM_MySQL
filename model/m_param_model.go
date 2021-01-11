package model

import (
	"time"

	"gorm.io/gorm"
)

// MParam ...
type MParam struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	ParamKey    string    `gorm:"size:255;not null" json:"paramKey"`
	ParamValue  string    `gorm:"size:255" json:"paramValue"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedBy   string    `gorm:"size:255;<-:create" json:"createdBy"`
	UpdatedBy   string    `gorm:"size:255" json:"updatedBy"`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
