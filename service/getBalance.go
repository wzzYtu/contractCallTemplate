package service

import (
	"context"
	"log"
	"math"
	"math/big"

	"contractCallTemplate/conf"
	"contractCallTemplate/init"
	"github.com/ethereum/go-ethereum/common"
)

// GetBalance 得到地址address中ETH余额
func GetBalance(address string) *big.Float {
	var client = init.NewClient(conf.Conf.Chain.ChainURL)

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return weiToEth(balance)
}

// GetBalanceOfBlockNumber 得到指定区块高度下地址address中ETH余额
func GetBalanceOfBlockNumber(address string, blockNumber int64) *big.Float {
	var client = init.NewClient(conf.Conf.Chain.ChainURL)

	blockNum := big.NewInt(blockNumber)
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, blockNum)
	if err != nil {
		log.Fatal(err)
	}
	return weiToEth(balance)
}

// GetBalanceOfPending 得到待处理的账户余额
func GetBalanceOfPending(address string) *big.Float {
	var client = init.NewClient(conf.Conf.Chain.ChainURL)
	account := common.HexToAddress(address)
	balance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	return weiToEth(balance)
}

func weiToEth(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}
