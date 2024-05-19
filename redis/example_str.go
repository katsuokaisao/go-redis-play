package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type strExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewStrExampleRepository(
	cli domain.BasicRedisRepository,
) domain.StrExampleRepository {
	return &strExampleRepository{
		keyFunc: func(id uint) string {
			return fmt.Sprintf("str:%d", id)
		},
		ttl: 30 * time.Minute,
		cli: cli,
	}
}

func (repo *strExampleRepository) Set(id uint, value string) error {
	return repo.cli.SetString(repo.keyFunc(id), value, repo.ttl)
}

func (repo *strExampleRepository) Get(id uint) (string, error) {
	return repo.cli.GetString(repo.keyFunc(id))
}

func (repo *strExampleRepository) MGet(ids ...uint) (map[uint]string, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = repo.keyFunc(id)
	}

	values, err := repo.cli.MGet(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint]string, len(ids))
	for i, id := range ids {
		if values[i] == nil {
			continue
		}
		res[id] = values[i].(string)
	}

	return res, nil
}

func (repo *strExampleRepository) Del(id uint) error {
	return repo.cli.Del(repo.keyFunc(id))
}

func (repo *strExampleRepository) Exists(id uint) (bool, error) {
	return repo.cli.Exists(repo.keyFunc(id))
}
