// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ICCTPBridge

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

// IDexSpanSwapPayload is an auto generated low-level Go binding around an user-defined struct.
type IDexSpanSwapPayload struct {
	Flags       []*big.Int
	MinToAmount *big.Int
	Tokens      []common.Address
	DataTx      [][]byte
}

// ICCTPBridgeMetaData contains all meta data concerning the ICCTPBridge contract.
var ICCTPBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CCTPNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUpdateType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongReturnToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"iDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"flags\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataTx\",\"type\":\"bytes[]\"}],\"internalType\":\"structIDexSpan.SwapPayload\",\"name\":\"swapPayload\",\"type\":\"tuple\"}],\"name\":\"swapAndIDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ICCTPBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ICCTPBridgeMetaData.ABI instead.
var ICCTPBridgeABI = ICCTPBridgeMetaData.ABI

// ICCTPBridge is an auto generated Go binding around an Ethereum contract.
type ICCTPBridge struct {
	ICCTPBridgeCaller     // Read-only binding to the contract
	ICCTPBridgeTransactor // Write-only binding to the contract
	ICCTPBridgeFilterer   // Log filterer for contract events
}

// ICCTPBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICCTPBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICCTPBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICCTPBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICCTPBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICCTPBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICCTPBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICCTPBridgeSession struct {
	Contract     *ICCTPBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICCTPBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICCTPBridgeCallerSession struct {
	Contract *ICCTPBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ICCTPBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICCTPBridgeTransactorSession struct {
	Contract     *ICCTPBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ICCTPBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICCTPBridgeRaw struct {
	Contract *ICCTPBridge // Generic contract binding to access the raw methods on
}

// ICCTPBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICCTPBridgeCallerRaw struct {
	Contract *ICCTPBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ICCTPBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICCTPBridgeTransactorRaw struct {
	Contract *ICCTPBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICCTPBridge creates a new instance of ICCTPBridge, bound to a specific deployed contract.
func NewICCTPBridge(address common.Address, backend bind.ContractBackend) (*ICCTPBridge, error) {
	contract, err := bindICCTPBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICCTPBridge{ICCTPBridgeCaller: ICCTPBridgeCaller{contract: contract}, ICCTPBridgeTransactor: ICCTPBridgeTransactor{contract: contract}, ICCTPBridgeFilterer: ICCTPBridgeFilterer{contract: contract}}, nil
}

// NewICCTPBridgeCaller creates a new read-only instance of ICCTPBridge, bound to a specific deployed contract.
func NewICCTPBridgeCaller(address common.Address, caller bind.ContractCaller) (*ICCTPBridgeCaller, error) {
	contract, err := bindICCTPBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICCTPBridgeCaller{contract: contract}, nil
}

// NewICCTPBridgeTransactor creates a new write-only instance of ICCTPBridge, bound to a specific deployed contract.
func NewICCTPBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ICCTPBridgeTransactor, error) {
	contract, err := bindICCTPBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICCTPBridgeTransactor{contract: contract}, nil
}

// NewICCTPBridgeFilterer creates a new log filterer instance of ICCTPBridge, bound to a specific deployed contract.
func NewICCTPBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ICCTPBridgeFilterer, error) {
	contract, err := bindICCTPBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICCTPBridgeFilterer{contract: contract}, nil
}

// bindICCTPBridge binds a generic wrapper to an already deployed contract.
func bindICCTPBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICCTPBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICCTPBridge *ICCTPBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICCTPBridge.Contract.ICCTPBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICCTPBridge *ICCTPBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.ICCTPBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICCTPBridge *ICCTPBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.ICCTPBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICCTPBridge *ICCTPBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICCTPBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICCTPBridge *ICCTPBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICCTPBridge *ICCTPBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.contract.Transact(opts, method, params...)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x67d70e13.
//
// Solidity: function iDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data) payable returns()
func (_ICCTPBridge *ICCTPBridgeTransactor) IDepositUSDC(opts *bind.TransactOpts, txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _ICCTPBridge.contract.Transact(opts, "iDepositUSDC", txType, partnerId, destChainIdBytes, amount, recipient, data)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x67d70e13.
//
// Solidity: function iDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data) payable returns()
func (_ICCTPBridge *ICCTPBridgeSession) IDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.IDepositUSDC(&_ICCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x67d70e13.
//
// Solidity: function iDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data) payable returns()
func (_ICCTPBridge *ICCTPBridgeTransactorSession) IDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.IDepositUSDC(&_ICCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data)
}

// SwapAndIDepositUSDC is a paid mutator transaction binding the contract method 0x9d9c44c9.
//
// Solidity: function swapAndIDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data, (uint256[],uint256,address[],bytes[]) swapPayload) payable returns()
func (_ICCTPBridge *ICCTPBridgeTransactor) SwapAndIDepositUSDC(opts *bind.TransactOpts, txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte, swapPayload IDexSpanSwapPayload) (*types.Transaction, error) {
	return _ICCTPBridge.contract.Transact(opts, "swapAndIDepositUSDC", txType, partnerId, destChainIdBytes, amount, recipient, data, swapPayload)
}

// SwapAndIDepositUSDC is a paid mutator transaction binding the contract method 0x9d9c44c9.
//
// Solidity: function swapAndIDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data, (uint256[],uint256,address[],bytes[]) swapPayload) payable returns()
func (_ICCTPBridge *ICCTPBridgeSession) SwapAndIDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte, swapPayload IDexSpanSwapPayload) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.SwapAndIDepositUSDC(&_ICCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data, swapPayload)
}

// SwapAndIDepositUSDC is a paid mutator transaction binding the contract method 0x9d9c44c9.
//
// Solidity: function swapAndIDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data, (uint256[],uint256,address[],bytes[]) swapPayload) payable returns()
func (_ICCTPBridge *ICCTPBridgeTransactorSession) SwapAndIDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte, swapPayload IDexSpanSwapPayload) (*types.Transaction, error) {
	return _ICCTPBridge.Contract.SwapAndIDepositUSDC(&_ICCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data, swapPayload)
}

// ICCTPBridgeIUSDCDepositedIterator is returned from FilterIUSDCDeposited and is used to iterate over the raw logs and unpacked data for IUSDCDeposited events raised by the ICCTPBridge contract.
type ICCTPBridgeIUSDCDepositedIterator struct {
	Event *ICCTPBridgeIUSDCDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICCTPBridgeIUSDCDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICCTPBridgeIUSDCDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICCTPBridgeIUSDCDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICCTPBridgeIUSDCDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICCTPBridgeIUSDCDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICCTPBridgeIUSDCDeposited represents a IUSDCDeposited event raised by the ICCTPBridge contract.
type ICCTPBridgeIUSDCDeposited struct {
	TxType           uint8
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	UsdcNonce        *big.Int
	SrcToken         common.Address
	Recipient        [32]byte
	Depositor        common.Address
	Data             []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIUSDCDeposited is a free log retrieval operation binding the contract event 0xa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a526.
//
// Solidity: event iUSDCDeposited(uint8 txType, uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor, bytes data)
func (_ICCTPBridge *ICCTPBridgeFilterer) FilterIUSDCDeposited(opts *bind.FilterOpts) (*ICCTPBridgeIUSDCDepositedIterator, error) {

	logs, sub, err := _ICCTPBridge.contract.FilterLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return &ICCTPBridgeIUSDCDepositedIterator{contract: _ICCTPBridge.contract, event: "iUSDCDeposited", logs: logs, sub: sub}, nil
}

// WatchIUSDCDeposited is a free log subscription operation binding the contract event 0xa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a526.
//
// Solidity: event iUSDCDeposited(uint8 txType, uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor, bytes data)
func (_ICCTPBridge *ICCTPBridgeFilterer) WatchIUSDCDeposited(opts *bind.WatchOpts, sink chan<- *ICCTPBridgeIUSDCDeposited) (event.Subscription, error) {

	logs, sub, err := _ICCTPBridge.contract.WatchLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICCTPBridgeIUSDCDeposited)
				if err := _ICCTPBridge.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIUSDCDeposited is a log parse operation binding the contract event 0xa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a526.
//
// Solidity: event iUSDCDeposited(uint8 txType, uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor, bytes data)
func (_ICCTPBridge *ICCTPBridgeFilterer) ParseIUSDCDeposited(log types.Log) (*ICCTPBridgeIUSDCDeposited, error) {
	event := new(ICCTPBridgeIUSDCDeposited)
	if err := _ICCTPBridge.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
