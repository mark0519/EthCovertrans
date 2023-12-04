package ethio

import (
	"EthCovertrans/src/ethio/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// "https://cloudflare-eth.com"
//const EthGateway = "https://sut0ne.tk/v1/sepolia"

// const EthGateway = "https://eth-sepolia.g.alchemy.com/v2/tTrWBB8FMZ7wfeBfv3gjYc7w9-pq_jb2"
const EthGateway = "wss://eth-sepolia.g.alchemy.com/v2/tTrWBB8FMZ7wfeBfv3gjYc7w9-pq_jb2"
const FaucetPrivatekeyStr = "983ec812c710bd1a3ef13bfd089cf8c7cf672f8bf17a7b9be51318c8314120aa"
const EtherscanAPIKey = "WE5VDDZE6QVKYC194QM44QVUEWUPTCGH8I"
const EtherscanAPIURL = "https://api-sepolia.etherscan.io/api"
const MsgSliceLen = 32 // 每次发送的消息比特数
const MsgSliceBytesLen = 4
const ContractAddress = "0x7d54615Cb5f7d30d244b0F6cC8BB0681D42236bD"
const KeyFile = ""

var Client *ethclient.Client
var FaucetAc *util.SendAddrData
var passwd string

func init() {
	Client = initETHClient()
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

func initKeyData() {
	// TODO: 加密文件解密,获取 []KeyFileData
}
