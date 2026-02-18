package data 

import "frontendmasters.com/movies/models"

// set the signature of what we need
// force the implementation in postgresql for this project

type MovieStorage interface {
	// GetTopMovies() []models.Movie
	// return movies with error is better 
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	// GetMovieById(id int) (models.Movie, error)
	// SearchMoviesByName(name string) ([]models.Movie, error)
	// GetAllGenres() ([]models.Genre, error)

}