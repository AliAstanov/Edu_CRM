package models

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID        uuid.UUID `json:"id"`
	Name  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateGroup struct {
	Name string `json:"full_name" binding:"required"`
}

type GetGroups struct {
	Groups []Group `json:"groups"`
	Count    int       `json:"count"`
}

type UpdateGroupReq struct {
	FullName string `json:"name"`
}

