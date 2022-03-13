package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func AddressCheck(address string, client *ethclient.Client) {
	//var (
	//	client = init.NewClient(conf.Conf.Chain.ChainURL)
	//)
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	account := common.HexToAddress(address)

	if re.MatchString(address) {
		fmt.Printf("%v is valid.\n", address)
	} else {
		fmt.Printf("%v is not valid.\n", address)
	}

	bytecode, err := client.CodeAt(context.Background(), account, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	if len(bytecode) > 0 {
		fmt.Printf("%v is contract address.\n", address)
	} else {
		fmt.Printf("%v is not contract address.\n", address)
	}
}
