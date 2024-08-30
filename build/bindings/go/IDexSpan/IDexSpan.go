// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDexSpan

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

// IDexSpanMetaData contains all meta data concerning the IDexSpan contract.
var IDexSpanMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"srcStablefromtoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcStableFromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destTokenEthPriceTimesGasPrice\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnWithGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateGasAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"parts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"flags\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"destTokenEthPriceTimesGasPrices\",\"type\":\"uint256[]\"}],\"name\":\"getExpectedReturnWithGasMulti\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"returnAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"estimateGasAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"name\":\"setBridgeAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handlerAddress\",\"type\":\"address\"}],\"name\":\"setHandlerAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserveAddress\",\"type\":\"address\"}],\"name\":\"setReserveAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dataTx\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapper\",\"type\":\"bool\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"flags\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataTx\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"isWrapper\",\"type\":\"bool\"}],\"name\":\"swapMulti\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"flags\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataTx\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"isWrapper\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapMultiWithRecipient\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"dataTx\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isWrapper\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapWithRecipient\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IDexSpanABI is the input ABI used to generate the binding from.
// Deprecated: Use IDexSpanMetaData.ABI instead.
var IDexSpanABI = IDexSpanMetaData.ABI

// IDexSpan is an auto generated Go binding around an Ethereum contract.
type IDexSpan struct {
	IDexSpanCaller     // Read-only binding to the contract
	IDexSpanTransactor // Write-only binding to the contract
	IDexSpanFilterer   // Log filterer for contract events
}

// IDexSpanCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDexSpanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDexSpanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDexSpanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDexSpanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDexSpanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDexSpanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDexSpanSession struct {
	Contract     *IDexSpan         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDexSpanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDexSpanCallerSession struct {
	Contract *IDexSpanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IDexSpanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDexSpanTransactorSession struct {
	Contract     *IDexSpanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IDexSpanRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDexSpanRaw struct {
	Contract *IDexSpan // Generic contract binding to access the raw methods on
}

// IDexSpanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDexSpanCallerRaw struct {
	Contract *IDexSpanCaller // Generic read-only contract binding to access the raw methods on
}

// IDexSpanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDexSpanTransactorRaw struct {
	Contract *IDexSpanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDexSpan creates a new instance of IDexSpan, bound to a specific deployed contract.
func NewIDexSpan(address common.Address, backend bind.ContractBackend) (*IDexSpan, error) {
	contract, err := bindIDexSpan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDexSpan{IDexSpanCaller: IDexSpanCaller{contract: contract}, IDexSpanTransactor: IDexSpanTransactor{contract: contract}, IDexSpanFilterer: IDexSpanFilterer{contract: contract}}, nil
}

// NewIDexSpanCaller creates a new read-only instance of IDexSpan, bound to a specific deployed contract.
func NewIDexSpanCaller(address common.Address, caller bind.ContractCaller) (*IDexSpanCaller, error) {
	contract, err := bindIDexSpan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDexSpanCaller{contract: contract}, nil
}

// NewIDexSpanTransactor creates a new write-only instance of IDexSpan, bound to a specific deployed contract.
func NewIDexSpanTransactor(address common.Address, transactor bind.ContractTransactor) (*IDexSpanTransactor, error) {
	contract, err := bindIDexSpan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDexSpanTransactor{contract: contract}, nil
}

// NewIDexSpanFilterer creates a new log filterer instance of IDexSpan, bound to a specific deployed contract.
func NewIDexSpanFilterer(address common.Address, filterer bind.ContractFilterer) (*IDexSpanFilterer, error) {
	contract, err := bindIDexSpan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDexSpanFilterer{contract: contract}, nil
}

// bindIDexSpan binds a generic wrapper to an already deployed contract.
func bindIDexSpan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDexSpanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDexSpan *IDexSpanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDexSpan.Contract.IDexSpanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDexSpan *IDexSpanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDexSpan.Contract.IDexSpanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDexSpan *IDexSpanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDexSpan.Contract.IDexSpanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDexSpan *IDexSpanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDexSpan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDexSpan *IDexSpanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDexSpan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDexSpan *IDexSpanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDexSpan.Contract.contract.Transact(opts, method, params...)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanCaller) GetExpectedReturn(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	var out []interface{}
	err := _IDexSpan.contract.Call(opts, &out, "getExpectedReturn", fromToken, destToken, amount, parts, flags)

	outstruct := new(struct {
		ReturnAmount *big.Int
		Distribution []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReturnAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Distribution = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanSession) GetExpectedReturn(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _IDexSpan.Contract.GetExpectedReturn(&_IDexSpan.CallOpts, fromToken, destToken, amount, parts, flags)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanCallerSession) GetExpectedReturn(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _IDexSpan.Contract.GetExpectedReturn(&_IDexSpan.CallOpts, fromToken, destToken, amount, parts, flags)
}

// GetExpectedReturnETH is a free data retrieval call binding the contract method 0x17e4e528.
//
// Solidity: function getExpectedReturnETH(address srcStablefromtoken, uint256 srcStableFromTokenAmount, uint256 parts, uint256 flags) view returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanCaller) GetExpectedReturnETH(opts *bind.CallOpts, srcStablefromtoken common.Address, srcStableFromTokenAmount *big.Int, parts *big.Int, flags *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IDexSpan.contract.Call(opts, &out, "getExpectedReturnETH", srcStablefromtoken, srcStableFromTokenAmount, parts, flags)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedReturnETH is a free data retrieval call binding the contract method 0x17e4e528.
//
// Solidity: function getExpectedReturnETH(address srcStablefromtoken, uint256 srcStableFromTokenAmount, uint256 parts, uint256 flags) view returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanSession) GetExpectedReturnETH(srcStablefromtoken common.Address, srcStableFromTokenAmount *big.Int, parts *big.Int, flags *big.Int) (*big.Int, error) {
	return _IDexSpan.Contract.GetExpectedReturnETH(&_IDexSpan.CallOpts, srcStablefromtoken, srcStableFromTokenAmount, parts, flags)
}

// GetExpectedReturnETH is a free data retrieval call binding the contract method 0x17e4e528.
//
// Solidity: function getExpectedReturnETH(address srcStablefromtoken, uint256 srcStableFromTokenAmount, uint256 parts, uint256 flags) view returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanCallerSession) GetExpectedReturnETH(srcStablefromtoken common.Address, srcStableFromTokenAmount *big.Int, parts *big.Int, flags *big.Int) (*big.Int, error) {
	return _IDexSpan.Contract.GetExpectedReturnETH(&_IDexSpan.CallOpts, srcStablefromtoken, srcStableFromTokenAmount, parts, flags)
}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanCaller) GetExpectedReturnWithGas(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	var out []interface{}
	err := _IDexSpan.contract.Call(opts, &out, "getExpectedReturnWithGas", fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)

	outstruct := new(struct {
		ReturnAmount      *big.Int
		EstimateGasAmount *big.Int
		Distribution      []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReturnAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EstimateGasAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Distribution = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanSession) GetExpectedReturnWithGas(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _IDexSpan.Contract.GetExpectedReturnWithGas(&_IDexSpan.CallOpts, fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)
}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanCallerSession) GetExpectedReturnWithGas(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _IDexSpan.Contract.GetExpectedReturnWithGas(&_IDexSpan.CallOpts, fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)
}

// GetExpectedReturnWithGasMulti is a free data retrieval call binding the contract method 0x7b33701a.
//
// Solidity: function getExpectedReturnWithGasMulti(address[] tokens, uint256 amount, uint256[] parts, uint256[] flags, uint256[] destTokenEthPriceTimesGasPrices) view returns(uint256[] returnAmounts, uint256 estimateGasAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanCaller) GetExpectedReturnWithGasMulti(opts *bind.CallOpts, tokens []common.Address, amount *big.Int, parts []*big.Int, flags []*big.Int, destTokenEthPriceTimesGasPrices []*big.Int) (struct {
	ReturnAmounts     []*big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	var out []interface{}
	err := _IDexSpan.contract.Call(opts, &out, "getExpectedReturnWithGasMulti", tokens, amount, parts, flags, destTokenEthPriceTimesGasPrices)

	outstruct := new(struct {
		ReturnAmounts     []*big.Int
		EstimateGasAmount *big.Int
		Distribution      []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReturnAmounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.EstimateGasAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Distribution = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetExpectedReturnWithGasMulti is a free data retrieval call binding the contract method 0x7b33701a.
//
// Solidity: function getExpectedReturnWithGasMulti(address[] tokens, uint256 amount, uint256[] parts, uint256[] flags, uint256[] destTokenEthPriceTimesGasPrices) view returns(uint256[] returnAmounts, uint256 estimateGasAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanSession) GetExpectedReturnWithGasMulti(tokens []common.Address, amount *big.Int, parts []*big.Int, flags []*big.Int, destTokenEthPriceTimesGasPrices []*big.Int) (struct {
	ReturnAmounts     []*big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _IDexSpan.Contract.GetExpectedReturnWithGasMulti(&_IDexSpan.CallOpts, tokens, amount, parts, flags, destTokenEthPriceTimesGasPrices)
}

// GetExpectedReturnWithGasMulti is a free data retrieval call binding the contract method 0x7b33701a.
//
// Solidity: function getExpectedReturnWithGasMulti(address[] tokens, uint256 amount, uint256[] parts, uint256[] flags, uint256[] destTokenEthPriceTimesGasPrices) view returns(uint256[] returnAmounts, uint256 estimateGasAmount, uint256[] distribution)
func (_IDexSpan *IDexSpanCallerSession) GetExpectedReturnWithGasMulti(tokens []common.Address, amount *big.Int, parts []*big.Int, flags []*big.Int, destTokenEthPriceTimesGasPrices []*big.Int) (struct {
	ReturnAmounts     []*big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _IDexSpan.Contract.GetExpectedReturnWithGasMulti(&_IDexSpan.CallOpts, tokens, amount, parts, flags, destTokenEthPriceTimesGasPrices)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address _bridgeAddress) returns(bool)
func (_IDexSpan *IDexSpanTransactor) SetBridgeAddress(opts *bind.TransactOpts, _bridgeAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "setBridgeAddress", _bridgeAddress)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address _bridgeAddress) returns(bool)
func (_IDexSpan *IDexSpanSession) SetBridgeAddress(_bridgeAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SetBridgeAddress(&_IDexSpan.TransactOpts, _bridgeAddress)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address _bridgeAddress) returns(bool)
func (_IDexSpan *IDexSpanTransactorSession) SetBridgeAddress(_bridgeAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SetBridgeAddress(&_IDexSpan.TransactOpts, _bridgeAddress)
}

// SetHandlerAddress is a paid mutator transaction binding the contract method 0xa6c2b900.
//
// Solidity: function setHandlerAddress(address _handlerAddress) returns(bool)
func (_IDexSpan *IDexSpanTransactor) SetHandlerAddress(opts *bind.TransactOpts, _handlerAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "setHandlerAddress", _handlerAddress)
}

// SetHandlerAddress is a paid mutator transaction binding the contract method 0xa6c2b900.
//
// Solidity: function setHandlerAddress(address _handlerAddress) returns(bool)
func (_IDexSpan *IDexSpanSession) SetHandlerAddress(_handlerAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SetHandlerAddress(&_IDexSpan.TransactOpts, _handlerAddress)
}

// SetHandlerAddress is a paid mutator transaction binding the contract method 0xa6c2b900.
//
// Solidity: function setHandlerAddress(address _handlerAddress) returns(bool)
func (_IDexSpan *IDexSpanTransactorSession) SetHandlerAddress(_handlerAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SetHandlerAddress(&_IDexSpan.TransactOpts, _handlerAddress)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(address _reserveAddress) returns(bool)
func (_IDexSpan *IDexSpanTransactor) SetReserveAddress(opts *bind.TransactOpts, _reserveAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "setReserveAddress", _reserveAddress)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(address _reserveAddress) returns(bool)
func (_IDexSpan *IDexSpanSession) SetReserveAddress(_reserveAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SetReserveAddress(&_IDexSpan.TransactOpts, _reserveAddress)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(address _reserveAddress) returns(bool)
func (_IDexSpan *IDexSpanTransactorSession) SetReserveAddress(_reserveAddress common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SetReserveAddress(&_IDexSpan.TransactOpts, _reserveAddress)
}

// Swap is a paid mutator transaction binding the contract method 0x67c52efe.
//
// Solidity: function swap(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256 flags, bytes dataTx, bool isWrapper) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, flags *big.Int, dataTx []byte, isWrapper bool) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "swap", fromToken, destToken, amount, minReturn, flags, dataTx, isWrapper)
}

// Swap is a paid mutator transaction binding the contract method 0x67c52efe.
//
// Solidity: function swap(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256 flags, bytes dataTx, bool isWrapper) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanSession) Swap(fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, flags *big.Int, dataTx []byte, isWrapper bool) (*types.Transaction, error) {
	return _IDexSpan.Contract.Swap(&_IDexSpan.TransactOpts, fromToken, destToken, amount, minReturn, flags, dataTx, isWrapper)
}

// Swap is a paid mutator transaction binding the contract method 0x67c52efe.
//
// Solidity: function swap(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256 flags, bytes dataTx, bool isWrapper) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactorSession) Swap(fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, flags *big.Int, dataTx []byte, isWrapper bool) (*types.Transaction, error) {
	return _IDexSpan.Contract.Swap(&_IDexSpan.TransactOpts, fromToken, destToken, amount, minReturn, flags, dataTx, isWrapper)
}

// SwapMulti is a paid mutator transaction binding the contract method 0xd835b069.
//
// Solidity: function swapMulti(address[] tokens, uint256 amount, uint256 minReturn, uint256[] flags, bytes[] dataTx, bool isWrapper) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactor) SwapMulti(opts *bind.TransactOpts, tokens []common.Address, amount *big.Int, minReturn *big.Int, flags []*big.Int, dataTx [][]byte, isWrapper bool) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "swapMulti", tokens, amount, minReturn, flags, dataTx, isWrapper)
}

// SwapMulti is a paid mutator transaction binding the contract method 0xd835b069.
//
// Solidity: function swapMulti(address[] tokens, uint256 amount, uint256 minReturn, uint256[] flags, bytes[] dataTx, bool isWrapper) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanSession) SwapMulti(tokens []common.Address, amount *big.Int, minReturn *big.Int, flags []*big.Int, dataTx [][]byte, isWrapper bool) (*types.Transaction, error) {
	return _IDexSpan.Contract.SwapMulti(&_IDexSpan.TransactOpts, tokens, amount, minReturn, flags, dataTx, isWrapper)
}

// SwapMulti is a paid mutator transaction binding the contract method 0xd835b069.
//
// Solidity: function swapMulti(address[] tokens, uint256 amount, uint256 minReturn, uint256[] flags, bytes[] dataTx, bool isWrapper) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactorSession) SwapMulti(tokens []common.Address, amount *big.Int, minReturn *big.Int, flags []*big.Int, dataTx [][]byte, isWrapper bool) (*types.Transaction, error) {
	return _IDexSpan.Contract.SwapMulti(&_IDexSpan.TransactOpts, tokens, amount, minReturn, flags, dataTx, isWrapper)
}

// SwapMultiWithRecipient is a paid mutator transaction binding the contract method 0xe738aa8d.
//
// Solidity: function swapMultiWithRecipient(address[] tokens, uint256 amount, uint256 minReturn, uint256[] flags, bytes[] dataTx, bool isWrapper, address recipient) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactor) SwapMultiWithRecipient(opts *bind.TransactOpts, tokens []common.Address, amount *big.Int, minReturn *big.Int, flags []*big.Int, dataTx [][]byte, isWrapper bool, recipient common.Address) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "swapMultiWithRecipient", tokens, amount, minReturn, flags, dataTx, isWrapper, recipient)
}

// SwapMultiWithRecipient is a paid mutator transaction binding the contract method 0xe738aa8d.
//
// Solidity: function swapMultiWithRecipient(address[] tokens, uint256 amount, uint256 minReturn, uint256[] flags, bytes[] dataTx, bool isWrapper, address recipient) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanSession) SwapMultiWithRecipient(tokens []common.Address, amount *big.Int, minReturn *big.Int, flags []*big.Int, dataTx [][]byte, isWrapper bool, recipient common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SwapMultiWithRecipient(&_IDexSpan.TransactOpts, tokens, amount, minReturn, flags, dataTx, isWrapper, recipient)
}

// SwapMultiWithRecipient is a paid mutator transaction binding the contract method 0xe738aa8d.
//
// Solidity: function swapMultiWithRecipient(address[] tokens, uint256 amount, uint256 minReturn, uint256[] flags, bytes[] dataTx, bool isWrapper, address recipient) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactorSession) SwapMultiWithRecipient(tokens []common.Address, amount *big.Int, minReturn *big.Int, flags []*big.Int, dataTx [][]byte, isWrapper bool, recipient common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SwapMultiWithRecipient(&_IDexSpan.TransactOpts, tokens, amount, minReturn, flags, dataTx, isWrapper, recipient)
}

// SwapWithRecipient is a paid mutator transaction binding the contract method 0xe31719d6.
//
// Solidity: function swapWithRecipient(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256 flags, bytes dataTx, bool isWrapper, address recipient) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactor) SwapWithRecipient(opts *bind.TransactOpts, fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, flags *big.Int, dataTx []byte, isWrapper bool, recipient common.Address) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "swapWithRecipient", fromToken, destToken, amount, minReturn, flags, dataTx, isWrapper, recipient)
}

// SwapWithRecipient is a paid mutator transaction binding the contract method 0xe31719d6.
//
// Solidity: function swapWithRecipient(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256 flags, bytes dataTx, bool isWrapper, address recipient) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanSession) SwapWithRecipient(fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, flags *big.Int, dataTx []byte, isWrapper bool, recipient common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SwapWithRecipient(&_IDexSpan.TransactOpts, fromToken, destToken, amount, minReturn, flags, dataTx, isWrapper, recipient)
}

// SwapWithRecipient is a paid mutator transaction binding the contract method 0xe31719d6.
//
// Solidity: function swapWithRecipient(address fromToken, address destToken, uint256 amount, uint256 minReturn, uint256 flags, bytes dataTx, bool isWrapper, address recipient) payable returns(uint256 returnAmount)
func (_IDexSpan *IDexSpanTransactorSession) SwapWithRecipient(fromToken common.Address, destToken common.Address, amount *big.Int, minReturn *big.Int, flags *big.Int, dataTx []byte, isWrapper bool, recipient common.Address) (*types.Transaction, error) {
	return _IDexSpan.Contract.SwapWithRecipient(&_IDexSpan.TransactOpts, fromToken, destToken, amount, minReturn, flags, dataTx, isWrapper, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) payable returns(bool)
func (_IDexSpan *IDexSpanTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDexSpan.contract.Transact(opts, "withdraw", tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) payable returns(bool)
func (_IDexSpan *IDexSpanSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDexSpan.Contract.Withdraw(&_IDexSpan.TransactOpts, tokenAddress, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amount) payable returns(bool)
func (_IDexSpan *IDexSpanTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IDexSpan.Contract.Withdraw(&_IDexSpan.TransactOpts, tokenAddress, recipient, amount)
}
