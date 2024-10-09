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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_minGasThreshhold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_depositNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGateway\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRefundData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageExcecutionFailedWithLowGas\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"CommunityPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"DepositInfoUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"}],\"name\":\"FundsDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"FundsDepositedWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"FundsPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"execFlag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"execData\",\"type\":\"bytes\"}],\"name\":\"FundsPaidWithMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseType\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseType\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdcNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"iUSDCDeposited\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_THRESHHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESOURCE_SETTER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"communityPause\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositPause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executeRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"iDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"initiatewithdrawal\",\"type\":\"bool\"}],\"name\":\"iDepositInfoUpdate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"partnerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destChainIdBytes\",\"type\":\"bytes32\"}],\"internalType\":\"structIAssetForwarder.DepositData\",\"name\":\"depositData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"destToken\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"iDepositMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"requestSender\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"packet\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"iReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIAssetForwarder.RelayData[]\",\"name\":\"relayData\",\"type\":\"tuple[]\"}],\"name\":\"iRelayBatched\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage\",\"name\":\"relayData\",\"type\":\"tuple\"}],\"name\":\"iRelayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcChainId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"internalType\":\"structIAssetForwarder.RelayDataMessage[]\",\"name\":\"relayData\",\"type\":\"tuple[]\"}],\"name\":\"iRelayMessageBatched\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isCommunityPauseEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_depositPause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_relayPause\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStakeAmountMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayPause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerMiddlewareBase\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleCommunityPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_depositUnpause\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_relayUnpause\",\"type\":\"bool\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gatewayContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_routerMiddlewareBase\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minPauseStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPauseStakeAmount\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600a805462ffffff191660011790553480156200002057600080fd5b5060405162005012380380620050128339810160408190526200004391620001a2565b6001805460ff191681179055600380546001600160a01b0319166001600160a01b03861617905582516020840120600255600682905562000086600033620000eb565b620000b27f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07833620000eb565b620000de7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c33620000eb565b60045550620002a8915050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1662000188576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620001473390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b634e487b7160e01b600052604160045260246000fd5b60008060008060808587031215620001b957600080fd5b84516001600160a01b0381168114620001d157600080fd5b602086810151919550906001600160401b0380821115620001f157600080fd5b818801915088601f8301126200020657600080fd5b8151818111156200021b576200021b6200018c565b604051601f8201601f19908116603f011681019083821181831017156200024657620002466200018c565b816040528281528b868487010111156200025f57600080fd5b600093505b8284101562000283578484018601518185018701529285019262000264565b6000928101909501919091525050506040860151606090960151949790965092505050565b614d5a80620002b86000396000f3fe60806040526004361061026a5760003560e01c806391d1485411610153578063db618e2a116100cb578063f215f1481161007f578063f627df9411610064578063f627df94146106b3578063fa7d593d146106c8578063fd5ad37c146106e857600080fd5b8063f215f1481461068a578063f452ed4d146106a057600080fd5b8063ddeb5094116100b0578063ddeb509414610602578063de35f5cb14610622578063eb0cde1d1461063857600080fd5b8063db618e2a146105ab578063ddd224f1146105df57600080fd5b8063b19941a911610122578063cc2205ea11610107578063cc2205ea14610544578063d547741f14610557578063d9dc86941461057757600080fd5b8063b19941a914610518578063c44e947e1461052e57600080fd5b806391d1485414610472578063a217fddf146104c3578063ac9650d8146104d8578063ad7c17ee1461050557600080fd5b8063567e98f9116101e65780636e22558d116101b55780637a4e4ecf1161019a5780637a4e4ecf1461041d5780637d0da5621461043d5780638a27fecb1461045d57600080fd5b80636e22558d146103eb5780636fb003da1461040a57600080fd5b8063567e98f91461039a5780635ac62700146103b057806364778c1f146103d05780636696821b146103e357600080fd5b8063248a9ca31161023d57806336568abe1161022257806336568abe1461034a5780634463182f1461036a5780634b7b94761461038457600080fd5b8063248a9ca3146102ec5780632f2ff15d1461032a57600080fd5b806301ffc9a71461026f5780630421caf0146102a457806311a4c778146102b95780631aa6485a146102cc575b600080fd5b34801561027b57600080fd5b5061028f61028a366004613d7a565b610718565b60405190151581526020015b60405180910390f35b6102b76102b2366004613f8d565b6107b1565b005b6102b76102c73660046140d9565b610a05565b3480156102d857600080fd5b506102b76102e73660046141d3565b611188565b3480156102f857600080fd5b5061031c610307366004614268565b60009081526020819052604090206001015490565b60405190815260200161029b565b34801561033657600080fd5b506102b7610345366004614281565b6113f3565b34801561035657600080fd5b506102b7610365366004614281565b61141d565b34801561037657600080fd5b50600a5461028f9060ff1681565b34801561039057600080fd5b5061031c60065481565b3480156103a657600080fd5b5061031c60095481565b3480156103bc57600080fd5b506102b76103cb3660046142b1565b6114d0565b6102b76103de3660046143a2565b611642565b6102b7611a06565b3480156103f757600080fd5b50600a5461028f90610100900460ff1681565b6102b76104183660046143be565b611ca3565b34801561042957600080fd5b506102b76104383660046143f3565b6121c3565b34801561044957600080fd5b506102b761045836600461442d565b6123b6565b34801561046957600080fd5b506102b76123ea565b34801561047e57600080fd5b5061028f61048d366004614281565b60009182526020828152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b3480156104cf57600080fd5b5061031c600081565b3480156104e457600080fd5b506104f86104f336600461445b565b6124c6565b60405161029b919061453e565b6102b76105133660046145be565b612638565b34801561052457600080fd5b5061031c60075481565b34801561053a57600080fd5b5061031c60025481565b6102b7610552366004614608565b6128fa565b34801561056357600080fd5b506102b7610572366004614281565b612d9a565b34801561058357600080fd5b5061031c7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c81565b3480156105b757600080fd5b5061031c7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb07881565b3480156105eb57600080fd5b5061031c6ec097ce7bc90715b34b9f100000000081565b34801561060e57600080fd5b506102b761061d36600461442d565b612dbf565b34801561062e57600080fd5b5061031c60045481565b34801561064457600080fd5b506003546106659073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161029b565b34801561069657600080fd5b5061031c60085481565b6102b76106ae3660046146a2565b612df3565b3480156106bf57600080fd5b506102b761303b565b3480156106d457600080fd5b50600a5461028f9062010000900460ff1681565b3480156106f457600080fd5b5061028f610703366004614268565b60056020526000908152604090205460ff1681565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b0000000000000000000000000000000000000000000000000000000014806107ab57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6107b9613079565b6107e6600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a54610100900460ff161561085d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4465706f7369743a20706175736564000000000000000000000000000000000060448201526064015b60405180910390fd5b6ec097ce7bc90715b34b9f1000000000846020015111156108aa576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108e1846060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b156109285734846020015114610923576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61095d565b61095d33308660200151876060015173ffffffffffffffffffffffffffffffffffffffff166130e7909392919063ffffffff16565b7f3dbc28a2fa93575c89d951d683c45ddb951a2ecf6bc9b9704a61589fa0fcb70f846000015185602001518660a0015187604001516004600081546109a190614747565b918290555060608a015160808b01516040516109c8979695949392918b918d908c9061477f565b60405180910390a16109ff600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b50505050565b610a0d613079565b610a3a600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a5462010000900460ff1615610aad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f52656c61793a20706175736564000000000000000000000000000000000000006044820152606401610854565b8051600090815b818110156110ed576ec097ce7bc90715b34b9f1000000000848281518110610ade57610ade614811565b6020026020010151600001511115610b22576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848281518110610b3657610b36614811565b602002602001015160000151858381518110610b5457610b54614811565b602002602001015160200151868481518110610b7257610b72614811565b602002602001015160400151878581518110610b9057610b90614811565b602002602001015160600151888681518110610bae57610bae614811565b602002602001015160800151308a8881518110610bcd57610bcd614811565b602002602001015160a00151604051602001610bef9796959493929190614840565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915060ff1615610c4357506110db565b600081815260056020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558451610ccd90869084908110610c9257610c92614811565b60200260200101516060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15610d605734858381518110610ce557610ce5614811565b60200260200101516000015185610cfc91906148a0565b1115610d34576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110610d4657610d46614811565b60200260200101516000015184610d5d91906148a0565b93505b610d75858381518110610c9257610c92614811565b15610e2c576000858381518110610d8e57610d8e614811565b60200260200101516080015173ffffffffffffffffffffffffffffffffffffffff16868481518110610dc257610dc2614811565b602090810291909101015151604051600081818185875af1925050503d8060008114610e0a576040519150601f19603f3d011682016040523d82523d6000602084013e610e0f565b606091505b5090915050600181151514610e2657610e266148b3565b50610eb0565b610eb033868481518110610e4257610e42614811565b602002602001015160800151878581518110610e6057610e60614811565b602002602001015160000151888681518110610e7e57610e7e614811565b60200260200101516060015173ffffffffffffffffffffffffffffffffffffffff166130e7909392919063ffffffff16565b60606000610edb878581518110610ec957610ec9614811565b6020026020010151608001513b151590565b8015610f0557506000878581518110610ef657610ef6614811565b602002602001015160a0015151115b1561108557868481518110610f1c57610f1c614811565b60200260200101516080015173ffffffffffffffffffffffffffffffffffffffff165a63d00a2d5f60e01b898781518110610f5957610f59614811565b6020026020010151606001518a8881518110610f7757610f77614811565b6020026020010151600001518b8981518110610f9557610f95614811565b602002602001015160a00151604051602401610fb3939291906148e2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317909252905161103c9190614920565b60006040518083038160008787f1925050503d806000811461107a576040519150601f19603f3d011682016040523d82523d6000602084013e61107f565b606091505b50925090505b7f21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad83336004600081546110b790614747565b91829055506040516110cf939291908690889061493c565b60405180910390a15050505b806110e581614747565b915050610ab4565b508134111561115457600033611103843461497f565b604051600081818185875af1925050503d806000811461113f576040519150601f19603f3d011682016040523d82523d6000602084013e611144565b606091505b505090508061115257600080fd5b505b5050611185600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b50565b60035473ffffffffffffffffffffffffffffffffffffffff1633146111d9576040517ffc9dfe8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84846040516111e9929190614992565b60405180910390206002541461122b576040517f23dfe1fb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806000858060200190518101906112449190614a08565b9250925092506000825190508151811461128a576040517f94809d2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b818110156113e7576112e18482815181106112aa576112aa614811565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b61134757611342858483815181106112fb576112fb614811565b602002602001015186848151811061131557611315614811565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166131c39092919063ffffffff16565b6113d5565b60008573ffffffffffffffffffffffffffffffffffffffff1684838151811061137257611372614811565b602002602001015160405160006040518083038185875af1925050503d80600081146113ba576040519150601f19603f3d011682016040523d82523d6000602084013e6113bf565b606091505b50909150506001811515146113d357600080fd5b505b806113df81614747565b91505061128d565b50505050505050505050565b60008281526020819052604090206001015461140e81613219565b6114188383613223565b505050565b73ffffffffffffffffffffffffffffffffffffffff811633146114c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c6600000000000000000000000000000000006064820152608401610854565b6114cc8282613313565b5050565b7f8b9e7a9f25b0aca3f51c01b8fee30790fb16f4d4deded8385ae6643d054bb0786114fa81613219565b8660010361154757600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8816179055611639565b8660020361157057848460405161155f929190614992565b604051908190039020600255611639565b86600303611639578183111561162e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604560248201527f6d696e50617573655374616b65416d6f756e74206d757374206265206c65737360448201527f207468616e206f7220657175616c20746f206d617850617573655374616b654160648201527f6d6f756e74000000000000000000000000000000000000000000000000000000608482015260a401610854565b600783905560088290555b50505050505050565b61164a613079565b611677600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a5462010000900460ff16156116ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f52656c61793a20706175736564000000000000000000000000000000000000006044820152606401610854565b80516ec097ce7bc90715b34b9f10000000001015611734576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516020808301516040808501516060808701516080808901518551978801989098529386019490945284015273ffffffffffffffffffffffffffffffffffffffff9182169083015290911660a08201523060c082015260009060e001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915060ff1615611811576040517f7448c64c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260056020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905560608201516118809073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b1561193b57815134146118bf576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151825160405160009273ffffffffffffffffffffffffffffffffffffffff1691908381818185875af1925050503d806000811461191c576040519150601f19603f3d011682016040523d82523d6000602084013e611921565b606091505b509091505060018115151461193557600080fd5b5061196c565b60808201518251606084015161196c9273ffffffffffffffffffffffffffffffffffffffff909116913391906130e7565b7f0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a8004464813360046000815461199e90614747565b91829055506040805193845273ffffffffffffffffffffffffffffffffffffffff90921660208401529082015260600160405180910390a150611185600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b600a54610100900460ff1615611a78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4465706f7369743a2070617573656400000000000000000000000000000000006044820152606401610854565b600a5462010000900460ff1615611aeb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f52656c61793a20706175736564000000000000000000000000000000000000006044820152606401610854565b600a5460ff16611b57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f436f6d6d756e6974792070617573652069732064697361626c656400000000006044820152606401610854565b60075415801590611b69575060085415155b611bcf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f536574205374616b6520416d6f756e742052616e6765000000000000000000006044820152606401610854565b6007543410158015611be357506008543411155b611c49576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f5374616b6520616d6f756e74206f7574206f662072616e6765000000000000006044820152606401610854565b600034600954611c5991906148a0565b60098190559050611c6b6001806133ca565b60405134815233907f9593b43c20e09177a4170170ac564123ad8138e040e21eec96d1ae9db9ee5d6d9060200160405180910390a250565b611cab613079565b611cd8600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a5462010000900460ff1615611d4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f52656c61793a20706175736564000000000000000000000000000000000000006044820152606401610854565b80516ec097ce7bc90715b34b9f10000000001015611d95576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80516020808301516040808501516060860151608087015160a08801519351600097611dc8979096953092909101614840565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915060ff1615611e48576040517f7448c64c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260056020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556060820151611eb79073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15611f725781513414611ef6576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6080820151825160405160009273ffffffffffffffffffffffffffffffffffffffff1691908381818185875af1925050503d8060008114611f53576040519150601f19603f3d011682016040523d82523d6000602084013e611f58565b606091505b5090915050600181151514611f6c57600080fd5b50611fa3565b608082015182516060840151611fa39273ffffffffffffffffffffffffffffffffffffffff909116913391906130e7565b60606000611fb584608001513b151590565b8015611fc6575060008460a0015151115b1561213f57836080015173ffffffffffffffffffffffffffffffffffffffff165a6060860151865160a08801516040517fd00a2d5f000000000000000000000000000000000000000000000000000000009361202893909290916024016148e2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516120b19190614920565b60006040518083038160008787f1925050503d80600081146120ef576040519150601f19603f3d011682016040523d82523d6000602084013e6120f4565b606091505b50925090508015801561210857506006545a105b1561213f576040517f9e82ca7900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f21937deaa62558dad619c8d730a7d1d7ef41731fc194c32973511e1455cb37ad833360046000815461217190614747565b9182905550604051612189939291908690889061493c565b60405180910390a1505050611185600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b60006121ce81613219565b6121d6613079565b612203600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff84160361229557604051600090339084908381818185875af1925050503d8060008114612276576040519150601f19603f3d011682016040523d82523d6000602084013e61227b565b606091505b509091505060018115151461228f57600080fd5b50612387565b6040513360248201526044810183905273ffffffffffffffffffffffffffffffffffffffff841690606401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052516123419190614920565b6000604051808303816000865af19150503d806000811461237e576040519150601f19603f3d011682016040523d82523d6000602084013e612383565b606091505b5050505b611418600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c6123e081613219565b611418838361358c565b60006123f581613219565b600954471015612461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f496e73756666696369656e742066756e647300000000000000000000000000006044820152606401610854565b60098054600091829055604051909190339083908381818185875af1925050503d80600081146124ad576040519150601f19603f3d011682016040523d82523d6000602084013e6124b2565b606091505b509091505060018115151461141857600080fd5b60608167ffffffffffffffff8111156124e1576124e1613dbc565b60405190808252806020026020018201604052801561251457816020015b60608152602001906001900390816124ff5790505b50905060005b82811015612631576000803086868581811061253857612538614811565b905060200281019061254a9190614ad7565b604051612558929190614992565b600060405180830381855af49150503d8060008114612593576040519150601f19603f3d011682016040523d82523d6000602084013e612598565b606091505b5091509150816125fe576044815110156125b157600080fd5b600481019050808060200190518101906125cb9190614b3c565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108549190614baa565b8084848151811061261157612611614811565b60200260200101819052505050808061262990614747565b91505061251a565b5092915050565b612640613079565b61266d600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a54610100900460ff16156126df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4465706f7369743a2070617573656400000000000000000000000000000000006044820152606401610854565b80156127755734156126f057600080fd5b7f86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b18460008460046000815461272490614747565b91829055506040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152600160808201523360a082015260c00160405180910390a16128cb565b6ec097ce7bc90715b34b9f10000000008311156127be576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff85160361282d57348314612828576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61284f565b61284f73ffffffffffffffffffffffffffffffffffffffff85163330866130e7565b7f86896302632bf6dc8a3ac0ae7ddf17d5a5d5c1ca1aad37b4b920a587c51135b184848460046000815461288290614747565b91829055506040805173ffffffffffffffffffffffffffffffffffffffff90951685526020850193909352918301526060820152600060808201523360a082015260c0016109c8565b6109ff600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b612902613079565b61292f600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a5462010000900460ff16156129a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f52656c61793a20706175736564000000000000000000000000000000000000006044820152606401610854565b8051600090815b818110156110ed576ec097ce7bc90715b34b9f10000000008482815181106129d3576129d3614811565b6020026020010151600001511115612a17576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000848281518110612a2b57612a2b614811565b602002602001015160000151858381518110612a4957612a49614811565b602002602001015160200151868481518110612a6757612a67614811565b602002602001015160400151878581518110612a8557612a85614811565b602002602001015160600151888681518110612aa357612aa3614811565b60200260200101516080015130604051602001612b04969594939291909586526020860194909452604085019290925273ffffffffffffffffffffffffffffffffffffffff908116606085015290811660808401521660a082015260c00190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600590935291205490915060ff1615612b585750612d88565b600081815260056020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558451612ba790869084908110610c9257610c92614811565b15612c3a5734858381518110612bbf57612bbf614811565b60200260200101516000015185612bd691906148a0565b1115612c0e576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b848281518110612c2057612c20614811565b60200260200101516000015184612c3791906148a0565b93505b612c4f858381518110610c9257610c92614811565b15612d06576000858381518110612c6857612c68614811565b60200260200101516080015173ffffffffffffffffffffffffffffffffffffffff16868481518110612c9c57612c9c614811565b602090810291909101015151604051600081818185875af1925050503d8060008114612ce4576040519150601f19603f3d011682016040523d82523d6000602084013e612ce9565b606091505b5090915050600181151514612d0057612d006148b3565b50612d1c565b612d1c33868481518110610e4257610e42614811565b7f0f3ca0b27903ec13ef88a7ea8be837cc19b0d7f71a735f2083215739a80044648133600460008154612d4e90614747565b91829055506040805193845273ffffffffffffffffffffffffffffffffffffffff90921660208401529082015260600160405180910390a1505b80612d9281614747565b9150506129a9565b600082815260208190526040902060010154612db581613219565b6114188383613313565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c612de981613219565b61141883836133ca565b612dfb613079565b612e28600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b600a54610100900460ff1615612e9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4465706f7369743a2070617573656400000000000000000000000000000000006044820152606401610854565b6ec097ce7bc90715b34b9f100000000083602001511115612ee7576040517f0625040100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f1e836060015173ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1490565b15612f655734836020015114612f60576040517f2c5211c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f9a565b612f9a33308560200151866060015173ffffffffffffffffffffffffffffffffffffffff166130e7909392919063ffffffff16565b7f6f223106c8e3df857d691613d18d1478cc7c629a1fdf16c7b461d36729fcc7ad836000015184602001518560a001518660400151600460008154612fde90614747565b9182905550606089015160808a015160405161300497969594939291908a908c90614bbd565b60405180910390a1611418600180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681179055565b600061304681613219565b50600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00811660ff90911615179055565b60015460ff166130e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610854565b565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526109ff9085907f23b872dd00000000000000000000000000000000000000000000000000000000906084015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261373b565b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526114189084907fa9059cbb0000000000000000000000000000000000000000000000000000000090606401613141565b611185813361384a565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166114cc5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556132b53390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16156114cc5760008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b81156134a857600a54610100900460ff1615613442576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4465706f736974732061726520616c72656164792070617573656400000000006044820152606401610854565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010017905560408051338152600060208201527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d910160405180910390a15b80156114cc57600a5462010000900460ff1615613521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4465706f736974732061726520616c72656164792070617573656400000000006044820152606401610854565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff166201000017905560408051338152600160208201527fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d91015b60405180910390a15050565b811561366557600a54610100900460ff16613603576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4465706f736974732061726520616c726561647920756e7061757365640000006044820152606401610854565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560408051338152600060208201527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c910160405180910390a15b80156114cc57600a5462010000900460ff166136dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4465706f736974732061726520616c726561647920756e7061757365640000006044820152606401610854565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffff16905560408051338152600160208201527f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9101613580565b600061379d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166139029092919063ffffffff16565b90508051600014806137be5750808060200190518101906137be9190614c3a565b611418576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610854565b60008281526020818152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff166114cc5761388881613919565b613893836020613938565b6040516020016138a4929190614c57565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261085491600401614baa565b60606139118484600085613b82565b949350505050565b60606107ab73ffffffffffffffffffffffffffffffffffffffff831660145b60606000613947836002614cd8565b6139529060026148a0565b67ffffffffffffffff81111561396a5761396a613dbc565b6040519080825280601f01601f191660200182016040528015613994576020820181803683370190505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106139cb576139cb614811565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110613a2e57613a2e614811565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000613a6a846002614cd8565b613a759060016148a0565b90505b6001811115613b12577f303132333435363738396162636465660000000000000000000000000000000085600f1660108110613ab657613ab6614811565b1a60f81b828281518110613acc57613acc614811565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c93613b0b81614cef565b9050613a78565b508315613b7b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610854565b9392505050565b606082471015613c14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610854565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051613c3d9190614920565b60006040518083038185875af1925050503d8060008114613c7a576040519150601f19603f3d011682016040523d82523d6000602084013e613c7f565b606091505b5091509150613c9087838387613c9b565b979650505050505050565b60608315613d31578251600003613d2a5773ffffffffffffffffffffffffffffffffffffffff85163b613d2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610854565b5081613911565b6139118383815115613d465781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108549190614baa565b600060208284031215613d8c57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114613b7b57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff81118282101715613e0e57613e0e613dbc565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e5b57613e5b613dbc565b604052919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461118557600080fd5b600060c08284031215613e9757600080fd5b613e9f613deb565b90508135815260208201356020820152604082013560408201526060820135613ec781613e63565b60608201526080820135613eda81613e63565b8060808301525060a082013560a082015292915050565b600067ffffffffffffffff821115613f0b57613f0b613dbc565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112613f4857600080fd5b8135613f5b613f5682613ef1565b613e14565b818152846020838601011115613f7057600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806101208587031215613fa457600080fd5b613fae8686613e85565b935060c085013567ffffffffffffffff80821115613fcb57600080fd5b613fd788838901613f37565b945060e0870135915080821115613fed57600080fd5b613ff988838901613f37565b935061010087013591508082111561401057600080fd5b5061401d87828801613f37565b91505092959194509250565b600067ffffffffffffffff82111561404357614043613dbc565b5060051b60200190565b600060c0828403121561405f57600080fd5b614067613deb565b9050813581526020820135602082015260408201356040820152606082013561408f81613e63565b606082015260808201356140a281613e63565b608082015260a082013567ffffffffffffffff8111156140c157600080fd5b6140cd84828501613f37565b60a08301525092915050565b600060208083850312156140ec57600080fd5b823567ffffffffffffffff8082111561410457600080fd5b818501915085601f83011261411857600080fd5b8135614126613f5682614029565b81815260059190911b8301840190848101908883111561414557600080fd5b8585015b8381101561417d578035858111156141615760008081fd5b61416f8b89838a010161404d565b845250918601918601614149565b5098975050505050505050565b60008083601f84011261419c57600080fd5b50813567ffffffffffffffff8111156141b457600080fd5b6020830191508360208285010111156141cc57600080fd5b9250929050565b6000806000806000606086880312156141eb57600080fd5b853567ffffffffffffffff8082111561420357600080fd5b61420f89838a0161418a565b9097509550602088013591508082111561422857600080fd5b61423489838a01613f37565b9450604088013591508082111561424a57600080fd5b506142578882890161418a565b969995985093965092949392505050565b60006020828403121561427a57600080fd5b5035919050565b6000806040838503121561429457600080fd5b8235915060208301356142a681613e63565b809150509250929050565b60008060008060008060a087890312156142ca57600080fd5b8635955060208701356142dc81613e63565b9450604087013567ffffffffffffffff8111156142f857600080fd5b61430489828a0161418a565b979a9699509760608101359660809091013595509350505050565b600060a0828403121561433157600080fd5b60405160a0810181811067ffffffffffffffff8211171561435457614354613dbc565b8060405250809150823581526020830135602082015260408301356040820152606083013561438281613e63565b6060820152608083013561439581613e63565b6080919091015292915050565b600060a082840312156143b457600080fd5b613b7b838361431f565b6000602082840312156143d057600080fd5b813567ffffffffffffffff8111156143e757600080fd5b6139118482850161404d565b6000806040838503121561440657600080fd5b823561441181613e63565b946020939093013593505050565b801515811461118557600080fd5b6000806040838503121561444057600080fd5b823561444b8161441f565b915060208301356142a68161441f565b6000806020838503121561446e57600080fd5b823567ffffffffffffffff8082111561448657600080fd5b818501915085601f83011261449a57600080fd5b8135818111156144a957600080fd5b8660208260051b85010111156144be57600080fd5b60209290920196919550909350505050565b60005b838110156144eb5781810151838201526020016144d3565b50506000910152565b6000815180845261450c8160208601602086016144d0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156145b1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261459f8583516144f4565b94509285019290850190600101614565565b5092979650505050505050565b600080600080608085870312156145d457600080fd5b84356145df81613e63565b9350602085013592506040850135915060608501356145fd8161441f565b939692955090935050565b6000602080838503121561461b57600080fd5b823567ffffffffffffffff81111561463257600080fd5b8301601f8101851361464357600080fd5b8035614651613f5682614029565b81815260a0918202830184019184820191908884111561467057600080fd5b938501935b8385101561469657614687898661431f565b83529384019391850191614675565b50979650505050505050565b600080600061010084860312156146b857600080fd5b6146c28585613e85565b925060c084013567ffffffffffffffff808211156146df57600080fd5b6146eb87838801613f37565b935060e086013591508082111561470157600080fd5b5061470e86828701613f37565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361477857614778614718565b5060010190565b60006101408c83528b60208401528a604084015289606084015288608084015273ffffffffffffffffffffffffffffffffffffffff80891660a08501528160c08501526147ce828501896144f4565b90871660e085015283810361010085015290506147eb81866144f4565b905082810361012084015261480081856144f4565b9d9c50505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b878152866020820152856040820152600073ffffffffffffffffffffffffffffffffffffffff8087166060840152808616608084015280851660a08401525060e060c083015261489360e08301846144f4565b9998505050505050505050565b808201808211156107ab576107ab614718565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff8416815282602082015260606040820152600061491760608301846144f4565b95945050505050565b600082516149328184602087016144d0565b9190910192915050565b85815273ffffffffffffffffffffffffffffffffffffffff85166020820152836040820152821515606082015260a060808201526000613c9060a08301846144f4565b818103818111156107ab576107ab614718565b8183823760009101908152919050565b600082601f8301126149b357600080fd5b815160206149c3613f5683614029565b82815260059290921b840181019181810190868411156149e257600080fd5b8286015b848110156149fd57805183529183019183016149e6565b509695505050505050565b600080600060608486031215614a1d57600080fd5b8351614a2881613e63565b8093505060208085015167ffffffffffffffff80821115614a4857600080fd5b818701915087601f830112614a5c57600080fd5b8151614a6a613f5682614029565b81815260059190911b8301840190848101908a831115614a8957600080fd5b938501935b82851015614ab0578451614aa181613e63565b82529385019390850190614a8e565b60408a01519097509450505080831115614ac957600080fd5b505061470e868287016149a2565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614b0c57600080fd5b83018035915067ffffffffffffffff821115614b2757600080fd5b6020019150368190038213156141cc57600080fd5b600060208284031215614b4e57600080fd5b815167ffffffffffffffff811115614b6557600080fd5b8201601f81018413614b7657600080fd5b8051614b84613f5682613ef1565b818152856020838501011115614b9957600080fd5b6149178260208301602086016144d0565b602081526000613b7b60208301846144f4565b60006101208b83528a602084015289604084015288606084015287608084015273ffffffffffffffffffffffffffffffffffffffff80881660a085015280871660c0850152508060e0840152614c15818401866144f4565b9050828103610100840152614c2a81856144f4565b9c9b505050505050505050505050565b600060208284031215614c4c57600080fd5b8151613b7b8161441f565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351614c8f8160178501602088016144d0565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351614ccc8160288401602088016144d0565b01602801949350505050565b80820281158282048414176107ab576107ab614718565b600081614cfe57614cfe614718565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea2646970667358221220e80edf992dcd3c9099bd1a4f508922bd339a44d16a8d0b73ad4fb43f9d152e6064736f6c63430008140033",
}

// AssetForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use AssetForwarderMetaData.ABI instead.
var AssetForwarderABI = AssetForwarderMetaData.ABI

// AssetForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AssetForwarderMetaData.Bin instead.
var AssetForwarderBin = AssetForwarderMetaData.Bin

// DeployAssetForwarder deploys a new Ethereum contract, binding an instance of AssetForwarder to it.
func DeployAssetForwarder(auth *bind.TransactOpts, backend bind.ContractBackend, _gatewayContract common.Address, _routerMiddlewareBase []byte, _minGasThreshhold *big.Int, _depositNonce *big.Int) (common.Address, *types.Transaction, *AssetForwarder, error) {
	parsed, err := AssetForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AssetForwarderBin), backend, _gatewayContract, _routerMiddlewareBase, _minGasThreshhold, _depositNonce)
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

// DepositPause is a free data retrieval call binding the contract method 0x6e22558d.
//
// Solidity: function depositPause() view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) DepositPause(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "depositPause")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositPause is a free data retrieval call binding the contract method 0x6e22558d.
//
// Solidity: function depositPause() view returns(bool)
func (_AssetForwarder *AssetForwarderSession) DepositPause() (bool, error) {
	return _AssetForwarder.Contract.DepositPause(&_AssetForwarder.CallOpts)
}

// DepositPause is a free data retrieval call binding the contract method 0x6e22558d.
//
// Solidity: function depositPause() view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) DepositPause() (bool, error) {
	return _AssetForwarder.Contract.DepositPause(&_AssetForwarder.CallOpts)
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

// RelayPause is a free data retrieval call binding the contract method 0xfa7d593d.
//
// Solidity: function relayPause() view returns(bool)
func (_AssetForwarder *AssetForwarderCaller) RelayPause(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AssetForwarder.contract.Call(opts, &out, "relayPause")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RelayPause is a free data retrieval call binding the contract method 0xfa7d593d.
//
// Solidity: function relayPause() view returns(bool)
func (_AssetForwarder *AssetForwarderSession) RelayPause() (bool, error) {
	return _AssetForwarder.Contract.RelayPause(&_AssetForwarder.CallOpts)
}

// RelayPause is a free data retrieval call binding the contract method 0xfa7d593d.
//
// Solidity: function relayPause() view returns(bool)
func (_AssetForwarder *AssetForwarderCallerSession) RelayPause() (bool, error) {
	return _AssetForwarder.Contract.RelayPause(&_AssetForwarder.CallOpts)
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

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns()
func (_AssetForwarder *AssetForwarderTransactor) IReceive(opts *bind.TransactOpts, requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iReceive", requestSender, packet, arg2)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns()
func (_AssetForwarder *AssetForwarderSession) IReceive(requestSender string, packet []byte, arg2 string) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IReceive(&_AssetForwarder.TransactOpts, requestSender, packet, arg2)
}

// IReceive is a paid mutator transaction binding the contract method 0x1aa6485a.
//
// Solidity: function iReceive(string requestSender, bytes packet, string ) returns()
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

// IRelayBatched is a paid mutator transaction binding the contract method 0xcc2205ea.
//
// Solidity: function iRelayBatched((uint256,bytes32,uint256,address,address)[] relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactor) IRelayBatched(opts *bind.TransactOpts, relayData []IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "iRelayBatched", relayData)
}

// IRelayBatched is a paid mutator transaction binding the contract method 0xcc2205ea.
//
// Solidity: function iRelayBatched((uint256,bytes32,uint256,address,address)[] relayData) payable returns()
func (_AssetForwarder *AssetForwarderSession) IRelayBatched(relayData []IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayBatched(&_AssetForwarder.TransactOpts, relayData)
}

// IRelayBatched is a paid mutator transaction binding the contract method 0xcc2205ea.
//
// Solidity: function iRelayBatched((uint256,bytes32,uint256,address,address)[] relayData) payable returns()
func (_AssetForwarder *AssetForwarderTransactorSession) IRelayBatched(relayData []IAssetForwarderRelayData) (*types.Transaction, error) {
	return _AssetForwarder.Contract.IRelayBatched(&_AssetForwarder.TransactOpts, relayData)
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

// Pause is a paid mutator transaction binding the contract method 0xddeb5094.
//
// Solidity: function pause(bool _depositPause, bool _relayPause) returns()
func (_AssetForwarder *AssetForwarderTransactor) Pause(opts *bind.TransactOpts, _depositPause bool, _relayPause bool) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "pause", _depositPause, _relayPause)
}

// Pause is a paid mutator transaction binding the contract method 0xddeb5094.
//
// Solidity: function pause(bool _depositPause, bool _relayPause) returns()
func (_AssetForwarder *AssetForwarderSession) Pause(_depositPause bool, _relayPause bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Pause(&_AssetForwarder.TransactOpts, _depositPause, _relayPause)
}

// Pause is a paid mutator transaction binding the contract method 0xddeb5094.
//
// Solidity: function pause(bool _depositPause, bool _relayPause) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Pause(_depositPause bool, _relayPause bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Pause(&_AssetForwarder.TransactOpts, _depositPause, _relayPause)
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

// Unpause is a paid mutator transaction binding the contract method 0x7d0da562.
//
// Solidity: function unpause(bool _depositUnpause, bool _relayUnpause) returns()
func (_AssetForwarder *AssetForwarderTransactor) Unpause(opts *bind.TransactOpts, _depositUnpause bool, _relayUnpause bool) (*types.Transaction, error) {
	return _AssetForwarder.contract.Transact(opts, "unpause", _depositUnpause, _relayUnpause)
}

// Unpause is a paid mutator transaction binding the contract method 0x7d0da562.
//
// Solidity: function unpause(bool _depositUnpause, bool _relayUnpause) returns()
func (_AssetForwarder *AssetForwarderSession) Unpause(_depositUnpause bool, _relayUnpause bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Unpause(&_AssetForwarder.TransactOpts, _depositUnpause, _relayUnpause)
}

// Unpause is a paid mutator transaction binding the contract method 0x7d0da562.
//
// Solidity: function unpause(bool _depositUnpause, bool _relayUnpause) returns()
func (_AssetForwarder *AssetForwarderTransactorSession) Unpause(_depositUnpause bool, _relayUnpause bool) (*types.Transaction, error) {
	return _AssetForwarder.Contract.Unpause(&_AssetForwarder.TransactOpts, _depositUnpause, _relayUnpause)
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
	Account   common.Address
	PauseType *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address account, uint256 pauseType)
func (_AssetForwarder *AssetForwarderFilterer) FilterPaused(opts *bind.FilterOpts) (*AssetForwarderPausedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderPausedIterator{contract: _AssetForwarder.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address account, uint256 pauseType)
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address account, uint256 pauseType)
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
	Account   common.Address
	PauseType *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address account, uint256 pauseType)
func (_AssetForwarder *AssetForwarderFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AssetForwarderUnpausedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderUnpausedIterator{contract: _AssetForwarder.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address account, uint256 pauseType)
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address account, uint256 pauseType)
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
func (_AssetForwarder *AssetForwarderFilterer) FilterIUSDCDeposited(opts *bind.FilterOpts) (*AssetForwarderIUSDCDepositedIterator, error) {

	logs, sub, err := _AssetForwarder.contract.FilterLogs(opts, "iUSDCDeposited")
	if err != nil {
		return nil, err
	}
	return &AssetForwarderIUSDCDepositedIterator{contract: _AssetForwarder.contract, event: "iUSDCDeposited", logs: logs, sub: sub}, nil
}

// WatchIUSDCDeposited is a free log subscription operation binding the contract event 0xa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a526.
//
// Solidity: event iUSDCDeposited(uint8 txType, uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor, bytes data)
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

// ParseIUSDCDeposited is a log parse operation binding the contract event 0xa7315553e8495e03079ca4879c06cf2c7c07e2e99f024ec9537c541bdea5a526.
//
// Solidity: event iUSDCDeposited(uint8 txType, uint256 partnerId, uint256 amount, bytes32 destChainIdBytes, uint256 usdcNonce, address srcToken, bytes32 recipient, address depositor, bytes data)
func (_AssetForwarder *AssetForwarderFilterer) ParseIUSDCDeposited(log types.Log) (*AssetForwarderIUSDCDeposited, error) {
	event := new(AssetForwarderIUSDCDeposited)
	if err := _AssetForwarder.contract.UnpackLog(event, "iUSDCDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
