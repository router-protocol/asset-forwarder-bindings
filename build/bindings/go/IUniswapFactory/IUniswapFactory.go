// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IUniswapFactory

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

// IUniswapFactoryMetaData contains all meta data concerning the IUniswapFactory contract.
var IUniswapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getExchange\",\"outputs\":[{\"internalType\":\"contractIUniswapExchange\",\"name\":\"exchange\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapFactoryMetaData.ABI instead.
var IUniswapFactoryABI = IUniswapFactoryMetaData.ABI

// IUniswapFactory is an auto generated Go binding around an Ethereum contract.
type IUniswapFactory struct {
	IUniswapFactoryCaller     // Read-only binding to the contract
	IUniswapFactoryTransactor // Write-only binding to the contract
	IUniswapFactoryFilterer   // Log filterer for contract events
}

// IUniswapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapFactorySession struct {
	Contract     *IUniswapFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniswapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapFactoryCallerSession struct {
	Contract *IUniswapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IUniswapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapFactoryTransactorSession struct {
	Contract     *IUniswapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IUniswapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapFactoryRaw struct {
	Contract *IUniswapFactory // Generic contract binding to access the raw methods on
}

// IUniswapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapFactoryCallerRaw struct {
	Contract *IUniswapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapFactoryTransactorRaw struct {
	Contract *IUniswapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapFactory creates a new instance of IUniswapFactory, bound to a specific deployed contract.
func NewIUniswapFactory(address common.Address, backend bind.ContractBackend) (*IUniswapFactory, error) {
	contract, err := bindIUniswapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapFactory{IUniswapFactoryCaller: IUniswapFactoryCaller{contract: contract}, IUniswapFactoryTransactor: IUniswapFactoryTransactor{contract: contract}, IUniswapFactoryFilterer: IUniswapFactoryFilterer{contract: contract}}, nil
}

// NewIUniswapFactoryCaller creates a new read-only instance of IUniswapFactory, bound to a specific deployed contract.
func NewIUniswapFactoryCaller(address common.Address, caller bind.ContractCaller) (*IUniswapFactoryCaller, error) {
	contract, err := bindIUniswapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapFactoryCaller{contract: contract}, nil
}

// NewIUniswapFactoryTransactor creates a new write-only instance of IUniswapFactory, bound to a specific deployed contract.
func NewIUniswapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapFactoryTransactor, error) {
	contract, err := bindIUniswapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapFactoryTransactor{contract: contract}, nil
}

// NewIUniswapFactoryFilterer creates a new log filterer instance of IUniswapFactory, bound to a specific deployed contract.
func NewIUniswapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapFactoryFilterer, error) {
	contract, err := bindIUniswapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapFactoryFilterer{contract: contract}, nil
}

// bindIUniswapFactory binds a generic wrapper to an already deployed contract.
func bindIUniswapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUniswapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapFactory *IUniswapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapFactory.Contract.IUniswapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapFactory *IUniswapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapFactory.Contract.IUniswapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapFactory *IUniswapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapFactory.Contract.IUniswapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapFactory *IUniswapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapFactory *IUniswapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapFactory *IUniswapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapFactory.Contract.contract.Transact(opts, method, params...)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) view returns(address exchange)
func (_IUniswapFactory *IUniswapFactoryCaller) GetExchange(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IUniswapFactory.contract.Call(opts, &out, "getExchange", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) view returns(address exchange)
func (_IUniswapFactory *IUniswapFactorySession) GetExchange(token common.Address) (common.Address, error) {
	return _IUniswapFactory.Contract.GetExchange(&_IUniswapFactory.CallOpts, token)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) view returns(address exchange)
func (_IUniswapFactory *IUniswapFactoryCallerSession) GetExchange(token common.Address) (common.Address, error) {
	return _IUniswapFactory.Contract.GetExchange(&_IUniswapFactory.CallOpts, token)
}
