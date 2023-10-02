package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/cart/service/internal/biz"

	pb "github.com/kx-boutique/api/cart/service/v1"
)

type CartService struct {
	pb.UnimplementedCartServer

	cc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(cc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}

func (s *CartService) ViewCart(ctx context.Context, req *pb.ViewCartReq) (reply *pb.ViewCartResp, err error) {
	reply = &pb.ViewCartResp{Items: make([]*pb.ViewCartResp_Item, 0)}
	c, err := s.cc.GetCart(ctx, req.UserId)
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&pb.ViewCartResp_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, nil
}
