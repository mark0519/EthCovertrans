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
	sendData := allcrypto.InitSendAddrData(sendKey)

	recvAddr := "C5118e577B22d793A618286655505F8beB21a2DC"
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
	value.SetString("1000", 10)
	txHash := io.sendETH(value)
	fmt.Println("[Sender] txHash:", txHash)
	GetReceiverBySenderAddr(sendData.Address)

}
