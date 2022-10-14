package store

import (
	"context"

	"github.com/morning-night-dream/platform/pkg/ent"
)

type Auth struct {
	db *ent.Client
}

func NewAuth(db *ent.Client) *Auth {
	return &Auth{
		db: db,
	}
}

func (a *Auth) Save(ctx context.Context) error {
	return nil
}
