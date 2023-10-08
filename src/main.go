package main

import (
	"EthCovertrans/src/allcrypto"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client := initEthClient()
	testGetBalance(client) // 检查和ETH网关（gateway）的连接
	//key := testGetKey()
	//addr, _ := createMsgAccount(key)
	//fmt.Println("钱包地址：" + addr.Hex())

}

func testGetKey() []byte {
	mainKey := allcrypto.GetMainKey()
	fmt.Printf("共享主密钥(MainKey): %x\n", mainKey)
	transHash := "0x8879c312dd900030f514b7fe44abcf64bf7c825c02b9a0a9e3bbc723bab2d004"
	key := allcrypto.GetTransKey(mainKey, transHash)
	fmt.Printf("通信私钥：: %x\n", key)
	return key
}

func getKeyByToken(mainKey [32]byte, transHash string) []byte {
	key := allcrypto.GetTransKey(mainKey, transHash)
	return key
}

func initEthClient() *ethclient.Client {
	client, err := ethclient.Dial("https://sut0ne.tk/v1/sepolia")
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func testGetBalance(c *ethclient.Client) {
	addr := "0xa4528e245F87CBA1D650403d196eF505EE4D0a2B"
	account := common.HexToAddress(addr)
	balance, err := c.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account: %s\n", account)
	fmt.Printf("Balance: %d wei\n", balance)
}

func createMsgAccount(key []byte) (common.Address, ecdsa.PublicKey) {
	// 假设你有一个 Ethereum 私钥 key,用来他生成公钥并计算钱包地址
	//fmt.Println("私钥：" + hexutil.Encode(key))
	privateKeyBytes := key
	privateKey := new(ecdsa.PrivateKey)
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.Curve = crypto.S256()
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())
	publicKey := privateKey.PublicKey
	// 获取公钥地址
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	//fmt.Println("钱包地址：" + address.Hex())
	return address, publicKey
}
