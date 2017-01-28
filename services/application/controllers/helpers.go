package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beautiful-code/sal/common/messages"
	"github.com/beautiful-code/sal/common/utils"

	"github.com/beautiful-code/sal/services/application/app"
	"github.com/beautiful-code/sal/services/application/models"
)

func GetUserMessage(r *http.Request) (messages.UserMessage, error) {
	var userMessage messages.UserMessage
	client := &http.Client{}
	userEndpoint := fmt.Sprintf("http://%s/user", app.Data.Config.UserService)
	utils.Info.Printf("userEndpoint=%s\n", userEndpoint)

	req, _ := http.NewRequest("GET", userEndpoint, nil)
	req.Header.Set("Authorization", r.Header.Get("Authorization"))

	res, _ := client.Do(req)

	//block forever at the next line
	content, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(content, &userMessage)
	if err != nil {
		return userMessage, err
	}

	// Zero value of uint is 0
	if userMessage.Data.ID == 0 {
		return userMessage, errors.New("Retrieved User ID is 0.")
	}

	return userMessage, nil
}

func CheckUserOwnsApplication(r *http.Request, app_id uint) error {
	userMessage, gum_err := GetUserMessage(r)
	if gum_err != nil {
		return gum_err
	}

	// Get the Application object
	var application models.Application
	app.Data.DB.Where("id = ?", fmt.Sprintf("%v", app_id)).First(&application)

	if application.UserId == userMessage.Data.ID {
		return nil
	} else {
		return errors.New("The application does not belong to the current user.")
	}

}
