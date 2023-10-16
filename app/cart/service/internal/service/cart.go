package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/kx-boutique/api/cart/service/v1"
	"github.com/kx-boutique/app/cart/service/internal/biz"
	"github.com/shopspring/decimal"
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
	id := s.uc.NewCart(ctx, req)

	return &pb.NewCartResp{
		Id: id.String(),
	}, nil
}

func (s *CartService) AddItemToCart(ctx context.Context, req *pb.AddItemToCartReq) (*pb.AddItemToCartResp, error) {
	ci := s.uc.AddItemToCart(ctx, req)

	return &pb.AddItemToCartResp{
		ProductId: ci.ProductId.String(),
		Qty:       ci.Qty,
	}, nil
}

func (s *CartService) RemoveItemFromCart(ctx context.Context, req *pb.RemoveItemFromCartReq) (*pb.RemoveItemFromCartResp, error) {
	productId := s.uc.RemoveItemFromCart(ctx, req)

	return &pb.RemoveItemFromCartResp{
		ProductId: productId.String(),
	}, nil
}

func (s *CartService) ViewCart(ctx context.Context, _ *emptypb.Empty) (*pb.ViewCartResp, error) {
	var items []*pb.ViewCartResp_Item

	totalPrice := decimal.NewFromFloat(0.0)
	c := s.uc.ViewCart(ctx)
	for _, e := range c.Product {
		totalPriceThisItem := decimal.NewFromFloat(e.Price).Mul(decimal.NewFromInt32(e.Qty))
		totalPrice = totalPrice.Add(totalPriceThisItem)

		items = append(items, &pb.ViewCartResp_Item{
			Id:        e.Id.String(),
			ProductId: e.ProductId.String(),
			Name:      e.Name,
			Price:     e.Price,
			Qty:       e.Qty,
		})
	}
	return &pb.ViewCartResp{
		CartId:     c.CartId.String(),
		TotalPrice: totalPrice.InexactFloat64(),
		Items:      items,
	}, nil
}
