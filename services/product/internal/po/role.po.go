package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; unique; primary_key; AUTO_INCREMENT"`
	RoleName string `gorm:"column:role_name; type:varchar(255); not null;"`
	RoleCode string `gorm:"column:role_code; type:varchar(255); not null;"`
	RoleNote string `gorm:"column:role_note; type:varchar(1000); not null;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
