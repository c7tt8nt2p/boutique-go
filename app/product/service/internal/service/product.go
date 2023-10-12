package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/product/service/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	p := s.uc.CreateProduct(ctx, req)

	return &pb.CreateProductResp{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
		UnitPrice:   p.UnitPrice,
	}, nil
}

func (s *ProductService) GetProductById(ctx context.Context, req *pb.GetProductByIdReq) (*pb.GetProductByIdResp, error) {
	p := s.uc.GetProductById(ctx, req)

	return &pb.GetProductByIdResp{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
		UnitPrice:   p.UnitPrice,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}, nil
}

func (s *ProductService) ValidatePurchasableProduct(ctx context.Context, req *pb.ValidatePurchasableProductReq) (*pb.ValidatePurchasableProductResp, error) {
	ok := s.uc.ValidatePurchasableProduct(ctx, req)
	if ok {
		return &pb.ValidatePurchasableProductResp{
			Validated: true,
		}, nil
	} else {
		return &pb.ValidatePurchasableProductResp{
			Validated:         false,
			ValidationMessage: "product is not purchasable due to invalid product or quantity",
		}, nil
	}

}
