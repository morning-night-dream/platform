package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/morning-night-dream/platform/internal/errors"
	"github.com/morning-night-dream/platform/internal/model"
)

type Client struct {
	// lock  sync.Mutex
	cache  map[string]Cache
	client *redis.Client
}

// キャッシュへの保存期間は1週間　TODO: 環境変数設定
const ttl = 7 * 24 * time.Hour

type Cache struct {
	model.Auth
	CreatedAt time.Time
}

func NewClient(url string) *Client {
	var opt *redis.Options

	if model.Env.IsProd() {
		var err error
		opt, err = redis.ParseURL(url)
		if err != nil {
			panic(err)
		}
	} else {
		opt = &redis.Options{
			Addr:     url,
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}

	client := redis.NewClient(opt)

	return &Client{
		cache:  make(map[string]Cache),
		client: client,
	}
}

func (c *Client) Get(ctx context.Context, key string) (model.Auth, error) {
	// c.lock.Lock()
	// defer c.lock.Unlock()

	// if val, ok := c.cache[key]; ok && val.CreatedAt.Before(time.Now().Add(ttl)) {
	// 	return val.Auth, nil
	// }

	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return model.Auth{}, errors.NewNotFoundError("mis cache", err)
	}

	var value model.Auth

	if err := json.Unmarshal([]byte(val), &value); err != nil {
		return model.Auth{}, errors.NewValidationError("failed to unmarshal json", err)
	}

	return value, nil
}

func (c *Client) Set(ctx context.Context, key string, val model.Auth) error {
	// c.lock.Lock()
	// defer c.lock.Unlock()

	// c.cache[key] = Cache{
	// 	Auth:      val,
	// 	CreatedAt: time.Now(),
	// }

	value, err := json.Marshal(val)
	if err != nil {
		return err
	}

	if err := c.client.Set(ctx, key, string(value), ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(ctx context.Context, key string) error {
	// c.lock.Lock()
	// defer c.lock.Unlock()

	// delete(c.cache, key)

	if err := c.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
