package entity

type Student struct {
	ID      int64  `json:"id" db:"id"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
	Picture string `json:"picture" db:"picture"`
}

type EnrollWebinarParam struct {
	WebinarID int64 `json:"webinar_id" db:"webinar_id"`
	TeacherID int64 `json:"teacher_id" db:"teacher_id"`
}

type CancelEnrollmentWebinarParam struct {
	WebinarID int64 `json:"webinar_id" db:"webinar_id"`
	TeacherID int64 `json:"teacher_id" db:"teacher_id"`
}
