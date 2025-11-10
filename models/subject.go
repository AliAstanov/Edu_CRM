package models

import (
	"github.com/google/uuid"
)

type Subject struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateSubject struct {
	Name string `json:"name" binding:"required"`
}

type GetSubjects struct {
	Subjects []Subject `json:"subjects"`
	Count    int       `json:"count"`
}

type UpdateSubjectReq struct{
	Name string `json:"name"`
}