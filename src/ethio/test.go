package ethio

import (
	"EthCovertrans/src/allcrypto"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func TestETHIO() {
	io := new(ETHSender)

	// Sender 账户私钥
	sendKey := allcrypto.NewPrivateKey()
	//sendKey := crypto.ToECDSAUnsafe(common.FromHex("46927aa4aef15bcb8233c953a0a62e0a53334adc27f89767cc82b2e9841a723d"))
	sendData := allcrypto.InitSendAddrData(sendKey)
	recvAddr := "a4528e245F87CBA1D650403d196eF505EE4D0a2B"
	addressBytes, _ := hex.DecodeString(recvAddr)
	recvData := new(allcrypto.RecvAddrData)
	recvData = &allcrypto.RecvAddrData{
		AddrData: &allcrypto.AddrData{
			PublicKey: nil,
			Address:   common.Address(addressBytes),
		},
	}

	// ETHSender初始化
	io.newETHSender(sendData, recvData)

	// 发送信息，目标地址是0xa4528e245F87CBA1D650403d196eF505EE4D0a2B
	value := new(big.Int)
	value.SetString("10", 10)
	txHash := io.sendETH(value)
	fmt.Println("[Sender] txHash:", txHash)
	waitForTx(txHash)
	//addressBytes2, _ := hex.DecodeString("c8D69B351aCD4508EFa0F67c163e864F76D0A522")
	//GetReceiverBySenderAddr(common.Address(addressBytes2))

}
