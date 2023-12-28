package models

import "time"

type User struct {
	UserId    int       `json:"UserId" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Profile   Profile   `json:"profile" gorm:"embedded"`
}

// User represents a user in the dating app
type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UserId    int       `json:"userId"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResult struct {
	User
}
