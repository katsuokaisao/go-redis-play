package redis

import (
	"encoding/base64"
	"fmt"
	"strconv"
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

func (r *basicRedisRepository) MGetString(keys ...string) ([]*string, error) {
	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*string, 0, len(keys))
	for _, v := range values {
		if v == nil {
			res = append(res, nil)
			continue
		}
		s := v.(string)
		res = append(res, &s)
	}

	return res, nil
}

func (r *basicRedisRepository) SetInt64(key string, value int64, ttl time.Duration) error {
	return r.cli.Set(key, fmt.Sprintf("%d", value), ttl).Err()
}

func (r *basicRedisRepository) GetInt64(key string) (int64, error) {
	return r.cli.Get(key).Int64()
}

func (r *basicRedisRepository) MGetInt64(keys ...string) ([]*int64, error) {
	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*int64, 0, len(keys))
	for _, v := range values {
		if v == nil {
			res = append(res, nil)
			continue
		}
		s := v.(string)
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, &i)
	}

	return res, nil
}

func (r *basicRedisRepository) SetFloat64(key string, value float64, ttl time.Duration) error {
	return r.cli.Set(key, fmt.Sprintf("%f", value), ttl).Err()
}

func (r *basicRedisRepository) GetFloat64(key string) (float64, error) {
	return r.cli.Get(key).Float64()
}

func (r *basicRedisRepository) MGetFloat64(keys ...string) ([]*float64, error) {
	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*float64, 0, len(keys))
	for _, v := range values {
		if v == nil {
			res = append(res, nil)
			continue
		}
		s := v.(string)
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, &f)
	}

	return res, nil
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

func (r *basicRedisRepository) MGetBool(keys ...string) ([]*bool, error) {
	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*bool, 0, len(keys))
	for _, v := range values {
		if v == nil {
			res = append(res, nil)
			continue
		}
		i := v.(int)
		b := i == 1
		res = append(res, &b)
	}

	return res, nil
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

func (r *basicRedisRepository) MGetBytes(keys ...string) ([][]byte, error) {
	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	res := make([][]byte, 0, len(keys))
	for _, v := range values {
		if v == nil {
			res = append(res, nil)
			continue
		}
		s := v.(string)
		b, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return nil, err
		}
		res = append(res, b)
	}

	return res, nil
}

func (r *basicRedisRepository) SetTime(key string, value time.Time, ttl time.Duration) error {
	return r.cli.Set(key, value.Format(time.RFC3339), ttl).Err()
}

func (r *basicRedisRepository) GetTime(key string) (time.Time, error) {
	return r.cli.Get(key).Time()
}

func (r *basicRedisRepository) MGetTime(keys ...string) ([]*time.Time, error) {
	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	res := make([]*time.Time, 0, len(keys))
	for _, v := range values {
		if v == nil {
			res = append(res, nil)
			continue
		}
		s := v.(string)
		t, err := time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return nil, err
		}
		res = append(res, &t)
	}

	return res, nil
}

func (r *basicRedisRepository) Exists(key string) (bool, error) {
	res, err := r.cli.Exists(key).Result()
	return res == 1, err
}

func (r *basicRedisRepository) Del(key string) error {
	return r.cli.Del(key).Err()
}
