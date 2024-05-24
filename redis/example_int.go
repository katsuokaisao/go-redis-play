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
	return repo.cli.Set(repo.keyFunc(id), value, repo.ttl)
}

func (repo *intExampleRepository) MSet(values map[uint]int64) error {
	args := make(map[string]interface{})
	for id, value := range values {
		args[repo.keyFunc(id)] = value
	}

	return repo.cli.MSet(args)
}

func (repo *intExampleRepository) Get(id uint) (int64, error) {
	return repo.cli.GetInt64(repo.keyFunc(id))
}

func (repo *intExampleRepository) MGet(ids ...uint) (map[uint]int64, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = repo.keyFunc(id)
	}

	values, err := repo.cli.MGetInt64(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint]int64, len(ids))
	for i, id := range ids {
		if values[i] == nil {
			continue
		}
		res[id] = *values[i]
	}

	return res, nil
}

func (repo *intExampleRepository) Del(id uint) error {
	return repo.cli.Del(repo.keyFunc(id))
}

func (repo *intExampleRepository) Exists(id uint) (bool, error) {
	return repo.cli.Exists(repo.keyFunc(id))
}
