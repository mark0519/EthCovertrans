package allcrypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type recvAddrData struct {
	addrData
}

type RecvAddrTable struct {
	table [][]recvAddrData
	n     int // 交易嵌入位数
}

func InitRecvAddrTable(n int) *RecvAddrTable {
	sum := 1 << n                        // 2的n次幂
	table := make([][]recvAddrData, sum) // 创建sum行切片
	for i := range table {
		// 为每一行创建一个切片，可以根据需要指定初始长度
		table[i] = make([]recvAddrData, 0) // 这里初始长度为 0
	}
	return &RecvAddrTable{
		table: table,
		n:     n,
	}
}

func NewPrivateKey() *ecdsa.PrivateKey {
	curve := crypto.S256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func newRecvAddrData() *recvAddrData {
	privateKey := NewPrivateKey()
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	return &recvAddrData{
		addrData: addrData{
			publicKey: *publicKey,
			address:   crypto.PubkeyToAddress(*publicKey).Hex(),
		},
	}
}

func (rt RecvAddrTable) FillRecvAddrTable(psk []byte, least int) {
	for i := 0; i < least; i++ {
		recv := newRecvAddrData()
		msg := calcMsg(recv, psk, rt.n)
		rt.table[msg] = append(rt.table[msg], *recv)
	}
}

func (rt RecvAddrTable) ShowRecvAddrTable() {
	for i := range rt.table {
		fmt.Println("============table[", i, "]==============")
		for j := range rt.table[i] {
			fmt.Printf("publicKey :0x%x\n", rt.table[i][j].publicKey)
			fmt.Println("Addr      :", rt.table[i][j].address)
			fmt.Println("---")
		}
	}
}

func getHashByAddrPsk(addr string, psk string) []byte {
	// 通过钱包地址计算交易哈希
	data := []byte(addr + psk)
	// 计算data的SHA256
	hasher := sha256.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)
	return hashBytes
}

func calcMsg(recv *recvAddrData, psk []byte, n int) int {
	// 通过哈希计算 地址集合addrList
	hash := getHashByAddrPsk(recv.address, string(psk))
	var result byte
	result = (1 << n) - 1
	lowestTwoBits := hash[15] & result
	return int(lowestTwoBits)
}
