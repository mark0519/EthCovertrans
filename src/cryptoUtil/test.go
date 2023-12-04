package cryptoUtil

import (
	"crypto/ecdsa"
	"fmt"
)

func TestAddr(psk *ecdsa.PrivateKey) {
	rt := InitRecvAddrData(psk, 4)
	fmt.Println("============RecvAddrData==============")
	fmt.Println("publicKey :", rt.PublicKey)
	fmt.Println("addr      :", rt.Address)
	fmt.Println("msg       :", rt.Msg)

	sk1 := NewPrivateKey()
	pk1 := sk1.Public().(*ecdsa.PublicKey)

	pk2_frompk1 := derivationPublicKey(pk1, psk)

	fmt.Printf("pk2_frompk1:%v\n", pk2_frompk1)
}
