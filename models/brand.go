package models

import "time"

// CreateBrandModel ...
type CreateBrandModel struct {
	Name        string `json:"name" binding:"required" minLength:"2" maxLength:"255" example:"John"`
	Discription string `json:"discription" example:"bu brand uzAuto tomonidan yaratilgan"`
	Year        string `json:"year" example:"2005"`
}

// Brand ...
type Brand struct {
	Id          string     `json:"brand_id"`
	Name        string     `json:"name" binding:"required" minLength:"2" maxLength:"255" example:"John"`
	Discription string     `json:"discription" example:"bu brand uzAuto tomonidan yaratilgan"`
	Year        string     `json:"year" example:"2005"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}

// UpdateBrandModel ...
type UpdateBrandModel struct {
	BrandId     string `json:"brand_id"`
	Name        string `json:"name" binding:"required" minLength:"2" maxLength:"255" example:"John"`
	Discription string `json:"discription" example:"bu brand uzAuto tomonidan yaratilgan"`
	Year        string `json:"year" example:"2005"`
}
