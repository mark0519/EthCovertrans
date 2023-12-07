package ethio

import (
	"EthCovertrans/src/ethio/util"
	"crypto/ecdsa"
	"math/big"
)

func TestMsgSenderFactory(msgstr string, psk *ecdsa.PrivateKey, orignSenderSK *ecdsa.PrivateKey) *ecdsa.PublicKey {
	// 创建ETHSender实例

	msgIntSlice := sliceMsg(msgstr)
	var times = len(msgIntSlice)
	orignSender := util.InitSendAddrData(orignSenderSK)
	senders := *newSenderList(times, psk, orignSender)
	recvers := *newRecverList(times, psk)

	msgSenders := make([]MsgSender, times)
	for i := 0; i < times; i++ {
		msgSenders[i] = *newMsgSender(&(senders)[i], &(recvers)[i], big.NewInt(int64(recvers[i].Msg^msgIntSlice[i])))
	}

	doSend(&msgSenders)
	defer util.UpdateContract(senders[times].PublicKey)

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

	psk := util.NewPrivateKey()
	senderSK := util.NewPrivateKey()
	sender := util.InitSendAddrData(senderSK)
	newPK := TestMsgSenderFactory("hello world!!", psk, sender.GetSendAddrDataPrivateKey())
	msg := MsgRecverFactory(psk, sender.PublicKey, newPK)
	print("[+] msg:", msg)
}
