package redis

import (
	proto "badge-service/protobuf"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	SetBadge(ctx context.Context, key string, badge *proto.Badge, expiration time.Duration) error
	SetMerchantBadge(ctx context.Context, key string, ShopBadge *proto.ShopBadge, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(addr string, password string, db int) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %w", err)
	}

	return &redisClient{client: client}, nil
}

// Get implements RedisClient.
func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// SetBadge implements RedisClient.
func (r *redisClient) SetBadge(ctx context.Context, key string, badge *proto.Badge, expiration time.Duration) error {
	badgeData, err := json.Marshal(badge)
	log.Println(badgeData)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	return r.client.Set(ctx, key, string(badgeData), expiration).Err()
}

// SetMerchantBadge implements RedisClient.
func (r *redisClient) SetMerchantBadge(ctx context.Context, key string, ShopBadge *proto.ShopBadge, expiration time.Duration) error {
	shopData, err := json.Marshal(ShopBadge)
	log.Println(shopData)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	return r.client.Set(ctx, key, string(shopData), expiration).Err()
}

func (r *redisClient) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("error deleting key from Redis: %w", err)
	}
	return nil
}
