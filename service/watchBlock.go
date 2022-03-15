package service

import (
	"context"
	"contractCallTemplate/conf"
	initialize "contractCallTemplate/init"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
)

func WatchBlock() {
	client := initialize.NewClient(conf.Conf.Chain.ChainWss)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("block hash:", block.Hash().Hex())                // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println("block number:", block.Number().Uint64())         // 3477413
			fmt.Println("block time:", block.Time())                      // 1529525947
			fmt.Println("block nonce:", block.Nonce())                    // 130524141876765836
			fmt.Println("transaction number:", len(block.Transactions())) // 7
		}
	}
}
