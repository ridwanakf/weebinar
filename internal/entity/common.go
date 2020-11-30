package entity

type GoogleAuth struct {
	ID      string `json:"id" db:"id"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
	Picture string `json:"picture" db:"picture"`
}
