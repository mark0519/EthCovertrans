package ethio

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

func GetReceiverBySenderAddr(address common.Address) {
	// client初始化
	//client, err := ethclient.Dial("https://sut0ne.tk/v1/sepolia")
	//client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/KvQyzbw_h3XnPpqfWoZ9GcvPAB0iPoDk")
	////client, err := ethclient.Dial("https://cloudflare-eth.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	client, err := rpc.Dial("https://eth-sepolia.g.alchemy.com/v2/KvQyzbw_h3XnPpqfWoZ9GcvPAB0iPoDk")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 查询交易
	query := rpc.BatchElem{
		Method: "eth_getLogs",
		Args: []interface{}{
			map[string]interface{}{
				"fromBlock": "0x0",    // 创世区块
				"toBlock":   "latest", // 最新区块
				"address":   address,
			},
		},
		Result: new([]types.Log),
	}

	err = client.BatchCall([]rpc.BatchElem{query})
	if err != nil {
		log.Fatal(err)
	}

	logs := *(query.Result.(*[]types.Log))

	for _, data := range logs {
		fmt.Printf("交易data: %s\n", data.Data)
		fmt.Printf("交易哈希: %s\n", data.TxHash.Hex())
	}
}
