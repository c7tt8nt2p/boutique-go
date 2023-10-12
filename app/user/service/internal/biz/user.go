package biz

import (
	"context"
	"github.com/google/uuid"
	authv1 "github.com/kx-boutique/api/auth/service/v1"
	cartv1 "github.com/kx-boutique/api/cart/service/v1"
	pb "github.com/kx-boutique/api/user/service/v1"
	entModel "github.com/kx-boutique/app/user/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/pkg/errors"
	"github.com/kx-boutique/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetEntClient() *ent.Client
	SaveUser(ctx context.Context, client *ent.Client, u *entModel.User) *entModel.User
	FindById(ctx context.Context, client *ent.Client, id uuid.UUID) *entModel.User
	FindIdByEmail(ctx context.Context, client *ent.Client, email string) uuid.UUID
	DeleteById(ctx context.Context, client *ent.Client, id uuid.UUID) uuid.UUID
}

type UserUseCase struct {
	authClient authv1.AuthClient
	cartClient cartv1.CartClient
	repo       UserRepo
	log        *log.Helper
}

func NewUserUseCase(cartClient cartv1.CartClient, authClient authv1.AuthClient, repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		cartClient: cartClient,
		authClient: authClient,
		repo:       repo,
		log:        log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) GetMe(ctx context.Context) *entModel.User {
	me := util.Me(ctx)
	id := util.ParseUUID(me.UserId)
	return uc.repo.FindById(ctx, uc.repo.GetEntClient(), id)
}

func (uc *UserUseCase) RegisterNewUser(ctx context.Context, req *pb.CreateUserReq) *entModel.User {
	user := uc.repo.SaveUser(ctx, uc.repo.GetEntClient(), &entModel.User{
		Name:  req.Name,
		Email: req.Email,
	})

	// new auth
	_, err1 := uc.authClient.Register(ctx, &authv1.RegisterReq{
		UserId:   user.Id.String(),
		Password: req.Password,
	})
	if err1 != nil {
		_ = uc.repo.DeleteById(ctx, uc.repo.GetEntClient(), user.Id)
		panic(errors.AppInternalErr(err1.Error()))
	}

	// new cart
	_, err2 := uc.cartClient.NewCart(ctx, &cartv1.NewCartReq{
		UserId: user.Id.String(),
	})
	if err2 != nil {
		_ = uc.repo.DeleteById(ctx, uc.repo.GetEntClient(), user.Id)
		panic(errors.AppInternalErr(err2.Error()))
	}

	return user
}

func (uc *UserUseCase) GetUserById(ctx context.Context, req *pb.GetUserByIdReq) *entModel.User {
	id := util.ParseUUID(req.Id)
	return uc.repo.FindById(ctx, uc.repo.GetEntClient(), id)
}

func (uc *UserUseCase) GetIdByEmail(ctx context.Context, req *pb.GetIdByEmailReq) uuid.UUID {
	return uc.repo.FindIdByEmail(ctx, uc.repo.GetEntClient(), req.Email)
}
