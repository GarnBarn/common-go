package model

import "gorm.io/gorm"

type Assignment struct {
	gorm.Model
	Name         string
	Author       string
	Description  string
	ReminderTime string
	DueDate      int
	TagID        int
	Tag          *Tag
}
