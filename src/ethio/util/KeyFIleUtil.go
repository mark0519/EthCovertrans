package util

import (
	. "EthCovertrans/src/cryptoUtil"
	"crypto/ecdsa"
)

type KeyFileData struct {
	psk     *ecdsa.PrivateKey
	Sender  *SendAddrData
	Recvers *[]RecvAddrData
}

func encryptKeyFileData() {
	// TODO: 加密KeyFileData存入文件
}

func decryptKeyFileData() {
	// TODO: 从文件读取KeyFileData并解密
}
