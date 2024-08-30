// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VoyagerDepositHandler

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

// VoyagerDepositHandlerMetaData contains all meta data concerning the VoyagerDepositHandler contract.
var VoyagerDepositHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"FundsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ISendEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESOURCE_SETTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"relayerFeePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"depositor\",\"type\":\"bytes\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"relayerFeePct\",\"type\":\"int64\"}],\"name\":\"iSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620021303803806200213083398101604081905262000034916200016a565b6001805461ffff1916811790556001600160a01b0381166080526200005b600033620000ba565b620000877f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07833620000ba565b620000b37f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c33620000ba565b506200019c565b620000c68282620000ca565b5050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16620000c6576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620001263390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b6000602082840312156200017d57600080fd5b81516001600160a01b03811681146200019557600080fd5b9392505050565b608051611f55620001db60003960008181610175015281816105c20152818161068b01528181610a7f0152818161118001526111d60152611f556000f3fe6080604052600436106101295760003560e01c80638456cb59116100a5578063d547741f11610074578063db618e2a11610059578063db618e2a14610383578063ddd224f1146103b7578063de35f5cb146103da57600080fd5b8063d547741f1461032f578063d9dc86941461034f57600080fd5b80638456cb59146102a157806391d14854146102b65780639418617c14610307578063a217fddf1461031a57600080fd5b80632f2ff15d116100fc5780633f4ba83a116100e15780633f4ba83a1461024f5780635c975abb14610264578063839006f21461028157600080fd5b80632f2ff15d1461020f57806336568abe1461022f57600080fd5b806301ffc9a71461012e57806317fcb39b14610163578063248a9ca3146101bc57806326d1172a146101fa575b600080fd5b34801561013a57600080fd5b5061014e6101493660046118c6565b6103f0565b60405190151581526020015b60405180910390f35b34801561016f57600080fd5b506101977f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b3480156101c857600080fd5b506101ec6101d7366004611908565b60009081526020819052604090206001015490565b60405190815260200161015a565b61020d61020836600461194a565b610489565b005b34801561021b57600080fd5b5061020d61022a3660046119fa565b6107c7565b34801561023b57600080fd5b5061020d61024a3660046119fa565b6107f1565b34801561025b57600080fd5b5061020d6108a4565b34801561027057600080fd5b50600154610100900460ff1661014e565b34801561028d57600080fd5b5061020d61029c366004611a26565b6108e1565b3480156102ad57600080fd5b5061020d6109f7565b3480156102c257600080fd5b5061014e6102d13660046119fa565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b61020d610315366004611b1b565b610a31565b34801561032657600080fd5b506101ec600081565b34801561033b57600080fd5b5061020d61034a3660046119fa565b610bfa565b34801561035b57600080fd5b506101ec7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c81565b34801561038f57600080fd5b506101ec7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07881565b3480156103c357600080fd5b506101ec6ec097ce7bc90715b34b9f100000000081565b3480156103e657600080fd5b506101ec60025481565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061048357507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610491610c1f565b6104be600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6104c6610c8d565b6706f05b59d3b200006104db8260070b610cff565b10610547576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e76616c69642072656c61796572206665650000000000000000000000000060448201526064015b60405180910390fd5b6ec097ce7bc90715b34b9f10000000008211156105c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f416d6f756e7420746f6f206c6172676500000000000000000000000000000000604482015260640161053e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614801561061b5750600034115b1561070f57813414610689576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f6d73672e76616c7565206d757374206d6174636820616d6f756e740000000000604482015260640161053e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b1580156106f157600080fd5b505af1158015610705573d6000803e3d6000fd5b5050505050610731565b61073173ffffffffffffffffffffffffffffffffffffffff8416333085610d16565b600280546000918261074283611be5565b9091555090507f99f98099e0db79801cb9821f6ccfd8ea7abdd2ee8ae98f397a66bb0b48192d3c8346898585898c8c3360405161078799989796959493929190611c1e565b60405180910390a1506107bf600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b505050505050565b6000828152602081905260409020600101546107e281610df8565b6107ec8383610e02565b505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610896576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c660000000000000000000000000000000000606482015260840161053e565b6108a08282610ef2565b5050565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c6108ce81610df8565b6108d6610fa9565b6108de61101a565b50565b6108e9610c1f565b610916600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526109c890339073ffffffffffffffffffffffffffffffffffffffff8416906370a0823190602401602060405180830381865afa158015610986573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109aa9190611cbc565b73ffffffffffffffffffffffffffffffffffffffff84169190611097565b6108de600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c610a2181610df8565b610a29610c8d565b6108de6110ed565b610a39610c1f565b610a66600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b610a6e610c8d565b6000610a7b836014015190565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610b0257610af373ffffffffffffffffffffffffffffffffffffffff851633308b610d16565b610afd8189611149565b610b24565b610b2473ffffffffffffffffffffffffffffffffffffffff851633838b610d16565b6000888888888888888d604051602001610b45989796959493929190611d4b565b604051602081830303815290604052805190602001209050600060026000815480929190610b7290611be5565b909155506040805184815233602082015246818301526060810183905290519192507f1556007120ca92333c476f1b666d8d1d2e6f0dc3a56a92bcc3ddb88ac009e21a919081900360800190a1505050610bf1600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b50505050505050565b600082815260208190526040902060010154610c1581610df8565b6107ec8383610ef2565b60015460ff16610c8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161053e565b565b600154610100900460ff1615610c8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161053e565b600080821215610d125781600003610483565b5090565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610df29085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261128a565b50505050565b6108de8133611399565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166108a05760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610e943390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16156108a05760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600154610100900460ff16610c8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161053e565b611022610fa9565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526107ec9084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401610d70565b6110f5610c8d565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861106d3390565b73ffffffffffffffffffffffffffffffffffffffff82163b156111a7576108a073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168383611097565b6040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632e1a7d4d90602401600060405180830381600087803b15801561122f57600080fd5b505af1158015611243573d6000803e3d6000fd5b505060405173ffffffffffffffffffffffffffffffffffffffff8516925083156108fc02915083906000818181858888f193505050501580156107ec573d6000803e3d6000fd5b60006112ec826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166114519092919063ffffffff16565b905080516000148061130d57508080602001905181019061130d9190611dbd565b6107ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161053e565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166108a0576113d781611468565b6113e2836020611487565b6040516020016113f3929190611ddf565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261053e91600401611e60565b606061146084846000856116d1565b949350505050565b606061048373ffffffffffffffffffffffffffffffffffffffff831660145b60606000611496836002611e73565b6114a1906002611eb0565b67ffffffffffffffff8111156114b9576114b9611a41565b6040519080825280601f01601f1916602001820160405280156114e3576020820181803683370190505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061151a5761151a611ec8565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061157d5761157d611ec8565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006115b9846002611e73565b6115c4906001611eb0565b90505b6001811115611661577f303132333435363738396162636465660000000000000000000000000000000085600f166010811061160557611605611ec8565b1a60f81b82828151811061161b5761161b611ec8565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c9361165a81611ef7565b90506115c7565b5083156116ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161053e565b9392505050565b606082471015611763576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161053e565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161178c9190611f2c565b60006040518083038185875af1925050503d80600081146117c9576040519150601f19603f3d011682016040523d82523d6000602084013e6117ce565b606091505b50915091506117df878383876117ea565b979650505050505050565b6060831561187d5782516118765773ffffffffffffffffffffffffffffffffffffffff85163b611876576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161053e565b5081611460565b61146083838151156118925781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053e9190611e60565b6000602082840312156118d857600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146116ca57600080fd5b60006020828403121561191a57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461194557600080fd5b919050565b60008060008060008060a0878903121561196357600080fd5b86359550602087013567ffffffffffffffff8082111561198257600080fd5b818901915089601f83011261199657600080fd5b8135818111156119a557600080fd5b8a60208285010111156119b757600080fd5b6020830197508096505050506119cf60408801611921565b92506060870135915060808701358060070b81146119ec57600080fd5b809150509295509295509295565b60008060408385031215611a0d57600080fd5b82359150611a1d60208401611921565b90509250929050565b600060208284031215611a3857600080fd5b6116ca82611921565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112611a8157600080fd5b813567ffffffffffffffff80821115611a9c57611a9c611a41565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611ae257611ae2611a41565b81604052838152866020858801011115611afb57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600080600060e0888a031215611b3657600080fd5b87359650602088013595506040880135945060608801359350611b5b60808901611921565b925060a088013567ffffffffffffffff80821115611b7857600080fd5b611b848b838c01611a70565b935060c08a0135915080821115611b9a57600080fd5b50611ba78a828b01611a70565b91505092959891949750929550565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611c1757611c17611bb6565b5060010190565b60006101008b83528a60208401528960408401528860070b606084015287608084015273ffffffffffffffffffffffffffffffffffffffff80881660a08501528160c08501528582850152610120915085878386013760008487018301529390931660e083015250601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910101979650505050505050565b600060208284031215611cce57600080fd5b5051919050565b60005b83811015611cf0578181015183820152602001611cd8565b83811115610df25750506000910152565b60008151808452611d19816020860160208601611cd5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006101008a835289602084015288604084015287606084015273ffffffffffffffffffffffffffffffffffffffff871660808401528060a0840152611d9381840187611d01565b905082810360c0840152611da78186611d01565b9150508260e08301529998505050505050505050565b600060208284031215611dcf57600080fd5b815180151581146116ca57600080fd5b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611e17816017850160208801611cd5565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351611e54816028840160208801611cd5565b01602801949350505050565b6020815260006116ca6020830184611d01565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611eab57611eab611bb6565b500290565b60008219821115611ec357611ec3611bb6565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081611f0657611f06611bb6565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60008251611f3e818460208701611cd5565b919091019291505056fea164736f6c634300080a000a",
}

// VoyagerDepositHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use VoyagerDepositHandlerMetaData.ABI instead.
var VoyagerDepositHandlerABI = VoyagerDepositHandlerMetaData.ABI

// VoyagerDepositHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoyagerDepositHandlerMetaData.Bin instead.
var VoyagerDepositHandlerBin = VoyagerDepositHandlerMetaData.Bin

// DeployVoyagerDepositHandler deploys a new Ethereum contract, binding an instance of VoyagerDepositHandler to it.
func DeployVoyagerDepositHandler(auth *bind.TransactOpts, backend bind.ContractBackend, _wrappedNativeTokenAddress common.Address) (common.Address, *types.Transaction, *VoyagerDepositHandler, error) {
	parsed, err := VoyagerDepositHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoyagerDepositHandlerBin), backend, _wrappedNativeTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VoyagerDepositHandler{VoyagerDepositHandlerCaller: VoyagerDepositHandlerCaller{contract: contract}, VoyagerDepositHandlerTransactor: VoyagerDepositHandlerTransactor{contract: contract}, VoyagerDepositHandlerFilterer: VoyagerDepositHandlerFilterer{contract: contract}}, nil
}

// VoyagerDepositHandler is an auto generated Go binding around an Ethereum contract.
type VoyagerDepositHandler struct {
	VoyagerDepositHandlerCaller     // Read-only binding to the contract
	VoyagerDepositHandlerTransactor // Write-only binding to the contract
	VoyagerDepositHandlerFilterer   // Log filterer for contract events
}

// VoyagerDepositHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoyagerDepositHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoyagerDepositHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoyagerDepositHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoyagerDepositHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoyagerDepositHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoyagerDepositHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoyagerDepositHandlerSession struct {
	Contract     *VoyagerDepositHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VoyagerDepositHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoyagerDepositHandlerCallerSession struct {
	Contract *VoyagerDepositHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// VoyagerDepositHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoyagerDepositHandlerTransactorSession struct {
	Contract     *VoyagerDepositHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// VoyagerDepositHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoyagerDepositHandlerRaw struct {
	Contract *VoyagerDepositHandler // Generic contract binding to access the raw methods on
}

// VoyagerDepositHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoyagerDepositHandlerCallerRaw struct {
	Contract *VoyagerDepositHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// VoyagerDepositHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoyagerDepositHandlerTransactorRaw struct {
	Contract *VoyagerDepositHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoyagerDepositHandler creates a new instance of VoyagerDepositHandler, bound to a specific deployed contract.
func NewVoyagerDepositHandler(address common.Address, backend bind.ContractBackend) (*VoyagerDepositHandler, error) {
	contract, err := bindVoyagerDepositHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandler{VoyagerDepositHandlerCaller: VoyagerDepositHandlerCaller{contract: contract}, VoyagerDepositHandlerTransactor: VoyagerDepositHandlerTransactor{contract: contract}, VoyagerDepositHandlerFilterer: VoyagerDepositHandlerFilterer{contract: contract}}, nil
}

// NewVoyagerDepositHandlerCaller creates a new read-only instance of VoyagerDepositHandler, bound to a specific deployed contract.
func NewVoyagerDepositHandlerCaller(address common.Address, caller bind.ContractCaller) (*VoyagerDepositHandlerCaller, error) {
	contract, err := bindVoyagerDepositHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerCaller{contract: contract}, nil
}

// NewVoyagerDepositHandlerTransactor creates a new write-only instance of VoyagerDepositHandler, bound to a specific deployed contract.
func NewVoyagerDepositHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*VoyagerDepositHandlerTransactor, error) {
	contract, err := bindVoyagerDepositHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerTransactor{contract: contract}, nil
}

// NewVoyagerDepositHandlerFilterer creates a new log filterer instance of VoyagerDepositHandler, bound to a specific deployed contract.
func NewVoyagerDepositHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*VoyagerDepositHandlerFilterer, error) {
	contract, err := bindVoyagerDepositHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerFilterer{contract: contract}, nil
}

// bindVoyagerDepositHandler binds a generic wrapper to an already deployed contract.
func bindVoyagerDepositHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoyagerDepositHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoyagerDepositHandler *VoyagerDepositHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoyagerDepositHandler.Contract.VoyagerDepositHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoyagerDepositHandler *VoyagerDepositHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.VoyagerDepositHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoyagerDepositHandler *VoyagerDepositHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.VoyagerDepositHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoyagerDepositHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.DEFAULTADMINROLE(&_VoyagerDepositHandler.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.DEFAULTADMINROLE(&_VoyagerDepositHandler.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) MAXTRANSFERSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "MAX_TRANSFER_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _VoyagerDepositHandler.Contract.MAXTRANSFERSIZE(&_VoyagerDepositHandler.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _VoyagerDepositHandler.Contract.MAXTRANSFERSIZE(&_VoyagerDepositHandler.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) PAUSER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "PAUSER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) PAUSER() ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.PAUSER(&_VoyagerDepositHandler.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) PAUSER() ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.PAUSER(&_VoyagerDepositHandler.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) RESOURCESETTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "RESOURCE_SETTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) RESOURCESETTER() ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.RESOURCESETTER(&_VoyagerDepositHandler.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) RESOURCESETTER() ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.RESOURCESETTER(&_VoyagerDepositHandler.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) DepositNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "depositNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) DepositNonce() (*big.Int, error) {
	return _VoyagerDepositHandler.Contract.DepositNonce(&_VoyagerDepositHandler.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) DepositNonce() (*big.Int, error) {
	return _VoyagerDepositHandler.Contract.DepositNonce(&_VoyagerDepositHandler.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.GetRoleAdmin(&_VoyagerDepositHandler.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _VoyagerDepositHandler.Contract.GetRoleAdmin(&_VoyagerDepositHandler.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VoyagerDepositHandler.Contract.HasRole(&_VoyagerDepositHandler.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _VoyagerDepositHandler.Contract.HasRole(&_VoyagerDepositHandler.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) Paused() (bool, error) {
	return _VoyagerDepositHandler.Contract.Paused(&_VoyagerDepositHandler.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) Paused() (bool, error) {
	return _VoyagerDepositHandler.Contract.Paused(&_VoyagerDepositHandler.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoyagerDepositHandler.Contract.SupportsInterface(&_VoyagerDepositHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _VoyagerDepositHandler.Contract.SupportsInterface(&_VoyagerDepositHandler.CallOpts, interfaceId)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VoyagerDepositHandler.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) WrappedNativeToken() (common.Address, error) {
	return _VoyagerDepositHandler.Contract.WrappedNativeToken(&_VoyagerDepositHandler.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_VoyagerDepositHandler *VoyagerDepositHandlerCallerSession) WrappedNativeToken() (common.Address, error) {
	return _VoyagerDepositHandler.Contract.WrappedNativeToken(&_VoyagerDepositHandler.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.GrantRole(&_VoyagerDepositHandler.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.GrantRole(&_VoyagerDepositHandler.TransactOpts, role, account)
}

// IRelay is a paid mutator transaction binding the contract method 0x9418617c.
//
// Solidity: function iRelay(uint256 amount, bytes32 srcChainId, uint256 relayerFeePct, uint256 depositId, address destToken, bytes recipient, bytes depositor) payable returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) IRelay(opts *bind.TransactOpts, amount *big.Int, srcChainId [32]byte, relayerFeePct *big.Int, depositId *big.Int, destToken common.Address, recipient []byte, depositor []byte) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "iRelay", amount, srcChainId, relayerFeePct, depositId, destToken, recipient, depositor)
}

// IRelay is a paid mutator transaction binding the contract method 0x9418617c.
//
// Solidity: function iRelay(uint256 amount, bytes32 srcChainId, uint256 relayerFeePct, uint256 depositId, address destToken, bytes recipient, bytes depositor) payable returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) IRelay(amount *big.Int, srcChainId [32]byte, relayerFeePct *big.Int, depositId *big.Int, destToken common.Address, recipient []byte, depositor []byte) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.IRelay(&_VoyagerDepositHandler.TransactOpts, amount, srcChainId, relayerFeePct, depositId, destToken, recipient, depositor)
}

// IRelay is a paid mutator transaction binding the contract method 0x9418617c.
//
// Solidity: function iRelay(uint256 amount, bytes32 srcChainId, uint256 relayerFeePct, uint256 depositId, address destToken, bytes recipient, bytes depositor) payable returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) IRelay(amount *big.Int, srcChainId [32]byte, relayerFeePct *big.Int, depositId *big.Int, destToken common.Address, recipient []byte, depositor []byte) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.IRelay(&_VoyagerDepositHandler.TransactOpts, amount, srcChainId, relayerFeePct, depositId, destToken, recipient, depositor)
}

// ISend is a paid mutator transaction binding the contract method 0x26d1172a.
//
// Solidity: function iSend(bytes32 destChainIdBytes, bytes recipient, address srcToken, uint256 amount, int64 relayerFeePct) payable returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) ISend(opts *bind.TransactOpts, destChainIdBytes [32]byte, recipient []byte, srcToken common.Address, amount *big.Int, relayerFeePct int64) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "iSend", destChainIdBytes, recipient, srcToken, amount, relayerFeePct)
}

// ISend is a paid mutator transaction binding the contract method 0x26d1172a.
//
// Solidity: function iSend(bytes32 destChainIdBytes, bytes recipient, address srcToken, uint256 amount, int64 relayerFeePct) payable returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) ISend(destChainIdBytes [32]byte, recipient []byte, srcToken common.Address, amount *big.Int, relayerFeePct int64) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.ISend(&_VoyagerDepositHandler.TransactOpts, destChainIdBytes, recipient, srcToken, amount, relayerFeePct)
}

// ISend is a paid mutator transaction binding the contract method 0x26d1172a.
//
// Solidity: function iSend(bytes32 destChainIdBytes, bytes recipient, address srcToken, uint256 amount, int64 relayerFeePct) payable returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) ISend(destChainIdBytes [32]byte, recipient []byte, srcToken common.Address, amount *big.Int, relayerFeePct int64) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.ISend(&_VoyagerDepositHandler.TransactOpts, destChainIdBytes, recipient, srcToken, amount, relayerFeePct)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) Pause() (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.Pause(&_VoyagerDepositHandler.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) Pause() (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.Pause(&_VoyagerDepositHandler.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.RenounceRole(&_VoyagerDepositHandler.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.RenounceRole(&_VoyagerDepositHandler.TransactOpts, role, account)
}

// Rescue is a paid mutator transaction binding the contract method 0x839006f2.
//
// Solidity: function rescue(address token) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) Rescue(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "rescue", token)
}

// Rescue is a paid mutator transaction binding the contract method 0x839006f2.
//
// Solidity: function rescue(address token) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) Rescue(token common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.Rescue(&_VoyagerDepositHandler.TransactOpts, token)
}

// Rescue is a paid mutator transaction binding the contract method 0x839006f2.
//
// Solidity: function rescue(address token) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) Rescue(token common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.Rescue(&_VoyagerDepositHandler.TransactOpts, token)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.RevokeRole(&_VoyagerDepositHandler.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.RevokeRole(&_VoyagerDepositHandler.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoyagerDepositHandler.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerSession) Unpause() (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.Unpause(&_VoyagerDepositHandler.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_VoyagerDepositHandler *VoyagerDepositHandlerTransactorSession) Unpause() (*types.Transaction, error) {
	return _VoyagerDepositHandler.Contract.Unpause(&_VoyagerDepositHandler.TransactOpts)
}

// VoyagerDepositHandlerFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerFundsDepositedIterator struct {
	Event *VoyagerDepositHandlerFundsDeposited // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerFundsDeposited)
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
		it.Event = new(VoyagerDepositHandlerFundsDeposited)
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
func (it *VoyagerDepositHandlerFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerFundsDeposited represents a FundsDeposited event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerFundsDeposited struct {
	Amount           *big.Int
	OriginChainId    *big.Int
	DestChainIdBytes [32]byte
	RelayerFeePct    int64
	DepositId        *big.Int
	SrcToken         common.Address
	Recipient        []byte
	Depositor        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0x99f98099e0db79801cb9821f6ccfd8ea7abdd2ee8ae98f397a66bb0b48192d3c.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, bytes32 destChainIdBytes, int64 relayerFeePct, uint256 depositId, address srcToken, bytes recipient, address depositor)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterFundsDeposited(opts *bind.FilterOpts) (*VoyagerDepositHandlerFundsDepositedIterator, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerFundsDepositedIterator{contract: _VoyagerDepositHandler.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0x99f98099e0db79801cb9821f6ccfd8ea7abdd2ee8ae98f397a66bb0b48192d3c.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, bytes32 destChainIdBytes, int64 relayerFeePct, uint256 depositId, address srcToken, bytes recipient, address depositor)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerFundsDeposited) (event.Subscription, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerFundsDeposited)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
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

// ParseFundsDeposited is a log parse operation binding the contract event 0x99f98099e0db79801cb9821f6ccfd8ea7abdd2ee8ae98f397a66bb0b48192d3c.
//
// Solidity: event FundsDeposited(uint256 amount, uint256 originChainId, bytes32 destChainIdBytes, int64 relayerFeePct, uint256 depositId, address srcToken, bytes recipient, address depositor)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseFundsDeposited(log types.Log) (*VoyagerDepositHandlerFundsDeposited, error) {
	event := new(VoyagerDepositHandlerFundsDeposited)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerFundsPaidIterator is returned from FilterFundsPaid and is used to iterate over the raw logs and unpacked data for FundsPaid events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerFundsPaidIterator struct {
	Event *VoyagerDepositHandlerFundsPaid // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerFundsPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerFundsPaid)
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
		it.Event = new(VoyagerDepositHandlerFundsPaid)
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
func (it *VoyagerDepositHandlerFundsPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerFundsPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerFundsPaid represents a FundsPaid event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerFundsPaid struct {
	MessageHash [32]byte
	Forwarder   common.Address
	ChainId     *big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaid is a free log retrieval operation binding the contract event 0x1556007120ca92333c476f1b666d8d1d2e6f0dc3a56a92bcc3ddb88ac009e21a.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 chainId, uint256 nonce)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterFundsPaid(opts *bind.FilterOpts) (*VoyagerDepositHandlerFundsPaidIterator, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerFundsPaidIterator{contract: _VoyagerDepositHandler.contract, event: "FundsPaid", logs: logs, sub: sub}, nil
}

// WatchFundsPaid is a free log subscription operation binding the contract event 0x1556007120ca92333c476f1b666d8d1d2e6f0dc3a56a92bcc3ddb88ac009e21a.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 chainId, uint256 nonce)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchFundsPaid(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerFundsPaid) (event.Subscription, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerFundsPaid)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "FundsPaid", log); err != nil {
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

// ParseFundsPaid is a log parse operation binding the contract event 0x1556007120ca92333c476f1b666d8d1d2e6f0dc3a56a92bcc3ddb88ac009e21a.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 chainId, uint256 nonce)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseFundsPaid(log types.Log) (*VoyagerDepositHandlerFundsPaid, error) {
	event := new(VoyagerDepositHandlerFundsPaid)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "FundsPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerISendEventIterator is returned from FilterISendEvent and is used to iterate over the raw logs and unpacked data for ISendEvent events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerISendEventIterator struct {
	Event *VoyagerDepositHandlerISendEvent // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerISendEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerISendEvent)
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
		it.Event = new(VoyagerDepositHandlerISendEvent)
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
func (it *VoyagerDepositHandlerISendEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerISendEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerISendEvent represents a ISendEvent event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerISendEvent struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterISendEvent is a free log retrieval operation binding the contract event 0xc3156a2ab1789002b2e60f9db0027bff415cec7c8eb357e1463958c9b0c87fd7.
//
// Solidity: event ISendEvent(address arg0)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterISendEvent(opts *bind.FilterOpts) (*VoyagerDepositHandlerISendEventIterator, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "ISendEvent")
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerISendEventIterator{contract: _VoyagerDepositHandler.contract, event: "ISendEvent", logs: logs, sub: sub}, nil
}

// WatchISendEvent is a free log subscription operation binding the contract event 0xc3156a2ab1789002b2e60f9db0027bff415cec7c8eb357e1463958c9b0c87fd7.
//
// Solidity: event ISendEvent(address arg0)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchISendEvent(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerISendEvent) (event.Subscription, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "ISendEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerISendEvent)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "ISendEvent", log); err != nil {
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

// ParseISendEvent is a log parse operation binding the contract event 0xc3156a2ab1789002b2e60f9db0027bff415cec7c8eb357e1463958c9b0c87fd7.
//
// Solidity: event ISendEvent(address arg0)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseISendEvent(log types.Log) (*VoyagerDepositHandlerISendEvent, error) {
	event := new(VoyagerDepositHandlerISendEvent)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "ISendEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerPausedIterator struct {
	Event *VoyagerDepositHandlerPaused // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerPaused)
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
		it.Event = new(VoyagerDepositHandlerPaused)
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
func (it *VoyagerDepositHandlerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerPaused represents a Paused event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterPaused(opts *bind.FilterOpts) (*VoyagerDepositHandlerPausedIterator, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerPausedIterator{contract: _VoyagerDepositHandler.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerPaused) (event.Subscription, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerPaused)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParsePaused(log types.Log) (*VoyagerDepositHandlerPaused, error) {
	event := new(VoyagerDepositHandlerPaused)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerRoleAdminChangedIterator struct {
	Event *VoyagerDepositHandlerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerRoleAdminChanged)
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
		it.Event = new(VoyagerDepositHandlerRoleAdminChanged)
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
func (it *VoyagerDepositHandlerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerRoleAdminChanged represents a RoleAdminChanged event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VoyagerDepositHandlerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerRoleAdminChangedIterator{contract: _VoyagerDepositHandler.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerRoleAdminChanged)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseRoleAdminChanged(log types.Log) (*VoyagerDepositHandlerRoleAdminChanged, error) {
	event := new(VoyagerDepositHandlerRoleAdminChanged)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerRoleGrantedIterator struct {
	Event *VoyagerDepositHandlerRoleGranted // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerRoleGranted)
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
		it.Event = new(VoyagerDepositHandlerRoleGranted)
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
func (it *VoyagerDepositHandlerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerRoleGranted represents a RoleGranted event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoyagerDepositHandlerRoleGrantedIterator, error) {

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

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerRoleGrantedIterator{contract: _VoyagerDepositHandler.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerRoleGranted)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseRoleGranted(log types.Log) (*VoyagerDepositHandlerRoleGranted, error) {
	event := new(VoyagerDepositHandlerRoleGranted)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerRoleRevokedIterator struct {
	Event *VoyagerDepositHandlerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerRoleRevoked)
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
		it.Event = new(VoyagerDepositHandlerRoleRevoked)
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
func (it *VoyagerDepositHandlerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerRoleRevoked represents a RoleRevoked event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VoyagerDepositHandlerRoleRevokedIterator, error) {

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

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerRoleRevokedIterator{contract: _VoyagerDepositHandler.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerRoleRevoked)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseRoleRevoked(log types.Log) (*VoyagerDepositHandlerRoleRevoked, error) {
	event := new(VoyagerDepositHandlerRoleRevoked)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoyagerDepositHandlerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerUnpausedIterator struct {
	Event *VoyagerDepositHandlerUnpaused // Event containing the contract specifics and raw log

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
func (it *VoyagerDepositHandlerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoyagerDepositHandlerUnpaused)
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
		it.Event = new(VoyagerDepositHandlerUnpaused)
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
func (it *VoyagerDepositHandlerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoyagerDepositHandlerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoyagerDepositHandlerUnpaused represents a Unpaused event raised by the VoyagerDepositHandler contract.
type VoyagerDepositHandlerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VoyagerDepositHandlerUnpausedIterator, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VoyagerDepositHandlerUnpausedIterator{contract: _VoyagerDepositHandler.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VoyagerDepositHandlerUnpaused) (event.Subscription, error) {

	logs, sub, err := _VoyagerDepositHandler.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoyagerDepositHandlerUnpaused)
				if err := _VoyagerDepositHandler.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_VoyagerDepositHandler *VoyagerDepositHandlerFilterer) ParseUnpaused(log types.Log) (*VoyagerDepositHandlerUnpaused, error) {
	event := new(VoyagerDepositHandlerUnpaused)
	if err := _VoyagerDepositHandler.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
