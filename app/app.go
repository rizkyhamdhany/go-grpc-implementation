package app

import "github.com/jinzhu/gorm"

type App struct {
	DB *gorm.DB
}