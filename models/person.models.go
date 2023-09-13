package models


type Person struct {
	ID       int    `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
	Email	string       `json:"email"`
	
}

type CreatePersonModel struct {
	Name string `json:"name" validate:"required"`
	Email	string       `json:"email"`
}
type UpdatePersonModel struct {
	ID       int    `json:"id,omitempty"`
	Name string `json:"name"`
	Email	string       `json:"email"`
}