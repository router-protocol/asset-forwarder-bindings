// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IUniswapV2Exchange

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

// IUniswapV2ExchangeMetaData contains all meta data concerning the IUniswapV2Exchange contract.
var IUniswapV2ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IUniswapV2ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniswapV2ExchangeMetaData.ABI instead.
var IUniswapV2ExchangeABI = IUniswapV2ExchangeMetaData.ABI

// IUniswapV2Exchange is an auto generated Go binding around an Ethereum contract.
type IUniswapV2Exchange struct {
	IUniswapV2ExchangeCaller     // Read-only binding to the contract
	IUniswapV2ExchangeTransactor // Write-only binding to the contract
	IUniswapV2ExchangeFilterer   // Log filterer for contract events
}

// IUniswapV2ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniswapV2ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniswapV2ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniswapV2ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniswapV2ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniswapV2ExchangeSession struct {
	Contract     *IUniswapV2Exchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IUniswapV2ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniswapV2ExchangeCallerSession struct {
	Contract *IUniswapV2ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IUniswapV2ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniswapV2ExchangeTransactorSession struct {
	Contract     *IUniswapV2ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IUniswapV2ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniswapV2ExchangeRaw struct {
	Contract *IUniswapV2Exchange // Generic contract binding to access the raw methods on
}

// IUniswapV2ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniswapV2ExchangeCallerRaw struct {
	Contract *IUniswapV2ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// IUniswapV2ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniswapV2ExchangeTransactorRaw struct {
	Contract *IUniswapV2ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniswapV2Exchange creates a new instance of IUniswapV2Exchange, bound to a specific deployed contract.
func NewIUniswapV2Exchange(address common.Address, backend bind.ContractBackend) (*IUniswapV2Exchange, error) {
	contract, err := bindIUniswapV2Exchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2Exchange{IUniswapV2ExchangeCaller: IUniswapV2ExchangeCaller{contract: contract}, IUniswapV2ExchangeTransactor: IUniswapV2ExchangeTransactor{contract: contract}, IUniswapV2ExchangeFilterer: IUniswapV2ExchangeFilterer{contract: contract}}, nil
}

// NewIUniswapV2ExchangeCaller creates a new read-only instance of IUniswapV2Exchange, bound to a specific deployed contract.
func NewIUniswapV2ExchangeCaller(address common.Address, caller bind.ContractCaller) (*IUniswapV2ExchangeCaller, error) {
	contract, err := bindIUniswapV2Exchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2ExchangeCaller{contract: contract}, nil
}

// NewIUniswapV2ExchangeTransactor creates a new write-only instance of IUniswapV2Exchange, bound to a specific deployed contract.
func NewIUniswapV2ExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniswapV2ExchangeTransactor, error) {
	contract, err := bindIUniswapV2Exchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2ExchangeTransactor{contract: contract}, nil
}

// NewIUniswapV2ExchangeFilterer creates a new log filterer instance of IUniswapV2Exchange, bound to a specific deployed contract.
func NewIUniswapV2ExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniswapV2ExchangeFilterer, error) {
	contract, err := bindIUniswapV2Exchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniswapV2ExchangeFilterer{contract: contract}, nil
}

// bindIUniswapV2Exchange binds a generic wrapper to an already deployed contract.
func bindIUniswapV2Exchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUniswapV2ExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Exchange *IUniswapV2ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Exchange.Contract.IUniswapV2ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Exchange *IUniswapV2ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.IUniswapV2ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Exchange *IUniswapV2ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.IUniswapV2ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniswapV2Exchange *IUniswapV2ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniswapV2Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_IUniswapV2Exchange *IUniswapV2ExchangeCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _IUniswapV2Exchange.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_IUniswapV2Exchange *IUniswapV2ExchangeSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IUniswapV2Exchange.Contract.GetReserves(&_IUniswapV2Exchange.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_IUniswapV2Exchange *IUniswapV2ExchangeCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _IUniswapV2Exchange.Contract.GetReserves(&_IUniswapV2Exchange.CallOpts)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IUniswapV2Exchange.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.Skim(&_IUniswapV2Exchange.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.Skim(&_IUniswapV2Exchange.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IUniswapV2Exchange.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.Swap(&_IUniswapV2Exchange.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.Swap(&_IUniswapV2Exchange.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniswapV2Exchange.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeSession) Sync() (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.Sync(&_IUniswapV2Exchange.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_IUniswapV2Exchange *IUniswapV2ExchangeTransactorSession) Sync() (*types.Transaction, error) {
	return _IUniswapV2Exchange.Contract.Sync(&_IUniswapV2Exchange.TransactOpts)
}
