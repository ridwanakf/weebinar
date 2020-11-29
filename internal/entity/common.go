package entity

type User struct {
	ID      int64  `json:"id" db:"id"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
	Role    string `json:"role" db:"role"`
	Picture string `json:"picture" db:"picture"`
}
