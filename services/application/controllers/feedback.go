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

func CreateFeedback(w http.ResponseWriter, r *http.Request) {
	utils.Info.Printf("Request to create a Feedback record.\n")
	var feedbackMessage messages.FeedbackMessage

	err := json.NewDecoder(r.Body).Decode(&feedbackMessage)
	if err != nil {
		utils.DisplayAppError(w, err, "Invalid Application JSON data", 500)
		return
	}

	feedback := models.Feedback{
		Desc:          feedbackMessage.Data.Desc,
		ApplicationId: feedbackMessage.Data.ApplicationId,
		Email:         feedbackMessage.Data.Email,
	}

	valid, err := govalidator.ValidateStruct(feedback)

	if valid {
		if app.Data.DB.NewRecord(&feedback) {
			result := app.Data.DB.Create(&feedback)
			if result.Error != nil {
				err := "Failed to write a feedback record to the database."
				utils.DisplayAppError(w, errors.New(err), err, 500)
				return
			}
		}
	} else {
		utils.DisplayAppError(w, err, "Unprocessable Entity", 422)
		return
	}

	utils.DisplayAppOK(w, "Created feedback record.", http.StatusCreated)

}

func ListFeedbacks(w http.ResponseWriter, r *http.Request) {
	var requestMessage messages.FeedbackListRequestMessage

	var err error
	err = json.NewDecoder(r.Body).Decode(&requestMessage)
	if err != nil {
		utils.DisplayAppError(w, err, "Invalid Application JSON data", 500)
		return
	}

	err = CheckUserOwnsApplication(r, requestMessage.Data.ID)
	if err != nil {
		utils.DisplayAppError(w, err, "Unprocessable Entity", 422)
		return
	} else {
		var feedbacks []models.Feedback
		app.Data.DB.Where("application_id = ?", fmt.Sprintf("%v", requestMessage.Data.ID)).Find(&feedbacks)

		var records []messages.FeedbackRecord = make([]messages.FeedbackRecord, len(feedbacks))
		for i, e := range feedbacks {
			var record messages.FeedbackRecord
			record.ID = e.ID
			record.Desc = e.Desc
			record.Email = e.Email
			record.ApplicationId = e.ApplicationId
			records[i] = record
		}

		j, err := json.Marshal(messages.FeedbackListResponseMessage{Data: records})
		if err != nil {
			utils.DisplayAppError(w, err, "An unexpected error has occurred", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}
