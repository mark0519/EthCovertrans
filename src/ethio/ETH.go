package ethio

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// "https://cloudflare-eth.com"
// "https://sut0ne.tk/v1/sepolia"
const EthGateway = "https://eth-sepolia.g.alchemy.com/v2/KvQyzbw_h3XnPpqfWoZ9GcvPAB0iPoDk"
const FaucetPrivatekeyStr = "46927aa4aef15bcb8233c953a0a62e0a53334adc27f89767cc82b2e9841a723d"

func initETHClient() *ethclient.Client {
	client, err := ethclient.Dial(EthGateway)
	if err != nil {
		panic(err)
	}
	return client
}

func initRPCClient() *rpc.Client {
	client, err := rpc.Dial(EthGateway)
	if err != nil {
		panic(err)
	}
	return client
}
