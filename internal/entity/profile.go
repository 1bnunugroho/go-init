package entity

import (
	"time"
)

// User represents a profile.
type Profile struct {
	ID   		string `json:"id"`
	Bio  		string `json:"password"`
	Image 		string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetUserName returns the user name.
/*
func (u User) GetBio() string {
	return "null"
}

// GetUserName returns the user name.
func (u User) GetImage() string {
	return "null"
}
*/