package allcrypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
)

type RecvAddrData struct {
	*AddrData
}

type recvAddrTable struct {
	table [][]RecvAddrData
	n     int // 交易嵌入位数
}

func initRecvAddrTable(n int) *recvAddrTable {
	sum := 1 << n                        // 2的n次幂
	table := make([][]RecvAddrData, sum) // 创建sum行切片
	for i := range table {
		// 为每一行创建一个切片，可以根据需要指定初始长度
		table[i] = make([]RecvAddrData, 0) // 这里初始长度为 0
	}
	return &recvAddrTable{
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

func NewRecvAddrData() *RecvAddrData {
	privateKey := NewPrivateKey()
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	return &RecvAddrData{
		AddrData: &AddrData{
			PublicKey: publicKey,
			Address:   crypto.PubkeyToAddress(*publicKey),
		},
	}
}

func (rt recvAddrTable) fillRecvAddrTable(psk []byte, least int) {
	for i := 0; i < least; i++ {
		recv := NewRecvAddrData()
		msg := calcMsg(recv, psk, rt.n)
		rt.table[msg] = append(rt.table[msg], *recv)
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

func calcMsg(recv *RecvAddrData, psk []byte, n int) int {
	// 通过哈希计算 地址集合addrList
	hash := getHashByAddrPsk(recv.Address.Hex(), string(psk))
	var result byte
	result = (1 << n) - 1
	lowestTwoBits := hash[15] & result
	return int(lowestTwoBits)
}
