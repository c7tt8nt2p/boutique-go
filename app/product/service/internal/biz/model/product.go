package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Stock       int32
	UnitPrice   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
