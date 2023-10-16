package biz

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/kx-boutique/api/product/service/v1"
	entModel "github.com/kx-boutique/app/product/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type ProductRepo interface {
	GetEntClient() *ent.Client
	Save(ctx context.Context, client *ent.Client, pp *entModel.Product) *entModel.Product
	FindById(ctx context.Context, client *ent.Client, id uuid.UUID) *entModel.Product
	UpdateStockById(ctx context.Context, client *ent.Client, id uuid.UUID, toStock int32) *entModel.Product
	SubtractStockById(ctx context.Context, client *ent.Client, id uuid.UUID, toStock int32) *entModel.Product
	IsPurchasable(ctx context.Context, client *ent.Client, id uuid.UUID, qty int32) bool
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
	return uc.repo.Save(ctx, uc.repo.GetEntClient(), pe)
}

func (uc *ProductUseCase) GetProductById(ctx context.Context, req *pb.GetProductReq) *entModel.Product {
	id := util.ParseUUID(req.Id)

	return uc.repo.FindById(ctx, uc.repo.GetEntClient(), id)
}

func (uc *ProductUseCase) UpdateProductStock(ctx context.Context, req *pb.UpdateProductStockReq) *entModel.Product {
	id := util.ParseUUID(req.Id)
	return uc.repo.UpdateStockById(ctx, uc.repo.GetEntClient(), id, req.Stock)
}

func (uc *ProductUseCase) SubtractProductStock(ctx context.Context, req *pb.SubtractProductStockReq) *entModel.Product {
	id := util.ParseUUID(req.Id)
	return uc.repo.SubtractStockById(ctx, uc.repo.GetEntClient(), id, req.SubtractStock)
}

func (uc *ProductUseCase) ValidatePurchasableProduct(ctx context.Context, req *pb.ValidatePurchasableProductReq) bool {
	id := util.ParseUUID(req.Id)
	ok := uc.repo.IsPurchasable(ctx, uc.repo.GetEntClient(), id, req.Qty)
	return ok
}
