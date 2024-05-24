package cmd

import (
	"fmt"

	"github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var incrDecrCmd = &cobra.Command{
	Use:   "incr-decr-sample",
	Short: "Incr and decr example",
	Long:  `Incr and decr example`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		incrDecrRepo := redis.NewIncrDecrRepository(basicCli)

		id := uint(1)
		incr, err := incrDecrRepo.Incr(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Incr: %d\n", incr)

		{
			value, err := incrDecrRepo.Get(id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Get: %d\n", value)
		}

		decr, err := incrDecrRepo.Decr(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Decr: %d\n", decr)

		{
			value, err := incrDecrRepo.Get(id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Get: %d\n", value)
		}

		incrBy, err := incrDecrRepo.IncrBy(id, 10)
		if err != nil {
			panic(err)
		}
		fmt.Printf("IncrBy: %d\n", incrBy)

		{
			value, err := incrDecrRepo.Get(id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Get: %d\n", value)
		}

		decrBy, err := incrDecrRepo.DecrBy(id, 5)
		if err != nil {
			panic(err)
		}
		fmt.Printf("DecrBy: %d\n", decrBy)

		{
			value, err := incrDecrRepo.Get(id)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Get: %d\n", value)
		}

		if err = incrDecrRepo.Delete(id); err != nil {
			panic(err)
		}
	},
}
