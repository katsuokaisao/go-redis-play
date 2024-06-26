package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type boolExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewBoolExampleRepository(cli domain.BasicRedisRepository) domain.BoolExampleRepository {
	return &boolExampleRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("bool:%d", id)
		},
		ttl: -1 * time.Second,
	}
}

func (r *boolExampleRepository) Set(id uint, value bool) error {
	return r.cli.Set(r.keyFunc(id), value, r.ttl)
}

func (r *boolExampleRepository) MSet(values map[uint]bool) error {
	args := make(map[string]interface{})
	for id, value := range values {
		args[r.keyFunc(id)] = value
	}

	return r.cli.MSet(args)
}

func (r *boolExampleRepository) Get(id uint) (bool, error) {
	return r.cli.GetBool(r.keyFunc(id))
}

func (r *boolExampleRepository) MGet(ids ...uint) (map[uint]bool, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = r.keyFunc(id)
	}

	values, err := r.cli.MGetBool(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint]bool, len(ids))
	for i, id := range ids {
		if values[i] == nil {
			continue
		}
		res[id] = *values[i]
	}

	return res, nil
}

func (r *boolExampleRepository) Del(id uint) error {
	return r.cli.Unlink(r.keyFunc(id))
}

func (r *boolExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
