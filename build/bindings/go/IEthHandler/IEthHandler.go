// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IEthHandler

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

// IEthHandlerMetaData contains all meta data concerning the IEthHandler contract.
var IEthHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"WETH\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IEthHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthHandlerMetaData.ABI instead.
var IEthHandlerABI = IEthHandlerMetaData.ABI

// IEthHandler is an auto generated Go binding around an Ethereum contract.
type IEthHandler struct {
	IEthHandlerCaller     // Read-only binding to the contract
	IEthHandlerTransactor // Write-only binding to the contract
	IEthHandlerFilterer   // Log filterer for contract events
}

// IEthHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthHandlerSession struct {
	Contract     *IEthHandler      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEthHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthHandlerCallerSession struct {
	Contract *IEthHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IEthHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthHandlerTransactorSession struct {
	Contract     *IEthHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IEthHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthHandlerRaw struct {
	Contract *IEthHandler // Generic contract binding to access the raw methods on
}

// IEthHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthHandlerCallerRaw struct {
	Contract *IEthHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IEthHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthHandlerTransactorRaw struct {
	Contract *IEthHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthHandler creates a new instance of IEthHandler, bound to a specific deployed contract.
func NewIEthHandler(address common.Address, backend bind.ContractBackend) (*IEthHandler, error) {
	contract, err := bindIEthHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthHandler{IEthHandlerCaller: IEthHandlerCaller{contract: contract}, IEthHandlerTransactor: IEthHandlerTransactor{contract: contract}, IEthHandlerFilterer: IEthHandlerFilterer{contract: contract}}, nil
}

// NewIEthHandlerCaller creates a new read-only instance of IEthHandler, bound to a specific deployed contract.
func NewIEthHandlerCaller(address common.Address, caller bind.ContractCaller) (*IEthHandlerCaller, error) {
	contract, err := bindIEthHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthHandlerCaller{contract: contract}, nil
}

// NewIEthHandlerTransactor creates a new write-only instance of IEthHandler, bound to a specific deployed contract.
func NewIEthHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthHandlerTransactor, error) {
	contract, err := bindIEthHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthHandlerTransactor{contract: contract}, nil
}

// NewIEthHandlerFilterer creates a new log filterer instance of IEthHandler, bound to a specific deployed contract.
func NewIEthHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthHandlerFilterer, error) {
	contract, err := bindIEthHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthHandlerFilterer{contract: contract}, nil
}

// bindIEthHandler binds a generic wrapper to an already deployed contract.
func bindIEthHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IEthHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthHandler *IEthHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthHandler.Contract.IEthHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthHandler *IEthHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthHandler.Contract.IEthHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthHandler *IEthHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthHandler.Contract.IEthHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthHandler *IEthHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthHandler *IEthHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthHandler *IEthHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthHandler.Contract.contract.Transact(opts, method, params...)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address WETH, uint256 ) returns()
func (_IEthHandler *IEthHandlerTransactor) Withdraw(opts *bind.TransactOpts, WETH common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IEthHandler.contract.Transact(opts, "withdraw", WETH, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address WETH, uint256 ) returns()
func (_IEthHandler *IEthHandlerSession) Withdraw(WETH common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IEthHandler.Contract.Withdraw(&_IEthHandler.TransactOpts, WETH, arg1)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address WETH, uint256 ) returns()
func (_IEthHandler *IEthHandlerTransactorSession) Withdraw(WETH common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _IEthHandler.Contract.Withdraw(&_IEthHandler.TransactOpts, WETH, arg1)
}
