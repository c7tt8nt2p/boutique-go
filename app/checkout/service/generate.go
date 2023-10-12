package generate

//go:generate kratos proto server ../../../api/checkout/service/v1/checkout.proto -t ./internal/service

//go:generate make api

//go:generate make proto

//go:generate make wire

//go:generate make build
