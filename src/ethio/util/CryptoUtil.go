package util

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
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
	// 私钥转SendAddrData
	ad := PrivateKeyToAddrData(sk)
	return &SendAddrData{
		AddrData:   ad,
		privateKey: sk,
	}
}

func (sad *SendAddrData) GetSendAddrDataPrivateKey() *ecdsa.PrivateKey {
	// 获取SendAddrData中的私钥
	return sad.privateKey
}

func DerivationSendAddrData(oldKey *SendAddrData, psk *ecdsa.PrivateKey) *SendAddrData {
	// SendAddrData派生
	newKeyInt := new(big.Int).Mul(oldKey.privateKey.D, psk.D)
	newKeyInt = newKeyInt.Mod(newKeyInt, crypto.S256().Params().N)
	newX, newY := crypto.S256().ScalarBaseMult(newKeyInt.Bytes())
	curve := crypto.S256()
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
	// 私钥转AddrData
	pk := sk.Public().(*ecdsa.PublicKey)
	return &AddrData{
		PublicKey: pk,
		Address:   crypto.PubkeyToAddress(*pk),
	}
}

func DerivationPublicKey(oldKey *ecdsa.PublicKey, psk *ecdsa.PrivateKey) *ecdsa.PublicKey {
	// 公钥派生
	newX, newY := crypto.S256().ScalarMult(oldKey.X, oldKey.Y, psk.D.Bytes())
	newKey := ecdsa.PublicKey{
		Curve: oldKey.Curve,
		X:     newX,
		Y:     newY,
	}
	return &newKey
}

func NewPrivateKey() *ecdsa.PrivateKey {
	// 新建随机私钥
	curve := crypto.S256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func newRecvAddrData() *RecvAddrData {
	// 新建随机RecvAddrData
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
	// 初始化随机RecvAddrData，n为每次发送消息比特数
	recv := newRecvAddrData()
	recv.Msg = CalcMsg(recv.Address, psk, n)
	return recv
}

func CalcMsg(addr common.Address, psk *ecdsa.PrivateKey, n int) (msg int32) {
	// 计算随机出的RecvAddrData对应的数据，用于与交易金额异或
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

func D2PrivateKey(D *big.Int) *ecdsa.PrivateKey {
	X, Y := crypto.S256().ScalarBaseMult(D.Bytes())
	return &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			X:     X,
			Y:     Y,
			Curve: crypto.S256(),
		},
		D: D,
	}
}

func XY2PublicKey(X, Y *big.Int) *ecdsa.PublicKey {
	return &ecdsa.PublicKey{
		Curve: crypto.S256(),
		X:     X,
		Y:     Y,
	}
}

func SignMessage(privateKey *ecdsa.PrivateKey, message []byte) (uint8, [32]byte, [32]byte) {
	// 计算哈希值
	hash := sha3.NewLegacyKeccak256()
	hash.Write(message)
	hashed := hash.Sum(nil)

	// 签名
	sig, err := crypto.Sign(hashed, privateKey)
	r := [32]byte(sig[:32])
	s := [32]byte(sig[32:64])
	v := sig[64] + 27

	if err != nil {
		log.Fatal(err)
	}
	return v, r, s
}
