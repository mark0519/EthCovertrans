package main

import (
	"EthCovertrans/src/allcrypto"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {
	pskInt, _ := rand.Int(rand.Reader, secp256k1.S256().Params().N)
	psk := pskInt.Bytes()
	allcrypto.TestAddr(psk)
}

//func initEthClient() *ethclient.Client {
//	client, err := ethclient.Dial("https://sut0ne.tk/v1/sepolia")
//	//client, err := ethclient.Dial("https://cloudflare-eth.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	return client
//}
//
//func testGetBalance(c *ethclient.Client) {
//	addr := "0xa4528e245F87CBA1D650403d196eF505EE4D0a2B"
//	account := common.HexToAddress(addr)
//	balance, err := c.BalanceAt(context.Background(), account, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Account: %s\n", account)
//	fmt.Printf("Balance: %d wei\n", balance)
//}
