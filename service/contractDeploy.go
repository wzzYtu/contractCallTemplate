package service

import (
	"context"
	"contractCallTemplate/conf"
	"contractCallTemplate/init"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

type DeployFunc func(*bind.TransactOpts, bind.ContractBackend, ...interface{}) (common.Address, *types.Transaction, interface{}, error)

// DeployContract 在以太坊上部署合约
// @param
func DeployContract(deployFunc DeployFunc, initPara ...interface{}) common.Address {
	var (
		client = initialize.NewClient(conf.Conf.Chain.ChainURL)
		auth   = initialize.InitAuth(initialize.ChainID(conf.Conf.Chain.ChainID))
	)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasPriceStr := gasPrice.String()
	auth.GasPrice, _ = new(big.Int).SetString(gasPriceStr, 10)
	contractAddr, _, _, err := deployFunc(auth, client, initPara)
	if err != nil {
		initialize.ExitWithErr(err)
	}
	fmt.Println("------ Contract Address ------")
	fmt.Println("contractAddr", contractAddr)
	return contractAddr
}
