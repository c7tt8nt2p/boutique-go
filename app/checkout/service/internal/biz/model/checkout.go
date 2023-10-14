package model

import (
	"github.com/google/uuid"
	"time"
)

type Checkout struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	TotalPrice float64
	CreatedAt  time.Time
}
