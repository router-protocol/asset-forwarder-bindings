// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAugustusSwapper

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

// IAugustusSwapperMetaData contains all meta data concerning the IAugustusSwapper contract.
var IAugustusSwapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getTokenTransferProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAugustusSwapperABI is the input ABI used to generate the binding from.
// Deprecated: Use IAugustusSwapperMetaData.ABI instead.
var IAugustusSwapperABI = IAugustusSwapperMetaData.ABI

// IAugustusSwapper is an auto generated Go binding around an Ethereum contract.
type IAugustusSwapper struct {
	IAugustusSwapperCaller     // Read-only binding to the contract
	IAugustusSwapperTransactor // Write-only binding to the contract
	IAugustusSwapperFilterer   // Log filterer for contract events
}

// IAugustusSwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAugustusSwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugustusSwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAugustusSwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugustusSwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAugustusSwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAugustusSwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAugustusSwapperSession struct {
	Contract     *IAugustusSwapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAugustusSwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAugustusSwapperCallerSession struct {
	Contract *IAugustusSwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IAugustusSwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAugustusSwapperTransactorSession struct {
	Contract     *IAugustusSwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IAugustusSwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAugustusSwapperRaw struct {
	Contract *IAugustusSwapper // Generic contract binding to access the raw methods on
}

// IAugustusSwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAugustusSwapperCallerRaw struct {
	Contract *IAugustusSwapperCaller // Generic read-only contract binding to access the raw methods on
}

// IAugustusSwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAugustusSwapperTransactorRaw struct {
	Contract *IAugustusSwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAugustusSwapper creates a new instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapper(address common.Address, backend bind.ContractBackend) (*IAugustusSwapper, error) {
	contract, err := bindIAugustusSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapper{IAugustusSwapperCaller: IAugustusSwapperCaller{contract: contract}, IAugustusSwapperTransactor: IAugustusSwapperTransactor{contract: contract}, IAugustusSwapperFilterer: IAugustusSwapperFilterer{contract: contract}}, nil
}

// NewIAugustusSwapperCaller creates a new read-only instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapperCaller(address common.Address, caller bind.ContractCaller) (*IAugustusSwapperCaller, error) {
	contract, err := bindIAugustusSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapperCaller{contract: contract}, nil
}

// NewIAugustusSwapperTransactor creates a new write-only instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IAugustusSwapperTransactor, error) {
	contract, err := bindIAugustusSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapperTransactor{contract: contract}, nil
}

// NewIAugustusSwapperFilterer creates a new log filterer instance of IAugustusSwapper, bound to a specific deployed contract.
func NewIAugustusSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IAugustusSwapperFilterer, error) {
	contract, err := bindIAugustusSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAugustusSwapperFilterer{contract: contract}, nil
}

// bindIAugustusSwapper binds a generic wrapper to an already deployed contract.
func bindIAugustusSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAugustusSwapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAugustusSwapper *IAugustusSwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAugustusSwapper.Contract.IAugustusSwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAugustusSwapper *IAugustusSwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.IAugustusSwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAugustusSwapper *IAugustusSwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.IAugustusSwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAugustusSwapper *IAugustusSwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAugustusSwapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAugustusSwapper *IAugustusSwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAugustusSwapper *IAugustusSwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAugustusSwapper.Contract.contract.Transact(opts, method, params...)
}

// GetTokenTransferProxy is a free data retrieval call binding the contract method 0xd2c4b598.
//
// Solidity: function getTokenTransferProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCaller) GetTokenTransferProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAugustusSwapper.contract.Call(opts, &out, "getTokenTransferProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenTransferProxy is a free data retrieval call binding the contract method 0xd2c4b598.
//
// Solidity: function getTokenTransferProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperSession) GetTokenTransferProxy() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetTokenTransferProxy(&_IAugustusSwapper.CallOpts)
}

// GetTokenTransferProxy is a free data retrieval call binding the contract method 0xd2c4b598.
//
// Solidity: function getTokenTransferProxy() view returns(address)
func (_IAugustusSwapper *IAugustusSwapperCallerSession) GetTokenTransferProxy() (common.Address, error) {
	return _IAugustusSwapper.Contract.GetTokenTransferProxy(&_IAugustusSwapper.CallOpts)
}
