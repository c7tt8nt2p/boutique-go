package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
