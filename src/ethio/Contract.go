package ethio

import (
	"EthCovertrans/src/ethio/contract"
	"EthCovertrans/src/ethio/util"
	"crypto/ecdsa"
	"crypto/elliptic"
	"log"
	"math/big"
)

func UpdateContract(pubkey *ecdsa.PublicKey) {
	// pubkey newPK
	// KeyData.PK oldPK

	GroupPK := util.PrivateKeyToAddrData(KeyData.Psk).Address
	myPK := util.PrivateKeyToAddrData(KeyData.Sender).PublicKey
	oldkey := contract.EllipticCurveKeyStorageECKey{
		X: myPK.X,
		Y: myPK.Y,
	}
	newkey := contract.EllipticCurveKeyStorageECKey{
		X: pubkey.X,
		Y: pubkey.Y,
	}
	v, r, s := util.SignMessage(KeyData.Psk, FaucetAc.Address.Bytes())

	_, err := EthContract.UpdateSenderPK(AuthTransact, GroupPK, oldkey, newkey, v, r, s)
	if err != nil {
		return
	}

}

func DiffKeyData() {
	localPK := *KeyData.Recvers
	contractPK := *GetRecvers(KeyData.Psk)
	if len(localPK) != len(contractPK) { // 新用户注册
		// append contractPK中新的公钥写入Recvers
		log.Print("[Sender] New User Register ... ")
		tmp := append(localPK, contractPK[len(localPK):]...)
		KeyData.Recvers = &tmp
	}
	for i := 0; i < len(localPK); i++ {
		if (localPK)[i].X.Cmp((contractPK)[i].X) != 0 || (localPK)[i].Y.Cmp((contractPK)[i].Y) != 0 {
			// 接收消息
			MsgRecverFactory(KeyData.Psk, localPK[i], contractPK[i])
			// 更新本地公钥
			(*KeyData.Recvers)[i].X = (contractPK)[i].X
			(*KeyData.Recvers)[i].Y = (contractPK)[i].Y
		}
	}
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

func RegisterRecv(psk *ecdsa.PrivateKey, Recv *ecdsa.PublicKey) {
	GroupPK := util.PrivateKeyToAddrData(psk).Address
	tmp := contract.EllipticCurveKeyStorageECKey{
		X: big.NewInt(0),
		Y: big.NewInt(0),
	}
	key := contract.EllipticCurveKeyStorageECKey{
		X: Recv.X,
		Y: Recv.Y,
	}
	v, r, s := util.SignMessage(KeyData.Psk, FaucetAc.Address.Bytes())

	_, err := EthContract.UpdateSenderPK(AuthTransact, GroupPK, tmp, key, v, r, s)
	if err != nil {
		return
	}
}
