package controllers

import (
	"html/template"
	"net/http"

	"functions/functions"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	// Initialize a struct to handle potential errors
	d := functions.Data{}

	// Retrieve form values from the HTTP request
	banner := r.PostFormValue("banners")
	input := r.PostFormValue("text")

	// Check if the 'banners' parameter is valid
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		d.ErrNum = http.StatusUnauthorized
		d.ErrTxt = "Unauthorized Action Try Again"
		functions.HandlerError(w, r, &d) // Handle the error and return
		return
	}
	input2 := functions.CheckInput(input)  


	// Call functions to get banner and process input
	getBanner := functions.GetBanner(banner)
	output := functions.ReadInput(input2, getBanner)

	// Parse the HTML template file
	tmpl, err := template.ParseFiles("templates/ascii-art.html")
	if err != nil {
		d.ErrNum = http.StatusInternalServerError
		d.ErrTxt = "Template parsing error"
		functions.HandlerError(w, r, &d) // Handle the error and return
		return
	}

	// Execute the template with the processed output
	err = tmpl.Execute(w, output)
	if err != nil {
		d.ErrNum = http.StatusInternalServerError
		d.ErrTxt = "Template execution error"
		functions.HandlerError(w, r, &d) // Handle the error and return
		return
	}
}
