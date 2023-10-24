package allcrypto

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type sendAddrData struct {
	addrData
	privateKey ecdsa.PrivateKey
}

type sendAddrList []sendAddrData

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

func privateKeyToAddrData(sk *ecdsa.PrivateKey) *addrData {
	pk := sk.Public().(*ecdsa.PublicKey)
	return &addrData{
		publicKey: *pk,
		address:   crypto.PubkeyToAddress(*pk).Hex(),
	}
}

func initSendAddrList(n int, psk []byte) *sendAddrList {
	sl := make([]sendAddrData, n)
	sk := newPrivateKey()
	sl[0] = sendAddrData{
		addrData:   *privateKeyToAddrData(sk),
		privateKey: *sk,
	}
	for i := 1; i < n; i++ {
		sk = derivationPrivateKey(sk, psk)
		sl[i] = sendAddrData{
			addrData:   *privateKeyToAddrData(sk),
			privateKey: *sk,
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
