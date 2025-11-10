package models

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	StudentId   uuid.UUID `json:"student_id"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	GroupID     string    `json:"group_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetStudens struct {
	Students []Student `json:"students"`
	Count    int       `json:"count"`
}

type CreateStudent struct {
	FullName string `json:"full_name" binding:"required"` //binding:"required" — bu gin validatori uchun. POST/PUT so‘rovda FullName va Phone bo‘sh bo‘lsa, xato beradi
	PhoneNumber    string `json:"phone_number" binding:"required"`
	GroupID  string `json:"group_id"`
}

type UpdateStudentReq struct {
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	GroupID     string    `json:"group_id"`
}