package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/katsuokaisao/go-redis-play/domain"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json-sample",
	Short: "JSON example",
	Long:  `JSON example`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		jsonCli := iredis.NewJsonExampleRepository(basicCli)

		testDt := map[uint]domain.Example{
			1: {
				Str:  "hello",
				Int:  1,
				Flt:  1.1,
				Byte: []byte("hello"),
				Bool: true,
				Tm:   time.Now(),
			},
			100: {
				Str:  "world",
				Int:  100,
				Flt:  100.1,
				Byte: []byte("world"),
				Bool: false,
				Tm:   time.Now().Add(time.Hour),
			},
		}

		for k := range testDt {
			v := testDt[k]
			if err := jsonCli.Set(k, &v); err != nil {
				panic(err)
			}
			b, err := json.MarshalIndent(testDt[k], "", "   ")
			if err != nil {
				panic(err)
			}
			fmt.Printf("set id %d sets %s\n", k, string(b))
		}

		for k := range testDt {
			exists, err := jsonCli.Exists(k)
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
			value, err := jsonCli.Get(k)
			if err != nil {
				panic(err)
			}
			b, err := json.MarshalIndent(value, "", "   ")
			if err != nil {
				panic(err)
			}
			fmt.Printf("get id %d has value: %s\n", k, string(b))
		}

		values, err := jsonCli.MGet(1, 100)
		if err != nil {
			panic(err)
		}
		for k, v := range values {
			b, err := json.MarshalIndent(v, "", "   ")
			if err != nil {
				panic(err)
			}
			fmt.Printf("mget id %d has value: %s\n", k, string(b))
		}

		for k := range testDt {
			if err := jsonCli.Del(k); err != nil {
				panic(err)
			}
			fmt.Printf("del id %d deleted\n", k)
		}

		for k := range testDt {
			exists, err := jsonCli.Exists(k)
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
			value, err := jsonCli.Get(k)
			if err != nil {
				if errors.Is(err, redis.Nil) {
					fmt.Printf("get id %d does not exist\n", k)
					continue
				}
				panic(err)
			}
			fmt.Printf("get id %d has value: %v\n", k, value)
		}
	},
}
