package cmd

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var strCmd = &cobra.Command{
	Use: "str-sample",
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		strCli := iredis.NewStrExampleRepository(basicCli)

		testDt := map[uint]string{
			1:   "Hello World!",
			100: "Hello Redis!",
		}

		for k, v := range testDt {
			if err := strCli.Set(k, v); err != nil {
				panic(err)
			}
			fmt.Printf("id %d sets %s\n", k, v)
		}

		for k := range testDt {
			exists, err := strCli.Exists(k)
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
			value, err := strCli.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("id %d has value: %s\n", k, value)
		}

		for k := range testDt {
			if err := strCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := strCli.Exists(k)
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
			value, err := strCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("id: %d, does not exist\n", k)
					continue
				}
				panic(err)
			}
			fmt.Printf("id %d has value: %s\n", k, value)
		}
	},
}
