package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/product/service/internal/biz"
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
	entity, err := s.uc.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResp{
		Id:          entity.Id,
		Name:        entity.Name,
		Description: entity.Description,
		Stock:       entity.Stock,
		UnitPrice:   entity.UnitPrice,
	}, nil
}

func (s *ProductService) GetProductById(ctx context.Context, req *pb.GetProductByIdReq) (*pb.GetProductByIdResp, error) {
	entity, err := s.uc.GetProductById(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.GetProductByIdResp{
		Id:          entity.Id,
		Name:        entity.Name,
		Description: entity.Description,
		Stock:       entity.Stock,
		UnitPrice:   entity.UnitPrice,
	}, nil
}

func (s *ProductService) ValidatePurchasableProduct(ctx context.Context, req *pb.ValidatePurchasableProductReq) (*pb.ValidatePurchasableProductResp, error) {
	err := s.uc.ValidatePurchasableProduct(ctx, req)
	if err != nil {
		return &pb.ValidatePurchasableProductResp{
			Validated:         false,
			ValidationMessage: err.Error(),
		}, nil
	}

	return &pb.ValidatePurchasableProductResp{
		Validated: true,
	}, nil
}
