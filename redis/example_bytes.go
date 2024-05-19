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

func (r *bytesExampleRepository) MGet(ids ...uint) (map[uint][]byte, error) {
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = r.keyFunc(id)
	}

	values, err := r.cli.MGet(keys...)
	if err != nil {
		return nil, err
	}

	res := make(map[uint][]byte, len(ids))
	for i, id := range ids {
		if values[i] == nil {
			continue
		}
		res[id] = values[i].([]byte)
	}

	return res, nil
}

func (r *bytesExampleRepository) Del(id uint) error {
	return r.cli.Del(r.keyFunc(id))
}

func (r *bytesExampleRepository) Exists(id uint) (bool, error) {
	return r.cli.Exists(r.keyFunc(id))
}
