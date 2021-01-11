package model

import (
	"time"

	"gorm.io/gorm"
)

// MUser ...
type MUser struct {
	ID                    int64     `gorm:"primary_key;auto_increment" json:"id"`
	UserName              string    `gorm:"size:255;not null;unique" json:"userName"`
	Password              string    `gorm:"size:255;not null" json:"password"`
	AccountNonExpired     bool      `gorm:"default:true" json:"accountNonExpired"`
	AccountNonLocked      bool      `gorm:"default:true" json:"accountNonLocked"`
	CredentialsNonExpired bool      `gorm:"default:true" json:"credentialsNonExpired"`
	Enabled               bool      `gorm:"default:true" json:"enabled"`
	Email                 string    `json:"email"`
	BranchID              int64     `gorm:"not null" json:"branchId"`
	MBranch               MBranch   `gorm:"foreignKey:BranchID"`
	UserRoleID            int64     `gorm:"not null" json:"userRoleId"`
	MUserRole             MUserRole `gorm:"foreignKey:UserRoleID"`
	ChangePwdCounter      int32     `json:"changePwdCounter"`
	FailCounter           int32     `json:"failCounter"`
	LastLogin             time.Time `json:"lastLogin"`
	ResetToken            string    `json:"resetToken"`
	CreatedBy             string    `gorm:"size:255;<-:create" json:"createdBy"`
	UpdatedBy             string    `gorm:"size:255" json:"updatedBy"`
	CreatedAt             time.Time `gorm:"<-:create"`
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
