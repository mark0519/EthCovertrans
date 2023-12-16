package util

import (
	"EthCovertrans/src/ethio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"math/big"
	"os"
)

// 文件加密密钥（口令的md5）
var FileAesKey []byte

type KeyFileDataAndFaucet struct {
	KeyFileData KeyFileData
	Faucet      string
}

type KeyFileData struct {
	Psk     *ecdsa.PrivateKey
	Sender  *ecdsa.PrivateKey
	Recvers *[]*ecdsa.PublicKey
}
type XY struct {
	X *big.Int
	Y *big.Int
}
type KeyFileDataForD struct {
	PskD    *big.Int
	SenderD *big.Int
	Recvers *[]XY
}

func dXY2KeyFileData(dKeyFileData *KeyFileDataForD) *KeyFileData {
	// 将KeyFileDataForD转换为KeyFileData

	psk := D2PrivateKey(dKeyFileData.PskD)
	sender := D2PrivateKey(dKeyFileData.SenderD)
	recvers := make([]*ecdsa.PublicKey, len(*dKeyFileData.Recvers))
	for i, xy := range *dKeyFileData.Recvers {
		recvers[i] = XY2PublicKey(xy.X, xy.Y)
	}
	return &KeyFileData{
		Psk:     psk,
		Sender:  sender,
		Recvers: &recvers,
	}
}

func keyFileData2DXY(keyFileData *KeyFileData) *KeyFileDataForD {
	// 将KeyFileData转换为KeyFileDataForD

	recvers := make([]XY, len(*keyFileData.Recvers))
	for i, r := range *keyFileData.Recvers {
		recvers[i].X = r.X
		recvers[i].Y = r.Y
	}
	return &KeyFileDataForD{
		PskD:    keyFileData.Psk.D,
		SenderD: keyFileData.Sender.D,
		Recvers: &recvers,
	}
}

func aesEncrypt(data []byte, key []byte) []byte {
	// AES加密
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext
}

func aesDecrypt(ciphertext []byte, key []byte) []byte {
	// AES解密
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	if len(ciphertext) < aes.BlockSize {
		log.Fatal("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

func EncryptKeyFileData(keyFileF KeyFileDataAndFaucet, path string) {
	// 加密KeyFileData存入文件

	// keyFileData转换为KeyFileDataForD
	keyFileForD := keyFileData2DXY(&keyFileF.KeyFileData)

	// 将结构体转换为JSON格式的[]byte
	keyDataBytes, err := json.Marshal(keyFileForD)
	if err != nil {
		log.Fatal("[Sender] Error marshaling:", err)
	}

	encData := append([]byte(keyFileF.Faucet), keyDataBytes...)

	// 32字节的AES密钥（AES-256）
	encrypted := aesEncrypt(encData, FileAesKey)
	//fmt.Print(encrypted)
	// 写入文件

	// 计算FileAesKey的sha256并取后4字节作为加密文件头MagicHeader
	hasher := sha256.New()
	hasher.Write(FileAesKey)
	MagicHeader := hasher.Sum(nil)
	MagicHeader = MagicHeader[len(MagicHeader)-4:]

	// 合并MagicHeader和加密后的数据
	data := append(MagicHeader, encrypted...)

	// 写入文件
	err = os.WriteFile(path, data, 0777)
	if err != nil {
		log.Fatal("Error writing file:", err)
	}
}

func DecryptKeyFileData(path string) KeyFileDataAndFaucet {
	// 从文件读取KeyFileData并解密

	// 读取配置文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open config file:", err)
	}
	defer file.Close()

	// 获取文件大小
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Failed to get file size:", err)
	}

	// 分配足够的空间来存储文件内容
	fileBytes := make([]byte, stat.Size())

	// 读取文件内容
	_, err = io.ReadFull(file, fileBytes)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}
	//fmt.Print("===================================")
	//fmt.Print(fileBytes)

	// 检查口令
	hasher := sha256.New()
	hasher.Write(FileAesKey)
	MagicHeader := hasher.Sum(nil)
	MagicHeader = MagicHeader[len(MagicHeader)-4:]
	equal := bytes.Equal(MagicHeader, fileBytes[:4])
	if !equal {
		log.Fatal("[Sender] Wrong File Password")
	}
	// 解密 32字节的AES密钥（AES-256）
	encData := fileBytes[4:]
	keyDataBytes := aesDecrypt(encData, FileAesKey)
	if err != nil {
		log.Fatal("[Sender] AES Decrypt Error:", err)
	}

	// 从解密后的数据中分离出水龙头私钥
	faucetPKBytes := keyDataBytes[:64]
	faucetPK := string(faucetPKBytes)

	// 创建一个 KeyFileDataForD 实例
	keyFileD := new(KeyFileDataForD)

	// 使用 json.Unmarshal 将 JSON 格式的字节切片转换回 KeyFileData 结构体
	err = json.Unmarshal(keyDataBytes[64:], &keyFileD)
	if err != nil {
		log.Fatal("[Sender] Error unmarshaling:", err)
	}

	// 将KeyFileDataForD转换为KeyFileData
	keyFile := dXY2KeyFileData(keyFileD)

	// 返回 KeyFileDataAndFaucet
	KeyFileAndF := new(KeyFileDataAndFaucet)
	KeyFileAndF.KeyFileData = *keyFile
	KeyFileAndF.Faucet = faucetPK
	return *KeyFileAndF
}

func GenerateKeyFile(pskFileName string, privateKeyFileName string, KeyFile string) KeyFileDataAndFaucet {
	// 初始化KeyFile并返回水龙头私钥
	// 如果已经存在ethCoverTrans.key，则读取psk和sender私钥, 返回 KeyFileData
	// 如果不存在ethCoverTrans.key，则从psk.key和privateKey.key获取psk和sender私钥，并生成加密密钥文件ethCoverTrans.key。同时删除psk.key,返回 KeyFileData
	// 返回 KeyFileDataAndFaucet

	// 使用 os.Stat 检查文件是否存在
	if _, err := os.Stat("ethCoverTrans.key"); err == nil { // 如果文件存在
		log.Print("[Sender] Loading init key file: ethCoverTrans.key ...")
		// 读取文件
		keyFileDataF := DecryptKeyFileData(KeyFile)
		// 返回 psk, senderPrivateKey, FaucetPrivateKey
		return keyFileDataF
	} else if os.IsNotExist(err) { // 如果文件不存在,则从psk.key初始化
		log.Print("[Sender] NO ethCoverTrans.key, Loading init pskFile&privateKeyFile ...")
		// 读取psk文件
		pskData, err := os.ReadFile(pskFileName)
		if err != nil {
			log.Fatal("[Sender] Not Found init psk file: ", pskFileName)
		}
		// 将读取的字节切片转换为十六进制字符串
		pskHexStr := string(pskData)
		psk := crypto.ToECDSAUnsafe(common.FromHex(pskHexStr))

		// 读取本人私钥
		privateKeyData, err := os.ReadFile(privateKeyFileName)
		if err != nil {
			log.Fatal("[Sender] Not Found init privateKey file: ", privateKeyFileName)
		}
		// 将读取的字节切片转换为十六进制字符串
		privateKeyHexStr := string(privateKeyData)
		privateKey := crypto.ToECDSAUnsafe(common.FromHex(privateKeyHexStr))

		var recvers []*ecdsa.PublicKey

		// 生成初始化ethCoverTrans.key文件
		keyData := KeyFileData{
			Psk:     psk,
			Sender:  privateKey,
			Recvers: &recvers, // 公钥列表
		}
		// 输入水龙头私钥FaucetPrivatekey
		fmt.Print("[Sender] Please Enter the ETH Faucet Private Key: ")
		var FaucetPrivatekeyStr string
		_, err = fmt.Scanln(&FaucetPrivatekeyStr)
		if err != nil {
			log.Fatal("[Sender] Error reading Faucet Private Key:", err)
		}
		if len(FaucetPrivatekeyStr) != 64 {
			log.Fatal("[Sender] Error reading Faucet Private Key: length error")
		}

		keyDataF := KeyFileDataAndFaucet{
			KeyFileData: keyData,
			Faucet:      FaucetPrivatekeyStr,
		}

		EncryptKeyFileData(keyDataF, KeyFile) // 加密并保存
		log.Print("[Sender] Generate ethCoverTrans.key Success")

		// 删除初始化psk文件
		err = os.Remove(pskFileName)
		if err != nil {
			log.Fatal("[Sender] Remove psk.key Err:", err)
		}
		log.Print("[Sender] Remove psk.key Success")

		// 注册本人公钥
		ethio.RegisterRecv(keyData.Psk, PrivateKeyToAddrData(keyData.Sender).PublicKey)

		return keyDataF
	} else {
		log.Fatal("[Sender] Unknown Init File Error:", err)
	}
	return KeyFileDataAndFaucet{} // 不可到达
}
