package ethio

import (
	"EthCovertrans/src/ethio/contract"
	"EthCovertrans/src/ethio/util"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"math/big"
)

func UpdateContract(newSendAddrData util.SendAddrData) {
	// pubkey newPK
	// KeyData.PK oldPK

	GroupPK := util.PrivateKeyToAddrData(KeyData.Psk).Address
	myPK := util.PrivateKeyToAddrData(KeyData.Sender).PublicKey
	oldkey := contract.EllipticCurveKeyStorageECKey{
		X: myPK.X,
		Y: myPK.Y,
	}
	newkey := contract.EllipticCurveKeyStorageECKey{
		X: newSendAddrData.PublicKey.X,
		Y: newSendAddrData.PublicKey.Y,
	}
	v, r, s := util.SignMessage(KeyData.Psk, FaucetAc.Address.Bytes())

	nonce, _ := Client.NonceAt(context.Background(), FaucetAc.Address, nil)
	AuthTransact.Nonce = new(big.Int).SetUint64(nonce)
	_, err := EthContract.UpdateSenderPK(AuthTransact, GroupPK, oldkey, newkey, v, r, s)
	if err != nil {
		log.Fatal("[Sender] UpdateContract Error: ", err)
	}
	log.Printf("[Sender] UpdateContract Success")
	KeyData.Sender = newSendAddrData.GetSendAddrDataPrivateKey()
	keyDataF := util.KeyFileDataAndFaucet{
		KeyFileData: *KeyData,
		Faucet:      FaucetPrivatekeyStr,
	}
	util.EncryptKeyFileData(keyDataF, "ethCoverTrans.key")
	log.Printf("[Sender] Local sendersk update Success")
}

func ForceUpdateContract() {
	// 仅供调试使用

	id := 2 // 本地私钥对应公钥在合约上的位置
	GroupPK := util.PrivateKeyToAddrData(KeyData.Psk).Address
	ctkey := (*GetRecvers(KeyData.Psk))[id]
	myPK := util.PrivateKeyToAddrData(KeyData.Sender).PublicKey
	oldkey := contract.EllipticCurveKeyStorageECKey{
		X: ctkey.X,
		Y: ctkey.Y,
	}
	newkey := contract.EllipticCurveKeyStorageECKey{
		X: myPK.X,
		Y: myPK.Y,
	}
	v, r, s := util.SignMessage(KeyData.Psk, FaucetAc.Address.Bytes())

	_, err := EthContract.UpdateSenderPK(AuthTransact, GroupPK, oldkey, newkey, v, r, s)
	if err != nil {
		log.Fatal("[Sender] UpdateContract Error: ", err)
	}
	log.Printf("[Sender] UpdateContract Success ... ")
}

func ForceUpdateLocal() {
	localPK := *KeyData.Recvers
	contractPK := *GetRecvers(KeyData.Psk)
	if len(localPK) != len(contractPK) { // 新用户注册
		// append contractPK中新的公钥写入Recvers
		tmp := append(localPK, contractPK[len(localPK):]...)
		KeyData.Recvers = &tmp
		log.Print("[Sender] New User Register")
	}
	for i := 0; i < len(localPK); i++ {
		if (localPK)[i].X.Cmp((contractPK)[i].X) != 0 || (localPK)[i].Y.Cmp((contractPK)[i].Y) != 0 {
			// 接收消息
			(*KeyData.Recvers)[i].X = (contractPK)[i].X
			(*KeyData.Recvers)[i].Y = (contractPK)[i].Y
		}
	}
	keyDataF := util.KeyFileDataAndFaucet{
		KeyFileData: *KeyData,
		Faucet:      FaucetPrivatekeyStr,
	}
	util.EncryptKeyFileData(keyDataF, "ethCoverTrans.key")
	log.Printf("[Sender] Force update KeyData Success ... ")
}

func DiffKeyData() {
	localPK := *KeyData.Recvers
	contractPK := *GetRecvers(KeyData.Psk)
	if len(localPK) != len(contractPK) { // 新用户注册
		// append contractPK中新的公钥写入Recvers
		tmp := append(localPK, contractPK[len(localPK):]...)
		KeyData.Recvers = &tmp
		log.Print("[Sender] New User Register")
	}
	for i := 0; i < len(localPK); i++ {
		if (localPK)[i].X.Cmp((contractPK)[i].X) != 0 || (localPK)[i].Y.Cmp((contractPK)[i].Y) != 0 {
			log.Println("[Sender] New Message!!")
			// 接收消息
			MsgRecverFactory(KeyData.Psk, localPK[i], contractPK[i])
			log.Println("[Sender] Recv Msg Success")
			// 更新本地公钥
			(*KeyData.Recvers)[i].X = (contractPK)[i].X
			(*KeyData.Recvers)[i].Y = (contractPK)[i].Y
			log.Println("[Sender] DiffKeyData Update Recvers")
		}
	}
	log.Printf("[Sender] DiffKeyData Success ... ")
	keyDataF := util.KeyFileDataAndFaucet{
		KeyFileData: *KeyData,
		Faucet:      FaucetPrivatekeyStr,
	}
	util.EncryptKeyFileData(keyDataF, "ethCoverTrans.key")
}

func addGroupPK4owner(psk *ecdsa.PrivateKey, ownerPk *ecdsa.PublicKey) {
	GroupPK := util.PrivateKeyToAddrData(psk).Address
	key := contract.EllipticCurveKeyStorageECKey{
		X: ownerPk.X,
		Y: ownerPk.Y,
	}

	nonce, _ := Client.NonceAt(context.Background(), FaucetAc.Address, nil)
	AuthTransact.Nonce = new(big.Int).SetUint64(nonce)
	_, err := EthContract.AddGroupPK(AuthTransact, GroupPK, key)
	if err != nil {
		return
	}
}

func GetRecvers(psk *ecdsa.PrivateKey) *[]*ecdsa.PublicKey {
	GroupPK := util.PrivateKeyToAddrData(psk).Address
	SenderPk, err := EthContract.GetSenderPK(nil, GroupPK)
	if err != nil {
		return nil
	}
	Recvers := make([]*ecdsa.PublicKey, len(SenderPk))
	for i := 0; i < len(SenderPk); i++ {
		Recvers[i] = &ecdsa.PublicKey{
			X:     SenderPk[i].X,
			Y:     SenderPk[i].Y,
			Curve: elliptic.P256(),
		}
	}
	return &Recvers
}

func RegisterRecv() {
	GroupPK := util.PrivateKeyToAddrData(KeyData.Psk).Address
	tmp := contract.EllipticCurveKeyStorageECKey{
		X: big.NewInt(0),
		Y: big.NewInt(0),
	}
	myPK := util.PrivateKeyToAddrData(KeyData.Sender).PublicKey
	key := contract.EllipticCurveKeyStorageECKey{
		X: myPK.X,
		Y: myPK.Y,
	}
	v, r, s := util.SignMessage(KeyData.Psk, FaucetAc.Address.Bytes())

	nonce, _ := Client.NonceAt(context.Background(), FaucetAc.Address, nil)
	AuthTransact.Nonce = new(big.Int).SetUint64(nonce)
	_, err := EthContract.UpdateSenderPK(AuthTransact, GroupPK, tmp, key, v, r, s)
	if err != nil {
		return
	}
	log.Print("[Sender] Register Success ... ")
}
