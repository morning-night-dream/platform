package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/morning-night-dream/platform/pkg/ent"
	"github.com/morning-night-dream/platform/pkg/ent/auth"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

const cost = 12

type Auth struct {
	db *ent.Client
}

func NewAuth(db *ent.Client) *Auth {
	return &Auth{
		db: db,
	}
}

func (a *Auth) Save(ctx context.Context, auth model.Auth) error {
	id, err := uuid.Parse(auth.UserID)
	if err != nil {
		return errors.Wrap(err, "failed to uuid parse")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(auth.Password), cost)

	if err := a.db.Auth.Create().
		SetID(id).
		SetLoginID(auth.LoginID).
		SetEmail(auth.Email).
		SetPassword(string(hashed)).
		Exec(ctx); err != nil {
		return errors.Wrap(err, "failed to save")
	}

	return nil
}

func (a *Auth) FindFromIDPass(ctx context.Context, id, pass string) (model.Auth, error) {
	res, err := a.db.Auth.Query().
		Where(
			auth.LoginIDEQ(id),
		).
		First(ctx)
	if err != nil {
		return model.Auth{}, errors.Wrap(err, "not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(pass)); err != nil {
		return model.Auth{}, errors.Wrap(err, "not found")
	}

	return model.Auth{
		UserID:   res.ID.String(),
		LoginID:  res.LoginID,
		Email:    "", // 機密情報なのであえて隠しておく
		Password: "", // 機密情報なのであえて隠しておく
	}, nil
}
