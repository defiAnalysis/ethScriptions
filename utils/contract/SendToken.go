// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SendTokenMetaData contains all meta data concerning the SendToken contract.
var SendTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"sendToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SendTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use SendTokenMetaData.ABI instead.
var SendTokenABI = SendTokenMetaData.ABI

// SendToken is an auto generated Go binding around an Ethereum contract.
type SendToken struct {
	SendTokenCaller     // Read-only binding to the contract
	SendTokenTransactor // Write-only binding to the contract
	SendTokenFilterer   // Log filterer for contract events
}

// SendTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SendTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SendTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SendTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SendTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SendTokenSession struct {
	Contract     *SendToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SendTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SendTokenCallerSession struct {
	Contract *SendTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SendTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SendTokenTransactorSession struct {
	Contract     *SendTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SendTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SendTokenRaw struct {
	Contract *SendToken // Generic contract binding to access the raw methods on
}

// SendTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SendTokenCallerRaw struct {
	Contract *SendTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SendTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SendTokenTransactorRaw struct {
	Contract *SendTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSendToken creates a new instance of SendToken, bound to a specific deployed contract.
func NewSendToken(address common.Address, backend bind.ContractBackend) (*SendToken, error) {
	contract, err := bindSendToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SendToken{SendTokenCaller: SendTokenCaller{contract: contract}, SendTokenTransactor: SendTokenTransactor{contract: contract}, SendTokenFilterer: SendTokenFilterer{contract: contract}}, nil
}

// NewSendTokenCaller creates a new read-only instance of SendToken, bound to a specific deployed contract.
func NewSendTokenCaller(address common.Address, caller bind.ContractCaller) (*SendTokenCaller, error) {
	contract, err := bindSendToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SendTokenCaller{contract: contract}, nil
}

// NewSendTokenTransactor creates a new write-only instance of SendToken, bound to a specific deployed contract.
func NewSendTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SendTokenTransactor, error) {
	contract, err := bindSendToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SendTokenTransactor{contract: contract}, nil
}

// NewSendTokenFilterer creates a new log filterer instance of SendToken, bound to a specific deployed contract.
func NewSendTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SendTokenFilterer, error) {
	contract, err := bindSendToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SendTokenFilterer{contract: contract}, nil
}

// bindSendToken binds a generic wrapper to an already deployed contract.
func bindSendToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SendTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendToken *SendTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SendToken.Contract.SendTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendToken *SendTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendToken.Contract.SendTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendToken *SendTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendToken.Contract.SendTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SendToken *SendTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SendToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SendToken *SendTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SendToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SendToken *SendTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SendToken.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SendToken *SendTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SendToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SendToken *SendTokenSession) Owner() (common.Address, error) {
	return _SendToken.Contract.Owner(&_SendToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SendToken *SendTokenCallerSession) Owner() (common.Address, error) {
	return _SendToken.Contract.Owner(&_SendToken.CallOpts)
}

// SendToken is a paid mutator transaction binding the contract method 0x0a333a0f.
//
// Solidity: function sendToken(address token, address[] addrs, uint256[] values) returns()
func (_SendToken *SendTokenTransactor) SendToken(opts *bind.TransactOpts, token common.Address, addrs []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _SendToken.contract.Transact(opts, "sendToken", token, addrs, values)
}

// SendToken is a paid mutator transaction binding the contract method 0x0a333a0f.
//
// Solidity: function sendToken(address token, address[] addrs, uint256[] values) returns()
func (_SendToken *SendTokenSession) SendToken(token common.Address, addrs []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _SendToken.Contract.SendToken(&_SendToken.TransactOpts, token, addrs, values)
}

// SendToken is a paid mutator transaction binding the contract method 0x0a333a0f.
//
// Solidity: function sendToken(address token, address[] addrs, uint256[] values) returns()
func (_SendToken *SendTokenTransactorSession) SendToken(token common.Address, addrs []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _SendToken.Contract.SendToken(&_SendToken.TransactOpts, token, addrs, values)
}
