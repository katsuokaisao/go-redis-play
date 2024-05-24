package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type floatExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewFloatExampleRepository(cli domain.BasicRedisRepository) domain.FloatExampleRepository {
	return &floatExampleRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("float:%d", id)
		},
		ttl: -1 * time.Second,
	}
}

func (r *floatExampleRepository) Set(id uint, value float64) error {
	return r.cli.Set(r.keyFunc(id), value, r.ttl)
}

func (r *floatExampleRepository) Get(id uint) (float64, error) {
	return r.cli.GetFloat64(r.keyFunc(id))
}

func (r *floatExampleRepository) MGet(ids ...uint) (map[uint]float64, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = r.keyFunc(id)
	}

	values, err := r.cli.MGetFloat64(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint]float64, len(ids))
	for i, id := range ids {
		if values[i] == nil {
			continue
		}
		res[id] = *values[i]
	}

	return res, nil
}

func (r *floatExampleRepository) Del(id uint) error {
	return r.cli.Del(r.keyFunc(id))
}

func (r *floatExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
