package allcrypto

import "crypto/ecdsa"

type addrData struct {
	publicKey ecdsa.PublicKey
	address   string
}
