// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITokenMessenger

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

// ITokenMessengerMetaData contains all meta data concerning the ITokenMessenger contract.
var ITokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_burnToken\",\"type\":\"address\"}],\"name\":\"depositForBurn\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"depositForBurnWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_destCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_mintRecipient\",\"type\":\"bytes32\"}],\"name\":\"replaceDepositForBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITokenMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMessengerMetaData.ABI instead.
var ITokenMessengerABI = ITokenMessengerMetaData.ABI

// ITokenMessenger is an auto generated Go binding around an Ethereum contract.
type ITokenMessenger struct {
	ITokenMessengerCaller     // Read-only binding to the contract
	ITokenMessengerTransactor // Write-only binding to the contract
	ITokenMessengerFilterer   // Log filterer for contract events
}

// ITokenMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMessengerSession struct {
	Contract     *ITokenMessenger  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMessengerCallerSession struct {
	Contract *ITokenMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ITokenMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMessengerTransactorSession struct {
	Contract     *ITokenMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ITokenMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMessengerRaw struct {
	Contract *ITokenMessenger // Generic contract binding to access the raw methods on
}

// ITokenMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMessengerCallerRaw struct {
	Contract *ITokenMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactorRaw struct {
	Contract *ITokenMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMessenger creates a new instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessenger(address common.Address, backend bind.ContractBackend) (*ITokenMessenger, error) {
	contract, err := bindITokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMessenger{ITokenMessengerCaller: ITokenMessengerCaller{contract: contract}, ITokenMessengerTransactor: ITokenMessengerTransactor{contract: contract}, ITokenMessengerFilterer: ITokenMessengerFilterer{contract: contract}}, nil
}

// NewITokenMessengerCaller creates a new read-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*ITokenMessengerCaller, error) {
	contract, err := bindITokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerCaller{contract: contract}, nil
}

// NewITokenMessengerTransactor creates a new write-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMessengerTransactor, error) {
	contract, err := bindITokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerTransactor{contract: contract}, nil
}

// NewITokenMessengerFilterer creates a new log filterer instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMessengerFilterer, error) {
	contract, err := bindITokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerFilterer{contract: contract}, nil
}

// bindITokenMessenger binds a generic wrapper to an already deployed contract.
func bindITokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITokenMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.ITokenMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transact(opts, method, params...)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 _amount, uint32 _destinationDomain, bytes32 _mintRecipient, address _burnToken) returns(uint64)
func (_ITokenMessenger *ITokenMessengerTransactor) DepositForBurn(opts *bind.TransactOpts, _amount *big.Int, _destinationDomain uint32, _mintRecipient [32]byte, _burnToken common.Address) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "depositForBurn", _amount, _destinationDomain, _mintRecipient, _burnToken)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 _amount, uint32 _destinationDomain, bytes32 _mintRecipient, address _burnToken) returns(uint64)
func (_ITokenMessenger *ITokenMessengerSession) DepositForBurn(_amount *big.Int, _destinationDomain uint32, _mintRecipient [32]byte, _burnToken common.Address) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurn(&_ITokenMessenger.TransactOpts, _amount, _destinationDomain, _mintRecipient, _burnToken)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 _amount, uint32 _destinationDomain, bytes32 _mintRecipient, address _burnToken) returns(uint64)
func (_ITokenMessenger *ITokenMessengerTransactorSession) DepositForBurn(_amount *big.Int, _destinationDomain uint32, _mintRecipient [32]byte, _burnToken common.Address) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurn(&_ITokenMessenger.TransactOpts, _amount, _destinationDomain, _mintRecipient, _burnToken)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 _amount, uint32 _destinationDomain, bytes32 _mintRecipient, address _burnToken, bytes32 _destinationCaller) returns(uint64)
func (_ITokenMessenger *ITokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, _amount *big.Int, _destinationDomain uint32, _mintRecipient [32]byte, _burnToken common.Address, _destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", _amount, _destinationDomain, _mintRecipient, _burnToken, _destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 _amount, uint32 _destinationDomain, bytes32 _mintRecipient, address _burnToken, bytes32 _destinationCaller) returns(uint64)
func (_ITokenMessenger *ITokenMessengerSession) DepositForBurnWithCaller(_amount *big.Int, _destinationDomain uint32, _mintRecipient [32]byte, _burnToken common.Address, _destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, _amount, _destinationDomain, _mintRecipient, _burnToken, _destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 _amount, uint32 _destinationDomain, bytes32 _mintRecipient, address _burnToken, bytes32 _destinationCaller) returns(uint64)
func (_ITokenMessenger *ITokenMessengerTransactorSession) DepositForBurnWithCaller(_amount *big.Int, _destinationDomain uint32, _mintRecipient [32]byte, _burnToken common.Address, _destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, _amount, _destinationDomain, _mintRecipient, _burnToken, _destinationCaller)
}

// ReplaceDepositForBurn is a paid mutator transaction binding the contract method 0x29a78e33.
//
// Solidity: function replaceDepositForBurn(bytes originalMessage, bytes originalAttestation, bytes32 _destCaller, bytes32 _mintRecipient) returns()
func (_ITokenMessenger *ITokenMessengerTransactor) ReplaceDepositForBurn(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, _destCaller [32]byte, _mintRecipient [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "replaceDepositForBurn", originalMessage, originalAttestation, _destCaller, _mintRecipient)
}

// ReplaceDepositForBurn is a paid mutator transaction binding the contract method 0x29a78e33.
//
// Solidity: function replaceDepositForBurn(bytes originalMessage, bytes originalAttestation, bytes32 _destCaller, bytes32 _mintRecipient) returns()
func (_ITokenMessenger *ITokenMessengerSession) ReplaceDepositForBurn(originalMessage []byte, originalAttestation []byte, _destCaller [32]byte, _mintRecipient [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ReplaceDepositForBurn(&_ITokenMessenger.TransactOpts, originalMessage, originalAttestation, _destCaller, _mintRecipient)
}

// ReplaceDepositForBurn is a paid mutator transaction binding the contract method 0x29a78e33.
//
// Solidity: function replaceDepositForBurn(bytes originalMessage, bytes originalAttestation, bytes32 _destCaller, bytes32 _mintRecipient) returns()
func (_ITokenMessenger *ITokenMessengerTransactorSession) ReplaceDepositForBurn(originalMessage []byte, originalAttestation []byte, _destCaller [32]byte, _mintRecipient [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ReplaceDepositForBurn(&_ITokenMessenger.TransactOpts, originalMessage, originalAttestation, _destCaller, _mintRecipient)
}
