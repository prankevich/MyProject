package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	FullName  string    `json:"full_name" db:"full_name"`
	Username  string    `json:"user_name" db:"user_name"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"create_at"`
	UpdatedAt time.Time `json:"updated_at" db:"update_at"`
	Role      Role      `json:"role" db:"role"`
}

type Role string

const (
	RoleUser  = "USER"
	RoleAdmin = "ADMIN"
)
