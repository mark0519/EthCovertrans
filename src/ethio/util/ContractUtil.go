package util

import "crypto/ecdsa"

func UpdateContract(key *ecdsa.PublicKey) {
	// TODO: 更新合约公钥，更新本地公钥
}

func diffKeyData() {
	// TODO: 从合约获取公钥，与本地公钥比对，返回不同的公钥
}
