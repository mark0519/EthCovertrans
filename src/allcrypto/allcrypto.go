package allcrypto

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
