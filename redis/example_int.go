package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type intExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewIntExampleRepository(
	cli domain.BasicRedisRepository,
) domain.IntExampleRepository {
	return &intExampleRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("int:%d", id)
		},
		ttl: 30 * time.Minute,
	}
}

func (repo *intExampleRepository) Set(id uint, value int64) error {
	return repo.cli.SetInt64(repo.keyFunc(id), value, repo.ttl)
}

func (repo *intExampleRepository) Get(id uint) (int64, error) {
	return repo.cli.GetInt64(repo.keyFunc(id))
}

func (repo *intExampleRepository) Del(id uint) error {
	return repo.cli.Del(repo.keyFunc(id))
}

func (repo *intExampleRepository) Exists(id uint) (bool, error) {
	return repo.cli.Exists(repo.keyFunc(id))
}
