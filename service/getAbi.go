package service

import (
	"fmt"
	"github.com/nanmu42/etherscan-api"
)

func GetContractABI(addr string) (string, error) {
	client := etherscan.New(etherscan.Mainnet, "4KUCSHA1KJA4J7PNBNZY2A5RQATEUIUJYP")
	contractAbi, err := client.ContractABI(addr)
	if err != nil {
		fmt.Println("get contract ABI fail:", err)
		return "", err
	}
	return contractAbi, nil
}
