package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type incrDecrRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewIncrDecrRepository(cli domain.BasicRedisRepository) domain.IncrDecrRepository {
	return &incrDecrRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("incr-decr:%d", id)
		},
		ttl: -1 * time.Second,
	}
}

func (repo *incrDecrRepository) Get(id uint) (int64, error) {
	return repo.cli.GetInt64(repo.keyFunc(id))
}

func (repo *incrDecrRepository) Delete(id uint) error {
	return repo.cli.Del(repo.keyFunc(id))
}

func (repo *incrDecrRepository) Incr(id uint) (int64, error) {
	return repo.cli.Incr(repo.keyFunc(id))
}

func (repo *incrDecrRepository) Decr(id uint) (int64, error) {
	return repo.cli.Decr(repo.keyFunc(id))
}

func (repo *incrDecrRepository) IncrBy(id uint, value int64) (int64, error) {
	return repo.cli.IncrBy(repo.keyFunc(id), value)
}

func (repo *incrDecrRepository) DecrBy(id uint, value int64) (int64, error) {
	return repo.cli.DecrBy(repo.keyFunc(id), value)
}
