package service

//
//import (
//	"context"
//	"contractCallTemplate/conf"
//	initialize "contractCallTemplate/init"
//	"fmt"
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/crypto/sha3"
//	"log"
//	"math/big"
//)
//
//func SendERC20(fromAddress, toAddress, tokenAddr string, value float64, precision int64) {
//	var client = initialize.NewClient(conf.Conf.Chain.ChainURL)
//	fromAccount := common.HexToAddress(fromAddress)
//	privateKey, err := crypto.HexToECDSA(conf.Conf.Wallet.PrivStrOfSend)
//	if err != nil {
//		log.Fatal("获取私钥失败", err)
//	}
//	toAccount := common.HexToAddress(toAddress)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAccount)
//	if err != nil {
//		log.Fatal("获取nonce值失败", err)
//	}
//	signer := types.LatestSignerForChainID(new(big.Int).SetUint64(uint64(conf.Conf.Chain.ChainID)))
//	valueETH := big.NewInt(0)
//
//	tokenAddress := common.HexToAddress(tokenAddr)
//
//	transferFnSignature := []byte("transfer(address,uint256)")
//	hash := sha3.NewKeccak256()
//	hash.Write(transferFnSignature)
//	methodID := hash.Sum(nil)[:4]
//	fmt.Println(hexutil.Encode(methodID))
//
//	paddedAddress := common.LeftPadBytes(toAccount.Bytes(), 32)
//
//	values := precisionfunc(value, precision)
//	bigstr := values.String()
//	paddedAmount := []byte(bigstr)
//	var data []byte
//	data = append(data, methodID...)
//	data = append(data, paddedAddress...)
//	data = append(data, paddedAmount...)
//
//	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
//		To:   &toAccount,
//		Data: data,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal("获取gas price失败", err)
//	}
//
//	oTx := types.NewTx(&types.LegacyTx {
//		Nonce:    nonce,
//		To:       &toAccount,
//		Value:    valueETH,
//		Gas:      gasLimit,
//		GasPrice: gasPrice,
//		Data:     data,
//	})
//	signature, err := crypto.Sign(signer.Hash(oTx).Bytes(), privateKey)
//	if err != nil {
//		initialize.ExitWithErr(err)
//	}
//	txWithSig, err := oTx.WithSignature(signer, signature)
//	if err != nil {
//		initialize.ExitWithErr(err)
//	}
//	err = client.SendTransaction(context.Background(), txWithSig)
//	if err != nil {
//		initialize.ExitWithErr(err)
//	}
//}
//
//func precisionfunc(val float64, precision int64) *big.Int {
//	bigval := new(big.Float)
//	bigval.SetFloat64(val)
//	coin := new(big.Float)
//	coin.SetInt(big.NewInt(precision))
//	bigval.Mul(bigval, coin)
//	result := new(big.Int)
//	bigval.Int(result)
//	return result
//}
