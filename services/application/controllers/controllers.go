package controllers

import (
	"encoding/json"
	"errors"
	//"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/beautiful-code/sal/common/messages"
	"github.com/beautiful-code/sal/common/utils"

	"github.com/beautiful-code/sal/services/application/app"
	"github.com/beautiful-code/sal/services/application/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// Get the current user object
	userMessage := getUserMessage(r)

	// Decode the incoming Application json
	var applicationMessage messages.ApplicationMessage
	err := json.NewDecoder(r.Body).Decode(&applicationMessage)
	if err != nil {
		utils.DisplayAppError(w, err, "Invalid Application JSON data", 500)
		return
	}

	application := model.Application{
		Name:   applicationMessage.Data.Name,
		UserId: userMessage.Data.ID,
	}

	valid, err := govalidator.ValidateStruct(application)

	if valid {
		if app.Data.DB.NewRecord(&application) {
			result := app.Data.DB.Create(&application)
			if result.Error != nil {
				err := "Failed to write an application record to the database."
				utils.DisplayAppError(w, errors.New(err), err, 500)
				return
			}
		}
	} else {
		utils.DisplayAppError(w, err, "Unprocessable Entity", 422)
		return
	}

	utils.DisplayAppOK(w, "Created application record.", http.StatusCreated)

}

func getUserMessage(r *http.Request) messages.UserMessage {
	var userMessage messages.UserMessage
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/user", nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	res, _ := client.Do(req)

	//block forever at the next line
	content, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(content, &userMessage)
	return userMessage
}
