package model

import "github.com/google/uuid"

type Auth struct {
	Id           uuid.UUID
	PasswordHash string
	UserId       uuid.UUID
}
