package cryptoUtil

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

func derivationPublicKey(oldKey *ecdsa.PublicKey, psk *ecdsa.PrivateKey) *ecdsa.PublicKey {
	newX, newY := secp256k1.S256().ScalarMult(oldKey.X, oldKey.Y, psk.D.Bytes())
	newKey := ecdsa.PublicKey{
		Curve: oldKey.Curve,
		X:     newX,
		Y:     newY,
	}
	return &newKey
}
