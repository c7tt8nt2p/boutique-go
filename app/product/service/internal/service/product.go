package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/product/service/internal/biz"

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

func (s *ProductService) ViewProduct(ctx context.Context, req *pb.ViewProductReq) (*pb.ViewProductResp, error) {
	product, err := s.uc.GetProductById(ctx, req.GetId())
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
