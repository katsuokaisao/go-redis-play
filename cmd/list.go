package cmd

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list-sample",
	Short: "List sample",
	Long:  `List sample`,
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		listCli := iredis.NewListExampleRepository(basicCli)

		testDt := []domain.Example{
			{Str: "a", Int: 1, Flt: 1.1, Bool: true, Byte: []byte("a"), Tm: time.Now().Add(3 * time.Hour)},
			{Str: "b", Int: 2, Flt: 2.2, Bool: false, Byte: []byte("b"), Tm: time.Now().Add(4 * time.Hour)},
			{Str: "c", Int: 3, Flt: 3.3, Bool: true, Byte: []byte("c"), Tm: time.Now().Add(5 * time.Hour)},
			{Str: "d", Int: 4, Flt: 4.4, Bool: false, Byte: []byte("d"), Tm: time.Now().Add(6 * time.Hour)},
		}

		// Queue: LPush -> RPop
		if err := listCli.LPush(testDt); err != nil {
			panic(err)
		}
		fmt.Println("LPush")

		cnt, err := listCli.LLen()
		if err != nil {
			panic(err)
		}
		fmt.Printf("LLen %d\n", cnt)

		for i := 0; i < len(testDt); i++ {
			value, err := listCli.RPop()
			if err != nil {
				panic(err)
			}
			printJSON(value, "RPop")
		}

		// Stack: RPush -> LPop
		if err := listCli.LPush(testDt); err != nil {
			panic(err)
		}
		fmt.Println("LPush")

		cnt, err = listCli.LLen()
		if err != nil {
			panic(err)
		}
		fmt.Printf("LLen %d\n", cnt)

		for i := 0; i < len(testDt); i++ {
			value, err := listCli.LPop()
			if err != nil {
				panic(err)
			}
			printJSON(value, "LPop")
		}

		// LRange test
		if err := listCli.LPush(testDt); err != nil {
			panic(err)
		}
		fmt.Println("LPush")

		cnt, err = listCli.LLen()
		if err != nil {
			panic(err)
		}
		fmt.Printf("LLen %d\n", cnt)

		values, err := listCli.LRange(0, 2)
		if err != nil {
			panic(err)
		}
		printJSON(values, "LRange: 0, 2")

		values, err = listCli.LRange(1, 3)
		if err != nil {
			panic(err)
		}
		printJSON(values, "LRange: 1, 3")

		values, err = listCli.LRange(2, 4)
		if err != nil {
			panic(err)
		}
		printJSON(values, "LRange: 2, 4")

		if err := listCli.Del(); err != nil {
			panic(err)
		}
		fmt.Println("Del")
	},
}
