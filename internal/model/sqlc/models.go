// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"
)

type Secret struct {
	ID           int64     `json:"id"`
	Content      string    `json:"content"`
	Creator      string    `json:"creator"`
	Token        string    `json:"token"`
	Hashpassword string    `json:"hashpassword"`
	Isview       bool      `json:"isview"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Hashpassword string    `json:"hashpassword"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Isactive     bool      `json:"isactive"`
	CreatedAt    time.Time `json:"created_at"`
}
