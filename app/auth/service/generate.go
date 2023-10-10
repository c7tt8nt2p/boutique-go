package generate

//go:generate kratos proto server ../../../api/auth/service/v1/auth.proto -t ./internal/service

//go:generate make api

//go:generate make proto

//go:generate make wire

//go:generate make build
