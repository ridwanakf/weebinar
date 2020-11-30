package entity

type Student struct {
	ID      int64  `json:"id" db:"id"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
	Picture string `json:"picture" db:"picture"`
}

type EnrollWebinarParam struct {
}

type CancelEnrollmentWebinarParam struct {
}
