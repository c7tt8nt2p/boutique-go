package generate

//go:generate kratos proto server ../../../api/user/service/v1/cart.proto -t ./internal/service

//go:generate make api

//go:generate make proto

//go:generate make wire

//go:generate make build
