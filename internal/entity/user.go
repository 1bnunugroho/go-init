package entity

import (
	"time"
)

// User represents a user.
type User struct {
	ID   		string
	Email 		string
	Username 	string
	Password 	string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c User) TableName() string {
	return "uzer"
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u User) GetEmail() string {
	return u.Email
}

// GetUserName returns the user name.
func (u User) GetUserName() string {
	return u.Username
}