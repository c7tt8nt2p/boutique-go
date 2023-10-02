package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/cart/service/internal/biz"

	v1 "github.com/kx-boutique/api/cart/service/v1"
)

type CartService struct {
	v1.UnimplementedCartServer

	cc  *biz.CartUseCase
	log *log.Helper
}

func NewCartService(cc *biz.CartUseCase, logger log.Logger) *CartService {
	return &CartService{
		cc:  cc,
		log: log.NewHelper(log.With(logger, "module", "service/cart"))}
}

func (s *CartService) ViewCart(ctx context.Context, req *v1.ViewCartReq) (reply *v1.ViewCartResp, err error) {
	reply = &v1.ViewCartResp{Items: make([]*v1.ViewCartResp_Item, 0)}
	c, err := s.cc.GetCart(ctx, req.UserId)
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&v1.ViewCartResp_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, nil
}
