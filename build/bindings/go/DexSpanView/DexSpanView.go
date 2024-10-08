// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DexSpanView

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

// DexSpanViewMetaData contains all meta data concerning the DexSpanView contract.
var DexSpanViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"parts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destTokenEthPriceTimesGasPrice\",\"type\":\"uint256\"}],\"name\":\"getExpectedReturnWithGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimateGasAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeAddress\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wnativeAddress\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001c600033610021565b6100d1565b61002b828261002f565b5050565b60008281526003602090815260408083206001600160a01b038516845290915290205460ff1661002b5760008281526003602090815260408083206001600160a01b03851684529091529020805460ff1916600117905561008d3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b611c5880620000e16000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806336568abe11610081578063a217fddf1161005b578063a217fddf1461021d578063d547741f14610225578063e2a4ac2d1461023857600080fd5b806336568abe146101a25780638373f265146101b557806391d14854146101d757600080fd5b8063248a9ca3116100b2578063248a9ca3146101175780632e476337146101485780632f2ff15d1461018d57600080fd5b806301ffc9a7146100ce578063085e2c5b146100f6575b600080fd5b6100e16100dc366004611761565b610258565b60405190151581526020015b60405180910390f35b6101096101043660046117c5565b6102f1565b6040516100ed929190611851565b61013a610125366004611872565b60009081526003602052604090206001015490565b6040519081526020016100ed565b6002546101689073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ed565b6101a061019b36600461188b565b610314565b005b6101a06101b036600461188b565b61033e565b6101c86101c33660046118bb565b6103f6565b6040516100ed93929190611914565b6100e16101e536600461188b565b600091825260036020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61013a600081565b6101a061023336600461188b565b6107f9565b6001546101689073ffffffffffffffffffffffffffffffffffffffff1681565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806102eb57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b60006060610304878787878760006103f6565b9199919850909650505050505050565b60008281526003602052604090206001015461032f8161081e565b610339838361082b565b505050565b73ffffffffffffffffffffffffffffffffffffffff811633146103e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b6103f2828261091f565b5050565b6040805160018082528183019092526000918291606091602080830190803683370190505090508773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff160361045c57866000925092506107ed565b6000610467866109da565b60408051600180825281830190925291925060009190816020015b60608152602001906001900390816104825790505090506104a1611711565b6000805b60018110156106785760606104d68f8f8f8f8a87600181106104c9576104c961196b565b602002015163ffffffff16565b8584600181106104e8576104e861196b565b602002015290506000610524670de0b6b3a764000061051e8d8887600181106105135761051361196b565b6020020151906109f7565b90610a0a565b90506105318d60016119c9565b67ffffffffffffffff8111156105495761054961193c565b604051908082528060200260200182016040528015610572578160200160208202803683370190505b508684815181106105855761058561196b565b602002602001018190525060005b825181101561066257818382815181106105af576105af61196b565b60200260200101516105c191906119dc565b8785815181106105d3576105d361196b565b60200260200101518260016105e891906119c9565b815181106105f8576105f861196b565b602002602001018181525050848061064e5750600087858151811061061f5761061f61196b565b602002602001015182600161063491906119c9565b815181106106445761064461196b565b6020026020010151135b94508061065a81611a03565b915050610593565b505050808061067090611a03565b9150506104a5565b50806107595760005b60018110156107575760015b6106988c60016119c9565b811015610744578482815181106106b1576106b161196b565b602002602001015181815181106106ca576106ca61196b565b6020026020010151600003610732577fffff6f1bf04115e2c5b54376aa16b901ce32309909cb1f00000000000000000085838151811061070c5761070c61196b565b602002602001015182815181106107255761072561196b565b6020026020010181815250505b8061073c81611a03565b91505061068d565b508061074f81611a03565b915050610681565b505b6107638a84610a16565b9050809550506107e36040518061014001604052808f73ffffffffffffffffffffffffffffffffffffffff1681526020018e73ffffffffffffffffffffffffffffffffffffffff1681526020018d81526020018c81526020018b81526020018a81526020018781526020018581526020018481526020018681525061113b565b9097509550505050505b96509650969350505050565b6000828152600360205260409020600101546108148161081e565b610339838361091f565b6108288133611395565b50565b600082815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166103f257600082815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556108c13390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600082815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16156103f257600082815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6109e261172f565b5050604080516020810190915261144f815290565b6000610a038284611a3b565b9392505050565b6000610a038284611a52565b8051600090606090828167ffffffffffffffff811115610a3857610a3861193c565b604051908082528060200260200182016040528015610a6b57816020015b6060815260200190600190039081610a565790505b50905060008267ffffffffffffffff811115610a8957610a8961193c565b604051908082528060200260200182016040528015610abc57816020015b6060815260200190600190039081610aa75790505b50905060005b83811015610bb157610ad58860016119c9565b67ffffffffffffffff811115610aed57610aed61193c565b604051908082528060200260200182016040528015610b16578160200160208202803683370190505b50838281518110610b2957610b2961196b565b6020908102919091010152610b3f8860016119c9565b67ffffffffffffffff811115610b5757610b5761193c565b604051908082528060200260200182016040528015610b80578160200160208202803683370190505b50828281518110610b9357610b9361196b565b60200260200101819052508080610ba990611a03565b915050610ac2565b5060005b878111610ce75786600081518110610bcf57610bcf61196b565b60200260200101518181518110610be857610be861196b565b602002602001015183600081518110610c0357610c0361196b565b60200260200101518281518110610c1c57610c1c61196b565b602090810291909101015260015b84811015610c9b577fffff6f1bf04115e2c5b54376aa16b901ce32309909cb1f000000000000000000848281518110610c6557610c6561196b565b60200260200101518381518110610c7e57610c7e61196b565b602090810291909101015280610c9381611a03565b915050610c2a565b50600082600081518110610cb157610cb161196b565b60200260200101518281518110610cca57610cca61196b565b602090810291909101015280610cdf81611a03565b915050610bb5565b5060015b83811015610fa95760005b888111610f965783610d09600184611a8d565b81518110610d1957610d1961196b565b60200260200101518181518110610d3257610d3261196b565b6020026020010151848381518110610d4c57610d4c61196b565b60200260200101518281518110610d6557610d6561196b565b60200260200101818152505080838381518110610d8457610d8461196b565b60200260200101518281518110610d9d57610d9d61196b565b602090810291909101015260015b818111610f8357848381518110610dc457610dc461196b565b60200260200101518281518110610ddd57610ddd61196b565b6020026020010151898481518110610df757610df761196b565b60200260200101518281518110610e1057610e1061196b565b602002602001015186600186610e269190611a8d565b81518110610e3657610e3661196b565b60200260200101518385610e4a9190611a8d565b81518110610e5a57610e5a61196b565b6020026020010151610e6c9190611aa0565b1315610f7157888381518110610e8457610e8461196b565b60200260200101518181518110610e9d57610e9d61196b565b602002602001015185600185610eb39190611a8d565b81518110610ec357610ec361196b565b60200260200101518284610ed79190611a8d565b81518110610ee757610ee761196b565b6020026020010151610ef99190611aa0565b858481518110610f0b57610f0b61196b565b60200260200101518381518110610f2457610f2461196b565b6020908102919091010152610f398183611a8d565b848481518110610f4b57610f4b61196b565b60200260200101518381518110610f6457610f6461196b565b6020026020010181815250505b80610f7b81611a03565b915050610dab565b5080610f8e81611a03565b915050610cf6565b5080610fa181611a03565b915050610ceb565b506040805160018082528183019092529060208083019080368337019050509350866000610fd8600186611a8d565b90505b811561108557828181518110610ff357610ff361196b565b6020026020010151828151811061100c5761100c61196b565b60200260200101518261101f9190611a8d565b8682815181106110315761103161196b565b60200260200101818152505082818151811061104f5761104f61196b565b602002602001015182815181106110685761106861196b565b60200260200101519150808061107d90611ac8565b915050610fdb565b507fffff6f1bf04115e2c5b54376aa16b901ce32309909cb1f000000000000000000836110b3600187611a8d565b815181106110c3576110c361196b565b602002602001015189815181106110dc576110dc61196b565b60200260200101511461112b57826110f5600186611a8d565b815181106111055761110561196b565b6020026020010151888151811061111e5761111e61196b565b602002602001015161112e565b60005b9550505050509250929050565b6040805160208101909152600181526000908190815b600181101561138e5760008560c0015182815181106111725761117261196b565b6020026020010151111561137c5784606001518560c00151828151811061119b5761119b61196b565b602002602001015114806111c257508181600181106111bc576111bc61196b565b60200201515b806111d857506080850151658000000000001615155b156112d45761120385610100015182600181106111f7576111f761196b565b602002015184906114a3565b925060008560e00151828151811061121d5761121d61196b565b60200260200101518660c00151838151811061123b5761123b61196b565b6020026020010151815181106112535761125361196b565b602002602001015190506112cc61128b670de0b6b3a764000061051e8960a001518a610100015187600181106105135761051361196b565b7fffff6f1bf04115e2c5b54376aa16b901ce32309909cb1f00000000000000000083146112b857826112bb565b60005b6112c59190611aa0565b86906114a3565b94505061137c565b600080611339876000015188602001516113208a6060015161051e8c60c0015189815181106113055761130561196b565b60200260200101518d604001516109f790919063ffffffff16565b60018b610120015188600181106104c9576104c961196b565b909250905061134885826114a3565b9450611377826000815181106113605761136061196b565b6020026020010151876114a390919063ffffffff16565b955050505b8061138681611a03565b915050611151565b5050915091565b600082815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166103f2576113d5816114af565b6113e08360206114ce565b6040516020016113f1929190611b21565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526103df91600401611ba2565b606060008267ffffffffffffffff81111561146c5761146c61193c565b604051908082528060200260200182016040528015611495578160200160208202803683370190505b509660009650945050505050565b6000610a0382846119c9565b60606102eb73ffffffffffffffffffffffffffffffffffffffff831660145b606060006114dd836002611a3b565b6114e89060026119c9565b67ffffffffffffffff8111156115005761150061193c565b6040519080825280601f01601f19166020018201604052801561152a576020820181803683370190505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106115615761156161196b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106115c4576115c461196b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000611600846002611a3b565b61160b9060016119c9565b90505b60018111156116a8577f303132333435363738396162636465660000000000000000000000000000000085600f166010811061164c5761164c61196b565b1a60f81b8282815181106116625761166261196b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c936116a181611ac8565b905061160e565b508315610a03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016103df565b60405180602001604052806001906020820280368337509192915050565b60405180602001604052806001905b61175781526020019060019003908161173e5790505090565b61175f611bf3565b565b60006020828403121561177357600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610a0357600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461082857600080fd5b600080600080600060a086880312156117dd57600080fd5b85356117e8816117a3565b945060208601356117f8816117a3565b94979496505050506040830135926060810135926080909101359150565b600081518084526020808501945080840160005b838110156118465781518752958201959082019060010161182a565b509495945050505050565b82815260406020820152600061186a6040830184611816565b949350505050565b60006020828403121561188457600080fd5b5035919050565b6000806040838503121561189e57600080fd5b8235915060208301356118b0816117a3565b809150509250929050565b60008060008060008060c087890312156118d457600080fd5b86356118df816117a3565b955060208701356118ef816117a3565b95989597505050506040840135936060810135936080820135935060a0909101359150565b8381528260208201526060604082015260006119336060830184611816565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156102eb576102eb61199a565b81810360008312801583831316838312821617156119fc576119fc61199a565b5092915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a3457611a3461199a565b5060010190565b80820281158282048414176102eb576102eb61199a565b600082611a88577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b818103818111156102eb576102eb61199a565b8082018281126000831280158216821582161715611ac057611ac061199a565b505092915050565b600081611ad757611ad761199a565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60005b83811015611b18578181015183820152602001611b00565b50506000910152565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611b59816017850160208801611afd565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351611b96816028840160208801611afd565b01602801949350505050565b6020815260008251806020840152611bc1816040850160208701611afd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052605160045260246000fdfea2646970667358221220030730c7568dad2d9b8d7d1e768bceb00256af02f4a9b5747a54c7a2116af3ec64736f6c63430008140033",
}

// DexSpanViewABI is the input ABI used to generate the binding from.
// Deprecated: Use DexSpanViewMetaData.ABI instead.
var DexSpanViewABI = DexSpanViewMetaData.ABI

// DexSpanViewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DexSpanViewMetaData.Bin instead.
var DexSpanViewBin = DexSpanViewMetaData.Bin

// DeployDexSpanView deploys a new Ethereum contract, binding an instance of DexSpanView to it.
func DeployDexSpanView(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DexSpanView, error) {
	parsed, err := DexSpanViewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DexSpanViewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DexSpanView{DexSpanViewCaller: DexSpanViewCaller{contract: contract}, DexSpanViewTransactor: DexSpanViewTransactor{contract: contract}, DexSpanViewFilterer: DexSpanViewFilterer{contract: contract}}, nil
}

// DexSpanView is an auto generated Go binding around an Ethereum contract.
type DexSpanView struct {
	DexSpanViewCaller     // Read-only binding to the contract
	DexSpanViewTransactor // Write-only binding to the contract
	DexSpanViewFilterer   // Log filterer for contract events
}

// DexSpanViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexSpanViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSpanViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexSpanViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSpanViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexSpanViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexSpanViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexSpanViewSession struct {
	Contract     *DexSpanView      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexSpanViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexSpanViewCallerSession struct {
	Contract *DexSpanViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DexSpanViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexSpanViewTransactorSession struct {
	Contract     *DexSpanViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DexSpanViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexSpanViewRaw struct {
	Contract *DexSpanView // Generic contract binding to access the raw methods on
}

// DexSpanViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexSpanViewCallerRaw struct {
	Contract *DexSpanViewCaller // Generic read-only contract binding to access the raw methods on
}

// DexSpanViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexSpanViewTransactorRaw struct {
	Contract *DexSpanViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexSpanView creates a new instance of DexSpanView, bound to a specific deployed contract.
func NewDexSpanView(address common.Address, backend bind.ContractBackend) (*DexSpanView, error) {
	contract, err := bindDexSpanView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DexSpanView{DexSpanViewCaller: DexSpanViewCaller{contract: contract}, DexSpanViewTransactor: DexSpanViewTransactor{contract: contract}, DexSpanViewFilterer: DexSpanViewFilterer{contract: contract}}, nil
}

// NewDexSpanViewCaller creates a new read-only instance of DexSpanView, bound to a specific deployed contract.
func NewDexSpanViewCaller(address common.Address, caller bind.ContractCaller) (*DexSpanViewCaller, error) {
	contract, err := bindDexSpanView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexSpanViewCaller{contract: contract}, nil
}

// NewDexSpanViewTransactor creates a new write-only instance of DexSpanView, bound to a specific deployed contract.
func NewDexSpanViewTransactor(address common.Address, transactor bind.ContractTransactor) (*DexSpanViewTransactor, error) {
	contract, err := bindDexSpanView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexSpanViewTransactor{contract: contract}, nil
}

// NewDexSpanViewFilterer creates a new log filterer instance of DexSpanView, bound to a specific deployed contract.
func NewDexSpanViewFilterer(address common.Address, filterer bind.ContractFilterer) (*DexSpanViewFilterer, error) {
	contract, err := bindDexSpanView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexSpanViewFilterer{contract: contract}, nil
}

// bindDexSpanView binds a generic wrapper to an already deployed contract.
func bindDexSpanView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DexSpanViewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexSpanView *DexSpanViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DexSpanView.Contract.DexSpanViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexSpanView *DexSpanViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexSpanView.Contract.DexSpanViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexSpanView *DexSpanViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexSpanView.Contract.DexSpanViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexSpanView *DexSpanViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DexSpanView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexSpanView *DexSpanViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexSpanView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexSpanView *DexSpanViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexSpanView.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DexSpanView *DexSpanViewCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DexSpanView *DexSpanViewSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DexSpanView.Contract.DEFAULTADMINROLE(&_DexSpanView.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DexSpanView *DexSpanViewCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DexSpanView.Contract.DEFAULTADMINROLE(&_DexSpanView.CallOpts)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_DexSpanView *DexSpanViewCaller) GetExpectedReturn(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "getExpectedReturn", fromToken, destToken, amount, parts, flags)

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
func (_DexSpanView *DexSpanViewSession) GetExpectedReturn(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _DexSpanView.Contract.GetExpectedReturn(&_DexSpanView.CallOpts, fromToken, destToken, amount, parts, flags)
}

// GetExpectedReturn is a free data retrieval call binding the contract method 0x085e2c5b.
//
// Solidity: function getExpectedReturn(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags) view returns(uint256 returnAmount, uint256[] distribution)
func (_DexSpanView *DexSpanViewCallerSession) GetExpectedReturn(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int) (struct {
	ReturnAmount *big.Int
	Distribution []*big.Int
}, error) {
	return _DexSpanView.Contract.GetExpectedReturn(&_DexSpanView.CallOpts, fromToken, destToken, amount, parts, flags)
}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_DexSpanView *DexSpanViewCaller) GetExpectedReturnWithGas(opts *bind.CallOpts, fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "getExpectedReturnWithGas", fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)

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
func (_DexSpanView *DexSpanViewSession) GetExpectedReturnWithGas(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _DexSpanView.Contract.GetExpectedReturnWithGas(&_DexSpanView.CallOpts, fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)
}

// GetExpectedReturnWithGas is a free data retrieval call binding the contract method 0x8373f265.
//
// Solidity: function getExpectedReturnWithGas(address fromToken, address destToken, uint256 amount, uint256 parts, uint256 flags, uint256 destTokenEthPriceTimesGasPrice) view returns(uint256 returnAmount, uint256 estimateGasAmount, uint256[] distribution)
func (_DexSpanView *DexSpanViewCallerSession) GetExpectedReturnWithGas(fromToken common.Address, destToken common.Address, amount *big.Int, parts *big.Int, flags *big.Int, destTokenEthPriceTimesGasPrice *big.Int) (struct {
	ReturnAmount      *big.Int
	EstimateGasAmount *big.Int
	Distribution      []*big.Int
}, error) {
	return _DexSpanView.Contract.GetExpectedReturnWithGas(&_DexSpanView.CallOpts, fromToken, destToken, amount, parts, flags, destTokenEthPriceTimesGasPrice)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DexSpanView *DexSpanViewCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DexSpanView *DexSpanViewSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DexSpanView.Contract.GetRoleAdmin(&_DexSpanView.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DexSpanView *DexSpanViewCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DexSpanView.Contract.GetRoleAdmin(&_DexSpanView.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DexSpanView *DexSpanViewCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DexSpanView *DexSpanViewSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DexSpanView.Contract.HasRole(&_DexSpanView.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DexSpanView *DexSpanViewCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DexSpanView.Contract.HasRole(&_DexSpanView.CallOpts, role, account)
}

// NativeAddress is a free data retrieval call binding the contract method 0x2e476337.
//
// Solidity: function nativeAddress() view returns(address)
func (_DexSpanView *DexSpanViewCaller) NativeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "nativeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeAddress is a free data retrieval call binding the contract method 0x2e476337.
//
// Solidity: function nativeAddress() view returns(address)
func (_DexSpanView *DexSpanViewSession) NativeAddress() (common.Address, error) {
	return _DexSpanView.Contract.NativeAddress(&_DexSpanView.CallOpts)
}

// NativeAddress is a free data retrieval call binding the contract method 0x2e476337.
//
// Solidity: function nativeAddress() view returns(address)
func (_DexSpanView *DexSpanViewCallerSession) NativeAddress() (common.Address, error) {
	return _DexSpanView.Contract.NativeAddress(&_DexSpanView.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DexSpanView *DexSpanViewCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DexSpanView *DexSpanViewSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DexSpanView.Contract.SupportsInterface(&_DexSpanView.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DexSpanView *DexSpanViewCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DexSpanView.Contract.SupportsInterface(&_DexSpanView.CallOpts, interfaceId)
}

// WnativeAddress is a free data retrieval call binding the contract method 0xe2a4ac2d.
//
// Solidity: function wnativeAddress() view returns(address)
func (_DexSpanView *DexSpanViewCaller) WnativeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexSpanView.contract.Call(opts, &out, "wnativeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WnativeAddress is a free data retrieval call binding the contract method 0xe2a4ac2d.
//
// Solidity: function wnativeAddress() view returns(address)
func (_DexSpanView *DexSpanViewSession) WnativeAddress() (common.Address, error) {
	return _DexSpanView.Contract.WnativeAddress(&_DexSpanView.CallOpts)
}

// WnativeAddress is a free data retrieval call binding the contract method 0xe2a4ac2d.
//
// Solidity: function wnativeAddress() view returns(address)
func (_DexSpanView *DexSpanViewCallerSession) WnativeAddress() (common.Address, error) {
	return _DexSpanView.Contract.WnativeAddress(&_DexSpanView.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.Contract.GrantRole(&_DexSpanView.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.Contract.GrantRole(&_DexSpanView.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.Contract.RenounceRole(&_DexSpanView.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.Contract.RenounceRole(&_DexSpanView.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.Contract.RevokeRole(&_DexSpanView.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DexSpanView *DexSpanViewTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DexSpanView.Contract.RevokeRole(&_DexSpanView.TransactOpts, role, account)
}

// DexSpanViewRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DexSpanView contract.
type DexSpanViewRoleAdminChangedIterator struct {
	Event *DexSpanViewRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DexSpanViewRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexSpanViewRoleAdminChanged)
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
		it.Event = new(DexSpanViewRoleAdminChanged)
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
func (it *DexSpanViewRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexSpanViewRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexSpanViewRoleAdminChanged represents a RoleAdminChanged event raised by the DexSpanView contract.
type DexSpanViewRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DexSpanView *DexSpanViewFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DexSpanViewRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DexSpanView.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DexSpanViewRoleAdminChangedIterator{contract: _DexSpanView.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DexSpanView *DexSpanViewFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DexSpanViewRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _DexSpanView.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexSpanViewRoleAdminChanged)
				if err := _DexSpanView.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DexSpanView *DexSpanViewFilterer) ParseRoleAdminChanged(log types.Log) (*DexSpanViewRoleAdminChanged, error) {
	event := new(DexSpanViewRoleAdminChanged)
	if err := _DexSpanView.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexSpanViewRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DexSpanView contract.
type DexSpanViewRoleGrantedIterator struct {
	Event *DexSpanViewRoleGranted // Event containing the contract specifics and raw log

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
func (it *DexSpanViewRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexSpanViewRoleGranted)
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
		it.Event = new(DexSpanViewRoleGranted)
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
func (it *DexSpanViewRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexSpanViewRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexSpanViewRoleGranted represents a RoleGranted event raised by the DexSpanView contract.
type DexSpanViewRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DexSpanView *DexSpanViewFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DexSpanViewRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DexSpanView.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DexSpanViewRoleGrantedIterator{contract: _DexSpanView.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DexSpanView *DexSpanViewFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DexSpanViewRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DexSpanView.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexSpanViewRoleGranted)
				if err := _DexSpanView.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DexSpanView *DexSpanViewFilterer) ParseRoleGranted(log types.Log) (*DexSpanViewRoleGranted, error) {
	event := new(DexSpanViewRoleGranted)
	if err := _DexSpanView.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexSpanViewRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DexSpanView contract.
type DexSpanViewRoleRevokedIterator struct {
	Event *DexSpanViewRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DexSpanViewRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexSpanViewRoleRevoked)
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
		it.Event = new(DexSpanViewRoleRevoked)
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
func (it *DexSpanViewRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexSpanViewRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexSpanViewRoleRevoked represents a RoleRevoked event raised by the DexSpanView contract.
type DexSpanViewRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DexSpanView *DexSpanViewFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DexSpanViewRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DexSpanView.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DexSpanViewRoleRevokedIterator{contract: _DexSpanView.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DexSpanView *DexSpanViewFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DexSpanViewRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _DexSpanView.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexSpanViewRoleRevoked)
				if err := _DexSpanView.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DexSpanView *DexSpanViewFilterer) ParseRoleRevoked(log types.Log) (*DexSpanViewRoleRevoked, error) {
	event := new(DexSpanViewRoleRevoked)
	if err := _DexSpanView.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
