package model

import (
	"time"

	"gorm.io/gorm"
)

// MUserRole ...
type MUserRole struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	RoleCode    string    `gorm:"size:255;" json:"roleCode"`
	Description string    `gorm:"size:255;" json:"description"`
	Flag        bool      `gorm:"default:true" json:"flag"`
	CreatedBy   string    `gorm:"size:255;<-:create" json:"createdBy"`
	UpdatedBy   string    `gorm:"size:255" json:"updatedBy"`
	CreatedAt   time.Time `gorm:"<-:create"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
