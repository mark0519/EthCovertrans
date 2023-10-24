package ethio

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type senderAccount struct {
	privateKey ecdsa.PrivateKey
	publicKey  ecdsa.PublicKey
	address    common.Address
	//balance    big.Int
}

type targetAccount struct {
	privateKey ecdsa.PrivateKey
	publicKey  ecdsa.PublicKey
	address    common.Address
}

type ETHSender struct {
	client   *ethclient.Client
	senderAc senderAccount
	targetAc targetAccount
	faucetAc senderAccount
}

func (esdr *ETHSender) NewETHSender(privKey ecdsa.PrivateKey) {
	// 初始化ETHSender
	// privKey是发送方的私钥

	// client初始化
	//c, err := ethclient.Dial("https://sut0ne.tk/v1/sepolia")
	c, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/KvQyzbw_h3XnPpqfWoZ9GcvPAB0iPoDk")
	//c, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	esdr.client = c

	// senderAc初始化
	esdr.senderAc.privateKey = privKey
	esdr.senderAc.publicKey = privKey.PublicKey
	esdr.senderAc.address = crypto.PubkeyToAddress(esdr.senderAc.publicKey)

	// 水龙头Faucet初始化
	// faucet私钥
	faucetPrivKeyStr := "46927aa4aef15bcb8233c953a0a62e0a53334adc27f89767cc82b2e9841a723d"
	esdr.faucetAc.privateKey = *crypto.ToECDSAUnsafe(common.FromHex(faucetPrivKeyStr))
	esdr.faucetAc.publicKey = esdr.faucetAc.privateKey.PublicKey
	esdr.faucetAc.address = crypto.PubkeyToAddress(esdr.faucetAc.publicKey)

	// 发送者sender余额初始化
	esdr.getBalance(esdr.senderAc.address) // 余额如果为0，自动请求faucet获得单位gas
}

func (esdr *ETHSender) getBalance(account common.Address) {
	// 发送者sender余额初始化
	// 获取sepoliaETH账户余额，余额如果为0，自动请求faucet获得单位gas
	balance, err := esdr.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("[Sender]", err)
	}
	log.Printf("[Sender] Sender Account: %s\n", account)
	if balance.Cmp(big.NewInt(0)) == 0 {
		// 余额为0，请求faucet
		log.Printf("[Sender] Sender Balance == 0 wei")
		log.Printf("[Sender] Request Faucets ...")

		// 请求faucet，获得400000000000000单位gas，返回交易哈希
		txHash := esdr.createTx(esdr.faucetAc, targetAccount(esdr.senderAc), big.NewInt(40000000000))
		log.Printf("[Sender] Request Faucets TxHash: %s\n", txHash)

		// 查询余额，如果还是0，递归请求faucet
		esdr.getBalance(account)
	} else {
		log.Printf("[Sender] Sender Balance: %d wei\n", balance)
	}
}

func (esdr *ETHSender) SendMsg(toAddr common.Address) string {
	// 发起0交易,传递信息
	// toAddr: 目标地址
	// 返回交易哈希

	// 创建交易,转账0 wei
	txHash := esdr.createTx(esdr.senderAc, targetAccount{address: toAddr}, big.NewInt(0))
	log.Printf("[Sender] SendETH TxHash: %s\n", txHash)

	// 返回交易哈希
	return txHash
}

func (esdr *ETHSender) createTx(fromAC senderAccount, toAC targetAccount, value *big.Int) string {
	// 创建交易,从fromAC发送到toAC,发送金额为value

	// 获取nonce
	fromAddress := fromAC.address
	nonce, err := esdr.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("[Sender]", err)
	}
	//log.Printf("[Sender] Sender Nonce: %d\n", nonce)

	// 获取gasPrice
	gasPrice, err := esdr.client.SuggestGasPrice(context.Background())
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(3*75000))
	if err != nil {
		log.Fatal("[Sender]", err)
	}
	//fmt.Println(gasPrice)
	// 设置gas上限
	gasLimit := uint64(25000)

	// 转账目标
	toAddress := toAC.address

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
	chanID, err := esdr.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("[Sender] chanID error:", err)
	}
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chanID), &fromAC.privateKey)
	if err != nil {
		log.Fatal("[Sender] signedTx error:", err)
	}

	// 广播节点
	if err := esdr.client.SendTransaction(context.Background(), signedTx); err != nil {
		log.Fatal("[Sender] SendTransaction error:", err)
	}

	// 返回交易hash
	return signedTx.Hash().Hex()
}
