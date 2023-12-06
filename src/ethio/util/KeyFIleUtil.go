package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"log"
	"os"
)

type KeyFileData struct {
	Psk     *ecdsa.PrivateKey
	Sender  *ecdsa.PrivateKey
	Recvers *[]*ecdsa.PublicKey
}

// MarshalJSON 实现自定义的 JSON 序列化方法
func (kfd *KeyFileData) MarshalJSON() ([]byte, error) {
	type Alias KeyFileData
	b1, _ := x509.MarshalECPrivateKey(kfd.Psk)

	pskBlock := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: b1,
	}
	pskPEM := pem.EncodeToMemory(&pskBlock)
	b2, _ := x509.MarshalECPrivateKey(kfd.Sender)
	senderBlock := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: b2,
	}
	senderPEM := pem.EncodeToMemory(&senderBlock)

	var recversPEM []string
	if kfd.Recvers != nil {
		for _, pubKey := range *kfd.Recvers {
			pubASN1, err := x509.MarshalPKIXPublicKey(pubKey)
			if err != nil {
				return nil, err
			}
			pubBlock := pem.Block{
				Type:  "EC PUBLIC KEY",
				Bytes: pubASN1,
			}
			recversPEM = append(recversPEM, string(pem.EncodeToMemory(&pubBlock)))
		}
	}

	return json.Marshal(&struct {
		Psk     string   `json:"psk"`
		Sender  string   `json:"sender"`
		Recvers []string `json:"recvers"`
	}{
		Psk:     string(pskPEM),
		Sender:  string(senderPEM),
		Recvers: recversPEM,
	})
}

// UnmarshalJSON 实现自定义的 JSON 反序列化方法
func (kfd *KeyFileData) UnmarshalJSON(data []byte) error {
	type Alias KeyFileData
	aux := &struct {
		Psk     string   `json:"psk"`
		Sender  string   `json:"sender"`
		Recvers []string `json:"recvers"`
	}{}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	pskBlock, _ := pem.Decode([]byte(aux.Psk))
	pskKey, err := x509.ParseECPrivateKey(pskBlock.Bytes)
	if err != nil {
		return err
	}
	kfd.Psk = pskKey

	senderBlock, _ := pem.Decode([]byte(aux.Sender))
	senderKey, err := x509.ParseECPrivateKey(senderBlock.Bytes)
	if err != nil {
		return err
	}
	kfd.Sender = senderKey

	var recvers []*ecdsa.PublicKey
	for _, recver := range aux.Recvers {
		block, _ := pem.Decode([]byte(recver))
		if block == nil {
			return fmt.Errorf("error decoding PEM block")
		}
		pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return err
		}
		ecPubKey, ok := pubKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error converting to ECDSA public key")
		}
		recvers = append(recvers, ecPubKey)
	}
	kfd.Recvers = &recvers

	return nil
}

func aesEncrypt(data []byte, key []byte) []byte {
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

	// 将结构体转换为JSON格式的[]byte
	keyDataBytes, err := json.Marshal(keyFile)
	if err != nil {
		log.Panic("[Sender] Error marshaling:", err)
	}
	// 32字节的AES密钥（AES-256）
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
	aesKey := []byte("12345678901234567890123456789012")
	keyDataBytes := aesDecrypt(fileBytes, aesKey)
	if err != nil {
		log.Panic("[Sender] AES Decrypt Error:", err)
	}
	// 创建一个 KeyFileData 实例
	keyFile := new(KeyFileData)

	// 使用 json.Unmarshal 将 JSON 格式的字节切片转换回 KeyFileData 结构体
	err = json.Unmarshal(keyDataBytes, &keyFile)
	if err != nil {
		log.Panic("[Sender] Error unmarshaling:", err)
	}
	return *keyFile
}

func GetPskFromFile() *ecdsa.PrivateKey {
	// 读取psk文件
	pskData, err := os.ReadFile("psk.txt")
	if err != nil {
		log.Fatal("[Sender] Not Found psk.txt(~/EthCovertrans/psk.txt) or psk.txt is empty")
	}
	// 将读取的字节切片转换为十六进制字符串
	pskHexStr := string(pskData)
	psk := crypto.ToECDSAUnsafe(common.FromHex(pskHexStr))
	//fmt.Print(psk)
	return psk
}
