package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/morning-night-dream/platform/internal/model"
)

type Client struct {
	lock  sync.Mutex
	cache map[string]Cache
}

const ttl = 60 * time.Minute

type Cache struct {
	model.Auth
	CreatedAt time.Time
}

func NewClient() *Client {
	return &Client{
		cache: make(map[string]Cache),
	}
}

func (c *Client) Get(ctx context.Context, key string) (model.Auth, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if val, ok := c.cache[key]; ok && val.CreatedAt.Before(time.Now().Add(ttl)) {
		return val.Auth, nil
	}

	return model.Auth{}, errors.New("mis cache")
}

func (c *Client) Set(ctx context.Context, key string, val model.Auth) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[key] = Cache{
		Auth:      val,
		CreatedAt: time.Now(),
	}

	return nil
}

func (c *Client) Delete(ctx context.Context, key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.cache, key)

	return nil
}
