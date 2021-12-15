package models

import (
	"time"
)

type UserRole string

// User statuses enum.
const (
	UserRoleCustomer UserRole = "customer"
	UserRoleOwner    UserRole = "owner"
)

type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	Language  string   `json:"language"`
	Role      UserRole `json:"role"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
