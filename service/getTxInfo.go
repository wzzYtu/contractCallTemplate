package service

import (
	"context"
	"contractCallTemplate/conf"
	"contractCallTemplate/init"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func GetAllTxByBlockHash(blockHash string) types.Transactions {
	blockInfo := GetBlockByHash(blockHash)
	return rangeBlock(blockInfo)
}

func GetAllTxByBlockNum(blockNum *big.Int) types.Transactions {
	blockInfo := GetBlockInfoByNum(blockNum)
	return rangeBlock(blockInfo)
}

func rangeBlock(block *types.Block) types.Transactions {
	var transactions types.Transactions
	for _, TxInfo := range block.Transactions() {
		transactions = append(transactions, TxInfo)
	}
	return transactions
}

func GetTxByTxHash(txHash string) (bool, *types.Transaction) {
	var client = initialize.NewClient(conf.Conf.Chain.ChainURL)
	Hash := common.HexToHash(txHash)
	tx, isPending, err := client.TransactionByHash(context.Background(), Hash)
	if err != nil {
		log.Fatal(err)
	}
	return isPending, tx
}
