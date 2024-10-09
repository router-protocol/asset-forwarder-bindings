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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedNativeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_minGasThreshhold\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRefundData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageExcecutionFailedWithLowGas\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"CommunityPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"DepositInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDepositedWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"FundsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"execFlag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"execData\",\"type\":\"bytes\"}],\"name\":\"FundsPaidWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_THRESHHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESOURCE_SETTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"communityPause\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"destDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executeRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"iDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"}],\"name\":\"iDepositInfoUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"iDepositMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"iDepositUSDC\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"requestSender\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"packet\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"iReceive\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage[]\",\"name\":\"relayData\",\"type\":\"tuple[]\"}],\"name\":\"iRelayMessageBatched\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCommunityPauseEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerMiddlewareBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_destChainIdBytes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"internalType\":\"structIAssetForwarder.DestDetails[]\",\"name\":\"_destDetails\",\"type\":\"tuple[]\"}],\"name\":\"setDestDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleCommunityPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenMessenger\",\"outputs\":[{\"internalType\":\"contractITokenMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minPauseStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPauseStakeAmount\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenMessenger\",\"type\":\"address\"}],\"name\":\"updateTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedNativeToken\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600d805460ff191660011790553480156200001e57600080fd5b506040516200522b3803806200522b8339810160408190526200004191620001e3565b6001805461ffff1916811790556001600160a01b03868116608052600580546001600160a01b03199081168684161790915560038054821688841617905560048054909116918616919091179055815160208301206002556009819055620000ab6000336200010f565b620000d77f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb078336200010f565b620001037f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c336200010f565b5050505050506200030b565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16620001ac576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556200016b3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b80516001600160a01b0381168114620001c857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060c08789031215620001fd57600080fd5b6200020887620001b0565b9550602062000219818901620001b0565b95506200022960408901620001b0565b94506200023960608901620001b0565b60808901519094506001600160401b03808211156200025757600080fd5b818a0191508a601f8301126200026c57600080fd5b815181811115620002815762000281620001cd565b604051601f8201601f19908116603f01168101908382118183101715620002ac57620002ac620001cd565b816040528281528d86848701011115620002c557600080fd5b600093505b82841015620002e95784840186015181850187015292850192620002ca565b600086848301015280975050505050505060a087015190509295509295509295565b608051614eda620003516000396000818161036501528181610ba601528181610c3a01528181612d0c01528181612d8d01528181612faa015261303e0152614eda6000f3fe6080604052600436106102d15760003560e01c806378e0f9bd11610179578063c44e947e116100d6578063de35f5cb1161008a578063f452ed4d11610064578063f452ed4d14610834578063f627df9414610847578063fd5ad37c1461085c57600080fd5b8063de35f5cb146107db578063eb0cde1d146107f1578063f215f1481461081e57600080fd5b8063d9dc8694116100bb578063d9dc869414610750578063db618e2a14610784578063ddd224f1146107b857600080fd5b8063c44e947e1461071a578063d547741f1461073057600080fd5b8063981a8a021161012d578063ac9650d811610112578063ac9650d8146106c4578063ad7c17ee146106f1578063b19941a91461070457600080fd5b8063981a8a0214610645578063a217fddf146106af57600080fd5b80638456cb591161015e5780638456cb59146105ca5780638a27fecb146105df57806391d14854146105f457600080fd5b806378e0f9bd1461058a5780637a4e4ecf146105aa57600080fd5b80633e413bee11610232578063567e98f9116101e657806364778c1f116101c057806364778c1f1461055c5780636696821b1461056f5780636fb003da1461057757600080fd5b8063567e98f9146105095780635ac627001461051f5780635c975abb1461053f57600080fd5b80634463182f116102175780634463182f146104ac57806346117830146104c65780634b7b9476146104f357600080fd5b80633e413bee1461046a5780633f4ba83a1461049757600080fd5b80631aa6485a116102895780632f2ff15d1161026e5780632f2ff15d1461041757806336568abe146104375780633e28c7d21461045757600080fd5b80631aa6485a146103ac578063248a9ca3146103d957600080fd5b80630421caf0116102ba5780630421caf01461032d57806311a4c7781461034057806317fcb39b1461035357600080fd5b80630171958a146102d657806301ffc9a7146102f8575b600080fd5b3480156102e257600080fd5b506102f66102f1366004613f75565b61088c565b005b34801561030457600080fd5b5061031861031336600461402e565b610a08565b60405190151581526020015b60405180910390f35b6102f661033b366004614195565b610aa1565b6102f661034e3660046142bd565b610d3e565b34801561035f57600080fd5b506103877f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610324565b3480156103b857600080fd5b506103cc6103c73660046143b7565b611456565b60405161032491906144ba565b3480156103e557600080fd5b506104096103f43660046144cd565b60009081526020819052604090206001015490565b604051908152602001610324565b34801561042357600080fd5b506102f66104323660046144e6565b6116d9565b34801561044357600080fd5b506102f66104523660046144e6565b611703565b6102f6610465366004614516565b6117b6565b34801561047657600080fd5b506004546103879073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104a357600080fd5b506102f6611b0d565b3480156104b857600080fd5b50600d546103189060ff1681565b3480156104d257600080fd5b506005546103879073ffffffffffffffffffffffffffffffffffffffff1681565b3480156104ff57600080fd5b5061040960095481565b34801561051557600080fd5b50610409600c5481565b34801561052b57600080fd5b506102f661053a366004614548565b611b47565b34801561054b57600080fd5b50600154610100900460ff16610318565b6102f661056a3660046145b6565b611cb9565b6102f6612015565b6102f6610585366004614635565b6121d2565b34801561059657600080fd5b506102f66105a536600461466a565b61268a565b3480156105b657600080fd5b506102f66105c5366004614687565b6126fc565b3480156105d657600080fd5b506102f66128f6565b3480156105eb57600080fd5b506102f6612930565b34801561060057600080fd5b5061031861060f3660046144e6565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b34801561065157600080fd5b5061068b6106603660046144cd565b60076020526000908152604090208054600182015460029092015463ffffffff909116919060ff1683565b6040805163ffffffff90941684526020840192909252151590820152606001610324565b3480156106bb57600080fd5b50610409600081565b3480156106d057600080fd5b506106e46106df3660046146b3565b612a0f565b6040516103249190614728565b6102f66106ff3660046147a8565b612b81565b34801561071057600080fd5b50610409600a5481565b34801561072657600080fd5b5061040960025481565b34801561073c57600080fd5b506102f661074b3660046144e6565b612e80565b34801561075c57600080fd5b506104097f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c81565b34801561079057600080fd5b506104097f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07881565b3480156107c457600080fd5b506104096ec097ce7bc90715b34b9f100000000081565b3480156107e757600080fd5b5061040960065481565b3480156107fd57600080fd5b506003546103879073ffffffffffffffffffffffffffffffffffffffff1681565b34801561082a57600080fd5b50610409600b5481565b6102f66108423660046147f2565b612ea5565b34801561085357600080fd5b506102f6613141565b34801561086857600080fd5b506103186108773660046144cd565b60086020526000908152604090205460ff1681565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb0786108b68161317f565b8151835114610926576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f696e76616c6964206c656e67746800000000000000000000000000000000000060448201526064015b60405180910390fd5b60005b8251811015610a025782818151811061094457610944614868565b60200260200101516007600086848151811061096257610962614868565b6020908102919091018101518252818101929092526040908101600020835181547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff9091161781559183015160018301559190910151600290910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055806109fa816148c6565b915050610929565b50505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b000000000000000000000000000000000000000000000000000000001480610a9b57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610aa9613189565b610ad6600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b610ade6131f7565b6ec097ce7bc90715b34b9f100000000084602001511115610b2b576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b62846060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15610c675734846020015114610ba4576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015610c0c57600080fd5b505af1158015610c20573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016606087015250610c9c9050565b610c9c33308660200151876060015173ffffffffffffffffffffffffffffffffffffffff16613269909392919063ffffffff16565b7f3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f846000015185602001518660a001518760400151600660008154610ce0906148c6565b918290555060608a015160808b0151604051610d07979695949392918b918d908c906148fe565b60405180910390a1610a02600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b610d46613189565b610d73600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b610d7b6131f7565b8051600090815b818110156113bb576ec097ce7bc90715b34b9f1000000000848281518110610dac57610dac614868565b6020026020010151600001511115610df0576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848281518110610e0457610e04614868565b602002602001015160000151858381518110610e2257610e22614868565b602002602001015160200151868481518110610e4057610e40614868565b602002602001015160400151878581518110610e5e57610e5e614868565b602002602001015160600151888681518110610e7c57610e7c614868565b602002602001015160800151308a8881518110610e9b57610e9b614868565b602002602001015160a00151604051602001610ebd9796959493929190614990565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600890935291205490915060ff1615610f1157506113a9565b600081815260086020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558451610f9b90869084908110610f6057610f60614868565b60200260200101516060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b1561102e5734858381518110610fb357610fb3614868565b60200260200101516000015185610fca91906149f0565b1115611002576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84828151811061101457611014614868565b6020026020010151600001518461102b91906149f0565b93505b611043858381518110610f6057610f60614868565b156110fa57600085838151811061105c5761105c614868565b60200260200101516080015173ffffffffffffffffffffffffffffffffffffffff1686848151811061109057611090614868565b602090810291909101015151604051600081818185875af1925050503d80600081146110d8576040519150601f19603f3d011682016040523d82523d6000602084013e6110dd565b606091505b50909150506001811515146110f4576110f4614a03565b5061117e565b61117e3386848151811061111057611110614868565b60200260200101516080015187858151811061112e5761112e614868565b60200260200101516000015188868151811061114c5761114c614868565b60200260200101516060015173ffffffffffffffffffffffffffffffffffffffff16613269909392919063ffffffff16565b606060006111a987858151811061119757611197614868565b6020026020010151608001513b151590565b80156111d3575060008785815181106111c4576111c4614868565b602002602001015160a0015151115b15611353578684815181106111ea576111ea614868565b60200260200101516080015173ffffffffffffffffffffffffffffffffffffffff165a63d00a2d5f60e01b89878151811061122757611227614868565b6020026020010151606001518a888151811061124557611245614868565b6020026020010151600001518b898151811061126357611263614868565b602002602001015160a0015160405160240161128193929190614a32565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252905161130a9190614a70565b60006040518083038160008787f1925050503d8060008114611348576040519150601f19603f3d011682016040523d82523d6000602084013e61134d565b606091505b50925090505b7f21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad8333600660008154611385906148c6565b918290555060405161139d9392919086908890614a8c565b60405180910390a15050505b806113b3816148c6565b915050610d82565b5081341115611422576000336113d18434614acf565b604051600081818185875af1925050503d806000811461140d576040519150601f19603f3d011682016040523d82523d6000602084013e611412565b606091505b505090508061142057600080fd5b505b5050611453600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b50565b60035460609073ffffffffffffffffffffffffffffffffffffffff1633146114aa576040517ffc9dfe8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b85856040516114ba929190614ae2565b6040518091039020600254146114fc576040517f23dfe1fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806000868060200190518101906115159190614b58565b9250925092506000825190508151811461155b576040517f94809d2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b818110156116bb576115b284828151811061157b5761157b614868565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b61161857611613858483815181106115cc576115cc614868565b60200260200101518684815181106115e6576115e6614868565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166133459092919063ffffffff16565b6116a9565b60008573ffffffffffffffffffffffffffffffffffffffff1684838151811061164357611643614868565b602002602001015160405160006040518083038185875af1925050503d806000811461168b576040519150601f19603f3d011682016040523d82523d6000602084013e611690565b606091505b50909150506001811515146116a7576116a7614a03565b505b806116b3816148c6565b91505061155e565b50506040805160208101909152600081529998505050505050505050565b6000828152602081905260409020600101546116f48161317f565b6116fe838361339b565b505050565b73ffffffffffffffffffffffffffffffffffffffff811633146117a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c660000000000000000000000000000000000606482015260840161091d565b6117b2828261348b565b5050565b6117be613189565b6117eb600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6117f36131f7565b60008381526007602052604090206002015460ff16801561182b575060045473ffffffffffffffffffffffffffffffffffffffff1615155b6118b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f75736463206e6f7420737570706f7274656420656974686572206f6e2073726360448201527f206f6e2064737420636861696e00000000000000000000000000000000000000606482015260840161091d565b6000838152600760205260409020600101543414611901576040517f58d620b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6ec097ce7bc90715b34b9f100000000081111561194a576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60045461196f9073ffffffffffffffffffffffffffffffffffffffff16333084613269565b6005546004546119999173ffffffffffffffffffffffffffffffffffffffff918216911683613542565b600554600084815260076020526040808220546004805492517f6fd3504e00000000000000000000000000000000000000000000000000000000815290810186905263ffffffff90911660248201526044810186905273ffffffffffffffffffffffffffffffffffffffff918216606482015291921690636fd3504e906084016020604051808303816000875af1158015611a38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a5c9190614c27565b600454604080518881526020810186905290810187905267ffffffffffffffff8316606082015273ffffffffffffffffffffffffffffffffffffffff909116608082015260a081018590523360c08201529091507f297a8bc8b87367a63661d6429dbab51be5cefd71ce6a3050fa900a8f276d66d99060e00160405180910390a150610a02600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c611b378161317f565b611b3f61363b565b6114536136ac565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb078611b718161317f565b86600103611bbe57600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8816179055611cb0565b86600203611be7578484604051611bd6929190614ae2565b604051908190039020600255611cb0565b86600303611cb05781831115611ca5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604560248201527f6d696e50617573655374616b65416d6f756e74206d757374206265206c65737360448201527f207468616e206f7220657175616c20746f206d617850617573655374616b654160648201527f6d6f756e74000000000000000000000000000000000000000000000000000000608482015260a40161091d565b600a839055600b8290555b50505050505050565b611cc1613189565b611cee600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b611cf66131f7565b80516ec097ce7bc90715b34b9f10000000001015611d40576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516020808301516040808501516060808701516080808901518551978801989098529386019490945284015273ffffffffffffffffffffffffffffffffffffffff9182169083015290911660a08201523060c082015260009060e001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600890935291205490915060ff1615611e1d576040517f7448c64c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260086020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556060820151611e8c9073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15611f4a5781513414611ecb576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151825160405160009273ffffffffffffffffffffffffffffffffffffffff1691908381818185875af1925050503d8060008114611f28576040519150601f19603f3d011682016040523d82523d6000602084013e611f2d565b606091505b5090915050600181151514611f4457611f44614a03565b50611f7b565b608082015182516060840151611f7b9273ffffffffffffffffffffffffffffffffffffffff90911691339190613269565b7f0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a80044648133600660008154611fad906148c6565b91829055506040805193845273ffffffffffffffffffffffffffffffffffffffff90921660208401529082015260600160405180910390a150611453600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b61201d6131f7565b600d5460ff16612089576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f436f6d6d756e6974792070617573652069732064697361626c65640000000000604482015260640161091d565b600a541580159061209b5750600b5415155b612101576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f536574205374616b6520416d6f756e742052616e676500000000000000000000604482015260640161091d565b600a5434101580156121155750600b543411155b61217b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f5374616b6520616d6f756e74206f7574206f662072616e676500000000000000604482015260640161091d565b600034600c5461218b91906149f0565b600c819055905061219a613729565b60405134815233907f9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d9060200160405180910390a250565b6121da613189565b612207600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b61220f6131f7565b80516ec097ce7bc90715b34b9f10000000001015612259576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516020808301516040808501516060860151608087015160a0880151935160009761228c979096953092909101614990565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600890935291205490915060ff161561230c576040517f7448c64c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260086020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055606082015161237b9073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b1561243957815134146123ba576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151825160405160009273ffffffffffffffffffffffffffffffffffffffff1691908381818185875af1925050503d8060008114612417576040519150601f19603f3d011682016040523d82523d6000602084013e61241c565b606091505b509091505060018115151461243357612433614a03565b5061246a565b60808201518251606084015161246a9273ffffffffffffffffffffffffffffffffffffffff90911691339190613269565b6060600061247c84608001513b151590565b801561248d575060008460a0015151115b1561260657836080015173ffffffffffffffffffffffffffffffffffffffff165a6060860151865160a08801516040517fd00a2d5f00000000000000000000000000000000000000000000000000000000936124ef9390929091602401614a32565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516125789190614a70565b60006040518083038160008787f1925050503d80600081146125b6576040519150601f19603f3d011682016040523d82523d6000602084013e6125bb565b606091505b5092509050801580156125cf57506009545a105b15612606576040517f9e82ca7900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad8333600660008154612638906148c6565b91829055506040516126509392919086908890614a8c565b60405180910390a1505050611453600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb0786126b48161317f565b50600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006127078161317f565b61270f613189565b61273c600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff8416036127d157604051600090339084908381818185875af1925050503d80600081146127af576040519150601f19603f3d011682016040523d82523d6000602084013e6127b4565b606091505b50909150506001811515146127cb576127cb614a03565b506128c7565b6040513360248201526044810183905260009073ffffffffffffffffffffffffffffffffffffffff851690606401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052516128809190614a70565b6000604051808303816000865af19150503d80600081146128bd576040519150601f19603f3d011682016040523d82523d6000602084013e6128c2565b606091505b505050505b6116fe600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c6129208161317f565b6129286131f7565b611453613729565b600061293b8161317f565b600c544710156129a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e73756666696369656e742066756e64730000000000000000000000000000604482015260640161091d565b600c8054600091829055604051909190339083908381818185875af1925050503d80600081146129f3576040519150601f19603f3d011682016040523d82523d6000602084013e6129f8565b606091505b50909150506001811515146116fe576116fe614a03565b60608167ffffffffffffffff811115612a2a57612a2a613dc4565b604051908082528060200260200182016040528015612a5d57816020015b6060815260200190600190039081612a485790505b50905060005b82811015612b7a5760008030868685818110612a8157612a81614868565b9050602002810190612a939190614c51565b604051612aa1929190614ae2565b600060405180830381855af49150503d8060008114612adc576040519150601f19603f3d011682016040523d82523d6000602084013e612ae1565b606091505b509150915081612b4757604481511015612afa57600080fd5b60048101905080806020019051810190612b149190614cb6565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091d91906144ba565b80848481518110612b5a57612b5a614868565b602002602001018190525050508080612b72906148c6565b915050612a63565b5092915050565b612b89613189565b612bb6600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b612bbe6131f7565b8015612c57573415612bd257612bd2614a03565b7f86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b184600084600660008154612c06906148c6565b91829055506040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152600160808201523360a082015260c00160405180910390a1612e51565b6ec097ce7bc90715b34b9f1000000000831115612ca0576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff851603612db357348314612d0a576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015612d7257600080fd5b505af1158015612d86573d6000803e3d6000fd5b50505050507f00000000000000000000000000000000000000000000000000000000000000009350612dd5565b612dd573ffffffffffffffffffffffffffffffffffffffff8516333086613269565b7f86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b1848484600660008154612e08906148c6565b91829055506040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152600060808201523360a082015260c001610d07565b610a02600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b600082815260208190526040902060010154612e9b8161317f565b6116fe838361348b565b612ead613189565b612eda600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b612ee26131f7565b6ec097ce7bc90715b34b9f100000000083602001511115612f2f576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f66836060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b1561306b5734836020015114612fa8576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561301057600080fd5b505af1158015613024573d6000803e3d6000fd5b50505073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166060860152506130a09050565b6130a033308560200151866060015173ffffffffffffffffffffffffffffffffffffffff16613269909392919063ffffffff16565b7f6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad836000015184602001518560a0015186604001516006600081546130e4906148c6565b9182905550606089015160808a015160405161310a97969594939291908a908c90614d24565b60405180910390a16116fe600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b600061314c8161317f565b50600d80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00811660ff90911615179055565b6114538133613785565b60015460ff166131f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161091d565b565b600154610100900460ff16156131f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015260640161091d565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610a029085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261383d565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526116fe9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064016132c3565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166117b25760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905561342d3390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16156117b25760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff83811660248301526000919085169063dd62ed3e90604401602060405180830381865afa1580156135b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135dc9190614da1565b9050610a02847f095ea7b3000000000000000000000000000000000000000000000000000000008561360e86866149f0565b60405173ffffffffffffffffffffffffffffffffffffffff909216602483015260448201526064016132c3565b600154610100900460ff166131f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161091d565b6136b461363b565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6137316131f7565b600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136ff3390565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166117b2576137c38161394c565b6137ce83602061396b565b6040516020016137df929190614dba565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261091d916004016144ba565b600061389f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16613bb59092919063ffffffff16565b90508051600014806138c05750808060200190518101906138c09190614e3b565b6116fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161091d565b6060610a9b73ffffffffffffffffffffffffffffffffffffffff831660145b6060600061397a836002614e58565b6139859060026149f0565b67ffffffffffffffff81111561399d5761399d613dc4565b6040519080825280601f01601f1916602001820160405280156139c7576020820181803683370190505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106139fe576139fe614868565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110613a6157613a61614868565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000613a9d846002614e58565b613aa89060016149f0565b90505b6001811115613b45577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110613ae957613ae9614868565b1a60f81b828281518110613aff57613aff614868565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93613b3e81614e6f565b9050613aab565b508315613bae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161091d565b9392505050565b6060613bc48484600085613bcc565b949350505050565b606082471015613c5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161091d565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051613c879190614a70565b60006040518083038185875af1925050503d8060008114613cc4576040519150601f19603f3d011682016040523d82523d6000602084013e613cc9565b606091505b5091509150613cda87838387613ce5565b979650505050505050565b60608315613d7b578251600003613d745773ffffffffffffffffffffffffffffffffffffffff85163b613d74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161091d565b5081613bc4565b613bc48383815115613d905781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091d91906144ba565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613e1657613e16613dc4565b60405290565b60405160c0810167ffffffffffffffff81118282101715613e1657613e16613dc4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e8657613e86613dc4565b604052919050565b600067ffffffffffffffff821115613ea857613ea8613dc4565b5060051b60200190565b801515811461145357600080fd5b600082601f830112613ed157600080fd5b81356020613ee6613ee183613e8e565b613e3f565b82815260609283028501820192828201919087851115613f0557600080fd5b8387015b85811015613f685781818a031215613f215760008081fd5b613f29613df3565b813563ffffffff81168114613f3e5760008081fd5b81528186013586820152604080830135613f5781613eb2565b908201528452928401928101613f09565b5090979650505050505050565b60008060408385031215613f8857600080fd5b823567ffffffffffffffff80821115613fa057600080fd5b818501915085601f830112613fb457600080fd5b81356020613fc4613ee183613e8e565b82815260059290921b84018101918181019089841115613fe357600080fd5b948201945b8386101561400157853582529482019490820190613fe8565b9650508601359250508082111561401757600080fd5b5061402485828601613ec0565b9150509250929050565b60006020828403121561404057600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114613bae57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461145357600080fd5b600060c082840312156140a457600080fd5b6140ac613e1c565b905081358152602082013560208201526040820135604082015260608201356140d481614070565b606082015260808201356140e781614070565b8060808301525060a082013560a082015292915050565b600067ffffffffffffffff82111561411857614118613dc4565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261415557600080fd5b8135614163613ee1826140fe565b81815284602083860101111561417857600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008061012085870312156141ac57600080fd5b6141b68686614092565b935060c085013567ffffffffffffffff808211156141d357600080fd5b6141df88838901614144565b945060e08701359150808211156141f557600080fd5b61420188838901614144565b935061010087013591508082111561421857600080fd5b5061422587828801614144565b91505092959194509250565b600060c0828403121561424357600080fd5b61424b613e1c565b9050813581526020820135602082015260408201356040820152606082013561427381614070565b6060820152608082013561428681614070565b608082015260a082013567ffffffffffffffff8111156142a557600080fd5b6142b184828501614144565b60a08301525092915050565b600060208083850312156142d057600080fd5b823567ffffffffffffffff808211156142e857600080fd5b818501915085601f8301126142fc57600080fd5b813561430a613ee182613e8e565b81815260059190911b8301840190848101908883111561432957600080fd5b8585015b83811015614361578035858111156143455760008081fd5b6143538b89838a0101614231565b84525091860191860161432d565b5098975050505050505050565b60008083601f84011261438057600080fd5b50813567ffffffffffffffff81111561439857600080fd5b6020830191508360208285010111156143b057600080fd5b9250929050565b6000806000806000606086880312156143cf57600080fd5b853567ffffffffffffffff808211156143e757600080fd5b6143f389838a0161436e565b9097509550602088013591508082111561440c57600080fd5b61441889838a01614144565b9450604088013591508082111561442e57600080fd5b5061443b8882890161436e565b969995985093965092949392505050565b60005b8381101561446757818101518382015260200161444f565b50506000910152565b6000815180845261448881602086016020860161444c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000613bae6020830184614470565b6000602082840312156144df57600080fd5b5035919050565b600080604083850312156144f957600080fd5b82359150602083013561450b81614070565b809150509250929050565b6000806000806080858703121561452c57600080fd5b5050823594602084013594506040840135936060013592509050565b60008060008060008060a0878903121561456157600080fd5b86359550602087013561457381614070565b9450604087013567ffffffffffffffff81111561458f57600080fd5b61459b89828a0161436e565b979a9699509760608101359660809091013595509350505050565b600060a082840312156145c857600080fd5b60405160a0810181811067ffffffffffffffff821117156145eb576145eb613dc4565b8060405250823581526020830135602082015260408301356040820152606083013561461681614070565b6060820152608083013561462981614070565b60808201529392505050565b60006020828403121561464757600080fd5b813567ffffffffffffffff81111561465e57600080fd5b613bc484828501614231565b60006020828403121561467c57600080fd5b8135613bae81614070565b6000806040838503121561469a57600080fd5b82356146a581614070565b946020939093013593505050565b600080602083850312156146c657600080fd5b823567ffffffffffffffff808211156146de57600080fd5b818501915085601f8301126146f257600080fd5b81358181111561470157600080fd5b8660208260051b850101111561471657600080fd5b60209290920196919550909350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561479b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0888603018452614789858351614470565b9450928501929085019060010161474f565b5092979650505050505050565b600080600080608085870312156147be57600080fd5b84356147c981614070565b9350602085013592506040850135915060608501356147e781613eb2565b939692955090935050565b6000806000610100848603121561480857600080fd5b6148128585614092565b925060c084013567ffffffffffffffff8082111561482f57600080fd5b61483b87838801614144565b935060e086013591508082111561485157600080fd5b5061485e86828701614144565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036148f7576148f7614897565b5060010190565b60006101408c83528b60208401528a604084015289606084015288608084015273ffffffffffffffffffffffffffffffffffffffff80891660a08501528160c085015261494d82850189614470565b90871660e0850152838103610100850152905061496a8186614470565b905082810361012084015261497f8185614470565b9d9c50505050505050505050505050565b878152866020820152856040820152600073ffffffffffffffffffffffffffffffffffffffff8087166060840152808616608084015280851660a08401525060e060c08301526149e360e0830184614470565b9998505050505050505050565b80820180821115610a9b57610a9b614897565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000614a676060830184614470565b95945050505050565b60008251614a8281846020870161444c565b9190910192915050565b85815273ffffffffffffffffffffffffffffffffffffffff85166020820152836040820152821515606082015260a060808201526000613cda60a0830184614470565b81810381811115610a9b57610a9b614897565b8183823760009101908152919050565b600082601f830112614b0357600080fd5b81516020614b13613ee183613e8e565b82815260059290921b84018101918181019086841115614b3257600080fd5b8286015b84811015614b4d5780518352918301918301614b36565b509695505050505050565b600080600060608486031215614b6d57600080fd5b8351614b7881614070565b8093505060208085015167ffffffffffffffff80821115614b9857600080fd5b818701915087601f830112614bac57600080fd5b8151614bba613ee182613e8e565b81815260059190911b8301840190848101908a831115614bd957600080fd5b938501935b82851015614c00578451614bf181614070565b82529385019390850190614bde565b60408a01519097509450505080831115614c1957600080fd5b505061485e86828701614af2565b600060208284031215614c3957600080fd5b815167ffffffffffffffff81168114613bae57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614c8657600080fd5b83018035915067ffffffffffffffff821115614ca157600080fd5b6020019150368190038213156143b057600080fd5b600060208284031215614cc857600080fd5b815167ffffffffffffffff811115614cdf57600080fd5b8201601f81018413614cf057600080fd5b8051614cfe613ee1826140fe565b818152856020838501011115614d1357600080fd5b614a6782602083016020860161444c565b60006101208b83528a602084015289604084015288606084015287608084015273ffffffffffffffffffffffffffffffffffffffff80881660a085015280871660c0850152508060e0840152614d7c81840186614470565b9050828103610100840152614d918185614470565b9c9b505050505050505050505050565b600060208284031215614db357600080fd5b5051919050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351614df281601785016020880161444c565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351614e2f81602884016020880161444c565b01602801949350505050565b600060208284031215614e4d57600080fd5b8151613bae81613eb2565b8082028115828204841417610a9b57610a9b614897565b600081614e7e57614e7e614897565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea264697066735822122067ad1429b803df7e954937e6b439bf278c2cca18e3a40bdda00727b96963509064736f6c63430008140033",
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

// IRelayMessageBatched is a paid mutator transaction binding the contract method 0x11a4c778.
//
// Solidity: function iRelayMessageBatched((uint256,bytes32,uint256,address,address,bytes)[] relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IRelayMessageBatched(opts *bind.TransactOpts, relayData []IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iRelayMessageBatched", relayData)
}

// IRelayMessageBatched is a paid mutator transaction binding the contract method 0x11a4c778.
//
// Solidity: function iRelayMessageBatched((uint256,bytes32,uint256,address,address,bytes)[] relayData) payable returns()
func (_AssetForwarder *AssetForwarderSession) IRelayMessageBatched(relayData []IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayMessageBatched(&_AssetForwarder.TransactOpts, relayData)
}

// IRelayMessageBatched is a paid mutator transaction binding the contract method 0x11a4c778.
//
// Solidity: function iRelayMessageBatched((uint256,bytes32,uint256,address,address,bytes)[] relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IRelayMessageBatched(relayData []IAssetForwarderRelayDataMessage) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayMessageBatched(&_AssetForwarder.TransactOpts, relayData)
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
