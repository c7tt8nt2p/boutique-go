package biz

import (
	"context"
	"github.com/kx-boutique/app/product/service/internal/data"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductUseCase struct {
	repo data.ProductRepo
	log  *log.Helper
}

func NewProductUseCase(repo data.ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/product"))}
}

func (uc *ProductUseCase) GetProductById(ctx context.Context, id string) (*data.ProductEntity, error) {
	return uc.repo.GetProductById(ctx, id)
}
