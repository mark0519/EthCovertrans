package ethio

import (
	"EthCovertrans/src/ethio/util"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
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

	psk := util.GetPskFromFile()     // 从文件读取psk
	senderSK := util.NewPrivateKey() // sender
	sender := util.InitSendAddrData(senderSK)

	// 初始化公钥列表切片
	var recvers []*ecdsa.PublicKey
	// TODO: sender.PublicKey 应该从合约获取
	recvers = append(recvers, sender.PublicKey)

	// 创建KeyData实例
	keyData := util.KeyFileData{
		Psk:     psk,
		Sender:  senderSK,
		Recvers: &recvers, // 公钥列表
	}

	// 将结构体转换为JSON格式的[]byte
	keyDataBytes, err := json.Marshal(keyData)
	if err != nil {
		log.Panic("[Sender] Error marshaling:", err)
	}
	keyFile := new(util.KeyFileData)
	// 使用 json.Unmarshal 将 JSON 格式的字节切片转换回 KeyFileData 结构体
	err = json.Unmarshal(keyDataBytes, &keyFile)
	if err != nil {
		log.Panic("[Sender] Error unmarshaling:", err)
	}

	util.EncryptKeyFileData(keyData, "ethCoverTrans.key")      // 加密并保存
	DecKeyData := util.DecryptKeyFileData("ethCoverTrans.key") // 读取并解密

	fmt.Printf("psk1: %v\n", keyData.Psk)
	fmt.Printf("psk2: %v\n", DecKeyData.Psk)
	fmt.Printf("Sender1: %v\n", keyData.Sender)
	fmt.Printf("Sender2: %v\n", DecKeyData.Sender)
	fmt.Printf("Recvers1: %v\n", keyData.Recvers)
	fmt.Printf("Recvers2: %v\n", DecKeyData.Recvers)

	//newPK := TestMsgSenderFactory("hello world!!", psk, sender.GetSendAddrDataPrivateKey())
	//msg := MsgRecverFactory(psk, sender.PublicKey, newPK)
	//print("[+] msg:", msg)
}
