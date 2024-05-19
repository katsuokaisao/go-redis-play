package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type jsonExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewJsonExampleRepository(cli domain.BasicRedisRepository) domain.JSONExampleRepository {
	return &jsonExampleRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("json:%d", id)
		},
		ttl: -1 * time.Second,
	}
}

func (r *jsonExampleRepository) Set(id uint, value *domain.Example) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	s := string(b)
	return r.cli.SetString(r.keyFunc(id), s, r.ttl)
}

func (r *jsonExampleRepository) Get(id uint) (*domain.Example, error) {
	s, err := r.cli.GetString(r.keyFunc(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get string: %w", err)
	}
	var e domain.Example
	if err := json.Unmarshal([]byte(s), &e); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return &e, nil
}

func (r *jsonExampleRepository) Del(id uint) error {
	return r.cli.Del(r.keyFunc(id))
}

func (r *jsonExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
