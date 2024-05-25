package redis

import (
	"encoding/json"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type listExampleRepository struct {
	cli domain.BasicRedisRepository
	key string
	ttl time.Duration
}

func NewListExampleRepository(
	cli domain.BasicRedisRepository,
) domain.ListExampleRepository {
	return &listExampleRepository{
		cli: cli,
		key: "list",
		ttl: 30 * time.Minute,
	}
}

func (repo *listExampleRepository) LPush(values []domain.Example) error {
	args := make([]interface{}, len(values))
	for i, value := range values {
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		args[i] = string(b)
	}

	return repo.cli.LPush(repo.key, args...)
}

func (repo *listExampleRepository) RPush(values []domain.Example) error {
	args := make([]interface{}, len(values))
	for i, value := range values {
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		args[i] = string(b)
	}

	return repo.cli.RPush(repo.key, args...)
}

func (repo *listExampleRepository) LPop() (*domain.Example, error) {
	value, err := repo.cli.LPop(repo.key)
	if err != nil {
		return nil, err
	}

	var example domain.Example
	if err := json.Unmarshal([]byte(value), &example); err != nil {
		return nil, err
	}

	return &example, nil
}

func (repo *listExampleRepository) RPop() (*domain.Example, error) {
	value, err := repo.cli.RPop(repo.key)
	if err != nil {
		return nil, err
	}

	var example domain.Example
	if err := json.Unmarshal([]byte(value), &example); err != nil {
		return nil, err
	}

	return &example, nil
}

func (repo *listExampleRepository) LLen() (int64, error) {
	return repo.cli.LLen(repo.key)
}

func (repo *listExampleRepository) LRange(start, stop int64) ([]domain.Example, error) {
	values, err := repo.cli.LRange(repo.key, start, stop)
	if err != nil {
		return nil, err
	}

	examples := make([]domain.Example, len(values))
	for i, value := range values {
		if err := json.Unmarshal([]byte(value), &examples[i]); err != nil {
			return nil, err
		}
	}

	return examples, nil
}

func (repo *listExampleRepository) Del() error {
	return repo.cli.Del(repo.key)
}