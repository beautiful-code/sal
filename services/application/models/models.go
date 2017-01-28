package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model
	Name   string `valid:"required"`
	UserId uint   `valid:"required"`
}

type Feedback struct {
	gorm.Model
	Desc          string `valid:"required" sql:"not null;type:text"`
	ApplicationId uint   `valid:"required"`
	Email         string
}
