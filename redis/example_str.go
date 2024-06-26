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
	return repo.cli.Set(repo.keyFunc(id), value, repo.ttl)
}

func (repo *strExampleRepository) MSet(values map[uint]string) error {
	args := make(map[string]interface{})
	for id, value := range values {
		args[repo.keyFunc(id)] = value
	}

	return repo.cli.MSet(args)
}

func (repo *strExampleRepository) Get(id uint) (string, error) {
	return repo.cli.GetString(repo.keyFunc(id))
}

func (repo *strExampleRepository) MGet(ids ...uint) (map[uint]string, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = repo.keyFunc(id)
	}

	values, err := repo.cli.MGetString(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint]string, len(ids))
	for i, id := range ids {
		if len(*values[i]) == 0 {
			continue
		}
		res[id] = *values[i]
	}

	return res, nil
}

func (repo *strExampleRepository) Del(id uint) error {
	return repo.cli.Unlink(repo.keyFunc(id))
}

func (repo *strExampleRepository) Exists(id uint) (bool, error) {
	return repo.cli.Exists(repo.keyFunc(id))
}
