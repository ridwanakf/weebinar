package entity

import (
	"time"
)

type Webinar struct {
	ID           int64     `json:"id" db:"id"`
	TeacherID    int64     `json:"teacher_id" db:"teacher_id"`
	Title        string    `json:"title" db:"title"`
	Desc         string    `json:"description" db:"description"`
	Link         string    `json:"link" db:"link"`
	Category     string    `json:"category" db:"category"`
	Schedule     string    `json:"schedule" db:"schedule"`
	Participants []Student `json:"participants"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type ListParticipants struct {
	StudentID int64 `json:"student_id" db:"student_id"`
	WebinarID int64 `json:"webinar_id" db:"webinar_id"`
	Status    int32 `json:"status" db:"status"`
}

type CreateWebinarParam struct {
	TeacherID int64     `json:"teacher_id" db:"teacher_id"`
	Title     string    `json:"title" db:"title"`
	Desc      string    `json:"description" db:"description"`
	Link      string    `json:"link" db:"link"`
	Category  string    `json:"category" db:"category"`
	Schedule  string    `json:"schedule" db:"schedule"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UpdateWebinarParam struct {
	ID       int64  `db:"id"`
	Title    string `json:"title" db:"title"`
	Desc     string `json:"description" db:"description"`
	Link     string `json:"link" db:"link"`
	Category string `json:"category" db:"category"`
	Schedule string `json:"schedule" db:"schedule"`
}

type DeleteWebinarParam struct {
	ID int64 `json:"id" db:"id"`
}

func (c CreateWebinarParam) Validate() error {
	return nil
}

func (u UpdateWebinarParam) Validate() error {
	return nil
}

func (d DeleteWebinarParam) Validate() error {
	return nil
}