package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/user/models"
)

// Register add a new User document
// Handler for HTTP Post - "/register"
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource

	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}

	userModel := dataResource.Data
	dataStore := common.NewDataStore()
	defer dataStore.Close()

	hpass, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)

	if err != nil {
		panic("bcrypt err!")
	}

	user := model.User{
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Email:     userModel.Email,
		Password:  string(hpass),
	}

	valid, err := govalidator.ValidateStruct(user)

	if valid {
		// TODO: Handle the errors
		// Create User record
		result := dataStore.Session.Create(&user)
		err := "User record not saved"

		if result.Error != nil {
			if strings.ContainsAny(result.Error.Error(), "Duplicate entry & for key 'user_email_index'") {
				err = "Email address is already taken!"
			}
		}

		if dataStore.Session.NewRecord(&user) {
			common.DisplayAppError(
				w,
				errors.New(err),
				"Failed to write to the database.",
				500,
			)
			return
		}
	} else {
		common.DisplayAppError(
			w,
			err,
			"Unprocessable Entity",
			422,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	j, _ := json.Marshal(map[string]string{"msg": "Created user record."})
	w.Write(j)
}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}

	loginUser := dataResource.Data
	dataStore := common.NewDataStore()
	defer dataStore.Close()

	var user model.User
	dataStore.Session.Where("email = ?", loginUser.Email).First(&user)

	// Authenticate the login user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(user.Email)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Eror while generating the access token",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.Password = ""
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
