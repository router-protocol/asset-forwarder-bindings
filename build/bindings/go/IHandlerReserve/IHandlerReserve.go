// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IHandlerReserve

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

// IHandlerReserveMetaData contains all meta data concerning the IHandlerReserve contract.
var IHandlerReserveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"_contractToLP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"_lpToContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IHandlerReserveABI is the input ABI used to generate the binding from.
// Deprecated: Use IHandlerReserveMetaData.ABI instead.
var IHandlerReserveABI = IHandlerReserveMetaData.ABI

// IHandlerReserve is an auto generated Go binding around an Ethereum contract.
type IHandlerReserve struct {
	IHandlerReserveCaller     // Read-only binding to the contract
	IHandlerReserveTransactor // Write-only binding to the contract
	IHandlerReserveFilterer   // Log filterer for contract events
}

// IHandlerReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHandlerReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHandlerReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHandlerReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHandlerReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHandlerReserveSession struct {
	Contract     *IHandlerReserve  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IHandlerReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHandlerReserveCallerSession struct {
	Contract *IHandlerReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IHandlerReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHandlerReserveTransactorSession struct {
	Contract     *IHandlerReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IHandlerReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHandlerReserveRaw struct {
	Contract *IHandlerReserve // Generic contract binding to access the raw methods on
}

// IHandlerReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHandlerReserveCallerRaw struct {
	Contract *IHandlerReserveCaller // Generic read-only contract binding to access the raw methods on
}

// IHandlerReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHandlerReserveTransactorRaw struct {
	Contract *IHandlerReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHandlerReserve creates a new instance of IHandlerReserve, bound to a specific deployed contract.
func NewIHandlerReserve(address common.Address, backend bind.ContractBackend) (*IHandlerReserve, error) {
	contract, err := bindIHandlerReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHandlerReserve{IHandlerReserveCaller: IHandlerReserveCaller{contract: contract}, IHandlerReserveTransactor: IHandlerReserveTransactor{contract: contract}, IHandlerReserveFilterer: IHandlerReserveFilterer{contract: contract}}, nil
}

// NewIHandlerReserveCaller creates a new read-only instance of IHandlerReserve, bound to a specific deployed contract.
func NewIHandlerReserveCaller(address common.Address, caller bind.ContractCaller) (*IHandlerReserveCaller, error) {
	contract, err := bindIHandlerReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHandlerReserveCaller{contract: contract}, nil
}

// NewIHandlerReserveTransactor creates a new write-only instance of IHandlerReserve, bound to a specific deployed contract.
func NewIHandlerReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*IHandlerReserveTransactor, error) {
	contract, err := bindIHandlerReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHandlerReserveTransactor{contract: contract}, nil
}

// NewIHandlerReserveFilterer creates a new log filterer instance of IHandlerReserve, bound to a specific deployed contract.
func NewIHandlerReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*IHandlerReserveFilterer, error) {
	contract, err := bindIHandlerReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHandlerReserveFilterer{contract: contract}, nil
}

// bindIHandlerReserve binds a generic wrapper to an already deployed contract.
func bindIHandlerReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHandlerReserveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHandlerReserve *IHandlerReserveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHandlerReserve.Contract.IHandlerReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHandlerReserve *IHandlerReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.IHandlerReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHandlerReserve *IHandlerReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.IHandlerReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHandlerReserve *IHandlerReserveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHandlerReserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHandlerReserve *IHandlerReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHandlerReserve *IHandlerReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.contract.Transact(opts, method, params...)
}

// ContractToLP is a paid mutator transaction binding the contract method 0x5fae9245.
//
// Solidity: function _contractToLP(address token) returns(address)
func (_IHandlerReserve *IHandlerReserveTransactor) ContractToLP(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IHandlerReserve.contract.Transact(opts, "_contractToLP", token)
}

// ContractToLP is a paid mutator transaction binding the contract method 0x5fae9245.
//
// Solidity: function _contractToLP(address token) returns(address)
func (_IHandlerReserve *IHandlerReserveSession) ContractToLP(token common.Address) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.ContractToLP(&_IHandlerReserve.TransactOpts, token)
}

// ContractToLP is a paid mutator transaction binding the contract method 0x5fae9245.
//
// Solidity: function _contractToLP(address token) returns(address)
func (_IHandlerReserve *IHandlerReserveTransactorSession) ContractToLP(token common.Address) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.ContractToLP(&_IHandlerReserve.TransactOpts, token)
}

// LpToContract is a paid mutator transaction binding the contract method 0xaa50ec12.
//
// Solidity: function _lpToContract(address token) returns(address)
func (_IHandlerReserve *IHandlerReserveTransactor) LpToContract(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IHandlerReserve.contract.Transact(opts, "_lpToContract", token)
}

// LpToContract is a paid mutator transaction binding the contract method 0xaa50ec12.
//
// Solidity: function _lpToContract(address token) returns(address)
func (_IHandlerReserve *IHandlerReserveSession) LpToContract(token common.Address) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.LpToContract(&_IHandlerReserve.TransactOpts, token)
}

// LpToContract is a paid mutator transaction binding the contract method 0xaa50ec12.
//
// Solidity: function _lpToContract(address token) returns(address)
func (_IHandlerReserve *IHandlerReserveTransactorSession) LpToContract(token common.Address) (*types.Transaction, error) {
	return _IHandlerReserve.Contract.LpToContract(&_IHandlerReserve.TransactOpts, token)
}
