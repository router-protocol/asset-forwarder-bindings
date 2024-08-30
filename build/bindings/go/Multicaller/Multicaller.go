// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultiCaller

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

// MultiCallerMetaData contains all meta data concerning the MultiCaller contract.
var MultiCallerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061058b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ac9650d814610030575b600080fd5b61004361003e3660046101d4565b610059565b60405161005091906102b7565b60405180910390f35b60608167ffffffffffffffff81111561007457610074610337565b6040519080825280602002602001820160405280156100a757816020015b60608152602001906001900390816100925790505b50905060005b828110156101cd57600080308686858181106100cb576100cb610366565b90506020028101906100dd9190610395565b6040516100eb929190610401565b600060405180830381855af49150503d8060008114610126576040519150601f19603f3d011682016040523d82523d6000602084013e61012b565b606091505b50915091508161019a5760448151101561014457600080fd5b6004810190508080602001905181019061015e9190610411565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019191906104dc565b60405180910390fd5b808484815181106101ad576101ad610366565b6020026020010181905250505080806101c5906104f6565b9150506100ad565b5092915050565b600080602083850312156101e757600080fd5b823567ffffffffffffffff808211156101ff57600080fd5b818501915085601f83011261021357600080fd5b81358181111561022257600080fd5b8660208260051b850101111561023757600080fd5b60209290920196919550909350505050565b60005b8381101561026457818101518382015260200161024c565b50506000910152565b60008151808452610285816020860160208601610249565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561032a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261031885835161026d565b945092850192908501906001016102de565b5092979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126103ca57600080fd5b83018035915067ffffffffffffffff8211156103e557600080fd5b6020019150368190038213156103fa57600080fd5b9250929050565b8183823760009101908152919050565b60006020828403121561042357600080fd5b815167ffffffffffffffff8082111561043b57600080fd5b818401915084601f83011261044f57600080fd5b81518181111561046157610461610337565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156104a7576104a7610337565b816040528281528760208487010111156104c057600080fd5b6104d1836020830160208801610249565b979650505050505050565b6020815260006104ef602083018461026d565b9392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361054e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea264697066735822122094728da134cbf93fb77ee24bfbc268b74c556d11aa42868710e607fcef24619964736f6c63430008140033",
}

// MultiCallerABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiCallerMetaData.ABI instead.
var MultiCallerABI = MultiCallerMetaData.ABI

// MultiCallerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiCallerMetaData.Bin instead.
var MultiCallerBin = MultiCallerMetaData.Bin

// DeployMultiCaller deploys a new Ethereum contract, binding an instance of MultiCaller to it.
func DeployMultiCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiCaller, error) {
	parsed, err := MultiCallerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiCaller{MultiCallerCaller: MultiCallerCaller{contract: contract}, MultiCallerTransactor: MultiCallerTransactor{contract: contract}, MultiCallerFilterer: MultiCallerFilterer{contract: contract}}, nil
}

// MultiCaller is an auto generated Go binding around an Ethereum contract.
type MultiCaller struct {
	MultiCallerCaller     // Read-only binding to the contract
	MultiCallerTransactor // Write-only binding to the contract
	MultiCallerFilterer   // Log filterer for contract events
}

// MultiCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiCallerSession struct {
	Contract     *MultiCaller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCallerCallerSession struct {
	Contract *MultiCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MultiCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiCallerTransactorSession struct {
	Contract     *MultiCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MultiCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiCallerRaw struct {
	Contract *MultiCaller // Generic contract binding to access the raw methods on
}

// MultiCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCallerCallerRaw struct {
	Contract *MultiCallerCaller // Generic read-only contract binding to access the raw methods on
}

// MultiCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiCallerTransactorRaw struct {
	Contract *MultiCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiCaller creates a new instance of MultiCaller, bound to a specific deployed contract.
func NewMultiCaller(address common.Address, backend bind.ContractBackend) (*MultiCaller, error) {
	contract, err := bindMultiCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiCaller{MultiCallerCaller: MultiCallerCaller{contract: contract}, MultiCallerTransactor: MultiCallerTransactor{contract: contract}, MultiCallerFilterer: MultiCallerFilterer{contract: contract}}, nil
}

// NewMultiCallerCaller creates a new read-only instance of MultiCaller, bound to a specific deployed contract.
func NewMultiCallerCaller(address common.Address, caller bind.ContractCaller) (*MultiCallerCaller, error) {
	contract, err := bindMultiCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallerCaller{contract: contract}, nil
}

// NewMultiCallerTransactor creates a new write-only instance of MultiCaller, bound to a specific deployed contract.
func NewMultiCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiCallerTransactor, error) {
	contract, err := bindMultiCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallerTransactor{contract: contract}, nil
}

// NewMultiCallerFilterer creates a new log filterer instance of MultiCaller, bound to a specific deployed contract.
func NewMultiCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiCallerFilterer, error) {
	contract, err := bindMultiCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiCallerFilterer{contract: contract}, nil
}

// bindMultiCaller binds a generic wrapper to an already deployed contract.
func bindMultiCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MultiCallerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCaller *MultiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCaller.Contract.MultiCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCaller *MultiCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCaller.Contract.MultiCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCaller *MultiCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCaller.Contract.MultiCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCaller *MultiCallerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCaller *MultiCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCaller *MultiCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCaller.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MultiCaller *MultiCallerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _MultiCaller.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MultiCaller *MultiCallerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MultiCaller.Contract.Multicall(&_MultiCaller.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_MultiCaller *MultiCallerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _MultiCaller.Contract.Multicall(&_MultiCaller.TransactOpts, data)
}
