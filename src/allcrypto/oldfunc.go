package allcrypto

//
//import (
//	"crypto/ecdsa"
//	"fmt"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/crypto/secp256k1"
//	"math/big"
//)
//
//var addrList [][]addrData
//
///*
//addrList[00] = [data,data,data]
//addrList[01] = [data,data,data]
//addrList[10] = [data,data,data]
//addrList[11] = [data,data,data]
//data = [privKey,PubKey,addr]
//*/
//
//func InitAddrTable(n int) [][]addrData {
//	// n 交易嵌入位数
//	addrList = getListByN(n)
//	return addrList
//}
//
//func getPubKeyByPriv(privKey []byte) []byte {
//	// 通过私钥计算公钥（ pk = sk*G ）
//	privateKeyBytes := privKey
//	privateKey := new(ecdsa.PrivateKey)
//	privateKey.D = new(big.Int).SetBytes(privateKeyBytes)
//	privateKey.PublicKey.Curve = crypto.S256()
//	privateKey.PublicKey.X, privateKey.PublicKey.Y = privateKey.PublicKey.Curve.ScalarBaseMult(privateKey.D.Bytes())
//	publicKey := privateKey.PublicKey
//	return crypto.FromECDSAPub(&publicKey)
//}
//
//func getAddrByPub(pubKey []byte) string {
//	// 通过公钥计算钱包地址
//	p, _ := crypto.UnmarshalPubkey(pubKey)
//	address := crypto.PubkeyToAddress(*p)
//	return address.Hex()
//}
//
////func getHashByAddrPsk(addr string, psk string) [16]byte {
////	// 通过钱包地址计算交易哈希
////	data := []byte(addr + psk)
////	// 计算data的md5
////	hashT := md5.Sum(data)
////	return hashT
////}
//
//func getListByN(n int) [][]addrData {
//	sum := 1 << n                   // 2的n次幂
//	rows := make([][]addrData, sum) // 创建sum行切片
//	for i := range rows {
//		// 为每一行创建一个切片，可以根据需要指定初始长度
//		rows[i] = make([]addrData, 0) // 这里初始长度为 0
//	}
//	return rows
//}
//
//func appendListByPriv(table [][]addrData, privKey []byte, psk []byte, n int) [][]addrData {
//	// 通过哈希计算 地址集合addrList
//	pubKey := getPubKeyByPriv(privKey)
//	addr := getAddrByPub(pubKey)
//	hash := getHashByAddrPsk(addr, string(psk))
//	//fmt.Printf("%x\n", hash)
//	var result byte
//	result = (1 << n) - 1
//	lowestTwoBits := hash[15] & result
//	//fmt.Printf("%x\n", lowestTwoBits)
//	//fmt.Printf("%x\n", pubKey)
//	addrData := addrData{
//		privateKey: privKey,
//		publicKey:  pubKey,
//		address:    addr,
//	}
//	table[lowestTwoBits] = append(table[lowestTwoBits], addrData)
//	return table
//}
//
//func newPrivKey(oldKey []byte, psk []byte) []byte {
//	var oldKeyInt big.Int
//	oldKeyInt.SetBytes(oldKey)
//	var pskInt big.Int
//	pskInt.SetBytes(psk)
//
//	newKey := new(big.Int).Mul(&oldKeyInt, &pskInt)
//	newKey = newKey.Mod(newKey, secp256k1.S256().Params().N)
//	return newKey.Bytes()
//}
//
//func SetAddrDatas(table [][]addrData, privKey []byte, psk []byte, n int, times int) []byte {
//	// 循环计算times次并保存到table里，返回最后一次计算之后的下一次还没使用的私钥
//	key := privKey
//	for i := 0; i < times; i++ {
//		appendListByPriv(table, key, psk, n)
//		key = newPrivKey(key, psk)
//	}
//	return key
//}
//
//func ShowAddrTable(table [][]addrData) {
//	for i := range table {
//		fmt.Println("============table[", i, "]==============")
//		for j := range table[i] {
//			fmt.Printf("publicKey :0x%x\n", table[i][j].publicKey)
//			fmt.Printf("privateKey:0x%x\n", table[i][j].privateKey)
//			fmt.Println("Addr      :", table[i][j].address)
//			fmt.Println("---")
//		}
//	}
//}
//
////func getAddrDatas() {
////
////}
////
////func GetAddrByMsg(msg int, n int) addrData {
////
////}

//func TestEcdh() {
//	// ECDH（Elliptic Curve Diffie-Hellman）算法是一种基于椭圆曲线的密钥交换协议
//	// aliceKey和bobKey都是各自的私钥
//	// alicePubkey和bobPubkey是各自的公钥
//	// aliceShared和bobShared一致，为计算出的共享密钥（项目中的主密钥）
//	fmt.Println("[Test ECDH]=================================")
//	aliceKey, err := ecdh.P256().GenerateKey(rand.Reader)
//	if err != nil {
//		panic(err)
//	}
//	bobKey, err := ecdh.P256().GenerateKey(rand.Reader)
//	if err != nil {
//		panic(err)
//	}
//	alicePubkey := aliceKey.PublicKey()
//	shared, _ := bobKey.ECDH(alicePubkey)
//	bobShared := sha256.Sum256(shared)
//	fmt.Printf("秘钥哈希(Bob)  %x\n", bobShared)
//	// 秘钥哈希(Bob)  a74e7949e71ead5f3bd4de031e2ad45c3f5b80b48ccf50e50eb86f4bdb025c3a
//	bobPubkey := bobKey.PublicKey()
//	shared, _ = aliceKey.ECDH(bobPubkey)
//	aliceShared := sha256.Sum256(shared)
//	fmt.Printf("秘钥哈希(Alice)  %x\n", aliceShared)
//	// 秘钥哈希(Alice)  a74e7949e71ead5f3bd4de031e2ad45c3f5b80b48ccf50e50eb86f4bdb025c3a
//
//}
//
//func TestEcdhExchange() [32]byte {
//	fmt.Println("[Test Ecdh Key Exchange]=================================")
//	// Alice和Bob 公私钥对生成
//	aliceKey, err := ecdh.P256().GenerateKey(rand.Reader)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Alice私钥： %x\n", aliceKey.Bytes())
//	alicePubkey := aliceKey.PublicKey()
//	fmt.Printf("Alice公钥： %x\n", alicePubkey.Bytes())
//
//	bobKey, err := ecdh.P256().GenerateKey(rand.Reader)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("Bob私钥： %x\n", bobKey.Bytes())
//	bobPubkey := bobKey.PublicKey()
//	fmt.Printf("Bob公钥： %x\n", bobPubkey.Bytes())
//
//	// 密钥交换：双方交换公钥，Bob获得alicePubkey，Alice获得bobPubkey
//
//	// Bob利用自己的私钥（Bob私钥）和Alice公钥计算共享主密钥mainKey
//	shared, _ := bobKey.ECDH(alicePubkey)
//	bobShared := sha256.Sum256(shared)
//	fmt.Printf("Bob计算出的共享主密钥mainKey  %x\n", bobShared)
//
//	// Alice利用自己的私钥（Alice私钥）和Bob公钥计算共享主密钥mainKey
//	shared, _ = aliceKey.ECDH(bobPubkey)
//	aliceShared := sha256.Sum256(shared)
//	fmt.Printf("Alice计算出的共享主密钥mainKey  %x\n", aliceShared)
//
//	// 返回mainKey
//	return aliceShared
//}
//
//func GetMainKey() [32]byte {
//	// 获得主密钥（链下）
//
//	// Alice和Bob 公私钥对生成
//	aliceKey, err := ecdh.P256().GenerateKey(rand.Reader)
//	if err != nil {
//		panic(err)
//	}
//	//fmt.Printf("Alice私钥： %x\n", aliceKey.Bytes())
//	alicePubkey := aliceKey.PublicKey()
//	//fmt.Printf("Alice公钥： %x\n", alicePubkey.Bytes())
//
//	bobKey, err := ecdh.P256().GenerateKey(rand.Reader)
//	if err != nil {
//		panic(err)
//	}
//	//fmt.Printf("Bob私钥： %x\n", bobKey.Bytes())
//	bobPubkey := bobKey.PublicKey()
//	//fmt.Printf("Bob公钥： %x\n", bobPubkey.Bytes())
//
//	// 密钥交换：双方交换公钥，Bob获得alicePubkey，Alice获得bobPubkey
//
//	// Bob利用自己的私钥（Bob私钥）和Alice公钥计算共享主密钥mainKey
//	shared, _ := bobKey.ECDH(alicePubkey)
//	bobShared := sha256.Sum256(shared)
//	//fmt.Printf("Bob计算出的共享主密钥mainKey  %x\n", bobShared)
//
//	// Alice利用自己的私钥（Alice私钥）和Bob公钥计算共享主密钥mainKey
//	shared, _ = aliceKey.ECDH(bobPubkey)
//	aliceShared := sha256.Sum256(shared)
//	//fmt.Printf("Alice计算出的共享主密钥mainKey  %x\n", aliceShared)
//
//	// 返回mainKey
//	if bobShared == aliceShared {
//		return aliceShared
//	} else {
//		panic("密钥计算错误")
//	}
//}
//
//func GetTransKey(mainKey [32]byte, transactionHash string) []byte {
//	// 使用pbkdf2密钥派生算法计算每次的通信密钥tansKey
//
//	// transactionHash就是上次交易的Transaction Hash
//	// transactionHash因该是字符串类型，形如：
//	// '0x8879c312dd900030f514b7fe44abcf64bf7c825c02b9a0a9e3bbc723bab2d004'
//	if len(transactionHash) != 66 {
//		panic("Transaction Hash错误")
//	}
//	// token 为transactionHash的后16位
//	token := transactionHash[len(transactionHash)-16:]
//
//	newToken := make([]byte, len(token))
//	copy(newToken, token)
//	newMainKey := mainKey[:]
//	// pbkdf2密钥派生算法，迭代次数4，迭代方式Sha256,得到每次交易密钥长度32
//	transKey := pbkdf2.Key(newMainKey, newToken, 4, 32, sha256.New)
//	return transKey
//}
//
//func getMainKey() {
//	/*
//
//	 */
//
//}
