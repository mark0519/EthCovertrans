package allcrypto

import (
	"crypto/ecdsa"
	"crypto/md5"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type addrData struct {
	privateKey []byte
	publicKey  []byte
	address    string
}

var addrList [][]addrData

/*
addrList[00] = [data,data,data]
addrList[01] = [data,data,data]
addrList[10] = [data,data,data]
addrList[11] = [data,data,data]
data = [privKey,PubKey,addr]
*/

func InitAddrTable(n int) [][]addrData {
	// n 交易嵌入位数
	addrList = getListByN(n)
	return addrList
}

func getPubKeyByPriv(privKey []byte) []byte {
	// 通过私钥计算公钥（ pk = sk*G ）
	privateKeyBytes := privKey
	privateKey := new(ecdsa.PrivateKey)
	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
	privateKey.PublicKey.Curve = crypto.S256()
	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())
	publicKey := privateKey.PublicKey
	return crypto.FromECDSAPub(&publicKey)
}

func getAddrByPub(pubKey []byte) string {
	// 通过公钥计算钱包地址
	p, _ := crypto.UnmarshalPubkey(pubKey)
	address := crypto.PubkeyToAddress(*p)
	return address.Hex()
}

func getHashByAddrPsk(addr string, psk string) [16]byte {
	// 通过钱包地址计算交易哈希
	data := []byte(addr + psk)
	// 计算data的md5
	hashT := md5.Sum(data)
	return hashT
}

func getListByN(n int) [][]addrData {
	sum := 1 << n                   // 2的n次幂
	rows := make([][]addrData, sum) // 创建sum行切片
	for i := range rows {
		// 为每一行创建一个切片，可以根据需要指定初始长度
		rows[i] = make([]addrData, 0) // 这里初始长度为 0
	}
	return rows
}

func appendListByPriv(table [][]addrData, privKey []byte, psk []byte, n int) [][]addrData {
	// 通过哈希计算 地址集合addrList
	pubKey := getPubKeyByPriv(privKey)
	addr := getAddrByPub(pubKey)
	hash := getHashByAddrPsk(addr, string(psk))
	//fmt.Printf("%x\n", hash)
	var result byte
	result = (1 << n) - 1
	lowestTwoBits := hash[15] & result
	//fmt.Printf("%x\n", lowestTwoBits)
	//fmt.Printf("%x\n", pubKey)
	addrData := addrData{
		privateKey: privKey,
		publicKey:  pubKey,
		address:    addr,
	}
	table[lowestTwoBits] = append(table[lowestTwoBits], addrData)
	return table
}

func newPrivKey(oldKey []byte, psk []byte) []byte {
	var oldKeyInt big.Int
	oldKeyInt.SetBytes(oldKey)
	var pskInt big.Int
	pskInt.SetBytes(psk)

	newKey := new(big.Int).Mul(&oldKeyInt, &pskInt)
	newKey = newKey.Mod(newKey, secp256k1.S256().Params().N)
	return newKey.Bytes()
}

func SetAddrDatas(table [][]addrData, privKey []byte, psk []byte, n int, times int) []byte {
	// 循环计算times次并保存到table里，返回最后一次计算之后的下一次还没使用的私钥
	key := privKey
	for i := 0; i < times; i++ {
		appendListByPriv(table, key, psk, n)
		key = newPrivKey(key, psk)
	}
	return key
}

func ShowAddrTable(table [][]addrData) {
	for i := range table {
		fmt.Println("============table[", i, "]==============")
		for j := range table[i] {
			fmt.Printf("publicKey :0x%x\n", table[i][j].publicKey)
			fmt.Printf("privateKey:0x%x\n", table[i][j].privateKey)
			fmt.Println("Addr      :", table[i][j].address)
			fmt.Println("---")
		}
	}
}

//func getAddrDatas() {
//
//}
//
//func GetAddrByMsg(msg int, n int) addrData {
//
//}
