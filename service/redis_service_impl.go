package service

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisServiceImpl struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisService(rdb *redis.Client, ctx context.Context) RedisService {
	return &RedisServiceImpl{
		rdb: rdb,
		ctx: ctx,
	}
}

func (r *RedisServiceImpl) Put(key string, value string) error {
	err := r.rdb.Set(r.ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisServiceImpl) Remove(key string) error {
	err := r.rdb.Del(r.ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisServiceImpl) Get(key string) (*string, error) {
	val, err := r.rdb.Get(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}
