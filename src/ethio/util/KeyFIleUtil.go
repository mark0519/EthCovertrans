package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"math/big"
	"os"
)

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
		log.Panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext
}

func aesDecrypt(ciphertext []byte, key []byte) []byte {
	// AES解密
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		log.Panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

func EncryptKeyFileData(keyFile KeyFileData, path string) {
	// 加密KeyFileData存入文件

	// keyFileData转换为KeyFileDataForD
	keyFileForD := keyFileData2DXY(&keyFile)

	// 将结构体转换为JSON格式的[]byte
	keyDataBytes, err := json.Marshal(keyFileForD)
	if err != nil {
		log.Panic("[Sender] Error marshaling:", err)
	}
	// 32字节的AES密钥（AES-256）
	// TODO: 修改密钥获取方式
	aesKey := []byte("12345678901234567890123456789012")

	encrypted := aesEncrypt(keyDataBytes, aesKey)
	//fmt.Print(encrypted)
	// 写入文件
	err = os.WriteFile(path, encrypted, 0777)
	if err != nil {
		log.Panic("Error writing file:", err)
	}
}

func DecryptKeyFileData(path string) KeyFileData {
	// 从文件读取KeyFileData并解密

	// 读取配置文件
	file, err := os.Open(path)
	if err != nil {
		log.Panic("Failed to open config file:", err)
	}
	defer file.Close()

	// 获取文件大小
	stat, err := file.Stat()
	if err != nil {
		log.Panic("Failed to get file size:", err)
	}

	// 分配足够的空间来存储文件内容
	fileBytes := make([]byte, stat.Size())

	// 读取文件内容
	_, err = io.ReadFull(file, fileBytes)
	if err != nil {
		log.Panic("Failed to read file:", err)
	}
	//fmt.Print("===================================")
	//fmt.Print(fileBytes)
	// 解密
	// 32字节的AES密钥（AES-256）
	// TODO: 修改密钥获取方式
	aesKey := []byte("12345678901234567890123456789012")
	keyDataBytes := aesDecrypt(fileBytes, aesKey)
	if err != nil {
		log.Panic("[Sender] AES Decrypt Error:", err)
	}
	// 创建一个 KeyFileDataForD 实例
	keyFileD := new(KeyFileDataForD)

	// 使用 json.Unmarshal 将 JSON 格式的字节切片转换回 KeyFileData 结构体
	err = json.Unmarshal(keyDataBytes, &keyFileD)
	if err != nil {
		log.Panic("[Sender] Error unmarshaling:", err)
	}

	// 将KeyFileDataForD转换为KeyFileData
	keyFile := dXY2KeyFileData(keyFileD)
	return *keyFile
}

func GenerateKeyFile(fileName string) *ecdsa.PrivateKey {
	// 首次使用初始化，生成加密密钥文件
	// 读取psk文件
	pskData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("[Sender] Not Found init file")
	}
	// TODO: 读取本人私钥
	// 将读取的字节切片转换为十六进制字符串
	pskHexStr := string(pskData)
	psk := crypto.ToECDSAUnsafe(common.FromHex(pskHexStr))
	// TODO: 删除初始化文件
	// TODO: 与合约交互，注册本人公钥
	return psk
}
