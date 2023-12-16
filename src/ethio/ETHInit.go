package ethio

import (
	"EthCovertrans/src/ethio/contract"
	"EthCovertrans/src/ethio/util"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

// "https://cloudflare-eth.com"
// const EthGateway = "https://sut0ne.tk/v1/sepolia"
// const FaucetPrivatekeyStr = "983ec812c710bd1a3ef13bfd089cf8c7cf672f8bf17a7b9be51318c8314120aa"
// const EthGateway = "wss://eth-sepolia.g.alchemy.com/v2/tTrWBB8FMZ7wfeBfv3gjYc7w9-pq_jb2"
// const EtherscanAPIKey = "WE5VDDZE6QVKYC194QM44QVUEWUPTCGH8I"
// const EtherscanAPIURL = "https://api-sepolia.etherscan.io/api"
// const MsgSliceLen = 32 // 每次发送的消息比特数
// const MsgSliceBytesLen = 4
// const ContractAddress = "0x7d54615Cb5f7d30d244b0F6cC8BB0681D42236bD"
// const KeyFile = "ethCoverTrans.key"

var FaucetPrivatekeyStr string
var EthGateway string
var EtherscanAPIKey string
var EtherscanAPIURL string
var MsgSliceLen int // 每次发送的消息比特数
var MsgSliceBytesLen int
var ContractAddress string
var KeyFile string
var Client *ethclient.Client
var FaucetAc *util.SendAddrData
var KeyData *util.KeyFileData
var ChainId *big.Int
var EthContract *contract.Contract
var AuthTransact *bind.TransactOpts

func init() {
	initConfigFile()         // config初始化 ，必须放在第一个初始化
	Client = initETHClient() // ETHClient初始化
	initKeyDataFromFile()    // KeyData初始化 ，必须在Faucet初始化之前
	initFaucet()             // Faucet 初始化 ，必须在KeyData初始化之后
	initContract()           // EthContract 初始化
}

func initFaucet() {
	// 水龙头Faucet初始化
	faucetSk := crypto.ToECDSAUnsafe(common.FromHex(FaucetPrivatekeyStr))
	FaucetAc = util.InitSendAddrData(faucetSk)
}

func initETHClient() *ethclient.Client {
	client, err := ethclient.Dial(EthGateway)
	if err != nil {
		panic(err)
	}
	return client
}

func initKeyDataFromFile() {
	// 获取 []KeyFileData 和水龙头私钥 （没有私钥需要先设置）

	// 输入文件加密口令
	fmt.Print("[Sender] Please Enter the File Encryption Password: ")
	var inputPassword string
	_, err := fmt.Scanln(&inputPassword)
	if err != nil {
		log.Panic("[Sender] Error reading password:", err)
	}
	// 计算口令md5作为加密密钥FileAesKey
	hasher := md5.New()
	hasher.Write([]byte(inputPassword))
	util.FileAesKey = hasher.Sum(nil)

	var data util.KeyFileDataAndFaucet
	if _, err := os.Stat("ethCoverTrans.key"); err == nil { // 如果文件存在
		log.Print("[Sender] Loading init key file: ethCoverTrans.key ...")
		// 读取文件
		data = util.DecryptKeyFileData(KeyFile)
		// 返回 psk, senderPrivateKey, FaucetPrivateKey
	} else if os.IsNotExist(err) {
		data = util.GenerateKeyFile("psk.key", "private.key", KeyFile)
		RegisterRecv(data.KeyFileData.Psk, util.PrivateKeyToAddrData(data.KeyFileData.Sender).PublicKey) // 注册公钥
	}

	KeyData = &data.KeyFileData
	FaucetPrivatekeyStr = data.Faucet
	log.Print("[Sender] Get Faucet Private Key: ", FaucetPrivatekeyStr)
}

func initConfigFile() {
	// config文件初始化
	log.Print("[Sender] Loading Config File ... ")
	// 读取 Config 文件
	configFilePath := "config.json"
	configDataBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal("[Sender] Not Found config file: ", configFilePath)
	}

	type Config struct {
		EthGateway       string `json:"EthGateway"`
		EtherscanAPIKey  string `json:"EtherscanAPIKey"`
		EtherscanAPIURL  string `json:"EtherscanAPIURL"`
		MsgSliceLen      int    `json:"MsgSliceLen"`
		MsgSliceBytesLen int    `json:"MsgSliceBytesLen"`
		ContractAddress  string `json:"ContractAddress"`
		KeyFile          string `json:"KeyFile"`
		ChainId          int    `json:"ChainId"`
	}
	// 解析 JSON 数据到结构体
	var config Config
	err = json.Unmarshal(configDataBytes, &config)
	if err != nil {
		log.Panic("Unmarshal Config.JSON Error: ", err)
	}
	EthGateway = config.EthGateway
	EtherscanAPIKey = config.EtherscanAPIKey
	EtherscanAPIURL = config.EtherscanAPIURL
	MsgSliceLen = config.MsgSliceLen
	MsgSliceBytesLen = config.MsgSliceBytesLen
	ContractAddress = config.ContractAddress
	KeyFile = config.KeyFile
	ChainId = big.NewInt(int64(config.ChainId))
	log.Print("[Sender] Config File Loading Completed ")
}

func initContract() {
	var err error
	gasPrice, _ := Client.SuggestGasPrice(context.Background())
	nonce, _ := Client.NonceAt(context.Background(), FaucetAc.Address, nil)
	AuthTransact, err = bind.NewKeyedTransactorWithChainID(FaucetAc.GetSendAddrDataPrivateKey(), ChainId) // 将链ID替换为相应的链ID
	if err != nil {
		log.Fatal(err)
	}
	// 设置 gas 限制和 gas 价格
	AuthTransact.Nonce = new(big.Int).SetUint64(nonce)
	AuthTransact.GasLimit = uint64(300000)
	AuthTransact.GasPrice = gasPrice // 1 Gwei
	// 创建已部署合约的绑定实例
	EthContract, err = contract.NewContract(common.HexToAddress(ContractAddress), Client)
	if err != nil {
		log.Fatal(err)
	}
}
