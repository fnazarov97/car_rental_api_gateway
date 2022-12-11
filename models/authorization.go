package models

import "time"

// LoginModel ...
type LoginModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse ...
type TokenResponse struct {
	Token string `json:"token"`
}

type CreateUserModel struct {
	Id       string `json:"id" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"user_type" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type User struct {
	Id        string     `json:"id" binding:"required"`
	Fname     string     `json:"fname" binding:"required"`
	Lname     string     `json:"lname" binding:"required"`
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	UserType  string     `json:"user_type" binding:"required"`
	Address   string     `json:"address" binding:"required"`
	Phone     string     `json:"phone" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type UpdateUserModel struct {
	Id       string `json:"id" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"user_type" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}
