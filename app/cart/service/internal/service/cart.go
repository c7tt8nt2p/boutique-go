package service

import (
	"context"

	v1 "github.com/go-kratos/beer-shop/api/cart/service/v1"
)

func (s *CartService) ViewCart(ctx context.Context, req *v1.ViewCartReq) (reply *v1.ViewCartReply, err error) {
	reply = &v1.ViewCartReply{Items: make([]*v1.ViewCartReply_Item, 0)}
	c, err := s.cc.GetCart(ctx, req.UserId)
	if err != nil {
		//fixme convert to error msg
		return reply, err
	}
	for _, x := range c.Items {
		reply.Items = append(reply.Items,
			&v1.ViewCartReply_Item{
				ItemId:   x.Id,
				Quantity: x.Quantity,
			})
	}
	return reply, nil
}
