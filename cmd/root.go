package cmd

import (
	"encoding/json"
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
	rootCmd.AddCommand(floatCmd)
	rootCmd.AddCommand(boolCmd)
	rootCmd.AddCommand(bytesCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(jsonCmd)
	rootCmd.AddCommand(incrDecrCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(basicCmd)
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

func printJSON(o interface{}, msg string) error {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n%s\n", msg, string(b))
	return nil
}