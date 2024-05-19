package domain

type Config struct {
	Redis   *RedisCfg
	IsLocal bool `env:"IS_LOCAL"`
}

type RedisCfg struct {
	Addr     string `env:"REDIS_ADDR"`
	Username string `env:"REDIS_USERNAME"`
	Password string `env:"REDIS_PASSWORD"`
}
