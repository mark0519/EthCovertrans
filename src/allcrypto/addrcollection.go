package allcrypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type addrData struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    string
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

func appendListByPriv(table [][]addrData, privKey []byte, psk []byte, n int) byte {
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
		PrivateKey: privKey,
		PublicKey:  pubKey,
		Address:    addr,
	}
	table[lowestTwoBits] = append(table[lowestTwoBits], addrData)
	return lowestTwoBits
}

func nextPrivKey(oldKey []byte, psk []byte) []byte {
	var oldKeyInt big.Int
	oldKeyInt.SetBytes(oldKey)
	var pskInt big.Int
	pskInt.SetBytes(psk)

	newKey := new(big.Int).Mul(&oldKeyInt, &pskInt)
	newKey = newKey.Mod(newKey, secp256k1.S256().Params().N)
	return newKey.Bytes()
}

func newPrivKey() []byte {
	// 随机生成钥对
	// 生成一个随机的 ECDSA 密钥对
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// 将私钥转为字节切片
	privateKeyBytes := crypto.FromECDSA(privateKey)
	return privateKeyBytes
}

func SetAddrDatas(table [][]addrData, psk []byte, n int, times int) {
	// 初始化，循环计算times次并保存到table里
	for i := 0; i < times; i++ {
		key := newPrivKey()
		appendListByPriv(table, key, psk, n)
	}
}

func getAddrDatas(table [][]addrData, psk []byte, n int, ans byte) {
	// 递归计算新addr,直到符合条件ans
	key := newPrivKey()
	res := appendListByPriv(table, key, psk, n)
	if res != ans {
		getAddrDatas(table, psk, n, ans)
	}

}

func ShowAddrTable(table [][]addrData) {
	for i := range table {
		fmt.Println("============table[", i, "]==============")
		for j := range table[i] {
			fmt.Printf("publicKey :0x%x\n", table[i][j].PublicKey)
			fmt.Printf("privateKey:0x%x\n", table[i][j].PrivateKey)
			fmt.Println("Addr      :", table[i][j].Address)
			fmt.Println("---")
		}
	}
}

func FindAddr(table [][]addrData, psk []byte, msgInt int) addrData {
	// 根据msgInt查找对应的addrData
	result := table[msgInt]
	if len(result) == 0 {
		getAddrDatas(table, psk, 2, byte(msgInt))
	}
	popAddrData := table[msgInt][0]
	table[msgInt] = table[msgInt][1:]
	return popAddrData
}

//func getAddrDatas() {
//
//}
//
//func GetAddrByMsg(msg int, n int) addrData {
//
//}
