package models

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTeacher struct {
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type GetTeachers struct {
	Teachers []Teacher `json:"teachers"`
	Count    int       `json:"count"`
}

type UpdateTeacherReq struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

