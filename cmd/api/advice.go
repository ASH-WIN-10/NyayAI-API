package main

import "net/http"

func (app *application) getLegalAdvice(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Prompt string  `json:"prompt"`
		Region *string `json:"region,omitempty"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Region == nil {
		input.Region = new(string)
		*input.Region = "India" // Default region "India"
	}
	input.Prompt = input.Prompt + "\n\nProvide legal advice relevant to " + *input.Region + "."

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
