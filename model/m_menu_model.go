package model

import (
	"time"

	"gorm.io/gorm"
)

// MMenu ...
type MMenu struct {
	ID           int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MenuName     string    `gorm:"size:255;" json:"menuName"`
	ParentMenuID int64     `json:"parentMenuID"`
	AccessLink   string    `json:"accessLink"`
	MenuChildren []MMenu   `gorm:"-" json:"menus"`
	CreatedBy    string    `gorm:"size:255;<-:create" json:"createdBy"`
	UpdatedBy    string    `gorm:"size:255" json:"updatedBy"`
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
