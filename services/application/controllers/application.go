package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	//"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/beautiful-code/sal/common/messages"
	"github.com/beautiful-code/sal/common/utils"

	"github.com/beautiful-code/sal/services/application/app"
	"github.com/beautiful-code/sal/services/application/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	utils.Info.Printf("Request to create an application.\n")
	// Get the current user object
	userMessage, gum_err := GetUserMessage(r)
	if gum_err != nil {
		utils.DisplayAppError(w, gum_err, "Unable to get the UserMessage object.", 500)
		return
	}

	// Decode the incoming Application json
	var applicationMessage messages.ApplicationMessage
	err := json.NewDecoder(r.Body).Decode(&applicationMessage)
	if err != nil {
		utils.DisplayAppError(w, err, "Invalid Application JSON data", 500)
		return
	}

	application := models.Application{
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

func List(w http.ResponseWriter, r *http.Request) {
	// Get the current user object
	userMessage, gum_err := GetUserMessage(r)
	if gum_err != nil {
		utils.DisplayAppError(w, gum_err, "Unable to get the UserMessage object.", 500)
		return
	}

	var applications []models.Application
	app.Data.DB.Where("user_id = ?", fmt.Sprintf("%v", userMessage.Data.ID)).Find(&applications)

	var records []messages.ApplicationRecord = make([]messages.ApplicationRecord, len(applications))
	for i, e := range applications {
		var record messages.ApplicationRecord
		record.ID = e.ID
		record.Name = e.Name
		records[i] = record
	}

	j, err := json.Marshal(messages.ApplicationListResponseMessage{Data: records})
	if err != nil {
		utils.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}
