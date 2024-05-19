package domain

import "time"

type BasicRedisRepository interface {
	Ping(retryNum int) error
	Close() error
	SetString(key, value string, ttl time.Duration) error
	GetString(key string) (string, error)
	MGetString(keys ...string) ([]*string, error)
	SetInt64(key string, value int64, ttl time.Duration) error
	GetInt64(key string) (int64, error)
	MGetInt64(keys ...string) ([]*int64, error)
	SetFloat64(key string, value float64, ttl time.Duration) error
	GetFloat64(key string) (float64, error)
	MGetFloat64(keys ...string) ([]*float64, error)
	SetBool(key string, value bool, ttl time.Duration) error
	GetBool(key string) (bool, error)
	MGetBool(keys ...string) ([]*bool, error)
	SetBytes(key string, value []byte, ttl time.Duration) error
	GetBytes(key string) ([]byte, error)
	MGetBytes(keys ...string) ([][]byte, error)
	SetTime(key string, value time.Time, ttl time.Duration) error
	GetTime(key string) (time.Time, error)
	MGetTime(keys ...string) ([]*time.Time, error)
	Exists(key string) (bool, error)
	Del(key string) error
}

type StrExampleRepository interface {
	Set(id uint, value string) error
	Get(id uint) (string, error)
	MGet(ids ...uint) (map[uint]string, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type IntExampleRepository interface {
	Set(id uint, value int64) error
	Get(id uint) (int64, error)
	MGet(ids ...uint) (map[uint]int64, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type FloatExampleRepository interface {
	Set(id uint, value float64) error
	Get(id uint) (float64, error)
	MGet(ids ...uint) (map[uint]float64, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type BoolExampleRepository interface {
	Set(id uint, value bool) error
	Get(id uint) (bool, error)
	MGet(ids ...uint) (map[uint]bool, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type BytesExampleRepository interface {
	Set(id uint, value []byte) error
	Get(id uint) ([]byte, error)
	MGet(ids ...uint) (map[uint][]byte, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type TimeExampleRepository interface {
	Set(id uint, value time.Time) error
	Get(id uint) (time.Time, error)
	MGet(ids ...uint) (map[uint]time.Time, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}

type JSONExampleRepository interface {
	Set(id uint, value *Example) error
	Get(id uint) (*Example, error)
	MGet(ids ...uint) (map[uint]*Example, error)
	Del(id uint) error
	Exists(id uint) (bool, error)
}
