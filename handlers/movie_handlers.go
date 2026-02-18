package handlers

import (
	"encoding/json"
	"net/http"

	// based on how you named your package in go.mod
	"frontendmasters.com/movies/data"
	"frontendmasters.com/movies/logger"
	// "frontendmasters.com/movies/models"
)

// defining structrue of movie handler
// enables abstraction
type MovieHandler struct {
	// logger
	// by capitalising first char, it exports it?
	Storage data.MovieStorage
	Logger  *logger.Logger
}

// interface {} refers to a type that takes in any value
func (h *MovieHandler) writeJsonResponse(w http.ResponseWriter, data interface{}) error {
	// write header: helps browser know we are sending json
	w.Header().Set("Content-Type", "application/json")
	// encode data into json,
	// writes to writer
	// then send to output
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// TODO: log error
		h.Logger.Error("Failed to encode response", err)
		// sends error_msg to writer (w)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return err
	}
	// go if conditions can have multiple expressions
	// so long as the last expression is boolean
	return nil
}

// objective:
// localhost/home/movies/top
// then i want to see json of top 10 movies

// every http handler has the same signature
// http response writer, http request (incoming)

// there is no classes in go; achieves similar functinality
// by using methods defined on types (structs)

// func (receiver_name ReceiverType) MethodName(parameters) return type {
// }
// receiver_name = this / self
// recevierType = struct that the method is bound to
func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		h.Logger.Error("Failed to get top movies", err)
	}
	//  some dummy data
	// movies := []models.Movie {
	// 	{
	// 		ID:          1,
	// 		TMDB_ID:     101,
	// 		Title:       "The Hacker",
	// 		ReleaseYear: 2022,
	// 		Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
	// 		Keywords:    []string{"hacking", "cybercrime"},
	// 		Casting:     []models.Actor{{ID: 1, FirstName: "Jane Doe"}},
	// 	},
	// 	{
	// 		ID:          2,
	// 		TMDB_ID:     102,
	// 		Title:       "Space Dreams",
	// 		ReleaseYear: 2020,
	// 		Genres:      []models.Genre{{ID: 2, Name: "Sci-Fi"}},
	// 		Keywords:    []string{"space", "exploration"},
	// 		Casting:     []models.Actor{{ID: 2, FirstName: "John Star"}},
	// 	},
	// 	{
	// 		ID:          3,
	// 		TMDB_ID:     103,
	// 		Title:       "The Lost City",
	// 		ReleaseYear: 2019,
	// 		Genres:      []models.Genre{{ID: 3, Name: "Adventure"}},
	// 		Keywords:    []string{"jungle", "treasure"},
	// 		Casting:     []models.Actor{{ID: 3, FirstName: "Lara Hunt"}},
	// 	},
	// }
	// to send data
	h.writeJsonResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		h.Logger.Error("Failed to get top movies", err)
	}
	h.writeJsonResponse(w, movies)
}

// func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
// 	//  some dummy data
// 	movie := []models.Movie{
// 		{
// 			ID:          1,
// 			TMDB_ID:     101,
// 			Title:       "The Hacker",
// 			ReleaseYear: 2022,
// 			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
// 			Keywords:    []string{"hacking", "cybercrime"},
// 			Casting:     []models.Actor{{ID: 1, FirstName: "Jane Doe"}},
// 		},
// 	}
// 	// to send data
// 	h.writeJsonResponse(w, movie)

// }
