package ethio

import (
	"EthCovertrans/src/allcrypto"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

type ETHSender struct {
	sendAc   *allcrypto.SendAddrData
	recvAc   *allcrypto.RecvAddrData
	faucetAc *allcrypto.SendAddrData
}

func (esdr *ETHSender) newETHSender(send *allcrypto.SendAddrData, recv *allcrypto.RecvAddrData) {
	// 初始化ETHSender

	//Client = initETHClient()
	// sendAc初始化
	esdr.sendAc = send

	// 水龙头Faucet初始化
	faucetSk := crypto.ToECDSAUnsafe(common.FromHex(FaucetPrivatekeyStr))
	esdr.faucetAc = allcrypto.InitSendAddrData(faucetSk)

	// 发送者sender余额初始化
	esdr.initSenderBalance() // 余额如果为0，自动请求faucet获得单位gas

	esdr.recvAc = recv
}

func (esdr *ETHSender) initSenderBalance() {
	// 发送者sender余额初始化
	// 获取sepoliaETH账户余额，余额如果为0，自动请求faucet获得单位gas

	balance, err := Client.BalanceAt(context.Background(), esdr.sendAc.Address, nil)
	if err != nil {
		log.Fatal("[Sender]", err)
	}
	log.Printf("[Sender] Sender Account: %s\n", esdr.sendAc.Address)
	if balance.Cmp(big.NewInt(0)) == 0 {
		// 余额为0，请求faucet
		log.Printf("[Sender] Sender Balance == 0 wei")
		log.Printf("[Sender] Request Faucets ...")

		// 请求faucet，获得400000000000000单位wei，返回交易哈希
		gas := big.NewInt(10000000000000000)
		txHash := esdr.supplyFromFaucet(gas)
		log.Printf("[Sender] Request Faucets TxHash: %s\n", txHash)
		waitForTx(txHash)
	} else {
		log.Printf("[Sender] Sender Balance: %d wei\n", balance)
	}
}
func (esdr *ETHSender) supplyFromFaucet(value *big.Int) string {
	// 从水龙头faucet请求value wei
	txHash := createTx(esdr.faucetAc, esdr.sendAc.AddrData, value)
	// 返回交易哈希
	return txHash
}

func (esdr *ETHSender) sendETH(value *big.Int) string {
	// 创建交易,转账value wei
	txHash := createTx(esdr.sendAc, esdr.recvAc.AddrData, value)
	log.Printf("[Sender] SendETH TxHash: %s\n", txHash)

	// 返回交易哈希
	return txHash
}

func waitForTx(txHash string) {
	// 等待交易完成
	for {
		if getStatusByTxhash(txHash) {
			break
		}
	}
	//fmt.Println("success")
}

func getStatusByTxhash(txHash string) bool {
	// 根据交易哈希查询交易状态
	receipt, err := Client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return false
	}
	return receipt.Status == types.ReceiptStatusSuccessful
}

func createTx(fromAC *allcrypto.SendAddrData, toAC *allcrypto.AddrData, value *big.Int) string {
	// 创建交易,从fromAC发送到toAC,发送金额为value

	// 获取nonce
	fromAddress := fromAC.Address
	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("[Sender]", err)
	}
	//log.Printf("[Sender] Sender Nonce: %d\n", nonce)

	// 获取gasPrice
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(3))
	if err != nil {
		log.Fatal("[Sender]", err)
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
