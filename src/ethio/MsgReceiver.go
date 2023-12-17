package ethio

import (
	"EthCovertrans/src/ethio/util"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"
)

// ApiData 定义接收数据结构体ApiData
type ApiData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		TransHash string `json:"hash"`
		BlockHash string `json:"blockHash"`
		From      string `json:"from"`
		To        string `json:"to"`
		Value     string `json:"value"`
	}
}

type MsgReceiver struct {
	recvAc    *util.AddrData
	recvData  *ApiData
	latestIdx int  // 查找到的最新的一笔交易的idx
	msgend    bool // 本次消息接收是否结束
}

func NewMsgReceiver(publicKey *ecdsa.PublicKey) (recvr *MsgReceiver) {
	// 初始化ETHReceiver ，初始化发送者addr, 从EtherscanAPI查询addr的所有发出交易,定位最新一次addr作为From的交易
	recvr = new(MsgReceiver)
	recvr.recvAc = &util.AddrData{
		Address:   crypto.PubkeyToAddress(*publicKey),
		PublicKey: publicKey,
	}

	recvr.latestIdx = -1
	for recvr.latestIdx == -1 {
		recvr.msgend = recvr.waitForInfo()
		if recvr.msgend {
			recvr.getLatestTransIdx()
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	return
}

func (recvr *MsgReceiver) GetLatestToAddress() common.Address {
	// 返回最新的一笔交易的接收者to地址

	// 没找到recvr.recvAc.Address作为From的交易
	if recvr.latestIdx == -1 {
		log.Fatal("[Receiver] Get Trans (From:", recvr.recvAc.Address.String(), ") Failed")
	}

	toAddr := recvr.recvData.Result[recvr.latestIdx].To
	toAddressByte, _ := hex.DecodeString(toAddr[2:]) // 去掉地址开头0x
	return common.Address(toAddressByte)
}

func (recvr *MsgReceiver) GetLatestTransValue() *big.Int {
	// 返回最新的一笔交易的交易金额Value 单位Wei

	// 没找到recvr.recvAc.Address作为From的交易
	if recvr.latestIdx == -1 {
		log.Fatal("[Receiver] Get Trans (From:", recvr.recvAc.Address.String(), ") Failed")
	}
	value := recvr.recvData.Result[recvr.latestIdx].Value
	n := new(big.Int)
	n, _ = n.SetString(value, 10)
	return n
}

func (recvr *MsgReceiver) getLatestTransIdx() {
	// 返回最新的recvr.recvAc.Address作为From的交易的idx

	for idx := 0; idx < len(recvr.recvData.Result); idx++ {
		// 如果from地址是发送者自己的地址，那么to地址就是接收者的地址
		// EqualFold 不区分大小写比较
		if strings.EqualFold(recvr.recvData.Result[idx].From, recvr.recvAc.Address.String()) {
			// API倒序，第一个找到的就是最新的
			recvr.latestIdx = idx
			break
		}
	}
}

func (recvr *MsgReceiver) waitForInfo() bool {
	// API限制5s一次，等待5s，尝试5次
	for i := 0; i < 5; i++ {
		msg := recvr.getReceiverInfo()
		if msg == "OK" {
			return true
		} else if msg == "No transactions found" {
			log.Println("[Receiver] Get nothing by EtherscanAPI")
			return false
		} else {
			log.Println("[Receiver] Get receiver Info By EtherscanAPI Failed, Retry in 5s ...")
			time.Sleep(5 * time.Second)
		}
	}
	log.Fatal("[Receiver] Unable to get receiver info by EtherscanAPI")
	return false
}

func (recvr *MsgReceiver) getReceiverInfo() string {
	// 向EtherscanAPI查询addr的所有发出交易

	// 初始化请求
	req, err := http.NewRequest("GET", EtherscanAPIURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 设置参数
	params := req.URL.Query()
	params.Add("module", "account")
	params.Add("action", "txlist")
	params.Add("address", recvr.recvAc.Address.String()) // 发送者addr
	params.Add("startblock", "0")
	params.Add("endblock", "99999999")
	params.Add("sort", "desc") // 正序aes 倒序desc
	params.Add("apikey", EtherscanAPIKey)
	req.URL.RawQuery = params.Encode()

	log.Println("[Receiver] Api url:", req.URL.String())

	// 设置超时限制 timeout 为5s
	var t int64 = 20
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	defer cancel()
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
	}

	body, _ := io.ReadAll(resp.Body)

	data := new(ApiData)
	err = json.Unmarshal(body, data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	recvr.recvData = data
	log.Print("[Receiver] Get data.Message By EtherscanAPI:", data.Message)
	//fmt.Println(data.Message)
	return data.Message
}

func doRecv(psk *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) []byte {

	recver := NewMsgReceiver(publicKey)
	if recver.msgend {
		value := recver.GetLatestTransValue()
		toAddr := recver.GetLatestToAddress()
		orignMsg := util.CalcMsg(toAddr, psk, MsgSliceLen)
		msgInt := orignMsg ^ int32(value.Int64())
		msg := make([]byte, MsgSliceLen)
		binary.LittleEndian.PutUint32(msg, uint32(msgInt))
		return msg
	} else {
		log.Println("[Receiver] This message end!!")
		return []byte("|")
	}

}

func MsgRecverFactory(psk *ecdsa.PrivateKey, orignPublicKey *ecdsa.PublicKey, lastPublicKey *ecdsa.PublicKey) []byte {
	var msg []byte
	for derivationKey := util.DerivationPublicKey(orignPublicKey, psk); derivationKey.X.Cmp(lastPublicKey.X) != 0 || derivationKey.Y.Cmp(lastPublicKey.Y) != 0; derivationKey = util.DerivationPublicKey(derivationKey, psk) {
		msg = append(msg, doRecv(psk, derivationKey)...)
		log.Println("[Receiver] Recv msg:", string(msg))
	}
	return msg
}
