package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/kx-boutique/app/product/service/internal/biz"
	"github.com/kx-boutique/pkg/errors"

	pb "github.com/kx-boutique/api/product/service/v1"
)

type ProductService struct {
	pb.UnimplementedProductServer

	uc  *biz.ProductUseCase
	log *log.Helper
}

func NewProductService(uc *biz.ProductUseCase, logger log.Logger) *ProductService {
	return &ProductService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/product")),
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	product, err := s.uc.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResp{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
	}, nil
}

func (s *ProductService) ViewProduct(ctx context.Context, req *pb.ViewProductReq) (*pb.ViewProductResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.ErrValidationFailed("Id is invalid.")
	}

	product, err := s.uc.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.ViewProductResp{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Stock:       product.Stock,
	}, nil
}
