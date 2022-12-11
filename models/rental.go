package models

import "time"

//CreateRentalModel ...
type CreateRentalModel struct {
	RentalId   string `json:"rental_id" binding:"required"`
	CarId      string `json:"car_id" binding:"required"`
	CustomerId string `json:"customer_id" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
	Payment    string `json:"payment" binding:"required"`
}

// Rental ...
type Rental struct {
	RentalId   string     `json:"rental_id" binding:"required"`
	CarId      string     `json:"car_id" binding:"required"`
	CustomerId string     `json:"customer_id" binding:"required"`
	StartDate  string     `json:"start_date" binding:"required"`
	EndDate    string     `json:"end_date" binding:"required"`
	Payment    string     `json:"payment" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

// PackedRentalModel ...
type PackedRentalModel struct {
	RentalId  string     `json:"rental_id" binding:"required"`
	Car       Car        `json:"car"`
	Customer  User       `json:"user"`
	StartDate string     `json:"start_date" binding:"required"`
	EndDate   string     `json:"end_date" binding:"required"`
	Payment   string     `json:"payment" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type UpdateRentalModel struct {
	RentalId   string `json:"rental_id" binding:"required"`
	CarId      string `json:"car_id" binding:"required"`
	CustomerId string `json:"customer_id" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
	Payment    string `json:"payment" binding:"required"`
}
