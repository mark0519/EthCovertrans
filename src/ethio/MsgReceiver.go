package ethio

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"ethio/util"
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
	latestIdx int // 查找到的最新的一笔交易的idx
}

func NewMsgReceiver(addr common.Address) (recvr *MsgReceiver) {
	// 初始化ETHReceiver ，初始化发送者addr, 从EtherscanAPI查询addr的所有发出交易,定位最新一次addr作为From的交易

	recvr.recvAc = &util.AddrData{Address: addr}
	recvr.latestIdx = -1
	recvr.waitForInfo()
	recvr.getLatestTransIdx()
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

func (recvr *MsgReceiver) waitForInfo() {
	// API限制5s一次，等待5s
	for {
		if recvr.getReceiverInfo() {
			break
		}
		log.Println("[Receiver] Get Receiver Info By EtherscanAPI Failed, Retry in 5s ...")
		time.Sleep(5 * time.Second)
	}
}

func (recvr *MsgReceiver) getReceiverInfo() bool {
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
	//fmt.Printf(req.URL.String())

	// 设置超时限制 timeout 为5s
	var t int64 = 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	defer cancel()
	resp, _ := http.DefaultClient.Do(req.WithContext(ctx))

	body, _ := io.ReadAll(resp.Body)

	var data *ApiData

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	recvr.recvData = data
	return data.Message == "OK"
}

func doRecv(psk *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) []byte {
	addr := crypto.PubkeyToAddress(*publicKey)
	recver := NewMsgReceiver(addr)
	value := recver.GetLatestTransValue()
	toAddr := recver.GetLatestToAddress()
	orignMsg := util.CalcMsg(toAddr, psk, MsgSliceLen)
	msgInt := orignMsg ^ int32(value.Int64())
	msg := make([]byte, MsgSliceLen)
	binary.LittleEndian.PutUint32(msg, uint32(msgInt))
	return msg
}

func MsgRecverFactory(psk *ecdsa.PrivateKey, orignPublicKey *ecdsa.PublicKey, lastPublicKey *ecdsa.PublicKey) []byte {
	var msg []byte
	for derivationKey := util.DerivationPublicKey(orignPublicKey, psk); derivationKey != lastPublicKey; derivationKey = util.DerivationPublicKey(derivationKey, psk) {
		msg = append(msg, doRecv(psk, derivationKey)...)
		log.Println("[Receiver] Recv Msg:", string(msg))
	}
	return msg
}
