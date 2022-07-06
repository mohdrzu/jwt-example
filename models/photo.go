package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	UserID uint
	Url    string
}

type CreatePhoto struct {
	Url string `json:"url"`
}

type UserPhotoResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo    []Photo
}
