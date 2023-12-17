package main

import (
	"EthCovertrans/src/ethio"
	"fmt"
	"os"
	"time"
)

func listener() {
	// 监听消息
	for {
		ethio.DiffKeyData()
		time.Sleep(30 * time.Second)
	}
}

func main() {
	// 检查参数的数量
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("EthCovertrans Usage:")
		fmt.Println("	Usage for register: EthConvertrans -r")
		fmt.Println("	Usage for listening: EthConvertrans -l")
		fmt.Println("	Usage for send message: EthConvertrans -s <your_msg>")
		fmt.Println("	Usage for force update history: EthConvertrans -f")
		os.Exit(1)
	}

	// 遍历参数
	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch arg {
		case "-s":
			// 检查下一个参数是否存在
			if i+1 < len(os.Args) {
				// 发送消息
				msg := os.Args[i+1]
				ethio.MsgSenderFactory(msg, ethio.KeyData.Psk, ethio.KeyData.Sender)
			} else {
				fmt.Println("Missing value for -s")
				os.Exit(1)
			}
		case "-l":
			// 接收消息
			listener()
		case "-f":
			// 强制修改本地公钥
			ethio.ForceUpdateLocal()
		case "-r":
			// 注册本地公钥
			ethio.RegisterRecv()
		}

	}
}
