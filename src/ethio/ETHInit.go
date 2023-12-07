package ethio

import (
	"EthCovertrans/src/ethio/util"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

// "https://cloudflare-eth.com"
// const EthGateway = "https://sut0ne.tk/v1/sepolia"
//const FaucetPrivatekeyStr = "983ec812c710bd1a3ef13bfd089cf8c7cf672f8bf17a7b9be51318c8314120aa"
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

func Init() {
	initConfigFile()         // config初始化 ，必须放在第一个初始化
	Client = initETHClient() // ETHClient初始化
	initKeyDataFromFile()    // KeyData初始化 ，必须在Faucet初始化之前
	initFaucet()             // Faucet 初始化 ，必须在KeyData初始化之后
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

	data := util.GenerateKeyFile("psk.key", "privateKey.key", KeyFile)
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
}
