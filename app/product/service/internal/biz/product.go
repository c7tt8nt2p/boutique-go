package biz

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/product/service/internal/data"
	"github.com/kx-boutique/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductUseCase struct {
	repo data.ProductRepo
	log  *log.Helper
}

func NewProductUseCase(repo data.ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/product"))}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*data.ProductEntity, error) {
	pe := &data.ProductEntity{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		UnitPrice:   req.UnitPrice,
	}
	return uc.repo.SaveProduct(ctx, pe)
}

func (uc *ProductUseCase) GetProductById(ctx context.Context, req *pb.ViewProductReq) (*data.ProductEntity, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, errors.ErrValidationFailed("Id is not valid.")
	}

	return uc.repo.GetProductById(ctx, id)
}
