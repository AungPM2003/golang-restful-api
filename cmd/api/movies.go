package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"aung.greenlight.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	//create a struct for what we expect the input will be
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// this is decoding with json.New Decoder
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//this is decoding with json.Un,marshal
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// err = json.Unmarshal(body, &input)
	// if err != nil {
	// 	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	// 	return
	// }

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "asdfadf",
		Runtime:   102,
		Genres:    []string{"drama", "horror"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
