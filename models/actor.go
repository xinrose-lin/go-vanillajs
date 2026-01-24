package models

type Actor struct {
	ID int
	FirstName string 
	LastName string 
	// optional: allow if actor has no picture
	// * (pointer) included so that its nullable 
	ImageURL *string 
}

