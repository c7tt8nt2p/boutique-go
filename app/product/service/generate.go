package generate

//go:generate kratos proto server ../../../api/product/service/v1/product.proto -t ./internal/service

//go:generate make api

//go:generate make proto

//go:generate make wire

//go:generate make ent
