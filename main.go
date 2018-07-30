package main

import (
	"context"
	"log"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Configuration
type Config struct {
	TxHash *string
	Server *string
}

var conf Config

func init() {
	conf.Server = kingpin.Flag("server", "eth server").Required().String()
	conf.TxHash = kingpin.Flag("txhash", "TX Hash").Required().String()
}

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()

	// Dial Ethereum service
	conn, err := ethclient.Dial(*conf.Server)
	if err != nil {
		log.Fatalf("Error dialing %s:%s", *conf.TxHash, err.Error())
	}

	ctx := context.Background()

	// Get Transaction info
	tx, isPending, err := conn.TransactionByHash(ctx, common.HexToHash(*conf.TxHash))
	if err != nil {
		log.Fatalf("Error looking for-> %s :%s", *conf.TxHash, err.Error())
	}
	if !isPending {
		spew.Dump(tx)
	}
}
