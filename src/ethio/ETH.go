package ethio

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// "https://cloudflare-eth.com"
//const EthGateway = "https://sut0ne.tk/v1/sepolia"

// const EthGateway = "https://eth-sepolia.g.alchemy.com/v2/tTrWBB8FMZ7wfeBfv3gjYc7w9-pq_jb2"
const EthGateway = "wss://eth-sepolia.g.alchemy.com/v2/tTrWBB8FMZ7wfeBfv3gjYc7w9-pq_jb2"
const FaucetPrivatekeyStr = "983ec812c710bd1a3ef13bfd089cf8c7cf672f8bf17a7b9be51318c8314120aa"
const EtherscanAPIKey = "WE5VDDZE6QVKYC194QM44QVUEWUPTCGH8I"
const EtherscanAPIURL = "https://api-sepolia.etherscan.io/api"
const MsgSliceLen = 8 // 每次发送的消息比特数

var Client *ethclient.Client

func init() {
	Client = initETHClient()
}

func initETHClient() *ethclient.Client {
	client, err := ethclient.Dial(EthGateway)
	if err != nil {
		panic(err)
	}
	return client
}
