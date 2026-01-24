package models

type Movie struct {
	ID int 
	// api 
	TMDB_ID int
	Title string 
	Tagline string 
	ReleaseYear int 
	Genres []Genre 
	Overview *string
	Score *float32
	Popularity *float32
	Keywords []string
	Language *string

	PosterURL *string
	TrailerURL *string
	Casting []Actor
//  Unlike primitive types, which hold default values like 0 or "" when empty, 
// pointers can be nil,
}

// temp := Movie()
// in database, if there is "null"
//when read into model, it will be "nil"
