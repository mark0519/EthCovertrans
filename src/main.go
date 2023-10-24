package main

import (
	"EthCovertrans/src/allcrypto"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type sharedData struct {
	pubKeyStr string
	psk       string
	n         int
}

func main() {
	//client := initEthClient()
	//testGetBalance(client) // 检查和ETH网关（gateway）的连接

	sharedData := sharedData{
		pubKeyStr: "04023b1d8cfbdfe2c5fb8cf1623bb5766c57c8458bad2f6ab2f4cecd70ad682b9de61c22438917d9ee2059f84b604b982ed375b596f3940461076851b60e0a191e",
		psk:       "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
		n:         2,
	}

	sender(sharedData)

}

func sender(data sharedData) {
	// Private Key: 0x93d5d04256882aaad507ff09f510969f347758109793448aa79e1b4dbe5f6efa
	// Public Key : 0x04023b1d8cfbdfe2c5fb8cf1623bb5766c57c8458bad2f6ab2f4cecd70ad682b9de61c22438917d9ee2059f84b604b982ed375b596f3940461076851b60e0a191e
	// Address    : 0xa4528e245F87CBA1D650403d196eF505EE4D0a2B
	PrivateKeyStr := "93d5d04256882aaad507ff09f510969f347758109793448aa79e1b4dbe5f6efa"
	sk, _ := hex.DecodeString(PrivateKeyStr)
	pskStr := data.psk
	psk, _ := hex.DecodeString(pskStr)
	n := data.n // 默认为2
	addrTable := allcrypto.InitAddrTable(n)
	times := 20
	// 预先计算20个地址
	key := allcrypto.SetAddrDatas(addrTable, sk, psk, n, times)
	allcrypto.ShowAddrTable(addrTable)
	fmt.Printf("Netx key: %x\n", key)
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

//func createMsgAccount(key []byte) (common.Address, ecdsa.PublicKey) {
//	// 假设你有一个 Ethereum 私钥 key,用来他生成公钥并计算钱包地址
//	//fmt.Println("私钥：" + hexutil.Encode(key))
//	privateKeyBytes := key
//	privateKey := new(ecdsa.PrivateKey)
//	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
//	privateKey.PublicKey.Curve = crypto.S256()
//	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())
//	publicKey := privateKey.PublicKey
//	// 获取公钥地址
//	address := crypto.PubkeyToAddress(privateKey.PublicKey)
//	//fmt.Println("钱包地址：" + address.Hex())
//
//	//crypto.
//
//	return address, publicKey
//}

//func testGetKey() []byte {
//	mainKey := allcrypto.GetMainKey()
//	fmt.Printf("共享主密钥(MainKey): %x\n", mainKey)
//	transHash := "0x8879c312dd900030f514b7fe44abcf64bf7c825c02b9a0a9e3bbc723bab2d004"
//	key := allcrypto.GetTransKey(mainKey, transHash)
//	fmt.Printf("通信私钥：: %x\n", key)
//	return key
//}
//
//func getKeyByToken(mainKey [32]byte, transHash string) []byte {
//	key := allcrypto.GetTransKey(mainKey, transHash)
//	return key
//}
//
//func test() {
//	// 生成一个私钥
//	//privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//
//	// 生成一个随机数作为标量 k
//	k1, err := rand.Int(rand.Reader, secp256k1.S256().Params().N)
//	if err != nil {
//		log.Fatal(err)
//	}
//	k2, err := rand.Int(rand.Reader, secp256k1.S256().Params().N)
//	if err != nil {
//		log.Fatal(err)
//	}
//	k := new(big.Int).Mul(k1, k2)
//	k = k.Mod(k, secp256k1.S256().Params().N)
//
//	// 计算 psk*k
//	x, y := secp256k1.S256().ScalarBaseMult(k1.Bytes())
//	//x, y := secp256k1.S256().ScalarMult(big.NewInt(1), big.NewInt(2), k1.Bytes())
//	x1, y1 := secp256k1.S256().ScalarMult(x, y, k2.Bytes())
//
//	x2, y2 := secp256k1.S256().ScalarBaseMult(k.Bytes())
//
//	fmt.Println("psk*sk:")
//	fmt.Println("x:", x1)
//	fmt.Println("y:", y1)
//	fmt.Println("x:", x2)
//	fmt.Println("y:", y2)
//
//}
