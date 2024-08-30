// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Interactor

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

// InteractorMetaData contains all meta data concerning the Interactor contract.
var InteractorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenSent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"handleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061040b806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d00a2d5f14610030575b600080fd5b61004361003e36600461021b565b610045565b005b6000808280602001905181019061005c91906102db565b915091506000632d1e0c02826040516024016100789190610368565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090506000808473ffffffffffffffffffffffffffffffffffffffff16836040516100e391906103b9565b6000604051808303816000865af19150503d8060008114610120576040519150601f19603f3d011682016040523d82523d6000602084013e610125565b606091505b5050505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461015457600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156101cd576101cd610157565b604052919050565b600067ffffffffffffffff8211156101ef576101ef610157565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b60008060006060848603121561023057600080fd5b833561023b81610132565b925060208401359150604084013567ffffffffffffffff81111561025e57600080fd5b8401601f8101861361026f57600080fd5b803561028261027d826101d5565b610186565b81815287602083850101111561029757600080fd5b816020840160208301376000602083830101528093505050509250925092565b60005b838110156102d25781810151838201526020016102ba565b50506000910152565b600080604083850312156102ee57600080fd5b82516102f981610132565b602084015190925067ffffffffffffffff81111561031657600080fd5b8301601f8101851361032757600080fd5b805161033561027d826101d5565b81815286602083850101111561034a57600080fd5b61035b8260208301602086016102b7565b8093505050509250929050565b60208152600082518060208401526103878160408501602087016102b7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600082516103cb8184602087016102b7565b919091019291505056fea2646970667358221220355260580d73932545f6cb894eefe93ddcc408231148bde4150a5ded290231c664736f6c63430008140033",
}

// InteractorABI is the input ABI used to generate the binding from.
// Deprecated: Use InteractorMetaData.ABI instead.
var InteractorABI = InteractorMetaData.ABI

// InteractorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InteractorMetaData.Bin instead.
var InteractorBin = InteractorMetaData.Bin

// DeployInteractor deploys a new Ethereum contract, binding an instance of Interactor to it.
func DeployInteractor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Interactor, error) {
	parsed, err := InteractorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InteractorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Interactor{InteractorCaller: InteractorCaller{contract: contract}, InteractorTransactor: InteractorTransactor{contract: contract}, InteractorFilterer: InteractorFilterer{contract: contract}}, nil
}

// Interactor is an auto generated Go binding around an Ethereum contract.
type Interactor struct {
	InteractorCaller     // Read-only binding to the contract
	InteractorTransactor // Write-only binding to the contract
	InteractorFilterer   // Log filterer for contract events
}

// InteractorCaller is an auto generated read-only Go binding around an Ethereum contract.
type InteractorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InteractorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InteractorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InteractorSession struct {
	Contract     *Interactor       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InteractorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InteractorCallerSession struct {
	Contract *InteractorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InteractorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InteractorTransactorSession struct {
	Contract     *InteractorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InteractorRaw is an auto generated low-level Go binding around an Ethereum contract.
type InteractorRaw struct {
	Contract *Interactor // Generic contract binding to access the raw methods on
}

// InteractorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InteractorCallerRaw struct {
	Contract *InteractorCaller // Generic read-only contract binding to access the raw methods on
}

// InteractorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InteractorTransactorRaw struct {
	Contract *InteractorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInteractor creates a new instance of Interactor, bound to a specific deployed contract.
func NewInteractor(address common.Address, backend bind.ContractBackend) (*Interactor, error) {
	contract, err := bindInteractor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Interactor{InteractorCaller: InteractorCaller{contract: contract}, InteractorTransactor: InteractorTransactor{contract: contract}, InteractorFilterer: InteractorFilterer{contract: contract}}, nil
}

// NewInteractorCaller creates a new read-only instance of Interactor, bound to a specific deployed contract.
func NewInteractorCaller(address common.Address, caller bind.ContractCaller) (*InteractorCaller, error) {
	contract, err := bindInteractor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InteractorCaller{contract: contract}, nil
}

// NewInteractorTransactor creates a new write-only instance of Interactor, bound to a specific deployed contract.
func NewInteractorTransactor(address common.Address, transactor bind.ContractTransactor) (*InteractorTransactor, error) {
	contract, err := bindInteractor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InteractorTransactor{contract: contract}, nil
}

// NewInteractorFilterer creates a new log filterer instance of Interactor, bound to a specific deployed contract.
func NewInteractorFilterer(address common.Address, filterer bind.ContractFilterer) (*InteractorFilterer, error) {
	contract, err := bindInteractor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InteractorFilterer{contract: contract}, nil
}

// bindInteractor binds a generic wrapper to an already deployed contract.
func bindInteractor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InteractorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interactor *InteractorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interactor.Contract.InteractorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interactor *InteractorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interactor.Contract.InteractorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interactor *InteractorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interactor.Contract.InteractorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interactor *InteractorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interactor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interactor *InteractorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interactor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interactor *InteractorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interactor.Contract.contract.Transact(opts, method, params...)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) returns()
func (_Interactor *InteractorTransactor) HandleMessage(opts *bind.TransactOpts, tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Interactor.contract.Transact(opts, "handleMessage", tokenSent, amount, message)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) returns()
func (_Interactor *InteractorSession) HandleMessage(tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Interactor.Contract.HandleMessage(&_Interactor.TransactOpts, tokenSent, amount, message)
}

// HandleMessage is a paid mutator transaction binding the contract method 0xd00a2d5f.
//
// Solidity: function handleMessage(address tokenSent, uint256 amount, bytes message) returns()
func (_Interactor *InteractorTransactorSession) HandleMessage(tokenSent common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _Interactor.Contract.HandleMessage(&_Interactor.TransactOpts, tokenSent, amount, message)
}
