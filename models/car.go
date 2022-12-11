package models

import "time"

// Car ...
type Car struct {
	CarId     string     `json:"car_id"`
	Model     string     `json:"model" example:"cobolt"`
	Color     string     `json:"color" example:"oq"`
	Year      string     `json:"year" example:"2018"`
	Mileage   string     `json:"mileage" example:"800 km"`
	BrandId   string     `json:"brand_id" example:"1"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateCarModel ...
type CreateCarModel struct {
	Model   string `json:"model" example:"cobolt"`
	Color   string `json:"color" example:"oq"`
	Year    string `json:"year" example:"2018"`
	Mileage string `json:"mileage" example:"800 km"`
	BrandId string `json:"brand_id" example:"1"`
}

// PackedCarModel ...
type PackedCarModel struct {
	CarId     string     `json:"car_id"`
	Model     string     `json:"model" example:"cobolt"`
	Color     string     `json:"color" example:"oq"`
	Year      string     `json:"year" example:"2018"`
	Mileage   string     `json:"mileage" example:"800 km"`
	Brand     Brand      `json:"brand"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type UpdateCarModel struct {
	CarId   string `json:"car_id"`
	Model   string `json:"model" example:"cobolt"`
	Color   string `json:"color" example:"oq"`
	Year    string `json:"year" example:"2018"`
	Mileage string `json:"mileage" example:"800 km"`
	BrandId string `json:"brand_id" example:"1"`
}
