package model

// MMenuRole ...
type MMenuRole struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	MMenuID     int64     `gorm:"not null" json:"menuId"`
	MMenu       MMenu     `gorm:"foreignKey:MMenuID" json:"menu"`
	MUserRoleID int64     `gorm:"not null" json:"userRoleId"`
	MUserRole   MUserRole `gorm:"foreignKey:MUserRoleID" json:"userRole"`
}
