package entity

import (
	"time"
)

// User represents a user.
type UsersProfile struct {
	username string    `json:"username"`
	name     string    `json:"name"`
	address  string    `json:"address"`
	bod      time.Time `json:"bod"`
	email    string    `json:"email"`
}
