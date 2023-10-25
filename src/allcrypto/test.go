package allcrypto

import (
	"crypto/ecdsa"
	"fmt"
)

func (rt recvAddrTable) showRecvAddrTable() {
	for i := range rt.table {
		fmt.Println("============table[", i, "]==============")
		for j := range rt.table[i] {
			fmt.Printf("publicKey :0x%x\n", rt.table[i][j].publicKey)
			fmt.Println("Addr      :", rt.table[i][j].address)
			fmt.Println("---")
		}
	}
}

func TestAddr(psk []byte) {
	rt := initRecvAddrTable(2)
	rt.fillRecvAddrTable(psk, 20)
	rt.showRecvAddrTable()

	sk1 := newPrivateKey()
	sk2 := derivationPrivateKey(sk1, psk)
	pk1 := sk1.Public().(*ecdsa.PublicKey)

	pk2_fromsk2 := sk2.Public()
	pk2_frompk1 := derivationPublicKey(pk1, psk)

	fmt.Printf("pk2_fromsk2:%v\n", pk2_fromsk2)
	fmt.Printf("pk2_frompk1:%v\n", pk2_frompk1)

	sl := initSendAddrList(20, psk)
	fmt.Println("============SendAddrList==============")
	for i := range *sl {
		fmt.Println("publicKey :", (*sl)[i].publicKey)
		fmt.Println("addr      :", (*sl)[i].address)
		fmt.Println("privateKey:", (*sl)[i].privateKey)
	}
}
