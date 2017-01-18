package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	_ "strings"

	"github.com/asaskevich/govalidator"

	"github.com/beautiful-code/sal/common"
	"github.com/beautiful-code/sal/services/app/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	dataStore := common.NewDataStore()
	defer dataStore.Close()

	application := model.Application{
		Name:   "My App",
		UserId: 1,
	}

	valid, err := govalidator.ValidateStruct(application)

	if valid {
		dataStore.Session.Create(&application)

		err := "Application record not saved"
		if dataStore.Session.NewRecord(&application) {
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
	j, _ := json.Marshal(map[string]string{"msg": "Created application record."})
	w.Write(j)

}
