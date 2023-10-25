package main

import (
	"EthCovertrans/src/ethio"
)

func main() {
	//pskInt, _ := rand.Int(rand.Reader, secp256k1.S256().Params().N)
	//psk := pskInt.Bytes()
	//allcrypto.TestAddr(psk)
	ethio.TestETHIO()

}

//func initEthClient() *ethclient.Client {
//	//client, err := ethclient.Dial("https://sut0ne.tk/v1/sepolia")
//	client, err := ethclient.Dial(ethio.EthGateway)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return client
//}
//
//func testGetBalance(c *ethclient.Client) {
//	addr := "0x0477a578618bB6E33AB017b441275d86C3E9a165"
//	account := common.HexToAddress(addr)
//	balance, err := c.BalanceAt(context.Background(), account, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Account: %s\n", account)
//	fmt.Printf("Balance: %d wei\n", balance)
//}
