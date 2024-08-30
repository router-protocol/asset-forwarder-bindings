// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IERC20PermitUpgradeable

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

// IERC20PermitUpgradeableMetaData contains all meta data concerning the IERC20PermitUpgradeable contract.
var IERC20PermitUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20PermitUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20PermitUpgradeableMetaData.ABI instead.
var IERC20PermitUpgradeableABI = IERC20PermitUpgradeableMetaData.ABI

// IERC20PermitUpgradeable is an auto generated Go binding around an Ethereum contract.
type IERC20PermitUpgradeable struct {
	IERC20PermitUpgradeableCaller     // Read-only binding to the contract
	IERC20PermitUpgradeableTransactor // Write-only binding to the contract
	IERC20PermitUpgradeableFilterer   // Log filterer for contract events
}

// IERC20PermitUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20PermitUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20PermitUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20PermitUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20PermitUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20PermitUpgradeableSession struct {
	Contract     *IERC20PermitUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC20PermitUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20PermitUpgradeableCallerSession struct {
	Contract *IERC20PermitUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IERC20PermitUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20PermitUpgradeableTransactorSession struct {
	Contract     *IERC20PermitUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IERC20PermitUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20PermitUpgradeableRaw struct {
	Contract *IERC20PermitUpgradeable // Generic contract binding to access the raw methods on
}

// IERC20PermitUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20PermitUpgradeableCallerRaw struct {
	Contract *IERC20PermitUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20PermitUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20PermitUpgradeableTransactorRaw struct {
	Contract *IERC20PermitUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20PermitUpgradeable creates a new instance of IERC20PermitUpgradeable, bound to a specific deployed contract.
func NewIERC20PermitUpgradeable(address common.Address, backend bind.ContractBackend) (*IERC20PermitUpgradeable, error) {
	contract, err := bindIERC20PermitUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitUpgradeable{IERC20PermitUpgradeableCaller: IERC20PermitUpgradeableCaller{contract: contract}, IERC20PermitUpgradeableTransactor: IERC20PermitUpgradeableTransactor{contract: contract}, IERC20PermitUpgradeableFilterer: IERC20PermitUpgradeableFilterer{contract: contract}}, nil
}

// NewIERC20PermitUpgradeableCaller creates a new read-only instance of IERC20PermitUpgradeable, bound to a specific deployed contract.
func NewIERC20PermitUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC20PermitUpgradeableCaller, error) {
	contract, err := bindIERC20PermitUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitUpgradeableCaller{contract: contract}, nil
}

// NewIERC20PermitUpgradeableTransactor creates a new write-only instance of IERC20PermitUpgradeable, bound to a specific deployed contract.
func NewIERC20PermitUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20PermitUpgradeableTransactor, error) {
	contract, err := bindIERC20PermitUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitUpgradeableTransactor{contract: contract}, nil
}

// NewIERC20PermitUpgradeableFilterer creates a new log filterer instance of IERC20PermitUpgradeable, bound to a specific deployed contract.
func NewIERC20PermitUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20PermitUpgradeableFilterer, error) {
	contract, err := bindIERC20PermitUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20PermitUpgradeableFilterer{contract: contract}, nil
}

// bindIERC20PermitUpgradeable binds a generic wrapper to an already deployed contract.
func bindIERC20PermitUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20PermitUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20PermitUpgradeable.Contract.IERC20PermitUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.Contract.IERC20PermitUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.Contract.IERC20PermitUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20PermitUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC20PermitUpgradeable.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20PermitUpgradeable.Contract.DOMAINSEPARATOR(&_IERC20PermitUpgradeable.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IERC20PermitUpgradeable.Contract.DOMAINSEPARATOR(&_IERC20PermitUpgradeable.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20PermitUpgradeable.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20PermitUpgradeable.Contract.Nonces(&_IERC20PermitUpgradeable.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _IERC20PermitUpgradeable.Contract.Nonces(&_IERC20PermitUpgradeable.CallOpts, owner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.Contract.Permit(&_IERC20PermitUpgradeable.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_IERC20PermitUpgradeable *IERC20PermitUpgradeableTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _IERC20PermitUpgradeable.Contract.Permit(&_IERC20PermitUpgradeable.TransactOpts, owner, spender, value, deadline, v, r, s)
}
