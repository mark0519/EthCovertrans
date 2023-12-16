// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EllipticCurveKeyStorageECKey is an auto generated low-level Go binding around an user-defined struct.
type EllipticCurveKeyStorageECKey struct {
	X *big.Int
	Y *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupPK\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structEllipticCurveKeyStorage.ECKey\",\"name\":\"key\",\"type\":\"tuple\"}],\"name\":\"addGroupPK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupPK\",\"type\":\"address\"}],\"name\":\"getSenderPK\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structEllipticCurveKeyStorage.ECKey[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"groupPKToSenderPK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"groupPK\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structEllipticCurveKeyStorage.ECKey\",\"name\":\"oldKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structEllipticCurveKeyStorage.ECKey\",\"name\":\"newKey\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"updateSenderPK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610eda8061005c5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80634d8c4822146100595780637701f0b6146100755780638da5cb5b146100a6578063f1e47084146100c4578063fc6603ce146100e0575b5f80fd5b610073600480360381019061006e919061094e565b610110565b005b61008f600480360381019061008a91906109d8565b6103fc565b60405161009d929190610a25565b60405180910390f35b6100ae610436565b6040516100bb9190610a5b565b60405180910390f35b6100de60048036038101906100d99190610a74565b610459565b005b6100fa60048036038101906100f59190610ab2565b610562565b6040516101079190610bc1565b60405180910390f35b5f60015f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905011610192576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018990610c61565b60405180910390fd5b5f336040516020016101a49190610cc4565b6040516020818303038152906040528051906020012090505f6001828686866040515f81526020016040526040516101df9493929190610cfc565b6020604051602081039080840390855afa1580156101ff573d5f803e3d5ffd5b5050506020604051035190508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610279576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027090610d89565b60405180910390fd5b5f875f015114801561028e57505f8760200151145b156103105760015f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2086908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550506103f2565b5f61031b898961060e565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810361037f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037690610df1565b60405180910390fd5b8660015f8b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082815481106103cf576103ce610e0f565b5b905f5260205f2090600202015f820151815f015560208201518160010155905050505b5050505050505050565b6001602052815f5260405f208181548110610415575f80fd5b905f5260205f2090600202015f9150915050805f0154908060010154905082565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd90610e86565b60405180910390fd5b60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f01556020820151816001015550505050565b606060015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610603578382905f5260205f2090600202016040518060400160405290815f8201548152602001600182015481525050815260200190600101906105c0565b505050509050919050565b5f805f90505b60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208054905081101561074757825f015160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2082815481106106af576106ae610e0f565b5b905f5260205f2090600202015f015414801561072c5750826020015160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20828154811061071a57610719610e0f565b5b905f5260205f20906002020160010154145b1561073a578091505061076c565b8080600101915050610614565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90505b92915050565b5f604051905090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107a88261077f565b9050919050565b6107b88161079e565b81146107c2575f80fd5b50565b5f813590506107d3816107af565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610823826107dd565b810181811067ffffffffffffffff82111715610842576108416107ed565b5b80604052505050565b5f610854610772565b9050610860828261081a565b919050565b5f819050919050565b61087781610865565b8114610881575f80fd5b50565b5f813590506108928161086e565b92915050565b5f604082840312156108ad576108ac6107d9565b5b6108b7604061084b565b90505f6108c684828501610884565b5f8301525060206108d984828501610884565b60208301525092915050565b5f60ff82169050919050565b6108fa816108e5565b8114610904575f80fd5b50565b5f81359050610915816108f1565b92915050565b5f819050919050565b61092d8161091b565b8114610937575f80fd5b50565b5f8135905061094881610924565b92915050565b5f805f805f8061010087890312156109695761096861077b565b5b5f61097689828a016107c5565b965050602061098789828a01610898565b955050606061099889828a01610898565b94505060a06109a989828a01610907565b93505060c06109ba89828a0161093a565b92505060e06109cb89828a0161093a565b9150509295509295509295565b5f80604083850312156109ee576109ed61077b565b5b5f6109fb858286016107c5565b9250506020610a0c85828601610884565b9150509250929050565b610a1f81610865565b82525050565b5f604082019050610a385f830185610a16565b610a456020830184610a16565b9392505050565b610a558161079e565b82525050565b5f602082019050610a6e5f830184610a4c565b92915050565b5f8060608385031215610a8a57610a8961077b565b5b5f610a97858286016107c5565b9250506020610aa885828601610898565b9150509250929050565b5f60208284031215610ac757610ac661077b565b5b5f610ad4848285016107c5565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b0f81610865565b82525050565b604082015f820151610b295f850182610b06565b506020820151610b3c6020850182610b06565b50505050565b5f610b4d8383610b15565b60408301905092915050565b5f602082019050919050565b5f610b6f82610add565b610b798185610ae7565b9350610b8483610af7565b805f5b83811015610bb4578151610b9b8882610b42565b9750610ba683610b59565b925050600181019050610b87565b5085935050505092915050565b5f6020820190508181035f830152610bd98184610b65565b905092915050565b5f82825260208201905092915050565b7f4e6f20636f72726573706f6e64696e672045434b657920666f722067726f75705f8201527f504b000000000000000000000000000000000000000000000000000000000000602082015250565b5f610c4b602283610be1565b9150610c5682610bf1565b604082019050919050565b5f6020820190508181035f830152610c7881610c3f565b9050919050565b5f8160601b9050919050565b5f610c9582610c7f565b9050919050565b5f610ca682610c8b565b9050919050565b610cbe610cb98261079e565b610c9c565b82525050565b5f610ccf8284610cad565b60148201915081905092915050565b610ce78161091b565b82525050565b610cf6816108e5565b82525050565b5f608082019050610d0f5f830187610cde565b610d1c6020830186610ced565b610d296040830185610cde565b610d366060830184610cde565b95945050505050565b7f5369676e617475726520766572696669636174696f6e206661696c65640000005f82015250565b5f610d73601d83610be1565b9150610d7e82610d3f565b602082019050919050565b5f6020820190508181035f830152610da081610d67565b9050919050565b7f4f6c642045434b6579206e6f7420666f756e64000000000000000000000000005f82015250565b5f610ddb601383610be1565b9150610de682610da7565b602082019050919050565b5f6020820190508181035f830152610e0881610dcf565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e6f7420746865206f776e6572000000000000000000000000000000000000005f82015250565b5f610e70600d83610be1565b9150610e7b82610e3c565b602082019050919050565b5f6020820190508181035f830152610e9d81610e64565b905091905056fea2646970667358221220d5b758df89cf7a712904926d9e31661572f9a043ae2a9354905e7ad425823fec64736f6c63430008170033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetSenderPK is a free data retrieval call binding the contract method 0xfc6603ce.
//
// Solidity: function getSenderPK(address groupPK) view returns((uint256,uint256)[])
func (_Contract *ContractCaller) GetSenderPK(opts *bind.CallOpts, groupPK common.Address) ([]EllipticCurveKeyStorageECKey, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSenderPK", groupPK)

	if err != nil {
		return *new([]EllipticCurveKeyStorageECKey), err
	}

	out0 := *abi.ConvertType(out[0], new([]EllipticCurveKeyStorageECKey)).(*[]EllipticCurveKeyStorageECKey)

	return out0, err

}

// GetSenderPK is a free data retrieval call binding the contract method 0xfc6603ce.
//
// Solidity: function getSenderPK(address groupPK) view returns((uint256,uint256)[])
func (_Contract *ContractSession) GetSenderPK(groupPK common.Address) ([]EllipticCurveKeyStorageECKey, error) {
	return _Contract.Contract.GetSenderPK(&_Contract.CallOpts, groupPK)
}

// GetSenderPK is a free data retrieval call binding the contract method 0xfc6603ce.
//
// Solidity: function getSenderPK(address groupPK) view returns((uint256,uint256)[])
func (_Contract *ContractCallerSession) GetSenderPK(groupPK common.Address) ([]EllipticCurveKeyStorageECKey, error) {
	return _Contract.Contract.GetSenderPK(&_Contract.CallOpts, groupPK)
}

// GroupPKToSenderPK is a free data retrieval call binding the contract method 0x7701f0b6.
//
// Solidity: function groupPKToSenderPK(address , uint256 ) view returns(uint256 x, uint256 y)
func (_Contract *ContractCaller) GroupPKToSenderPK(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "groupPKToSenderPK", arg0, arg1)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GroupPKToSenderPK is a free data retrieval call binding the contract method 0x7701f0b6.
//
// Solidity: function groupPKToSenderPK(address , uint256 ) view returns(uint256 x, uint256 y)
func (_Contract *ContractSession) GroupPKToSenderPK(arg0 common.Address, arg1 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Contract.Contract.GroupPKToSenderPK(&_Contract.CallOpts, arg0, arg1)
}

// GroupPKToSenderPK is a free data retrieval call binding the contract method 0x7701f0b6.
//
// Solidity: function groupPKToSenderPK(address , uint256 ) view returns(uint256 x, uint256 y)
func (_Contract *ContractCallerSession) GroupPKToSenderPK(arg0 common.Address, arg1 *big.Int) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _Contract.Contract.GroupPKToSenderPK(&_Contract.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// AddGroupPK is a paid mutator transaction binding the contract method 0xf1e47084.
//
// Solidity: function addGroupPK(address groupPK, (uint256,uint256) key) returns()
func (_Contract *ContractTransactor) AddGroupPK(opts *bind.TransactOpts, groupPK common.Address, key EllipticCurveKeyStorageECKey) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addGroupPK", groupPK, key)
}

// AddGroupPK is a paid mutator transaction binding the contract method 0xf1e47084.
//
// Solidity: function addGroupPK(address groupPK, (uint256,uint256) key) returns()
func (_Contract *ContractSession) AddGroupPK(groupPK common.Address, key EllipticCurveKeyStorageECKey) (*types.Transaction, error) {
	return _Contract.Contract.AddGroupPK(&_Contract.TransactOpts, groupPK, key)
}

// AddGroupPK is a paid mutator transaction binding the contract method 0xf1e47084.
//
// Solidity: function addGroupPK(address groupPK, (uint256,uint256) key) returns()
func (_Contract *ContractTransactorSession) AddGroupPK(groupPK common.Address, key EllipticCurveKeyStorageECKey) (*types.Transaction, error) {
	return _Contract.Contract.AddGroupPK(&_Contract.TransactOpts, groupPK, key)
}

// UpdateSenderPK is a paid mutator transaction binding the contract method 0x4d8c4822.
//
// Solidity: function updateSenderPK(address groupPK, (uint256,uint256) oldKey, (uint256,uint256) newKey, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactor) UpdateSenderPK(opts *bind.TransactOpts, groupPK common.Address, oldKey EllipticCurveKeyStorageECKey, newKey EllipticCurveKeyStorageECKey, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSenderPK", groupPK, oldKey, newKey, v, r, s)
}

// UpdateSenderPK is a paid mutator transaction binding the contract method 0x4d8c4822.
//
// Solidity: function updateSenderPK(address groupPK, (uint256,uint256) oldKey, (uint256,uint256) newKey, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractSession) UpdateSenderPK(groupPK common.Address, oldKey EllipticCurveKeyStorageECKey, newKey EllipticCurveKeyStorageECKey, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSenderPK(&_Contract.TransactOpts, groupPK, oldKey, newKey, v, r, s)
}

// UpdateSenderPK is a paid mutator transaction binding the contract method 0x4d8c4822.
//
// Solidity: function updateSenderPK(address groupPK, (uint256,uint256) oldKey, (uint256,uint256) newKey, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contract *ContractTransactorSession) UpdateSenderPK(groupPK common.Address, oldKey EllipticCurveKeyStorageECKey, newKey EllipticCurveKeyStorageECKey, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSenderPK(&_Contract.TransactOpts, groupPK, oldKey, newKey, v, r, s)
}
