package biz

import (
	"context"
	"fmt"
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
	return uc.repo.Save(ctx, pe)
}

func (uc *ProductUseCase) GetProductById(ctx context.Context, req *pb.GetProductByIdReq) (*data.ProductEntity, error) {
	id, err := util.ParseUUID(req.Id)
	if err != nil {
		return nil, err
	}
	return uc.repo.FindById(ctx, id)
}

func (uc *ProductUseCase) ValidateProduct(ctx context.Context, req *pb.ValidatePurchasableProductReq) (bool, error) {
	id, err1 := util.ParseUUID(req.Id)
	if err1 != nil {
		return false, err1
	}

	ok, err2 := uc.repo.IsPurchasable(ctx, id, req.Qty)
	if err2 != nil {
		return false, err2
	}

	if !ok {
		return false, fmt.Errorf("product is not purchaseable due to invalid product or quantity")
	}

	return uc.repo.IsPurchasable(ctx, id, req.Qty)
}
