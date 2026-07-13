package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/config"
	goredis "github.com/redis/go-redis/v9"
)

type Client struct {
	Client *goredis.Client
}

func New(cfg config.RedisConfig) (*Client, error) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		_ = rdb.Close()
		return nil, err
	}

	return &Client{
		Client: rdb,
	}, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}

func (c *Client) Health(ctx context.Context) error {
	return c.Client.Ping(ctx).Err()
}
