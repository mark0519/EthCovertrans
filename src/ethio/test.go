package ethio

import (
	"EthCovertrans/src/ethio/util"
	"crypto/ecdsa"
	"fmt"
	"math/big"
)

func TestMsgSenderFactory(msgstr string, psk *ecdsa.PrivateKey, orignSenderSK *ecdsa.PrivateKey) *ecdsa.PublicKey {
	// 创建ETHSender实例

	msgIntSlice := SliceMsg(msgstr)
	var times = len(msgIntSlice)
	orignSender := util.InitSendAddrData(orignSenderSK)
	senders := *NewSenderList(times, psk, orignSender)
	recvers := *NewRecverList(times, psk)

	msgSenders := make([]MsgSender, times)
	for i := 0; i < times; i++ {
		msgSenders[i] = *NewMsgSender(&(senders)[i], &(recvers)[i], big.NewInt(int64(recvers[i].Msg^msgIntSlice[i])))
	}

	DoSend(&msgSenders)
	defer UpdateContract(senders[times].PublicKey)

	return senders[times].PublicKey
}

func Test() {
	//// 首次使用时初始化
	//psk := util.GetPskFromFile()
	//senderSK := util.NewPrivateKey()
	//sender := util.InitSendAddrData(senderSK)
	//var recvers []*ecdsa.PublicKey
	//recvers = append(recvers, sender.PublicKey)
	//
	//// 创建KeyData实例
	//keyData := util.KeyFileData{
	//	Psk:     psk,
	//	Sender:  senderSK,
	//	Recvers: &recvers, // 公钥列表
	//}
	//
	//util.EncryptKeyFileData(keyData, "ethCoverTrans.key")      // 加密并保存
	//DecKeyData := util.DecryptKeyFileData("ethCoverTrans.key") // 读取并解密
	//
	//fmt.Printf("psk1: %v\n", keyData.Psk)
	//fmt.Printf("psk2: %v\n", DecKeyData.Psk)
	//fmt.Printf("Sender1: %v\n", keyData.Sender)
	//fmt.Printf("Sender2: %v\n", DecKeyData.Sender)
	//fmt.Printf("Recvers1: %v\n", keyData.Recvers)
	//fmt.Printf("Recvers2: %v\n", DecKeyData.Recvers)

	//psk := util.NewPrivateKey()
	//senderSK := util.NewPrivateKey()
	//sender := util.InitSendAddrData(senderSK)
	//newPK := TestMsgSenderFactory("hello world!!", psk, sender.GetSendAddrDataPrivateKey())
	//msg := MsgRecverFactory(psk, sender.PublicKey, newPK)
	//print("[+] msg:", msg)

	//addGroupPK4owner(KeyData.Psk, util.PrivateKeyToAddrData(KeyData.Sender).PublicKey)
	//RegisterRecv(KeyData.Psk, util.PrivateKeyToAddrData(KeyData.Sender).PublicKey)
	recvs := GetRecvers(KeyData.Psk)
	for i, data := range *recvs {
		fmt.Printf("Recver%d: %v\n", i, data)
	}

}

func TestFileIO() {
	// 首次使用时初始化
	psk := util.NewPrivateKey()
	//psk := util.GenerateKeyFile("psk.key")
	senderSK := util.NewPrivateKey()
	sender := util.InitSendAddrData(senderSK)
	var recvers []*ecdsa.PublicKey
	recvers = append(recvers, sender.PublicKey)

	// 创建KeyData实例
	keyData := util.KeyFileData{
		Psk:     psk,
		Sender:  senderSK,
		Recvers: &recvers, // 公钥列表
	}
	keyDataF := util.KeyFileDataAndFaucet{
		KeyFileData: keyData,
		Faucet:      "983ec812c710bd1a3ef13bfd089cf8c7cf672f8bf17a7b9be51318c8314120aa",
	}

	util.EncryptKeyFileData(keyDataF, "ethCoverTrans.key")      // 加密并保存
	DecKeyDataF := util.DecryptKeyFileData("ethCoverTrans.key") // 读取并解密
	DecKeyData := DecKeyDataF.KeyFileData
	fmt.Printf("psk1: %v\n", keyData.Psk)
	fmt.Printf("psk2: %v\n", DecKeyData.Psk)
	fmt.Printf("Sender1: %v\n", keyData.Sender)
	fmt.Printf("Sender2: %v\n", DecKeyData.Sender)
	fmt.Printf("Recvers1: %v\n", keyData.Recvers)
	fmt.Printf("Recvers2: %v\n", DecKeyData.Recvers)
}

//func TestInit() {
//	fmt.Printf("FaucetPrivatekeyStr: %v\n", FaucetPrivatekeyStr)
//	fmt.Printf("EthGateway: %s\n", EthGateway)
//	fmt.Printf("PSK: %v\n", KeyData.Psk)
//	fmt.Printf("Sender: %v\n", KeyData.Sender)
//}
