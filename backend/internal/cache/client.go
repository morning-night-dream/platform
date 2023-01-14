package cache

import (
	"context"
	"errors"
	"sync"

	"github.com/morning-night-dream/platform/internal/model"
)

type Client struct {
	lock  sync.Mutex
	cache map[string]model.Auth
}

func NewClient() *Client {
	return &Client{
		cache: make(map[string]model.Auth),
	}
}

func (c *Client) Get(ctx context.Context, key string) (model.Auth, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if val, ok := c.cache[key]; ok {
		return val, nil
	}

	return model.Auth{}, errors.New("mis cache")
}

func (c *Client) Set(ctx context.Context, key string, val model.Auth) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[key] = val

	return nil
}
