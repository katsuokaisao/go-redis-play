package redis

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type bytesExampleRepository struct {
	cli     domain.BasicRedisRepository
	keyFunc func(id uint) string
	ttl     time.Duration
}

func NewBytesExampleRepository(cli domain.BasicRedisRepository) domain.BytesExampleRepository {
	return &bytesExampleRepository{
		cli: cli,
		keyFunc: func(id uint) string {
			return fmt.Sprintf("bytes:%d", id)
		},
		ttl: -1 * time.Second,
	}
}

func (r *bytesExampleRepository) Set(id uint, value []byte) error {
	return r.cli.SetBytes(r.keyFunc(id), value, r.ttl)
}

func (r *bytesExampleRepository) Get(id uint) ([]byte, error) {
	return r.cli.GetBytes(r.keyFunc(id))
}

func (r *bytesExampleRepository) Del(id uint) error {
	return r.cli.Del(r.keyFunc(id))
}

func (r *bytesExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
