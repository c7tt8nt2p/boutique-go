package biz

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/product/service/v1"
	entModel "github.com/kx-boutique/ent/model"
	"github.com/kx-boutique/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductRepo interface {
	Save(ctx context.Context, pp *entModel.Product) *entModel.Product
	FindById(ctx context.Context, id uuid.UUID) *entModel.Product
	IsPurchasable(ctx context.Context, id uuid.UUID, qty int32) bool
}

type ProductUseCase struct {
	repo ProductRepo
	log  *log.Helper
}

func NewProductUseCase(repo ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/product"))}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, req *pb.CreateProductReq) *entModel.Product {
	pe := &entModel.Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		UnitPrice:   req.UnitPrice,
	}
	return uc.repo.Save(ctx, pe)
}

func (uc *ProductUseCase) GetProductById(ctx context.Context, req *pb.GetProductByIdReq) *entModel.Product {
	id := util.ParseUUID(req.Id)

	return uc.repo.FindById(ctx, id)
}

func (uc *ProductUseCase) ValidatePurchasableProduct(ctx context.Context, req *pb.ValidatePurchasableProductReq) bool {
	id := util.ParseUUID(req.Id)
	ok := uc.repo.IsPurchasable(ctx, id, req.Qty)
	return ok
}
