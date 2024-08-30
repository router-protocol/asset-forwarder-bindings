// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IUniswapV2Factory

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

// IUniswapV2FactoryMetaData contains all meta data concerning the IUniswapV2Factory contract.
var IUniswapV2FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Exchange\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IUniswapV2FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV2FactoryMetaData.ABI instead.
var IUniswapV2FactoryABI = IUniswapV2FactoryMetaData.ABI

// IUniswapV2Factory is an auto generated Go binding around an Ethereum contract.
type IUniswapV2Factory struct {
	IUniswapV2FactoryCaller     // Read-only binding to the contract
	IUniswapV2FactoryTransactor // Write-only binding to the contract
	IUniswapV2FactoryFilterer   // Log filterer for contract events
}

// IUniswapV2FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV2FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV2FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV2FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV2FactorySession struct {
	Contract     *IUniswapV2Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IUniswapV2FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV2FactoryCallerSession struct {
	Contract *IUniswapV2FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IUniswapV2FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV2FactoryTransactorSession struct {
	Contract     *IUniswapV2FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IUniswapV2FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV2FactoryRaw struct {
	Contract *IUniswapV2Factory // Generic contract binding to access the raw methods on
}

// IUniswapV2FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV2FactoryCallerRaw struct {
	Contract *IUniswapV2FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV2FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV2FactoryTransactorRaw struct {
	Contract *IUniswapV2FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV2Factory creates a new instance of IUniswapV2Factory, bound to a specific deployed contract.
func NewIUniswapV2Factory(address common.Address, backend bind.ContractBackend) (*IUniswapV2Factory, error) {
	contract, err := bindIUniswapV2Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Factory{IUniswapV2FactoryCaller: IUniswapV2FactoryCaller{contract: contract}, IUniswapV2FactoryTransactor: IUniswapV2FactoryTransactor{contract: contract}, IUniswapV2FactoryFilterer: IUniswapV2FactoryFilterer{contract: contract}}, nil
}

// NewIUniswapV2FactoryCaller creates a new read-only instance of IUniswapV2Factory, bound to a specific deployed contract.
func NewIUniswapV2FactoryCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV2FactoryCaller, error) {
	contract, err := bindIUniswapV2Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2FactoryCaller{contract: contract}, nil
}

// NewIUniswapV2FactoryTransactor creates a new write-only instance of IUniswapV2Factory, bound to a specific deployed contract.
func NewIUniswapV2FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV2FactoryTransactor, error) {
	contract, err := bindIUniswapV2Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2FactoryTransactor{contract: contract}, nil
}

// NewIUniswapV2FactoryFilterer creates a new log filterer instance of IUniswapV2Factory, bound to a specific deployed contract.
func NewIUniswapV2FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV2FactoryFilterer, error) {
	contract, err := bindIUniswapV2Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2FactoryFilterer{contract: contract}, nil
}

// bindIUniswapV2Factory binds a generic wrapper to an already deployed contract.
func bindIUniswapV2Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUniswapV2FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Factory *IUniswapV2FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Factory.Contract.IUniswapV2FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Factory *IUniswapV2FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Factory.Contract.IUniswapV2FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Factory *IUniswapV2FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Factory.Contract.IUniswapV2FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Factory *IUniswapV2FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Factory *IUniswapV2FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Factory *IUniswapV2FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Factory.Contract.contract.Transact(opts, method, params...)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IUniswapV2Factory *IUniswapV2FactoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _IUniswapV2Factory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IUniswapV2Factory *IUniswapV2FactorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IUniswapV2Factory.Contract.GetPair(&_IUniswapV2Factory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_IUniswapV2Factory *IUniswapV2FactoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _IUniswapV2Factory.Contract.GetPair(&_IUniswapV2Factory.CallOpts, tokenA, tokenB)
}
