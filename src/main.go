package main

import (
	"EthCovertrans/src/ethio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "EthCovertrans",
	Short: "EthCovertrans",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send command",
	Long:  `This command sends something.`,
	Run: func(cmd *cobra.Command, args []string) {
		funcSend()
	},
}

var recvCmd = &cobra.Command{
	Use:   "recv",
	Short: "recv command",
	Long:  `This command recv something.`,
	Run: func(cmd *cobra.Command, args []string) {
		funcRecver()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(recvCmd)
}

func funcSend() { // 发送消息
	fmt.Println("Executing func1()")
}

func funcRecver() { // 接收消息
	fmt.Println("Executing func2()")
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "EthCovertrans",
		Short: "EthCovertrans",
		Run: func(cmd *cobra.Command, args []string) {
			// 显示帮助信息
			_ = cmd.Help()
		},
	}

	var sendCmd = &cobra.Command{
		Use:   "send",
		Short: "Send data",
		Run: func(cmd *cobra.Command, args []string) {
			funcSend()
		},
	}

	var recvCmd = &cobra.Command{
		Use:   "recv",
		Short: "Receive data",
		Run: func(cmd *cobra.Command, args []string) {
			funcRecver()
		},
	}

	rootCmd.AddCommand(sendCmd)
	rootCmd.AddCommand(recvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Listener() {
	for {
		ethio.DiffKeyData()
		time.Sleep(30 * time.Second)
	}
}
