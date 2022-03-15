package service

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"strings"
)

func UnpackTxInput(input, contractAddress string) (string, []interface{}, error) {
	// load contract ABI
	contractAbi, err := GetContractABI(contractAddress)
	if err != nil {

	}
	abi, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		log.Fatal(err)
		return "", nil, err
	}

	// decode txInput method signature
	decodedSig, err := hex.DecodeString(input[2:10])
	if err != nil {
		log.Fatal(err)
		return "", nil, err
	}

	// recover Method from signature and ABI
	method, err := abi.MethodById(decodedSig)
	if err != nil {
		log.Fatal("The recovery method from the signature and ABI failed：", err)
		return "", nil, err
	}
	decodedData, err := hex.DecodeString(input[10:])
	if err != nil {
		log.Fatal(err)
	}
	inputData, err := method.Inputs.Unpack(decodedData)
	if err != nil {
		log.Fatal(err)
		return "", nil, err
	}
	result := method.String()
	return result, inputData, nil
}
