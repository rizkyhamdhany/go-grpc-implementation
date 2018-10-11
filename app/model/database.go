package model

import "github.com/jinzhu/gorm"

//Setup Auto Migration
func DBMigrate(db *gorm.DB) *gorm.DB {
	return db
}