package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ehilmidag/hillmoviego/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "+v\n", input)

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   105,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

//json.Unmarshal() requires about 80% more memory (B/op) than json.Decoder, as well as
//being a tiny bit slower (ns/op).
// example usage for json.Unmarshal instead of decode
// func (app *application) exampleHandler(w(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 	Foo string `json:"foo"`
// 	}
// 	// Use io.ReadAll() to read the entire request body into a []byte slice.
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 	app.serverErrorResponse(w, r, err)
// 	return
// 	}
// 	// Use the json.Unmarshal() function to decode the JSON in the []byte slice to the
// 	// input struct. Again, notice that we are using a *pointer* to the input
// 	// struct as the decode destination.
// 	err = json.Unmarshal(body, &input)
// 	if err != nil {
// 	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
// 	return
// 	}
// 	fmt.Fprintf(w, "%+v\n", input)
// 	}
