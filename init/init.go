package initialize

import (
	"context"
	"contractCallTemplate/conf"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChainID uint64

var (
	privStr        = conf.Conf.Wallet.PrivStrOfDeploy
	PlatON_Dev_Url = conf.Conf.Chain.ChainURL
	chainID        = ChainID(conf.Conf.Chain.ChainID)
)

var (
	privilege          = common.HexToAddress("0x48bB7866a811Be354F49f6E287d2bF64F436C992")
	tokenTransferProxy = common.HexToAddress("0x80ce20967f303fad6d6cd819958a54db1a8d0a1c")
	token              = common.HexToAddress("0xa647253a501123792e3982968bc5232e0af95e32")
)

func ExitWithErr(err error) {
	fmt.Println(err)
	os.Exit(0)
}

func NewClient(url string) *ethclient.Client {
	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		ExitWithErr(err)
	}
	return client
}

func InitAuth(chainID ChainID) *bind.TransactOpts {
	var priv ecdsa.PrivateKey
	priv.D, _ = new(big.Int).SetString(privStr, 16)
	priv.Curve = crypto.S256()
	priv.X, priv.Y = priv.ScalarBaseMult(priv.D.Bytes())
	auth, err := bind.NewKeyedTransactorWithChainID(&priv, big.NewInt(int64(chainID)))
	if err != nil {
		ExitWithErr(err)
	}
	return auth
}

func GetAccount() (*ecdsa.PrivateKey, common.Address) {
	var priv ecdsa.PrivateKey
	priv.D, _ = new(big.Int).SetString(privStr, 16)
	priv.Curve = crypto.S256()
	priv.X, priv.Y = priv.ScalarBaseMult(priv.D.Bytes())
	return &priv, crypto.PubkeyToAddress(priv.PublicKey)
}

func NewAuth(chainID ChainID, priv *ecdsa.PrivateKey) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(priv, big.NewInt(int64(chainID)))
	if err != nil {
		ExitWithErr(err)
	}
	return auth
}
