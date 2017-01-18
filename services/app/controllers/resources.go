package controllers

import (
	"github.com/beautiful-code/sal/services/user/models"
)

//Models for JSON resources
type (
	// UserResource For Post - /register
	UserResource struct {
		Data UserModel `json:"data"`
	}

	// AuthUserResource Response for authorized user Post - /login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	// UserModel reperesents a user
	UserModel struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	// AuthUserModel for authorized user with access token
	AuthUserModel struct {
		User  model.User `json:"user"`
		Token string     `json:"token"`
	}
)
