package ethio

import (
	"EthCovertrans/src/ethio/contract"
	"EthCovertrans/src/ethio/util"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

func UpdateContract(pubkey *ecdsa.PublicKey) {
	// TODO: 更新合约公钥，更新本地公钥
	// pubkey newPK
	// KeyData.PK oldPK

}

func diffKeyData() {
	// TODO: 从合约获取公钥，与本地公钥比对
	localPK := KeyData.Recvers
	contractPK := GetRecvers(KeyData.Psk)
	if len(*localPK) != len(*contractPK) { // 新用户注册
		// TODO: 新用户注册 新的resrs
	}
	for i := 0; i < len(*localPK); i++ {
		if (*localPK)[i].X.Cmp((*contractPK)[i].X) != 0 || (*localPK)[i].Y.Cmp((*contractPK)[i].Y) != 0 {
			// TODO: 接收消息
		}
	}
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
