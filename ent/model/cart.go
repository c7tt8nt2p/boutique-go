package model

import "github.com/google/uuid"

type Cart struct {
	Id     uuid.UUID
	UserId uuid.UUID
}
