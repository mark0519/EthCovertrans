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
	//recvr := new(MsgReceiver)
	//addr := "0xeeab000c366718eb463b5ac12c53bde5b8aa7ca0"
	//recvr.recvAc = &util.AddrData{
	//	Address:   common.HexToAddress(addr),
	//	PublicKey: nil,
	//}
	//recvr.waitForInfo()
	//recvr.getLatestTransIdx()
	//recvr.GetLatestTransValue()
	//recvr.GetLatestToAddress()
	psk := util.NewPrivateKey()
	senderSK := util.NewPrivateKey()
	sender := util.InitSendAddrData(senderSK)
	newPK := TestMsgSenderFactory("stone", psk, sender.GetSendAddrDataPrivateKey())
	msg := MsgRecverFactory(psk, sender.PublicKey, newPK)
	print("[+] msg:", string(msg))
}
