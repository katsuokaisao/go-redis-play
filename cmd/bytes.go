package cmd

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v7"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var bytesCmd = &cobra.Command{
	Use:   "bytes-sample",
	Short: "Bytes example",
	Long:  `Bytes example`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		bytesCli := iredis.NewBytesExampleRepository(basicCli)

		testDt := map[uint][]byte{
			1:   []byte("hello"),
			100: []byte("world"),
		}

		for k, v := range testDt {
			if err := bytesCli.Set(k, v); err != nil {
				panic(err)
			}
			fmt.Printf("set id %d sets %s\n", k, v)
		}

		for k := range testDt {
			exists, err := bytesCli.Exists(k)
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
			value, err := bytesCli.Get(k)
			if err != nil {
				panic(err)
			}
			fmt.Printf("get id %d has value: %s\n", k, value)
		}

		values, err := bytesCli.MGet(1, 100)
		if err != nil {
			panic(err)
		}
		for k, v := range values {
			fmt.Printf("mget id %d has value: %s\n", k, v)
		}

		for k := range testDt {
			if err := bytesCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("del id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := bytesCli.Exists(k)
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
			value, err := bytesCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("get id %d does not exist\n", k)
					continue
				}
				panic(err)
			}
			fmt.Printf("get id %d has value: %s\n", k, value)
		}

	},
}
