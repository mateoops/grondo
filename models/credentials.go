package models

import (
	"gorm.io/gorm"
)

type UserWithPassword struct {
	gorm.Model
	Username string
	Password string
}

type UserWithSSHKey struct {
	gorm.Model
	Username   string
	SSHKey     string
	Passphrase string
}
