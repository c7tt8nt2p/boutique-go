package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/checkout/service/internal/biz"
)

type checkoutRepo struct {
	data *Data
	log  *log.Helper
}

func NewCheckoutRepo(data *Data, logger log.Logger) biz.CheckoutRepo {
	return &checkoutRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/checkout")),
	}
}
