package service

import (
	"contractCallTemplate/conf"
	initialize "contractCallTemplate/init"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

type newContractLoad func(common.Address, bind.ContractBackend) (interface{}, error)

func ContractLoad(contractAddr string, load newContractLoad) interface{} {
	client := initialize.NewClient(conf.Conf.Chain.ChainURL)
	address := common.HexToAddress(contractAddr)
	instance, err := load(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	fmt.Println("contract is:", instance)
	return instance
}
