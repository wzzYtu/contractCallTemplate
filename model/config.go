package model

type Conf struct {
	Chain    Chain
	Wallet   Wallet
	Contract Contract
}

type Chain struct {
	ChainURL string `toml:"chainURL"`
	ChainID  int    `toml:"chainID"`
	ChainWss string `toml:"chainWss"`
}

type Wallet struct {
	PrivStrOfDeploy string `toml:"privStrOfDeploy"`
	PrivStrOfSend   string `toml:"privStrOfSend"`
}

type Contract struct {
	Address string `toml:"address"`
}
