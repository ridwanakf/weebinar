package entity

type Student struct {
	ID      string `json:"id" db:"id"`
	Email   string `json:"email" db:"email"`
	Name    string `json:"name" db:"name"`
	Picture string `json:"picture" db:"picture"`
}

type EnrollWebinarParam struct {
	WebinarID int64  `json:"webinar_id" db:"webinar_id"`
	TeacherID string `json:"teacher_id" db:"teacher_id"`
}

type CancelEnrollmentWebinarParam struct {
	WebinarID int64  `json:"webinar_id" db:"webinar_id"`
	TeacherID string `json:"teacher_id" db:"teacher_id"`
}

func (e EnrollWebinarParam) Validate() error {
	return nil
}
func (c CancelEnrollmentWebinarParam) Validate() error {
	return nil
}
