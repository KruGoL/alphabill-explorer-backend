package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/alphabill-org/alphabill/internal/txsystem/money"
	"github.com/alphabill-org/alphabill/pkg/explorer"
)

func main() {
	fmt.Println("Starting AB Explorer")
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Usage: explorer <AB Partition url> <AB Explorer url> [<Block number>]")
		return
	}
	workDir := filepath.Dir(args[0]) //"/tmp/"
	fmt.Printf("filepath: %s\n", filepath.Dir(args[0]))
	fmt.Printf("AB Partition url: %s\n", args[1])
	fmt.Printf("AB Explorer url: %s\n", args[2])
	blockNumber := uint64(0)
	if len(args) > 3 {
		fmt.Printf("Block number: %s\n", args[3])
		blockNumber, _ = strconv.ParseUint(args[3], 10, 64)
	}
	err := explorer.Run(context.Background(), &explorer.Config{
		ABMoneySystemIdentifier: money.DefaultSystemIdentifier,
		AlphabillUrl:            args[1],
		ServerAddr:              args[2],
		DbFile:                  filepath.Join(workDir, "explorer.db"),
		BlockNumber:             blockNumber,
	})
	if err != nil {
		panic(err)
	}
}

//var defaultMoneySDR = &genesis.SystemDescriptionRecord{
//	SystemIdentifier: money.DefaultSystemIdentifier,
//	T2Timeout:        2500,
//	FeeCreditBill: &genesis.FeeCreditBill{
//		UnitId:         money.NewBillID(nil, []byte{2}),
//		OwnerPredicate: script.PredicateAlwaysTrue(),
//	},
//}
//
//var defaultInitialBillID = money.NewBillID(nil, []byte{1})
