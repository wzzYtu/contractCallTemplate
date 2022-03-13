package service

import (
	"contractCallTemplate/conf"
	"contractCallTemplate/init"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"context"
	"log"
	"math/big"
)

func GetBlockInfoByNum(blockNumber *big.Int) *types.Block {
	var client = init.NewClient(conf.Conf.Chain.ChainURL)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	return block
}

func GetLastBlockInfo() {
	var client = init.NewClient(conf.Conf.Chain.ChainURL)
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	GetBlockInfoByNum(header.Number)
}

func GetBlockByHash(blockHash string) *types.Block {
	var client = init.NewClient(conf.Conf.Chain.ChainURL)
	bHash := common.HexToHash(blockHash)
	block, err := client.BlockByHash(context.Background(), bHash)
	if err != nil {
		log.Fatal(err)
	}
	return block
}
