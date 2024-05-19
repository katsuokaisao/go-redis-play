package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time-example",
	Short: "Time example",
	Long:  `Time example`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		timeCli := iredis.NewTimeExampleRepository(basicCli)

		testDt := map[uint]time.Time{
			1:   time.Now(),
			100: time.Now().Add(time.Hour),
		}

		for k, v := range testDt {
			if err := timeCli.Set(k, v); err != nil {
				panic(err)
			}
			fmt.Printf("id %d sets %s\n", k, v)
		}

		for k := range testDt {
			exists, err := timeCli.Exists(k)
			if err != nil {
				panic(err)
			}
			if exists {
				fmt.Printf("id %d exists\n", k)
			} else {
				fmt.Printf("id %d does not exist\n", k)
			}
		}

		for k := range testDt {
			value, err := timeCli.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("id %d has value: %s\n", k, value)
		}

		for k := range testDt {
			if err := timeCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := timeCli.Exists(k)
			if err != nil {
				panic(err)
			}
			if exists {
				fmt.Printf("id %d exists\n", k)
			} else {
				fmt.Printf("id %d does not exist\n", k)
			}
		}

		for k := range testDt {
			value, err := timeCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("id %d does not exist\n", k)
					continue
				}
				panic(err)
			}
			fmt.Printf("id %d has value: %s\n", k, value)
		}
	},
}
