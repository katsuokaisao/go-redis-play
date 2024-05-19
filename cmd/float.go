package cmd

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var floatCmd = &cobra.Command{
	Use:   "float-sample",
	Short: "Float example",
	Long:  `Float example`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		floatCli := iredis.NewFloatExampleRepository(basicCli)

		testDt := map[uint]float64{
			1:   100.1,
			100: 200.2,
		}

		for k, v := range testDt {
			if err := floatCli.Set(k, v); err != nil {
				panic(err)
			}
			fmt.Printf("id %d sets %f\n", k, v)
		}

		for k := range testDt {
			exists, err := floatCli.Exists(k)
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
			value, err := floatCli.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("id %d has value: %f\n", k, value)
		}

		values, err := floatCli.MGet(1, 100)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mget values: %v\n", values)

		for k := range testDt {
			if err := floatCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := floatCli.Exists(k)
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
			value, err := floatCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("id %d does not exist\n", k)
				} else {
					panic(err)
				}
			} else {
				fmt.Printf("id %d has value: %f\n", k, value)
			}
		}
	},
}
