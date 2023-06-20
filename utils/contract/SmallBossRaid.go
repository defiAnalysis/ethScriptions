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

// SmallBossRaidBossInfo is an auto generated low-level Go binding around an user-defined struct.
type SmallBossRaidBossInfo struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}

// SmallBossRaidMetaData contains all meta data concerning the SmallBossRaid contract.
var SmallBossRaidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turn\",\"type\":\"uint256\"}],\"name\":\"SmallBattleBossEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"SmallBattleEndEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"}],\"name\":\"SmallBattleRaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"SmallBattleSkillEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"turn\",\"type\":\"uint256\"}],\"name\":\"SmallBossBeginEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"SmallHarvestEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minBossid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_maxBossid\",\"type\":\"uint256\"}],\"name\":\"SmallNewRound\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BossId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BossInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumSmallBossRaid.bossstatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"actives\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"battleBoss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"battleFlags\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginBattle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginFirst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beginRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bloodRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBossInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumSmallBossRaid.bossstatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structSmallBossRaid.BossInfo\",\"name\":\"_bossInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBossRaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getCryptoFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_flag\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bloodVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDamage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"enumSmallBossRaid.bossstatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structSmallBossRaid.BossInfo\",\"name\":\"_bossInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userHarm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nextTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"}],\"name\":\"getRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getSpendNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sig\",\"type\":\"bytes[]\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"level\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundBossId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundTurns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftAddr\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setBossInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sig\",\"type\":\"bytes[]\"}],\"name\":\"validSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoonerNFT\",\"outputs\":[{\"internalType\":\"contractCryptoZooNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SmallBossRaidABI is the input ABI used to generate the binding from.
// Deprecated: Use SmallBossRaidMetaData.ABI instead.
var SmallBossRaidABI = SmallBossRaidMetaData.ABI

// SmallBossRaid is an auto generated Go binding around an Ethereum contract.
type SmallBossRaid struct {
	SmallBossRaidCaller     // Read-only binding to the contract
	SmallBossRaidTransactor // Write-only binding to the contract
	SmallBossRaidFilterer   // Log filterer for contract events
}

// SmallBossRaidCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmallBossRaidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallBossRaidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmallBossRaidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallBossRaidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmallBossRaidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallBossRaidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmallBossRaidSession struct {
	Contract     *SmallBossRaid    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmallBossRaidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmallBossRaidCallerSession struct {
	Contract *SmallBossRaidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SmallBossRaidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmallBossRaidTransactorSession struct {
	Contract     *SmallBossRaidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SmallBossRaidRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmallBossRaidRaw struct {
	Contract *SmallBossRaid // Generic contract binding to access the raw methods on
}

// SmallBossRaidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmallBossRaidCallerRaw struct {
	Contract *SmallBossRaidCaller // Generic read-only contract binding to access the raw methods on
}

// SmallBossRaidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmallBossRaidTransactorRaw struct {
	Contract *SmallBossRaidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmallBossRaid creates a new instance of SmallBossRaid, bound to a specific deployed contract.
func NewSmallBossRaid(address common.Address, backend bind.ContractBackend) (*SmallBossRaid, error) {
	contract, err := bindSmallBossRaid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmallBossRaid{SmallBossRaidCaller: SmallBossRaidCaller{contract: contract}, SmallBossRaidTransactor: SmallBossRaidTransactor{contract: contract}, SmallBossRaidFilterer: SmallBossRaidFilterer{contract: contract}}, nil
}

// NewSmallBossRaidCaller creates a new read-only instance of SmallBossRaid, bound to a specific deployed contract.
func NewSmallBossRaidCaller(address common.Address, caller bind.ContractCaller) (*SmallBossRaidCaller, error) {
	contract, err := bindSmallBossRaid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidCaller{contract: contract}, nil
}

// NewSmallBossRaidTransactor creates a new write-only instance of SmallBossRaid, bound to a specific deployed contract.
func NewSmallBossRaidTransactor(address common.Address, transactor bind.ContractTransactor) (*SmallBossRaidTransactor, error) {
	contract, err := bindSmallBossRaid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidTransactor{contract: contract}, nil
}

// NewSmallBossRaidFilterer creates a new log filterer instance of SmallBossRaid, bound to a specific deployed contract.
func NewSmallBossRaidFilterer(address common.Address, filterer bind.ContractFilterer) (*SmallBossRaidFilterer, error) {
	contract, err := bindSmallBossRaid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidFilterer{contract: contract}, nil
}

// bindSmallBossRaid binds a generic wrapper to an already deployed contract.
func bindSmallBossRaid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmallBossRaidMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmallBossRaid *SmallBossRaidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmallBossRaid.Contract.SmallBossRaidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmallBossRaid *SmallBossRaidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.SmallBossRaidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmallBossRaid *SmallBossRaidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.SmallBossRaidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmallBossRaid *SmallBossRaidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmallBossRaid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmallBossRaid *SmallBossRaidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmallBossRaid *SmallBossRaidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.contract.Transact(opts, method, params...)
}

// BossId is a free data retrieval call binding the contract method 0x7bd36711.
//
// Solidity: function BossId() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) BossId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "BossId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BossId is a free data retrieval call binding the contract method 0x7bd36711.
//
// Solidity: function BossId() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) BossId() (*big.Int, error) {
	return _SmallBossRaid.Contract.BossId(&_SmallBossRaid.CallOpts)
}

// BossId is a free data retrieval call binding the contract method 0x7bd36711.
//
// Solidity: function BossId() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) BossId() (*big.Int, error) {
	return _SmallBossRaid.Contract.BossId(&_SmallBossRaid.CallOpts)
}

// BossInfos is a free data retrieval call binding the contract method 0x13254c66.
//
// Solidity: function BossInfos(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 totalDamage, uint256 totalValue, uint8 status)
func (_SmallBossRaid *SmallBossRaidCaller) BossInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "BossInfos", arg0)

	outstruct := new(struct {
		StartTime   *big.Int
		EndTime     *big.Int
		BloodVolume *big.Int
		TotalDamage *big.Int
		TotalValue  *big.Int
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BloodVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalDamage = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// BossInfos is a free data retrieval call binding the contract method 0x13254c66.
//
// Solidity: function BossInfos(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 totalDamage, uint256 totalValue, uint8 status)
func (_SmallBossRaid *SmallBossRaidSession) BossInfos(arg0 *big.Int) (struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}, error) {
	return _SmallBossRaid.Contract.BossInfos(&_SmallBossRaid.CallOpts, arg0)
}

// BossInfos is a free data retrieval call binding the contract method 0x13254c66.
//
// Solidity: function BossInfos(uint256 ) view returns(uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 totalDamage, uint256 totalValue, uint8 status)
func (_SmallBossRaid *SmallBossRaidCallerSession) BossInfos(arg0 *big.Int) (struct {
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	TotalDamage *big.Int
	TotalValue  *big.Int
	Status      uint8
}, error) {
	return _SmallBossRaid.Contract.BossInfos(&_SmallBossRaid.CallOpts, arg0)
}

// Actives is a free data retrieval call binding the contract method 0x947535b4.
//
// Solidity: function actives(uint256 ) view returns(bool)
func (_SmallBossRaid *SmallBossRaidCaller) Actives(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "actives", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Actives is a free data retrieval call binding the contract method 0x947535b4.
//
// Solidity: function actives(uint256 ) view returns(bool)
func (_SmallBossRaid *SmallBossRaidSession) Actives(arg0 *big.Int) (bool, error) {
	return _SmallBossRaid.Contract.Actives(&_SmallBossRaid.CallOpts, arg0)
}

// Actives is a free data retrieval call binding the contract method 0x947535b4.
//
// Solidity: function actives(uint256 ) view returns(bool)
func (_SmallBossRaid *SmallBossRaidCallerSession) Actives(arg0 *big.Int) (bool, error) {
	return _SmallBossRaid.Contract.Actives(&_SmallBossRaid.CallOpts, arg0)
}

// BattleFlags is a free data retrieval call binding the contract method 0xb2d57d5c.
//
// Solidity: function battleFlags(uint256 , uint256 ) view returns(bool)
func (_SmallBossRaid *SmallBossRaidCaller) BattleFlags(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "battleFlags", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BattleFlags is a free data retrieval call binding the contract method 0xb2d57d5c.
//
// Solidity: function battleFlags(uint256 , uint256 ) view returns(bool)
func (_SmallBossRaid *SmallBossRaidSession) BattleFlags(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _SmallBossRaid.Contract.BattleFlags(&_SmallBossRaid.CallOpts, arg0, arg1)
}

// BattleFlags is a free data retrieval call binding the contract method 0xb2d57d5c.
//
// Solidity: function battleFlags(uint256 , uint256 ) view returns(bool)
func (_SmallBossRaid *SmallBossRaidCallerSession) BattleFlags(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _SmallBossRaid.Contract.BattleFlags(&_SmallBossRaid.CallOpts, arg0, arg1)
}

// BloodRound is a free data retrieval call binding the contract method 0x7271c460.
//
// Solidity: function bloodRound() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) BloodRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "bloodRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BloodRound is a free data retrieval call binding the contract method 0x7271c460.
//
// Solidity: function bloodRound() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) BloodRound() (*big.Int, error) {
	return _SmallBossRaid.Contract.BloodRound(&_SmallBossRaid.CallOpts)
}

// BloodRound is a free data retrieval call binding the contract method 0x7271c460.
//
// Solidity: function bloodRound() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) BloodRound() (*big.Int, error) {
	return _SmallBossRaid.Contract.BloodRound(&_SmallBossRaid.CallOpts)
}

// GetBossInfo is a free data retrieval call binding the contract method 0xff14c6b4.
//
// Solidity: function getBossInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo)
func (_SmallBossRaid *SmallBossRaidCaller) GetBossInfo(opts *bind.CallOpts) (SmallBossRaidBossInfo, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "getBossInfo")

	if err != nil {
		return *new(SmallBossRaidBossInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(SmallBossRaidBossInfo)).(*SmallBossRaidBossInfo)

	return out0, err

}

// GetBossInfo is a free data retrieval call binding the contract method 0xff14c6b4.
//
// Solidity: function getBossInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo)
func (_SmallBossRaid *SmallBossRaidSession) GetBossInfo() (SmallBossRaidBossInfo, error) {
	return _SmallBossRaid.Contract.GetBossInfo(&_SmallBossRaid.CallOpts)
}

// GetBossInfo is a free data retrieval call binding the contract method 0xff14c6b4.
//
// Solidity: function getBossInfo() view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo)
func (_SmallBossRaid *SmallBossRaidCallerSession) GetBossInfo() (SmallBossRaidBossInfo, error) {
	return _SmallBossRaid.Contract.GetBossInfo(&_SmallBossRaid.CallOpts)
}

// GetBossRaid is a free data retrieval call binding the contract method 0x2bf89616.
//
// Solidity: function getBossRaid() view returns(bool _status)
func (_SmallBossRaid *SmallBossRaidCaller) GetBossRaid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "getBossRaid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBossRaid is a free data retrieval call binding the contract method 0x2bf89616.
//
// Solidity: function getBossRaid() view returns(bool _status)
func (_SmallBossRaid *SmallBossRaidSession) GetBossRaid() (bool, error) {
	return _SmallBossRaid.Contract.GetBossRaid(&_SmallBossRaid.CallOpts)
}

// GetBossRaid is a free data retrieval call binding the contract method 0x2bf89616.
//
// Solidity: function getBossRaid() view returns(bool _status)
func (_SmallBossRaid *SmallBossRaidCallerSession) GetBossRaid() (bool, error) {
	return _SmallBossRaid.Contract.GetBossRaid(&_SmallBossRaid.CallOpts)
}

// GetCryptoFlag is a free data retrieval call binding the contract method 0x19956651.
//
// Solidity: function getCryptoFlag(uint256 _tokenid) view returns(bool _flag)
func (_SmallBossRaid *SmallBossRaidCaller) GetCryptoFlag(opts *bind.CallOpts, _tokenid *big.Int) (bool, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "getCryptoFlag", _tokenid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetCryptoFlag is a free data retrieval call binding the contract method 0x19956651.
//
// Solidity: function getCryptoFlag(uint256 _tokenid) view returns(bool _flag)
func (_SmallBossRaid *SmallBossRaidSession) GetCryptoFlag(_tokenid *big.Int) (bool, error) {
	return _SmallBossRaid.Contract.GetCryptoFlag(&_SmallBossRaid.CallOpts, _tokenid)
}

// GetCryptoFlag is a free data retrieval call binding the contract method 0x19956651.
//
// Solidity: function getCryptoFlag(uint256 _tokenid) view returns(bool _flag)
func (_SmallBossRaid *SmallBossRaidCallerSession) GetCryptoFlag(_tokenid *big.Int) (bool, error) {
	return _SmallBossRaid.Contract.GetCryptoFlag(&_SmallBossRaid.CallOpts, _tokenid)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _sender) view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo, uint256 _level, uint256 _userHarm, uint256 _nextTime)
func (_SmallBossRaid *SmallBossRaidCaller) GetInfo(opts *bind.CallOpts, _sender common.Address) (struct {
	BossInfo SmallBossRaidBossInfo
	Level    *big.Int
	UserHarm *big.Int
	NextTime *big.Int
}, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "getInfo", _sender)

	outstruct := new(struct {
		BossInfo SmallBossRaidBossInfo
		Level    *big.Int
		UserHarm *big.Int
		NextTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BossInfo = *abi.ConvertType(out[0], new(SmallBossRaidBossInfo)).(*SmallBossRaidBossInfo)
	outstruct.Level = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UserHarm = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NextTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _sender) view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo, uint256 _level, uint256 _userHarm, uint256 _nextTime)
func (_SmallBossRaid *SmallBossRaidSession) GetInfo(_sender common.Address) (struct {
	BossInfo SmallBossRaidBossInfo
	Level    *big.Int
	UserHarm *big.Int
	NextTime *big.Int
}, error) {
	return _SmallBossRaid.Contract.GetInfo(&_SmallBossRaid.CallOpts, _sender)
}

// GetInfo is a free data retrieval call binding the contract method 0xffdd5cf1.
//
// Solidity: function getInfo(address _sender) view returns((uint256,uint256,uint256,uint256,uint256,uint8) _bossInfo, uint256 _level, uint256 _userHarm, uint256 _nextTime)
func (_SmallBossRaid *SmallBossRaidCallerSession) GetInfo(_sender common.Address) (struct {
	BossInfo SmallBossRaidBossInfo
	Level    *big.Int
	UserHarm *big.Int
	NextTime *big.Int
}, error) {
	return _SmallBossRaid.Contract.GetInfo(&_SmallBossRaid.CallOpts, _sender)
}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 _level) view returns(uint256 _round)
func (_SmallBossRaid *SmallBossRaidCaller) GetRound(opts *bind.CallOpts, _level *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "getRound", _level)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 _level) view returns(uint256 _round)
func (_SmallBossRaid *SmallBossRaidSession) GetRound(_level *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.GetRound(&_SmallBossRaid.CallOpts, _level)
}

// GetRound is a free data retrieval call binding the contract method 0x8f1327c0.
//
// Solidity: function getRound(uint256 _level) view returns(uint256 _round)
func (_SmallBossRaid *SmallBossRaidCallerSession) GetRound(_level *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.GetRound(&_SmallBossRaid.CallOpts, _level)
}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _user) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) GetSpendNonce(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "getSpendNonce", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _user) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) GetSpendNonce(_user common.Address) (*big.Int, error) {
	return _SmallBossRaid.Contract.GetSpendNonce(&_SmallBossRaid.CallOpts, _user)
}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _user) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) GetSpendNonce(_user common.Address) (*big.Int, error) {
	return _SmallBossRaid.Contract.GetSpendNonce(&_SmallBossRaid.CallOpts, _user)
}

// Level is a free data retrieval call binding the contract method 0x6fd5ae15.
//
// Solidity: function level() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) Level(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "level")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Level is a free data retrieval call binding the contract method 0x6fd5ae15.
//
// Solidity: function level() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) Level() (*big.Int, error) {
	return _SmallBossRaid.Contract.Level(&_SmallBossRaid.CallOpts)
}

// Level is a free data retrieval call binding the contract method 0x6fd5ae15.
//
// Solidity: function level() view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) Level() (*big.Int, error) {
	return _SmallBossRaid.Contract.Level(&_SmallBossRaid.CallOpts)
}

// RoundBossId is a free data retrieval call binding the contract method 0xea254a6f.
//
// Solidity: function roundBossId(uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) RoundBossId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "roundBossId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundBossId is a free data retrieval call binding the contract method 0xea254a6f.
//
// Solidity: function roundBossId(uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) RoundBossId(arg0 *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.RoundBossId(&_SmallBossRaid.CallOpts, arg0)
}

// RoundBossId is a free data retrieval call binding the contract method 0xea254a6f.
//
// Solidity: function roundBossId(uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) RoundBossId(arg0 *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.RoundBossId(&_SmallBossRaid.CallOpts, arg0)
}

// RoundTurns is a free data retrieval call binding the contract method 0x9feee924.
//
// Solidity: function roundTurns(uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) RoundTurns(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "roundTurns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundTurns is a free data retrieval call binding the contract method 0x9feee924.
//
// Solidity: function roundTurns(uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) RoundTurns(arg0 *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.RoundTurns(&_SmallBossRaid.CallOpts, arg0)
}

// RoundTurns is a free data retrieval call binding the contract method 0x9feee924.
//
// Solidity: function roundTurns(uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) RoundTurns(arg0 *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.RoundTurns(&_SmallBossRaid.CallOpts, arg0)
}

// UserRewards is a free data retrieval call binding the contract method 0x63c2a20a.
//
// Solidity: function userRewards(address , uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCaller) UserRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "userRewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewards is a free data retrieval call binding the contract method 0x63c2a20a.
//
// Solidity: function userRewards(address , uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidSession) UserRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.UserRewards(&_SmallBossRaid.CallOpts, arg0, arg1)
}

// UserRewards is a free data retrieval call binding the contract method 0x63c2a20a.
//
// Solidity: function userRewards(address , uint256 ) view returns(uint256)
func (_SmallBossRaid *SmallBossRaidCallerSession) UserRewards(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SmallBossRaid.Contract.UserRewards(&_SmallBossRaid.CallOpts, arg0, arg1)
}

// ValidSignature is a free data retrieval call binding the contract method 0x178c47f5.
//
// Solidity: function validSignature(address _sender, uint256 _amount, bytes[] _sig) view returns(bool)
func (_SmallBossRaid *SmallBossRaidCaller) ValidSignature(opts *bind.CallOpts, _sender common.Address, _amount *big.Int, _sig [][]byte) (bool, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "validSignature", _sender, _amount, _sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidSignature is a free data retrieval call binding the contract method 0x178c47f5.
//
// Solidity: function validSignature(address _sender, uint256 _amount, bytes[] _sig) view returns(bool)
func (_SmallBossRaid *SmallBossRaidSession) ValidSignature(_sender common.Address, _amount *big.Int, _sig [][]byte) (bool, error) {
	return _SmallBossRaid.Contract.ValidSignature(&_SmallBossRaid.CallOpts, _sender, _amount, _sig)
}

// ValidSignature is a free data retrieval call binding the contract method 0x178c47f5.
//
// Solidity: function validSignature(address _sender, uint256 _amount, bytes[] _sig) view returns(bool)
func (_SmallBossRaid *SmallBossRaidCallerSession) ValidSignature(_sender common.Address, _amount *big.Int, _sig [][]byte) (bool, error) {
	return _SmallBossRaid.Contract.ValidSignature(&_SmallBossRaid.CallOpts, _sender, _amount, _sig)
}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_SmallBossRaid *SmallBossRaidCaller) ZoonerNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SmallBossRaid.contract.Call(opts, &out, "zoonerNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_SmallBossRaid *SmallBossRaidSession) ZoonerNFT() (common.Address, error) {
	return _SmallBossRaid.Contract.ZoonerNFT(&_SmallBossRaid.CallOpts)
}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_SmallBossRaid *SmallBossRaidCallerSession) ZoonerNFT() (common.Address, error) {
	return _SmallBossRaid.Contract.ZoonerNFT(&_SmallBossRaid.CallOpts)
}

// BattleBoss is a paid mutator transaction binding the contract method 0x331b9c15.
//
// Solidity: function battleBoss(uint256 _amount, uint256 _tokenId) returns()
func (_SmallBossRaid *SmallBossRaidTransactor) BattleBoss(opts *bind.TransactOpts, _amount *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "battleBoss", _amount, _tokenId)
}

// BattleBoss is a paid mutator transaction binding the contract method 0x331b9c15.
//
// Solidity: function battleBoss(uint256 _amount, uint256 _tokenId) returns()
func (_SmallBossRaid *SmallBossRaidSession) BattleBoss(_amount *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BattleBoss(&_SmallBossRaid.TransactOpts, _amount, _tokenId)
}

// BattleBoss is a paid mutator transaction binding the contract method 0x331b9c15.
//
// Solidity: function battleBoss(uint256 _amount, uint256 _tokenId) returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) BattleBoss(_amount *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BattleBoss(&_SmallBossRaid.TransactOpts, _amount, _tokenId)
}

// BeginBattle is a paid mutator transaction binding the contract method 0xccaf4430.
//
// Solidity: function beginBattle() returns()
func (_SmallBossRaid *SmallBossRaidTransactor) BeginBattle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "beginBattle")
}

// BeginBattle is a paid mutator transaction binding the contract method 0xccaf4430.
//
// Solidity: function beginBattle() returns()
func (_SmallBossRaid *SmallBossRaidSession) BeginBattle() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BeginBattle(&_SmallBossRaid.TransactOpts)
}

// BeginBattle is a paid mutator transaction binding the contract method 0xccaf4430.
//
// Solidity: function beginBattle() returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) BeginBattle() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BeginBattle(&_SmallBossRaid.TransactOpts)
}

// BeginFirst is a paid mutator transaction binding the contract method 0x827ae457.
//
// Solidity: function beginFirst() returns()
func (_SmallBossRaid *SmallBossRaidTransactor) BeginFirst(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "beginFirst")
}

// BeginFirst is a paid mutator transaction binding the contract method 0x827ae457.
//
// Solidity: function beginFirst() returns()
func (_SmallBossRaid *SmallBossRaidSession) BeginFirst() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BeginFirst(&_SmallBossRaid.TransactOpts)
}

// BeginFirst is a paid mutator transaction binding the contract method 0x827ae457.
//
// Solidity: function beginFirst() returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) BeginFirst() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BeginFirst(&_SmallBossRaid.TransactOpts)
}

// BeginRound is a paid mutator transaction binding the contract method 0x65fdb08e.
//
// Solidity: function beginRound() returns()
func (_SmallBossRaid *SmallBossRaidTransactor) BeginRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "beginRound")
}

// BeginRound is a paid mutator transaction binding the contract method 0x65fdb08e.
//
// Solidity: function beginRound() returns()
func (_SmallBossRaid *SmallBossRaidSession) BeginRound() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BeginRound(&_SmallBossRaid.TransactOpts)
}

// BeginRound is a paid mutator transaction binding the contract method 0x65fdb08e.
//
// Solidity: function beginRound() returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) BeginRound() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.BeginRound(&_SmallBossRaid.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0xedb2f394.
//
// Solidity: function harvest(uint256 amount, bytes[] _sig) returns()
func (_SmallBossRaid *SmallBossRaidTransactor) Harvest(opts *bind.TransactOpts, amount *big.Int, _sig [][]byte) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "harvest", amount, _sig)
}

// Harvest is a paid mutator transaction binding the contract method 0xedb2f394.
//
// Solidity: function harvest(uint256 amount, bytes[] _sig) returns()
func (_SmallBossRaid *SmallBossRaidSession) Harvest(amount *big.Int, _sig [][]byte) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.Harvest(&_SmallBossRaid.TransactOpts, amount, _sig)
}

// Harvest is a paid mutator transaction binding the contract method 0xedb2f394.
//
// Solidity: function harvest(uint256 amount, bytes[] _sig) returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) Harvest(amount *big.Int, _sig [][]byte) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.Harvest(&_SmallBossRaid.TransactOpts, amount, _sig)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _owners) returns()
func (_SmallBossRaid *SmallBossRaidTransactor) Initialize(opts *bind.TransactOpts, _owners []common.Address) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "initialize", _owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _owners) returns()
func (_SmallBossRaid *SmallBossRaidSession) Initialize(_owners []common.Address) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.Initialize(&_SmallBossRaid.TransactOpts, _owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _owners) returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) Initialize(_owners []common.Address) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.Initialize(&_SmallBossRaid.TransactOpts, _owners)
}

// SetAddr is a paid mutator transaction binding the contract method 0xebcb00e0.
//
// Solidity: function setAddr(address _tokenAddr, address _nftAddr) returns()
func (_SmallBossRaid *SmallBossRaidTransactor) SetAddr(opts *bind.TransactOpts, _tokenAddr common.Address, _nftAddr common.Address) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "setAddr", _tokenAddr, _nftAddr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xebcb00e0.
//
// Solidity: function setAddr(address _tokenAddr, address _nftAddr) returns()
func (_SmallBossRaid *SmallBossRaidSession) SetAddr(_tokenAddr common.Address, _nftAddr common.Address) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.SetAddr(&_SmallBossRaid.TransactOpts, _tokenAddr, _nftAddr)
}

// SetAddr is a paid mutator transaction binding the contract method 0xebcb00e0.
//
// Solidity: function setAddr(address _tokenAddr, address _nftAddr) returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) SetAddr(_tokenAddr common.Address, _nftAddr common.Address) (*types.Transaction, error) {
	return _SmallBossRaid.Contract.SetAddr(&_SmallBossRaid.TransactOpts, _tokenAddr, _nftAddr)
}

// SetBossInfo is a paid mutator transaction binding the contract method 0x6c6b5812.
//
// Solidity: function setBossInfo() returns()
func (_SmallBossRaid *SmallBossRaidTransactor) SetBossInfo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "setBossInfo")
}

// SetBossInfo is a paid mutator transaction binding the contract method 0x6c6b5812.
//
// Solidity: function setBossInfo() returns()
func (_SmallBossRaid *SmallBossRaidSession) SetBossInfo() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.SetBossInfo(&_SmallBossRaid.TransactOpts)
}

// SetBossInfo is a paid mutator transaction binding the contract method 0x6c6b5812.
//
// Solidity: function setBossInfo() returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) SetBossInfo() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.SetBossInfo(&_SmallBossRaid.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SmallBossRaid *SmallBossRaidTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBossRaid.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SmallBossRaid *SmallBossRaidSession) Withdraw() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.Withdraw(&_SmallBossRaid.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SmallBossRaid *SmallBossRaidTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SmallBossRaid.Contract.Withdraw(&_SmallBossRaid.TransactOpts)
}

// SmallBossRaidInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SmallBossRaid contract.
type SmallBossRaidInitializedIterator struct {
	Event *SmallBossRaidInitialized // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidInitialized)
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
		it.Event = new(SmallBossRaidInitialized)
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
func (it *SmallBossRaidInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidInitialized represents a Initialized event raised by the SmallBossRaid contract.
type SmallBossRaidInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterInitialized(opts *bind.FilterOpts) (*SmallBossRaidInitializedIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidInitializedIterator{contract: _SmallBossRaid.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SmallBossRaidInitialized) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidInitialized)
				if err := _SmallBossRaid.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseInitialized(log types.Log) (*SmallBossRaidInitialized, error) {
	event := new(SmallBossRaidInitialized)
	if err := _SmallBossRaid.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallBattleBossEventIterator is returned from FilterSmallBattleBossEvent and is used to iterate over the raw logs and unpacked data for SmallBattleBossEvent events raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleBossEventIterator struct {
	Event *SmallBossRaidSmallBattleBossEvent // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallBattleBossEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallBattleBossEvent)
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
		it.Event = new(SmallBossRaidSmallBattleBossEvent)
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
func (it *SmallBossRaidSmallBattleBossEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallBattleBossEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallBattleBossEvent represents a SmallBattleBossEvent event raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleBossEvent struct {
	Id          *big.Int
	User        common.Address
	Amount      *big.Int
	TotalAmount *big.Int
	Level       *big.Int
	Tokenid     *big.Int
	Round       *big.Int
	Turn        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSmallBattleBossEvent is a free log retrieval operation binding the contract event 0x9f30afb22be824dc3c6073d74f962ea39a3cc351a5eac44bdd0a7d5cab56d632.
//
// Solidity: event SmallBattleBossEvent(uint256 id, address user, uint256 amount, uint256 totalAmount, uint256 level, uint256 tokenid, uint256 round, uint256 turn)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallBattleBossEvent(opts *bind.FilterOpts) (*SmallBossRaidSmallBattleBossEventIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallBattleBossEvent")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallBattleBossEventIterator{contract: _SmallBossRaid.contract, event: "SmallBattleBossEvent", logs: logs, sub: sub}, nil
}

// WatchSmallBattleBossEvent is a free log subscription operation binding the contract event 0x9f30afb22be824dc3c6073d74f962ea39a3cc351a5eac44bdd0a7d5cab56d632.
//
// Solidity: event SmallBattleBossEvent(uint256 id, address user, uint256 amount, uint256 totalAmount, uint256 level, uint256 tokenid, uint256 round, uint256 turn)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallBattleBossEvent(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallBattleBossEvent) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallBattleBossEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallBattleBossEvent)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleBossEvent", log); err != nil {
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

// ParseSmallBattleBossEvent is a log parse operation binding the contract event 0x9f30afb22be824dc3c6073d74f962ea39a3cc351a5eac44bdd0a7d5cab56d632.
//
// Solidity: event SmallBattleBossEvent(uint256 id, address user, uint256 amount, uint256 totalAmount, uint256 level, uint256 tokenid, uint256 round, uint256 turn)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallBattleBossEvent(log types.Log) (*SmallBossRaidSmallBattleBossEvent, error) {
	event := new(SmallBossRaidSmallBattleBossEvent)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleBossEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallBattleEndEventIterator is returned from FilterSmallBattleEndEvent and is used to iterate over the raw logs and unpacked data for SmallBattleEndEvent events raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleEndEventIterator struct {
	Event *SmallBossRaidSmallBattleEndEvent // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallBattleEndEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallBattleEndEvent)
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
		it.Event = new(SmallBossRaidSmallBattleEndEvent)
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
func (it *SmallBossRaidSmallBattleEndEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallBattleEndEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallBattleEndEvent represents a SmallBattleEndEvent event raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleEndEvent struct {
	Id    *big.Int
	Level *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSmallBattleEndEvent is a free log retrieval operation binding the contract event 0x05a9bdb0ec1f7e1fb8afa7eeee54934d2ebe1202e1ddf5b4f7bacab854e8cc83.
//
// Solidity: event SmallBattleEndEvent(uint256 id, uint256 level)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallBattleEndEvent(opts *bind.FilterOpts) (*SmallBossRaidSmallBattleEndEventIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallBattleEndEvent")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallBattleEndEventIterator{contract: _SmallBossRaid.contract, event: "SmallBattleEndEvent", logs: logs, sub: sub}, nil
}

// WatchSmallBattleEndEvent is a free log subscription operation binding the contract event 0x05a9bdb0ec1f7e1fb8afa7eeee54934d2ebe1202e1ddf5b4f7bacab854e8cc83.
//
// Solidity: event SmallBattleEndEvent(uint256 id, uint256 level)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallBattleEndEvent(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallBattleEndEvent) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallBattleEndEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallBattleEndEvent)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleEndEvent", log); err != nil {
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

// ParseSmallBattleEndEvent is a log parse operation binding the contract event 0x05a9bdb0ec1f7e1fb8afa7eeee54934d2ebe1202e1ddf5b4f7bacab854e8cc83.
//
// Solidity: event SmallBattleEndEvent(uint256 id, uint256 level)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallBattleEndEvent(log types.Log) (*SmallBossRaidSmallBattleEndEvent, error) {
	event := new(SmallBossRaidSmallBattleEndEvent)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleEndEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallBattleRaidEventIterator is returned from FilterSmallBattleRaidEvent and is used to iterate over the raw logs and unpacked data for SmallBattleRaidEvent events raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleRaidEventIterator struct {
	Event *SmallBossRaidSmallBattleRaidEvent // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallBattleRaidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallBattleRaidEvent)
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
		it.Event = new(SmallBossRaidSmallBattleRaidEvent)
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
func (it *SmallBossRaidSmallBattleRaidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallBattleRaidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallBattleRaidEvent represents a SmallBattleRaidEvent event raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleRaidEvent struct {
	Id          *big.Int
	Level       *big.Int
	TotalDamage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSmallBattleRaidEvent is a free log retrieval operation binding the contract event 0x622c3aa8c4767752a000ab0dbf7e1346a394a489cbcfc89e03d7498194b6fee4.
//
// Solidity: event SmallBattleRaidEvent(uint256 id, uint256 level, uint256 totalDamage)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallBattleRaidEvent(opts *bind.FilterOpts) (*SmallBossRaidSmallBattleRaidEventIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallBattleRaidEvent")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallBattleRaidEventIterator{contract: _SmallBossRaid.contract, event: "SmallBattleRaidEvent", logs: logs, sub: sub}, nil
}

// WatchSmallBattleRaidEvent is a free log subscription operation binding the contract event 0x622c3aa8c4767752a000ab0dbf7e1346a394a489cbcfc89e03d7498194b6fee4.
//
// Solidity: event SmallBattleRaidEvent(uint256 id, uint256 level, uint256 totalDamage)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallBattleRaidEvent(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallBattleRaidEvent) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallBattleRaidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallBattleRaidEvent)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleRaidEvent", log); err != nil {
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

// ParseSmallBattleRaidEvent is a log parse operation binding the contract event 0x622c3aa8c4767752a000ab0dbf7e1346a394a489cbcfc89e03d7498194b6fee4.
//
// Solidity: event SmallBattleRaidEvent(uint256 id, uint256 level, uint256 totalDamage)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallBattleRaidEvent(log types.Log) (*SmallBossRaidSmallBattleRaidEvent, error) {
	event := new(SmallBossRaidSmallBattleRaidEvent)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleRaidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallBattleSkillEventIterator is returned from FilterSmallBattleSkillEvent and is used to iterate over the raw logs and unpacked data for SmallBattleSkillEvent events raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleSkillEventIterator struct {
	Event *SmallBossRaidSmallBattleSkillEvent // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallBattleSkillEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallBattleSkillEvent)
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
		it.Event = new(SmallBossRaidSmallBattleSkillEvent)
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
func (it *SmallBossRaidSmallBattleSkillEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallBattleSkillEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallBattleSkillEvent represents a SmallBattleSkillEvent event raised by the SmallBossRaid contract.
type SmallBossRaidSmallBattleSkillEvent struct {
	Round *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSmallBattleSkillEvent is a free log retrieval operation binding the contract event 0xaed7de7af460b0c3a741b7767e95a2f2788e798f9d9882c137dcb9edd8d3ef5a.
//
// Solidity: event SmallBattleSkillEvent(uint256 round)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallBattleSkillEvent(opts *bind.FilterOpts) (*SmallBossRaidSmallBattleSkillEventIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallBattleSkillEvent")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallBattleSkillEventIterator{contract: _SmallBossRaid.contract, event: "SmallBattleSkillEvent", logs: logs, sub: sub}, nil
}

// WatchSmallBattleSkillEvent is a free log subscription operation binding the contract event 0xaed7de7af460b0c3a741b7767e95a2f2788e798f9d9882c137dcb9edd8d3ef5a.
//
// Solidity: event SmallBattleSkillEvent(uint256 round)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallBattleSkillEvent(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallBattleSkillEvent) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallBattleSkillEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallBattleSkillEvent)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleSkillEvent", log); err != nil {
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

// ParseSmallBattleSkillEvent is a log parse operation binding the contract event 0xaed7de7af460b0c3a741b7767e95a2f2788e798f9d9882c137dcb9edd8d3ef5a.
//
// Solidity: event SmallBattleSkillEvent(uint256 round)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallBattleSkillEvent(log types.Log) (*SmallBossRaidSmallBattleSkillEvent, error) {
	event := new(SmallBossRaidSmallBattleSkillEvent)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBattleSkillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallBossBeginEventIterator is returned from FilterSmallBossBeginEvent and is used to iterate over the raw logs and unpacked data for SmallBossBeginEvent events raised by the SmallBossRaid contract.
type SmallBossRaidSmallBossBeginEventIterator struct {
	Event *SmallBossRaidSmallBossBeginEvent // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallBossBeginEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallBossBeginEvent)
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
		it.Event = new(SmallBossRaidSmallBossBeginEvent)
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
func (it *SmallBossRaidSmallBossBeginEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallBossBeginEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallBossBeginEvent represents a SmallBossBeginEvent event raised by the SmallBossRaid contract.
type SmallBossRaidSmallBossBeginEvent struct {
	Id          *big.Int
	Level       *big.Int
	StartTime   *big.Int
	EndTime     *big.Int
	BloodVolume *big.Int
	Value       *big.Int
	Round       *big.Int
	Turn        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSmallBossBeginEvent is a free log retrieval operation binding the contract event 0x34fd63fe00b0303512385b2068e1fffb8db5648e4d5304e751954a7794659a50.
//
// Solidity: event SmallBossBeginEvent(uint256 id, uint256 level, uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 value, uint256 round, uint256 turn)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallBossBeginEvent(opts *bind.FilterOpts) (*SmallBossRaidSmallBossBeginEventIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallBossBeginEvent")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallBossBeginEventIterator{contract: _SmallBossRaid.contract, event: "SmallBossBeginEvent", logs: logs, sub: sub}, nil
}

// WatchSmallBossBeginEvent is a free log subscription operation binding the contract event 0x34fd63fe00b0303512385b2068e1fffb8db5648e4d5304e751954a7794659a50.
//
// Solidity: event SmallBossBeginEvent(uint256 id, uint256 level, uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 value, uint256 round, uint256 turn)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallBossBeginEvent(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallBossBeginEvent) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallBossBeginEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallBossBeginEvent)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBossBeginEvent", log); err != nil {
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

// ParseSmallBossBeginEvent is a log parse operation binding the contract event 0x34fd63fe00b0303512385b2068e1fffb8db5648e4d5304e751954a7794659a50.
//
// Solidity: event SmallBossBeginEvent(uint256 id, uint256 level, uint256 startTime, uint256 endTime, uint256 bloodVolume, uint256 value, uint256 round, uint256 turn)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallBossBeginEvent(log types.Log) (*SmallBossRaidSmallBossBeginEvent, error) {
	event := new(SmallBossRaidSmallBossBeginEvent)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallBossBeginEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallHarvestEventIterator is returned from FilterSmallHarvestEvent and is used to iterate over the raw logs and unpacked data for SmallHarvestEvent events raised by the SmallBossRaid contract.
type SmallBossRaidSmallHarvestEventIterator struct {
	Event *SmallBossRaidSmallHarvestEvent // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallHarvestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallHarvestEvent)
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
		it.Event = new(SmallBossRaidSmallHarvestEvent)
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
func (it *SmallBossRaidSmallHarvestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallHarvestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallHarvestEvent represents a SmallHarvestEvent event raised by the SmallBossRaid contract.
type SmallBossRaidSmallHarvestEvent struct {
	Sender  common.Address
	OrderId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSmallHarvestEvent is a free log retrieval operation binding the contract event 0xb5daa28e2eb2c0a618fc3bf3c1f0fd3785345595f43d90ef1fabfd382de83b79.
//
// Solidity: event SmallHarvestEvent(address sender, uint256 orderId, uint256 time)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallHarvestEvent(opts *bind.FilterOpts) (*SmallBossRaidSmallHarvestEventIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallHarvestEvent")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallHarvestEventIterator{contract: _SmallBossRaid.contract, event: "SmallHarvestEvent", logs: logs, sub: sub}, nil
}

// WatchSmallHarvestEvent is a free log subscription operation binding the contract event 0xb5daa28e2eb2c0a618fc3bf3c1f0fd3785345595f43d90ef1fabfd382de83b79.
//
// Solidity: event SmallHarvestEvent(address sender, uint256 orderId, uint256 time)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallHarvestEvent(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallHarvestEvent) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallHarvestEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallHarvestEvent)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallHarvestEvent", log); err != nil {
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

// ParseSmallHarvestEvent is a log parse operation binding the contract event 0xb5daa28e2eb2c0a618fc3bf3c1f0fd3785345595f43d90ef1fabfd382de83b79.
//
// Solidity: event SmallHarvestEvent(address sender, uint256 orderId, uint256 time)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallHarvestEvent(log types.Log) (*SmallBossRaidSmallHarvestEvent, error) {
	event := new(SmallBossRaidSmallHarvestEvent)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallHarvestEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmallBossRaidSmallNewRoundIterator is returned from FilterSmallNewRound and is used to iterate over the raw logs and unpacked data for SmallNewRound events raised by the SmallBossRaid contract.
type SmallBossRaidSmallNewRoundIterator struct {
	Event *SmallBossRaidSmallNewRound // Event containing the contract specifics and raw log

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
func (it *SmallBossRaidSmallNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmallBossRaidSmallNewRound)
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
		it.Event = new(SmallBossRaidSmallNewRound)
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
func (it *SmallBossRaidSmallNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmallBossRaidSmallNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmallBossRaidSmallNewRound represents a SmallNewRound event raised by the SmallBossRaid contract.
type SmallBossRaidSmallNewRound struct {
	Round     *big.Int
	MaxTurn   *big.Int
	MinBossid *big.Int
	MaxBossid *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSmallNewRound is a free log retrieval operation binding the contract event 0x750d36bd1cdf46b79b6adad44c6668fd81c8d1315614653783eba277bba12e84.
//
// Solidity: event SmallNewRound(uint256 round, uint256 maxTurn, uint256 _minBossid, uint256 _maxBossid)
func (_SmallBossRaid *SmallBossRaidFilterer) FilterSmallNewRound(opts *bind.FilterOpts) (*SmallBossRaidSmallNewRoundIterator, error) {

	logs, sub, err := _SmallBossRaid.contract.FilterLogs(opts, "SmallNewRound")
	if err != nil {
		return nil, err
	}
	return &SmallBossRaidSmallNewRoundIterator{contract: _SmallBossRaid.contract, event: "SmallNewRound", logs: logs, sub: sub}, nil
}

// WatchSmallNewRound is a free log subscription operation binding the contract event 0x750d36bd1cdf46b79b6adad44c6668fd81c8d1315614653783eba277bba12e84.
//
// Solidity: event SmallNewRound(uint256 round, uint256 maxTurn, uint256 _minBossid, uint256 _maxBossid)
func (_SmallBossRaid *SmallBossRaidFilterer) WatchSmallNewRound(opts *bind.WatchOpts, sink chan<- *SmallBossRaidSmallNewRound) (event.Subscription, error) {

	logs, sub, err := _SmallBossRaid.contract.WatchLogs(opts, "SmallNewRound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmallBossRaidSmallNewRound)
				if err := _SmallBossRaid.contract.UnpackLog(event, "SmallNewRound", log); err != nil {
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

// ParseSmallNewRound is a log parse operation binding the contract event 0x750d36bd1cdf46b79b6adad44c6668fd81c8d1315614653783eba277bba12e84.
//
// Solidity: event SmallNewRound(uint256 round, uint256 maxTurn, uint256 _minBossid, uint256 _maxBossid)
func (_SmallBossRaid *SmallBossRaidFilterer) ParseSmallNewRound(log types.Log) (*SmallBossRaidSmallNewRound, error) {
	event := new(SmallBossRaidSmallNewRound)
	if err := _SmallBossRaid.contract.UnpackLog(event, "SmallNewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
