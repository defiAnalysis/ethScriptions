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

// CryptoZooNFTCryptoZoon is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTCryptoZoon struct {
	Tribe      uint8
	Generation *big.Int
	Exp        *big.Int
	Dna        *big.Int
	FarmTime   *big.Int
	BornTime   *big.Int
}

// CryptoZooNFTItemSale is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTItemSale struct {
	TokenId *big.Int
	Owner   common.Address
	Price   *big.Int
}

// CryptoZooNFTSynPair is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTSynPair struct {
	Elves1 CryptoZooNFTZoonInfo
	Elves2 CryptoZooNFTZoonInfo
	Elves3 CryptoZooNFTZoonInfo
	UpTs   *big.Int
}

// CryptoZooNFTUserZoonInfoList is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTUserZoonInfoList struct {
	ZoonInfo CryptoZooNFTZoonInfo
	Harm     *big.Int
	Physical *big.Int
	Total    *big.Int
	Skill    *big.Int
}

// CryptoZooNFTZoonInfo is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTZoonInfo struct {
	Zoon    CryptoZooNFTCryptoZoon
	Tokenid *big.Int
}

// CryptoZooNFTMetaData contains all meta data concerning the CryptoZooNFT contract.
var CryptoZooNFTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"CancelOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"}],\"name\":\"Evolve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"Exp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"FillOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"LayEgg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PlaceOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"SynthEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"}],\"name\":\"UpdateTribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGeneration\",\"type\":\"uint256\"}],\"name\":\"UpgradeGeneration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"Working\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"_tokenids\",\"type\":\"uint128[]\"}],\"name\":\"addRare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"delBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dna\",\"type\":\"uint256\"}],\"name\":\"evolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_exp\",\"type\":\"uint256\"}],\"name\":\"exp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeMarketRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getNFTInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"_zoon\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_token0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_token1\",\"type\":\"uint256\"}],\"name\":\"getPairs\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"_zoon\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getPhysical\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_physical\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getPhysicalAndMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_physical\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRareAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRarePhy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getSale\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ItemSale\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSynthHistory\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"elves1\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"elves2\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"elves3\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"upTs\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.SynPair[]\",\"name\":\"_info\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getUserSkill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_skill\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getZooner\",\"outputs\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isEvolved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"enumCryptoZooNFT.Tribe[]\",\"name\":\"tribes\",\"type\":\"uint8[]\"}],\"name\":\"layEgg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"listAllNftByOwner\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"zoonInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"harm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"physical\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skill\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.UserZoonInfoList[]\",\"name\":\"available\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"zoonInfo\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"harm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"physical\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skill\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.UserZoonInfoList[]\",\"name\":\"sell\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"listNftByOwner\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractManagerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketsSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"mulPhysical\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"payPhysical\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"placeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceEgg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceEggDogCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceEggUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenids\",\"type\":\"uint256[]\"}],\"name\":\"rebasePhysicalTotal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"_tokenids\",\"type\":\"uint128[]\"}],\"name\":\"reduceRare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenid\",\"type\":\"uint256[]\"}],\"name\":\"setBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zoonerERC20\",\"type\":\"address\"}],\"name\":\"setERC20AndManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totals\",\"type\":\"uint256[]\"}],\"name\":\"setPhysicalTotal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"setRare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenid\",\"type\":\"uint256[]\"}],\"name\":\"setReversion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_smallBoss\",\"type\":\"address\"}],\"name\":\"setSmallBoss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"setSync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"_tokenids\",\"type\":\"uint128[]\"}],\"name\":\"setUserRare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_skills\",\"type\":\"uint256[]\"}],\"name\":\"setUserSkill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smallBoss\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenids\",\"type\":\"uint256[]\"}],\"name\":\"synthesis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenSaleByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenSaleOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"working\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoonerERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"zoonerLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CryptoZooNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CryptoZooNFTMetaData.ABI instead.
var CryptoZooNFTABI = CryptoZooNFTMetaData.ABI

// CryptoZooNFT is an auto generated Go binding around an Ethereum contract.
type CryptoZooNFT struct {
	CryptoZooNFTCaller     // Read-only binding to the contract
	CryptoZooNFTTransactor // Write-only binding to the contract
	CryptoZooNFTFilterer   // Log filterer for contract events
}

// CryptoZooNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoZooNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoZooNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoZooNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoZooNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoZooNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoZooNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoZooNFTSession struct {
	Contract     *CryptoZooNFT     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoZooNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoZooNFTCallerSession struct {
	Contract *CryptoZooNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CryptoZooNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoZooNFTTransactorSession struct {
	Contract     *CryptoZooNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CryptoZooNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoZooNFTRaw struct {
	Contract *CryptoZooNFT // Generic contract binding to access the raw methods on
}

// CryptoZooNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoZooNFTCallerRaw struct {
	Contract *CryptoZooNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoZooNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoZooNFTTransactorRaw struct {
	Contract *CryptoZooNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoZooNFT creates a new instance of CryptoZooNFT, bound to a specific deployed contract.
func NewCryptoZooNFT(address common.Address, backend bind.ContractBackend) (*CryptoZooNFT, error) {
	contract, err := bindCryptoZooNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFT{CryptoZooNFTCaller: CryptoZooNFTCaller{contract: contract}, CryptoZooNFTTransactor: CryptoZooNFTTransactor{contract: contract}, CryptoZooNFTFilterer: CryptoZooNFTFilterer{contract: contract}}, nil
}

// NewCryptoZooNFTCaller creates a new read-only instance of CryptoZooNFT, bound to a specific deployed contract.
func NewCryptoZooNFTCaller(address common.Address, caller bind.ContractCaller) (*CryptoZooNFTCaller, error) {
	contract, err := bindCryptoZooNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTCaller{contract: contract}, nil
}

// NewCryptoZooNFTTransactor creates a new write-only instance of CryptoZooNFT, bound to a specific deployed contract.
func NewCryptoZooNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoZooNFTTransactor, error) {
	contract, err := bindCryptoZooNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTTransactor{contract: contract}, nil
}

// NewCryptoZooNFTFilterer creates a new log filterer instance of CryptoZooNFT, bound to a specific deployed contract.
func NewCryptoZooNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoZooNFTFilterer, error) {
	contract, err := bindCryptoZooNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTFilterer{contract: contract}, nil
}

// bindCryptoZooNFT binds a generic wrapper to an already deployed contract.
func bindCryptoZooNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CryptoZooNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoZooNFT *CryptoZooNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoZooNFT.Contract.CryptoZooNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoZooNFT *CryptoZooNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.CryptoZooNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoZooNFT *CryptoZooNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.CryptoZooNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoZooNFT *CryptoZooNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoZooNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoZooNFT *CryptoZooNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoZooNFT *CryptoZooNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CryptoZooNFT.Contract.BalanceOf(&_CryptoZooNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CryptoZooNFT.Contract.BalanceOf(&_CryptoZooNFT.CallOpts, owner)
}

// FeeMarketRate is a free data retrieval call binding the contract method 0x43aa2b6c.
//
// Solidity: function feeMarketRate() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) FeeMarketRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "feeMarketRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeMarketRate is a free data retrieval call binding the contract method 0x43aa2b6c.
//
// Solidity: function feeMarketRate() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) FeeMarketRate() (*big.Int, error) {
	return _CryptoZooNFT.Contract.FeeMarketRate(&_CryptoZooNFT.CallOpts)
}

// FeeMarketRate is a free data retrieval call binding the contract method 0x43aa2b6c.
//
// Solidity: function feeMarketRate() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) FeeMarketRate() (*big.Int, error) {
	return _CryptoZooNFT.Contract.FeeMarketRate(&_CryptoZooNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CryptoZooNFT *CryptoZooNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CryptoZooNFT.Contract.GetApproved(&_CryptoZooNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CryptoZooNFT.Contract.GetApproved(&_CryptoZooNFT.CallOpts, tokenId)
}

// GetNFTInfo is a free data retrieval call binding the contract method 0xd188929f.
//
// Solidity: function getNFTInfo(uint256 _tokenid) view returns(address _owner, (uint8,uint256,uint256,uint256,uint256,uint256) _zoon)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetNFTInfo(opts *bind.CallOpts, _tokenid *big.Int) (struct {
	Owner common.Address
	Zoon  CryptoZooNFTCryptoZoon
}, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getNFTInfo", _tokenid)

	outstruct := new(struct {
		Owner common.Address
		Zoon  CryptoZooNFTCryptoZoon
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Zoon = *abi.ConvertType(out[1], new(CryptoZooNFTCryptoZoon)).(*CryptoZooNFTCryptoZoon)

	return *outstruct, err

}

// GetNFTInfo is a free data retrieval call binding the contract method 0xd188929f.
//
// Solidity: function getNFTInfo(uint256 _tokenid) view returns(address _owner, (uint8,uint256,uint256,uint256,uint256,uint256) _zoon)
func (_CryptoZooNFT *CryptoZooNFTSession) GetNFTInfo(_tokenid *big.Int) (struct {
	Owner common.Address
	Zoon  CryptoZooNFTCryptoZoon
}, error) {
	return _CryptoZooNFT.Contract.GetNFTInfo(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetNFTInfo is a free data retrieval call binding the contract method 0xd188929f.
//
// Solidity: function getNFTInfo(uint256 _tokenid) view returns(address _owner, (uint8,uint256,uint256,uint256,uint256,uint256) _zoon)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetNFTInfo(_tokenid *big.Int) (struct {
	Owner common.Address
	Zoon  CryptoZooNFTCryptoZoon
}, error) {
	return _CryptoZooNFT.Contract.GetNFTInfo(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetPairs is a free data retrieval call binding the contract method 0xf2364e91.
//
// Solidity: function getPairs(uint256 _token0, uint256 _token1) view returns(((uint8,uint256,uint256,uint256,uint256,uint256),uint256) _zoon)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetPairs(opts *bind.CallOpts, _token0 *big.Int, _token1 *big.Int) (CryptoZooNFTZoonInfo, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getPairs", _token0, _token1)

	if err != nil {
		return *new(CryptoZooNFTZoonInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(CryptoZooNFTZoonInfo)).(*CryptoZooNFTZoonInfo)

	return out0, err

}

// GetPairs is a free data retrieval call binding the contract method 0xf2364e91.
//
// Solidity: function getPairs(uint256 _token0, uint256 _token1) view returns(((uint8,uint256,uint256,uint256,uint256,uint256),uint256) _zoon)
func (_CryptoZooNFT *CryptoZooNFTSession) GetPairs(_token0 *big.Int, _token1 *big.Int) (CryptoZooNFTZoonInfo, error) {
	return _CryptoZooNFT.Contract.GetPairs(&_CryptoZooNFT.CallOpts, _token0, _token1)
}

// GetPairs is a free data retrieval call binding the contract method 0xf2364e91.
//
// Solidity: function getPairs(uint256 _token0, uint256 _token1) view returns(((uint8,uint256,uint256,uint256,uint256,uint256),uint256) _zoon)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetPairs(_token0 *big.Int, _token1 *big.Int) (CryptoZooNFTZoonInfo, error) {
	return _CryptoZooNFT.Contract.GetPairs(&_CryptoZooNFT.CallOpts, _token0, _token1)
}

// GetPhysical is a free data retrieval call binding the contract method 0x48024bfd.
//
// Solidity: function getPhysical(uint256 _tokenid) view returns(uint256 _physical)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetPhysical(opts *bind.CallOpts, _tokenid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getPhysical", _tokenid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPhysical is a free data retrieval call binding the contract method 0x48024bfd.
//
// Solidity: function getPhysical(uint256 _tokenid) view returns(uint256 _physical)
func (_CryptoZooNFT *CryptoZooNFTSession) GetPhysical(_tokenid *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetPhysical(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetPhysical is a free data retrieval call binding the contract method 0x48024bfd.
//
// Solidity: function getPhysical(uint256 _tokenid) view returns(uint256 _physical)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetPhysical(_tokenid *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetPhysical(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetPhysicalAndMax is a free data retrieval call binding the contract method 0x4f0ff104.
//
// Solidity: function getPhysicalAndMax(uint256 _tokenid) view returns(uint256 _physical, uint256 _total)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetPhysicalAndMax(opts *bind.CallOpts, _tokenid *big.Int) (struct {
	Physical *big.Int
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getPhysicalAndMax", _tokenid)

	outstruct := new(struct {
		Physical *big.Int
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Physical = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPhysicalAndMax is a free data retrieval call binding the contract method 0x4f0ff104.
//
// Solidity: function getPhysicalAndMax(uint256 _tokenid) view returns(uint256 _physical, uint256 _total)
func (_CryptoZooNFT *CryptoZooNFTSession) GetPhysicalAndMax(_tokenid *big.Int) (struct {
	Physical *big.Int
	Total    *big.Int
}, error) {
	return _CryptoZooNFT.Contract.GetPhysicalAndMax(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetPhysicalAndMax is a free data retrieval call binding the contract method 0x4f0ff104.
//
// Solidity: function getPhysicalAndMax(uint256 _tokenid) view returns(uint256 _physical, uint256 _total)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetPhysicalAndMax(_tokenid *big.Int) (struct {
	Physical *big.Int
	Total    *big.Int
}, error) {
	return _CryptoZooNFT.Contract.GetPhysicalAndMax(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetRare is a free data retrieval call binding the contract method 0xa90b9c72.
//
// Solidity: function getRare(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetRare(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getRare", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRare is a free data retrieval call binding the contract method 0xa90b9c72.
//
// Solidity: function getRare(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) GetRare(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetRare(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetRare is a free data retrieval call binding the contract method 0xa90b9c72.
//
// Solidity: function getRare(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetRare(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetRare(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetRareAmount is a free data retrieval call binding the contract method 0x7329df29.
//
// Solidity: function getRareAmount(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetRareAmount(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getRareAmount", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRareAmount is a free data retrieval call binding the contract method 0x7329df29.
//
// Solidity: function getRareAmount(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) GetRareAmount(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetRareAmount(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetRareAmount is a free data retrieval call binding the contract method 0x7329df29.
//
// Solidity: function getRareAmount(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetRareAmount(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetRareAmount(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetRarePhy is a free data retrieval call binding the contract method 0x44b5a107.
//
// Solidity: function getRarePhy(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetRarePhy(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getRarePhy", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRarePhy is a free data retrieval call binding the contract method 0x44b5a107.
//
// Solidity: function getRarePhy(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) GetRarePhy(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetRarePhy(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetRarePhy is a free data retrieval call binding the contract method 0x44b5a107.
//
// Solidity: function getRarePhy(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetRarePhy(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetRarePhy(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetSale is a free data retrieval call binding the contract method 0xd8f6d596.
//
// Solidity: function getSale(uint256 _tokenId) view returns((uint256,address,uint256))
func (_CryptoZooNFT *CryptoZooNFTCaller) GetSale(opts *bind.CallOpts, _tokenId *big.Int) (CryptoZooNFTItemSale, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getSale", _tokenId)

	if err != nil {
		return *new(CryptoZooNFTItemSale), err
	}

	out0 := *abi.ConvertType(out[0], new(CryptoZooNFTItemSale)).(*CryptoZooNFTItemSale)

	return out0, err

}

// GetSale is a free data retrieval call binding the contract method 0xd8f6d596.
//
// Solidity: function getSale(uint256 _tokenId) view returns((uint256,address,uint256))
func (_CryptoZooNFT *CryptoZooNFTSession) GetSale(_tokenId *big.Int) (CryptoZooNFTItemSale, error) {
	return _CryptoZooNFT.Contract.GetSale(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetSale is a free data retrieval call binding the contract method 0xd8f6d596.
//
// Solidity: function getSale(uint256 _tokenId) view returns((uint256,address,uint256))
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetSale(_tokenId *big.Int) (CryptoZooNFTItemSale, error) {
	return _CryptoZooNFT.Contract.GetSale(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetSynthHistory is a free data retrieval call binding the contract method 0x3d4f6c5b.
//
// Solidity: function getSynthHistory(address sender) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),((uint8,uint256,uint256,uint256,uint256,uint256),uint256),((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256)[] _info)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetSynthHistory(opts *bind.CallOpts, sender common.Address) ([]CryptoZooNFTSynPair, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getSynthHistory", sender)

	if err != nil {
		return *new([]CryptoZooNFTSynPair), err
	}

	out0 := *abi.ConvertType(out[0], new([]CryptoZooNFTSynPair)).(*[]CryptoZooNFTSynPair)

	return out0, err

}

// GetSynthHistory is a free data retrieval call binding the contract method 0x3d4f6c5b.
//
// Solidity: function getSynthHistory(address sender) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),((uint8,uint256,uint256,uint256,uint256,uint256),uint256),((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256)[] _info)
func (_CryptoZooNFT *CryptoZooNFTSession) GetSynthHistory(sender common.Address) ([]CryptoZooNFTSynPair, error) {
	return _CryptoZooNFT.Contract.GetSynthHistory(&_CryptoZooNFT.CallOpts, sender)
}

// GetSynthHistory is a free data retrieval call binding the contract method 0x3d4f6c5b.
//
// Solidity: function getSynthHistory(address sender) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),((uint8,uint256,uint256,uint256,uint256,uint256),uint256),((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256)[] _info)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetSynthHistory(sender common.Address) ([]CryptoZooNFTSynPair, error) {
	return _CryptoZooNFT.Contract.GetSynthHistory(&_CryptoZooNFT.CallOpts, sender)
}

// GetUserSkill is a free data retrieval call binding the contract method 0xeb2043c1.
//
// Solidity: function getUserSkill(uint256 _tokenid) view returns(uint256 _skill)
func (_CryptoZooNFT *CryptoZooNFTCaller) GetUserSkill(opts *bind.CallOpts, _tokenid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getUserSkill", _tokenid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserSkill is a free data retrieval call binding the contract method 0xeb2043c1.
//
// Solidity: function getUserSkill(uint256 _tokenid) view returns(uint256 _skill)
func (_CryptoZooNFT *CryptoZooNFTSession) GetUserSkill(_tokenid *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetUserSkill(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetUserSkill is a free data retrieval call binding the contract method 0xeb2043c1.
//
// Solidity: function getUserSkill(uint256 _tokenid) view returns(uint256 _skill)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetUserSkill(_tokenid *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.GetUserSkill(&_CryptoZooNFT.CallOpts, _tokenid)
}

// GetZooner is a free data retrieval call binding the contract method 0x57eb46ee.
//
// Solidity: function getZooner(uint256 _tokenId) view returns((uint8,uint256,uint256,uint256,uint256,uint256))
func (_CryptoZooNFT *CryptoZooNFTCaller) GetZooner(opts *bind.CallOpts, _tokenId *big.Int) (CryptoZooNFTCryptoZoon, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "getZooner", _tokenId)

	if err != nil {
		return *new(CryptoZooNFTCryptoZoon), err
	}

	out0 := *abi.ConvertType(out[0], new(CryptoZooNFTCryptoZoon)).(*CryptoZooNFTCryptoZoon)

	return out0, err

}

// GetZooner is a free data retrieval call binding the contract method 0x57eb46ee.
//
// Solidity: function getZooner(uint256 _tokenId) view returns((uint8,uint256,uint256,uint256,uint256,uint256))
func (_CryptoZooNFT *CryptoZooNFTSession) GetZooner(_tokenId *big.Int) (CryptoZooNFTCryptoZoon, error) {
	return _CryptoZooNFT.Contract.GetZooner(&_CryptoZooNFT.CallOpts, _tokenId)
}

// GetZooner is a free data retrieval call binding the contract method 0x57eb46ee.
//
// Solidity: function getZooner(uint256 _tokenId) view returns((uint8,uint256,uint256,uint256,uint256,uint256))
func (_CryptoZooNFT *CryptoZooNFTCallerSession) GetZooner(_tokenId *big.Int) (CryptoZooNFTCryptoZoon, error) {
	return _CryptoZooNFT.Contract.GetZooner(&_CryptoZooNFT.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CryptoZooNFT.Contract.IsApprovedForAll(&_CryptoZooNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CryptoZooNFT.Contract.IsApprovedForAll(&_CryptoZooNFT.CallOpts, owner, operator)
}

// IsEvolved is a free data retrieval call binding the contract method 0x4918f947.
//
// Solidity: function isEvolved(uint256 ) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTCaller) IsEvolved(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "isEvolved", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEvolved is a free data retrieval call binding the contract method 0x4918f947.
//
// Solidity: function isEvolved(uint256 ) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTSession) IsEvolved(arg0 *big.Int) (bool, error) {
	return _CryptoZooNFT.Contract.IsEvolved(&_CryptoZooNFT.CallOpts, arg0)
}

// IsEvolved is a free data retrieval call binding the contract method 0x4918f947.
//
// Solidity: function isEvolved(uint256 ) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) IsEvolved(arg0 *big.Int) (bool, error) {
	return _CryptoZooNFT.Contract.IsEvolved(&_CryptoZooNFT.CallOpts, arg0)
}

// LatestTokenId is a free data retrieval call binding the contract method 0x8c0e8349.
//
// Solidity: function latestTokenId() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) LatestTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "latestTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTokenId is a free data retrieval call binding the contract method 0x8c0e8349.
//
// Solidity: function latestTokenId() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) LatestTokenId() (*big.Int, error) {
	return _CryptoZooNFT.Contract.LatestTokenId(&_CryptoZooNFT.CallOpts)
}

// LatestTokenId is a free data retrieval call binding the contract method 0x8c0e8349.
//
// Solidity: function latestTokenId() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) LatestTokenId() (*big.Int, error) {
	return _CryptoZooNFT.Contract.LatestTokenId(&_CryptoZooNFT.CallOpts)
}

// ListAllNftByOwner is a free data retrieval call binding the contract method 0x6e0630f8.
//
// Solidity: function listAllNftByOwner(address _owner) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256,uint256)[] available, (((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256,uint256)[] sell)
func (_CryptoZooNFT *CryptoZooNFTCaller) ListAllNftByOwner(opts *bind.CallOpts, _owner common.Address) (struct {
	Available []CryptoZooNFTUserZoonInfoList
	Sell      []CryptoZooNFTUserZoonInfoList
}, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "listAllNftByOwner", _owner)

	outstruct := new(struct {
		Available []CryptoZooNFTUserZoonInfoList
		Sell      []CryptoZooNFTUserZoonInfoList
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Available = *abi.ConvertType(out[0], new([]CryptoZooNFTUserZoonInfoList)).(*[]CryptoZooNFTUserZoonInfoList)
	outstruct.Sell = *abi.ConvertType(out[1], new([]CryptoZooNFTUserZoonInfoList)).(*[]CryptoZooNFTUserZoonInfoList)

	return *outstruct, err

}

// ListAllNftByOwner is a free data retrieval call binding the contract method 0x6e0630f8.
//
// Solidity: function listAllNftByOwner(address _owner) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256,uint256)[] available, (((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256,uint256)[] sell)
func (_CryptoZooNFT *CryptoZooNFTSession) ListAllNftByOwner(_owner common.Address) (struct {
	Available []CryptoZooNFTUserZoonInfoList
	Sell      []CryptoZooNFTUserZoonInfoList
}, error) {
	return _CryptoZooNFT.Contract.ListAllNftByOwner(&_CryptoZooNFT.CallOpts, _owner)
}

// ListAllNftByOwner is a free data retrieval call binding the contract method 0x6e0630f8.
//
// Solidity: function listAllNftByOwner(address _owner) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256,uint256)[] available, (((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256,uint256)[] sell)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) ListAllNftByOwner(_owner common.Address) (struct {
	Available []CryptoZooNFTUserZoonInfoList
	Sell      []CryptoZooNFTUserZoonInfoList
}, error) {
	return _CryptoZooNFT.Contract.ListAllNftByOwner(&_CryptoZooNFT.CallOpts, _owner)
}

// ListNftByOwner is a free data retrieval call binding the contract method 0x7c3232a8.
//
// Solidity: function listNftByOwner(address _owner) view returns(((uint8,uint256,uint256,uint256,uint256,uint256),uint256)[] result)
func (_CryptoZooNFT *CryptoZooNFTCaller) ListNftByOwner(opts *bind.CallOpts, _owner common.Address) ([]CryptoZooNFTZoonInfo, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "listNftByOwner", _owner)

	if err != nil {
		return *new([]CryptoZooNFTZoonInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]CryptoZooNFTZoonInfo)).(*[]CryptoZooNFTZoonInfo)

	return out0, err

}

// ListNftByOwner is a free data retrieval call binding the contract method 0x7c3232a8.
//
// Solidity: function listNftByOwner(address _owner) view returns(((uint8,uint256,uint256,uint256,uint256,uint256),uint256)[] result)
func (_CryptoZooNFT *CryptoZooNFTSession) ListNftByOwner(_owner common.Address) ([]CryptoZooNFTZoonInfo, error) {
	return _CryptoZooNFT.Contract.ListNftByOwner(&_CryptoZooNFT.CallOpts, _owner)
}

// ListNftByOwner is a free data retrieval call binding the contract method 0x7c3232a8.
//
// Solidity: function listNftByOwner(address _owner) view returns(((uint8,uint256,uint256,uint256,uint256,uint256),uint256)[] result)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) ListNftByOwner(_owner common.Address) ([]CryptoZooNFTZoonInfo, error) {
	return _CryptoZooNFT.Contract.ListNftByOwner(&_CryptoZooNFT.CallOpts, _owner)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTSession) Manager() (common.Address, error) {
	return _CryptoZooNFT.Contract.Manager(&_CryptoZooNFT.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) Manager() (common.Address, error) {
	return _CryptoZooNFT.Contract.Manager(&_CryptoZooNFT.CallOpts)
}

// MarketsSize is a free data retrieval call binding the contract method 0x9d443201.
//
// Solidity: function marketsSize() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) MarketsSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "marketsSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketsSize is a free data retrieval call binding the contract method 0x9d443201.
//
// Solidity: function marketsSize() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) MarketsSize() (*big.Int, error) {
	return _CryptoZooNFT.Contract.MarketsSize(&_CryptoZooNFT.CallOpts)
}

// MarketsSize is a free data retrieval call binding the contract method 0x9d443201.
//
// Solidity: function marketsSize() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) MarketsSize() (*big.Int, error) {
	return _CryptoZooNFT.Contract.MarketsSize(&_CryptoZooNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoZooNFT *CryptoZooNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoZooNFT *CryptoZooNFTSession) Name() (string, error) {
	return _CryptoZooNFT.Contract.Name(&_CryptoZooNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) Name() (string, error) {
	return _CryptoZooNFT.Contract.Name(&_CryptoZooNFT.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0xf40e8471.
//
// Solidity: function orders(address _seller) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) Orders(opts *bind.CallOpts, _seller common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "orders", _seller)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Orders is a free data retrieval call binding the contract method 0xf40e8471.
//
// Solidity: function orders(address _seller) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) Orders(_seller common.Address) (*big.Int, error) {
	return _CryptoZooNFT.Contract.Orders(&_CryptoZooNFT.CallOpts, _seller)
}

// Orders is a free data retrieval call binding the contract method 0xf40e8471.
//
// Solidity: function orders(address _seller) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) Orders(_seller common.Address) (*big.Int, error) {
	return _CryptoZooNFT.Contract.Orders(&_CryptoZooNFT.CallOpts, _seller)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTSession) Owner() (common.Address, error) {
	return _CryptoZooNFT.Contract.Owner(&_CryptoZooNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) Owner() (common.Address, error) {
	return _CryptoZooNFT.Contract.Owner(&_CryptoZooNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CryptoZooNFT *CryptoZooNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CryptoZooNFT.Contract.OwnerOf(&_CryptoZooNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CryptoZooNFT.Contract.OwnerOf(&_CryptoZooNFT.CallOpts, tokenId)
}

// PriceEgg is a free data retrieval call binding the contract method 0x8a10008c.
//
// Solidity: function priceEgg() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) PriceEgg(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "priceEgg")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceEgg is a free data retrieval call binding the contract method 0x8a10008c.
//
// Solidity: function priceEgg() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) PriceEgg() (*big.Int, error) {
	return _CryptoZooNFT.Contract.PriceEgg(&_CryptoZooNFT.CallOpts)
}

// PriceEgg is a free data retrieval call binding the contract method 0x8a10008c.
//
// Solidity: function priceEgg() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) PriceEgg() (*big.Int, error) {
	return _CryptoZooNFT.Contract.PriceEgg(&_CryptoZooNFT.CallOpts)
}

// PriceEggDogCoin is a free data retrieval call binding the contract method 0x7057f377.
//
// Solidity: function priceEggDogCoin() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) PriceEggDogCoin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "priceEggDogCoin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceEggDogCoin is a free data retrieval call binding the contract method 0x7057f377.
//
// Solidity: function priceEggDogCoin() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) PriceEggDogCoin() (*big.Int, error) {
	return _CryptoZooNFT.Contract.PriceEggDogCoin(&_CryptoZooNFT.CallOpts)
}

// PriceEggDogCoin is a free data retrieval call binding the contract method 0x7057f377.
//
// Solidity: function priceEggDogCoin() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) PriceEggDogCoin() (*big.Int, error) {
	return _CryptoZooNFT.Contract.PriceEggDogCoin(&_CryptoZooNFT.CallOpts)
}

// PriceEggUsdt is a free data retrieval call binding the contract method 0x82483fb8.
//
// Solidity: function priceEggUsdt() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) PriceEggUsdt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "priceEggUsdt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceEggUsdt is a free data retrieval call binding the contract method 0x82483fb8.
//
// Solidity: function priceEggUsdt() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) PriceEggUsdt() (*big.Int, error) {
	return _CryptoZooNFT.Contract.PriceEggUsdt(&_CryptoZooNFT.CallOpts)
}

// PriceEggUsdt is a free data retrieval call binding the contract method 0x82483fb8.
//
// Solidity: function priceEggUsdt() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) PriceEggUsdt() (*big.Int, error) {
	return _CryptoZooNFT.Contract.PriceEggUsdt(&_CryptoZooNFT.CallOpts)
}

// SmallBoss is a free data retrieval call binding the contract method 0xd5e5c62b.
//
// Solidity: function smallBoss() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCaller) SmallBoss(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "smallBoss")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmallBoss is a free data retrieval call binding the contract method 0xd5e5c62b.
//
// Solidity: function smallBoss() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTSession) SmallBoss() (common.Address, error) {
	return _CryptoZooNFT.Contract.SmallBoss(&_CryptoZooNFT.CallOpts)
}

// SmallBoss is a free data retrieval call binding the contract method 0xd5e5c62b.
//
// Solidity: function smallBoss() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) SmallBoss() (common.Address, error) {
	return _CryptoZooNFT.Contract.SmallBoss(&_CryptoZooNFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CryptoZooNFT.Contract.SupportsInterface(&_CryptoZooNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CryptoZooNFT.Contract.SupportsInterface(&_CryptoZooNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoZooNFT *CryptoZooNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoZooNFT *CryptoZooNFTSession) Symbol() (string, error) {
	return _CryptoZooNFT.Contract.Symbol(&_CryptoZooNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) Symbol() (string, error) {
	return _CryptoZooNFT.Contract.Symbol(&_CryptoZooNFT.CallOpts)
}

// TokenSaleByIndex is a free data retrieval call binding the contract method 0x598a4fc8.
//
// Solidity: function tokenSaleByIndex(uint256 index) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) TokenSaleByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "tokenSaleByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenSaleByIndex is a free data retrieval call binding the contract method 0x598a4fc8.
//
// Solidity: function tokenSaleByIndex(uint256 index) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) TokenSaleByIndex(index *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.TokenSaleByIndex(&_CryptoZooNFT.CallOpts, index)
}

// TokenSaleByIndex is a free data retrieval call binding the contract method 0x598a4fc8.
//
// Solidity: function tokenSaleByIndex(uint256 index) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) TokenSaleByIndex(index *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.TokenSaleByIndex(&_CryptoZooNFT.CallOpts, index)
}

// TokenSaleOfOwnerByIndex is a free data retrieval call binding the contract method 0x078e62d4.
//
// Solidity: function tokenSaleOfOwnerByIndex(address _seller, uint256 index) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) TokenSaleOfOwnerByIndex(opts *bind.CallOpts, _seller common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "tokenSaleOfOwnerByIndex", _seller, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenSaleOfOwnerByIndex is a free data retrieval call binding the contract method 0x078e62d4.
//
// Solidity: function tokenSaleOfOwnerByIndex(address _seller, uint256 index) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) TokenSaleOfOwnerByIndex(_seller common.Address, index *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.TokenSaleOfOwnerByIndex(&_CryptoZooNFT.CallOpts, _seller, index)
}

// TokenSaleOfOwnerByIndex is a free data retrieval call binding the contract method 0x078e62d4.
//
// Solidity: function tokenSaleOfOwnerByIndex(address _seller, uint256 index) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) TokenSaleOfOwnerByIndex(_seller common.Address, index *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.TokenSaleOfOwnerByIndex(&_CryptoZooNFT.CallOpts, _seller, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CryptoZooNFT *CryptoZooNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CryptoZooNFT *CryptoZooNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CryptoZooNFT.Contract.TokenURI(&_CryptoZooNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CryptoZooNFT.Contract.TokenURI(&_CryptoZooNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) TotalSupply() (*big.Int, error) {
	return _CryptoZooNFT.Contract.TotalSupply(&_CryptoZooNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _CryptoZooNFT.Contract.TotalSupply(&_CryptoZooNFT.CallOpts)
}

// ZoonerERC20 is a free data retrieval call binding the contract method 0x63a28caa.
//
// Solidity: function zoonerERC20() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCaller) ZoonerERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "zoonerERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoonerERC20 is a free data retrieval call binding the contract method 0x63a28caa.
//
// Solidity: function zoonerERC20() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTSession) ZoonerERC20() (common.Address, error) {
	return _CryptoZooNFT.Contract.ZoonerERC20(&_CryptoZooNFT.CallOpts)
}

// ZoonerERC20 is a free data retrieval call binding the contract method 0x63a28caa.
//
// Solidity: function zoonerERC20() view returns(address)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) ZoonerERC20() (common.Address, error) {
	return _CryptoZooNFT.Contract.ZoonerERC20(&_CryptoZooNFT.CallOpts)
}

// ZoonerLevel is a free data retrieval call binding the contract method 0x108287c4.
//
// Solidity: function zoonerLevel(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCaller) ZoonerLevel(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CryptoZooNFT.contract.Call(opts, &out, "zoonerLevel", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZoonerLevel is a free data retrieval call binding the contract method 0x108287c4.
//
// Solidity: function zoonerLevel(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTSession) ZoonerLevel(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.ZoonerLevel(&_CryptoZooNFT.CallOpts, _tokenId)
}

// ZoonerLevel is a free data retrieval call binding the contract method 0x108287c4.
//
// Solidity: function zoonerLevel(uint256 _tokenId) view returns(uint256)
func (_CryptoZooNFT *CryptoZooNFTCallerSession) ZoonerLevel(_tokenId *big.Int) (*big.Int, error) {
	return _CryptoZooNFT.Contract.ZoonerLevel(&_CryptoZooNFT.CallOpts, _tokenId)
}

// AddRare is a paid mutator transaction binding the contract method 0x1ce331b8.
//
// Solidity: function addRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) AddRare(opts *bind.TransactOpts, _tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "addRare", _tokenids)
}

// AddRare is a paid mutator transaction binding the contract method 0x1ce331b8.
//
// Solidity: function addRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) AddRare(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.AddRare(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// AddRare is a paid mutator transaction binding the contract method 0x1ce331b8.
//
// Solidity: function addRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) AddRare(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.AddRare(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Approve(&_CryptoZooNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Approve(&_CryptoZooNFT.TransactOpts, to, tokenId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) CancelOrder(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "cancelOrder", _tokenId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) CancelOrder(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.CancelOrder(&_CryptoZooNFT.TransactOpts, _tokenId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) CancelOrder(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.CancelOrder(&_CryptoZooNFT.TransactOpts, _tokenId)
}

// DelBlack is a paid mutator transaction binding the contract method 0x8666debb.
//
// Solidity: function delBlack(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) DelBlack(opts *bind.TransactOpts, _tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "delBlack", _tokenid)
}

// DelBlack is a paid mutator transaction binding the contract method 0x8666debb.
//
// Solidity: function delBlack(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) DelBlack(_tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.DelBlack(&_CryptoZooNFT.TransactOpts, _tokenid)
}

// DelBlack is a paid mutator transaction binding the contract method 0x8666debb.
//
// Solidity: function delBlack(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) DelBlack(_tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.DelBlack(&_CryptoZooNFT.TransactOpts, _tokenid)
}

// Evolve is a paid mutator transaction binding the contract method 0xfbc4e9a0.
//
// Solidity: function evolve(uint256 _tokenId, address _owner, uint256 _dna) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) Evolve(opts *bind.TransactOpts, _tokenId *big.Int, _owner common.Address, _dna *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "evolve", _tokenId, _owner, _dna)
}

// Evolve is a paid mutator transaction binding the contract method 0xfbc4e9a0.
//
// Solidity: function evolve(uint256 _tokenId, address _owner, uint256 _dna) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) Evolve(_tokenId *big.Int, _owner common.Address, _dna *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Evolve(&_CryptoZooNFT.TransactOpts, _tokenId, _owner, _dna)
}

// Evolve is a paid mutator transaction binding the contract method 0xfbc4e9a0.
//
// Solidity: function evolve(uint256 _tokenId, address _owner, uint256 _dna) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) Evolve(_tokenId *big.Int, _owner common.Address, _dna *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Evolve(&_CryptoZooNFT.TransactOpts, _tokenId, _owner, _dna)
}

// Exp is a paid mutator transaction binding the contract method 0xf5f565f8.
//
// Solidity: function exp(uint256 _tokenId, uint256 _exp) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) Exp(opts *bind.TransactOpts, _tokenId *big.Int, _exp *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "exp", _tokenId, _exp)
}

// Exp is a paid mutator transaction binding the contract method 0xf5f565f8.
//
// Solidity: function exp(uint256 _tokenId, uint256 _exp) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) Exp(_tokenId *big.Int, _exp *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Exp(&_CryptoZooNFT.TransactOpts, _tokenId, _exp)
}

// Exp is a paid mutator transaction binding the contract method 0xf5f565f8.
//
// Solidity: function exp(uint256 _tokenId, uint256 _exp) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) Exp(_tokenId *big.Int, _exp *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Exp(&_CryptoZooNFT.TransactOpts, _tokenId, _exp)
}

// FillOrder is a paid mutator transaction binding the contract method 0x67b830ad.
//
// Solidity: function fillOrder(uint256 _tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) FillOrder(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "fillOrder", _tokenId)
}

// FillOrder is a paid mutator transaction binding the contract method 0x67b830ad.
//
// Solidity: function fillOrder(uint256 _tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) FillOrder(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.FillOrder(&_CryptoZooNFT.TransactOpts, _tokenId)
}

// FillOrder is a paid mutator transaction binding the contract method 0x67b830ad.
//
// Solidity: function fillOrder(uint256 _tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) FillOrder(_tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.FillOrder(&_CryptoZooNFT.TransactOpts, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "initialize", _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Initialize(&_CryptoZooNFT.TransactOpts, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Initialize(&_CryptoZooNFT.TransactOpts, _name, _symbol)
}

// LayEgg is a paid mutator transaction binding the contract method 0xda676678.
//
// Solidity: function layEgg(address receiver, uint8[] tribes) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) LayEgg(opts *bind.TransactOpts, receiver common.Address, tribes []uint8) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "layEgg", receiver, tribes)
}

// LayEgg is a paid mutator transaction binding the contract method 0xda676678.
//
// Solidity: function layEgg(address receiver, uint8[] tribes) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) LayEgg(receiver common.Address, tribes []uint8) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.LayEgg(&_CryptoZooNFT.TransactOpts, receiver, tribes)
}

// LayEgg is a paid mutator transaction binding the contract method 0xda676678.
//
// Solidity: function layEgg(address receiver, uint8[] tribes) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) LayEgg(receiver common.Address, tribes []uint8) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.LayEgg(&_CryptoZooNFT.TransactOpts, receiver, tribes)
}

// MulPhysical is a paid mutator transaction binding the contract method 0x8b98aaab.
//
// Solidity: function mulPhysical(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) MulPhysical(opts *bind.TransactOpts, _tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "mulPhysical", _tokenid)
}

// MulPhysical is a paid mutator transaction binding the contract method 0x8b98aaab.
//
// Solidity: function mulPhysical(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) MulPhysical(_tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.MulPhysical(&_CryptoZooNFT.TransactOpts, _tokenid)
}

// MulPhysical is a paid mutator transaction binding the contract method 0x8b98aaab.
//
// Solidity: function mulPhysical(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) MulPhysical(_tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.MulPhysical(&_CryptoZooNFT.TransactOpts, _tokenid)
}

// PayPhysical is a paid mutator transaction binding the contract method 0x4c8f3c63.
//
// Solidity: function payPhysical(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) PayPhysical(opts *bind.TransactOpts, _tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "payPhysical", _tokenid)
}

// PayPhysical is a paid mutator transaction binding the contract method 0x4c8f3c63.
//
// Solidity: function payPhysical(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) PayPhysical(_tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.PayPhysical(&_CryptoZooNFT.TransactOpts, _tokenid)
}

// PayPhysical is a paid mutator transaction binding the contract method 0x4c8f3c63.
//
// Solidity: function payPhysical(uint256 _tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) PayPhysical(_tokenid *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.PayPhysical(&_CryptoZooNFT.TransactOpts, _tokenid)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(uint256 _tokenId, uint256 _price) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) PlaceOrder(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "placeOrder", _tokenId, _price)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(uint256 _tokenId, uint256 _price) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) PlaceOrder(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.PlaceOrder(&_CryptoZooNFT.TransactOpts, _tokenId, _price)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x843f61e2.
//
// Solidity: function placeOrder(uint256 _tokenId, uint256 _price) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) PlaceOrder(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.PlaceOrder(&_CryptoZooNFT.TransactOpts, _tokenId, _price)
}

// RebasePhysicalTotal is a paid mutator transaction binding the contract method 0xcaf95bf5.
//
// Solidity: function rebasePhysicalTotal(uint256[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) RebasePhysicalTotal(opts *bind.TransactOpts, _tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "rebasePhysicalTotal", _tokenids)
}

// RebasePhysicalTotal is a paid mutator transaction binding the contract method 0xcaf95bf5.
//
// Solidity: function rebasePhysicalTotal(uint256[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) RebasePhysicalTotal(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.RebasePhysicalTotal(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// RebasePhysicalTotal is a paid mutator transaction binding the contract method 0xcaf95bf5.
//
// Solidity: function rebasePhysicalTotal(uint256[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) RebasePhysicalTotal(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.RebasePhysicalTotal(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// ReduceRare is a paid mutator transaction binding the contract method 0x6bcdb317.
//
// Solidity: function reduceRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) ReduceRare(opts *bind.TransactOpts, _tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "reduceRare", _tokenids)
}

// ReduceRare is a paid mutator transaction binding the contract method 0x6bcdb317.
//
// Solidity: function reduceRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) ReduceRare(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.ReduceRare(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// ReduceRare is a paid mutator transaction binding the contract method 0x6bcdb317.
//
// Solidity: function reduceRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) ReduceRare(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.ReduceRare(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CryptoZooNFT *CryptoZooNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.RenounceOwnership(&_CryptoZooNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.RenounceOwnership(&_CryptoZooNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SafeTransferFrom(&_CryptoZooNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SafeTransferFrom(&_CryptoZooNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SafeTransferFrom0(&_CryptoZooNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SafeTransferFrom0(&_CryptoZooNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetApprovalForAll(&_CryptoZooNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetApprovalForAll(&_CryptoZooNFT.TransactOpts, operator, approved)
}

// SetBlack is a paid mutator transaction binding the contract method 0x00436d79.
//
// Solidity: function setBlack(uint256[] tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetBlack(opts *bind.TransactOpts, tokenid []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setBlack", tokenid)
}

// SetBlack is a paid mutator transaction binding the contract method 0x00436d79.
//
// Solidity: function setBlack(uint256[] tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetBlack(tokenid []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetBlack(&_CryptoZooNFT.TransactOpts, tokenid)
}

// SetBlack is a paid mutator transaction binding the contract method 0x00436d79.
//
// Solidity: function setBlack(uint256[] tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetBlack(tokenid []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetBlack(&_CryptoZooNFT.TransactOpts, tokenid)
}

// SetERC20AndManager is a paid mutator transaction binding the contract method 0xb118d156.
//
// Solidity: function setERC20AndManager(address _manager, address _zoonerERC20) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetERC20AndManager(opts *bind.TransactOpts, _manager common.Address, _zoonerERC20 common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setERC20AndManager", _manager, _zoonerERC20)
}

// SetERC20AndManager is a paid mutator transaction binding the contract method 0xb118d156.
//
// Solidity: function setERC20AndManager(address _manager, address _zoonerERC20) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetERC20AndManager(_manager common.Address, _zoonerERC20 common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetERC20AndManager(&_CryptoZooNFT.TransactOpts, _manager, _zoonerERC20)
}

// SetERC20AndManager is a paid mutator transaction binding the contract method 0xb118d156.
//
// Solidity: function setERC20AndManager(address _manager, address _zoonerERC20) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetERC20AndManager(_manager common.Address, _zoonerERC20 common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetERC20AndManager(&_CryptoZooNFT.TransactOpts, _manager, _zoonerERC20)
}

// SetPhysicalTotal is a paid mutator transaction binding the contract method 0x826d5866.
//
// Solidity: function setPhysicalTotal(uint256[] _tokenids, uint256[] _totals) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetPhysicalTotal(opts *bind.TransactOpts, _tokenids []*big.Int, _totals []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setPhysicalTotal", _tokenids, _totals)
}

// SetPhysicalTotal is a paid mutator transaction binding the contract method 0x826d5866.
//
// Solidity: function setPhysicalTotal(uint256[] _tokenids, uint256[] _totals) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetPhysicalTotal(_tokenids []*big.Int, _totals []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetPhysicalTotal(&_CryptoZooNFT.TransactOpts, _tokenids, _totals)
}

// SetPhysicalTotal is a paid mutator transaction binding the contract method 0x826d5866.
//
// Solidity: function setPhysicalTotal(uint256[] _tokenids, uint256[] _totals) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetPhysicalTotal(_tokenids []*big.Int, _totals []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetPhysicalTotal(&_CryptoZooNFT.TransactOpts, _tokenids, _totals)
}

// SetRare is a paid mutator transaction binding the contract method 0xc90fc418.
//
// Solidity: function setRare(uint256 _tokenid, uint256 _newRate) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetRare(opts *bind.TransactOpts, _tokenid *big.Int, _newRate *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setRare", _tokenid, _newRate)
}

// SetRare is a paid mutator transaction binding the contract method 0xc90fc418.
//
// Solidity: function setRare(uint256 _tokenid, uint256 _newRate) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetRare(_tokenid *big.Int, _newRate *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetRare(&_CryptoZooNFT.TransactOpts, _tokenid, _newRate)
}

// SetRare is a paid mutator transaction binding the contract method 0xc90fc418.
//
// Solidity: function setRare(uint256 _tokenid, uint256 _newRate) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetRare(_tokenid *big.Int, _newRate *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetRare(&_CryptoZooNFT.TransactOpts, _tokenid, _newRate)
}

// SetReversion is a paid mutator transaction binding the contract method 0x6fff5eca.
//
// Solidity: function setReversion(uint256[] tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetReversion(opts *bind.TransactOpts, tokenid []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setReversion", tokenid)
}

// SetReversion is a paid mutator transaction binding the contract method 0x6fff5eca.
//
// Solidity: function setReversion(uint256[] tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetReversion(tokenid []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetReversion(&_CryptoZooNFT.TransactOpts, tokenid)
}

// SetReversion is a paid mutator transaction binding the contract method 0x6fff5eca.
//
// Solidity: function setReversion(uint256[] tokenid) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetReversion(tokenid []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetReversion(&_CryptoZooNFT.TransactOpts, tokenid)
}

// SetSmallBoss is a paid mutator transaction binding the contract method 0x19f65ac2.
//
// Solidity: function setSmallBoss(address _smallBoss) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetSmallBoss(opts *bind.TransactOpts, _smallBoss common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setSmallBoss", _smallBoss)
}

// SetSmallBoss is a paid mutator transaction binding the contract method 0x19f65ac2.
//
// Solidity: function setSmallBoss(address _smallBoss) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetSmallBoss(_smallBoss common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetSmallBoss(&_CryptoZooNFT.TransactOpts, _smallBoss)
}

// SetSmallBoss is a paid mutator transaction binding the contract method 0x19f65ac2.
//
// Solidity: function setSmallBoss(address _smallBoss) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetSmallBoss(_smallBoss common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetSmallBoss(&_CryptoZooNFT.TransactOpts, _smallBoss)
}

// SetSync is a paid mutator transaction binding the contract method 0xf915c64c.
//
// Solidity: function setSync(uint256[] tokenIds) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetSync(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setSync", tokenIds)
}

// SetSync is a paid mutator transaction binding the contract method 0xf915c64c.
//
// Solidity: function setSync(uint256[] tokenIds) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetSync(tokenIds []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetSync(&_CryptoZooNFT.TransactOpts, tokenIds)
}

// SetSync is a paid mutator transaction binding the contract method 0xf915c64c.
//
// Solidity: function setSync(uint256[] tokenIds) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetSync(tokenIds []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetSync(&_CryptoZooNFT.TransactOpts, tokenIds)
}

// SetUserRare is a paid mutator transaction binding the contract method 0x09c806b5.
//
// Solidity: function setUserRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetUserRare(opts *bind.TransactOpts, _tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setUserRare", _tokenids)
}

// SetUserRare is a paid mutator transaction binding the contract method 0x09c806b5.
//
// Solidity: function setUserRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetUserRare(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetUserRare(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// SetUserRare is a paid mutator transaction binding the contract method 0x09c806b5.
//
// Solidity: function setUserRare(uint128[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetUserRare(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetUserRare(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// SetUserSkill is a paid mutator transaction binding the contract method 0x70b66c33.
//
// Solidity: function setUserSkill(uint256[] _tokenids, uint256[] _skills) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) SetUserSkill(opts *bind.TransactOpts, _tokenids []*big.Int, _skills []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "setUserSkill", _tokenids, _skills)
}

// SetUserSkill is a paid mutator transaction binding the contract method 0x70b66c33.
//
// Solidity: function setUserSkill(uint256[] _tokenids, uint256[] _skills) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) SetUserSkill(_tokenids []*big.Int, _skills []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetUserSkill(&_CryptoZooNFT.TransactOpts, _tokenids, _skills)
}

// SetUserSkill is a paid mutator transaction binding the contract method 0x70b66c33.
//
// Solidity: function setUserSkill(uint256[] _tokenids, uint256[] _skills) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) SetUserSkill(_tokenids []*big.Int, _skills []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.SetUserSkill(&_CryptoZooNFT.TransactOpts, _tokenids, _skills)
}

// Synthesis is a paid mutator transaction binding the contract method 0x75608c66.
//
// Solidity: function synthesis(uint256[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) Synthesis(opts *bind.TransactOpts, _tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "synthesis", _tokenids)
}

// Synthesis is a paid mutator transaction binding the contract method 0x75608c66.
//
// Solidity: function synthesis(uint256[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) Synthesis(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Synthesis(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// Synthesis is a paid mutator transaction binding the contract method 0x75608c66.
//
// Solidity: function synthesis(uint256[] _tokenids) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) Synthesis(_tokenids []*big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Synthesis(&_CryptoZooNFT.TransactOpts, _tokenids)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.TransferFrom(&_CryptoZooNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.TransferFrom(&_CryptoZooNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.TransferOwnership(&_CryptoZooNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.TransferOwnership(&_CryptoZooNFT.TransactOpts, newOwner)
}

// Working is a paid mutator transaction binding the contract method 0xd9b3f5fc.
//
// Solidity: function working(uint256 _tokenId, uint256 _time) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactor) Working(opts *bind.TransactOpts, _tokenId *big.Int, _time *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.contract.Transact(opts, "working", _tokenId, _time)
}

// Working is a paid mutator transaction binding the contract method 0xd9b3f5fc.
//
// Solidity: function working(uint256 _tokenId, uint256 _time) returns()
func (_CryptoZooNFT *CryptoZooNFTSession) Working(_tokenId *big.Int, _time *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Working(&_CryptoZooNFT.TransactOpts, _tokenId, _time)
}

// Working is a paid mutator transaction binding the contract method 0xd9b3f5fc.
//
// Solidity: function working(uint256 _tokenId, uint256 _time) returns()
func (_CryptoZooNFT *CryptoZooNFTTransactorSession) Working(_tokenId *big.Int, _time *big.Int) (*types.Transaction, error) {
	return _CryptoZooNFT.Contract.Working(&_CryptoZooNFT.TransactOpts, _tokenId, _time)
}

// CryptoZooNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CryptoZooNFT contract.
type CryptoZooNFTApprovalIterator struct {
	Event *CryptoZooNFTApproval // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTApproval)
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
		it.Event = new(CryptoZooNFTApproval)
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
func (it *CryptoZooNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTApproval represents a Approval event raised by the CryptoZooNFT contract.
type CryptoZooNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CryptoZooNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTApprovalIterator{contract: _CryptoZooNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTApproval)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseApproval(log types.Log) (*CryptoZooNFTApproval, error) {
	event := new(CryptoZooNFTApproval)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CryptoZooNFT contract.
type CryptoZooNFTApprovalForAllIterator struct {
	Event *CryptoZooNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTApprovalForAll)
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
		it.Event = new(CryptoZooNFTApprovalForAll)
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
func (it *CryptoZooNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTApprovalForAll represents a ApprovalForAll event raised by the CryptoZooNFT contract.
type CryptoZooNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CryptoZooNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTApprovalForAllIterator{contract: _CryptoZooNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTApprovalForAll)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseApprovalForAll(log types.Log) (*CryptoZooNFTApprovalForAll, error) {
	event := new(CryptoZooNFTApprovalForAll)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the CryptoZooNFT contract.
type CryptoZooNFTBuyIterator struct {
	Event *CryptoZooNFTBuy // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTBuy)
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
		it.Event = new(CryptoZooNFTBuy)
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
func (it *CryptoZooNFTBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTBuy represents a Buy event raised by the CryptoZooNFT contract.
type CryptoZooNFTBuy struct {
	TokenId *big.Int
	Buyer   common.Address
	Seller  common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0x07f87664c10527a8207b443cccf57f3c20f25bd1165eeae416be79890a35892c.
//
// Solidity: event Buy(uint256 indexed tokenId, address buyer, address seller, uint256 price)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterBuy(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTBuyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Buy", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTBuyIterator{contract: _CryptoZooNFT.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0x07f87664c10527a8207b443cccf57f3c20f25bd1165eeae416be79890a35892c.
//
// Solidity: event Buy(uint256 indexed tokenId, address buyer, address seller, uint256 price)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTBuy, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Buy", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTBuy)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0x07f87664c10527a8207b443cccf57f3c20f25bd1165eeae416be79890a35892c.
//
// Solidity: event Buy(uint256 indexed tokenId, address buyer, address seller, uint256 price)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseBuy(log types.Log) (*CryptoZooNFTBuy, error) {
	event := new(CryptoZooNFTBuy)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTCancelOrderIterator is returned from FilterCancelOrder and is used to iterate over the raw logs and unpacked data for CancelOrder events raised by the CryptoZooNFT contract.
type CryptoZooNFTCancelOrderIterator struct {
	Event *CryptoZooNFTCancelOrder // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTCancelOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTCancelOrder)
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
		it.Event = new(CryptoZooNFTCancelOrder)
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
func (it *CryptoZooNFTCancelOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTCancelOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTCancelOrder represents a CancelOrder event raised by the CryptoZooNFT contract.
type CryptoZooNFTCancelOrder struct {
	TokenId *big.Int
	Seller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelOrder is a free log retrieval operation binding the contract event 0x22369ba22944aadf9e9d6f4c51462417a50ea7876b9c62c7c46b5522e9c672cc.
//
// Solidity: event CancelOrder(uint256 indexed tokenId, address seller)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterCancelOrder(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTCancelOrderIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "CancelOrder", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTCancelOrderIterator{contract: _CryptoZooNFT.contract, event: "CancelOrder", logs: logs, sub: sub}, nil
}

// WatchCancelOrder is a free log subscription operation binding the contract event 0x22369ba22944aadf9e9d6f4c51462417a50ea7876b9c62c7c46b5522e9c672cc.
//
// Solidity: event CancelOrder(uint256 indexed tokenId, address seller)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchCancelOrder(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTCancelOrder, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "CancelOrder", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTCancelOrder)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "CancelOrder", log); err != nil {
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

// ParseCancelOrder is a log parse operation binding the contract event 0x22369ba22944aadf9e9d6f4c51462417a50ea7876b9c62c7c46b5522e9c672cc.
//
// Solidity: event CancelOrder(uint256 indexed tokenId, address seller)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseCancelOrder(log types.Log) (*CryptoZooNFTCancelOrder, error) {
	event := new(CryptoZooNFTCancelOrder)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "CancelOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTEvolveIterator is returned from FilterEvolve and is used to iterate over the raw logs and unpacked data for Evolve events raised by the CryptoZooNFT contract.
type CryptoZooNFTEvolveIterator struct {
	Event *CryptoZooNFTEvolve // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTEvolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTEvolve)
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
		it.Event = new(CryptoZooNFTEvolve)
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
func (it *CryptoZooNFTEvolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTEvolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTEvolve represents a Evolve event raised by the CryptoZooNFT contract.
type CryptoZooNFTEvolve struct {
	TokenId *big.Int
	Dna     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEvolve is a free log retrieval operation binding the contract event 0xaed9826fdb59b4dced9e7dfd3305624745170ecc0ae281bb1f3ca5a9279f14cb.
//
// Solidity: event Evolve(uint256 indexed tokenId, uint256 dna)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterEvolve(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTEvolveIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Evolve", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTEvolveIterator{contract: _CryptoZooNFT.contract, event: "Evolve", logs: logs, sub: sub}, nil
}

// WatchEvolve is a free log subscription operation binding the contract event 0xaed9826fdb59b4dced9e7dfd3305624745170ecc0ae281bb1f3ca5a9279f14cb.
//
// Solidity: event Evolve(uint256 indexed tokenId, uint256 dna)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchEvolve(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTEvolve, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Evolve", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTEvolve)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Evolve", log); err != nil {
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

// ParseEvolve is a log parse operation binding the contract event 0xaed9826fdb59b4dced9e7dfd3305624745170ecc0ae281bb1f3ca5a9279f14cb.
//
// Solidity: event Evolve(uint256 indexed tokenId, uint256 dna)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseEvolve(log types.Log) (*CryptoZooNFTEvolve, error) {
	event := new(CryptoZooNFTEvolve)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Evolve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTExpIterator is returned from FilterExp and is used to iterate over the raw logs and unpacked data for Exp events raised by the CryptoZooNFT contract.
type CryptoZooNFTExpIterator struct {
	Event *CryptoZooNFTExp // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTExpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTExp)
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
		it.Event = new(CryptoZooNFTExp)
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
func (it *CryptoZooNFTExpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTExpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTExp represents a Exp event raised by the CryptoZooNFT contract.
type CryptoZooNFTExp struct {
	TokenId *big.Int
	Exp     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExp is a free log retrieval operation binding the contract event 0x3056b9cd784671e7a0f5d6f930db4d14e1e876a87360c871cc4d88b9f54c44b2.
//
// Solidity: event Exp(uint256 indexed tokenId, uint256 exp)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterExp(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTExpIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Exp", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTExpIterator{contract: _CryptoZooNFT.contract, event: "Exp", logs: logs, sub: sub}, nil
}

// WatchExp is a free log subscription operation binding the contract event 0x3056b9cd784671e7a0f5d6f930db4d14e1e876a87360c871cc4d88b9f54c44b2.
//
// Solidity: event Exp(uint256 indexed tokenId, uint256 exp)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchExp(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTExp, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Exp", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTExp)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Exp", log); err != nil {
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

// ParseExp is a log parse operation binding the contract event 0x3056b9cd784671e7a0f5d6f930db4d14e1e876a87360c871cc4d88b9f54c44b2.
//
// Solidity: event Exp(uint256 indexed tokenId, uint256 exp)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseExp(log types.Log) (*CryptoZooNFTExp, error) {
	event := new(CryptoZooNFTExp)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Exp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTFillOrderIterator is returned from FilterFillOrder and is used to iterate over the raw logs and unpacked data for FillOrder events raised by the CryptoZooNFT contract.
type CryptoZooNFTFillOrderIterator struct {
	Event *CryptoZooNFTFillOrder // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTFillOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTFillOrder)
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
		it.Event = new(CryptoZooNFTFillOrder)
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
func (it *CryptoZooNFTFillOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTFillOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTFillOrder represents a FillOrder event raised by the CryptoZooNFT contract.
type CryptoZooNFTFillOrder struct {
	TokenId *big.Int
	Seller  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFillOrder is a free log retrieval operation binding the contract event 0xe172f6058a4b7a5493aea0902d867636f05695ffeb71801046c3f4e5002caa97.
//
// Solidity: event FillOrder(uint256 indexed tokenId, address seller)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterFillOrder(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTFillOrderIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "FillOrder", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTFillOrderIterator{contract: _CryptoZooNFT.contract, event: "FillOrder", logs: logs, sub: sub}, nil
}

// WatchFillOrder is a free log subscription operation binding the contract event 0xe172f6058a4b7a5493aea0902d867636f05695ffeb71801046c3f4e5002caa97.
//
// Solidity: event FillOrder(uint256 indexed tokenId, address seller)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchFillOrder(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTFillOrder, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "FillOrder", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTFillOrder)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "FillOrder", log); err != nil {
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

// ParseFillOrder is a log parse operation binding the contract event 0xe172f6058a4b7a5493aea0902d867636f05695ffeb71801046c3f4e5002caa97.
//
// Solidity: event FillOrder(uint256 indexed tokenId, address seller)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseFillOrder(log types.Log) (*CryptoZooNFTFillOrder, error) {
	event := new(CryptoZooNFTFillOrder)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "FillOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CryptoZooNFT contract.
type CryptoZooNFTInitializedIterator struct {
	Event *CryptoZooNFTInitialized // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTInitialized)
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
		it.Event = new(CryptoZooNFTInitialized)
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
func (it *CryptoZooNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTInitialized represents a Initialized event raised by the CryptoZooNFT contract.
type CryptoZooNFTInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterInitialized(opts *bind.FilterOpts) (*CryptoZooNFTInitializedIterator, error) {

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTInitializedIterator{contract: _CryptoZooNFT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTInitialized)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseInitialized(log types.Log) (*CryptoZooNFTInitialized, error) {
	event := new(CryptoZooNFTInitialized)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTLayEggIterator is returned from FilterLayEgg and is used to iterate over the raw logs and unpacked data for LayEgg events raised by the CryptoZooNFT contract.
type CryptoZooNFTLayEggIterator struct {
	Event *CryptoZooNFTLayEgg // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTLayEggIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTLayEgg)
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
		it.Event = new(CryptoZooNFTLayEgg)
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
func (it *CryptoZooNFTLayEggIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTLayEggIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTLayEgg represents a LayEgg event raised by the CryptoZooNFT contract.
type CryptoZooNFTLayEgg struct {
	TokenId *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLayEgg is a free log retrieval operation binding the contract event 0xa6e85fb9fbf4d686fbbb85f52a9ed2b7d1782836d6541b9efbef16807d390b73.
//
// Solidity: event LayEgg(uint256 indexed tokenId, address buyer)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterLayEgg(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTLayEggIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "LayEgg", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTLayEggIterator{contract: _CryptoZooNFT.contract, event: "LayEgg", logs: logs, sub: sub}, nil
}

// WatchLayEgg is a free log subscription operation binding the contract event 0xa6e85fb9fbf4d686fbbb85f52a9ed2b7d1782836d6541b9efbef16807d390b73.
//
// Solidity: event LayEgg(uint256 indexed tokenId, address buyer)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchLayEgg(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTLayEgg, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "LayEgg", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTLayEgg)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "LayEgg", log); err != nil {
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

// ParseLayEgg is a log parse operation binding the contract event 0xa6e85fb9fbf4d686fbbb85f52a9ed2b7d1782836d6541b9efbef16807d390b73.
//
// Solidity: event LayEgg(uint256 indexed tokenId, address buyer)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseLayEgg(log types.Log) (*CryptoZooNFTLayEgg, error) {
	event := new(CryptoZooNFTLayEgg)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "LayEgg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CryptoZooNFT contract.
type CryptoZooNFTOwnershipTransferredIterator struct {
	Event *CryptoZooNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTOwnershipTransferred)
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
		it.Event = new(CryptoZooNFTOwnershipTransferred)
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
func (it *CryptoZooNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTOwnershipTransferred represents a OwnershipTransferred event raised by the CryptoZooNFT contract.
type CryptoZooNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CryptoZooNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTOwnershipTransferredIterator{contract: _CryptoZooNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTOwnershipTransferred)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseOwnershipTransferred(log types.Log) (*CryptoZooNFTOwnershipTransferred, error) {
	event := new(CryptoZooNFTOwnershipTransferred)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTPlaceOrderIterator is returned from FilterPlaceOrder and is used to iterate over the raw logs and unpacked data for PlaceOrder events raised by the CryptoZooNFT contract.
type CryptoZooNFTPlaceOrderIterator struct {
	Event *CryptoZooNFTPlaceOrder // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTPlaceOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTPlaceOrder)
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
		it.Event = new(CryptoZooNFTPlaceOrder)
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
func (it *CryptoZooNFTPlaceOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTPlaceOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTPlaceOrder represents a PlaceOrder event raised by the CryptoZooNFT contract.
type CryptoZooNFTPlaceOrder struct {
	TokenId *big.Int
	Seller  common.Address
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPlaceOrder is a free log retrieval operation binding the contract event 0x88ba4afa23f9d5bf6fc21903b32126bc507e001fd6fd00affc837adeb5d4a607.
//
// Solidity: event PlaceOrder(uint256 indexed tokenId, address seller, uint256 price)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterPlaceOrder(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTPlaceOrderIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "PlaceOrder", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTPlaceOrderIterator{contract: _CryptoZooNFT.contract, event: "PlaceOrder", logs: logs, sub: sub}, nil
}

// WatchPlaceOrder is a free log subscription operation binding the contract event 0x88ba4afa23f9d5bf6fc21903b32126bc507e001fd6fd00affc837adeb5d4a607.
//
// Solidity: event PlaceOrder(uint256 indexed tokenId, address seller, uint256 price)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchPlaceOrder(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTPlaceOrder, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "PlaceOrder", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTPlaceOrder)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "PlaceOrder", log); err != nil {
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

// ParsePlaceOrder is a log parse operation binding the contract event 0x88ba4afa23f9d5bf6fc21903b32126bc507e001fd6fd00affc837adeb5d4a607.
//
// Solidity: event PlaceOrder(uint256 indexed tokenId, address seller, uint256 price)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParsePlaceOrder(log types.Log) (*CryptoZooNFTPlaceOrder, error) {
	event := new(CryptoZooNFTPlaceOrder)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "PlaceOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTSynthEventIterator is returned from FilterSynthEvent and is used to iterate over the raw logs and unpacked data for SynthEvent events raised by the CryptoZooNFT contract.
type CryptoZooNFTSynthEventIterator struct {
	Event *CryptoZooNFTSynthEvent // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTSynthEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTSynthEvent)
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
		it.Event = new(CryptoZooNFTSynthEvent)
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
func (it *CryptoZooNFTSynthEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTSynthEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTSynthEvent represents a SynthEvent event raised by the CryptoZooNFT contract.
type CryptoZooNFTSynthEvent struct {
	Sender  common.Address
	Tokenid *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSynthEvent is a free log retrieval operation binding the contract event 0xc280e84c64bd0bfb6cff73c04a4e139535020694698d11f1065222206689fdbe.
//
// Solidity: event SynthEvent(address _sender, uint256 _tokenid)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterSynthEvent(opts *bind.FilterOpts) (*CryptoZooNFTSynthEventIterator, error) {

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "SynthEvent")
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTSynthEventIterator{contract: _CryptoZooNFT.contract, event: "SynthEvent", logs: logs, sub: sub}, nil
}

// WatchSynthEvent is a free log subscription operation binding the contract event 0xc280e84c64bd0bfb6cff73c04a4e139535020694698d11f1065222206689fdbe.
//
// Solidity: event SynthEvent(address _sender, uint256 _tokenid)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchSynthEvent(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTSynthEvent) (event.Subscription, error) {

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "SynthEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTSynthEvent)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "SynthEvent", log); err != nil {
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

// ParseSynthEvent is a log parse operation binding the contract event 0xc280e84c64bd0bfb6cff73c04a4e139535020694698d11f1065222206689fdbe.
//
// Solidity: event SynthEvent(address _sender, uint256 _tokenid)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseSynthEvent(log types.Log) (*CryptoZooNFTSynthEvent, error) {
	event := new(CryptoZooNFTSynthEvent)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "SynthEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CryptoZooNFT contract.
type CryptoZooNFTTransferIterator struct {
	Event *CryptoZooNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTTransfer)
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
		it.Event = new(CryptoZooNFTTransfer)
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
func (it *CryptoZooNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTTransfer represents a Transfer event raised by the CryptoZooNFT contract.
type CryptoZooNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CryptoZooNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTTransferIterator{contract: _CryptoZooNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTTransfer)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseTransfer(log types.Log) (*CryptoZooNFTTransfer, error) {
	event := new(CryptoZooNFTTransfer)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the CryptoZooNFT contract.
type CryptoZooNFTUpdatePriceIterator struct {
	Event *CryptoZooNFTUpdatePrice // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTUpdatePrice)
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
		it.Event = new(CryptoZooNFTUpdatePrice)
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
func (it *CryptoZooNFTUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTUpdatePrice represents a UpdatePrice event raised by the CryptoZooNFT contract.
type CryptoZooNFTUpdatePrice struct {
	TokenId  *big.Int
	Seller   common.Address
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0xb2ba2579f847b115aea408a81cca74e27e514b4563c1251ebe78f8c5e9ad7176.
//
// Solidity: event UpdatePrice(uint256 indexed tokenId, address seller, uint256 newPrice)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterUpdatePrice(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTUpdatePriceIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "UpdatePrice", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTUpdatePriceIterator{contract: _CryptoZooNFT.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0xb2ba2579f847b115aea408a81cca74e27e514b4563c1251ebe78f8c5e9ad7176.
//
// Solidity: event UpdatePrice(uint256 indexed tokenId, address seller, uint256 newPrice)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTUpdatePrice, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "UpdatePrice", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTUpdatePrice)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
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

// ParseUpdatePrice is a log parse operation binding the contract event 0xb2ba2579f847b115aea408a81cca74e27e514b4563c1251ebe78f8c5e9ad7176.
//
// Solidity: event UpdatePrice(uint256 indexed tokenId, address seller, uint256 newPrice)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseUpdatePrice(log types.Log) (*CryptoZooNFTUpdatePrice, error) {
	event := new(CryptoZooNFTUpdatePrice)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTUpdateTribeIterator is returned from FilterUpdateTribe and is used to iterate over the raw logs and unpacked data for UpdateTribe events raised by the CryptoZooNFT contract.
type CryptoZooNFTUpdateTribeIterator struct {
	Event *CryptoZooNFTUpdateTribe // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTUpdateTribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTUpdateTribe)
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
		it.Event = new(CryptoZooNFTUpdateTribe)
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
func (it *CryptoZooNFTUpdateTribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTUpdateTribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTUpdateTribe represents a UpdateTribe event raised by the CryptoZooNFT contract.
type CryptoZooNFTUpdateTribe struct {
	TokenId *big.Int
	Tribe   uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTribe is a free log retrieval operation binding the contract event 0x77b4e20830068beb0bfa8065352474e6a4d6050ac7c77949eec24f5a46025b4e.
//
// Solidity: event UpdateTribe(uint256 indexed tokenId, uint8 tribe)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterUpdateTribe(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTUpdateTribeIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "UpdateTribe", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTUpdateTribeIterator{contract: _CryptoZooNFT.contract, event: "UpdateTribe", logs: logs, sub: sub}, nil
}

// WatchUpdateTribe is a free log subscription operation binding the contract event 0x77b4e20830068beb0bfa8065352474e6a4d6050ac7c77949eec24f5a46025b4e.
//
// Solidity: event UpdateTribe(uint256 indexed tokenId, uint8 tribe)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchUpdateTribe(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTUpdateTribe, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "UpdateTribe", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTUpdateTribe)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "UpdateTribe", log); err != nil {
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

// ParseUpdateTribe is a log parse operation binding the contract event 0x77b4e20830068beb0bfa8065352474e6a4d6050ac7c77949eec24f5a46025b4e.
//
// Solidity: event UpdateTribe(uint256 indexed tokenId, uint8 tribe)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseUpdateTribe(log types.Log) (*CryptoZooNFTUpdateTribe, error) {
	event := new(CryptoZooNFTUpdateTribe)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "UpdateTribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTUpgradeGenerationIterator is returned from FilterUpgradeGeneration and is used to iterate over the raw logs and unpacked data for UpgradeGeneration events raised by the CryptoZooNFT contract.
type CryptoZooNFTUpgradeGenerationIterator struct {
	Event *CryptoZooNFTUpgradeGeneration // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTUpgradeGenerationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTUpgradeGeneration)
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
		it.Event = new(CryptoZooNFTUpgradeGeneration)
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
func (it *CryptoZooNFTUpgradeGenerationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTUpgradeGenerationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTUpgradeGeneration represents a UpgradeGeneration event raised by the CryptoZooNFT contract.
type CryptoZooNFTUpgradeGeneration struct {
	TokenId       *big.Int
	NewGeneration *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpgradeGeneration is a free log retrieval operation binding the contract event 0xf51cc1eb6b47c35b6cb187ef8f47f5cdc8ef340406d9e52c1c9238bcd3487b2a.
//
// Solidity: event UpgradeGeneration(uint256 indexed tokenId, uint256 newGeneration)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterUpgradeGeneration(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTUpgradeGenerationIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "UpgradeGeneration", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTUpgradeGenerationIterator{contract: _CryptoZooNFT.contract, event: "UpgradeGeneration", logs: logs, sub: sub}, nil
}

// WatchUpgradeGeneration is a free log subscription operation binding the contract event 0xf51cc1eb6b47c35b6cb187ef8f47f5cdc8ef340406d9e52c1c9238bcd3487b2a.
//
// Solidity: event UpgradeGeneration(uint256 indexed tokenId, uint256 newGeneration)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchUpgradeGeneration(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTUpgradeGeneration, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "UpgradeGeneration", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTUpgradeGeneration)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "UpgradeGeneration", log); err != nil {
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

// ParseUpgradeGeneration is a log parse operation binding the contract event 0xf51cc1eb6b47c35b6cb187ef8f47f5cdc8ef340406d9e52c1c9238bcd3487b2a.
//
// Solidity: event UpgradeGeneration(uint256 indexed tokenId, uint256 newGeneration)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseUpgradeGeneration(log types.Log) (*CryptoZooNFTUpgradeGeneration, error) {
	event := new(CryptoZooNFTUpgradeGeneration)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "UpgradeGeneration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoZooNFTWorkingIterator is returned from FilterWorking and is used to iterate over the raw logs and unpacked data for Working events raised by the CryptoZooNFT contract.
type CryptoZooNFTWorkingIterator struct {
	Event *CryptoZooNFTWorking // Event containing the contract specifics and raw log

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
func (it *CryptoZooNFTWorkingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoZooNFTWorking)
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
		it.Event = new(CryptoZooNFTWorking)
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
func (it *CryptoZooNFTWorkingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoZooNFTWorkingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoZooNFTWorking represents a Working event raised by the CryptoZooNFT contract.
type CryptoZooNFTWorking struct {
	TokenId *big.Int
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWorking is a free log retrieval operation binding the contract event 0xfd01fe220c4c972c579a382864689384348521ab699de06e0a1f48f86e6964b2.
//
// Solidity: event Working(uint256 indexed tokenId, uint256 time)
func (_CryptoZooNFT *CryptoZooNFTFilterer) FilterWorking(opts *bind.FilterOpts, tokenId []*big.Int) (*CryptoZooNFTWorkingIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.FilterLogs(opts, "Working", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CryptoZooNFTWorkingIterator{contract: _CryptoZooNFT.contract, event: "Working", logs: logs, sub: sub}, nil
}

// WatchWorking is a free log subscription operation binding the contract event 0xfd01fe220c4c972c579a382864689384348521ab699de06e0a1f48f86e6964b2.
//
// Solidity: event Working(uint256 indexed tokenId, uint256 time)
func (_CryptoZooNFT *CryptoZooNFTFilterer) WatchWorking(opts *bind.WatchOpts, sink chan<- *CryptoZooNFTWorking, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CryptoZooNFT.contract.WatchLogs(opts, "Working", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoZooNFTWorking)
				if err := _CryptoZooNFT.contract.UnpackLog(event, "Working", log); err != nil {
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

// ParseWorking is a log parse operation binding the contract event 0xfd01fe220c4c972c579a382864689384348521ab699de06e0a1f48f86e6964b2.
//
// Solidity: event Working(uint256 indexed tokenId, uint256 time)
func (_CryptoZooNFT *CryptoZooNFTFilterer) ParseWorking(log types.Log) (*CryptoZooNFTWorking, error) {
	event := new(CryptoZooNFTWorking)
	if err := _CryptoZooNFT.contract.UnpackLog(event, "Working", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
