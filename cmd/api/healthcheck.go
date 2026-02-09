package main

import (
	"net/http"
)

// this func will check and show the app health
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map as go object that we want to send in http.Response with json format
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
