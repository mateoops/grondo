package models

import (
	"gorm.io/gorm"
)

type Machine struct {
	gorm.Model
	Address            string
	Port               uint
	UserWithPasswordID uint
	UserWithPassword   UserWithPassword
	UserWithSSHKeyID   uint
	UserWithSSHKey     UserWithSSHKey
}
