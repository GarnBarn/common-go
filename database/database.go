package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn(connectionString string) (*gorm.DB, error) {
	// Start DB Connection
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}
