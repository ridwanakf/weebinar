package entity

import "time"

type Webinar struct {
	ID           int64     `json:"" db:""`
	TeacherID    int64     `json:"" db:""`
	Title        string    `json:"" db:""`
	Desc         string    `json:"" db:""`
	Participants []Student `json:"" db:""`
	Schedule     time.Time `json:"" db:""`
	CreatedAt    time.Time `json:"" db:""`
}

type ListParticipants struct {
	StudentID int64 `json:"student_id" db:"student_id"`
	WebinarID int64 `json:"webinar_id" db:"webinar_id"`
	Status    int32 `json:"status" db:"status"`
}

type CreateWebinarParam struct {
}

type UpdateWebinarParam struct {
}

type DeleteWebinarParam struct {
}
