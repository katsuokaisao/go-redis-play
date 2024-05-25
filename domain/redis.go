package domain

import (
	"time"
)

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
	Incr(key string) (int64, error)
	Decr(key string) (int64, error)
	IncrBy(key string, value int64) (int64, error)
	DecrBy(key string, value int64) (int64, error)
	LPush(key string, values ...interface{}) error
	RPush(key string, values ...interface{}) error
	LPop(key string) (string, error)
	RPop(key string) (string, error)
	LLen(key string) (int64, error)
	LRange(key string, start, stop int64) ([]string, error)
	Exists(key string) (bool, error)
	Type(key string) (string, error)
	DBSize() (int64, error)
	// DEL & KEYS の使用は禁止
	// 代わりに UNLINK & SCAN を使用する
	Unlink(keys ...string) error
	Scan(match string, count int64) ([]string, error)
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

type IncrDecrRepository interface {
	Get(id uint) (int64, error)
	Delete(id uint) error
	Incr(id uint) (int64, error)
	Decr(id uint) (int64, error)
	IncrBy(id uint, value int64) (int64, error)
	DecrBy(id uint, value int64) (int64, error)
}

type ListExampleRepository interface {
	LPush(values []Example) error
	RPush(values []Example) error
	LPop() (*Example, error)
	RPop() (*Example, error)
	LLen() (int64, error)
	LRange(start, stop int64) ([]Example, error)
	Del() error
}

// Type, DBSize, Scan, Unlink test
type BasicExampleRepository interface {
	DBSize() (int64, error)
	ScanStr() ([]string, error)
	ScanList() ([]string, error)
	UnlinkStr(ids []uint) error
	UnlinkList() error
	Set(id uint, value string) error
	LPush(values []Example) error
}
