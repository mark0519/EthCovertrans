package allcrypto

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type SendAddrData struct {
	*AddrData
	privateKey *ecdsa.PrivateKey
}

type sendAddrList []SendAddrData

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

func derivationPrivateKey(oldKey *ecdsa.PrivateKey, psk []byte) *ecdsa.PrivateKey {
	var pskInt big.Int
	pskInt.SetBytes(psk)

	newKeyInt := new(big.Int).Mul(oldKey.D, &pskInt)
	newKeyInt = newKeyInt.Mod(newKeyInt, secp256k1.S256().Params().N)
	newX, newY := secp256k1.S256().ScalarBaseMult(newKeyInt.Bytes())
	curve := secp256k1.S256()
	newKey := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     newX,
			Y:     newY,
		},
		D: newKeyInt,
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

func InitSendAddrList(len int, psk []byte) *sendAddrList {
	sl := make([]SendAddrData, len)
	sk := NewPrivateKey()
	sl[0] = SendAddrData{
		AddrData:   PrivateKeyToAddrData(sk),
		privateKey: sk,
	}
	for i := 1; i < len; i++ {
		sk = derivationPrivateKey(sk, psk)
		sl[i] = SendAddrData{
			AddrData:   PrivateKeyToAddrData(sk),
			privateKey: sk,
		}
	}
	return (*sendAddrList)(&sl)
}

func derivationPublicKey(oldKey *ecdsa.PublicKey, psk []byte) *ecdsa.PublicKey {
	newX, newY := secp256k1.S256().ScalarMult(oldKey.X, oldKey.Y, psk)
	newKey := ecdsa.PublicKey{
		Curve: oldKey.Curve,
		X:     newX,
		Y:     newY,
	}
	return &newKey
}
