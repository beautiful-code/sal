package controllers

import (
	"encoding/json"
	"errors"
	//"fmt"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/common/messages"
	"github.com/beautiful-code/sal/common/utils"

	"github.com/beautiful-code/sal/services/user/app"
	"github.com/beautiful-code/sal/services/user/models"
)

// Register add a new User document
// Handler for HTTP Post - "/register"
func Register(w http.ResponseWriter, r *http.Request) {
	var userMessage messages.UserMessage

	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&userMessage)

	if err != nil {
		utils.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}

	userRecord := userMessage.Data

	hpass, err := bcrypt.GenerateFromPassword([]byte(userRecord.Password), bcrypt.DefaultCost)

	if err != nil {
		panic("bcrypt err!")
	}

	user := model.User{
		FirstName: userRecord.FirstName,
		LastName:  userRecord.LastName,
		Email:     userRecord.Email,
		Password:  string(hpass),
	}

	valid, err := govalidator.ValidateStruct(user)

	if valid {
		result := app.Data.DB.Create(&user)
		err := "User record not saved"

		if result.Error != nil {
			if strings.ContainsAny(result.Error.Error(), "Duplicate entry & for key 'user_email_index'") {
				err = "Email address is already taken!"
			}
		}

		if app.Data.DB.NewRecord(&user) {
			utils.DisplayAppError(w, errors.New(err), "Failed to write to the database.", 500)
			return
		}
	} else {
		utils.DisplayAppError(w, err, "Unprocessable Entity", 422)
		return
	}

	utils.DisplayAppOK(w, "Created user record.", http.StatusCreated)
}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/login"
func Login(w http.ResponseWriter, r *http.Request) {
	var userMessage messages.UserMessage
	var token string
	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&userMessage)
	if err != nil {
		utils.DisplayAppError(w, err, "Invalid Login data", 500)
		return
	}

	loginUser := userMessage.Data

	var user model.User
	app.Data.DB.Where("email = ?", loginUser.Email).First(&user)

	// Authenticate the login user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))

	if err != nil {
		utils.DisplayAppError(w, err, "Invalid login credentials", 401)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(user.Email)
	if err != nil {
		utils.DisplayAppError(w, err, "Eror while generating the access token", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.Password = ""
	authUserMessage := messages.AuthUserMessage{
		Token: token,
	}

	j, err := json.Marshal(authUserMessage)
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// Returns the user object when valid JWT token is present.
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get current user email from the request's context
	user_email := r.Context().Value(common.ContextUserEmailKey).(string)

	var user model.User
	app.Data.DB.Where("email = ?", user_email).First(&user)

	w.Header().Set("Content-Type", "application/json")
	// Clean-up the hashpassword to eliminate it from response JSON
	user.Password = ""
	userRecord := messages.UserRecord{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	j, err := json.Marshal(messages.UserMessage{Data: userRecord})
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
