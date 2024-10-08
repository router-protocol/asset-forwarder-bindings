// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CCTPBridge

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

// ICCTPBridgeDestDetails is an auto generated low-level Go binding around an user-defined struct.
type ICCTPBridgeDestDetails struct {
	DomainId       uint32
	BaseFee        *big.Int
	AtaCreationFee *big.Int
	IsSet          bool
}

// IDexSpanSwapPayload is an auto generated low-level Go binding around an user-defined struct.
type IDexSpanSwapPayload struct {
	Flags       []*big.Int
	MinToAmount *big.Int
	Tokens      []common.Address
	DataTx      [][]byte
}

// CCTPBridgeMetaData contains all meta data concerning the CCTPBridge contract.
var CCTPBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dexspan\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CCTPNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUpdateType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnequalLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongReturnToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESOURCE_SETTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"destDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ataCreationFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dexspan\",\"outputs\":[{\"internalType\":\"contractIDexSpan\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"iDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_destChainIdBytes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ataCreationFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"internalType\":\"structICCTPBridge.DestDetails[]\",\"name\":\"_destDetails\",\"type\":\"tuple[]\"}],\"name\":\"setDestDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"flags\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"dataTx\",\"type\":\"bytes[]\"}],\"internalType\":\"structIDexSpan.SwapPayload\",\"name\":\"swapPayload\",\"type\":\"tuple\"}],\"name\":\"swapAndIDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"utype\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002b8838038062002b888339810160408190526200003491620001aa565b60018054600380546001600160a01b038087166001600160a01b031992831617909255600280548884169216919091179055831662010000026001600160b01b031990911617811790556200008b600033620000ec565b620000b77f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07833620000ec565b620000e37f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c33620000ec565b505050620001f4565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1662000189576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620001483390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b80516001600160a01b0381168114620001a557600080fd5b919050565b600080600060608486031215620001c057600080fd5b620001cb846200018d565b9250620001db602085016200018d565b9150620001eb604085016200018d565b90509250925092565b61298480620002046000396000f3fe60806040526004361061016a5760003560e01c80637a4e4ecf116100cb578063981a8a021161007f578063d547741f11610059578063d547741f1461047e578063d9dc86941461049e578063db618e2a146104d257600080fd5b8063981a8a02146103e15780639d9c44c914610456578063a217fddf1461046957600080fd5b806384756a06116100b057806384756a061461035057806391d1485414610370578063944f778a146103c157600080fd5b80637a4e4ecf1461031b5780638456cb591461033b57600080fd5b80633f4ba83a116101225780635c975abb116101075780635c975abb146102b857806367d70e13146102d55780636cec044b146102e857600080fd5b80633f4ba83a14610276578063461178301461028b57600080fd5b80632f2ff15d116101535780632f2ff15d146101e257806336568abe146102045780633e413bee1461022457600080fd5b806301ffc9a71461016f578063248a9ca3146101a4575b600080fd5b34801561017b57600080fd5b5061018f61018a366004611e0f565b610506565b60405190151581526020015b60405180910390f35b3480156101b057600080fd5b506101d46101bf366004611e51565b60009081526020819052604090206001015490565b60405190815260200161019b565b3480156101ee57600080fd5b506102026101fd366004611e93565b61059f565b005b34801561021057600080fd5b5061020261021f366004611e93565b6105c9565b34801561023057600080fd5b506002546102519073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b34801561028257600080fd5b50610202610681565b34801561029757600080fd5b506003546102519073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102c457600080fd5b50600154610100900460ff1661018f565b6102026102e3366004612005565b6106be565b3480156102f457600080fd5b506001546102519062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b34801561032757600080fd5b50610202610336366004612078565b610896565b34801561034757600080fd5b50610202610964565b34801561035c57600080fd5b5061020261036b3660046120a2565b61099e565b34801561037c57600080fd5b5061018f61038b366004611e93565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156103cd57600080fd5b506102026103dc3660046121bd565b610af4565b3480156103ed57600080fd5b5061042d6103fc366004611e51565b600460205260009081526040902080546001820154600283015460039093015463ffffffff90921692909160ff1684565b6040805163ffffffff90951685526020850193909352918301521515606082015260800161019b565b610202610464366004612358565b610c44565b34801561047557600080fd5b506101d4600081565b34801561048a57600080fd5b50610202610499366004611e93565b61109e565b3480156104aa57600080fd5b506101d47f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c81565b3480156104de57600080fd5b506101d47f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07881565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061059957507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6000828152602081905260409020600101546105ba816110c3565b6105c483836110cd565b505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610673576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b61067d82826111bd565b5050565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c6106ab816110c3565b6106b3611274565b6106bb6112e7565b50565b6106c6611364565b6106f3600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6106fb6113d0565b6000848152600460205260409020600301548490879060ff1661074a576040517f25928b3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60ff8116158061075d57508060ff166001145b156107ce5760008281526004602052604090206001015415801590610792575060008281526004602052604090206001015434105b156107c9576040517f58d620b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61082a565b600082815260046020526040902060028101546001909101546107f191906124fe565b34101561082a576040517f58d620b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025461084f9073ffffffffffffffffffffffffffffffffffffffff16333088611442565b61085d88888887898861151e565b505061088e600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b505050505050565b60006108a1816110c3565b6108a9611364565b6108d6600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b604051600090339084908381818185875af1925050503d8060008114610918576040519150601f19603f3d011682016040523d82523d6000602084013e61091d565b606091505b509091505060018115151461093457610934612511565b506105c4600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c61098e816110c3565b6109966113d0565b6106bb61167b565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb0786109c8816110c3565b8260ff16600103610a1f576001805473ffffffffffffffffffffffffffffffffffffffff841662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909116179055505050565b8260ff16600203610a71576003805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b8260ff16600303610ac3576002805473ffffffffffffffffffffffffffffffffffffffff84167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116179055505050565b6040517e64280000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb078610b1e816110c3565b8151835114610b59576040517f1c091df100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8251811015610c3e57828181518110610b7757610b77612580565b602002602001015160046000868481518110610b9557610b95612580565b6020908102919091018101518252818101929092526040908101600020835181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff9091161781559183015160018301558201516002820155606090910151600390910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001691151591909117905580610c36816125af565b915050610b5c565b50505050565b610c4c611364565b610c79600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b610c816113d0565b6000858152600460205260409020600301548590889060ff16610cd0576040517f25928b3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60ff81161580610ce357508060ff166001145b15610d545760008281526004602052604090206001015415801590610d18575060008281526004602052604090206001015434105b15610d4f576040517f58d620b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610db0565b60008281526004602052604090206002810154600190910154610d7791906124fe565b341015610db0576040517f58d620b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040840151805173ffffffffffffffffffffffffffffffffffffffff90921691610ddf906001906125e7565b81518110610def57610def612580565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614610e44576040517ff099a50400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610e978460400151600081518110610e6057610e60612580565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15610f2757600088815260046020526040812060010154610eb890346125e7565b905060028b60ff1610610ee457600089815260046020526040902060020154610ee190826125e7565b90505b87811015610f1e576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b87915050610f94565b610f9433600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16898760400151600081518110610f6657610f66612580565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16611442909392919063ffffffff16565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e738aa8d8387604001518b89602001518a600001518b606001516001306040518963ffffffff1660e01b815260040161100f97969594939291906126d4565b60206040518083038185885af115801561102d573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906110529190612792565b90506110628b8b8b8a858b61151e565b50505050611095600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b50505050505050565b6000828152602081905260409020600101546110b9816110c3565b6105c483836111bd565b6106bb81336116d7565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661067d5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561115f3390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161561067d5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600154610100900460ff166112e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161066a565b565b6112ef611274565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60015460ff166112e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161066a565b600154610100900460ff16156112e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161066a565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610c3e9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261178f565b6003546002546115489173ffffffffffffffffffffffffffffffffffffffff91821691168461189e565b60035460008581526004602081905260408083205460025491517f6fd3504e00000000000000000000000000000000000000000000000000000000815292830187905263ffffffff1660248301526044820187905273ffffffffffffffffffffffffffffffffffffffff908116606483015291929190911690636fd3504e906084016020604051808303816000875af11580156115e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061160d91906127ab565b6002546040519192507fa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a5269161166a918a918a9188918b91889173ffffffffffffffffffffffffffffffffffffffff909116908c9033908c906127d5565b60405180910390a150505050505050565b6116836113d0565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861133a3390565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661067d5761171581611997565b6117208360206119b6565b604051602001611731929190612851565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261066a916004016128d2565b60006117f1826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611c009092919063ffffffff16565b905080516000148061181257508080602001905181019061181291906128e5565b6105c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161066a565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83811660248301526000919085169063dd62ed3e90604401602060405180830381865afa158015611914573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119389190612792565b9050610c3e847f095ea7b3000000000000000000000000000000000000000000000000000000008561196a86866124fe565b60405173ffffffffffffffffffffffffffffffffffffffff9092166024830152604482015260640161149c565b606061059973ffffffffffffffffffffffffffffffffffffffff831660145b606060006119c5836002612902565b6119d09060026124fe565b67ffffffffffffffff8111156119e8576119e8611ed0565b6040519080825280601f01601f191660200182016040528015611a12576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110611a4957611a49612580565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110611aac57611aac612580565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000611ae8846002612902565b611af39060016124fe565b90505b6001811115611b90577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110611b3457611b34612580565b1a60f81b828281518110611b4a57611b4a612580565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93611b8981612919565b9050611af6565b508315611bf9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161066a565b9392505050565b6060611c0f8484600085611c17565b949350505050565b606082471015611ca9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161066a565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611cd29190612564565b60006040518083038185875af1925050503d8060008114611d0f576040519150601f19603f3d011682016040523d82523d6000602084013e611d14565b606091505b5091509150611d2587838387611d30565b979650505050505050565b60608315611dc6578251600003611dbf5773ffffffffffffffffffffffffffffffffffffffff85163b611dbf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161066a565b5081611c0f565b611c0f8383815115611ddb5781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066a91906128d2565b600060208284031215611e2157600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114611bf957600080fd5b600060208284031215611e6357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611e8e57600080fd5b919050565b60008060408385031215611ea657600080fd5b82359150611eb660208401611e6a565b90509250929050565b803560ff81168114611e8e57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715611f2257611f22611ed0565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611f6f57611f6f611ed0565b604052919050565b600082601f830112611f8857600080fd5b813567ffffffffffffffff811115611fa257611fa2611ed0565b611fd360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611f28565b818152846020838601011115611fe857600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060c0878903121561201e57600080fd5b61202787611ebf565b95506020870135945060408701359350606087013592506080870135915060a087013567ffffffffffffffff81111561205f57600080fd5b61206b89828a01611f77565b9150509295509295509295565b6000806040838503121561208b57600080fd5b61209483611e6a565b946020939093013593505050565b600080604083850312156120b557600080fd5b6120be83611ebf565b9150611eb660208401611e6a565b600067ffffffffffffffff8211156120e6576120e6611ed0565b5060051b60200190565b80151581146106bb57600080fd5b600082601f83011261210f57600080fd5b8135602061212461211f836120cc565b611f28565b82815260079290921b8401810191818101908684111561214357600080fd5b8286015b848110156121b257608081890312156121605760008081fd5b612168611eff565b813563ffffffff8116811461217d5760008081fd5b81528185013585820152604080830135908201526060808301356121a0816120f0565b90820152835291830191608001612147565b509695505050505050565b600080604083850312156121d057600080fd5b823567ffffffffffffffff808211156121e857600080fd5b818501915085601f8301126121fc57600080fd5b8135602061220c61211f836120cc565b82815260059290921b8401810191818101908984111561222b57600080fd5b948201945b8386101561224957853582529482019490820190612230565b9650508601359250508082111561225f57600080fd5b5061226c858286016120fe565b9150509250929050565b600082601f83011261228757600080fd5b8135602061229761211f836120cc565b82815260059290921b840181019181810190868411156122b657600080fd5b8286015b848110156121b2576122cb81611e6a565b83529183019183016122ba565b600082601f8301126122e957600080fd5b813560206122f961211f836120cc565b82815260059290921b8401810191818101908684111561231857600080fd5b8286015b848110156121b257803567ffffffffffffffff81111561233c5760008081fd5b61234a8986838b0101611f77565b84525091830191830161231c565b600080600080600080600060e0888a03121561237357600080fd5b61237c88611ebf565b96506020880135955060408801359450606088013593506080880135925060a088013567ffffffffffffffff808211156123b557600080fd5b6123c18b838c01611f77565b935060c08a01359150808211156123d757600080fd5b908901906080828c0312156123eb57600080fd5b6123f3611eff565b818335111561240157600080fd5b823583018c601f82011261241457600080fd5b803561242261211f826120cc565b8082825260208201915060208360051b85010192508f83111561244457600080fd5b6020840193505b8284101561246657833582526020938401939091019061244b565b84525050506020838101359082015260408301358281111561248757600080fd5b6124938d828601612276565b6040830152506060830135828111156124ab57600080fd5b6124b78d8286016122d8565b60608301525080935050505092959891949750929550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610599576105996124cf565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60005b8381101561255b578181015183820152602001612543565b50506000910152565b60008251612576818460208701612540565b9190910192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036125e0576125e06124cf565b5060010190565b81810381811115610599576105996124cf565b600081518084526020808501945080840160005b8381101561262a5781518752958201959082019060010161260e565b509495945050505050565b6000815180845261264d816020860160208601612540565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600081518084526020808501808196508360051b8101915082860160005b858110156126c75782840389526126b5848351612635565b9885019893509084019060010161269d565b5091979650505050505050565b60e08082528851908201819052600090602090610100840190828c01845b8281101561272457815173ffffffffffffffffffffffffffffffffffffffff16845292840192908401906001016126f2565b5050508982850152886040850152838103606085015261274481896125fa565b9150508281036080840152612759818761267f565b91505061276a60a083018515159052565b73ffffffffffffffffffffffffffffffffffffffff831660c083015298975050505050505050565b6000602082840312156127a457600080fd5b5051919050565b6000602082840312156127bd57600080fd5b815167ffffffffffffffff81168114611bf957600080fd5b600061012060ff8c1683528a602084015289604084015288606084015267ffffffffffffffff8816608084015273ffffffffffffffffffffffffffffffffffffffff80881660a08501528660c085015280861660e0850152508061010084015261284181840185612635565b9c9b505050505050505050505050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351612889816017850160208801612540565b7f206973206d697373696e6720726f6c652000000000000000000000000000000060179184019182015283516128c6816028840160208801612540565b01602801949350505050565b602081526000611bf96020830184612635565b6000602082840312156128f757600080fd5b8151611bf9816120f0565b8082028115828204841417610599576105996124cf565b600081612928576129286124cf565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea2646970667358221220e7f1416334a58f92346dbd0217af7ed0e3f54325f1835ae36c6350318f720e1064736f6c63430008140033",
}

// CCTPBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CCTPBridgeMetaData.ABI instead.
var CCTPBridgeABI = CCTPBridgeMetaData.ABI

// CCTPBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CCTPBridgeMetaData.Bin instead.
var CCTPBridgeBin = CCTPBridgeMetaData.Bin

// DeployCCTPBridge deploys a new Ethereum contract, binding an instance of CCTPBridge to it.
func DeployCCTPBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _usdcAddress common.Address, _tokenMessenger common.Address, _dexspan common.Address) (common.Address, *types.Transaction, *CCTPBridge, error) {
	parsed, err := CCTPBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CCTPBridgeBin), backend, _usdcAddress, _tokenMessenger, _dexspan)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CCTPBridge{CCTPBridgeCaller: CCTPBridgeCaller{contract: contract}, CCTPBridgeTransactor: CCTPBridgeTransactor{contract: contract}, CCTPBridgeFilterer: CCTPBridgeFilterer{contract: contract}}, nil
}

// CCTPBridge is an auto generated Go binding around an Ethereum contract.
type CCTPBridge struct {
	CCTPBridgeCaller     // Read-only binding to the contract
	CCTPBridgeTransactor // Write-only binding to the contract
	CCTPBridgeFilterer   // Log filterer for contract events
}

// CCTPBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CCTPBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCTPBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CCTPBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCTPBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CCTPBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CCTPBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CCTPBridgeSession struct {
	Contract     *CCTPBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CCTPBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CCTPBridgeCallerSession struct {
	Contract *CCTPBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CCTPBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CCTPBridgeTransactorSession struct {
	Contract     *CCTPBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CCTPBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CCTPBridgeRaw struct {
	Contract *CCTPBridge // Generic contract binding to access the raw methods on
}

// CCTPBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CCTPBridgeCallerRaw struct {
	Contract *CCTPBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CCTPBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CCTPBridgeTransactorRaw struct {
	Contract *CCTPBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCCTPBridge creates a new instance of CCTPBridge, bound to a specific deployed contract.
func NewCCTPBridge(address common.Address, backend bind.ContractBackend) (*CCTPBridge, error) {
	contract, err := bindCCTPBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CCTPBridge{CCTPBridgeCaller: CCTPBridgeCaller{contract: contract}, CCTPBridgeTransactor: CCTPBridgeTransactor{contract: contract}, CCTPBridgeFilterer: CCTPBridgeFilterer{contract: contract}}, nil
}

// NewCCTPBridgeCaller creates a new read-only instance of CCTPBridge, bound to a specific deployed contract.
func NewCCTPBridgeCaller(address common.Address, caller bind.ContractCaller) (*CCTPBridgeCaller, error) {
	contract, err := bindCCTPBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeCaller{contract: contract}, nil
}

// NewCCTPBridgeTransactor creates a new write-only instance of CCTPBridge, bound to a specific deployed contract.
func NewCCTPBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CCTPBridgeTransactor, error) {
	contract, err := bindCCTPBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeTransactor{contract: contract}, nil
}

// NewCCTPBridgeFilterer creates a new log filterer instance of CCTPBridge, bound to a specific deployed contract.
func NewCCTPBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CCTPBridgeFilterer, error) {
	contract, err := bindCCTPBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeFilterer{contract: contract}, nil
}

// bindCCTPBridge binds a generic wrapper to an already deployed contract.
func bindCCTPBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CCTPBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CCTPBridge *CCTPBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPBridge.Contract.CCTPBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CCTPBridge *CCTPBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPBridge.Contract.CCTPBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CCTPBridge *CCTPBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPBridge.Contract.CCTPBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CCTPBridge *CCTPBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CCTPBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CCTPBridge *CCTPBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CCTPBridge *CCTPBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CCTPBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CCTPBridge.Contract.DEFAULTADMINROLE(&_CCTPBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CCTPBridge.Contract.DEFAULTADMINROLE(&_CCTPBridge.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCaller) PAUSER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "PAUSER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeSession) PAUSER() ([32]byte, error) {
	return _CCTPBridge.Contract.PAUSER(&_CCTPBridge.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCallerSession) PAUSER() ([32]byte, error) {
	return _CCTPBridge.Contract.PAUSER(&_CCTPBridge.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCaller) RESOURCESETTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "RESOURCE_SETTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeSession) RESOURCESETTER() ([32]byte, error) {
	return _CCTPBridge.Contract.RESOURCESETTER(&_CCTPBridge.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCallerSession) RESOURCESETTER() ([32]byte, error) {
	return _CCTPBridge.Contract.RESOURCESETTER(&_CCTPBridge.CallOpts)
}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 baseFee, uint256 ataCreationFee, bool isSet)
func (_CCTPBridge *CCTPBridgeCaller) DestDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DomainId       uint32
	BaseFee        *big.Int
	AtaCreationFee *big.Int
	IsSet          bool
}, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "destDetails", arg0)

	outstruct := new(struct {
		DomainId       uint32
		BaseFee        *big.Int
		AtaCreationFee *big.Int
		IsSet          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DomainId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BaseFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AtaCreationFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsSet = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 baseFee, uint256 ataCreationFee, bool isSet)
func (_CCTPBridge *CCTPBridgeSession) DestDetails(arg0 [32]byte) (struct {
	DomainId       uint32
	BaseFee        *big.Int
	AtaCreationFee *big.Int
	IsSet          bool
}, error) {
	return _CCTPBridge.Contract.DestDetails(&_CCTPBridge.CallOpts, arg0)
}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 baseFee, uint256 ataCreationFee, bool isSet)
func (_CCTPBridge *CCTPBridgeCallerSession) DestDetails(arg0 [32]byte) (struct {
	DomainId       uint32
	BaseFee        *big.Int
	AtaCreationFee *big.Int
	IsSet          bool
}, error) {
	return _CCTPBridge.Contract.DestDetails(&_CCTPBridge.CallOpts, arg0)
}

// Dexspan is a free data retrieval call binding the contract method 0x6cec044b.
//
// Solidity: function dexspan() view returns(address)
func (_CCTPBridge *CCTPBridgeCaller) Dexspan(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "dexspan")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dexspan is a free data retrieval call binding the contract method 0x6cec044b.
//
// Solidity: function dexspan() view returns(address)
func (_CCTPBridge *CCTPBridgeSession) Dexspan() (common.Address, error) {
	return _CCTPBridge.Contract.Dexspan(&_CCTPBridge.CallOpts)
}

// Dexspan is a free data retrieval call binding the contract method 0x6cec044b.
//
// Solidity: function dexspan() view returns(address)
func (_CCTPBridge *CCTPBridgeCallerSession) Dexspan() (common.Address, error) {
	return _CCTPBridge.Contract.Dexspan(&_CCTPBridge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CCTPBridge *CCTPBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CCTPBridge.Contract.GetRoleAdmin(&_CCTPBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CCTPBridge *CCTPBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CCTPBridge.Contract.GetRoleAdmin(&_CCTPBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CCTPBridge *CCTPBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CCTPBridge *CCTPBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CCTPBridge.Contract.HasRole(&_CCTPBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CCTPBridge *CCTPBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CCTPBridge.Contract.HasRole(&_CCTPBridge.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CCTPBridge *CCTPBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CCTPBridge *CCTPBridgeSession) Paused() (bool, error) {
	return _CCTPBridge.Contract.Paused(&_CCTPBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CCTPBridge *CCTPBridgeCallerSession) Paused() (bool, error) {
	return _CCTPBridge.Contract.Paused(&_CCTPBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CCTPBridge *CCTPBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CCTPBridge *CCTPBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPBridge.Contract.SupportsInterface(&_CCTPBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CCTPBridge *CCTPBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CCTPBridge.Contract.SupportsInterface(&_CCTPBridge.CallOpts, interfaceId)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_CCTPBridge *CCTPBridgeCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_CCTPBridge *CCTPBridgeSession) TokenMessenger() (common.Address, error) {
	return _CCTPBridge.Contract.TokenMessenger(&_CCTPBridge.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_CCTPBridge *CCTPBridgeCallerSession) TokenMessenger() (common.Address, error) {
	return _CCTPBridge.Contract.TokenMessenger(&_CCTPBridge.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_CCTPBridge *CCTPBridgeCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CCTPBridge.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_CCTPBridge *CCTPBridgeSession) Usdc() (common.Address, error) {
	return _CCTPBridge.Contract.Usdc(&_CCTPBridge.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_CCTPBridge *CCTPBridgeCallerSession) Usdc() (common.Address, error) {
	return _CCTPBridge.Contract.Usdc(&_CCTPBridge.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.GrantRole(&_CCTPBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.GrantRole(&_CCTPBridge.TransactOpts, role, account)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x67d70e13.
//
// Solidity: function iDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data) payable returns()
func (_CCTPBridge *CCTPBridgeTransactor) IDepositUSDC(opts *bind.TransactOpts, txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "iDepositUSDC", txType, partnerId, destChainIdBytes, amount, recipient, data)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x67d70e13.
//
// Solidity: function iDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data) payable returns()
func (_CCTPBridge *CCTPBridgeSession) IDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _CCTPBridge.Contract.IDepositUSDC(&_CCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x67d70e13.
//
// Solidity: function iDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data) payable returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) IDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte) (*types.Transaction, error) {
	return _CCTPBridge.Contract.IDepositUSDC(&_CCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CCTPBridge *CCTPBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CCTPBridge *CCTPBridgeSession) Pause() (*types.Transaction, error) {
	return _CCTPBridge.Contract.Pause(&_CCTPBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _CCTPBridge.Contract.Pause(&_CCTPBridge.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.RenounceRole(&_CCTPBridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.RenounceRole(&_CCTPBridge.TransactOpts, role, account)
}

// Rescue is a paid mutator transaction binding the contract method 0x7a4e4ecf.
//
// Solidity: function rescue(address token, uint256 amount) returns()
func (_CCTPBridge *CCTPBridgeTransactor) Rescue(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "rescue", token, amount)
}

// Rescue is a paid mutator transaction binding the contract method 0x7a4e4ecf.
//
// Solidity: function rescue(address token, uint256 amount) returns()
func (_CCTPBridge *CCTPBridgeSession) Rescue(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCTPBridge.Contract.Rescue(&_CCTPBridge.TransactOpts, token, amount)
}

// Rescue is a paid mutator transaction binding the contract method 0x7a4e4ecf.
//
// Solidity: function rescue(address token, uint256 amount) returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) Rescue(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CCTPBridge.Contract.Rescue(&_CCTPBridge.TransactOpts, token, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.RevokeRole(&_CCTPBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.RevokeRole(&_CCTPBridge.TransactOpts, role, account)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x944f778a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,uint256,bool)[] _destDetails) returns()
func (_CCTPBridge *CCTPBridgeTransactor) SetDestDetails(opts *bind.TransactOpts, _destChainIdBytes [][32]byte, _destDetails []ICCTPBridgeDestDetails) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "setDestDetails", _destChainIdBytes, _destDetails)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x944f778a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,uint256,bool)[] _destDetails) returns()
func (_CCTPBridge *CCTPBridgeSession) SetDestDetails(_destChainIdBytes [][32]byte, _destDetails []ICCTPBridgeDestDetails) (*types.Transaction, error) {
	return _CCTPBridge.Contract.SetDestDetails(&_CCTPBridge.TransactOpts, _destChainIdBytes, _destDetails)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x944f778a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,uint256,bool)[] _destDetails) returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) SetDestDetails(_destChainIdBytes [][32]byte, _destDetails []ICCTPBridgeDestDetails) (*types.Transaction, error) {
	return _CCTPBridge.Contract.SetDestDetails(&_CCTPBridge.TransactOpts, _destChainIdBytes, _destDetails)
}

// SwapAndIDepositUSDC is a paid mutator transaction binding the contract method 0x9d9c44c9.
//
// Solidity: function swapAndIDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data, (uint256[],uint256,address[],bytes[]) swapPayload) payable returns()
func (_CCTPBridge *CCTPBridgeTransactor) SwapAndIDepositUSDC(opts *bind.TransactOpts, txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte, swapPayload IDexSpanSwapPayload) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "swapAndIDepositUSDC", txType, partnerId, destChainIdBytes, amount, recipient, data, swapPayload)
}

// SwapAndIDepositUSDC is a paid mutator transaction binding the contract method 0x9d9c44c9.
//
// Solidity: function swapAndIDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data, (uint256[],uint256,address[],bytes[]) swapPayload) payable returns()
func (_CCTPBridge *CCTPBridgeSession) SwapAndIDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte, swapPayload IDexSpanSwapPayload) (*types.Transaction, error) {
	return _CCTPBridge.Contract.SwapAndIDepositUSDC(&_CCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data, swapPayload)
}

// SwapAndIDepositUSDC is a paid mutator transaction binding the contract method 0x9d9c44c9.
//
// Solidity: function swapAndIDepositUSDC(uint8 txType, uint256 partnerId, bytes32 destChainIdBytes, uint256 amount, bytes32 recipient, bytes data, (uint256[],uint256,address[],bytes[]) swapPayload) payable returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) SwapAndIDepositUSDC(txType uint8, partnerId *big.Int, destChainIdBytes [32]byte, amount *big.Int, recipient [32]byte, data []byte, swapPayload IDexSpanSwapPayload) (*types.Transaction, error) {
	return _CCTPBridge.Contract.SwapAndIDepositUSDC(&_CCTPBridge.TransactOpts, txType, partnerId, destChainIdBytes, amount, recipient, data, swapPayload)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CCTPBridge *CCTPBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CCTPBridge *CCTPBridgeSession) Unpause() (*types.Transaction, error) {
	return _CCTPBridge.Contract.Unpause(&_CCTPBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _CCTPBridge.Contract.Unpause(&_CCTPBridge.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x84756a06.
//
// Solidity: function update(uint8 utype, address _address) returns()
func (_CCTPBridge *CCTPBridgeTransactor) Update(opts *bind.TransactOpts, utype uint8, _address common.Address) (*types.Transaction, error) {
	return _CCTPBridge.contract.Transact(opts, "update", utype, _address)
}

// Update is a paid mutator transaction binding the contract method 0x84756a06.
//
// Solidity: function update(uint8 utype, address _address) returns()
func (_CCTPBridge *CCTPBridgeSession) Update(utype uint8, _address common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.Update(&_CCTPBridge.TransactOpts, utype, _address)
}

// Update is a paid mutator transaction binding the contract method 0x84756a06.
//
// Solidity: function update(uint8 utype, address _address) returns()
func (_CCTPBridge *CCTPBridgeTransactorSession) Update(utype uint8, _address common.Address) (*types.Transaction, error) {
	return _CCTPBridge.Contract.Update(&_CCTPBridge.TransactOpts, utype, _address)
}

// CCTPBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CCTPBridge contract.
type CCTPBridgePausedIterator struct {
	Event *CCTPBridgePaused // Event containing the contract specifics and raw log

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
func (it *CCTPBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPBridgePaused)
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
		it.Event = new(CCTPBridgePaused)
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
func (it *CCTPBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCTPBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCTPBridgePaused represents a Paused event raised by the CCTPBridge contract.
type CCTPBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CCTPBridge *CCTPBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*CCTPBridgePausedIterator, error) {

	logs, sub, err := _CCTPBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CCTPBridgePausedIterator{contract: _CCTPBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CCTPBridge *CCTPBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CCTPBridgePaused) (event.Subscription, error) {

	logs, sub, err := _CCTPBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCTPBridgePaused)
				if err := _CCTPBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CCTPBridge *CCTPBridgeFilterer) ParsePaused(log types.Log) (*CCTPBridgePaused, error) {
	event := new(CCTPBridgePaused)
	if err := _CCTPBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CCTPBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CCTPBridge contract.
type CCTPBridgeRoleAdminChangedIterator struct {
	Event *CCTPBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CCTPBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPBridgeRoleAdminChanged)
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
		it.Event = new(CCTPBridgeRoleAdminChanged)
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
func (it *CCTPBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCTPBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCTPBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the CCTPBridge contract.
type CCTPBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CCTPBridge *CCTPBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CCTPBridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CCTPBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeRoleAdminChangedIterator{contract: _CCTPBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CCTPBridge *CCTPBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CCTPBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CCTPBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCTPBridgeRoleAdminChanged)
				if err := _CCTPBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CCTPBridge *CCTPBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*CCTPBridgeRoleAdminChanged, error) {
	event := new(CCTPBridgeRoleAdminChanged)
	if err := _CCTPBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CCTPBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CCTPBridge contract.
type CCTPBridgeRoleGrantedIterator struct {
	Event *CCTPBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *CCTPBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPBridgeRoleGranted)
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
		it.Event = new(CCTPBridgeRoleGranted)
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
func (it *CCTPBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCTPBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCTPBridgeRoleGranted represents a RoleGranted event raised by the CCTPBridge contract.
type CCTPBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CCTPBridge *CCTPBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CCTPBridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _CCTPBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeRoleGrantedIterator{contract: _CCTPBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CCTPBridge *CCTPBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CCTPBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CCTPBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCTPBridgeRoleGranted)
				if err := _CCTPBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CCTPBridge *CCTPBridgeFilterer) ParseRoleGranted(log types.Log) (*CCTPBridgeRoleGranted, error) {
	event := new(CCTPBridgeRoleGranted)
	if err := _CCTPBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CCTPBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CCTPBridge contract.
type CCTPBridgeRoleRevokedIterator struct {
	Event *CCTPBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CCTPBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPBridgeRoleRevoked)
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
		it.Event = new(CCTPBridgeRoleRevoked)
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
func (it *CCTPBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCTPBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCTPBridgeRoleRevoked represents a RoleRevoked event raised by the CCTPBridge contract.
type CCTPBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CCTPBridge *CCTPBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CCTPBridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _CCTPBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeRoleRevokedIterator{contract: _CCTPBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CCTPBridge *CCTPBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CCTPBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CCTPBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCTPBridgeRoleRevoked)
				if err := _CCTPBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CCTPBridge *CCTPBridgeFilterer) ParseRoleRevoked(log types.Log) (*CCTPBridgeRoleRevoked, error) {
	event := new(CCTPBridgeRoleRevoked)
	if err := _CCTPBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CCTPBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CCTPBridge contract.
type CCTPBridgeUnpausedIterator struct {
	Event *CCTPBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *CCTPBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPBridgeUnpaused)
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
		it.Event = new(CCTPBridgeUnpaused)
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
func (it *CCTPBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCTPBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCTPBridgeUnpaused represents a Unpaused event raised by the CCTPBridge contract.
type CCTPBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CCTPBridge *CCTPBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CCTPBridgeUnpausedIterator, error) {

	logs, sub, err := _CCTPBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeUnpausedIterator{contract: _CCTPBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CCTPBridge *CCTPBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CCTPBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _CCTPBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCTPBridgeUnpaused)
				if err := _CCTPBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CCTPBridge *CCTPBridgeFilterer) ParseUnpaused(log types.Log) (*CCTPBridgeUnpaused, error) {
	event := new(CCTPBridgeUnpaused)
	if err := _CCTPBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CCTPBridgeIUSDCDepositedIterator is returned from FilterIUSDCDeposited and is used to iterate over the raw logs and unpacked data for IUSDCDeposited events raised by the CCTPBridge contract.
type CCTPBridgeIUSDCDepositedIterator struct {
	Event *CCTPBridgeIUSDCDeposited // Event containing the contract specifics and raw log

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
func (it *CCTPBridgeIUSDCDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CCTPBridgeIUSDCDeposited)
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
		it.Event = new(CCTPBridgeIUSDCDeposited)
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
func (it *CCTPBridgeIUSDCDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CCTPBridgeIUSDCDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CCTPBridgeIUSDCDeposited represents a IUSDCDeposited event raised by the CCTPBridge contract.
type CCTPBridgeIUSDCDeposited struct {
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
func (_CCTPBridge *CCTPBridgeFilterer) FilterIUSDCDeposited(opts *bind.FilterOpts) (*CCTPBridgeIUSDCDepositedIterator, error) {

	logs, sub, err := _CCTPBridge.contract.FilterLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return &CCTPBridgeIUSDCDepositedIterator{contract: _CCTPBridge.contract, event: "iUSDCDeposited", logs: logs, sub: sub}, nil
}

// WatchIUSDCDeposited is a free log subscription operation binding the contract event 0xa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a526.
//
// Solidity: event iUSDCDeposited(uint8 txType, uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor, bytes data)
func (_CCTPBridge *CCTPBridgeFilterer) WatchIUSDCDeposited(opts *bind.WatchOpts, sink chan<- *CCTPBridgeIUSDCDeposited) (event.Subscription, error) {

	logs, sub, err := _CCTPBridge.contract.WatchLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CCTPBridgeIUSDCDeposited)
				if err := _CCTPBridge.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
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
func (_CCTPBridge *CCTPBridgeFilterer) ParseIUSDCDeposited(log types.Log) (*CCTPBridgeIUSDCDeposited, error) {
	event := new(CCTPBridgeIUSDCDeposited)
	if err := _CCTPBridge.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
