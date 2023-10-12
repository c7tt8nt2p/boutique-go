package model

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Stock       int32
	UnitPrice   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
