package ethio

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func TestETHIO() {
	//sender := new(ETHSender)

	// Sender 账户私钥
	//sendKey := cryptoUtil.NewPrivateKey()
	//sendKey := crypto.ToECDSAUnsafe(common.FromHex("46927aa4aef15bcb8233c953a0a62e0a53334adc27f89767cc82b2e9841a723d"))
	//sendData := cryptoUtil.InitSendAddrData(sendKey)
	//recvAddr := "0477a578618bB6E33AB017b441275d86C3E9a165"
	//addressBytes, _ := hex.DecodeString(recvAddr)
	//recvData := new(cryptoUtil.RecvAddrData)
	//recvData = &cryptoUtil.RecvAddrData{
	//	AddrData: &cryptoUtil.AddrData{
	//		PublicKey: nil,
	//		Address:   common.Address(addressBytes),
	//	},
	//}
	//
	//// ETHSender初始化
	//sender.newETHSender(sendData, recvData)
	//
	//// 发送信息，目标地址是0xa4528e245F87CBA1D650403d196eF505EE4D0a2B
	//value := new(big.Int)
	//value.SetString("10", 10)
	//txHash := sender.sendETH(value)
	//fmt.Println("[Sender] txHash:", txHash)
	//waitForTx(txHash)

	receiver := new(ETHReceiver)
	fromAddr := "0477a578618bB6E33AB017b441275d86C3E9a165"
	fmt.Println("[Receiver] Trans From Address: ", fromAddr)
	fromAddressBytes, _ := hex.DecodeString(fromAddr)
	receiver.NewETHReceiver(common.Address(fromAddressBytes))

	addr := receiver.GetLatestToAddress()
	fmt.Println("[Receiver] Trans To Address: ", addr.String())
	value := receiver.GetLatestTransValue()
	fmt.Println("[Receiver] Trans Value: ", value.String(), " Wei")

}
