package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	Name   string `valid:"required"`
	UserId int64  `valid:"required"`
}
