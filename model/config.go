package model

type Conf struct {
	Chain    Chain
	Wallet   Wallet
	Contract Contract
}

type Chain struct {
	ChainURL string `toml:"chainURL"`
	ChainID  int    `toml:"chainID"`
}

type Wallet struct {
	PrivStr string `toml:"privStr"`
}

type Contract struct {
	Address string `toml:"address"`
}
