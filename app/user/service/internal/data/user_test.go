package data

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/kx-boutique/app/user/service/internal/biz"
	entModel "github.com/kx-boutique/app/user/service/internal/biz/model"
	ent "github.com/kx-boutique/ent/generated"
	"github.com/kx-boutique/ent/generated/enttest"
	"github.com/kx-boutique/ent/generated/migrate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newEntClientTest(t *testing.T) *ent.Client {
	dataSource := "postgresql://user:1234@localhost:5678/kx_boutique_test?sslmode=disable"
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	eclient := enttest.Open(t, dialect.Postgres, dataSource, opts...)
	return eclient
}

func newRepo(t *testing.T, client *ent.Client) (biz.UserRepo, func()) {
	data, cleanup, err := NewData(client, log.DefaultLogger)
	if err != nil {
		t.Error(err)
	}
	return NewUserRepo(data, log.DefaultLogger), cleanup
}

func TestA(t *testing.T) {
	eclient := newEntClientTest(t)
	repo, cleanup := newRepo(t, eclient)
	defer cleanup()

	ctx := context.Background()
	name := gofakeit.Name()
	email := gofakeit.Email()
	u := repo.SaveUser(ctx, repo.GetEntClient(), &entModel.User{
		Name:  name,
		Email: email,
	})

	result := eclient.User.GetX(ctx, u.Id)

	assert.NotNil(t, result)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, email, result.Email)
}

func TestB(t *testing.T) {
}
