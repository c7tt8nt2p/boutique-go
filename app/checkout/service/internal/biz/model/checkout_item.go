package model

import (
	"github.com/google/uuid"
)

type CheckoutItem struct {
	Id         uuid.UUID
	CheckoutId uuid.UUID
	CartItemId uuid.UUID
}
