// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FetchLiquidity

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

// FetchLiquiditydexResponse is an auto generated low-level Go binding around an user-defined struct.
type FetchLiquiditydexResponse struct {
	FromToken    common.Address
	DestToken    common.Address
	Reserve0     *big.Int
	Reserve1     *big.Int
	ExchangeCode *big.Int
}

// FetchLiquidityMetaData contains all meta data concerning the FetchLiquidity contract.
var FetchLiquidityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable[2][]\",\"name\":\"tokensIn\",\"type\":\"address[2][]\"},{\"internalType\":\"uint256[]\",\"name\":\"exchangeIds\",\"type\":\"uint256[]\"}],\"name\":\"getLiquidityReserves\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_destToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_exchangeCode\",\"type\":\"uint112\"}],\"internalType\":\"structFetchLiquidity.dexResponse[]\",\"name\":\"response\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109c2806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631e8f67a914610030575b600080fd5b61004361003e366004610689565b610059565b6040516100509190610751565b60405180910390f35b60606100f68383808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250604080516020808d02820181019092528b815294508b93508a925082919085015b828210156100ec57604080518082018252908084028701906002908390839080828437600092019190915250505081526001909101906020016100ae565b50505050506100ff565b95945050505050565b6060600083518351610111919061081c565b67ffffffffffffffff81111561012957610129610833565b6040519080825280602002602001820160405280156101a057816020015b6040805160a0810182526000808252602080830182905292820181905260608201819052608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816101475790505b5090506000805b85518110156102745760008682815181106101c4576101c4610862565b602002602001015190506000806101da83610280565b9150915060006101eb89848461034b565b905060005b815181101561024e5781818151811061020b5761020b610862565b60200260200101518882896102209190610891565b8151811061023057610230610862565b60200260200101819052508080610246906108a4565b9150506101f0565b50805161025b9087610891565b955050505050808061026c906108a4565b9150506101a7565b50909150505b92915050565b600080600183036102aa575073d590cc180601aecd6eeadd9b7f2b7611519544f492600192509050565b600283036102d1575073689abaeeed3f0bb3585773192e23224cac25dd4192600292509050565b600383036102f85750731cc79ecb3a6f9b6d6faf7298ec6d8667e814d59292600392509050565b6004830361031f575073462c98cae5affeed576c98a55daa922604e2d87592600492509050565b60058303610346575073c35dadb65012ec5796536bd9864ed8773abc74c492600592509050565b915091565b6060835167ffffffffffffffff81111561036757610367610833565b6040519080825280602002602001820160405280156103de57816020015b6040805160a0810182526000808252602080830182905292820181905260608201819052608082015282527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092019101816103855790505b50905060005b845181101561051e5760008060008061045e89868151811061040857610408610862565b602002602001015160006002811061042257610422610862565b60200201518a878151811061043957610439610862565b602002602001015160016002811061045357610453610862565b60200201518a610526565b5093509350935093506040518060a001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff168152602001836dffffffffffffffffffffffffffff168152602001826dffffffffffffffffffffffffffff168152602001886dffffffffffffffffffffffffffff168152508686815181106104fc576104fc610862565b6020026020010181905250505050508080610516906108a4565b9150506103e4565b509392505050565b6040517fe6a4390500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015283811660248301526000918291829182918291829188169063e6a4390590604401602060405180830381865afa1580156105a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105cb91906108dc565b905073ffffffffffffffffffffffffffffffffffffffff81161561066b578073ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b8152600401606060405180830381865afa158015610634573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610658919061093c565b508a9750899650909450925061067e9050565b8888600080600095509550955095509550505b939792965093509350565b6000806000806040858703121561069f57600080fd5b843567ffffffffffffffff808211156106b757600080fd5b818701915087601f8301126106cb57600080fd5b8135818111156106da57600080fd5b8860208260061b85010111156106ef57600080fd5b60209283019650945090860135908082111561070a57600080fd5b818701915087601f83011261071e57600080fd5b81358181111561072d57600080fd5b8860208260051b850101111561074257600080fd5b95989497505060200194505050565b602080825282518282018190526000919060409081850190868401855b828110156107e0578151805173ffffffffffffffffffffffffffffffffffffffff9081168652878201511687860152858101516dffffffffffffffffffffffffffff90811687870152606080830151821690870152608091820151169085015260a0909301929085019060010161076e565b5091979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761027a5761027a6107ed565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8082018082111561027a5761027a6107ed565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108d5576108d56107ed565b5060010190565b6000602082840312156108ee57600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461091257600080fd5b9392505050565b80516dffffffffffffffffffffffffffff8116811461093757600080fd5b919050565b60008060006060848603121561095157600080fd5b61095a84610919565b925061096860208501610919565b9150604084015163ffffffff8116811461098157600080fd5b80915050925092509256fea2646970667358221220811fb7072c7b0c0919172abbe17435dd696fd38c24c59b22873fb341c2ead9fe64736f6c63430008140033",
}

// FetchLiquidityABI is the input ABI used to generate the binding from.
// Deprecated: Use FetchLiquidityMetaData.ABI instead.
var FetchLiquidityABI = FetchLiquidityMetaData.ABI

// FetchLiquidityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FetchLiquidityMetaData.Bin instead.
var FetchLiquidityBin = FetchLiquidityMetaData.Bin

// DeployFetchLiquidity deploys a new Ethereum contract, binding an instance of FetchLiquidity to it.
func DeployFetchLiquidity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FetchLiquidity, error) {
	parsed, err := FetchLiquidityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FetchLiquidityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FetchLiquidity{FetchLiquidityCaller: FetchLiquidityCaller{contract: contract}, FetchLiquidityTransactor: FetchLiquidityTransactor{contract: contract}, FetchLiquidityFilterer: FetchLiquidityFilterer{contract: contract}}, nil
}

// FetchLiquidity is an auto generated Go binding around an Ethereum contract.
type FetchLiquidity struct {
	FetchLiquidityCaller     // Read-only binding to the contract
	FetchLiquidityTransactor // Write-only binding to the contract
	FetchLiquidityFilterer   // Log filterer for contract events
}

// FetchLiquidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type FetchLiquidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FetchLiquidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FetchLiquidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FetchLiquidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FetchLiquidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FetchLiquiditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FetchLiquiditySession struct {
	Contract     *FetchLiquidity   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FetchLiquidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FetchLiquidityCallerSession struct {
	Contract *FetchLiquidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FetchLiquidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FetchLiquidityTransactorSession struct {
	Contract     *FetchLiquidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FetchLiquidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type FetchLiquidityRaw struct {
	Contract *FetchLiquidity // Generic contract binding to access the raw methods on
}

// FetchLiquidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FetchLiquidityCallerRaw struct {
	Contract *FetchLiquidityCaller // Generic read-only contract binding to access the raw methods on
}

// FetchLiquidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FetchLiquidityTransactorRaw struct {
	Contract *FetchLiquidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFetchLiquidity creates a new instance of FetchLiquidity, bound to a specific deployed contract.
func NewFetchLiquidity(address common.Address, backend bind.ContractBackend) (*FetchLiquidity, error) {
	contract, err := bindFetchLiquidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FetchLiquidity{FetchLiquidityCaller: FetchLiquidityCaller{contract: contract}, FetchLiquidityTransactor: FetchLiquidityTransactor{contract: contract}, FetchLiquidityFilterer: FetchLiquidityFilterer{contract: contract}}, nil
}

// NewFetchLiquidityCaller creates a new read-only instance of FetchLiquidity, bound to a specific deployed contract.
func NewFetchLiquidityCaller(address common.Address, caller bind.ContractCaller) (*FetchLiquidityCaller, error) {
	contract, err := bindFetchLiquidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FetchLiquidityCaller{contract: contract}, nil
}

// NewFetchLiquidityTransactor creates a new write-only instance of FetchLiquidity, bound to a specific deployed contract.
func NewFetchLiquidityTransactor(address common.Address, transactor bind.ContractTransactor) (*FetchLiquidityTransactor, error) {
	contract, err := bindFetchLiquidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FetchLiquidityTransactor{contract: contract}, nil
}

// NewFetchLiquidityFilterer creates a new log filterer instance of FetchLiquidity, bound to a specific deployed contract.
func NewFetchLiquidityFilterer(address common.Address, filterer bind.ContractFilterer) (*FetchLiquidityFilterer, error) {
	contract, err := bindFetchLiquidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FetchLiquidityFilterer{contract: contract}, nil
}

// bindFetchLiquidity binds a generic wrapper to an already deployed contract.
func bindFetchLiquidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FetchLiquidityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FetchLiquidity *FetchLiquidityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FetchLiquidity.Contract.FetchLiquidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FetchLiquidity *FetchLiquidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FetchLiquidity.Contract.FetchLiquidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FetchLiquidity *FetchLiquidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FetchLiquidity.Contract.FetchLiquidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FetchLiquidity *FetchLiquidityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FetchLiquidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FetchLiquidity *FetchLiquidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FetchLiquidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FetchLiquidity *FetchLiquidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FetchLiquidity.Contract.contract.Transact(opts, method, params...)
}

// GetLiquidityReserves is a free data retrieval call binding the contract method 0x1e8f67a9.
//
// Solidity: function getLiquidityReserves(address[2][] tokensIn, uint256[] exchangeIds) view returns((address,address,uint112,uint112,uint112)[] response)
func (_FetchLiquidity *FetchLiquidityCaller) GetLiquidityReserves(opts *bind.CallOpts, tokensIn [][2]common.Address, exchangeIds []*big.Int) ([]FetchLiquiditydexResponse, error) {
	var out []interface{}
	err := _FetchLiquidity.contract.Call(opts, &out, "getLiquidityReserves", tokensIn, exchangeIds)

	if err != nil {
		return *new([]FetchLiquiditydexResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]FetchLiquiditydexResponse)).(*[]FetchLiquiditydexResponse)

	return out0, err

}

// GetLiquidityReserves is a free data retrieval call binding the contract method 0x1e8f67a9.
//
// Solidity: function getLiquidityReserves(address[2][] tokensIn, uint256[] exchangeIds) view returns((address,address,uint112,uint112,uint112)[] response)
func (_FetchLiquidity *FetchLiquiditySession) GetLiquidityReserves(tokensIn [][2]common.Address, exchangeIds []*big.Int) ([]FetchLiquiditydexResponse, error) {
	return _FetchLiquidity.Contract.GetLiquidityReserves(&_FetchLiquidity.CallOpts, tokensIn, exchangeIds)
}

// GetLiquidityReserves is a free data retrieval call binding the contract method 0x1e8f67a9.
//
// Solidity: function getLiquidityReserves(address[2][] tokensIn, uint256[] exchangeIds) view returns((address,address,uint112,uint112,uint112)[] response)
func (_FetchLiquidity *FetchLiquidityCallerSession) GetLiquidityReserves(tokensIn [][2]common.Address, exchangeIds []*big.Int) ([]FetchLiquiditydexResponse, error) {
	return _FetchLiquidity.Contract.GetLiquidityReserves(&_FetchLiquidity.CallOpts, tokensIn, exchangeIds)
}
