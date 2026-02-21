package data

import (
	"database/sql"
	"errors"
	// "strconv"

	"frontendmasters.com/movies/logger"
	"frontendmasters.com/movies/models"
	_ "github.com/lib/pq"
)

type MovieRepository struct {
	db *sql.DB
	logger *logger.Logger
}

// Factory
// may need multiple connections to a db or to diff dbs
// Factory class ensures consistent initialisation befr use
// allow you to swap in alt implementation 
func NewMovieRepository(db *sql.DB, log *logger.Logger) (*MovieRepository, error) {
	return &MovieRepository{
		db: 	db, 
		logger: log,
	}, nil
}

const defaultLimit = 20

// 
func (r *MovieRepository) GetTopMovies() ([]models.Movie, error) {
	// query := `
	// 	SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
	// 	       popularity, language, poster_url, trailer_url
	// 	FROM movies
	// 	ORDER BY popularity DESC
	// 	LIMIT` + quantity
	//
	//	^ concatenating var as a str -- risks SQL injection attack

	// INSTEAD use arguments 'LIMIT $1 OFFSET $2'
	// reads first and second arugment; so that SQL sees this as a arg/value
	// will not see it as SQL

	// Fetch movies
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		ORDER BY popularity DESC
		LIMIT $1
	`
		// LIMIT '\n drop table'
	return r.getMovies(query)
}

func (r *MovieRepository) GetRandomMovies() ([]models.Movie, error) {
	// Fetch movies
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		ORDER BY random() DESC
		LIMIT $1
	`
		// LIMIT '\n drop table'
	return r.getMovies(query)
}

func (r *MovieRepository) getMovies(query string) ([]models.Movie, error) {
	rows, err := r.db.Query(query, defaultLimit)
	if err != nil {
		r.logger.Error("Failed to query movies", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	//reads row by row 
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.TMDB_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
			&m.Overview, &m.Score, &m.Popularity, &m.Language,
			&m.PosterURL, &m.TrailerURL,
		); err != nil {
			r.logger.Error("Failed to scan movie row", err)
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func (r *MovieRepository) GetMovieById(id int) (models.Movie, error) {
	// Fetch movie
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		WHERE id = $1
	`
	row := r.db.QueryRow(query, id)

	var m models.Movie
	err := row.Scan(
		&m.ID, &m.TMDB_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
		&m.Overview, &m.Score, &m.Popularity, &m.Language,
		&m.PosterURL, &m.TrailerURL,
	)
	if err == sql.ErrNoRows {
		r.logger.Error("Movie not found", ErrMovieNotFound)
		return models.Movie{}, ErrMovieNotFound
	}
	if err != nil {
		r.logger.Error("Failed to query movie by ID", err)
		return models.Movie{}, err
	}

	// Fetch related data
	// if err := r.fetchMovieRelations(&m); err != nil {
	// 	return models.Movie{}, err
	// }

	return m, nil
}

var (
	ErrMovieNotFound = errors.New("movie not found")
)