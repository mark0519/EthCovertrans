package ethio

import (
	"crypto/ecdsa"
	"ethio/util"
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
		msgSenders[i] = MsgSender{
			sendAc: &(senders)[i],
			recvAc: &(recvers)[i],
			value:  big.NewInt(int64(recvers[i].Msg ^ msgIntSlice[i])),
		}
	}

	doSend(&msgSenders)
	defer util.UpdateContract(senders[times].PublicKey)

	return senders[times].PublicKey
}

func Test() {
	psk := util.NewPrivateKey()
	senderSK := util.NewPrivateKey()
	sender := util.InitSendAddrData(senderSK)
	newPK := TestMsgSenderFactory("hello world!", psk, sender.GetSendAddrDataPrivateKey())
	msg := MsgRecverFactory(psk, sender.PublicKey, newPK)
	print("[+] msg:", msg)
}
