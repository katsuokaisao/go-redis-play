package cmd

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var intCmd = &cobra.Command{
	Use:   "int-sample",
	Short: "int command",
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		intCli := iredis.NewIntExampleRepository(basicCli)

		testDt := map[uint]int64{
			1:   100,
			100: 200,
		}

		for k, v := range testDt {
			if err := intCli.Set(k, v); err != nil {
				panic(err)
			}
			fmt.Printf("set id %d sets %d\n", k, v)
		}

		for k := range testDt {
			exists, err := intCli.Exists(k)
			if err != nil {
				panic(err)
			}
			if exists {
				fmt.Printf("exists id %d exists\n", k)
			} else {
				fmt.Printf("exists id %d does not exist\n", k)
			}
		}

		for k := range testDt {
			value, err := intCli.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("get id %d has value: %d\n", k, value)
		}

		for k := range testDt {
			if err := intCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("del id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := intCli.Exists(k)
			if err != nil {
				panic(err)
			}
			if exists {
				fmt.Printf("exists id %d exists\n", k)
			} else {
				fmt.Printf("exists id %d does not exist\n", k)
			}
		}

		for k := range testDt {
			value, err := intCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("get id %d does not exist\n", k)
					continue
				}
				panic(err)
			}
			fmt.Printf("get id %d has value: %d\n", k, value)
		}

		if err := intCli.MSet(testDt); err != nil {
			panic(err)
		}
		fmt.Printf("mset values: %v\n", testDt)

		values, err := intCli.MGet(1, 100)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mget values: %v\n", values)

		for k := range testDt {
			if err := intCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("del id %d deleted\n", k)
		}
	},
}
