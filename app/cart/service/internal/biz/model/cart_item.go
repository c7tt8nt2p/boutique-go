package model

import (
	"github.com/google/uuid"
	"time"
)

type CartItem struct {
	Id         uuid.UUID
	CartId     uuid.UUID
	ProductId  uuid.UUID
	Qty        int32
	CheckedOut bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CartWithProducts struct {
	CartId  uuid.UUID
	Product []*CartProduct
}

type CartProduct struct {
	Id        uuid.UUID
	ProductId uuid.UUID
	Name      string
	Price     float64
	Qty       int32
}
