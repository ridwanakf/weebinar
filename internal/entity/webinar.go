package entity

import (
	"time"
)

type Webinar struct {
	ID             int64          `json:"id" db:"id" form:"id"`
	TeacherID      int64          `json:"teacher_id" db:"teacher_id" form:"teacher_id"`
	Title          string         `json:"title" db:"title" form:"title"`
	Desc           string         `json:"description" db:"description" form:"description"`
	Link           string         `json:"link" db:"link" form:"link"`
	Category       string         `json:"category" db:"category" form:"category"`
	ScheduleString string         `json:"schedule_string" db:"schedule_string" form:"schedule_string"`
	Schedule       time.Time      `json:"schedule" db:"schedule" form:"schedule"`
	Participants   []Participants `json:"participants"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
}

type Participants struct {
	Student `json:"student"`
	Status  int32 `json:"status" db:"status"`
}

type CreateWebinarParam struct {
	TeacherID      int64     `json:"teacher_id" db:"teacher_id" form:"teacher_id"`
	Title          string    `json:"title" db:"title" form:"title"`
	Desc           string    `json:"description" db:"description" form:"description"`
	Link           string    `json:"link" db:"link" form:"link"`
	Category       string    `json:"category" db:"category" form:"category"`
	ScheduleString string    `json:"schedule_string" db:"schedule_string" form:"schedule_string"`
	Schedule       time.Time `json:"schedule" db:"schedule" form:"schedule"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

type UpdateWebinarParam struct {
	ID             int64     `json:"id" db:"id" form:"id"`
	Title          string    `json:"title" db:"title" form:"title"`
	Desc           string    `json:"description" db:"description" form:"description"`
	Link           string    `json:"link" db:"link" form:"link"`
	Category       string    `json:"category" db:"category" form:"category"`
	ScheduleString string    `json:"schedule_string" db:"schedule_string" form:"schedule_string"`
	Schedule       time.Time `json:"schedule" db:"schedule" form:"schedule"`
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
