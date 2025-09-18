package main

import "net/http"

func (app *application) getLegalAdvice(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Prompt string `json:"prompt"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	advice, err := app.ai.GenerateLegalAdvice(input.Prompt)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"advice": advice}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
