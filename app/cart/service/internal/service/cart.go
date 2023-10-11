package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/cart/service/v1"
	"github.com/kx-boutique/app/cart/service/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CartService struct {
	pb.UnimplementedCartServer

	uc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(uc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}

func (s *CartService) NewCart(ctx context.Context, req *pb.NewCartReq) (*pb.NewCartResp, error) {
	id, err := s.uc.NewCart(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.NewCartResp{
		Id: id.String(),
	}, nil
}

func (s *CartService) AddItemToCart(ctx context.Context, req *pb.AddItemToCartReq) (*pb.AddItemToCartResp, error) {
	cie, err := s.uc.AddItemToCart(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.AddItemToCartResp{
		ProductId: cie.ProductId.String(),
		Qty:       cie.Qty,
	}, nil
}

func (s *CartService) ViewCart(ctx context.Context, req *emptypb.Empty) (*pb.ViewCartResp, error) {
	result, err := s.uc.ViewCart(ctx)
	if err != nil {
		return nil, err
	}

	var items []*pb.ViewCartResp_Item

	for _, e := range result {
		items = append(items, &pb.ViewCartResp_Item{
			Id:          e.Id.String(),
			Name:        e.Name,
			Description: e.Description,
		})
	}
	return &pb.ViewCartResp{Items: items}, nil
}
