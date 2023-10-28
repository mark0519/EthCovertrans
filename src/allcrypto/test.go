package allcrypto

import (
	"crypto/ecdsa"
	"fmt"
)

func TestAddr(psk []byte) {
	rt := InitRecvAddrData(psk, 4)
	fmt.Println("============RecvAddrData==============")
	fmt.Println("publicKey :", rt.PublicKey)
	fmt.Println("addr      :", rt.Address)
	fmt.Println("msg       :", rt.Msg)

	sk1 := NewPrivateKey()
	sk2 := derivationPrivateKey(sk1, psk)
	pk1 := sk1.Public().(*ecdsa.PublicKey)

	pk2_fromsk2 := sk2.Public()
	pk2_frompk1 := derivationPublicKey(pk1, psk)

	fmt.Printf("pk2_fromsk2:%v\n", pk2_fromsk2)
	fmt.Printf("pk2_frompk1:%v\n", pk2_frompk1)

	sl := InitSendAddrList(20, psk)
	fmt.Println("============SendAddrList==============")
	for i := range *sl {
		fmt.Println("publicKey :", (*sl)[i].PublicKey)
		fmt.Println("addr      :", (*sl)[i].Address)
		fmt.Println("privateKey:", (*sl)[i].privateKey)
	}
}
