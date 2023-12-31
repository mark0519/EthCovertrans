package ethio

import (
	"EthCovertrans/src/ethio/util"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

type MsgSender struct {
	sendAc *util.SendAddrData
	recvAc *util.RecvAddrData
	value  *big.Int
}

func (msdr *MsgSender) supplyFromFaucet(gas *big.Int) string {
	// 从水龙头faucet请求gas wei
	txHash := createTx(FaucetAc, msdr.sendAc.AddrData, gas)
	// 返回交易哈希
	return txHash
}

func (msdr *MsgSender) initSenderBalance() {
	// 发送者sender余额初始化
	// 获取sepoliaETH账户余额，余额如果为0，自动请求faucet获得单位gas

	balance, err := Client.BalanceAt(context.Background(), msdr.sendAc.Address, nil)
	if err != nil {
		log.Fatal("[Sender] get balance:", err)
	}
	log.Printf("[Sender] Sender Account: %s\n", msdr.sendAc.Address)
	if balance.Cmp(big.NewInt(0)) == 0 {
		// 余额为0，请求faucet
		log.Printf("[Sender] Sender Balance == 0 wei")
		log.Printf("[Sender] Request Faucets ...")

		// 请求faucet，获得1000000000000000单位wei，返回交易哈希
		gas := big.NewInt(1000000000000000)
		txHash := msdr.supplyFromFaucet(gas)
		log.Printf("[Sender] Request Faucets TxHash: %s\n", txHash)
		waitForTx(txHash)
	} else {
		log.Printf("[Sender] Sender Balance: %d wei\n", balance)
	}
}

func NewMsgSender(send *util.SendAddrData, recv *util.RecvAddrData, value *big.Int) (msdr *MsgSender) {
	msdr = new(MsgSender)
	// 发送方地址初始化
	msdr.sendAc = send
	// 发送方余额初始化
	msdr.initSenderBalance() // 余额如果为0，自动请求faucet获得单位gas
	// 接收方地址初始化
	msdr.recvAc = recv
	msdr.value = value
	return
}

func (msdr *MsgSender) sendETH() string {
	// 创建交易,转账value wei
	txHash := createTx(msdr.sendAc, msdr.recvAc.AddrData, msdr.value)
	log.Printf("[Sender] SendETH TxHash: %s\n", txHash)
	waitForTx(txHash)
	// 返回交易哈希
	return txHash
}

func waitForTx(txHash string) bool {
	log.Print("[Sender] Waiting For Transaction ...")
	for {
		// 根据交易哈希查询交易状态
		receipt, err := Client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		if receipt.Status == types.ReceiptStatusSuccessful {
			break
		}
		time.Sleep(3 * time.Second)
	}
	return true
}

func createTx(fromAC *util.SendAddrData, toAC *util.AddrData, value *big.Int) string {
	// 创建交易,从fromAC发送到toAC,发送金额为value

	// 获取nonce
	fromAddress := fromAC.Address
	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("[Sender] get nonce:", err)
	}
	//log.Printf("[Sender] Sender Nonce: %d\n", nonce)

	// 获取gasPrice
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(3))
	if err != nil {
		log.Fatal("[Sender] get gasPrice:", err)
	}
	//fmt.Println(gasPrice)
	// 设置gas上限
	gasLimit := uint64(25000)

	// 转账目标
	toAddress := toAC.Address

	// 创建交易,转账value wei
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     make([]byte, 0),
	})

	// 使用发送者私钥签名交易
	chanID, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("[Sender] chanID error:", err)
	}
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chanID), fromAC.GetSendAddrDataPrivateKey())
	if err != nil {
		log.Fatal("[Sender] signedTx error:", err)
	}

	// 广播节点
	if err := Client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal("[Sender] SendTransaction error:", err)
	}

	// 返回交易hash
	return signedTx.Hash().Hex()
}

func NewSenderList(times int, psk *ecdsa.PrivateKey, originSender *util.SendAddrData) *[]util.SendAddrData {
	// 创建发送者列表
	senderList := make([]util.SendAddrData, times+1)
	senderList[0] = *(util.DerivationSendAddrData(originSender, psk))

	// 多计算一个以更新合约
	for i := 1; i < times+1; i++ {
		senderList[i] = *(util.DerivationSendAddrData(&senderList[i-1], psk))
	}
	return &senderList
}

func NewRecverList(times int, psk *ecdsa.PrivateKey) *[]util.RecvAddrData {
	recverList := make([]util.RecvAddrData, times)
	for i := 0; i < times; i++ {
		recverList[i] = *(util.InitRecvAddrData(psk, MsgSliceLen))
	}
	return &recverList
}

func DoSend(msgSenders *[]MsgSender) {
	// 发送信息
	// todo: 并发数量限制
	for i := 0; i < len(*msgSenders); i++ {
		(*msgSenders)[i].sendETH()
	}
}

func MsgSenderFactory(msgstr string, psk *ecdsa.PrivateKey, orignSenderSK *ecdsa.PrivateKey) {
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
	log.Print("[Sender] Send msg Success:", msgstr)

	// 更新本地sender私钥
	// 更新合约公钥
	UpdateContract(senders[times])
}

func SliceMsg(msg string) []int32 {
	msgByteSlice := []byte(msg)
	if len(msgByteSlice)%MsgSliceBytesLen != 0 {
		padding := 4 - len(msgByteSlice)%MsgSliceBytesLen
		for i := 0; i < padding; i++ {
			msgByteSlice = append(msgByteSlice, 0)
		}
	}
	msgIntSlice := make([]int32, len(msgByteSlice)/MsgSliceBytesLen)
	for i := 0; i < len(msgByteSlice); i += MsgSliceBytesLen {
		end := i + MsgSliceBytesLen
		if end > len(msgByteSlice) {
			end = len(msgByteSlice)
		}
		subSlice := msgByteSlice[i:end]
		var num int32
		err := binary.Read(bytes.NewReader(subSlice), binary.LittleEndian, &num)
		if err != nil {
			log.Fatal("[Sender] binary.Read failed:", err)
		} else {
			msgIntSlice[i/MsgSliceBytesLen] = num
		}
	}
	return msgIntSlice
}
