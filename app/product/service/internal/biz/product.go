package biz

import (
	"context"
	pb "github.com/kx-boutique/api/product/service/v1"
	"github.com/kx-boutique/app/product/service/internal/data"
	"github.com/kx-boutique/pkg/util"

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

func (uc *ProductUseCase) GetProductById(ctx context.Context, req *pb.GetProductByIdReq) (*data.ProductEntity, error) {
	id, err := util.ParseUUID(req.Id)
	if err != nil {
		return nil, err
	}
	return uc.repo.FindProductById(ctx, id)
}
