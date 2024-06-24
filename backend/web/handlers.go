package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"web-page/backend/pet"
)

func (app *application) Index(
	w http.ResponseWriter,
	r *http.Request) {
	app.render(w, "index.page.gohtml", nil)
}

func (app *application) ShowPage(
	w http.ResponseWriter,
	r *http.Request) {
	pageName := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", pageName), nil)
}

func (app *application) CreateDogFactory(w http.ResponseWriter, r *http.Request) {
	// Parse the species from the request body (assuming JSON input)
	var input struct {
		Species string `json:"species"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new Pet using the factory function
	petObj := pet.NewPet(input.Species)

	// Respond with the newly created Pet in JSON format
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(petObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) CreateCatFactory(w http.ResponseWriter, r *http.Request) {
	// Parse the species from the request body (assuming JSON input)
	var input struct {
		Species string `json:"species"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new Pet using the factory function
	petObj := pet.NewPet(input.Species)

	// Respond with the newly created Pet in JSON format
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(petObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
