package cryptoUtil

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/crypto"
)

type RecvAddrData struct {
	*AddrData
	Msg int
}

func NewPrivateKey() *ecdsa.PrivateKey {
	curve := crypto.S256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func newRecvAddrData() *RecvAddrData {
	privateKey := NewPrivateKey()
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	return &RecvAddrData{
		AddrData: &AddrData{
			PublicKey: publicKey,
			Address:   crypto.PubkeyToAddress(*publicKey),
		},
	}
}

func InitRecvAddrData(psk *ecdsa.PrivateKey, n int) *RecvAddrData {
	recv := newRecvAddrData()
	recv.calcMsg(psk, n)
	return recv
}

func (recv *RecvAddrData) calcMsg(psk *ecdsa.PrivateKey, n int) {
	// 通过哈希计算 地址集合addrList
	// n为每次发送消息的位数
	data := []byte(recv.Address.Hex() + string(psk.D.Bytes()))
	hasher := sha256.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)
	// 获取最后n位
	bytesNum := n/8 + 1
	lastByte := hashBytes[len(hashBytes)-bytesNum]
	mask := uint64(1<<n - 1)
	result := uint64(lastByte) & mask
	recv.Msg = int(result)
}
