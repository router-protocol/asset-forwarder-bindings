// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IUniswapExchange

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

// IUniswapExchangeMetaData contains all meta data concerning the IUniswapExchange contract.
var IUniswapExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ethToTokenSwapInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethSold\",\"type\":\"uint256\"}],\"name\":\"getEthToTokenInputPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensSold\",\"type\":\"uint256\"}],\"name\":\"getTokenToEthInputPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBought\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"tokenToEthSwapInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethBought\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensSold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTokensBought\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minEthBought\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"}],\"name\":\"tokenToTokenSwapInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapExchangeMetaData.ABI instead.
var IUniswapExchangeABI = IUniswapExchangeMetaData.ABI

// IUniswapExchange is an auto generated Go binding around an Ethereum contract.
type IUniswapExchange struct {
	IUniswapExchangeCaller     // Read-only binding to the contract
	IUniswapExchangeTransactor // Write-only binding to the contract
	IUniswapExchangeFilterer   // Log filterer for contract events
}

// IUniswapExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapExchangeSession struct {
	Contract     *IUniswapExchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IUniswapExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapExchangeCallerSession struct {
	Contract *IUniswapExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IUniswapExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapExchangeTransactorSession struct {
	Contract     *IUniswapExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IUniswapExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapExchangeRaw struct {
	Contract *IUniswapExchange // Generic contract binding to access the raw methods on
}

// IUniswapExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapExchangeCallerRaw struct {
	Contract *IUniswapExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapExchangeTransactorRaw struct {
	Contract *IUniswapExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapExchange creates a new instance of IUniswapExchange, bound to a specific deployed contract.
func NewIUniswapExchange(address common.Address, backend bind.ContractBackend) (*IUniswapExchange, error) {
	contract, err := bindIUniswapExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapExchange{IUniswapExchangeCaller: IUniswapExchangeCaller{contract: contract}, IUniswapExchangeTransactor: IUniswapExchangeTransactor{contract: contract}, IUniswapExchangeFilterer: IUniswapExchangeFilterer{contract: contract}}, nil
}

// NewIUniswapExchangeCaller creates a new read-only instance of IUniswapExchange, bound to a specific deployed contract.
func NewIUniswapExchangeCaller(address common.Address, caller bind.ContractCaller) (*IUniswapExchangeCaller, error) {
	contract, err := bindIUniswapExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapExchangeCaller{contract: contract}, nil
}

// NewIUniswapExchangeTransactor creates a new write-only instance of IUniswapExchange, bound to a specific deployed contract.
func NewIUniswapExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapExchangeTransactor, error) {
	contract, err := bindIUniswapExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapExchangeTransactor{contract: contract}, nil
}

// NewIUniswapExchangeFilterer creates a new log filterer instance of IUniswapExchange, bound to a specific deployed contract.
func NewIUniswapExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapExchangeFilterer, error) {
	contract, err := bindIUniswapExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapExchangeFilterer{contract: contract}, nil
}

// bindIUniswapExchange binds a generic wrapper to an already deployed contract.
func bindIUniswapExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUniswapExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapExchange *IUniswapExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapExchange.Contract.IUniswapExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapExchange *IUniswapExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.IUniswapExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapExchange *IUniswapExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.IUniswapExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapExchange *IUniswapExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapExchange *IUniswapExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapExchange *IUniswapExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.contract.Transact(opts, method, params...)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 ethSold) view returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeCaller) GetEthToTokenInputPrice(opts *bind.CallOpts, ethSold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapExchange.contract.Call(opts, &out, "getEthToTokenInputPrice", ethSold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 ethSold) view returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeSession) GetEthToTokenInputPrice(ethSold *big.Int) (*big.Int, error) {
	return _IUniswapExchange.Contract.GetEthToTokenInputPrice(&_IUniswapExchange.CallOpts, ethSold)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 ethSold) view returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeCallerSession) GetEthToTokenInputPrice(ethSold *big.Int) (*big.Int, error) {
	return _IUniswapExchange.Contract.GetEthToTokenInputPrice(&_IUniswapExchange.CallOpts, ethSold)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokensSold) view returns(uint256 ethBought)
func (_IUniswapExchange *IUniswapExchangeCaller) GetTokenToEthInputPrice(opts *bind.CallOpts, tokensSold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IUniswapExchange.contract.Call(opts, &out, "getTokenToEthInputPrice", tokensSold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokensSold) view returns(uint256 ethBought)
func (_IUniswapExchange *IUniswapExchangeSession) GetTokenToEthInputPrice(tokensSold *big.Int) (*big.Int, error) {
	return _IUniswapExchange.Contract.GetTokenToEthInputPrice(&_IUniswapExchange.CallOpts, tokensSold)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokensSold) view returns(uint256 ethBought)
func (_IUniswapExchange *IUniswapExchangeCallerSession) GetTokenToEthInputPrice(tokensSold *big.Int) (*big.Int, error) {
	return _IUniswapExchange.Contract.GetTokenToEthInputPrice(&_IUniswapExchange.CallOpts, tokensSold)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 minTokens, uint256 deadline) payable returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeTransactor) EthToTokenSwapInput(opts *bind.TransactOpts, minTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapExchange.contract.Transact(opts, "ethToTokenSwapInput", minTokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 minTokens, uint256 deadline) payable returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeSession) EthToTokenSwapInput(minTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.EthToTokenSwapInput(&_IUniswapExchange.TransactOpts, minTokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 minTokens, uint256 deadline) payable returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeTransactorSession) EthToTokenSwapInput(minTokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.EthToTokenSwapInput(&_IUniswapExchange.TransactOpts, minTokens, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokensSold, uint256 minEth, uint256 deadline) returns(uint256 ethBought)
func (_IUniswapExchange *IUniswapExchangeTransactor) TokenToEthSwapInput(opts *bind.TransactOpts, tokensSold *big.Int, minEth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapExchange.contract.Transact(opts, "tokenToEthSwapInput", tokensSold, minEth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokensSold, uint256 minEth, uint256 deadline) returns(uint256 ethBought)
func (_IUniswapExchange *IUniswapExchangeSession) TokenToEthSwapInput(tokensSold *big.Int, minEth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.TokenToEthSwapInput(&_IUniswapExchange.TransactOpts, tokensSold, minEth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokensSold, uint256 minEth, uint256 deadline) returns(uint256 ethBought)
func (_IUniswapExchange *IUniswapExchangeTransactorSession) TokenToEthSwapInput(tokensSold *big.Int, minEth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.TokenToEthSwapInput(&_IUniswapExchange.TransactOpts, tokensSold, minEth, deadline)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokensSold, uint256 minTokensBought, uint256 minEthBought, uint256 deadline, address tokenAddr) returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeTransactor) TokenToTokenSwapInput(opts *bind.TransactOpts, tokensSold *big.Int, minTokensBought *big.Int, minEthBought *big.Int, deadline *big.Int, tokenAddr common.Address) (*types.Transaction, error) {
	return _IUniswapExchange.contract.Transact(opts, "tokenToTokenSwapInput", tokensSold, minTokensBought, minEthBought, deadline, tokenAddr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokensSold, uint256 minTokensBought, uint256 minEthBought, uint256 deadline, address tokenAddr) returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeSession) TokenToTokenSwapInput(tokensSold *big.Int, minTokensBought *big.Int, minEthBought *big.Int, deadline *big.Int, tokenAddr common.Address) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.TokenToTokenSwapInput(&_IUniswapExchange.TransactOpts, tokensSold, minTokensBought, minEthBought, deadline, tokenAddr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokensSold, uint256 minTokensBought, uint256 minEthBought, uint256 deadline, address tokenAddr) returns(uint256 tokensBought)
func (_IUniswapExchange *IUniswapExchangeTransactorSession) TokenToTokenSwapInput(tokensSold *big.Int, minTokensBought *big.Int, minEthBought *big.Int, deadline *big.Int, tokenAddr common.Address) (*types.Transaction, error) {
	return _IUniswapExchange.Contract.TokenToTokenSwapInput(&_IUniswapExchange.TransactOpts, tokensSold, minTokensBought, minEthBought, deadline, tokenAddr)
}
