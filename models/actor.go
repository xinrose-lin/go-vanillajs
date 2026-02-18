package models

type Actor struct {
	ID int				`json:"id"`
	FirstName string 	`json:"first_name"`
	LastName string 	`json:"last_name"`
	// optional: allow if actor has no picture
	// * (pointer) included so that its nullable 
	ImageURL *string 	`json:"image_url"`
}

