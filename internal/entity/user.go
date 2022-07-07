package entity

import (
	"time"
)

// User represents a user.
type User struct {
	ID   		string `json:"id"`
	Email 		string `json:"email"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	Bio  		string `json:"password"`
	Image 		string `json:"password"`
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

// GetUserName returns the user name.
func (u User) GetBio() string {
	return "null"
}

// GetUserName returns the user name.
func (u User) GetImage() string {
	return "null"
}