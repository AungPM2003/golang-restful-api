package main

import (
	"net/http"
)

// this func will check and show the app health
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map as go object that we want to send in http.Response with json format
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
