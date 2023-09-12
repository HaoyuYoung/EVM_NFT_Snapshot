// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// ERC721ATokenOwnership is an auto generated low-level Go binding around an user-defined struct.
type ERC721ATokenOwnership struct {
	Addr           common.Address
	StartTimestamp uint64
}

// NFTMetaData contains all meta data concerning the NFT contract.
var NFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxBatchSize_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectionSize_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountForAuctionAndDev_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountForDevs_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUCTION_DROP_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTION_DROP_PER_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTION_END_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTION_PRICE_CURVE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AUCTION_START_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowlist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowlistMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountForAuctionAndDev\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountForDevs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"auctionMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"devMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"mintlistPriceWei\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicPriceWei\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"publicSaleStartTime\",\"type\":\"uint32\"}],\"name\":\"endAuctionAndSetupNonAuctionSaleInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleStartTime\",\"type\":\"uint256\"}],\"name\":\"getAuctionPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getOwnershipData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structERC721A.TokenOwnership\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"publicPriceWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicSaleKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicSaleStartTime\",\"type\":\"uint256\"}],\"name\":\"isPublicSaleOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPerAddressDuringMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOwnerToExplicitlySet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"numberMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callerPublicSaleKey\",\"type\":\"uint256\"}],\"name\":\"publicSaleMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"auctionSaleStartTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"publicSaleStartTime\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"mintlistPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicPrice\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"publicSaleKey\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"numSlots\",\"type\":\"uint256[]\"}],\"name\":\"seedAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"name\":\"setAuctionSaleStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"setOwnersExplicit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"key\",\"type\":\"uint32\"}],\"name\":\"setPublicSaleKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTMetaData.ABI instead.
var NFTABI = NFTMetaData.ABI

// NFT is an auto generated Go binding around an Ethereum contract.
type NFT struct {
	NFTCaller     // Read-only binding to the contract
	NFTTransactor // Write-only binding to the contract
	NFTFilterer   // Log filterer for contract events
}

// NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTSession struct {
	Contract     *NFT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCallerSession struct {
	Contract *NFTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTTransactorSession struct {
	Contract     *NFTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTRaw struct {
	Contract *NFT // Generic contract binding to access the raw methods on
}

// NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCallerRaw struct {
	Contract *NFTCaller // Generic read-only contract binding to access the raw methods on
}

// NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTTransactorRaw struct {
	Contract *NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFT creates a new instance of NFT, bound to a specific deployed contract.
func NewNFT(address common.Address, backend bind.ContractBackend) (*NFT, error) {
	contract, err := bindNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFT{NFTCaller: NFTCaller{contract: contract}, NFTTransactor: NFTTransactor{contract: contract}, NFTFilterer: NFTFilterer{contract: contract}}, nil
}

// NewNFTCaller creates a new read-only instance of NFT, bound to a specific deployed contract.
func NewNFTCaller(address common.Address, caller bind.ContractCaller) (*NFTCaller, error) {
	contract, err := bindNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCaller{contract: contract}, nil
}

// NewNFTTransactor creates a new write-only instance of NFT, bound to a specific deployed contract.
func NewNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTTransactor, error) {
	contract, err := bindNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTTransactor{contract: contract}, nil
}

// NewNFTFilterer creates a new log filterer instance of NFT, bound to a specific deployed contract.
func NewNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTFilterer, error) {
	contract, err := bindNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTFilterer{contract: contract}, nil
}

// bindNFT binds a generic wrapper to an already deployed contract.
func bindNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transact(opts, method, params...)
}

// AUCTIONDROPINTERVAL is a free data retrieval call binding the contract method 0x5cae01d3.
//
// Solidity: function AUCTION_DROP_INTERVAL() view returns(uint256)
func (_NFT *NFTCaller) AUCTIONDROPINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "AUCTION_DROP_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AUCTIONDROPINTERVAL is a free data retrieval call binding the contract method 0x5cae01d3.
//
// Solidity: function AUCTION_DROP_INTERVAL() view returns(uint256)
func (_NFT *NFTSession) AUCTIONDROPINTERVAL() (*big.Int, error) {
	return _NFT.Contract.AUCTIONDROPINTERVAL(&_NFT.CallOpts)
}

// AUCTIONDROPINTERVAL is a free data retrieval call binding the contract method 0x5cae01d3.
//
// Solidity: function AUCTION_DROP_INTERVAL() view returns(uint256)
func (_NFT *NFTCallerSession) AUCTIONDROPINTERVAL() (*big.Int, error) {
	return _NFT.Contract.AUCTIONDROPINTERVAL(&_NFT.CallOpts)
}

// AUCTIONDROPPERSTEP is a free data retrieval call binding the contract method 0x59f369fe.
//
// Solidity: function AUCTION_DROP_PER_STEP() view returns(uint256)
func (_NFT *NFTCaller) AUCTIONDROPPERSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "AUCTION_DROP_PER_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AUCTIONDROPPERSTEP is a free data retrieval call binding the contract method 0x59f369fe.
//
// Solidity: function AUCTION_DROP_PER_STEP() view returns(uint256)
func (_NFT *NFTSession) AUCTIONDROPPERSTEP() (*big.Int, error) {
	return _NFT.Contract.AUCTIONDROPPERSTEP(&_NFT.CallOpts)
}

// AUCTIONDROPPERSTEP is a free data retrieval call binding the contract method 0x59f369fe.
//
// Solidity: function AUCTION_DROP_PER_STEP() view returns(uint256)
func (_NFT *NFTCallerSession) AUCTIONDROPPERSTEP() (*big.Int, error) {
	return _NFT.Contract.AUCTIONDROPPERSTEP(&_NFT.CallOpts)
}

// AUCTIONENDPRICE is a free data retrieval call binding the contract method 0xcaf8a6d1.
//
// Solidity: function AUCTION_END_PRICE() view returns(uint256)
func (_NFT *NFTCaller) AUCTIONENDPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "AUCTION_END_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AUCTIONENDPRICE is a free data retrieval call binding the contract method 0xcaf8a6d1.
//
// Solidity: function AUCTION_END_PRICE() view returns(uint256)
func (_NFT *NFTSession) AUCTIONENDPRICE() (*big.Int, error) {
	return _NFT.Contract.AUCTIONENDPRICE(&_NFT.CallOpts)
}

// AUCTIONENDPRICE is a free data retrieval call binding the contract method 0xcaf8a6d1.
//
// Solidity: function AUCTION_END_PRICE() view returns(uint256)
func (_NFT *NFTCallerSession) AUCTIONENDPRICE() (*big.Int, error) {
	return _NFT.Contract.AUCTIONENDPRICE(&_NFT.CallOpts)
}

// AUCTIONPRICECURVELENGTH is a free data retrieval call binding the contract method 0xf8a987d8.
//
// Solidity: function AUCTION_PRICE_CURVE_LENGTH() view returns(uint256)
func (_NFT *NFTCaller) AUCTIONPRICECURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "AUCTION_PRICE_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AUCTIONPRICECURVELENGTH is a free data retrieval call binding the contract method 0xf8a987d8.
//
// Solidity: function AUCTION_PRICE_CURVE_LENGTH() view returns(uint256)
func (_NFT *NFTSession) AUCTIONPRICECURVELENGTH() (*big.Int, error) {
	return _NFT.Contract.AUCTIONPRICECURVELENGTH(&_NFT.CallOpts)
}

// AUCTIONPRICECURVELENGTH is a free data retrieval call binding the contract method 0xf8a987d8.
//
// Solidity: function AUCTION_PRICE_CURVE_LENGTH() view returns(uint256)
func (_NFT *NFTCallerSession) AUCTIONPRICECURVELENGTH() (*big.Int, error) {
	return _NFT.Contract.AUCTIONPRICECURVELENGTH(&_NFT.CallOpts)
}

// AUCTIONSTARTPRICE is a free data retrieval call binding the contract method 0x7a1c4a56.
//
// Solidity: function AUCTION_START_PRICE() view returns(uint256)
func (_NFT *NFTCaller) AUCTIONSTARTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "AUCTION_START_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AUCTIONSTARTPRICE is a free data retrieval call binding the contract method 0x7a1c4a56.
//
// Solidity: function AUCTION_START_PRICE() view returns(uint256)
func (_NFT *NFTSession) AUCTIONSTARTPRICE() (*big.Int, error) {
	return _NFT.Contract.AUCTIONSTARTPRICE(&_NFT.CallOpts)
}

// AUCTIONSTARTPRICE is a free data retrieval call binding the contract method 0x7a1c4a56.
//
// Solidity: function AUCTION_START_PRICE() view returns(uint256)
func (_NFT *NFTCallerSession) AUCTIONSTARTPRICE() (*big.Int, error) {
	return _NFT.Contract.AUCTIONSTARTPRICE(&_NFT.CallOpts)
}

// Allowlist is a free data retrieval call binding the contract method 0xa7cd52cb.
//
// Solidity: function allowlist(address ) view returns(uint256)
func (_NFT *NFTCaller) Allowlist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "allowlist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowlist is a free data retrieval call binding the contract method 0xa7cd52cb.
//
// Solidity: function allowlist(address ) view returns(uint256)
func (_NFT *NFTSession) Allowlist(arg0 common.Address) (*big.Int, error) {
	return _NFT.Contract.Allowlist(&_NFT.CallOpts, arg0)
}

// Allowlist is a free data retrieval call binding the contract method 0xa7cd52cb.
//
// Solidity: function allowlist(address ) view returns(uint256)
func (_NFT *NFTCallerSession) Allowlist(arg0 common.Address) (*big.Int, error) {
	return _NFT.Contract.Allowlist(&_NFT.CallOpts, arg0)
}

// AmountForAuctionAndDev is a free data retrieval call binding the contract method 0x5666c880.
//
// Solidity: function amountForAuctionAndDev() view returns(uint256)
func (_NFT *NFTCaller) AmountForAuctionAndDev(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "amountForAuctionAndDev")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountForAuctionAndDev is a free data retrieval call binding the contract method 0x5666c880.
//
// Solidity: function amountForAuctionAndDev() view returns(uint256)
func (_NFT *NFTSession) AmountForAuctionAndDev() (*big.Int, error) {
	return _NFT.Contract.AmountForAuctionAndDev(&_NFT.CallOpts)
}

// AmountForAuctionAndDev is a free data retrieval call binding the contract method 0x5666c880.
//
// Solidity: function amountForAuctionAndDev() view returns(uint256)
func (_NFT *NFTCallerSession) AmountForAuctionAndDev() (*big.Int, error) {
	return _NFT.Contract.AmountForAuctionAndDev(&_NFT.CallOpts)
}

// AmountForDevs is a free data retrieval call binding the contract method 0xfbe1aa51.
//
// Solidity: function amountForDevs() view returns(uint256)
func (_NFT *NFTCaller) AmountForDevs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "amountForDevs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountForDevs is a free data retrieval call binding the contract method 0xfbe1aa51.
//
// Solidity: function amountForDevs() view returns(uint256)
func (_NFT *NFTSession) AmountForDevs() (*big.Int, error) {
	return _NFT.Contract.AmountForDevs(&_NFT.CallOpts)
}

// AmountForDevs is a free data retrieval call binding the contract method 0xfbe1aa51.
//
// Solidity: function amountForDevs() view returns(uint256)
func (_NFT *NFTCallerSession) AmountForDevs() (*big.Int, error) {
	return _NFT.Contract.AmountForDevs(&_NFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFT *NFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFT *NFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFT.Contract.BalanceOf(&_NFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFT *NFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFT.Contract.BalanceOf(&_NFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFT *NFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFT *NFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.GetApproved(&_NFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFT *NFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.GetApproved(&_NFT.CallOpts, tokenId)
}

// GetAuctionPrice is a free data retrieval call binding the contract method 0x917d009e.
//
// Solidity: function getAuctionPrice(uint256 _saleStartTime) view returns(uint256)
func (_NFT *NFTCaller) GetAuctionPrice(opts *bind.CallOpts, _saleStartTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getAuctionPrice", _saleStartTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAuctionPrice is a free data retrieval call binding the contract method 0x917d009e.
//
// Solidity: function getAuctionPrice(uint256 _saleStartTime) view returns(uint256)
func (_NFT *NFTSession) GetAuctionPrice(_saleStartTime *big.Int) (*big.Int, error) {
	return _NFT.Contract.GetAuctionPrice(&_NFT.CallOpts, _saleStartTime)
}

// GetAuctionPrice is a free data retrieval call binding the contract method 0x917d009e.
//
// Solidity: function getAuctionPrice(uint256 _saleStartTime) view returns(uint256)
func (_NFT *NFTCallerSession) GetAuctionPrice(_saleStartTime *big.Int) (*big.Int, error) {
	return _NFT.Contract.GetAuctionPrice(&_NFT.CallOpts, _saleStartTime)
}

// GetOwnershipData is a free data retrieval call binding the contract method 0x9231ab2a.
//
// Solidity: function getOwnershipData(uint256 tokenId) view returns((address,uint64))
func (_NFT *NFTCaller) GetOwnershipData(opts *bind.CallOpts, tokenId *big.Int) (ERC721ATokenOwnership, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getOwnershipData", tokenId)

	if err != nil {
		return *new(ERC721ATokenOwnership), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC721ATokenOwnership)).(*ERC721ATokenOwnership)

	return out0, err

}

// GetOwnershipData is a free data retrieval call binding the contract method 0x9231ab2a.
//
// Solidity: function getOwnershipData(uint256 tokenId) view returns((address,uint64))
func (_NFT *NFTSession) GetOwnershipData(tokenId *big.Int) (ERC721ATokenOwnership, error) {
	return _NFT.Contract.GetOwnershipData(&_NFT.CallOpts, tokenId)
}

// GetOwnershipData is a free data retrieval call binding the contract method 0x9231ab2a.
//
// Solidity: function getOwnershipData(uint256 tokenId) view returns((address,uint64))
func (_NFT *NFTCallerSession) GetOwnershipData(tokenId *big.Int) (ERC721ATokenOwnership, error) {
	return _NFT.Contract.GetOwnershipData(&_NFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFT *NFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFT *NFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFT.Contract.IsApprovedForAll(&_NFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFT *NFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFT.Contract.IsApprovedForAll(&_NFT.CallOpts, owner, operator)
}

// IsPublicSaleOn is a free data retrieval call binding the contract method 0x422030ba.
//
// Solidity: function isPublicSaleOn(uint256 publicPriceWei, uint256 publicSaleKey, uint256 publicSaleStartTime) view returns(bool)
func (_NFT *NFTCaller) IsPublicSaleOn(opts *bind.CallOpts, publicPriceWei *big.Int, publicSaleKey *big.Int, publicSaleStartTime *big.Int) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "isPublicSaleOn", publicPriceWei, publicSaleKey, publicSaleStartTime)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSaleOn is a free data retrieval call binding the contract method 0x422030ba.
//
// Solidity: function isPublicSaleOn(uint256 publicPriceWei, uint256 publicSaleKey, uint256 publicSaleStartTime) view returns(bool)
func (_NFT *NFTSession) IsPublicSaleOn(publicPriceWei *big.Int, publicSaleKey *big.Int, publicSaleStartTime *big.Int) (bool, error) {
	return _NFT.Contract.IsPublicSaleOn(&_NFT.CallOpts, publicPriceWei, publicSaleKey, publicSaleStartTime)
}

// IsPublicSaleOn is a free data retrieval call binding the contract method 0x422030ba.
//
// Solidity: function isPublicSaleOn(uint256 publicPriceWei, uint256 publicSaleKey, uint256 publicSaleStartTime) view returns(bool)
func (_NFT *NFTCallerSession) IsPublicSaleOn(publicPriceWei *big.Int, publicSaleKey *big.Int, publicSaleStartTime *big.Int) (bool, error) {
	return _NFT.Contract.IsPublicSaleOn(&_NFT.CallOpts, publicPriceWei, publicSaleKey, publicSaleStartTime)
}

// MaxPerAddressDuringMint is a free data retrieval call binding the contract method 0x8bc35c2f.
//
// Solidity: function maxPerAddressDuringMint() view returns(uint256)
func (_NFT *NFTCaller) MaxPerAddressDuringMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "maxPerAddressDuringMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerAddressDuringMint is a free data retrieval call binding the contract method 0x8bc35c2f.
//
// Solidity: function maxPerAddressDuringMint() view returns(uint256)
func (_NFT *NFTSession) MaxPerAddressDuringMint() (*big.Int, error) {
	return _NFT.Contract.MaxPerAddressDuringMint(&_NFT.CallOpts)
}

// MaxPerAddressDuringMint is a free data retrieval call binding the contract method 0x8bc35c2f.
//
// Solidity: function maxPerAddressDuringMint() view returns(uint256)
func (_NFT *NFTCallerSession) MaxPerAddressDuringMint() (*big.Int, error) {
	return _NFT.Contract.MaxPerAddressDuringMint(&_NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFT *NFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFT *NFTSession) Name() (string, error) {
	return _NFT.Contract.Name(&_NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFT *NFTCallerSession) Name() (string, error) {
	return _NFT.Contract.Name(&_NFT.CallOpts)
}

// NextOwnerToExplicitlySet is a free data retrieval call binding the contract method 0xd7224ba0.
//
// Solidity: function nextOwnerToExplicitlySet() view returns(uint256)
func (_NFT *NFTCaller) NextOwnerToExplicitlySet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "nextOwnerToExplicitlySet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOwnerToExplicitlySet is a free data retrieval call binding the contract method 0xd7224ba0.
//
// Solidity: function nextOwnerToExplicitlySet() view returns(uint256)
func (_NFT *NFTSession) NextOwnerToExplicitlySet() (*big.Int, error) {
	return _NFT.Contract.NextOwnerToExplicitlySet(&_NFT.CallOpts)
}

// NextOwnerToExplicitlySet is a free data retrieval call binding the contract method 0xd7224ba0.
//
// Solidity: function nextOwnerToExplicitlySet() view returns(uint256)
func (_NFT *NFTCallerSession) NextOwnerToExplicitlySet() (*big.Int, error) {
	return _NFT.Contract.NextOwnerToExplicitlySet(&_NFT.CallOpts)
}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address owner) view returns(uint256)
func (_NFT *NFTCaller) NumberMinted(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "numberMinted", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address owner) view returns(uint256)
func (_NFT *NFTSession) NumberMinted(owner common.Address) (*big.Int, error) {
	return _NFT.Contract.NumberMinted(&_NFT.CallOpts, owner)
}

// NumberMinted is a free data retrieval call binding the contract method 0xdc33e681.
//
// Solidity: function numberMinted(address owner) view returns(uint256)
func (_NFT *NFTCallerSession) NumberMinted(owner common.Address) (*big.Int, error) {
	return _NFT.Contract.NumberMinted(&_NFT.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFT *NFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFT *NFTSession) Owner() (common.Address, error) {
	return _NFT.Contract.Owner(&_NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFT *NFTCallerSession) Owner() (common.Address, error) {
	return _NFT.Contract.Owner(&_NFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFT *NFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFT *NFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.OwnerOf(&_NFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFT *NFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.OwnerOf(&_NFT.CallOpts, tokenId)
}

// SaleConfig is a free data retrieval call binding the contract method 0x90aa0b0f.
//
// Solidity: function saleConfig() view returns(uint32 auctionSaleStartTime, uint32 publicSaleStartTime, uint64 mintlistPrice, uint64 publicPrice, uint32 publicSaleKey)
func (_NFT *NFTCaller) SaleConfig(opts *bind.CallOpts) (struct {
	AuctionSaleStartTime uint32
	PublicSaleStartTime  uint32
	MintlistPrice        uint64
	PublicPrice          uint64
	PublicSaleKey        uint32
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "saleConfig")

	outstruct := new(struct {
		AuctionSaleStartTime uint32
		PublicSaleStartTime  uint32
		MintlistPrice        uint64
		PublicPrice          uint64
		PublicSaleKey        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AuctionSaleStartTime = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.PublicSaleStartTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.MintlistPrice = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PublicPrice = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.PublicSaleKey = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// SaleConfig is a free data retrieval call binding the contract method 0x90aa0b0f.
//
// Solidity: function saleConfig() view returns(uint32 auctionSaleStartTime, uint32 publicSaleStartTime, uint64 mintlistPrice, uint64 publicPrice, uint32 publicSaleKey)
func (_NFT *NFTSession) SaleConfig() (struct {
	AuctionSaleStartTime uint32
	PublicSaleStartTime  uint32
	MintlistPrice        uint64
	PublicPrice          uint64
	PublicSaleKey        uint32
}, error) {
	return _NFT.Contract.SaleConfig(&_NFT.CallOpts)
}

// SaleConfig is a free data retrieval call binding the contract method 0x90aa0b0f.
//
// Solidity: function saleConfig() view returns(uint32 auctionSaleStartTime, uint32 publicSaleStartTime, uint64 mintlistPrice, uint64 publicPrice, uint32 publicSaleKey)
func (_NFT *NFTCallerSession) SaleConfig() (struct {
	AuctionSaleStartTime uint32
	PublicSaleStartTime  uint32
	MintlistPrice        uint64
	PublicPrice          uint64
	PublicSaleKey        uint32
}, error) {
	return _NFT.Contract.SaleConfig(&_NFT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFT.Contract.SupportsInterface(&_NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFT.Contract.SupportsInterface(&_NFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFT *NFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFT *NFTSession) Symbol() (string, error) {
	return _NFT.Contract.Symbol(&_NFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFT *NFTCallerSession) Symbol() (string, error) {
	return _NFT.Contract.Symbol(&_NFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NFT *NFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NFT *NFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenByIndex(&_NFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NFT *NFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenByIndex(&_NFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NFT *NFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NFT *NFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenOfOwnerByIndex(&_NFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NFT *NFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenOfOwnerByIndex(&_NFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFT *NFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFT *NFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NFT.Contract.TokenURI(&_NFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NFT *NFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NFT.Contract.TokenURI(&_NFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFT *NFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFT *NFTSession) TotalSupply() (*big.Int, error) {
	return _NFT.Contract.TotalSupply(&_NFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFT *NFTCallerSession) TotalSupply() (*big.Int, error) {
	return _NFT.Contract.TotalSupply(&_NFT.CallOpts)
}

// AllowlistMint is a paid mutator transaction binding the contract method 0x41fbddbd.
//
// Solidity: function allowlistMint() payable returns()
func (_NFT *NFTTransactor) AllowlistMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "allowlistMint")
}

// AllowlistMint is a paid mutator transaction binding the contract method 0x41fbddbd.
//
// Solidity: function allowlistMint() payable returns()
func (_NFT *NFTSession) AllowlistMint() (*types.Transaction, error) {
	return _NFT.Contract.AllowlistMint(&_NFT.TransactOpts)
}

// AllowlistMint is a paid mutator transaction binding the contract method 0x41fbddbd.
//
// Solidity: function allowlistMint() payable returns()
func (_NFT *NFTTransactorSession) AllowlistMint() (*types.Transaction, error) {
	return _NFT.Contract.AllowlistMint(&_NFT.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFT *NFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFT *NFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Approve(&_NFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Approve(&_NFT.TransactOpts, to, tokenId)
}

// AuctionMint is a paid mutator transaction binding the contract method 0x4d3554c3.
//
// Solidity: function auctionMint(uint256 quantity) payable returns()
func (_NFT *NFTTransactor) AuctionMint(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "auctionMint", quantity)
}

// AuctionMint is a paid mutator transaction binding the contract method 0x4d3554c3.
//
// Solidity: function auctionMint(uint256 quantity) payable returns()
func (_NFT *NFTSession) AuctionMint(quantity *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.AuctionMint(&_NFT.TransactOpts, quantity)
}

// AuctionMint is a paid mutator transaction binding the contract method 0x4d3554c3.
//
// Solidity: function auctionMint(uint256 quantity) payable returns()
func (_NFT *NFTTransactorSession) AuctionMint(quantity *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.AuctionMint(&_NFT.TransactOpts, quantity)
}

// DevMint is a paid mutator transaction binding the contract method 0x375a069a.
//
// Solidity: function devMint(uint256 quantity) returns()
func (_NFT *NFTTransactor) DevMint(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "devMint", quantity)
}

// DevMint is a paid mutator transaction binding the contract method 0x375a069a.
//
// Solidity: function devMint(uint256 quantity) returns()
func (_NFT *NFTSession) DevMint(quantity *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.DevMint(&_NFT.TransactOpts, quantity)
}

// DevMint is a paid mutator transaction binding the contract method 0x375a069a.
//
// Solidity: function devMint(uint256 quantity) returns()
func (_NFT *NFTTransactorSession) DevMint(quantity *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.DevMint(&_NFT.TransactOpts, quantity)
}

// EndAuctionAndSetupNonAuctionSaleInfo is a paid mutator transaction binding the contract method 0x16e6e15a.
//
// Solidity: function endAuctionAndSetupNonAuctionSaleInfo(uint64 mintlistPriceWei, uint64 publicPriceWei, uint32 publicSaleStartTime) returns()
func (_NFT *NFTTransactor) EndAuctionAndSetupNonAuctionSaleInfo(opts *bind.TransactOpts, mintlistPriceWei uint64, publicPriceWei uint64, publicSaleStartTime uint32) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "endAuctionAndSetupNonAuctionSaleInfo", mintlistPriceWei, publicPriceWei, publicSaleStartTime)
}

// EndAuctionAndSetupNonAuctionSaleInfo is a paid mutator transaction binding the contract method 0x16e6e15a.
//
// Solidity: function endAuctionAndSetupNonAuctionSaleInfo(uint64 mintlistPriceWei, uint64 publicPriceWei, uint32 publicSaleStartTime) returns()
func (_NFT *NFTSession) EndAuctionAndSetupNonAuctionSaleInfo(mintlistPriceWei uint64, publicPriceWei uint64, publicSaleStartTime uint32) (*types.Transaction, error) {
	return _NFT.Contract.EndAuctionAndSetupNonAuctionSaleInfo(&_NFT.TransactOpts, mintlistPriceWei, publicPriceWei, publicSaleStartTime)
}

// EndAuctionAndSetupNonAuctionSaleInfo is a paid mutator transaction binding the contract method 0x16e6e15a.
//
// Solidity: function endAuctionAndSetupNonAuctionSaleInfo(uint64 mintlistPriceWei, uint64 publicPriceWei, uint32 publicSaleStartTime) returns()
func (_NFT *NFTTransactorSession) EndAuctionAndSetupNonAuctionSaleInfo(mintlistPriceWei uint64, publicPriceWei uint64, publicSaleStartTime uint32) (*types.Transaction, error) {
	return _NFT.Contract.EndAuctionAndSetupNonAuctionSaleInfo(&_NFT.TransactOpts, mintlistPriceWei, publicPriceWei, publicSaleStartTime)
}

// PublicSaleMint is a paid mutator transaction binding the contract method 0xcb91d8b3.
//
// Solidity: function publicSaleMint(uint256 quantity, uint256 callerPublicSaleKey) payable returns()
func (_NFT *NFTTransactor) PublicSaleMint(opts *bind.TransactOpts, quantity *big.Int, callerPublicSaleKey *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "publicSaleMint", quantity, callerPublicSaleKey)
}

// PublicSaleMint is a paid mutator transaction binding the contract method 0xcb91d8b3.
//
// Solidity: function publicSaleMint(uint256 quantity, uint256 callerPublicSaleKey) payable returns()
func (_NFT *NFTSession) PublicSaleMint(quantity *big.Int, callerPublicSaleKey *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.PublicSaleMint(&_NFT.TransactOpts, quantity, callerPublicSaleKey)
}

// PublicSaleMint is a paid mutator transaction binding the contract method 0xcb91d8b3.
//
// Solidity: function publicSaleMint(uint256 quantity, uint256 callerPublicSaleKey) payable returns()
func (_NFT *NFTTransactorSession) PublicSaleMint(quantity *big.Int, callerPublicSaleKey *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.PublicSaleMint(&_NFT.TransactOpts, quantity, callerPublicSaleKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFT *NFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFT *NFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFT.Contract.RenounceOwnership(&_NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFT *NFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFT.Contract.RenounceOwnership(&_NFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFT *NFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFT *NFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom0(&_NFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFT *NFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom0(&_NFT.TransactOpts, from, to, tokenId, _data)
}

// SeedAllowlist is a paid mutator transaction binding the contract method 0xb05863d5.
//
// Solidity: function seedAllowlist(address[] addresses, uint256[] numSlots) returns()
func (_NFT *NFTTransactor) SeedAllowlist(opts *bind.TransactOpts, addresses []common.Address, numSlots []*big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "seedAllowlist", addresses, numSlots)
}

// SeedAllowlist is a paid mutator transaction binding the contract method 0xb05863d5.
//
// Solidity: function seedAllowlist(address[] addresses, uint256[] numSlots) returns()
func (_NFT *NFTSession) SeedAllowlist(addresses []common.Address, numSlots []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SeedAllowlist(&_NFT.TransactOpts, addresses, numSlots)
}

// SeedAllowlist is a paid mutator transaction binding the contract method 0xb05863d5.
//
// Solidity: function seedAllowlist(address[] addresses, uint256[] numSlots) returns()
func (_NFT *NFTTransactorSession) SeedAllowlist(addresses []common.Address, numSlots []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SeedAllowlist(&_NFT.TransactOpts, addresses, numSlots)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFT *NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFT *NFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.Contract.SetApprovalForAll(&_NFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFT *NFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.Contract.SetApprovalForAll(&_NFT.TransactOpts, operator, approved)
}

// SetAuctionSaleStartTime is a paid mutator transaction binding the contract method 0x6ebc5601.
//
// Solidity: function setAuctionSaleStartTime(uint32 timestamp) returns()
func (_NFT *NFTTransactor) SetAuctionSaleStartTime(opts *bind.TransactOpts, timestamp uint32) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setAuctionSaleStartTime", timestamp)
}

// SetAuctionSaleStartTime is a paid mutator transaction binding the contract method 0x6ebc5601.
//
// Solidity: function setAuctionSaleStartTime(uint32 timestamp) returns()
func (_NFT *NFTSession) SetAuctionSaleStartTime(timestamp uint32) (*types.Transaction, error) {
	return _NFT.Contract.SetAuctionSaleStartTime(&_NFT.TransactOpts, timestamp)
}

// SetAuctionSaleStartTime is a paid mutator transaction binding the contract method 0x6ebc5601.
//
// Solidity: function setAuctionSaleStartTime(uint32 timestamp) returns()
func (_NFT *NFTTransactorSession) SetAuctionSaleStartTime(timestamp uint32) (*types.Transaction, error) {
	return _NFT.Contract.SetAuctionSaleStartTime(&_NFT.TransactOpts, timestamp)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_NFT *NFTTransactor) SetBaseURI(opts *bind.TransactOpts, baseURI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setBaseURI", baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_NFT *NFTSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _NFT.Contract.SetBaseURI(&_NFT.TransactOpts, baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string baseURI) returns()
func (_NFT *NFTTransactorSession) SetBaseURI(baseURI string) (*types.Transaction, error) {
	return _NFT.Contract.SetBaseURI(&_NFT.TransactOpts, baseURI)
}

// SetOwnersExplicit is a paid mutator transaction binding the contract method 0x2d20fb60.
//
// Solidity: function setOwnersExplicit(uint256 quantity) returns()
func (_NFT *NFTTransactor) SetOwnersExplicit(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setOwnersExplicit", quantity)
}

// SetOwnersExplicit is a paid mutator transaction binding the contract method 0x2d20fb60.
//
// Solidity: function setOwnersExplicit(uint256 quantity) returns()
func (_NFT *NFTSession) SetOwnersExplicit(quantity *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SetOwnersExplicit(&_NFT.TransactOpts, quantity)
}

// SetOwnersExplicit is a paid mutator transaction binding the contract method 0x2d20fb60.
//
// Solidity: function setOwnersExplicit(uint256 quantity) returns()
func (_NFT *NFTTransactorSession) SetOwnersExplicit(quantity *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SetOwnersExplicit(&_NFT.TransactOpts, quantity)
}

// SetPublicSaleKey is a paid mutator transaction binding the contract method 0x90028083.
//
// Solidity: function setPublicSaleKey(uint32 key) returns()
func (_NFT *NFTTransactor) SetPublicSaleKey(opts *bind.TransactOpts, key uint32) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setPublicSaleKey", key)
}

// SetPublicSaleKey is a paid mutator transaction binding the contract method 0x90028083.
//
// Solidity: function setPublicSaleKey(uint32 key) returns()
func (_NFT *NFTSession) SetPublicSaleKey(key uint32) (*types.Transaction, error) {
	return _NFT.Contract.SetPublicSaleKey(&_NFT.TransactOpts, key)
}

// SetPublicSaleKey is a paid mutator transaction binding the contract method 0x90028083.
//
// Solidity: function setPublicSaleKey(uint32 key) returns()
func (_NFT *NFTTransactorSession) SetPublicSaleKey(key uint32) (*types.Transaction, error) {
	return _NFT.Contract.SetPublicSaleKey(&_NFT.TransactOpts, key)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.TransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.TransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFT *NFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFT *NFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFT.Contract.TransferOwnership(&_NFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFT *NFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFT.Contract.TransferOwnership(&_NFT.TransactOpts, newOwner)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_NFT *NFTTransactor) WithdrawMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "withdrawMoney")
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_NFT *NFTSession) WithdrawMoney() (*types.Transaction, error) {
	return _NFT.Contract.WithdrawMoney(&_NFT.TransactOpts)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_NFT *NFTTransactorSession) WithdrawMoney() (*types.Transaction, error) {
	return _NFT.Contract.WithdrawMoney(&_NFT.TransactOpts)
}

// NFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NFT contract.
type NFTApprovalIterator struct {
	Event *NFTApproval // Event containing the contract specifics and raw log

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
func (it *NFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTApproval)
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
		it.Event = new(NFTApproval)
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
func (it *NFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTApproval represents a Approval event raised by the NFT contract.
type NFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFT *NFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NFTApprovalIterator, error) {

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

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTApprovalIterator{contract: _NFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFT *NFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTApproval)
				if err := _NFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_NFT *NFTFilterer) ParseApproval(log types.Log) (*NFTApproval, error) {
	event := new(NFTApproval)
	if err := _NFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFT contract.
type NFTApprovalForAllIterator struct {
	Event *NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTApprovalForAll)
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
		it.Event = new(NFTApprovalForAll)
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
func (it *NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTApprovalForAll represents a ApprovalForAll event raised by the NFT contract.
type NFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFT *NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTApprovalForAllIterator{contract: _NFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFT *NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTApprovalForAll)
				if err := _NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NFT *NFTFilterer) ParseApprovalForAll(log types.Log) (*NFTApprovalForAll, error) {
	event := new(NFTApprovalForAll)
	if err := _NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFT contract.
type NFTOwnershipTransferredIterator struct {
	Event *NFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTOwnershipTransferred)
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
		it.Event = new(NFTOwnershipTransferred)
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
func (it *NFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTOwnershipTransferred represents a OwnershipTransferred event raised by the NFT contract.
type NFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFT *NFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTOwnershipTransferredIterator{contract: _NFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFT *NFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTOwnershipTransferred)
				if err := _NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFT *NFTFilterer) ParseOwnershipTransferred(log types.Log) (*NFTOwnershipTransferred, error) {
	event := new(NFTOwnershipTransferred)
	if err := _NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFT contract.
type NFTTransferIterator struct {
	Event *NFTTransfer // Event containing the contract specifics and raw log

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
func (it *NFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTTransfer)
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
		it.Event = new(NFTTransfer)
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
func (it *NFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTTransfer represents a Transfer event raised by the NFT contract.
type NFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFT *NFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NFTTransferIterator, error) {

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

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTTransferIterator{contract: _NFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFT *NFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTTransfer)
				if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_NFT *NFTFilterer) ParseTransfer(log types.Log) (*NFTTransfer, error) {
	event := new(NFTTransfer)
	if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
