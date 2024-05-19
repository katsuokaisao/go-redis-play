package cmd

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	redisv7 "github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/katsuokaisao/go-redis-play/domain"
	"github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "",
}

func init() {
	rootCmd.AddCommand(strCmd)
	rootCmd.AddCommand(intCmd)
}

func initRedisRepository() domain.BasicRedisRepository {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	c := domain.Config{
		Redis: &domain.RedisCfg{},
	}
	if err := env.Parse(&c); err != nil {
		log.Fatalf("%+v\n", err)
	}
	if c.Redis == nil {
		panic("Redis config is nil")
	}

	var cli *redisv7.Client
	if c.IsLocal {
		log.Println("Local Redis")
		cli = redis.ProvideLocalRedis(c.Redis)
	} else {
		log.Println("Remote Redis")
		cli = redis.ProvideRemoteRedis(c.Redis)
	}

	repo := redis.NewBasicRedisRepository(cli)

	if err := repo.Ping(5); err != nil {
		panic(fmt.Sprintf("Failed to ping Redis: %v", err))
	}

	return repo
}

func Execute() error {
	return rootCmd.Execute()
}
