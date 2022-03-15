package service

import (
	"context"
	"contractCallTemplate/conf"
	"contractCallTemplate/init"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func SendETH(fromAddress, toAddress string, value float64) {
	var client = initialize.NewClient(conf.Conf.Chain.ChainURL)
	fromAccount := common.HexToAddress(fromAddress)
	privateKey, err := crypto.HexToECDSA(conf.Conf.Wallet.PrivStrOfSend)
	if err != nil {
		log.Fatal("获取私钥失败", err)
	}
	toAccount := common.HexToAddress(toAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAccount)
	if err != nil {
		log.Fatal("获取nonce值失败", err)
	}
	signer := types.LatestSignerForChainID(new(big.Int).SetUint64(uint64(conf.Conf.Chain.ChainID)))
	values := weiToETH(value)
	fmt.Println("values is:", values)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取gas price失败", err)
	}

	oTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAccount,
		Value:    values,
		Gas:      21000,
		GasPrice: gasPrice,
	})
	signature, err := crypto.Sign(signer.Hash(oTx).Bytes(), privateKey)
	if err != nil {
		initialize.ExitWithErr(err)
	}
	txWithSig, err := oTx.WithSignature(signer, signature)
	if err != nil {
		initialize.ExitWithErr(err)
	}
	err = client.SendTransaction(context.Background(), txWithSig)
	if err != nil {
		initialize.ExitWithErr(err)
	}

	fmt.Printf("tx sent: %s", txWithSig.Hash().Hex())
}

func weiToETH(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))
	bigval.Mul(bigval, coin)
	result := new(big.Int)
	bigval.Int(result)
	return result
}
