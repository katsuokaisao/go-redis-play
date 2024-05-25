package cmd

import (
	"fmt"
	"time"

	"github.com/katsuokaisao/go-redis-play/domain"
	iredis "github.com/katsuokaisao/go-redis-play/redis"
	"github.com/spf13/cobra"
)

var basicCmd = &cobra.Command{
	Use:   "basic-cmd-sample",
	Short: "Basic command example",
	Run: func(cmd *cobra.Command, args []string) {
		basicCli := initRedisRepository()
		defer basicCli.Close()

		basicExampleCli := iredis.NewBasicExampleRepository(basicCli)

		testStrDt := map[uint]string{
			1:   "Hello World!",
			100: "Hello Redis!",
		}
		tesListDt := []domain.Example{
			{Str: "a", Int: 1, Flt: 1.1, Bool: true, Byte: []byte("a"), Tm: time.Now().Add(3 * time.Hour)},
			{Str: "b", Int: 2, Flt: 2.2, Bool: false, Byte: []byte("b"), Tm: time.Now().Add(4 * time.Hour)},
			{Str: "c", Int: 3, Flt: 3.3, Bool: true, Byte: []byte("c"), Tm: time.Now().Add(5 * time.Hour)},
			{Str: "d", Int: 4, Flt: 4.4, Bool: false, Byte: []byte("d"), Tm: time.Now().Add(6 * time.Hour)},
		}

		for k, v := range testStrDt {
			if err := basicExampleCli.Set(k, v); err != nil {
				panic(err)
			}
		}

		if err := basicExampleCli.LPush(tesListDt); err != nil {
			panic(err)
		}

		dbSize, err := basicExampleCli.DBSize()
		if err != nil {
			panic(err)
		}
		fmt.Printf("DBSize: %d\n", dbSize)

		scanKeys, err := basicExampleCli.ScanStr()
		if err != nil {
			panic(err)
		}
		fmt.Printf("ScanStr: %v\n", scanKeys)

		scanKeys, err = basicExampleCli.ScanList()
		if err != nil {
			panic(err)
		}
		fmt.Printf("ScanList: %v\n", scanKeys)

		if err := basicExampleCli.UnlinkStr([]uint{1, 100}); err != nil {
			panic(err)
		}

		if err := basicExampleCli.UnlinkList(); err != nil {
			panic(err)
		}

		for {
			keys, err := basicCli.LRange("basic-list", 0, -1)
			if err != nil {
				panic(err)
			}
			if len(keys) == 0 {
				fmt.Println("unlinked all keys")
				break
			}

			fmt.Printf("List keys: %v\n", keys)
			time.Sleep(1 * time.Second)
		}
	},
}
