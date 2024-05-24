package cmd

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var boolCmd = &cobra.Command{
	Use:   "bool-sample",
	Short: "Bool example",
	Long:  `Bool example`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		boolCli := iredis.NewBoolExampleRepository(basicCli)

		testDt := map[uint]bool{
			1:   true,
			100: false,
		}

		for k, v := range testDt {
			if err := boolCli.Set(k, v); err != nil {
				panic(err)
			}
			fmt.Printf("set id %d sets %t\n", k, v)
		}

		for k := range testDt {
			exists, err := boolCli.Exists(k)
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
			value, err := boolCli.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("get id %d has value: %t\n", k, value)
		}

		for k := range testDt {
			if err := boolCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("del id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := boolCli.Exists(k)
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
			value, err := boolCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("get id %d does not exist\n", k)
					continue
				} else {
					panic(err)
				}
			}
			fmt.Printf("get id %d has value: %t\n", k, value)
		}

		if err := boolCli.MSet(testDt); err != nil {
			panic(err)
		}
		fmt.Printf("mset values: %v\n", testDt)

		keys := []uint{1, 100}
		values, err := boolCli.MGet(keys...)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mget values: %v\n", values)

		for k := range testDt {
			if err := boolCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("del id %d deleted\n", k)
		}
	},
}
