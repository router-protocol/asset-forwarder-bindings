// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AssetForwarder

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

// IAssetForwarderDepositData is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderDepositData struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestAmount       *big.Int
	SrcToken         common.Address
	RefundRecipient  common.Address
	DestChainIdBytes [32]byte
}

// IAssetForwarderDestDetails is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderDestDetails struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}

// IAssetForwarderRelayData is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderRelayData struct {
	Amount     *big.Int
	SrcChainId [32]byte
	DepositId  *big.Int
	DestToken  common.Address
	Recipient  common.Address
}

// IAssetForwarderRelayDataMessage is an auto generated low-level Go binding around an user-defined struct.
type IAssetForwarderRelayDataMessage struct {
	Amount     *big.Int
	SrcChainId [32]byte
	DepositId  *big.Int
	DestToken  common.Address
	Recipient  common.Address
	Message    []byte
}

// AssetForwarderMetaData contains all meta data concerning the AssetForwarder contract.
var AssetForwarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_minGasThreshhold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRefundData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageExcecutionFailedWithLowGas\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"CommunityPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"DepositInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDepositedWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"FundsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"execFlag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"execData\",\"type\":\"bytes\"}],\"name\":\"FundsPaidWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_THRESHHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESOURCE_SETTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"communityPause\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"destDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executeRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"iDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"}],\"name\":\"iDepositInfoUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"iDepositMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"iDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"requestSender\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"packet\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"iReceive\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCommunityPauseEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerMiddlewareBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_destChainIdBytes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"internalType\":\"structIAssetForwarder.DestDetails[]\",\"name\":\"_destDetails\",\"type\":\"tuple[]\"}],\"name\":\"setDestDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleCommunityPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minPauseStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPauseStakeAmount\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"}],\"name\":\"updateTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600d805460ff191660011790553480156200001e57600080fd5b5060405162004a2438038062004a248339810160408190526200004191620001e3565b6001805461ffff1916811790556001600160a01b03868116608052600580546001600160a01b03199081168684161790915560038054821688841617905560048054909116918616919091179055815160208301206002556009819055620000ab6000336200010f565b620000d77f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb078336200010f565b620001037f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c336200010f565b5050505050506200030b565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16620001ac576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556200016b3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b80516001600160a01b0381168114620001c857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060c08789031215620001fd57600080fd5b6200020887620001b0565b9550602062000219818901620001b0565b95506200022960408901620001b0565b94506200023960608901620001b0565b60808901519094506001600160401b03808211156200025757600080fd5b818a0191508a601f8301126200026c57600080fd5b815181811115620002815762000281620001cd565b604051601f8201601f19908116603f01168101908382118183101715620002ac57620002ac620001cd565b816040528281528d86848701011115620002c557600080fd5b600093505b82841015620002e95784840186015181850187015292850192620002ca565b600086848301015280975050505050505060a087015190509295509295509295565b6080516146d3620003516000396000818161034701528181610b8801528181610c1c015281816125d90152818161265a01528181612877015261290b01526146d36000f3fe6080604052600436106102c65760003560e01c806378e0f9bd11610179578063c44e947e116100d6578063de35f5cb1161008a578063f452ed4d11610064578063f452ed4d14610816578063f627df9414610829578063fd5ad37c1461083e57600080fd5b8063de35f5cb146107bd578063eb0cde1d146107d3578063f215f1481461080057600080fd5b8063d9dc8694116100bb578063d9dc869414610732578063db618e2a14610766578063ddd224f11461079a57600080fd5b8063c44e947e146106fc578063d547741f1461071257600080fd5b8063981a8a021161012d578063ac9650d811610112578063ac9650d8146106a6578063ad7c17ee146106d3578063b19941a9146106e657600080fd5b8063981a8a0214610627578063a217fddf1461069157600080fd5b80638456cb591161015e5780638456cb59146105ac5780638a27fecb146105c157806391d14854146105d657600080fd5b806378e0f9bd1461056c5780637a4e4ecf1461058c57600080fd5b80633f4ba83a116102275780635ac62700116101db57806364778c1f116101c057806364778c1f1461053e5780636696821b146105515780636fb003da1461055957600080fd5b80635ac62700146105015780635c975abb1461052157600080fd5b8063461178301161020c57806346117830146104a85780634b7b9476146104d5578063567e98f9146104eb57600080fd5b80633f4ba83a146104795780634463182f1461048e57600080fd5b8063248a9ca31161027e57806336568abe1161026357806336568abe146104195780633e28c7d2146104395780633e413bee1461044c57600080fd5b8063248a9ca3146103bb5780632f2ff15d146103f957600080fd5b80630421caf0116102af5780630421caf01461032257806317fcb39b146103355780631aa6485a1461038e57600080fd5b80630171958a146102cb57806301ffc9a7146102ed575b600080fd5b3480156102d757600080fd5b506102eb6102e6366004613842565b61086e565b005b3480156102f957600080fd5b5061030d6103083660046138fb565b6109ea565b60405190151581526020015b60405180910390f35b6102eb610330366004613a62565b610a83565b34801561034157600080fd5b506103697f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610319565b34801561039a57600080fd5b506103ae6103a9366004613b47565b610d20565b6040516103199190613c4a565b3480156103c757600080fd5b506103eb6103d6366004613c5d565b60009081526020819052604090206001015490565b604051908152602001610319565b34801561040557600080fd5b506102eb610414366004613c76565b610fa3565b34801561042557600080fd5b506102eb610434366004613c76565b610fcd565b6102eb610447366004613ca6565b611080565b34801561045857600080fd5b506004546103699073ffffffffffffffffffffffffffffffffffffffff1681565b34801561048557600080fd5b506102eb6113d7565b34801561049a57600080fd5b50600d5461030d9060ff1681565b3480156104b457600080fd5b506005546103699073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104e157600080fd5b506103eb60095481565b3480156104f757600080fd5b506103eb600c5481565b34801561050d57600080fd5b506102eb61051c366004613cd8565b611414565b34801561052d57600080fd5b50600154610100900460ff1661030d565b6102eb61054c366004613d46565b611586565b6102eb6118e2565b6102eb610567366004613dc5565b611a9f565b34801561057857600080fd5b506102eb610587366004613e76565b611f57565b34801561059857600080fd5b506102eb6105a7366004613e93565b611fc9565b3480156105b857600080fd5b506102eb6121c3565b3480156105cd57600080fd5b506102eb6121fd565b3480156105e257600080fd5b5061030d6105f1366004613c76565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561063357600080fd5b5061066d610642366004613c5d565b60076020526000908152604090208054600182015460029092015463ffffffff909116919060ff1683565b6040805163ffffffff90941684526020840192909252151590820152606001610319565b34801561069d57600080fd5b506103eb600081565b3480156106b257600080fd5b506106c66106c1366004613ebf565b6122dc565b6040516103199190613f34565b6102eb6106e1366004613fb4565b61244e565b3480156106f257600080fd5b506103eb600a5481565b34801561070857600080fd5b506103eb60025481565b34801561071e57600080fd5b506102eb61072d366004613c76565b61274d565b34801561073e57600080fd5b506103eb7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c81565b34801561077257600080fd5b506103eb7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07881565b3480156107a657600080fd5b506103eb6ec097ce7bc90715b34b9f100000000081565b3480156107c957600080fd5b506103eb60065481565b3480156107df57600080fd5b506003546103699073ffffffffffffffffffffffffffffffffffffffff1681565b34801561080c57600080fd5b506103eb600b5481565b6102eb610824366004613ffe565b612772565b34801561083557600080fd5b506102eb612a0e565b34801561084a57600080fd5b5061030d610859366004613c5d565b60086020526000908152604090205460ff1681565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07861089881612a4c565b8151835114610908576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e76616c6964206c656e67746800000000000000000000000000000000000060448201526064015b60405180910390fd5b60005b82518110156109e45782818151811061092657610926614074565b60200260200101516007600086848151811061094457610944614074565b6020908102919091018101518252818101929092526040908101600020835181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff9091161781559183015160018301559190910151600290910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055806109dc816140d2565b91505061090b565b50505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610a7d57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610a8b612a56565b610ab8600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b610ac0612ac4565b6ec097ce7bc90715b34b9f100000000084602001511115610b0d576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b44846060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15610c495734846020015114610b86576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015610bee57600080fd5b505af1158015610c02573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606087015250610c7e9050565b610c7e33308660200151876060015173ffffffffffffffffffffffffffffffffffffffff16612b36909392919063ffffffff16565b7f3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f846000015185602001518660a001518760400151600660008154610cc2906140d2565b918290555060608a015160808b0151604051610ce9979695949392918b918d908c9061410a565b60405180910390a16109e4600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b60035460609073ffffffffffffffffffffffffffffffffffffffff163314610d74576040517ffc9dfe8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8585604051610d8492919061419c565b604051809103902060025414610dc6576040517f23dfe1fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600086806020019051810190610ddf9190614212565b92509250925060008251905081518114610e25576040517f94809d2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015610f8557610e7c848281518110610e4557610e45614074565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b610ee257610edd85848381518110610e9657610e96614074565b6020026020010151868481518110610eb057610eb0614074565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16612c129092919063ffffffff16565b610f73565b60008573ffffffffffffffffffffffffffffffffffffffff16848381518110610f0d57610f0d614074565b602002602001015160405160006040518083038185875af1925050503d8060008114610f55576040519150601f19603f3d011682016040523d82523d6000602084013e610f5a565b606091505b5090915050600181151514610f7157610f716142e1565b505b80610f7d816140d2565b915050610e28565b50506040805160208101909152600081529998505050505050505050565b600082815260208190526040902060010154610fbe81612a4c565b610fc88383612c68565b505050565b73ffffffffffffffffffffffffffffffffffffffff81163314611072576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084016108ff565b61107c8282612d58565b5050565b611088612a56565b6110b5600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6110bd612ac4565b60008381526007602052604090206002015460ff1680156110f5575060045473ffffffffffffffffffffffffffffffffffffffff1615155b611181576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f75736463206e6f7420737570706f7274656420656974686572206f6e2073726360448201527f206f6e2064737420636861696e0000000000000000000000000000000000000060648201526084016108ff565b60008381526007602052604090206001015434146111cb576040517f58d620b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6ec097ce7bc90715b34b9f1000000000811115611214576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004546112399073ffffffffffffffffffffffffffffffffffffffff16333084612b36565b6005546004546112639173ffffffffffffffffffffffffffffffffffffffff918216911683612e0f565b600554600084815260076020526040808220546004805492517f6fd3504e00000000000000000000000000000000000000000000000000000000815290810186905263ffffffff90911660248201526044810186905273ffffffffffffffffffffffffffffffffffffffff918216606482015291921690636fd3504e906084016020604051808303816000875af1158015611302573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113269190614310565b600454604080518881526020810186905290810187905267ffffffffffffffff8316606082015273ffffffffffffffffffffffffffffffffffffffff909116608082015260a081018590523360c08201529091507f297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d99060e00160405180910390a1506109e4600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c61140181612a4c565b611409612f08565b611411612f79565b50565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07861143e81612a4c565b8660010361148b57600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff881617905561157d565b866002036114b45784846040516114a392919061419c565b60405190819003902060025561157d565b8660030361157d5781831115611572576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604560248201527f6d696e50617573655374616b65416d6f756e74206d757374206265206c65737360448201527f207468616e206f7220657175616c20746f206d617850617573655374616b654160648201527f6d6f756e74000000000000000000000000000000000000000000000000000000608482015260a4016108ff565b600a839055600b8290555b50505050505050565b61158e612a56565b6115bb600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6115c3612ac4565b80516ec097ce7bc90715b34b9f1000000000101561160d576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516020808301516040808501516060808701516080808901518551978801989098529386019490945284015273ffffffffffffffffffffffffffffffffffffffff9182169083015290911660a08201523060c082015260009060e001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600890935291205490915060ff16156116ea576040517f7448c64c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260086020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560608201516117599073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b156118175781513414611798576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151825160405160009273ffffffffffffffffffffffffffffffffffffffff1691908381818185875af1925050503d80600081146117f5576040519150601f19603f3d011682016040523d82523d6000602084013e6117fa565b606091505b5090915050600181151514611811576118116142e1565b50611848565b6080820151825160608401516118489273ffffffffffffffffffffffffffffffffffffffff90911691339190612b36565b7f0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464813360066000815461187a906140d2565b91829055506040805193845273ffffffffffffffffffffffffffffffffffffffff90921660208401529082015260600160405180910390a150611411600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b6118ea612ac4565b600d5460ff16611956576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f436f6d6d756e6974792070617573652069732064697361626c6564000000000060448201526064016108ff565b600a54158015906119685750600b5415155b6119ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f536574205374616b6520416d6f756e742052616e67650000000000000000000060448201526064016108ff565b600a5434101580156119e25750600b543411155b611a48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f5374616b6520616d6f756e74206f7574206f662072616e67650000000000000060448201526064016108ff565b600034600c54611a58919061433a565b600c8190559050611a67612ff6565b60405134815233907f9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d9060200160405180910390a250565b611aa7612a56565b611ad4600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b611adc612ac4565b80516ec097ce7bc90715b34b9f10000000001015611b26576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516020808301516040808501516060860151608087015160a08801519351600097611b5997909695309290910161434d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600890935291205490915060ff1615611bd9576040517f7448c64c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260086020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556060820151611c489073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15611d065781513414611c87576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151825160405160009273ffffffffffffffffffffffffffffffffffffffff1691908381818185875af1925050503d8060008114611ce4576040519150601f19603f3d011682016040523d82523d6000602084013e611ce9565b606091505b5090915050600181151514611d0057611d006142e1565b50611d37565b608082015182516060840151611d379273ffffffffffffffffffffffffffffffffffffffff90911691339190612b36565b60606000611d4984608001513b151590565b8015611d5a575060008460a0015151115b15611ed357836080015173ffffffffffffffffffffffffffffffffffffffff165a6060860151865160a08801516040517fd00a2d5f0000000000000000000000000000000000000000000000000000000093611dbc93909290916024016143ad565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909416939093179092529051611e4591906143eb565b60006040518083038160008787f1925050503d8060008114611e83576040519150601f19603f3d011682016040523d82523d6000602084013e611e88565b606091505b509250905080158015611e9c57506009545a105b15611ed3576040517f9e82ca7900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad8333600660008154611f05906140d2565b9182905550604051611f1d9392919086908890614407565b60405180910390a1505050611411600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb078611f8181612a4c565b50600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000611fd481612a4c565b611fdc612a56565b612009600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff84160361209e57604051600090339084908381818185875af1925050503d806000811461207c576040519150601f19603f3d011682016040523d82523d6000602084013e612081565b606091505b5090915050600181151514612098576120986142e1565b50612194565b6040513360248201526044810183905260009073ffffffffffffffffffffffffffffffffffffffff851690606401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790525161214d91906143eb565b6000604051808303816000865af19150503d806000811461218a576040519150601f19603f3d011682016040523d82523d6000602084013e61218f565b606091505b505050505b610fc8600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c6121ed81612a4c565b6121f5612ac4565b611411612ff6565b600061220881612a4c565b600c54471015612274576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e73756666696369656e742066756e6473000000000000000000000000000060448201526064016108ff565b600c8054600091829055604051909190339083908381818185875af1925050503d80600081146122c0576040519150601f19603f3d011682016040523d82523d6000602084013e6122c5565b606091505b5090915050600181151514610fc857610fc86142e1565b60608167ffffffffffffffff8111156122f7576122f7613691565b60405190808252806020026020018201604052801561232a57816020015b60608152602001906001900390816123155790505b50905060005b82811015612447576000803086868581811061234e5761234e614074565b9050602002810190612360919061444a565b60405161236e92919061419c565b600060405180830381855af49150503d80600081146123a9576040519150601f19603f3d011682016040523d82523d6000602084013e6123ae565b606091505b509150915081612414576044815110156123c757600080fd5b600481019050808060200190518101906123e191906144af565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ff9190613c4a565b8084848151811061242757612427614074565b60200260200101819052505050808061243f906140d2565b915050612330565b5092915050565b612456612a56565b612483600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b61248b612ac4565b801561252457341561249f5761249f6142e1565b7f86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1846000846006600081546124d3906140d2565b91829055506040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152600160808201523360a082015260c00160405180910390a161271e565b6ec097ce7bc90715b34b9f100000000083111561256d576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff851603612680573483146125d7576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561263f57600080fd5b505af1158015612653573d6000803e3d6000fd5b50505050507f000000000000000000000000000000000000000000000000000000000000000093506126a2565b6126a273ffffffffffffffffffffffffffffffffffffffff8516333086612b36565b7f86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b18484846006600081546126d5906140d2565b91829055506040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152600060808201523360a082015260c001610ce9565b6109e4600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b60008281526020819052604090206001015461276881612a4c565b610fc88383612d58565b61277a612a56565b6127a7600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6127af612ac4565b6ec097ce7bc90715b34b9f1000000000836020015111156127fc576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612833836060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b156129385734836020015114612875576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b1580156128dd57600080fd5b505af11580156128f1573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001660608601525061296d9050565b61296d33308560200151866060015173ffffffffffffffffffffffffffffffffffffffff16612b36909392919063ffffffff16565b7f6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad836000015184602001518560a0015186604001516006600081546129b1906140d2565b9182905550606089015160808a01516040516129d797969594939291908a908c9061451d565b60405180910390a1610fc8600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b6000612a1981612a4c565b50600d80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00811660ff90911615179055565b6114118133613052565b60015460ff16612ac2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016108ff565b565b600154610100900460ff1615612ac2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016108ff565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526109e49085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261310a565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610fc89084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401612b90565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661107c5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612cfa3390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161561107c5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83811660248301526000919085169063dd62ed3e90604401602060405180830381865afa158015612e85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ea9919061459a565b90506109e4847f095ea7b30000000000000000000000000000000000000000000000000000000085612edb868661433a565b60405173ffffffffffffffffffffffffffffffffffffffff90921660248301526044820152606401612b90565b600154610100900460ff16612ac2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016108ff565b612f81612f08565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b612ffe612ac4565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612fcc3390565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661107c5761309081613219565b61309b836020613238565b6040516020016130ac9291906145b3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a00000000000000000000000000000000000000000000000000000000082526108ff91600401613c4a565b600061316c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166134829092919063ffffffff16565b905080516000148061318d57508080602001905181019061318d9190614634565b610fc8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016108ff565b6060610a7d73ffffffffffffffffffffffffffffffffffffffff831660145b60606000613247836002614651565b61325290600261433a565b67ffffffffffffffff81111561326a5761326a613691565b6040519080825280601f01601f191660200182016040528015613294576020820181803683370190505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106132cb576132cb614074565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061332e5761332e614074565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600061336a846002614651565b61337590600161433a565b90505b6001811115613412577f303132333435363738396162636465660000000000000000000000000000000085600f16601081106133b6576133b6614074565b1a60f81b8282815181106133cc576133cc614074565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c9361340b81614668565b9050613378565b50831561347b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016108ff565b9392505050565b60606134918484600085613499565b949350505050565b60608247101561352b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016108ff565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161355491906143eb565b60006040518083038185875af1925050503d8060008114613591576040519150601f19603f3d011682016040523d82523d6000602084013e613596565b606091505b50915091506135a7878383876135b2565b979650505050505050565b606083156136485782516000036136415773ffffffffffffffffffffffffffffffffffffffff85163b613641576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016108ff565b5081613491565b613491838381511561365d5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ff9190613c4a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156136e3576136e3613691565b60405290565b60405160c0810167ffffffffffffffff811182821017156136e3576136e3613691565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561375357613753613691565b604052919050565b600067ffffffffffffffff82111561377557613775613691565b5060051b60200190565b801515811461141157600080fd5b600082601f83011261379e57600080fd5b813560206137b36137ae8361375b565b61370c565b828152606092830285018201928282019190878511156137d257600080fd5b8387015b858110156138355781818a0312156137ee5760008081fd5b6137f66136c0565b813563ffffffff8116811461380b5760008081fd5b815281860135868201526040808301356138248161377f565b9082015284529284019281016137d6565b5090979650505050505050565b6000806040838503121561385557600080fd5b823567ffffffffffffffff8082111561386d57600080fd5b818501915085601f83011261388157600080fd5b813560206138916137ae8361375b565b82815260059290921b840181019181810190898411156138b057600080fd5b948201945b838610156138ce578535825294820194908201906138b5565b965050860135925050808211156138e457600080fd5b506138f18582860161378d565b9150509250929050565b60006020828403121561390d57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461347b57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461141157600080fd5b600060c0828403121561397157600080fd5b6139796136e9565b905081358152602082013560208201526040820135604082015260608201356139a18161393d565b606082015260808201356139b48161393d565b8060808301525060a082013560a082015292915050565b600067ffffffffffffffff8211156139e5576139e5613691565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613a2257600080fd5b8135613a306137ae826139cb565b818152846020838601011115613a4557600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806101208587031215613a7957600080fd5b613a83868661395f565b935060c085013567ffffffffffffffff80821115613aa057600080fd5b613aac88838901613a11565b945060e0870135915080821115613ac257600080fd5b613ace88838901613a11565b9350610100870135915080821115613ae557600080fd5b50613af287828801613a11565b91505092959194509250565b60008083601f840112613b1057600080fd5b50813567ffffffffffffffff811115613b2857600080fd5b602083019150836020828501011115613b4057600080fd5b9250929050565b600080600080600060608688031215613b5f57600080fd5b853567ffffffffffffffff80821115613b7757600080fd5b613b8389838a01613afe565b90975095506020880135915080821115613b9c57600080fd5b613ba889838a01613a11565b94506040880135915080821115613bbe57600080fd5b50613bcb88828901613afe565b969995985093965092949392505050565b60005b83811015613bf7578181015183820152602001613bdf565b50506000910152565b60008151808452613c18816020860160208601613bdc565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061347b6020830184613c00565b600060208284031215613c6f57600080fd5b5035919050565b60008060408385031215613c8957600080fd5b823591506020830135613c9b8161393d565b809150509250929050565b60008060008060808587031215613cbc57600080fd5b5050823594602084013594506040840135936060013592509050565b60008060008060008060a08789031215613cf157600080fd5b863595506020870135613d038161393d565b9450604087013567ffffffffffffffff811115613d1f57600080fd5b613d2b89828a01613afe565b979a9699509760608101359660809091013595509350505050565b600060a08284031215613d5857600080fd5b60405160a0810181811067ffffffffffffffff82111715613d7b57613d7b613691565b80604052508235815260208301356020820152604083013560408201526060830135613da68161393d565b60608201526080830135613db98161393d565b60808201529392505050565b600060208284031215613dd757600080fd5b813567ffffffffffffffff80821115613def57600080fd5b9083019060c08286031215613e0357600080fd5b613e0b6136e9565b8235815260208301356020820152604083013560408201526060830135613e318161393d565b60608201526080830135613e448161393d565b608082015260a083013582811115613e5b57600080fd5b613e6787828601613a11565b60a08301525095945050505050565b600060208284031215613e8857600080fd5b813561347b8161393d565b60008060408385031215613ea657600080fd5b8235613eb18161393d565b946020939093013593505050565b60008060208385031215613ed257600080fd5b823567ffffffffffffffff80821115613eea57600080fd5b818501915085601f830112613efe57600080fd5b813581811115613f0d57600080fd5b8660208260051b8501011115613f2257600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015613fa7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452613f95858351613c00565b94509285019290850190600101613f5b565b5092979650505050505050565b60008060008060808587031215613fca57600080fd5b8435613fd58161393d565b935060208501359250604085013591506060850135613ff38161377f565b939692955090935050565b6000806000610100848603121561401457600080fd5b61401e858561395f565b925060c084013567ffffffffffffffff8082111561403b57600080fd5b61404787838801613a11565b935060e086013591508082111561405d57600080fd5b5061406a86828701613a11565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614103576141036140a3565b5060010190565b60006101408c83528b60208401528a604084015289606084015288608084015273ffffffffffffffffffffffffffffffffffffffff80891660a08501528160c085015261415982850189613c00565b90871660e085015283810361010085015290506141768186613c00565b905082810361012084015261418b8185613c00565b9d9c50505050505050505050505050565b8183823760009101908152919050565b600082601f8301126141bd57600080fd5b815160206141cd6137ae8361375b565b82815260059290921b840181019181810190868411156141ec57600080fd5b8286015b8481101561420757805183529183019183016141f0565b509695505050505050565b60008060006060848603121561422757600080fd5b83516142328161393d565b8093505060208085015167ffffffffffffffff8082111561425257600080fd5b818701915087601f83011261426657600080fd5b81516142746137ae8261375b565b81815260059190911b8301840190848101908a83111561429357600080fd5b938501935b828510156142ba5784516142ab8161393d565b82529385019390850190614298565b60408a015190975094505050808311156142d357600080fd5b505061406a868287016141ac565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121561432257600080fd5b815167ffffffffffffffff8116811461347b57600080fd5b80820180821115610a7d57610a7d6140a3565b878152866020820152856040820152600073ffffffffffffffffffffffffffffffffffffffff8087166060840152808616608084015280851660a08401525060e060c08301526143a060e0830184613c00565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006143e26060830184613c00565b95945050505050565b600082516143fd818460208701613bdc565b9190910192915050565b85815273ffffffffffffffffffffffffffffffffffffffff85166020820152836040820152821515606082015260a0608082015260006135a760a0830184613c00565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261447f57600080fd5b83018035915067ffffffffffffffff82111561449a57600080fd5b602001915036819003821315613b4057600080fd5b6000602082840312156144c157600080fd5b815167ffffffffffffffff8111156144d857600080fd5b8201601f810184136144e957600080fd5b80516144f76137ae826139cb565b81815285602083850101111561450c57600080fd5b6143e2826020830160208601613bdc565b60006101208b83528a602084015289604084015288606084015287608084015273ffffffffffffffffffffffffffffffffffffffff80881660a085015280871660c0850152508060e084015261457581840186613c00565b905082810361010084015261458a8185613c00565b9c9b505050505050505050505050565b6000602082840312156145ac57600080fd5b5051919050565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000008152600083516145eb816017850160208801613bdc565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351614628816028840160208801613bdc565b01602801949350505050565b60006020828403121561464657600080fd5b815161347b8161377f565b8082028115828204841417610a7d57610a7d6140a3565b600081614677576146776140a3565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea26469706673582212206fe7b0d9c72ceaa01d5d7b48e636feb44b93bf36652e8dc7ecba03e6c94e145664736f6c63430008140033",
}

// AssetForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetForwarderMetaData.ABI instead.
var AssetForwarderABI = AssetForwarderMetaData.ABI

// AssetForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssetForwarderMetaData.Bin instead.
var AssetForwarderBin = AssetForwarderMetaData.Bin

// DeployAssetForwarder deploys a new Ethereum contract, binding an instance of AssetForwarder to it.
func DeployAssetForwarder(auth *bind.TransactOpts, backend bind.ContractBackend, _wrappedNativeTokenAddress common.Address, _gatewayContract common.Address, _usdcAddress common.Address, _tokenMessenger common.Address, _routerMiddlewareBase []byte, _minGasThreshhold *big.Int) (common.Address, *types.Transaction, *AssetForwarder, error) {
	parsed, err := AssetForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssetForwarderBin), backend, _wrappedNativeTokenAddress, _gatewayContract, _usdcAddress, _tokenMessenger, _routerMiddlewareBase, _minGasThreshhold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetForwarder{AssetForwarderCaller: AssetForwarderCaller{contract: contract}, AssetForwarderTransactor: AssetForwarderTransactor{contract: contract}, AssetForwarderFilterer: AssetForwarderFilterer{contract: contract}}, nil
}

// AssetForwarder is an auto generated Go binding around an Ethereum contract.
type AssetForwarder struct {
	AssetForwarderCaller     // Read-only binding to the contract
	AssetForwarderTransactor // Write-only binding to the contract
	AssetForwarderFilterer   // Log filterer for contract events
}

// AssetForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetForwarderSession struct {
	Contract     *AssetForwarder   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetForwarderCallerSession struct {
	Contract *AssetForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AssetForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetForwarderTransactorSession struct {
	Contract     *AssetForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AssetForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetForwarderRaw struct {
	Contract *AssetForwarder // Generic contract binding to access the raw methods on
}

// AssetForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetForwarderCallerRaw struct {
	Contract *AssetForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// AssetForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetForwarderTransactorRaw struct {
	Contract *AssetForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetForwarder creates a new instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarder(address common.Address, backend bind.ContractBackend) (*AssetForwarder, error) {
	contract, err := bindAssetForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetForwarder{AssetForwarderCaller: AssetForwarderCaller{contract: contract}, AssetForwarderTransactor: AssetForwarderTransactor{contract: contract}, AssetForwarderFilterer: AssetForwarderFilterer{contract: contract}}, nil
}

// NewAssetForwarderCaller creates a new read-only instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarderCaller(address common.Address, caller bind.ContractCaller) (*AssetForwarderCaller, error) {
	contract, err := bindAssetForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderCaller{contract: contract}, nil
}

// NewAssetForwarderTransactor creates a new write-only instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetForwarderTransactor, error) {
	contract, err := bindAssetForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderTransactor{contract: contract}, nil
}

// NewAssetForwarderFilterer creates a new log filterer instance of AssetForwarder, bound to a specific deployed contract.
func NewAssetForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetForwarderFilterer, error) {
	contract, err := bindAssetForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFilterer{contract: contract}, nil
}

// bindAssetForwarder binds a generic wrapper to an already deployed contract.
func bindAssetForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AssetForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetForwarder *AssetForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetForwarder.Contract.AssetForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetForwarder *AssetForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.Contract.AssetForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetForwarder *AssetForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetForwarder.Contract.AssetForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetForwarder *AssetForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AssetForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetForwarder *AssetForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetForwarder *AssetForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetForwarder.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetForwarder.Contract.DEFAULTADMINROLE(&_AssetForwarder.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AssetForwarder.Contract.DEFAULTADMINROLE(&_AssetForwarder.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) MAXTRANSFERSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "MAX_TRANSFER_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _AssetForwarder.Contract.MAXTRANSFERSIZE(&_AssetForwarder.CallOpts)
}

// MAXTRANSFERSIZE is a free data retrieval call binding the contract method 0xddd224f1.
//
// Solidity: function MAX_TRANSFER_SIZE() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) MAXTRANSFERSIZE() (*big.Int, error) {
	return _AssetForwarder.Contract.MAXTRANSFERSIZE(&_AssetForwarder.CallOpts)
}

// MINGASTHRESHHOLD is a free data retrieval call binding the contract method 0x4b7b9476.
//
// Solidity: function MIN_GAS_THRESHHOLD() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) MINGASTHRESHHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "MIN_GAS_THRESHHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINGASTHRESHHOLD is a free data retrieval call binding the contract method 0x4b7b9476.
//
// Solidity: function MIN_GAS_THRESHHOLD() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) MINGASTHRESHHOLD() (*big.Int, error) {
	return _AssetForwarder.Contract.MINGASTHRESHHOLD(&_AssetForwarder.CallOpts)
}

// MINGASTHRESHHOLD is a free data retrieval call binding the contract method 0x4b7b9476.
//
// Solidity: function MIN_GAS_THRESHHOLD() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) MINGASTHRESHHOLD() (*big.Int, error) {
	return _AssetForwarder.Contract.MINGASTHRESHHOLD(&_AssetForwarder.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) PAUSER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "PAUSER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) PAUSER() ([32]byte, error) {
	return _AssetForwarder.Contract.PAUSER(&_AssetForwarder.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) PAUSER() ([32]byte, error) {
	return _AssetForwarder.Contract.PAUSER(&_AssetForwarder.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) RESOURCESETTER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "RESOURCE_SETTER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) RESOURCESETTER() ([32]byte, error) {
	return _AssetForwarder.Contract.RESOURCESETTER(&_AssetForwarder.CallOpts)
}

// RESOURCESETTER is a free data retrieval call binding the contract method 0xdb618e2a.
//
// Solidity: function RESOURCE_SETTER() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) RESOURCESETTER() ([32]byte, error) {
	return _AssetForwarder.Contract.RESOURCESETTER(&_AssetForwarder.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) DepositNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "depositNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) DepositNonce() (*big.Int, error) {
	return _AssetForwarder.Contract.DepositNonce(&_AssetForwarder.CallOpts)
}

// DepositNonce is a free data retrieval call binding the contract method 0xde35f5cb.
//
// Solidity: function depositNonce() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) DepositNonce() (*big.Int, error) {
	return _AssetForwarder.Contract.DepositNonce(&_AssetForwarder.CallOpts)
}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 fee, bool isSet)
func (_AssetForwarder *AssetForwarderCaller) DestDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "destDetails", arg0)

	outstruct := new(struct {
		DomainId uint32
		Fee      *big.Int
		IsSet    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DomainId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsSet = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 fee, bool isSet)
func (_AssetForwarder *AssetForwarderSession) DestDetails(arg0 [32]byte) (struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}, error) {
	return _AssetForwarder.Contract.DestDetails(&_AssetForwarder.CallOpts, arg0)
}

// DestDetails is a free data retrieval call binding the contract method 0x981a8a02.
//
// Solidity: function destDetails(bytes32 ) view returns(uint32 domainId, uint256 fee, bool isSet)
func (_AssetForwarder *AssetForwarderCallerSession) DestDetails(arg0 [32]byte) (struct {
	DomainId uint32
	Fee      *big.Int
	IsSet    bool
}, error) {
	return _AssetForwarder.Contract.DestDetails(&_AssetForwarder.CallOpts, arg0)
}

// ExecuteRecord is a free data retrieval call binding the contract method 0xfd5ad37c.
//
// Solidity: function executeRecord(bytes32 ) view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) ExecuteRecord(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "executeRecord", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecuteRecord is a free data retrieval call binding the contract method 0xfd5ad37c.
//
// Solidity: function executeRecord(bytes32 ) view returns(bool)
func (_AssetForwarder *AssetForwarderSession) ExecuteRecord(arg0 [32]byte) (bool, error) {
	return _AssetForwarder.Contract.ExecuteRecord(&_AssetForwarder.CallOpts, arg0)
}

// ExecuteRecord is a free data retrieval call binding the contract method 0xfd5ad37c.
//
// Solidity: function executeRecord(bytes32 ) view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) ExecuteRecord(arg0 [32]byte) (bool, error) {
	return _AssetForwarder.Contract.ExecuteRecord(&_AssetForwarder.CallOpts, arg0)
}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) GatewayContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "gatewayContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_AssetForwarder *AssetForwarderSession) GatewayContract() (common.Address, error) {
	return _AssetForwarder.Contract.GatewayContract(&_AssetForwarder.CallOpts)
}

// GatewayContract is a free data retrieval call binding the contract method 0xeb0cde1d.
//
// Solidity: function gatewayContract() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) GatewayContract() (common.Address, error) {
	return _AssetForwarder.Contract.GatewayContract(&_AssetForwarder.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetForwarder.Contract.GetRoleAdmin(&_AssetForwarder.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AssetForwarder.Contract.GetRoleAdmin(&_AssetForwarder.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetForwarder *AssetForwarderSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetForwarder.Contract.HasRole(&_AssetForwarder.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AssetForwarder.Contract.HasRole(&_AssetForwarder.CallOpts, role, account)
}

// IsCommunityPauseEnabled is a free data retrieval call binding the contract method 0x4463182f.
//
// Solidity: function isCommunityPauseEnabled() view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) IsCommunityPauseEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "isCommunityPauseEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommunityPauseEnabled is a free data retrieval call binding the contract method 0x4463182f.
//
// Solidity: function isCommunityPauseEnabled() view returns(bool)
func (_AssetForwarder *AssetForwarderSession) IsCommunityPauseEnabled() (bool, error) {
	return _AssetForwarder.Contract.IsCommunityPauseEnabled(&_AssetForwarder.CallOpts)
}

// IsCommunityPauseEnabled is a free data retrieval call binding the contract method 0x4463182f.
//
// Solidity: function isCommunityPauseEnabled() view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) IsCommunityPauseEnabled() (bool, error) {
	return _AssetForwarder.Contract.IsCommunityPauseEnabled(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMax is a free data retrieval call binding the contract method 0xf215f148.
//
// Solidity: function pauseStakeAmountMax() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) PauseStakeAmountMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "pauseStakeAmountMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PauseStakeAmountMax is a free data retrieval call binding the contract method 0xf215f148.
//
// Solidity: function pauseStakeAmountMax() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) PauseStakeAmountMax() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMax(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMax is a free data retrieval call binding the contract method 0xf215f148.
//
// Solidity: function pauseStakeAmountMax() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) PauseStakeAmountMax() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMax(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMin is a free data retrieval call binding the contract method 0xb19941a9.
//
// Solidity: function pauseStakeAmountMin() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) PauseStakeAmountMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "pauseStakeAmountMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PauseStakeAmountMin is a free data retrieval call binding the contract method 0xb19941a9.
//
// Solidity: function pauseStakeAmountMin() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) PauseStakeAmountMin() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMin(&_AssetForwarder.CallOpts)
}

// PauseStakeAmountMin is a free data retrieval call binding the contract method 0xb19941a9.
//
// Solidity: function pauseStakeAmountMin() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) PauseStakeAmountMin() (*big.Int, error) {
	return _AssetForwarder.Contract.PauseStakeAmountMin(&_AssetForwarder.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AssetForwarder *AssetForwarderSession) Paused() (bool, error) {
	return _AssetForwarder.Contract.Paused(&_AssetForwarder.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) Paused() (bool, error) {
	return _AssetForwarder.Contract.Paused(&_AssetForwarder.CallOpts)
}

// RouterMiddlewareBase is a free data retrieval call binding the contract method 0xc44e947e.
//
// Solidity: function routerMiddlewareBase() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCaller) RouterMiddlewareBase(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "routerMiddlewareBase")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RouterMiddlewareBase is a free data retrieval call binding the contract method 0xc44e947e.
//
// Solidity: function routerMiddlewareBase() view returns(bytes32)
func (_AssetForwarder *AssetForwarderSession) RouterMiddlewareBase() ([32]byte, error) {
	return _AssetForwarder.Contract.RouterMiddlewareBase(&_AssetForwarder.CallOpts)
}

// RouterMiddlewareBase is a free data retrieval call binding the contract method 0xc44e947e.
//
// Solidity: function routerMiddlewareBase() view returns(bytes32)
func (_AssetForwarder *AssetForwarderCallerSession) RouterMiddlewareBase() ([32]byte, error) {
	return _AssetForwarder.Contract.RouterMiddlewareBase(&_AssetForwarder.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetForwarder *AssetForwarderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetForwarder.Contract.SupportsInterface(&_AssetForwarder.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AssetForwarder.Contract.SupportsInterface(&_AssetForwarder.CallOpts, interfaceId)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) TokenMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "tokenMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_AssetForwarder *AssetForwarderSession) TokenMessenger() (common.Address, error) {
	return _AssetForwarder.Contract.TokenMessenger(&_AssetForwarder.CallOpts)
}

// TokenMessenger is a free data retrieval call binding the contract method 0x46117830.
//
// Solidity: function tokenMessenger() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) TokenMessenger() (common.Address, error) {
	return _AssetForwarder.Contract.TokenMessenger(&_AssetForwarder.CallOpts)
}

// TotalStakedAmount is a free data retrieval call binding the contract method 0x567e98f9.
//
// Solidity: function totalStakedAmount() view returns(uint256)
func (_AssetForwarder *AssetForwarderCaller) TotalStakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "totalStakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedAmount is a free data retrieval call binding the contract method 0x567e98f9.
//
// Solidity: function totalStakedAmount() view returns(uint256)
func (_AssetForwarder *AssetForwarderSession) TotalStakedAmount() (*big.Int, error) {
	return _AssetForwarder.Contract.TotalStakedAmount(&_AssetForwarder.CallOpts)
}

// TotalStakedAmount is a free data retrieval call binding the contract method 0x567e98f9.
//
// Solidity: function totalStakedAmount() view returns(uint256)
func (_AssetForwarder *AssetForwarderCallerSession) TotalStakedAmount() (*big.Int, error) {
	return _AssetForwarder.Contract.TotalStakedAmount(&_AssetForwarder.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_AssetForwarder *AssetForwarderSession) Usdc() (common.Address, error) {
	return _AssetForwarder.Contract.Usdc(&_AssetForwarder.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) Usdc() (common.Address, error) {
	return _AssetForwarder.Contract.Usdc(&_AssetForwarder.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_AssetForwarder *AssetForwarderCaller) WrappedNativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "wrappedNativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_AssetForwarder *AssetForwarderSession) WrappedNativeToken() (common.Address, error) {
	return _AssetForwarder.Contract.WrappedNativeToken(&_AssetForwarder.CallOpts)
}

// WrappedNativeToken is a free data retrieval call binding the contract method 0x17fcb39b.
//
// Solidity: function wrappedNativeToken() view returns(address)
func (_AssetForwarder *AssetForwarderCallerSession) WrappedNativeToken() (common.Address, error) {
	return _AssetForwarder.Contract.WrappedNativeToken(&_AssetForwarder.CallOpts)
}

// CommunityPause is a paid mutator transaction binding the contract method 0x6696821b.
//
// Solidity: function communityPause() payable returns()
func (_AssetForwarder *AssetForwarderTransactor) CommunityPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "communityPause")
}

// CommunityPause is a paid mutator transaction binding the contract method 0x6696821b.
//
// Solidity: function communityPause() payable returns()
func (_AssetForwarder *AssetForwarderSession) CommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.CommunityPause(&_AssetForwarder.TransactOpts)
}

// CommunityPause is a paid mutator transaction binding the contract method 0x6696821b.
//
// Solidity: function communityPause() payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) CommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.CommunityPause(&_AssetForwarder.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.GrantRole(&_AssetForwarder.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.GrantRole(&_AssetForwarder.TransactOpts, role, account)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDeposit(opts *bind.TransactOpts, depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDeposit", depositData, destToken, recipient)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDeposit(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDeposit(&_AssetForwarder.TransactOpts, depositData, destToken, recipient)
}

// IDeposit is a paid mutator transaction binding the contract method 0xf452ed4d.
//
// Solidity: function iDeposit((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDeposit(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDeposit(&_AssetForwarder.TransactOpts, depositData, destToken, recipient)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDepositInfoUpdate(opts *bind.TransactOpts, srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDepositInfoUpdate", srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDepositInfoUpdate(srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositInfoUpdate(&_AssetForwarder.TransactOpts, srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositInfoUpdate is a paid mutator transaction binding the contract method 0xad7c17ee.
//
// Solidity: function iDepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, bool initiatewithdrawal) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDepositInfoUpdate(srcToken common.Address, feeAmount *big.Int, depositId *big.Int, initiatewithdrawal bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositInfoUpdate(&_AssetForwarder.TransactOpts, srcToken, feeAmount, depositId, initiatewithdrawal)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDepositMessage(opts *bind.TransactOpts, depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDepositMessage", depositData, destToken, recipient, message)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDepositMessage(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositMessage(&_AssetForwarder.TransactOpts, depositData, destToken, recipient, message)
}

// IDepositMessage is a paid mutator transaction binding the contract method 0x0421caf0.
//
// Solidity: function iDepositMessage((uint256,uint256,uint256,address,address,bytes32) depositData, bytes destToken, bytes recipient, bytes message) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDepositMessage(depositData IAssetForwarderDepositData, destToken []byte, recipient []byte, message []byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositMessage(&_AssetForwarder.TransactOpts, depositData, destToken, recipient, message)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IDepositUSDC(opts *bind.TransactOpts, partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iDepositUSDC", partnerId, destChainIdBytes, recipient, amount)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_AssetForwarder *AssetForwarderSession) IDepositUSDC(partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositUSDC(&_AssetForwarder.TransactOpts, partnerId, destChainIdBytes, recipient, amount)
}

// IDepositUSDC is a paid mutator transaction binding the contract method 0x3e28c7d2.
//
// Solidity: function iDepositUSDC(uint256 partnerId, bytes32 destChainIdBytes, bytes32 recipient, uint256 amount) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IDepositUSDC(partnerId *big.Int, destChainIdBytes [32]byte, recipient [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IDepositUSDC(&_AssetForwarder.TransactOpts, partnerId, destChainIdBytes, recipient, amount)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns(bytes)
func (_AssetForwarder *AssetForwarderTransactor) IReceive(opts *bind.TransactOpts, requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iReceive", requestSender, packet, arg2)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns(bytes)
func (_AssetForwarder *AssetForwarderSession) IReceive(requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IReceive(&_AssetForwarder.TransactOpts, requestSender, packet, arg2)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns(bytes)
func (_AssetForwarder *AssetForwarderTransactorSession) IReceive(requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IReceive(&_AssetForwarder.TransactOpts, requestSender, packet, arg2)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IRelay(opts *bind.TransactOpts, relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iRelay", relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_AssetForwarder *AssetForwarderSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelay(&_AssetForwarder.TransactOpts, relayData)
}

// IRelay is a paid mutator transaction binding the contract method 0x64778c1f.
//
// Solidity: function iRelay((uint256,bytes32,uint256,address,address) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IRelay(relayData IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelay(&_AssetForwarder.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IRelayMessage(opts *bind.TransactOpts, relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iRelayMessage", relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_AssetForwarder *AssetForwarderSession) IRelayMessage(relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayMessage(&_AssetForwarder.TransactOpts, relayData)
}

// IRelayMessage is a paid mutator transaction binding the contract method 0x6fb003da.
//
// Solidity: function iRelayMessage((uint256,bytes32,uint256,address,address,bytes) relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IRelayMessage(relayData IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayMessage(&_AssetForwarder.TransactOpts, relayData)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AssetForwarder *AssetForwarderTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AssetForwarder *AssetForwarderSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Multicall(&_AssetForwarder.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_AssetForwarder *AssetForwarderTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Multicall(&_AssetForwarder.TransactOpts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AssetForwarder *AssetForwarderTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AssetForwarder *AssetForwarderSession) Pause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Pause(&_AssetForwarder.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Pause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Pause(&_AssetForwarder.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RenounceRole(&_AssetForwarder.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RenounceRole(&_AssetForwarder.TransactOpts, role, account)
}

// Rescue is a paid mutator transaction binding the contract method 0x7a4e4ecf.
//
// Solidity: function rescue(address token, uint256 amount) returns()
func (_AssetForwarder *AssetForwarderTransactor) Rescue(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "rescue", token, amount)
}

// Rescue is a paid mutator transaction binding the contract method 0x7a4e4ecf.
//
// Solidity: function rescue(address token, uint256 amount) returns()
func (_AssetForwarder *AssetForwarderSession) Rescue(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Rescue(&_AssetForwarder.TransactOpts, token, amount)
}

// Rescue is a paid mutator transaction binding the contract method 0x7a4e4ecf.
//
// Solidity: function rescue(address token, uint256 amount) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Rescue(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Rescue(&_AssetForwarder.TransactOpts, token, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RevokeRole(&_AssetForwarder.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.RevokeRole(&_AssetForwarder.TransactOpts, role, account)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x0171958a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,bool)[] _destDetails) returns()
func (_AssetForwarder *AssetForwarderTransactor) SetDestDetails(opts *bind.TransactOpts, _destChainIdBytes [][32]byte, _destDetails []IAssetForwarderDestDetails) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "setDestDetails", _destChainIdBytes, _destDetails)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x0171958a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,bool)[] _destDetails) returns()
func (_AssetForwarder *AssetForwarderSession) SetDestDetails(_destChainIdBytes [][32]byte, _destDetails []IAssetForwarderDestDetails) (*types.Transaction, error) {
	return _AssetForwarder.Contract.SetDestDetails(&_AssetForwarder.TransactOpts, _destChainIdBytes, _destDetails)
}

// SetDestDetails is a paid mutator transaction binding the contract method 0x0171958a.
//
// Solidity: function setDestDetails(bytes32[] _destChainIdBytes, (uint32,uint256,bool)[] _destDetails) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) SetDestDetails(_destChainIdBytes [][32]byte, _destDetails []IAssetForwarderDestDetails) (*types.Transaction, error) {
	return _AssetForwarder.Contract.SetDestDetails(&_AssetForwarder.TransactOpts, _destChainIdBytes, _destDetails)
}

// ToggleCommunityPause is a paid mutator transaction binding the contract method 0xf627df94.
//
// Solidity: function toggleCommunityPause() returns()
func (_AssetForwarder *AssetForwarderTransactor) ToggleCommunityPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "toggleCommunityPause")
}

// ToggleCommunityPause is a paid mutator transaction binding the contract method 0xf627df94.
//
// Solidity: function toggleCommunityPause() returns()
func (_AssetForwarder *AssetForwarderSession) ToggleCommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.ToggleCommunityPause(&_AssetForwarder.TransactOpts)
}

// ToggleCommunityPause is a paid mutator transaction binding the contract method 0xf627df94.
//
// Solidity: function toggleCommunityPause() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) ToggleCommunityPause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.ToggleCommunityPause(&_AssetForwarder.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AssetForwarder *AssetForwarderTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AssetForwarder *AssetForwarderSession) Unpause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Unpause(&_AssetForwarder.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Unpause() (*types.Transaction, error) {
	return _AssetForwarder.Contract.Unpause(&_AssetForwarder.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x5ac62700.
//
// Solidity: function update(uint256 index, address _gatewayContract, bytes _routerMiddlewareBase, uint256 minPauseStakeAmount, uint256 maxPauseStakeAmount) returns()
func (_AssetForwarder *AssetForwarderTransactor) Update(opts *bind.TransactOpts, index *big.Int, _gatewayContract common.Address, _routerMiddlewareBase []byte, minPauseStakeAmount *big.Int, maxPauseStakeAmount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "update", index, _gatewayContract, _routerMiddlewareBase, minPauseStakeAmount, maxPauseStakeAmount)
}

// Update is a paid mutator transaction binding the contract method 0x5ac62700.
//
// Solidity: function update(uint256 index, address _gatewayContract, bytes _routerMiddlewareBase, uint256 minPauseStakeAmount, uint256 maxPauseStakeAmount) returns()
func (_AssetForwarder *AssetForwarderSession) Update(index *big.Int, _gatewayContract common.Address, _routerMiddlewareBase []byte, minPauseStakeAmount *big.Int, maxPauseStakeAmount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Update(&_AssetForwarder.TransactOpts, index, _gatewayContract, _routerMiddlewareBase, minPauseStakeAmount, maxPauseStakeAmount)
}

// Update is a paid mutator transaction binding the contract method 0x5ac62700.
//
// Solidity: function update(uint256 index, address _gatewayContract, bytes _routerMiddlewareBase, uint256 minPauseStakeAmount, uint256 maxPauseStakeAmount) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Update(index *big.Int, _gatewayContract common.Address, _routerMiddlewareBase []byte, minPauseStakeAmount *big.Int, maxPauseStakeAmount *big.Int) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Update(&_AssetForwarder.TransactOpts, index, _gatewayContract, _routerMiddlewareBase, minPauseStakeAmount, maxPauseStakeAmount)
}

// UpdateTokenMessenger is a paid mutator transaction binding the contract method 0x78e0f9bd.
//
// Solidity: function updateTokenMessenger(address _tokenMessenger) returns()
func (_AssetForwarder *AssetForwarderTransactor) UpdateTokenMessenger(opts *bind.TransactOpts, _tokenMessenger common.Address) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "updateTokenMessenger", _tokenMessenger)
}

// UpdateTokenMessenger is a paid mutator transaction binding the contract method 0x78e0f9bd.
//
// Solidity: function updateTokenMessenger(address _tokenMessenger) returns()
func (_AssetForwarder *AssetForwarderSession) UpdateTokenMessenger(_tokenMessenger common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.UpdateTokenMessenger(&_AssetForwarder.TransactOpts, _tokenMessenger)
}

// UpdateTokenMessenger is a paid mutator transaction binding the contract method 0x78e0f9bd.
//
// Solidity: function updateTokenMessenger(address _tokenMessenger) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) UpdateTokenMessenger(_tokenMessenger common.Address) (*types.Transaction, error) {
	return _AssetForwarder.Contract.UpdateTokenMessenger(&_AssetForwarder.TransactOpts, _tokenMessenger)
}

// WithdrawStakeAmount is a paid mutator transaction binding the contract method 0x8a27fecb.
//
// Solidity: function withdrawStakeAmount() returns()
func (_AssetForwarder *AssetForwarderTransactor) WithdrawStakeAmount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "withdrawStakeAmount")
}

// WithdrawStakeAmount is a paid mutator transaction binding the contract method 0x8a27fecb.
//
// Solidity: function withdrawStakeAmount() returns()
func (_AssetForwarder *AssetForwarderSession) WithdrawStakeAmount() (*types.Transaction, error) {
	return _AssetForwarder.Contract.WithdrawStakeAmount(&_AssetForwarder.TransactOpts)
}

// WithdrawStakeAmount is a paid mutator transaction binding the contract method 0x8a27fecb.
//
// Solidity: function withdrawStakeAmount() returns()
func (_AssetForwarder *AssetForwarderTransactorSession) WithdrawStakeAmount() (*types.Transaction, error) {
	return _AssetForwarder.Contract.WithdrawStakeAmount(&_AssetForwarder.TransactOpts)
}

// AssetForwarderCommunityPausedIterator is returned from FilterCommunityPaused and is used to iterate over the raw logs and unpacked data for CommunityPaused events raised by the AssetForwarder contract.
type AssetForwarderCommunityPausedIterator struct {
	Event *AssetForwarderCommunityPaused // Event containing the contract specifics and raw log

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
func (it *AssetForwarderCommunityPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderCommunityPaused)
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
		it.Event = new(AssetForwarderCommunityPaused)
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
func (it *AssetForwarderCommunityPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderCommunityPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderCommunityPaused represents a CommunityPaused event raised by the AssetForwarder contract.
type AssetForwarderCommunityPaused struct {
	Pauser       common.Address
	StakedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCommunityPaused is a free log retrieval operation binding the contract event 0x9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d.
//
// Solidity: event CommunityPaused(address indexed pauser, uint256 stakedAmount)
func (_AssetForwarder *AssetForwarderFilterer) FilterCommunityPaused(opts *bind.FilterOpts, pauser []common.Address) (*AssetForwarderCommunityPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "CommunityPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderCommunityPausedIterator{contract: _AssetForwarder.contract, event: "CommunityPaused", logs: logs, sub: sub}, nil
}

// WatchCommunityPaused is a free log subscription operation binding the contract event 0x9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d.
//
// Solidity: event CommunityPaused(address indexed pauser, uint256 stakedAmount)
func (_AssetForwarder *AssetForwarderFilterer) WatchCommunityPaused(opts *bind.WatchOpts, sink chan<- *AssetForwarderCommunityPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "CommunityPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderCommunityPaused)
				if err := _AssetForwarder.contract.UnpackLog(event, "CommunityPaused", log); err != nil {
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

// ParseCommunityPaused is a log parse operation binding the contract event 0x9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d.
//
// Solidity: event CommunityPaused(address indexed pauser, uint256 stakedAmount)
func (_AssetForwarder *AssetForwarderFilterer) ParseCommunityPaused(log types.Log) (*AssetForwarderCommunityPaused, error) {
	event := new(AssetForwarderCommunityPaused)
	if err := _AssetForwarder.contract.UnpackLog(event, "CommunityPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderDepositInfoUpdateIterator is returned from FilterDepositInfoUpdate and is used to iterate over the raw logs and unpacked data for DepositInfoUpdate events raised by the AssetForwarder contract.
type AssetForwarderDepositInfoUpdateIterator struct {
	Event *AssetForwarderDepositInfoUpdate // Event containing the contract specifics and raw log

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
func (it *AssetForwarderDepositInfoUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderDepositInfoUpdate)
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
		it.Event = new(AssetForwarderDepositInfoUpdate)
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
func (it *AssetForwarderDepositInfoUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderDepositInfoUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderDepositInfoUpdate represents a DepositInfoUpdate event raised by the AssetForwarder contract.
type AssetForwarderDepositInfoUpdate struct {
	SrcToken           common.Address
	FeeAmount          *big.Int
	DepositId          *big.Int
	EventNonce         *big.Int
	Initiatewithdrawal bool
	Depositor          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositInfoUpdate is a free log retrieval operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) FilterDepositInfoUpdate(opts *bind.FilterOpts) (*AssetForwarderDepositInfoUpdateIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "DepositInfoUpdate")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderDepositInfoUpdateIterator{contract: _AssetForwarder.contract, event: "DepositInfoUpdate", logs: logs, sub: sub}, nil
}

// WatchDepositInfoUpdate is a free log subscription operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) WatchDepositInfoUpdate(opts *bind.WatchOpts, sink chan<- *AssetForwarderDepositInfoUpdate) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "DepositInfoUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderDepositInfoUpdate)
				if err := _AssetForwarder.contract.UnpackLog(event, "DepositInfoUpdate", log); err != nil {
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

// ParseDepositInfoUpdate is a log parse operation binding the contract event 0x86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1.
//
// Solidity: event DepositInfoUpdate(address srcToken, uint256 feeAmount, uint256 depositId, uint256 eventNonce, bool initiatewithdrawal, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) ParseDepositInfoUpdate(log types.Log) (*AssetForwarderDepositInfoUpdate, error) {
	event := new(AssetForwarderDepositInfoUpdate)
	if err := _AssetForwarder.contract.UnpackLog(event, "DepositInfoUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsDepositedIterator is returned from FilterFundsDeposited and is used to iterate over the raw logs and unpacked data for FundsDeposited events raised by the AssetForwarder contract.
type AssetForwarderFundsDepositedIterator struct {
	Event *AssetForwarderFundsDeposited // Event containing the contract specifics and raw log

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
func (it *AssetForwarderFundsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsDeposited)
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
		it.Event = new(AssetForwarderFundsDeposited)
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
func (it *AssetForwarderFundsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsDeposited represents a FundsDeposited event raised by the AssetForwarder contract.
type AssetForwarderFundsDeposited struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	DestAmount       *big.Int
	DepositId        *big.Int
	SrcToken         common.Address
	Depositor        common.Address
	Recipient        []byte
	DestToken        []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDeposited is a free log retrieval operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsDeposited(opts *bind.FilterOpts) (*AssetForwarderFundsDepositedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsDepositedIterator{contract: _AssetForwarder.contract, event: "FundsDeposited", logs: logs, sub: sub}, nil
}

// WatchFundsDeposited is a free log subscription operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsDeposited(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsDeposited) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsDeposited)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
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

// ParseFundsDeposited is a log parse operation binding the contract event 0x6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad.
//
// Solidity: event FundsDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, address depositor, bytes recipient, bytes destToken)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsDeposited(log types.Log) (*AssetForwarderFundsDeposited, error) {
	event := new(AssetForwarderFundsDeposited)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsDepositedWithMessageIterator is returned from FilterFundsDepositedWithMessage and is used to iterate over the raw logs and unpacked data for FundsDepositedWithMessage events raised by the AssetForwarder contract.
type AssetForwarderFundsDepositedWithMessageIterator struct {
	Event *AssetForwarderFundsDepositedWithMessage // Event containing the contract specifics and raw log

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
func (it *AssetForwarderFundsDepositedWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsDepositedWithMessage)
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
		it.Event = new(AssetForwarderFundsDepositedWithMessage)
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
func (it *AssetForwarderFundsDepositedWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsDepositedWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsDepositedWithMessage represents a FundsDepositedWithMessage event raised by the AssetForwarder contract.
type AssetForwarderFundsDepositedWithMessage struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	DestAmount       *big.Int
	DepositId        *big.Int
	SrcToken         common.Address
	Recipient        []byte
	Depositor        common.Address
	DestToken        []byte
	Message          []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDepositedWithMessage is a free log retrieval operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsDepositedWithMessage(opts *bind.FilterOpts) (*AssetForwarderFundsDepositedWithMessageIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsDepositedWithMessage")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsDepositedWithMessageIterator{contract: _AssetForwarder.contract, event: "FundsDepositedWithMessage", logs: logs, sub: sub}, nil
}

// WatchFundsDepositedWithMessage is a free log subscription operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsDepositedWithMessage(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsDepositedWithMessage) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsDepositedWithMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsDepositedWithMessage)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsDepositedWithMessage", log); err != nil {
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

// ParseFundsDepositedWithMessage is a log parse operation binding the contract event 0x3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f.
//
// Solidity: event FundsDepositedWithMessage(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 destAmount, uint256 depositId, address srcToken, bytes recipient, address depositor, bytes destToken, bytes message)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsDepositedWithMessage(log types.Log) (*AssetForwarderFundsDepositedWithMessage, error) {
	event := new(AssetForwarderFundsDepositedWithMessage)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsDepositedWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsPaidIterator is returned from FilterFundsPaid and is used to iterate over the raw logs and unpacked data for FundsPaid events raised by the AssetForwarder contract.
type AssetForwarderFundsPaidIterator struct {
	Event *AssetForwarderFundsPaid // Event containing the contract specifics and raw log

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
func (it *AssetForwarderFundsPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsPaid)
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
		it.Event = new(AssetForwarderFundsPaid)
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
func (it *AssetForwarderFundsPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsPaid represents a FundsPaid event raised by the AssetForwarder contract.
type AssetForwarderFundsPaid struct {
	MessageHash [32]byte
	Forwarder   common.Address
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaid is a free log retrieval operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsPaid(opts *bind.FilterOpts) (*AssetForwarderFundsPaidIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsPaidIterator{contract: _AssetForwarder.contract, event: "FundsPaid", logs: logs, sub: sub}, nil
}

// WatchFundsPaid is a free log subscription operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsPaid(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsPaid) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsPaid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsPaid)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaid", log); err != nil {
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

// ParseFundsPaid is a log parse operation binding the contract event 0x0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464.
//
// Solidity: event FundsPaid(bytes32 messageHash, address forwarder, uint256 nonce)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsPaid(log types.Log) (*AssetForwarderFundsPaid, error) {
	event := new(AssetForwarderFundsPaid)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderFundsPaidWithMessageIterator is returned from FilterFundsPaidWithMessage and is used to iterate over the raw logs and unpacked data for FundsPaidWithMessage events raised by the AssetForwarder contract.
type AssetForwarderFundsPaidWithMessageIterator struct {
	Event *AssetForwarderFundsPaidWithMessage // Event containing the contract specifics and raw log

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
func (it *AssetForwarderFundsPaidWithMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderFundsPaidWithMessage)
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
		it.Event = new(AssetForwarderFundsPaidWithMessage)
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
func (it *AssetForwarderFundsPaidWithMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderFundsPaidWithMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderFundsPaidWithMessage represents a FundsPaidWithMessage event raised by the AssetForwarder contract.
type AssetForwarderFundsPaidWithMessage struct {
	MessageHash [32]byte
	Forwarder   common.Address
	Nonce       *big.Int
	ExecFlag    bool
	ExecData    []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFundsPaidWithMessage is a free log retrieval operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_AssetForwarder *AssetForwarderFilterer) FilterFundsPaidWithMessage(opts *bind.FilterOpts) (*AssetForwarderFundsPaidWithMessageIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "FundsPaidWithMessage")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderFundsPaidWithMessageIterator{contract: _AssetForwarder.contract, event: "FundsPaidWithMessage", logs: logs, sub: sub}, nil
}

// WatchFundsPaidWithMessage is a free log subscription operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_AssetForwarder *AssetForwarderFilterer) WatchFundsPaidWithMessage(opts *bind.WatchOpts, sink chan<- *AssetForwarderFundsPaidWithMessage) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "FundsPaidWithMessage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderFundsPaidWithMessage)
				if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaidWithMessage", log); err != nil {
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

// ParseFundsPaidWithMessage is a log parse operation binding the contract event 0x21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad.
//
// Solidity: event FundsPaidWithMessage(bytes32 messageHash, address forwarder, uint256 nonce, bool execFlag, bytes execData)
func (_AssetForwarder *AssetForwarderFilterer) ParseFundsPaidWithMessage(log types.Log) (*AssetForwarderFundsPaidWithMessage, error) {
	event := new(AssetForwarderFundsPaidWithMessage)
	if err := _AssetForwarder.contract.UnpackLog(event, "FundsPaidWithMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AssetForwarder contract.
type AssetForwarderPausedIterator struct {
	Event *AssetForwarderPaused // Event containing the contract specifics and raw log

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
func (it *AssetForwarderPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderPaused)
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
		it.Event = new(AssetForwarderPaused)
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
func (it *AssetForwarderPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderPaused represents a Paused event raised by the AssetForwarder contract.
type AssetForwarderPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AssetForwarder *AssetForwarderFilterer) FilterPaused(opts *bind.FilterOpts) (*AssetForwarderPausedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderPausedIterator{contract: _AssetForwarder.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AssetForwarder *AssetForwarderFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AssetForwarderPaused) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderPaused)
				if err := _AssetForwarder.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AssetForwarder *AssetForwarderFilterer) ParsePaused(log types.Log) (*AssetForwarderPaused, error) {
	event := new(AssetForwarderPaused)
	if err := _AssetForwarder.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AssetForwarder contract.
type AssetForwarderRoleAdminChangedIterator struct {
	Event *AssetForwarderRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AssetForwarderRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderRoleAdminChanged)
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
		it.Event = new(AssetForwarderRoleAdminChanged)
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
func (it *AssetForwarderRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderRoleAdminChanged represents a RoleAdminChanged event raised by the AssetForwarder contract.
type AssetForwarderRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetForwarder *AssetForwarderFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AssetForwarderRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderRoleAdminChangedIterator{contract: _AssetForwarder.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AssetForwarder *AssetForwarderFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AssetForwarderRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderRoleAdminChanged)
				if err := _AssetForwarder.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AssetForwarder *AssetForwarderFilterer) ParseRoleAdminChanged(log types.Log) (*AssetForwarderRoleAdminChanged, error) {
	event := new(AssetForwarderRoleAdminChanged)
	if err := _AssetForwarder.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AssetForwarder contract.
type AssetForwarderRoleGrantedIterator struct {
	Event *AssetForwarderRoleGranted // Event containing the contract specifics and raw log

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
func (it *AssetForwarderRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderRoleGranted)
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
		it.Event = new(AssetForwarderRoleGranted)
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
func (it *AssetForwarderRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderRoleGranted represents a RoleGranted event raised by the AssetForwarder contract.
type AssetForwarderRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetForwarderRoleGrantedIterator, error) {

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

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderRoleGrantedIterator{contract: _AssetForwarder.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AssetForwarderRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderRoleGranted)
				if err := _AssetForwarder.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AssetForwarder *AssetForwarderFilterer) ParseRoleGranted(log types.Log) (*AssetForwarderRoleGranted, error) {
	event := new(AssetForwarderRoleGranted)
	if err := _AssetForwarder.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AssetForwarder contract.
type AssetForwarderRoleRevokedIterator struct {
	Event *AssetForwarderRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AssetForwarderRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderRoleRevoked)
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
		it.Event = new(AssetForwarderRoleRevoked)
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
func (it *AssetForwarderRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderRoleRevoked represents a RoleRevoked event raised by the AssetForwarder contract.
type AssetForwarderRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AssetForwarderRoleRevokedIterator, error) {

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

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AssetForwarderRoleRevokedIterator{contract: _AssetForwarder.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AssetForwarder *AssetForwarderFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AssetForwarderRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderRoleRevoked)
				if err := _AssetForwarder.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AssetForwarder *AssetForwarderFilterer) ParseRoleRevoked(log types.Log) (*AssetForwarderRoleRevoked, error) {
	event := new(AssetForwarderRoleRevoked)
	if err := _AssetForwarder.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AssetForwarder contract.
type AssetForwarderUnpausedIterator struct {
	Event *AssetForwarderUnpaused // Event containing the contract specifics and raw log

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
func (it *AssetForwarderUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderUnpaused)
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
		it.Event = new(AssetForwarderUnpaused)
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
func (it *AssetForwarderUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderUnpaused represents a Unpaused event raised by the AssetForwarder contract.
type AssetForwarderUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AssetForwarder *AssetForwarderFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AssetForwarderUnpausedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderUnpausedIterator{contract: _AssetForwarder.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AssetForwarder *AssetForwarderFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AssetForwarderUnpaused) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderUnpaused)
				if err := _AssetForwarder.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AssetForwarder *AssetForwarderFilterer) ParseUnpaused(log types.Log) (*AssetForwarderUnpaused, error) {
	event := new(AssetForwarderUnpaused)
	if err := _AssetForwarder.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetForwarderIUSDCDepositedIterator is returned from FilterIUSDCDeposited and is used to iterate over the raw logs and unpacked data for IUSDCDeposited events raised by the AssetForwarder contract.
type AssetForwarderIUSDCDepositedIterator struct {
	Event *AssetForwarderIUSDCDeposited // Event containing the contract specifics and raw log

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
func (it *AssetForwarderIUSDCDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetForwarderIUSDCDeposited)
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
		it.Event = new(AssetForwarderIUSDCDeposited)
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
func (it *AssetForwarderIUSDCDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetForwarderIUSDCDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetForwarderIUSDCDeposited represents a IUSDCDeposited event raised by the AssetForwarder contract.
type AssetForwarderIUSDCDeposited struct {
	PartnerId        *big.Int
	Amount           *big.Int
	DestChainIdBytes [32]byte
	UsdcNonce        *big.Int
	SrcToken         common.Address
	Recipient        [32]byte
	Depositor        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIUSDCDeposited is a free log retrieval operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) FilterIUSDCDeposited(opts *bind.FilterOpts) (*AssetForwarderIUSDCDepositedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderIUSDCDepositedIterator{contract: _AssetForwarder.contract, event: "iUSDCDeposited", logs: logs, sub: sub}, nil
}

// WatchIUSDCDeposited is a free log subscription operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) WatchIUSDCDeposited(opts *bind.WatchOpts, sink chan<- *AssetForwarderIUSDCDeposited) (event.Subscription, error) {

	logs, sub, err := _AssetForwarder.contract.WatchLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetForwarderIUSDCDeposited)
				if err := _AssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
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

// ParseIUSDCDeposited is a log parse operation binding the contract event 0x297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d9.
//
// Solidity: event iUSDCDeposited(uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor)
func (_AssetForwarder *AssetForwarderFilterer) ParseIUSDCDeposited(log types.Log) (*AssetForwarderIUSDCDeposited, error) {
	event := new(AssetForwarderIUSDCDeposited)
	if err := _AssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
