package util

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type AddrData struct {
	PublicKey *ecdsa.PublicKey
	Address   common.Address
}

type RecvAddrData struct {
	*AddrData
	Msg int32
}

type SendAddrData struct {
	*AddrData
	privateKey *ecdsa.PrivateKey
}

func InitSendAddrData(sk *ecdsa.PrivateKey) *SendAddrData {
	ad := PrivateKeyToAddrData(sk)
	return &SendAddrData{
		AddrData:   ad,
		privateKey: sk,
	}
}

func (sad *SendAddrData) GetSendAddrDataPrivateKey() *ecdsa.PrivateKey {
	return sad.privateKey
}

func DerivationSendAddrData(oldKey *SendAddrData, psk *ecdsa.PrivateKey) *SendAddrData {

	newKeyInt := new(big.Int).Mul(oldKey.privateKey.D, psk.D)
	newKeyInt = newKeyInt.Mod(newKeyInt, secp256k1.S256().Params().N)
	newX, newY := secp256k1.S256().ScalarBaseMult(newKeyInt.Bytes())
	curve := secp256k1.S256()
	newPublicKey := ecdsa.PublicKey{
		Curve: curve,
		X:     newX,
		Y:     newY,
	}
	newKey := SendAddrData{
		AddrData: &AddrData{
			PublicKey: &newPublicKey,
			Address:   crypto.PubkeyToAddress(newPublicKey),
		},
		privateKey: &ecdsa.PrivateKey{
			PublicKey: newPublicKey,
			D:         newKeyInt,
		},
	}
	return &newKey
}

func PrivateKeyToAddrData(sk *ecdsa.PrivateKey) *AddrData {
	pk := sk.Public().(*ecdsa.PublicKey)
	return &AddrData{
		PublicKey: pk,
		Address:   crypto.PubkeyToAddress(*pk),
	}
}

func DerivationPublicKey(oldKey *ecdsa.PublicKey, psk *ecdsa.PrivateKey) *ecdsa.PublicKey {
	newX, newY := secp256k1.S256().ScalarMult(oldKey.X, oldKey.Y, psk.D.Bytes())
	newKey := ecdsa.PublicKey{
		Curve: oldKey.Curve,
		X:     newX,
		Y:     newY,
	}
	return &newKey
}

func NewPrivateKey() *ecdsa.PrivateKey {
	curve := secp256k1.S256()
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
	recv.Msg = CalcMsg(recv.Address, psk, n)
	return recv
}

func CalcMsg(addr common.Address, psk *ecdsa.PrivateKey, n int) (msg int32) {
	// 通过哈希计算 地址集合addrList
	// n为每次发送消息的位数
	data := []byte(addr.Hex() + string(psk.D.Bytes()))
	hasher := sha256.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)
	// 获取最后n位
	bytesNum := n/8 + 1
	lastByte := hashBytes[len(hashBytes)-bytesNum]
	mask := int32(1<<n - 1)
	msg = int32(lastByte) & mask
	return
}
