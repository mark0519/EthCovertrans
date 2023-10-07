package main

import (
	"EthCovertrans/src/allcrypto"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//client, err := ethclient.Dial("https://sutOne.tk/v1/goerli")
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	account := common.HexToAddress("0xa4528e245F87CBA1D650403d196eF505EE4D0a2B")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

}

func getKey() [32]byte {
	mainKey := allcrypto.GetMainKey()
	fmt.Printf("共享主密钥(MainKey): %x\n", mainKey)
	transHash := "0x8879c312dd900030f514b7fe44abcf64bf7c825c02b9a0a9e3bbc723bab2d004"
	key := allcrypto.GetTransKey(mainKey, transHash)
	return key
}
