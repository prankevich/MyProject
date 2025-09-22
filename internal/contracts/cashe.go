package contracts

import (
	"context"
	"time"
)

type CacheI interface {
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string, response interface{}) error
}
