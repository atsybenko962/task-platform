package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt int64     `json:"created_at"`
	createdAt time.Time `json:"-"`
}

func (u *User) Scan() error {
	u.CreatedAt = u.createdAt.Unix()
	return nil
}

type UserWithPassword struct {
	User
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"-"`
}

func (u *UserWithPassword) Scan() error {
	return u.User.Scan()
}

type RefreshToken struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	TokenHash string    `json:"-"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
