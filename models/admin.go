package models

import "time"

type Admin struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // never return password
	CreatedAt time.Time `json:"created_at"`
}

type RegisterAdmin struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
