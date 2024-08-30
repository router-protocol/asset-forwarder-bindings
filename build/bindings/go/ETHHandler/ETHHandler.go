// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ETHHandler

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

// ETHHandlerMetaData contains all meta data concerning the ETHHandler contract.
var ETHHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610275806100206000396000f3fe6080604052600436106100225760003560e01c8063f3fef3a31461002e57600080fd5b3661002957005b600080fd5b34801561003a57600080fd5b5061004e6100493660046101cb565b610050565b005b6040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810182905273ffffffffffffffffffffffffffffffffffffffff831690632e1a7d4d90602401600060405180830381600087803b1580156100b857600080fd5b505af11580156100cc573d6000803e3d6000fd5b50506040805160008082526020820192839052935033925084916100f09190610210565b60006040518083038185875af1925050503d806000811461012d576040519150601f19603f3d011682016040523d82523d6000602084013e610132565b606091505b50509050806101c6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f736166655472616e736665724554483a20455448207472616e7366657220666160448201527f696c656400000000000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b505050565b600080604083850312156101de57600080fd5b823573ffffffffffffffffffffffffffffffffffffffff8116811461020257600080fd5b946020939093013593505050565b6000825160005b818110156102315760208186018101518583015201610217565b50600092019182525091905056fea26469706673582212207b508ea1e5b79e24ab5621a919d0aed420bfb5b15e56d45abacf3327158b4c9164736f6c63430008140033",
}

// ETHHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHHandlerMetaData.ABI instead.
var ETHHandlerABI = ETHHandlerMetaData.ABI

// ETHHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ETHHandlerMetaData.Bin instead.
var ETHHandlerBin = ETHHandlerMetaData.Bin

// DeployETHHandler deploys a new Ethereum contract, binding an instance of ETHHandler to it.
func DeployETHHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ETHHandler, error) {
	parsed, err := ETHHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ETHHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ETHHandler{ETHHandlerCaller: ETHHandlerCaller{contract: contract}, ETHHandlerTransactor: ETHHandlerTransactor{contract: contract}, ETHHandlerFilterer: ETHHandlerFilterer{contract: contract}}, nil
}

// ETHHandler is an auto generated Go binding around an Ethereum contract.
type ETHHandler struct {
	ETHHandlerCaller     // Read-only binding to the contract
	ETHHandlerTransactor // Write-only binding to the contract
	ETHHandlerFilterer   // Log filterer for contract events
}

// ETHHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHHandlerSession struct {
	Contract     *ETHHandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHHandlerCallerSession struct {
	Contract *ETHHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ETHHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHHandlerTransactorSession struct {
	Contract     *ETHHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ETHHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHHandlerRaw struct {
	Contract *ETHHandler // Generic contract binding to access the raw methods on
}

// ETHHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHHandlerCallerRaw struct {
	Contract *ETHHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// ETHHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHHandlerTransactorRaw struct {
	Contract *ETHHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHHandler creates a new instance of ETHHandler, bound to a specific deployed contract.
func NewETHHandler(address common.Address, backend bind.ContractBackend) (*ETHHandler, error) {
	contract, err := bindETHHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHHandler{ETHHandlerCaller: ETHHandlerCaller{contract: contract}, ETHHandlerTransactor: ETHHandlerTransactor{contract: contract}, ETHHandlerFilterer: ETHHandlerFilterer{contract: contract}}, nil
}

// NewETHHandlerCaller creates a new read-only instance of ETHHandler, bound to a specific deployed contract.
func NewETHHandlerCaller(address common.Address, caller bind.ContractCaller) (*ETHHandlerCaller, error) {
	contract, err := bindETHHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHHandlerCaller{contract: contract}, nil
}

// NewETHHandlerTransactor creates a new write-only instance of ETHHandler, bound to a specific deployed contract.
func NewETHHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHHandlerTransactor, error) {
	contract, err := bindETHHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHHandlerTransactor{contract: contract}, nil
}

// NewETHHandlerFilterer creates a new log filterer instance of ETHHandler, bound to a specific deployed contract.
func NewETHHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHHandlerFilterer, error) {
	contract, err := bindETHHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHHandlerFilterer{contract: contract}, nil
}

// bindETHHandler binds a generic wrapper to an already deployed contract.
func bindETHHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ETHHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHHandler *ETHHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHHandler.Contract.ETHHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHHandler *ETHHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHHandler.Contract.ETHHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHHandler *ETHHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHHandler.Contract.ETHHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHHandler *ETHHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHHandler *ETHHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHHandler *ETHHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHHandler.Contract.contract.Transact(opts, method, params...)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address weth, uint256 amount) returns()
func (_ETHHandler *ETHHandlerTransactor) Withdraw(opts *bind.TransactOpts, weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHHandler.contract.Transact(opts, "withdraw", weth, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address weth, uint256 amount) returns()
func (_ETHHandler *ETHHandlerSession) Withdraw(weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHHandler.Contract.Withdraw(&_ETHHandler.TransactOpts, weth, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address weth, uint256 amount) returns()
func (_ETHHandler *ETHHandlerTransactorSession) Withdraw(weth common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHHandler.Contract.Withdraw(&_ETHHandler.TransactOpts, weth, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ETHHandler *ETHHandlerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHHandler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ETHHandler *ETHHandlerSession) Receive() (*types.Transaction, error) {
	return _ETHHandler.Contract.Receive(&_ETHHandler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ETHHandler *ETHHandlerTransactorSession) Receive() (*types.Transaction, error) {
	return _ETHHandler.Contract.Receive(&_ETHHandler.TransactOpts)
}
