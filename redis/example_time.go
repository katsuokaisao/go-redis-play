package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type timeExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewTimeExampleRepository(cli domain.BasicRedisRepository) domain.TimeExampleRepository {
	return &timeExampleRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("time:%d", id)
		},
		ttl: -1 * time.Second,
	}
}

func (r *timeExampleRepository) Set(id uint, value time.Time) error {
	return r.cli.SetTime(r.keyFunc(id), value, r.ttl)
}

func (r *timeExampleRepository) Get(id uint) (time.Time, error) {
	return r.cli.GetTime(r.keyFunc(id))
}

func (r *timeExampleRepository) MGet(ids ...uint) (map[uint]time.Time, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = r.keyFunc(id)
	}

	values, err := r.cli.MGet(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint]time.Time, len(ids))
	for i, id := range ids {
		if values[i] == nil {
			continue
		}
		res[id] = values[i].(time.Time)
	}

	return res, nil
}

func (r *timeExampleRepository) Del(id uint) error {
	return r.cli.Del(r.keyFunc(id))
}

func (r *timeExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
