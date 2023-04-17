package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name          string
	Author        string
	Color         string
	ReminderTime  string
	Subscriber    string
	SecretKeyTotp string
}
