package cryptoUtil

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type AddrData struct {
	PublicKey *ecdsa.PublicKey
	Address   common.Address
}
