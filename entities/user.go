package entities

type User struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name" validate:"required"`
}
