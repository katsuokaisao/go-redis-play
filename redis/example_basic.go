package redis

import (
	"encoding/json"
	"fmt"

	"github.com/katsuokaisao/go-redis-play/domain"
)

type basicExampleRepository struct {
	cli        domain.BasicRedisRepository
	strKeyFunc func(id uint) string
	listKey    string
}

func NewBasicExampleRepository(cli domain.BasicRedisRepository) domain.BasicExampleRepository {
	return &basicExampleRepository{
		cli: cli,
		strKeyFunc: func(id uint) string {
			return fmt.Sprintf("basic-str:%d", id)
		},
		listKey: "basic-list",
	}
}

func (r *basicExampleRepository) DBSize() (int64, error) {
	return r.cli.DBSize()
}

func (r *basicExampleRepository) ScanStr() ([]string, error) {
	count := int64(10)
	match := "basic-str:*"
	return r.cli.Scan(match, count)
}

func (r *basicExampleRepository) ScanList() ([]string, error) {
	count := int64(10)
	match := "basic-list"
	return r.cli.Scan(match, count)
}

func (r *basicExampleRepository) UnlinkStr(ids []uint) error {
	keys := make([]string, 0, len(ids))
	for _, id := range ids {
		keys = append(keys, r.strKeyFunc(id))
	}

	return r.cli.Unlink(keys...)
}

func (r *basicExampleRepository) UnlinkList() error {
	return r.cli.Unlink(r.listKey)
}

func (r *basicExampleRepository) Set(id uint, value string) error {
	return r.cli.Set(r.strKeyFunc(id), value, -1)
}

func (r *basicExampleRepository) LPush(values []domain.Example) error {
	args := make([]interface{}, 0, len(values))
	for _, v := range values {
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		args = append(args, string(b))
	}

	return r.cli.LPush(r.listKey, args...)
}
