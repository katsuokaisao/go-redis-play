package redis

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/katsuokaisao/go-redis-play/domain"
)

type basicRedisRepository struct {
	cli *redis.Client
}

func NewBasicRedisRepository(cli *redis.Client) domain.BasicRedisRepository {
	return &basicRedisRepository{cli: cli}
}

func (r *basicRedisRepository) Ping(retryNum int) error {
	for i := 0; i < retryNum; i++ {
		if err := r.cli.Ping().Err(); err != nil {
			fmt.Printf("Ping error: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		res, err := r.cli.Ping().Result()
		if err != nil {
			fmt.Printf("Ping error: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("Ping result: %v\n", res)
		return nil
	}

	return fmt.Errorf("Ping failed")
}

func (r *basicRedisRepository) Close() error {
	return r.cli.Close()
}

func (r *basicRedisRepository) SetString(key string, value string, ttl time.Duration) error {
	return r.cli.Set(key, value, ttl).Err()
}

func (r *basicRedisRepository) GetString(key string) (string, error) {
	return r.cli.Get(key).Result()
}

func (r *basicRedisRepository) SetInt64(key string, value int64, ttl time.Duration) error {
	return r.cli.Set(key, fmt.Sprintf("%d", value), ttl).Err()
}

func (r *basicRedisRepository) GetInt64(key string) (int64, error) {
	return r.cli.Get(key).Int64()
}

func (r *basicRedisRepository) SetFloat64(key string, value float64, ttl time.Duration) error {
	return r.cli.Set(key, fmt.Sprintf("%f", value), ttl).Err()
}

func (r *basicRedisRepository) GetFloat64(key string) (float64, error) {
	return r.cli.Get(key).Float64()
}

func (r *basicRedisRepository) SetBool(key string, value bool, ttl time.Duration) error {
	v := "0"
	if value {
		v = "1"
	}
	return r.cli.Set(key, v, ttl).Err()
}

func (r *basicRedisRepository) GetBool(key string) (bool, error) {
	v, err := r.cli.Get(key).Int()
	if err != nil {
		return false, err
	}
	return v == 1, nil
}

func (r *basicRedisRepository) SetBytes(key string, value []byte, ttl time.Duration) error {
	s := base64.StdEncoding.EncodeToString([]byte(value))
	return r.cli.Set(key, s, ttl).Err()
}

func (r *basicRedisRepository) GetBytes(key string) ([]byte, error) {
	s, err := r.cli.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(s)
}

func (r *basicRedisRepository) SetTime(key string, value time.Time, ttl time.Duration) error {
	return r.cli.Set(key, value.Format(time.RFC3339), ttl).Err()
}

func (r *basicRedisRepository) GetTime(key string) (time.Time, error) {
	return r.cli.Get(key).Time()
}

func (r *basicRedisRepository) MGet(keys ...string) ([]interface{}, error) {
	return r.cli.MGet(keys...).Result()
}

func (r *basicRedisRepository) Exists(key string) (bool, error) {
	res, err := r.cli.Exists(key).Result()
	return res == 1, err
}

func (r *basicRedisRepository) Del(key string) error {
	return r.cli.Del(key).Err()
}
