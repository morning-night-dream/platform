package store

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/core/model"
	"github.com/morning-night-dream/platform/pkg/ent"
	"github.com/pkg/errors"
)

type Auth struct {
	db       *ent.Client
	firebase *FirebaseClient
}

func NewAuth(db *ent.Client, firebase *FirebaseClient) *Auth {
	return &Auth{
		db:       db,
		firebase: firebase,
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

	err = a.firebase.CreateUser(ctx, auth.UserID, auth.Email, auth.Password)
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

func (a *Auth) FindFromEmailPass(ctx context.Context, email, pass string) (model.Tokens, error) {
	res, err := a.firebase.Login(ctx, email, pass)
	if err != nil {
		return model.Tokens{}, errors.Wrap(err, "not found")
	}

	return res, nil
}

func (a *Auth) FindFromIDToken(ctx context.Context, token string) (model.Tokens, error) {
	res, err := a.firebase.RefreshToken(ctx, token)
	if err != nil {
		return model.Tokens{}, errors.Wrap(err, "not found")
	}

	return res, nil
}
