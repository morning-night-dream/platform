package store

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/internal/model"
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

	tx, err := a.db.Tx(ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction")
	}

	now := time.Now().UTC()

	err = tx.User.Create().
		SetID(id).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Exec(ctx)
	if err != nil {
		log.Printf("failed to save user %s", err)

		if re := tx.Rollback(); re != nil {
			return errors.Wrap(re, "")
		}

		return errors.Wrap(err, "")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(auth.Password), cost)

	err = tx.Auth.Create().
		SetID(id).
		SetLoginID(auth.LoginID).
		SetEmail(auth.Email).
		SetUserID(id).
		SetPassword(string(hashed)).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Exec(ctx)
	if err != nil {
		log.Printf("failed to save auth %s", err)

		if re := tx.Rollback(); re != nil {
			return errors.Wrap(re, "")
		}

		return errors.Wrap(err, "")
	}

	if ce := tx.Commit(); ce != nil {
		return errors.Wrap(ce, "")
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
