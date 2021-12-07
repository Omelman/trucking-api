package models

import (
	"time"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Language  string `json:"language" default:"en-us"`
	RoleID    int    `json:"role_id"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Role *Role `json:"-" bun:"-"`
}

type Role struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Editable bool   `json:"-"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
