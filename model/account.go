package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Uid  string
	Line string
}
