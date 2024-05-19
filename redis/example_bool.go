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
	return r.cli.SetBool(r.keyFunc(id), value, r.ttl)
}

func (r *boolExampleRepository) Get(id uint) (bool, error) {
	return r.cli.GetBool(r.keyFunc(id))
}

func (r *boolExampleRepository) Del(id uint) error {
	return r.cli.Del(r.keyFunc(id))
}

func (r *boolExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
