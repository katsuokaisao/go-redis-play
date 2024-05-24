package domain

import "time"

type BasicRedisRepository interface {
	Ping(retryNum int) error
	Close() error
	Set(key string, value interface{}, ttl time.Duration) error
	SetBytes(key string, value []byte, ttl time.Duration) error
	MSet(data map[string]interface{}) error
	MSetBytes(data map[string][]byte) error
	GetString(key string) (string, error)
	MGetString(keys ...string) ([]*string, error)
	GetInt64(key string) (int64, error)
	MGetInt64(keys ...string) ([]*int64, error)
	GetFloat64(key string) (float64, error)
	MGetFloat64(keys ...string) ([]*float64, error)
	GetBool(key string) (bool, error)
	MGetBool(keys ...string) ([]*bool, error)
	GetBytes(key string) ([]byte, error)
	MGetBytes(keys ...string) ([][]byte, error)
	GetTime(key string) (time.Time, error)
	MGetTime(keys ...string) ([]*time.Time, error)
	Exists(key string) (bool, error)
	Del(key string) error
}

type StrExampleRepository interface {
	Set(id uint, value string) error
	MSet(values map[uint]string) error
	Get(id uint) (string, error)
	MGet(ids ...uint) (map[uint]string, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type IntExampleRepository interface {
	Set(id uint, value int64) error
	MSet(values map[uint]int64) error
	Get(id uint) (int64, error)
	MGet(ids ...uint) (map[uint]int64, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type FloatExampleRepository interface {
	Set(id uint, value float64) error
	MSet(values map[uint]float64) error
	Get(id uint) (float64, error)
	MGet(ids ...uint) (map[uint]float64, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type BoolExampleRepository interface {
	Set(id uint, value bool) error
	MSet(values map[uint]bool) error
	Get(id uint) (bool, error)
	MGet(ids ...uint) (map[uint]bool, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type BytesExampleRepository interface {
	Set(id uint, value []byte) error
	MSet(values map[uint][]byte) error
	Get(id uint) ([]byte, error)
	MGet(ids ...uint) (map[uint][]byte, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type TimeExampleRepository interface {
	Set(id uint, value time.Time) error
	MSet(values map[uint]time.Time) error
	Get(id uint) (time.Time, error)
	MGet(ids ...uint) (map[uint]time.Time, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type JSONExampleRepository interface {
	Set(id uint, value *Example) error
	MSet(values map[uint]Example) error
	Get(id uint) (*Example, error)
	MGet(ids ...uint) (map[uint]*Example, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}
