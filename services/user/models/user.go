package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Email     string `gorm:"unique_index:user_email_index" valid:"email,required"`
	Password  string `valid:"required"`
}
