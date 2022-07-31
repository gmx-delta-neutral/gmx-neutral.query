// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolcommitter

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
)

// IPoolCommitterBalance is an auto generated low-level Go binding around an user-defined struct.
type IPoolCommitterBalance struct {
	LongTokens       *big.Int
	ShortTokens      *big.Int
	SettlementTokens *big.Int
}

// PoolcommitterMetaData contains all meta data concerning the Poolcommitter contract.
var PoolcommitterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AggregateBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_burningFee\",\"type\":\"uint256\"}],\"name\":\"BurningFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_changeInterval\",\"type\":\"uint256\"}],\"name\":\"ChangeIntervalSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumIPoolCommitter.CommitType\",\"name\":\"commitType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"appropriateUpdateIntervalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fromAggregateBalance\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"payForClaim\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"mintingFee\",\"type\":\"bytes16\"}],\"name\":\"CreateCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"updateIntervalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"burningFee\",\"type\":\"bytes16\"}],\"name\":\"ExecutedCommitsForInterval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_feeController\",\"type\":\"address\"}],\"name\":\"FeeControllerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_mintingFee\",\"type\":\"uint256\"}],\"name\":\"MintingFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"PoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LONG_INDEX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BURNING_FEE\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CHANGE_INTERVAL\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ITERATIONS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_MINTING_FEE\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORT_INDEX\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoClaim\",\"outputs\":[{\"internalType\":\"contractIAutoClaim\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burnFeeHistory\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burningFee\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"changeInterval\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"args\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lastPriceTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBalance\",\"type\":\"uint256\"}],\"name\":\"executeCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAggregateBalance\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTokens\",\"type\":\"uint256\"}],\"internalType\":\"structIPoolCommitter.Balance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppropriateUpdateIntervalId\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBurningFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_autoClaim\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeController\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_invariantCheck\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_burningFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_changeInterval\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPoolCommitter.CommitType\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"isBurn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPoolCommitter.CommitType\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"isLong\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPoolCommitter.CommitType\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"isMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPoolCommitter.CommitType\",\"name\":\"t\",\"type\":\"uint8\"}],\"name\":\"isShort\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastUpdatedIntervalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leveragedPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintingFee\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingLongBurnPoolTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingMintSettlementAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingShortBurnPoolTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"priceHistory\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"longPrice\",\"type\":\"bytes16\"},{\"internalType\":\"bytes16\",\"name\":\"shortPrice\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burningFee\",\"type\":\"uint256\"}],\"name\":\"setBurningFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_changeInterval\",\"type\":\"uint256\"}],\"name\":\"setChangeInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeController\",\"type\":\"address\"}],\"name\":\"setFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintingFee\",\"type\":\"uint256\"}],\"name\":\"setMintingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leveragedPool\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalPoolCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"longMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBurnPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBurnPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBurnLongMintPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBurnShortMintPoolTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unAggregatedCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"updateAggregateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateIntervalId\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userAggregateBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"longTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"longMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBurnPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBurnPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBurnLongMintPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBurnShortMintPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updateIntervalId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolcommitterABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolcommitterMetaData.ABI instead.
var PoolcommitterABI = PoolcommitterMetaData.ABI

// Poolcommitter is an auto generated Go binding around an Ethereum contract.
type Poolcommitter struct {
	PoolcommitterCaller     // Read-only binding to the contract
	PoolcommitterTransactor // Write-only binding to the contract
	PoolcommitterFilterer   // Log filterer for contract events
}

// PoolcommitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolcommitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolcommitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolcommitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolcommitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolcommitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolcommitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolcommitterSession struct {
	Contract     *Poolcommitter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolcommitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolcommitterCallerSession struct {
	Contract *PoolcommitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PoolcommitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolcommitterTransactorSession struct {
	Contract     *PoolcommitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PoolcommitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolcommitterRaw struct {
	Contract *Poolcommitter // Generic contract binding to access the raw methods on
}

// PoolcommitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolcommitterCallerRaw struct {
	Contract *PoolcommitterCaller // Generic read-only contract binding to access the raw methods on
}

// PoolcommitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolcommitterTransactorRaw struct {
	Contract *PoolcommitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolcommitter creates a new instance of Poolcommitter, bound to a specific deployed contract.
func NewPoolcommitter(address common.Address, backend bind.ContractBackend) (*Poolcommitter, error) {
	contract, err := bindPoolcommitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolcommitter{PoolcommitterCaller: PoolcommitterCaller{contract: contract}, PoolcommitterTransactor: PoolcommitterTransactor{contract: contract}, PoolcommitterFilterer: PoolcommitterFilterer{contract: contract}}, nil
}

// NewPoolcommitterCaller creates a new read-only instance of Poolcommitter, bound to a specific deployed contract.
func NewPoolcommitterCaller(address common.Address, caller bind.ContractCaller) (*PoolcommitterCaller, error) {
	contract, err := bindPoolcommitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterCaller{contract: contract}, nil
}

// NewPoolcommitterTransactor creates a new write-only instance of Poolcommitter, bound to a specific deployed contract.
func NewPoolcommitterTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolcommitterTransactor, error) {
	contract, err := bindPoolcommitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterTransactor{contract: contract}, nil
}

// NewPoolcommitterFilterer creates a new log filterer instance of Poolcommitter, bound to a specific deployed contract.
func NewPoolcommitterFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolcommitterFilterer, error) {
	contract, err := bindPoolcommitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterFilterer{contract: contract}, nil
}

// bindPoolcommitter binds a generic wrapper to an already deployed contract.
func bindPoolcommitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolcommitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolcommitter *PoolcommitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolcommitter.Contract.PoolcommitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolcommitter *PoolcommitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolcommitter.Contract.PoolcommitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolcommitter *PoolcommitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolcommitter.Contract.PoolcommitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolcommitter *PoolcommitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolcommitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolcommitter *PoolcommitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolcommitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolcommitter *PoolcommitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolcommitter.Contract.contract.Transact(opts, method, params...)
}

// LONGINDEX is a free data retrieval call binding the contract method 0x062e0d69.
//
// Solidity: function LONG_INDEX() view returns(uint128)
func (_Poolcommitter *PoolcommitterCaller) LONGINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "LONG_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LONGINDEX is a free data retrieval call binding the contract method 0x062e0d69.
//
// Solidity: function LONG_INDEX() view returns(uint128)
func (_Poolcommitter *PoolcommitterSession) LONGINDEX() (*big.Int, error) {
	return _Poolcommitter.Contract.LONGINDEX(&_Poolcommitter.CallOpts)
}

// LONGINDEX is a free data retrieval call binding the contract method 0x062e0d69.
//
// Solidity: function LONG_INDEX() view returns(uint128)
func (_Poolcommitter *PoolcommitterCallerSession) LONGINDEX() (*big.Int, error) {
	return _Poolcommitter.Contract.LONGINDEX(&_Poolcommitter.CallOpts)
}

// MAXBURNINGFEE is a free data retrieval call binding the contract method 0x6fa1268e.
//
// Solidity: function MAX_BURNING_FEE() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) MAXBURNINGFEE(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "MAX_BURNING_FEE")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// MAXBURNINGFEE is a free data retrieval call binding the contract method 0x6fa1268e.
//
// Solidity: function MAX_BURNING_FEE() view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) MAXBURNINGFEE() ([16]byte, error) {
	return _Poolcommitter.Contract.MAXBURNINGFEE(&_Poolcommitter.CallOpts)
}

// MAXBURNINGFEE is a free data retrieval call binding the contract method 0x6fa1268e.
//
// Solidity: function MAX_BURNING_FEE() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) MAXBURNINGFEE() ([16]byte, error) {
	return _Poolcommitter.Contract.MAXBURNINGFEE(&_Poolcommitter.CallOpts)
}

// MAXCHANGEINTERVAL is a free data retrieval call binding the contract method 0xf3fa35a0.
//
// Solidity: function MAX_CHANGE_INTERVAL() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) MAXCHANGEINTERVAL(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "MAX_CHANGE_INTERVAL")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// MAXCHANGEINTERVAL is a free data retrieval call binding the contract method 0xf3fa35a0.
//
// Solidity: function MAX_CHANGE_INTERVAL() view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) MAXCHANGEINTERVAL() ([16]byte, error) {
	return _Poolcommitter.Contract.MAXCHANGEINTERVAL(&_Poolcommitter.CallOpts)
}

// MAXCHANGEINTERVAL is a free data retrieval call binding the contract method 0xf3fa35a0.
//
// Solidity: function MAX_CHANGE_INTERVAL() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) MAXCHANGEINTERVAL() ([16]byte, error) {
	return _Poolcommitter.Contract.MAXCHANGEINTERVAL(&_Poolcommitter.CallOpts)
}

// MAXITERATIONS is a free data retrieval call binding the contract method 0x1d92c26d.
//
// Solidity: function MAX_ITERATIONS() view returns(uint8)
func (_Poolcommitter *PoolcommitterCaller) MAXITERATIONS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "MAX_ITERATIONS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXITERATIONS is a free data retrieval call binding the contract method 0x1d92c26d.
//
// Solidity: function MAX_ITERATIONS() view returns(uint8)
func (_Poolcommitter *PoolcommitterSession) MAXITERATIONS() (uint8, error) {
	return _Poolcommitter.Contract.MAXITERATIONS(&_Poolcommitter.CallOpts)
}

// MAXITERATIONS is a free data retrieval call binding the contract method 0x1d92c26d.
//
// Solidity: function MAX_ITERATIONS() view returns(uint8)
func (_Poolcommitter *PoolcommitterCallerSession) MAXITERATIONS() (uint8, error) {
	return _Poolcommitter.Contract.MAXITERATIONS(&_Poolcommitter.CallOpts)
}

// MAXMINTINGFEE is a free data retrieval call binding the contract method 0x88b78aaa.
//
// Solidity: function MAX_MINTING_FEE() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) MAXMINTINGFEE(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "MAX_MINTING_FEE")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// MAXMINTINGFEE is a free data retrieval call binding the contract method 0x88b78aaa.
//
// Solidity: function MAX_MINTING_FEE() view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) MAXMINTINGFEE() ([16]byte, error) {
	return _Poolcommitter.Contract.MAXMINTINGFEE(&_Poolcommitter.CallOpts)
}

// MAXMINTINGFEE is a free data retrieval call binding the contract method 0x88b78aaa.
//
// Solidity: function MAX_MINTING_FEE() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) MAXMINTINGFEE() ([16]byte, error) {
	return _Poolcommitter.Contract.MAXMINTINGFEE(&_Poolcommitter.CallOpts)
}

// SHORTINDEX is a free data retrieval call binding the contract method 0xc40f6aa9.
//
// Solidity: function SHORT_INDEX() view returns(uint128)
func (_Poolcommitter *PoolcommitterCaller) SHORTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "SHORT_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTINDEX is a free data retrieval call binding the contract method 0xc40f6aa9.
//
// Solidity: function SHORT_INDEX() view returns(uint128)
func (_Poolcommitter *PoolcommitterSession) SHORTINDEX() (*big.Int, error) {
	return _Poolcommitter.Contract.SHORTINDEX(&_Poolcommitter.CallOpts)
}

// SHORTINDEX is a free data retrieval call binding the contract method 0xc40f6aa9.
//
// Solidity: function SHORT_INDEX() view returns(uint128)
func (_Poolcommitter *PoolcommitterCallerSession) SHORTINDEX() (*big.Int, error) {
	return _Poolcommitter.Contract.SHORTINDEX(&_Poolcommitter.CallOpts)
}

// AutoClaim is a free data retrieval call binding the contract method 0xc70db323.
//
// Solidity: function autoClaim() view returns(address)
func (_Poolcommitter *PoolcommitterCaller) AutoClaim(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "autoClaim")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AutoClaim is a free data retrieval call binding the contract method 0xc70db323.
//
// Solidity: function autoClaim() view returns(address)
func (_Poolcommitter *PoolcommitterSession) AutoClaim() (common.Address, error) {
	return _Poolcommitter.Contract.AutoClaim(&_Poolcommitter.CallOpts)
}

// AutoClaim is a free data retrieval call binding the contract method 0xc70db323.
//
// Solidity: function autoClaim() view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) AutoClaim() (common.Address, error) {
	return _Poolcommitter.Contract.AutoClaim(&_Poolcommitter.CallOpts)
}

// BurnFeeHistory is a free data retrieval call binding the contract method 0x1aabc7e7.
//
// Solidity: function burnFeeHistory(uint256 ) view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) BurnFeeHistory(opts *bind.CallOpts, arg0 *big.Int) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "burnFeeHistory", arg0)

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// BurnFeeHistory is a free data retrieval call binding the contract method 0x1aabc7e7.
//
// Solidity: function burnFeeHistory(uint256 ) view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) BurnFeeHistory(arg0 *big.Int) ([16]byte, error) {
	return _Poolcommitter.Contract.BurnFeeHistory(&_Poolcommitter.CallOpts, arg0)
}

// BurnFeeHistory is a free data retrieval call binding the contract method 0x1aabc7e7.
//
// Solidity: function burnFeeHistory(uint256 ) view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) BurnFeeHistory(arg0 *big.Int) ([16]byte, error) {
	return _Poolcommitter.Contract.BurnFeeHistory(&_Poolcommitter.CallOpts, arg0)
}

// BurningFee is a free data retrieval call binding the contract method 0x90f712cc.
//
// Solidity: function burningFee() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) BurningFee(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "burningFee")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// BurningFee is a free data retrieval call binding the contract method 0x90f712cc.
//
// Solidity: function burningFee() view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) BurningFee() ([16]byte, error) {
	return _Poolcommitter.Contract.BurningFee(&_Poolcommitter.CallOpts)
}

// BurningFee is a free data retrieval call binding the contract method 0x90f712cc.
//
// Solidity: function burningFee() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) BurningFee() ([16]byte, error) {
	return _Poolcommitter.Contract.BurningFee(&_Poolcommitter.CallOpts)
}

// ChangeInterval is a free data retrieval call binding the contract method 0xe20acc79.
//
// Solidity: function changeInterval() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) ChangeInterval(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "changeInterval")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// ChangeInterval is a free data retrieval call binding the contract method 0xe20acc79.
//
// Solidity: function changeInterval() view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) ChangeInterval() ([16]byte, error) {
	return _Poolcommitter.Contract.ChangeInterval(&_Poolcommitter.CallOpts)
}

// ChangeInterval is a free data retrieval call binding the contract method 0xe20acc79.
//
// Solidity: function changeInterval() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) ChangeInterval() ([16]byte, error) {
	return _Poolcommitter.Contract.ChangeInterval(&_Poolcommitter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Poolcommitter *PoolcommitterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Poolcommitter *PoolcommitterSession) Factory() (common.Address, error) {
	return _Poolcommitter.Contract.Factory(&_Poolcommitter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) Factory() (common.Address, error) {
	return _Poolcommitter.Contract.Factory(&_Poolcommitter.CallOpts)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Poolcommitter *PoolcommitterCaller) FeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "feeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Poolcommitter *PoolcommitterSession) FeeController() (common.Address, error) {
	return _Poolcommitter.Contract.FeeController(&_Poolcommitter.CallOpts)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) FeeController() (common.Address, error) {
	return _Poolcommitter.Contract.FeeController(&_Poolcommitter.CallOpts)
}

// GetAggregateBalance is a free data retrieval call binding the contract method 0x54da4992.
//
// Solidity: function getAggregateBalance(address user) view returns((uint256,uint256,uint256))
func (_Poolcommitter *PoolcommitterCaller) GetAggregateBalance(opts *bind.CallOpts, user common.Address) (IPoolCommitterBalance, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "getAggregateBalance", user)

	if err != nil {
		return *new(IPoolCommitterBalance), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolCommitterBalance)).(*IPoolCommitterBalance)

	return out0, err

}

// GetAggregateBalance is a free data retrieval call binding the contract method 0x54da4992.
//
// Solidity: function getAggregateBalance(address user) view returns((uint256,uint256,uint256))
func (_Poolcommitter *PoolcommitterSession) GetAggregateBalance(user common.Address) (IPoolCommitterBalance, error) {
	return _Poolcommitter.Contract.GetAggregateBalance(&_Poolcommitter.CallOpts, user)
}

// GetAggregateBalance is a free data retrieval call binding the contract method 0x54da4992.
//
// Solidity: function getAggregateBalance(address user) view returns((uint256,uint256,uint256))
func (_Poolcommitter *PoolcommitterCallerSession) GetAggregateBalance(user common.Address) (IPoolCommitterBalance, error) {
	return _Poolcommitter.Contract.GetAggregateBalance(&_Poolcommitter.CallOpts, user)
}

// GetAppropriateUpdateIntervalId is a free data retrieval call binding the contract method 0xd4c4e283.
//
// Solidity: function getAppropriateUpdateIntervalId() view returns(uint128)
func (_Poolcommitter *PoolcommitterCaller) GetAppropriateUpdateIntervalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "getAppropriateUpdateIntervalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAppropriateUpdateIntervalId is a free data retrieval call binding the contract method 0xd4c4e283.
//
// Solidity: function getAppropriateUpdateIntervalId() view returns(uint128)
func (_Poolcommitter *PoolcommitterSession) GetAppropriateUpdateIntervalId() (*big.Int, error) {
	return _Poolcommitter.Contract.GetAppropriateUpdateIntervalId(&_Poolcommitter.CallOpts)
}

// GetAppropriateUpdateIntervalId is a free data retrieval call binding the contract method 0xd4c4e283.
//
// Solidity: function getAppropriateUpdateIntervalId() view returns(uint128)
func (_Poolcommitter *PoolcommitterCallerSession) GetAppropriateUpdateIntervalId() (*big.Int, error) {
	return _Poolcommitter.Contract.GetAppropriateUpdateIntervalId(&_Poolcommitter.CallOpts)
}

// GetBurningFee is a free data retrieval call binding the contract method 0xfa871665.
//
// Solidity: function getBurningFee() view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) GetBurningFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "getBurningFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBurningFee is a free data retrieval call binding the contract method 0xfa871665.
//
// Solidity: function getBurningFee() view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) GetBurningFee() (*big.Int, error) {
	return _Poolcommitter.Contract.GetBurningFee(&_Poolcommitter.CallOpts)
}

// GetBurningFee is a free data retrieval call binding the contract method 0xfa871665.
//
// Solidity: function getBurningFee() view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) GetBurningFee() (*big.Int, error) {
	return _Poolcommitter.Contract.GetBurningFee(&_Poolcommitter.CallOpts)
}

// GetMintingFee is a free data retrieval call binding the contract method 0xb59d88a9.
//
// Solidity: function getMintingFee() view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) GetMintingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "getMintingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintingFee is a free data retrieval call binding the contract method 0xb59d88a9.
//
// Solidity: function getMintingFee() view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) GetMintingFee() (*big.Int, error) {
	return _Poolcommitter.Contract.GetMintingFee(&_Poolcommitter.CallOpts)
}

// GetMintingFee is a free data retrieval call binding the contract method 0xb59d88a9.
//
// Solidity: function getMintingFee() view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) GetMintingFee() (*big.Int, error) {
	return _Poolcommitter.Contract.GetMintingFee(&_Poolcommitter.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Poolcommitter *PoolcommitterCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Poolcommitter *PoolcommitterSession) Governance() (common.Address, error) {
	return _Poolcommitter.Contract.Governance(&_Poolcommitter.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) Governance() (common.Address, error) {
	return _Poolcommitter.Contract.Governance(&_Poolcommitter.CallOpts)
}

// InvariantCheck is a free data retrieval call binding the contract method 0x4509017e.
//
// Solidity: function invariantCheck() view returns(address)
func (_Poolcommitter *PoolcommitterCaller) InvariantCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "invariantCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InvariantCheck is a free data retrieval call binding the contract method 0x4509017e.
//
// Solidity: function invariantCheck() view returns(address)
func (_Poolcommitter *PoolcommitterSession) InvariantCheck() (common.Address, error) {
	return _Poolcommitter.Contract.InvariantCheck(&_Poolcommitter.CallOpts)
}

// InvariantCheck is a free data retrieval call binding the contract method 0x4509017e.
//
// Solidity: function invariantCheck() view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) InvariantCheck() (common.Address, error) {
	return _Poolcommitter.Contract.InvariantCheck(&_Poolcommitter.CallOpts)
}

// IsBurn is a free data retrieval call binding the contract method 0x8d7269fa.
//
// Solidity: function isBurn(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCaller) IsBurn(opts *bind.CallOpts, t uint8) (bool, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "isBurn", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBurn is a free data retrieval call binding the contract method 0x8d7269fa.
//
// Solidity: function isBurn(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterSession) IsBurn(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsBurn(&_Poolcommitter.CallOpts, t)
}

// IsBurn is a free data retrieval call binding the contract method 0x8d7269fa.
//
// Solidity: function isBurn(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCallerSession) IsBurn(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsBurn(&_Poolcommitter.CallOpts, t)
}

// IsLong is a free data retrieval call binding the contract method 0x882d882b.
//
// Solidity: function isLong(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCaller) IsLong(opts *bind.CallOpts, t uint8) (bool, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "isLong", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLong is a free data retrieval call binding the contract method 0x882d882b.
//
// Solidity: function isLong(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterSession) IsLong(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsLong(&_Poolcommitter.CallOpts, t)
}

// IsLong is a free data retrieval call binding the contract method 0x882d882b.
//
// Solidity: function isLong(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCallerSession) IsLong(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsLong(&_Poolcommitter.CallOpts, t)
}

// IsMint is a free data retrieval call binding the contract method 0x6333091a.
//
// Solidity: function isMint(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCaller) IsMint(opts *bind.CallOpts, t uint8) (bool, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "isMint", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMint is a free data retrieval call binding the contract method 0x6333091a.
//
// Solidity: function isMint(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterSession) IsMint(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsMint(&_Poolcommitter.CallOpts, t)
}

// IsMint is a free data retrieval call binding the contract method 0x6333091a.
//
// Solidity: function isMint(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCallerSession) IsMint(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsMint(&_Poolcommitter.CallOpts, t)
}

// IsShort is a free data retrieval call binding the contract method 0xc3bed53f.
//
// Solidity: function isShort(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCaller) IsShort(opts *bind.CallOpts, t uint8) (bool, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "isShort", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShort is a free data retrieval call binding the contract method 0xc3bed53f.
//
// Solidity: function isShort(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterSession) IsShort(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsShort(&_Poolcommitter.CallOpts, t)
}

// IsShort is a free data retrieval call binding the contract method 0xc3bed53f.
//
// Solidity: function isShort(uint8 t) pure returns(bool)
func (_Poolcommitter *PoolcommitterCallerSession) IsShort(t uint8) (bool, error) {
	return _Poolcommitter.Contract.IsShort(&_Poolcommitter.CallOpts, t)
}

// LastUpdatedIntervalId is a free data retrieval call binding the contract method 0x55c10624.
//
// Solidity: function lastUpdatedIntervalId(address ) view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) LastUpdatedIntervalId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "lastUpdatedIntervalId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedIntervalId is a free data retrieval call binding the contract method 0x55c10624.
//
// Solidity: function lastUpdatedIntervalId(address ) view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) LastUpdatedIntervalId(arg0 common.Address) (*big.Int, error) {
	return _Poolcommitter.Contract.LastUpdatedIntervalId(&_Poolcommitter.CallOpts, arg0)
}

// LastUpdatedIntervalId is a free data retrieval call binding the contract method 0x55c10624.
//
// Solidity: function lastUpdatedIntervalId(address ) view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) LastUpdatedIntervalId(arg0 common.Address) (*big.Int, error) {
	return _Poolcommitter.Contract.LastUpdatedIntervalId(&_Poolcommitter.CallOpts, arg0)
}

// LeveragedPool is a free data retrieval call binding the contract method 0xe323aac6.
//
// Solidity: function leveragedPool() view returns(address)
func (_Poolcommitter *PoolcommitterCaller) LeveragedPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "leveragedPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LeveragedPool is a free data retrieval call binding the contract method 0xe323aac6.
//
// Solidity: function leveragedPool() view returns(address)
func (_Poolcommitter *PoolcommitterSession) LeveragedPool() (common.Address, error) {
	return _Poolcommitter.Contract.LeveragedPool(&_Poolcommitter.CallOpts)
}

// LeveragedPool is a free data retrieval call binding the contract method 0xe323aac6.
//
// Solidity: function leveragedPool() view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) LeveragedPool() (common.Address, error) {
	return _Poolcommitter.Contract.LeveragedPool(&_Poolcommitter.CallOpts)
}

// MintingFee is a free data retrieval call binding the contract method 0x5a64ad95.
//
// Solidity: function mintingFee() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCaller) MintingFee(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "mintingFee")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// MintingFee is a free data retrieval call binding the contract method 0x5a64ad95.
//
// Solidity: function mintingFee() view returns(bytes16)
func (_Poolcommitter *PoolcommitterSession) MintingFee() ([16]byte, error) {
	return _Poolcommitter.Contract.MintingFee(&_Poolcommitter.CallOpts)
}

// MintingFee is a free data retrieval call binding the contract method 0x5a64ad95.
//
// Solidity: function mintingFee() view returns(bytes16)
func (_Poolcommitter *PoolcommitterCallerSession) MintingFee() ([16]byte, error) {
	return _Poolcommitter.Contract.MintingFee(&_Poolcommitter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Poolcommitter *PoolcommitterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Poolcommitter *PoolcommitterSession) Paused() (bool, error) {
	return _Poolcommitter.Contract.Paused(&_Poolcommitter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Poolcommitter *PoolcommitterCallerSession) Paused() (bool, error) {
	return _Poolcommitter.Contract.Paused(&_Poolcommitter.CallOpts)
}

// PendingLongBurnPoolTokens is a free data retrieval call binding the contract method 0xaf914976.
//
// Solidity: function pendingLongBurnPoolTokens() view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) PendingLongBurnPoolTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "pendingLongBurnPoolTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingLongBurnPoolTokens is a free data retrieval call binding the contract method 0xaf914976.
//
// Solidity: function pendingLongBurnPoolTokens() view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) PendingLongBurnPoolTokens() (*big.Int, error) {
	return _Poolcommitter.Contract.PendingLongBurnPoolTokens(&_Poolcommitter.CallOpts)
}

// PendingLongBurnPoolTokens is a free data retrieval call binding the contract method 0xaf914976.
//
// Solidity: function pendingLongBurnPoolTokens() view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) PendingLongBurnPoolTokens() (*big.Int, error) {
	return _Poolcommitter.Contract.PendingLongBurnPoolTokens(&_Poolcommitter.CallOpts)
}

// PendingMintSettlementAmount is a free data retrieval call binding the contract method 0xce660916.
//
// Solidity: function pendingMintSettlementAmount() view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) PendingMintSettlementAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "pendingMintSettlementAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingMintSettlementAmount is a free data retrieval call binding the contract method 0xce660916.
//
// Solidity: function pendingMintSettlementAmount() view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) PendingMintSettlementAmount() (*big.Int, error) {
	return _Poolcommitter.Contract.PendingMintSettlementAmount(&_Poolcommitter.CallOpts)
}

// PendingMintSettlementAmount is a free data retrieval call binding the contract method 0xce660916.
//
// Solidity: function pendingMintSettlementAmount() view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) PendingMintSettlementAmount() (*big.Int, error) {
	return _Poolcommitter.Contract.PendingMintSettlementAmount(&_Poolcommitter.CallOpts)
}

// PendingShortBurnPoolTokens is a free data retrieval call binding the contract method 0x97ac250f.
//
// Solidity: function pendingShortBurnPoolTokens() view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) PendingShortBurnPoolTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "pendingShortBurnPoolTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingShortBurnPoolTokens is a free data retrieval call binding the contract method 0x97ac250f.
//
// Solidity: function pendingShortBurnPoolTokens() view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) PendingShortBurnPoolTokens() (*big.Int, error) {
	return _Poolcommitter.Contract.PendingShortBurnPoolTokens(&_Poolcommitter.CallOpts)
}

// PendingShortBurnPoolTokens is a free data retrieval call binding the contract method 0x97ac250f.
//
// Solidity: function pendingShortBurnPoolTokens() view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) PendingShortBurnPoolTokens() (*big.Int, error) {
	return _Poolcommitter.Contract.PendingShortBurnPoolTokens(&_Poolcommitter.CallOpts)
}

// PriceHistory is a free data retrieval call binding the contract method 0xb27a0484.
//
// Solidity: function priceHistory(uint256 ) view returns(bytes16 longPrice, bytes16 shortPrice)
func (_Poolcommitter *PoolcommitterCaller) PriceHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LongPrice  [16]byte
	ShortPrice [16]byte
}, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "priceHistory", arg0)

	outstruct := new(struct {
		LongPrice  [16]byte
		ShortPrice [16]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LongPrice = *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)
	outstruct.ShortPrice = *abi.ConvertType(out[1], new([16]byte)).(*[16]byte)

	return *outstruct, err

}

// PriceHistory is a free data retrieval call binding the contract method 0xb27a0484.
//
// Solidity: function priceHistory(uint256 ) view returns(bytes16 longPrice, bytes16 shortPrice)
func (_Poolcommitter *PoolcommitterSession) PriceHistory(arg0 *big.Int) (struct {
	LongPrice  [16]byte
	ShortPrice [16]byte
}, error) {
	return _Poolcommitter.Contract.PriceHistory(&_Poolcommitter.CallOpts, arg0)
}

// PriceHistory is a free data retrieval call binding the contract method 0xb27a0484.
//
// Solidity: function priceHistory(uint256 ) view returns(bytes16 longPrice, bytes16 shortPrice)
func (_Poolcommitter *PoolcommitterCallerSession) PriceHistory(arg0 *big.Int) (struct {
	LongPrice  [16]byte
	ShortPrice [16]byte
}, error) {
	return _Poolcommitter.Contract.PriceHistory(&_Poolcommitter.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Poolcommitter *PoolcommitterCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Poolcommitter *PoolcommitterSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Poolcommitter.Contract.Tokens(&_Poolcommitter.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Poolcommitter *PoolcommitterCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Poolcommitter.Contract.Tokens(&_Poolcommitter.CallOpts, arg0)
}

// TotalPoolCommitments is a free data retrieval call binding the contract method 0xfe589858.
//
// Solidity: function totalPoolCommitments(uint256 ) view returns(uint256 longMintSettlement, uint256 longBurnPoolTokens, uint256 shortMintSettlement, uint256 shortBurnPoolTokens, uint256 shortBurnLongMintPoolTokens, uint256 longBurnShortMintPoolTokens)
func (_Poolcommitter *PoolcommitterCaller) TotalPoolCommitments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
}, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "totalPoolCommitments", arg0)

	outstruct := new(struct {
		LongMintSettlement          *big.Int
		LongBurnPoolTokens          *big.Int
		ShortMintSettlement         *big.Int
		ShortBurnPoolTokens         *big.Int
		ShortBurnLongMintPoolTokens *big.Int
		LongBurnShortMintPoolTokens *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LongMintSettlement = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LongBurnPoolTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ShortMintSettlement = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortBurnPoolTokens = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ShortBurnLongMintPoolTokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LongBurnShortMintPoolTokens = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TotalPoolCommitments is a free data retrieval call binding the contract method 0xfe589858.
//
// Solidity: function totalPoolCommitments(uint256 ) view returns(uint256 longMintSettlement, uint256 longBurnPoolTokens, uint256 shortMintSettlement, uint256 shortBurnPoolTokens, uint256 shortBurnLongMintPoolTokens, uint256 longBurnShortMintPoolTokens)
func (_Poolcommitter *PoolcommitterSession) TotalPoolCommitments(arg0 *big.Int) (struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
}, error) {
	return _Poolcommitter.Contract.TotalPoolCommitments(&_Poolcommitter.CallOpts, arg0)
}

// TotalPoolCommitments is a free data retrieval call binding the contract method 0xfe589858.
//
// Solidity: function totalPoolCommitments(uint256 ) view returns(uint256 longMintSettlement, uint256 longBurnPoolTokens, uint256 shortMintSettlement, uint256 shortBurnPoolTokens, uint256 shortBurnLongMintPoolTokens, uint256 longBurnShortMintPoolTokens)
func (_Poolcommitter *PoolcommitterCallerSession) TotalPoolCommitments(arg0 *big.Int) (struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
}, error) {
	return _Poolcommitter.Contract.TotalPoolCommitments(&_Poolcommitter.CallOpts, arg0)
}

// UnAggregatedCommitments is a free data retrieval call binding the contract method 0xb9fc75b0.
//
// Solidity: function unAggregatedCommitments(address , uint256 ) view returns(uint256)
func (_Poolcommitter *PoolcommitterCaller) UnAggregatedCommitments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "unAggregatedCommitments", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnAggregatedCommitments is a free data retrieval call binding the contract method 0xb9fc75b0.
//
// Solidity: function unAggregatedCommitments(address , uint256 ) view returns(uint256)
func (_Poolcommitter *PoolcommitterSession) UnAggregatedCommitments(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Poolcommitter.Contract.UnAggregatedCommitments(&_Poolcommitter.CallOpts, arg0, arg1)
}

// UnAggregatedCommitments is a free data retrieval call binding the contract method 0xb9fc75b0.
//
// Solidity: function unAggregatedCommitments(address , uint256 ) view returns(uint256)
func (_Poolcommitter *PoolcommitterCallerSession) UnAggregatedCommitments(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Poolcommitter.Contract.UnAggregatedCommitments(&_Poolcommitter.CallOpts, arg0, arg1)
}

// UpdateIntervalId is a free data retrieval call binding the contract method 0x9d03132e.
//
// Solidity: function updateIntervalId() view returns(uint128)
func (_Poolcommitter *PoolcommitterCaller) UpdateIntervalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "updateIntervalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateIntervalId is a free data retrieval call binding the contract method 0x9d03132e.
//
// Solidity: function updateIntervalId() view returns(uint128)
func (_Poolcommitter *PoolcommitterSession) UpdateIntervalId() (*big.Int, error) {
	return _Poolcommitter.Contract.UpdateIntervalId(&_Poolcommitter.CallOpts)
}

// UpdateIntervalId is a free data retrieval call binding the contract method 0x9d03132e.
//
// Solidity: function updateIntervalId() view returns(uint128)
func (_Poolcommitter *PoolcommitterCallerSession) UpdateIntervalId() (*big.Int, error) {
	return _Poolcommitter.Contract.UpdateIntervalId(&_Poolcommitter.CallOpts)
}

// UserAggregateBalance is a free data retrieval call binding the contract method 0xea348c9a.
//
// Solidity: function userAggregateBalance(address ) view returns(uint256 longTokens, uint256 shortTokens, uint256 settlementTokens)
func (_Poolcommitter *PoolcommitterCaller) UserAggregateBalance(opts *bind.CallOpts, arg0 common.Address) (struct {
	LongTokens       *big.Int
	ShortTokens      *big.Int
	SettlementTokens *big.Int
}, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "userAggregateBalance", arg0)

	outstruct := new(struct {
		LongTokens       *big.Int
		ShortTokens      *big.Int
		SettlementTokens *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LongTokens = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ShortTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SettlementTokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserAggregateBalance is a free data retrieval call binding the contract method 0xea348c9a.
//
// Solidity: function userAggregateBalance(address ) view returns(uint256 longTokens, uint256 shortTokens, uint256 settlementTokens)
func (_Poolcommitter *PoolcommitterSession) UserAggregateBalance(arg0 common.Address) (struct {
	LongTokens       *big.Int
	ShortTokens      *big.Int
	SettlementTokens *big.Int
}, error) {
	return _Poolcommitter.Contract.UserAggregateBalance(&_Poolcommitter.CallOpts, arg0)
}

// UserAggregateBalance is a free data retrieval call binding the contract method 0xea348c9a.
//
// Solidity: function userAggregateBalance(address ) view returns(uint256 longTokens, uint256 shortTokens, uint256 settlementTokens)
func (_Poolcommitter *PoolcommitterCallerSession) UserAggregateBalance(arg0 common.Address) (struct {
	LongTokens       *big.Int
	ShortTokens      *big.Int
	SettlementTokens *big.Int
}, error) {
	return _Poolcommitter.Contract.UserAggregateBalance(&_Poolcommitter.CallOpts, arg0)
}

// UserCommitments is a free data retrieval call binding the contract method 0x9a546848.
//
// Solidity: function userCommitments(address , uint256 ) view returns(uint256 longMintSettlement, uint256 longBurnPoolTokens, uint256 shortMintSettlement, uint256 shortBurnPoolTokens, uint256 shortBurnLongMintPoolTokens, uint256 longBurnShortMintPoolTokens, uint256 updateIntervalId)
func (_Poolcommitter *PoolcommitterCaller) UserCommitments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
	UpdateIntervalId            *big.Int
}, error) {
	var out []interface{}
	err := _Poolcommitter.contract.Call(opts, &out, "userCommitments", arg0, arg1)

	outstruct := new(struct {
		LongMintSettlement          *big.Int
		LongBurnPoolTokens          *big.Int
		ShortMintSettlement         *big.Int
		ShortBurnPoolTokens         *big.Int
		ShortBurnLongMintPoolTokens *big.Int
		LongBurnShortMintPoolTokens *big.Int
		UpdateIntervalId            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LongMintSettlement = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LongBurnPoolTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ShortMintSettlement = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ShortBurnPoolTokens = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ShortBurnLongMintPoolTokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LongBurnShortMintPoolTokens = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UpdateIntervalId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserCommitments is a free data retrieval call binding the contract method 0x9a546848.
//
// Solidity: function userCommitments(address , uint256 ) view returns(uint256 longMintSettlement, uint256 longBurnPoolTokens, uint256 shortMintSettlement, uint256 shortBurnPoolTokens, uint256 shortBurnLongMintPoolTokens, uint256 longBurnShortMintPoolTokens, uint256 updateIntervalId)
func (_Poolcommitter *PoolcommitterSession) UserCommitments(arg0 common.Address, arg1 *big.Int) (struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
	UpdateIntervalId            *big.Int
}, error) {
	return _Poolcommitter.Contract.UserCommitments(&_Poolcommitter.CallOpts, arg0, arg1)
}

// UserCommitments is a free data retrieval call binding the contract method 0x9a546848.
//
// Solidity: function userCommitments(address , uint256 ) view returns(uint256 longMintSettlement, uint256 longBurnPoolTokens, uint256 shortMintSettlement, uint256 shortBurnPoolTokens, uint256 shortBurnLongMintPoolTokens, uint256 longBurnShortMintPoolTokens, uint256 updateIntervalId)
func (_Poolcommitter *PoolcommitterCallerSession) UserCommitments(arg0 common.Address, arg1 *big.Int) (struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
	UpdateIntervalId            *big.Int
}, error) {
	return _Poolcommitter.Contract.UserCommitments(&_Poolcommitter.CallOpts, arg0, arg1)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address user) returns()
func (_Poolcommitter *PoolcommitterTransactor) Claim(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "claim", user)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address user) returns()
func (_Poolcommitter *PoolcommitterSession) Claim(user common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.Claim(&_Poolcommitter.TransactOpts, user)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address user) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) Claim(user common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.Claim(&_Poolcommitter.TransactOpts, user)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 args) payable returns()
func (_Poolcommitter *PoolcommitterTransactor) Commit(opts *bind.TransactOpts, args [32]byte) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "commit", args)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 args) payable returns()
func (_Poolcommitter *PoolcommitterSession) Commit(args [32]byte) (*types.Transaction, error) {
	return _Poolcommitter.Contract.Commit(&_Poolcommitter.TransactOpts, args)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 args) payable returns()
func (_Poolcommitter *PoolcommitterTransactorSession) Commit(args [32]byte) (*types.Transaction, error) {
	return _Poolcommitter.Contract.Commit(&_Poolcommitter.TransactOpts, args)
}

// ExecuteCommitments is a paid mutator transaction binding the contract method 0x8f5b8e87.
//
// Solidity: function executeCommitments(uint256 lastPriceTimestamp, uint256 updateInterval, uint256 longBalance, uint256 shortBalance) returns(uint256, uint256, uint256, uint256, uint256)
func (_Poolcommitter *PoolcommitterTransactor) ExecuteCommitments(opts *bind.TransactOpts, lastPriceTimestamp *big.Int, updateInterval *big.Int, longBalance *big.Int, shortBalance *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "executeCommitments", lastPriceTimestamp, updateInterval, longBalance, shortBalance)
}

// ExecuteCommitments is a paid mutator transaction binding the contract method 0x8f5b8e87.
//
// Solidity: function executeCommitments(uint256 lastPriceTimestamp, uint256 updateInterval, uint256 longBalance, uint256 shortBalance) returns(uint256, uint256, uint256, uint256, uint256)
func (_Poolcommitter *PoolcommitterSession) ExecuteCommitments(lastPriceTimestamp *big.Int, updateInterval *big.Int, longBalance *big.Int, shortBalance *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.ExecuteCommitments(&_Poolcommitter.TransactOpts, lastPriceTimestamp, updateInterval, longBalance, shortBalance)
}

// ExecuteCommitments is a paid mutator transaction binding the contract method 0x8f5b8e87.
//
// Solidity: function executeCommitments(uint256 lastPriceTimestamp, uint256 updateInterval, uint256 longBalance, uint256 shortBalance) returns(uint256, uint256, uint256, uint256, uint256)
func (_Poolcommitter *PoolcommitterTransactorSession) ExecuteCommitments(lastPriceTimestamp *big.Int, updateInterval *big.Int, longBalance *big.Int, shortBalance *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.ExecuteCommitments(&_Poolcommitter.TransactOpts, lastPriceTimestamp, updateInterval, longBalance, shortBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0x7403c6cd.
//
// Solidity: function initialize(address _factory, address _autoClaim, address _factoryOwner, address _feeController, address _invariantCheck, uint256 _mintingFee, uint256 _burningFee, uint256 _changeInterval) returns()
func (_Poolcommitter *PoolcommitterTransactor) Initialize(opts *bind.TransactOpts, _factory common.Address, _autoClaim common.Address, _factoryOwner common.Address, _feeController common.Address, _invariantCheck common.Address, _mintingFee *big.Int, _burningFee *big.Int, _changeInterval *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "initialize", _factory, _autoClaim, _factoryOwner, _feeController, _invariantCheck, _mintingFee, _burningFee, _changeInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0x7403c6cd.
//
// Solidity: function initialize(address _factory, address _autoClaim, address _factoryOwner, address _feeController, address _invariantCheck, uint256 _mintingFee, uint256 _burningFee, uint256 _changeInterval) returns()
func (_Poolcommitter *PoolcommitterSession) Initialize(_factory common.Address, _autoClaim common.Address, _factoryOwner common.Address, _feeController common.Address, _invariantCheck common.Address, _mintingFee *big.Int, _burningFee *big.Int, _changeInterval *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.Initialize(&_Poolcommitter.TransactOpts, _factory, _autoClaim, _factoryOwner, _feeController, _invariantCheck, _mintingFee, _burningFee, _changeInterval)
}

// Initialize is a paid mutator transaction binding the contract method 0x7403c6cd.
//
// Solidity: function initialize(address _factory, address _autoClaim, address _factoryOwner, address _feeController, address _invariantCheck, uint256 _mintingFee, uint256 _burningFee, uint256 _changeInterval) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) Initialize(_factory common.Address, _autoClaim common.Address, _factoryOwner common.Address, _feeController common.Address, _invariantCheck common.Address, _mintingFee *big.Int, _burningFee *big.Int, _changeInterval *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.Initialize(&_Poolcommitter.TransactOpts, _factory, _autoClaim, _factoryOwner, _feeController, _invariantCheck, _mintingFee, _burningFee, _changeInterval)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Poolcommitter *PoolcommitterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Poolcommitter *PoolcommitterSession) Pause() (*types.Transaction, error) {
	return _Poolcommitter.Contract.Pause(&_Poolcommitter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Poolcommitter *PoolcommitterTransactorSession) Pause() (*types.Transaction, error) {
	return _Poolcommitter.Contract.Pause(&_Poolcommitter.TransactOpts)
}

// SetBurningFee is a paid mutator transaction binding the contract method 0xf9dc9720.
//
// Solidity: function setBurningFee(uint256 _burningFee) returns()
func (_Poolcommitter *PoolcommitterTransactor) SetBurningFee(opts *bind.TransactOpts, _burningFee *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "setBurningFee", _burningFee)
}

// SetBurningFee is a paid mutator transaction binding the contract method 0xf9dc9720.
//
// Solidity: function setBurningFee(uint256 _burningFee) returns()
func (_Poolcommitter *PoolcommitterSession) SetBurningFee(_burningFee *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetBurningFee(&_Poolcommitter.TransactOpts, _burningFee)
}

// SetBurningFee is a paid mutator transaction binding the contract method 0xf9dc9720.
//
// Solidity: function setBurningFee(uint256 _burningFee) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) SetBurningFee(_burningFee *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetBurningFee(&_Poolcommitter.TransactOpts, _burningFee)
}

// SetChangeInterval is a paid mutator transaction binding the contract method 0x384564e3.
//
// Solidity: function setChangeInterval(uint256 _changeInterval) returns()
func (_Poolcommitter *PoolcommitterTransactor) SetChangeInterval(opts *bind.TransactOpts, _changeInterval *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "setChangeInterval", _changeInterval)
}

// SetChangeInterval is a paid mutator transaction binding the contract method 0x384564e3.
//
// Solidity: function setChangeInterval(uint256 _changeInterval) returns()
func (_Poolcommitter *PoolcommitterSession) SetChangeInterval(_changeInterval *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetChangeInterval(&_Poolcommitter.TransactOpts, _changeInterval)
}

// SetChangeInterval is a paid mutator transaction binding the contract method 0x384564e3.
//
// Solidity: function setChangeInterval(uint256 _changeInterval) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) SetChangeInterval(_changeInterval *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetChangeInterval(&_Poolcommitter.TransactOpts, _changeInterval)
}

// SetFeeController is a paid mutator transaction binding the contract method 0x3ed4c678.
//
// Solidity: function setFeeController(address _feeController) returns()
func (_Poolcommitter *PoolcommitterTransactor) SetFeeController(opts *bind.TransactOpts, _feeController common.Address) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "setFeeController", _feeController)
}

// SetFeeController is a paid mutator transaction binding the contract method 0x3ed4c678.
//
// Solidity: function setFeeController(address _feeController) returns()
func (_Poolcommitter *PoolcommitterSession) SetFeeController(_feeController common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetFeeController(&_Poolcommitter.TransactOpts, _feeController)
}

// SetFeeController is a paid mutator transaction binding the contract method 0x3ed4c678.
//
// Solidity: function setFeeController(address _feeController) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) SetFeeController(_feeController common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetFeeController(&_Poolcommitter.TransactOpts, _feeController)
}

// SetMintingFee is a paid mutator transaction binding the contract method 0x238a4709.
//
// Solidity: function setMintingFee(uint256 _mintingFee) returns()
func (_Poolcommitter *PoolcommitterTransactor) SetMintingFee(opts *bind.TransactOpts, _mintingFee *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "setMintingFee", _mintingFee)
}

// SetMintingFee is a paid mutator transaction binding the contract method 0x238a4709.
//
// Solidity: function setMintingFee(uint256 _mintingFee) returns()
func (_Poolcommitter *PoolcommitterSession) SetMintingFee(_mintingFee *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetMintingFee(&_Poolcommitter.TransactOpts, _mintingFee)
}

// SetMintingFee is a paid mutator transaction binding the contract method 0x238a4709.
//
// Solidity: function setMintingFee(uint256 _mintingFee) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) SetMintingFee(_mintingFee *big.Int) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetMintingFee(&_Poolcommitter.TransactOpts, _mintingFee)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _leveragedPool) returns()
func (_Poolcommitter *PoolcommitterTransactor) SetPool(opts *bind.TransactOpts, _leveragedPool common.Address) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "setPool", _leveragedPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _leveragedPool) returns()
func (_Poolcommitter *PoolcommitterSession) SetPool(_leveragedPool common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetPool(&_Poolcommitter.TransactOpts, _leveragedPool)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address _leveragedPool) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) SetPool(_leveragedPool common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.SetPool(&_Poolcommitter.TransactOpts, _leveragedPool)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Poolcommitter *PoolcommitterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Poolcommitter *PoolcommitterSession) Unpause() (*types.Transaction, error) {
	return _Poolcommitter.Contract.Unpause(&_Poolcommitter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Poolcommitter *PoolcommitterTransactorSession) Unpause() (*types.Transaction, error) {
	return _Poolcommitter.Contract.Unpause(&_Poolcommitter.TransactOpts)
}

// UpdateAggregateBalance is a paid mutator transaction binding the contract method 0x9fa21dca.
//
// Solidity: function updateAggregateBalance(address user) returns()
func (_Poolcommitter *PoolcommitterTransactor) UpdateAggregateBalance(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Poolcommitter.contract.Transact(opts, "updateAggregateBalance", user)
}

// UpdateAggregateBalance is a paid mutator transaction binding the contract method 0x9fa21dca.
//
// Solidity: function updateAggregateBalance(address user) returns()
func (_Poolcommitter *PoolcommitterSession) UpdateAggregateBalance(user common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.UpdateAggregateBalance(&_Poolcommitter.TransactOpts, user)
}

// UpdateAggregateBalance is a paid mutator transaction binding the contract method 0x9fa21dca.
//
// Solidity: function updateAggregateBalance(address user) returns()
func (_Poolcommitter *PoolcommitterTransactorSession) UpdateAggregateBalance(user common.Address) (*types.Transaction, error) {
	return _Poolcommitter.Contract.UpdateAggregateBalance(&_Poolcommitter.TransactOpts, user)
}

// PoolcommitterAggregateBalanceUpdatedIterator is returned from FilterAggregateBalanceUpdated and is used to iterate over the raw logs and unpacked data for AggregateBalanceUpdated events raised by the Poolcommitter contract.
type PoolcommitterAggregateBalanceUpdatedIterator struct {
	Event *PoolcommitterAggregateBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *PoolcommitterAggregateBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterAggregateBalanceUpdated)
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
		it.Event = new(PoolcommitterAggregateBalanceUpdated)
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
func (it *PoolcommitterAggregateBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterAggregateBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterAggregateBalanceUpdated represents a AggregateBalanceUpdated event raised by the Poolcommitter contract.
type PoolcommitterAggregateBalanceUpdated struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAggregateBalanceUpdated is a free log retrieval operation binding the contract event 0xdefa4140341b4fd283af41700779361dc94f34ae245a186b3b9c187033c06554.
//
// Solidity: event AggregateBalanceUpdated(address indexed user)
func (_Poolcommitter *PoolcommitterFilterer) FilterAggregateBalanceUpdated(opts *bind.FilterOpts, user []common.Address) (*PoolcommitterAggregateBalanceUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "AggregateBalanceUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterAggregateBalanceUpdatedIterator{contract: _Poolcommitter.contract, event: "AggregateBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregateBalanceUpdated is a free log subscription operation binding the contract event 0xdefa4140341b4fd283af41700779361dc94f34ae245a186b3b9c187033c06554.
//
// Solidity: event AggregateBalanceUpdated(address indexed user)
func (_Poolcommitter *PoolcommitterFilterer) WatchAggregateBalanceUpdated(opts *bind.WatchOpts, sink chan<- *PoolcommitterAggregateBalanceUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "AggregateBalanceUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterAggregateBalanceUpdated)
				if err := _Poolcommitter.contract.UnpackLog(event, "AggregateBalanceUpdated", log); err != nil {
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

// ParseAggregateBalanceUpdated is a log parse operation binding the contract event 0xdefa4140341b4fd283af41700779361dc94f34ae245a186b3b9c187033c06554.
//
// Solidity: event AggregateBalanceUpdated(address indexed user)
func (_Poolcommitter *PoolcommitterFilterer) ParseAggregateBalanceUpdated(log types.Log) (*PoolcommitterAggregateBalanceUpdated, error) {
	event := new(PoolcommitterAggregateBalanceUpdated)
	if err := _Poolcommitter.contract.UnpackLog(event, "AggregateBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterBurningFeeSetIterator is returned from FilterBurningFeeSet and is used to iterate over the raw logs and unpacked data for BurningFeeSet events raised by the Poolcommitter contract.
type PoolcommitterBurningFeeSetIterator struct {
	Event *PoolcommitterBurningFeeSet // Event containing the contract specifics and raw log

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
func (it *PoolcommitterBurningFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterBurningFeeSet)
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
		it.Event = new(PoolcommitterBurningFeeSet)
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
func (it *PoolcommitterBurningFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterBurningFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterBurningFeeSet represents a BurningFeeSet event raised by the Poolcommitter contract.
type PoolcommitterBurningFeeSet struct {
	BurningFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurningFeeSet is a free log retrieval operation binding the contract event 0x9bd544d4dd4e5359b6126812d7f52828638adf3d7beab61310fab58b27d811cc.
//
// Solidity: event BurningFeeSet(uint256 indexed _burningFee)
func (_Poolcommitter *PoolcommitterFilterer) FilterBurningFeeSet(opts *bind.FilterOpts, _burningFee []*big.Int) (*PoolcommitterBurningFeeSetIterator, error) {

	var _burningFeeRule []interface{}
	for _, _burningFeeItem := range _burningFee {
		_burningFeeRule = append(_burningFeeRule, _burningFeeItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "BurningFeeSet", _burningFeeRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterBurningFeeSetIterator{contract: _Poolcommitter.contract, event: "BurningFeeSet", logs: logs, sub: sub}, nil
}

// WatchBurningFeeSet is a free log subscription operation binding the contract event 0x9bd544d4dd4e5359b6126812d7f52828638adf3d7beab61310fab58b27d811cc.
//
// Solidity: event BurningFeeSet(uint256 indexed _burningFee)
func (_Poolcommitter *PoolcommitterFilterer) WatchBurningFeeSet(opts *bind.WatchOpts, sink chan<- *PoolcommitterBurningFeeSet, _burningFee []*big.Int) (event.Subscription, error) {

	var _burningFeeRule []interface{}
	for _, _burningFeeItem := range _burningFee {
		_burningFeeRule = append(_burningFeeRule, _burningFeeItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "BurningFeeSet", _burningFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterBurningFeeSet)
				if err := _Poolcommitter.contract.UnpackLog(event, "BurningFeeSet", log); err != nil {
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

// ParseBurningFeeSet is a log parse operation binding the contract event 0x9bd544d4dd4e5359b6126812d7f52828638adf3d7beab61310fab58b27d811cc.
//
// Solidity: event BurningFeeSet(uint256 indexed _burningFee)
func (_Poolcommitter *PoolcommitterFilterer) ParseBurningFeeSet(log types.Log) (*PoolcommitterBurningFeeSet, error) {
	event := new(PoolcommitterBurningFeeSet)
	if err := _Poolcommitter.contract.UnpackLog(event, "BurningFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterChangeIntervalSetIterator is returned from FilterChangeIntervalSet and is used to iterate over the raw logs and unpacked data for ChangeIntervalSet events raised by the Poolcommitter contract.
type PoolcommitterChangeIntervalSetIterator struct {
	Event *PoolcommitterChangeIntervalSet // Event containing the contract specifics and raw log

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
func (it *PoolcommitterChangeIntervalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterChangeIntervalSet)
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
		it.Event = new(PoolcommitterChangeIntervalSet)
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
func (it *PoolcommitterChangeIntervalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterChangeIntervalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterChangeIntervalSet represents a ChangeIntervalSet event raised by the Poolcommitter contract.
type PoolcommitterChangeIntervalSet struct {
	ChangeInterval *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangeIntervalSet is a free log retrieval operation binding the contract event 0x1f609e633399cb8530717d1970d31255fae2af322c2c3b90793cc1c1a7fe88be.
//
// Solidity: event ChangeIntervalSet(uint256 indexed _changeInterval)
func (_Poolcommitter *PoolcommitterFilterer) FilterChangeIntervalSet(opts *bind.FilterOpts, _changeInterval []*big.Int) (*PoolcommitterChangeIntervalSetIterator, error) {

	var _changeIntervalRule []interface{}
	for _, _changeIntervalItem := range _changeInterval {
		_changeIntervalRule = append(_changeIntervalRule, _changeIntervalItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "ChangeIntervalSet", _changeIntervalRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterChangeIntervalSetIterator{contract: _Poolcommitter.contract, event: "ChangeIntervalSet", logs: logs, sub: sub}, nil
}

// WatchChangeIntervalSet is a free log subscription operation binding the contract event 0x1f609e633399cb8530717d1970d31255fae2af322c2c3b90793cc1c1a7fe88be.
//
// Solidity: event ChangeIntervalSet(uint256 indexed _changeInterval)
func (_Poolcommitter *PoolcommitterFilterer) WatchChangeIntervalSet(opts *bind.WatchOpts, sink chan<- *PoolcommitterChangeIntervalSet, _changeInterval []*big.Int) (event.Subscription, error) {

	var _changeIntervalRule []interface{}
	for _, _changeIntervalItem := range _changeInterval {
		_changeIntervalRule = append(_changeIntervalRule, _changeIntervalItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "ChangeIntervalSet", _changeIntervalRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterChangeIntervalSet)
				if err := _Poolcommitter.contract.UnpackLog(event, "ChangeIntervalSet", log); err != nil {
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

// ParseChangeIntervalSet is a log parse operation binding the contract event 0x1f609e633399cb8530717d1970d31255fae2af322c2c3b90793cc1c1a7fe88be.
//
// Solidity: event ChangeIntervalSet(uint256 indexed _changeInterval)
func (_Poolcommitter *PoolcommitterFilterer) ParseChangeIntervalSet(log types.Log) (*PoolcommitterChangeIntervalSet, error) {
	event := new(PoolcommitterChangeIntervalSet)
	if err := _Poolcommitter.contract.UnpackLog(event, "ChangeIntervalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Poolcommitter contract.
type PoolcommitterClaimIterator struct {
	Event *PoolcommitterClaim // Event containing the contract specifics and raw log

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
func (it *PoolcommitterClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterClaim)
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
		it.Event = new(PoolcommitterClaim)
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
func (it *PoolcommitterClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterClaim represents a Claim event raised by the Poolcommitter contract.
type PoolcommitterClaim struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address indexed user)
func (_Poolcommitter *PoolcommitterFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*PoolcommitterClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterClaimIterator{contract: _Poolcommitter.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address indexed user)
func (_Poolcommitter *PoolcommitterFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *PoolcommitterClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterClaim)
				if err := _Poolcommitter.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x0c7ef932d3b91976772937f18d5ef9b39a9930bef486b576c374f047c4b512dc.
//
// Solidity: event Claim(address indexed user)
func (_Poolcommitter *PoolcommitterFilterer) ParseClaim(log types.Log) (*PoolcommitterClaim, error) {
	event := new(PoolcommitterClaim)
	if err := _Poolcommitter.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterCreateCommitIterator is returned from FilterCreateCommit and is used to iterate over the raw logs and unpacked data for CreateCommit events raised by the Poolcommitter contract.
type PoolcommitterCreateCommitIterator struct {
	Event *PoolcommitterCreateCommit // Event containing the contract specifics and raw log

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
func (it *PoolcommitterCreateCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterCreateCommit)
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
		it.Event = new(PoolcommitterCreateCommit)
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
func (it *PoolcommitterCreateCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterCreateCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterCreateCommit represents a CreateCommit event raised by the Poolcommitter contract.
type PoolcommitterCreateCommit struct {
	User                        common.Address
	Amount                      *big.Int
	CommitType                  uint8
	AppropriateUpdateIntervalId *big.Int
	FromAggregateBalance        bool
	PayForClaim                 bool
	MintingFee                  [16]byte
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterCreateCommit is a free log retrieval operation binding the contract event 0x04f9dae5361dfaa5b3f324fc0af5e71bcb4d76e8324ad9ce78d70496c548d503.
//
// Solidity: event CreateCommit(address indexed user, uint256 indexed amount, uint8 indexed commitType, uint256 appropriateUpdateIntervalId, bool fromAggregateBalance, bool payForClaim, bytes16 mintingFee)
func (_Poolcommitter *PoolcommitterFilterer) FilterCreateCommit(opts *bind.FilterOpts, user []common.Address, amount []*big.Int, commitType []uint8) (*PoolcommitterCreateCommitIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var commitTypeRule []interface{}
	for _, commitTypeItem := range commitType {
		commitTypeRule = append(commitTypeRule, commitTypeItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "CreateCommit", userRule, amountRule, commitTypeRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterCreateCommitIterator{contract: _Poolcommitter.contract, event: "CreateCommit", logs: logs, sub: sub}, nil
}

// WatchCreateCommit is a free log subscription operation binding the contract event 0x04f9dae5361dfaa5b3f324fc0af5e71bcb4d76e8324ad9ce78d70496c548d503.
//
// Solidity: event CreateCommit(address indexed user, uint256 indexed amount, uint8 indexed commitType, uint256 appropriateUpdateIntervalId, bool fromAggregateBalance, bool payForClaim, bytes16 mintingFee)
func (_Poolcommitter *PoolcommitterFilterer) WatchCreateCommit(opts *bind.WatchOpts, sink chan<- *PoolcommitterCreateCommit, user []common.Address, amount []*big.Int, commitType []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var commitTypeRule []interface{}
	for _, commitTypeItem := range commitType {
		commitTypeRule = append(commitTypeRule, commitTypeItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "CreateCommit", userRule, amountRule, commitTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterCreateCommit)
				if err := _Poolcommitter.contract.UnpackLog(event, "CreateCommit", log); err != nil {
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

// ParseCreateCommit is a log parse operation binding the contract event 0x04f9dae5361dfaa5b3f324fc0af5e71bcb4d76e8324ad9ce78d70496c548d503.
//
// Solidity: event CreateCommit(address indexed user, uint256 indexed amount, uint8 indexed commitType, uint256 appropriateUpdateIntervalId, bool fromAggregateBalance, bool payForClaim, bytes16 mintingFee)
func (_Poolcommitter *PoolcommitterFilterer) ParseCreateCommit(log types.Log) (*PoolcommitterCreateCommit, error) {
	event := new(PoolcommitterCreateCommit)
	if err := _Poolcommitter.contract.UnpackLog(event, "CreateCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterExecutedCommitsForIntervalIterator is returned from FilterExecutedCommitsForInterval and is used to iterate over the raw logs and unpacked data for ExecutedCommitsForInterval events raised by the Poolcommitter contract.
type PoolcommitterExecutedCommitsForIntervalIterator struct {
	Event *PoolcommitterExecutedCommitsForInterval // Event containing the contract specifics and raw log

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
func (it *PoolcommitterExecutedCommitsForIntervalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterExecutedCommitsForInterval)
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
		it.Event = new(PoolcommitterExecutedCommitsForInterval)
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
func (it *PoolcommitterExecutedCommitsForIntervalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterExecutedCommitsForIntervalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterExecutedCommitsForInterval represents a ExecutedCommitsForInterval event raised by the Poolcommitter contract.
type PoolcommitterExecutedCommitsForInterval struct {
	UpdateIntervalId *big.Int
	BurningFee       [16]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutedCommitsForInterval is a free log retrieval operation binding the contract event 0x260463f9bbb6d40c58e9fc4e09a71c3f80fc18d818d6d9d5df86b2443722a59f.
//
// Solidity: event ExecutedCommitsForInterval(uint256 indexed updateIntervalId, bytes16 burningFee)
func (_Poolcommitter *PoolcommitterFilterer) FilterExecutedCommitsForInterval(opts *bind.FilterOpts, updateIntervalId []*big.Int) (*PoolcommitterExecutedCommitsForIntervalIterator, error) {

	var updateIntervalIdRule []interface{}
	for _, updateIntervalIdItem := range updateIntervalId {
		updateIntervalIdRule = append(updateIntervalIdRule, updateIntervalIdItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "ExecutedCommitsForInterval", updateIntervalIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterExecutedCommitsForIntervalIterator{contract: _Poolcommitter.contract, event: "ExecutedCommitsForInterval", logs: logs, sub: sub}, nil
}

// WatchExecutedCommitsForInterval is a free log subscription operation binding the contract event 0x260463f9bbb6d40c58e9fc4e09a71c3f80fc18d818d6d9d5df86b2443722a59f.
//
// Solidity: event ExecutedCommitsForInterval(uint256 indexed updateIntervalId, bytes16 burningFee)
func (_Poolcommitter *PoolcommitterFilterer) WatchExecutedCommitsForInterval(opts *bind.WatchOpts, sink chan<- *PoolcommitterExecutedCommitsForInterval, updateIntervalId []*big.Int) (event.Subscription, error) {

	var updateIntervalIdRule []interface{}
	for _, updateIntervalIdItem := range updateIntervalId {
		updateIntervalIdRule = append(updateIntervalIdRule, updateIntervalIdItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "ExecutedCommitsForInterval", updateIntervalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterExecutedCommitsForInterval)
				if err := _Poolcommitter.contract.UnpackLog(event, "ExecutedCommitsForInterval", log); err != nil {
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

// ParseExecutedCommitsForInterval is a log parse operation binding the contract event 0x260463f9bbb6d40c58e9fc4e09a71c3f80fc18d818d6d9d5df86b2443722a59f.
//
// Solidity: event ExecutedCommitsForInterval(uint256 indexed updateIntervalId, bytes16 burningFee)
func (_Poolcommitter *PoolcommitterFilterer) ParseExecutedCommitsForInterval(log types.Log) (*PoolcommitterExecutedCommitsForInterval, error) {
	event := new(PoolcommitterExecutedCommitsForInterval)
	if err := _Poolcommitter.contract.UnpackLog(event, "ExecutedCommitsForInterval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterFeeControllerSetIterator is returned from FilterFeeControllerSet and is used to iterate over the raw logs and unpacked data for FeeControllerSet events raised by the Poolcommitter contract.
type PoolcommitterFeeControllerSetIterator struct {
	Event *PoolcommitterFeeControllerSet // Event containing the contract specifics and raw log

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
func (it *PoolcommitterFeeControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterFeeControllerSet)
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
		it.Event = new(PoolcommitterFeeControllerSet)
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
func (it *PoolcommitterFeeControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterFeeControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterFeeControllerSet represents a FeeControllerSet event raised by the Poolcommitter contract.
type PoolcommitterFeeControllerSet struct {
	FeeController common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeControllerSet is a free log retrieval operation binding the contract event 0x4645d48e2e83635dc48865c27880f748bc610aab2e41ca3358057b565ff56a57.
//
// Solidity: event FeeControllerSet(address indexed _feeController)
func (_Poolcommitter *PoolcommitterFilterer) FilterFeeControllerSet(opts *bind.FilterOpts, _feeController []common.Address) (*PoolcommitterFeeControllerSetIterator, error) {

	var _feeControllerRule []interface{}
	for _, _feeControllerItem := range _feeController {
		_feeControllerRule = append(_feeControllerRule, _feeControllerItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "FeeControllerSet", _feeControllerRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterFeeControllerSetIterator{contract: _Poolcommitter.contract, event: "FeeControllerSet", logs: logs, sub: sub}, nil
}

// WatchFeeControllerSet is a free log subscription operation binding the contract event 0x4645d48e2e83635dc48865c27880f748bc610aab2e41ca3358057b565ff56a57.
//
// Solidity: event FeeControllerSet(address indexed _feeController)
func (_Poolcommitter *PoolcommitterFilterer) WatchFeeControllerSet(opts *bind.WatchOpts, sink chan<- *PoolcommitterFeeControllerSet, _feeController []common.Address) (event.Subscription, error) {

	var _feeControllerRule []interface{}
	for _, _feeControllerItem := range _feeController {
		_feeControllerRule = append(_feeControllerRule, _feeControllerItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "FeeControllerSet", _feeControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterFeeControllerSet)
				if err := _Poolcommitter.contract.UnpackLog(event, "FeeControllerSet", log); err != nil {
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

// ParseFeeControllerSet is a log parse operation binding the contract event 0x4645d48e2e83635dc48865c27880f748bc610aab2e41ca3358057b565ff56a57.
//
// Solidity: event FeeControllerSet(address indexed _feeController)
func (_Poolcommitter *PoolcommitterFilterer) ParseFeeControllerSet(log types.Log) (*PoolcommitterFeeControllerSet, error) {
	event := new(PoolcommitterFeeControllerSet)
	if err := _Poolcommitter.contract.UnpackLog(event, "FeeControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterMintingFeeSetIterator is returned from FilterMintingFeeSet and is used to iterate over the raw logs and unpacked data for MintingFeeSet events raised by the Poolcommitter contract.
type PoolcommitterMintingFeeSetIterator struct {
	Event *PoolcommitterMintingFeeSet // Event containing the contract specifics and raw log

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
func (it *PoolcommitterMintingFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterMintingFeeSet)
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
		it.Event = new(PoolcommitterMintingFeeSet)
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
func (it *PoolcommitterMintingFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterMintingFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterMintingFeeSet represents a MintingFeeSet event raised by the Poolcommitter contract.
type PoolcommitterMintingFeeSet struct {
	MintingFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintingFeeSet is a free log retrieval operation binding the contract event 0x07aefbcc0d8cdf34e241279c78d93d1df9540b5a60d320437958df12ef59e463.
//
// Solidity: event MintingFeeSet(uint256 indexed _mintingFee)
func (_Poolcommitter *PoolcommitterFilterer) FilterMintingFeeSet(opts *bind.FilterOpts, _mintingFee []*big.Int) (*PoolcommitterMintingFeeSetIterator, error) {

	var _mintingFeeRule []interface{}
	for _, _mintingFeeItem := range _mintingFee {
		_mintingFeeRule = append(_mintingFeeRule, _mintingFeeItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "MintingFeeSet", _mintingFeeRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterMintingFeeSetIterator{contract: _Poolcommitter.contract, event: "MintingFeeSet", logs: logs, sub: sub}, nil
}

// WatchMintingFeeSet is a free log subscription operation binding the contract event 0x07aefbcc0d8cdf34e241279c78d93d1df9540b5a60d320437958df12ef59e463.
//
// Solidity: event MintingFeeSet(uint256 indexed _mintingFee)
func (_Poolcommitter *PoolcommitterFilterer) WatchMintingFeeSet(opts *bind.WatchOpts, sink chan<- *PoolcommitterMintingFeeSet, _mintingFee []*big.Int) (event.Subscription, error) {

	var _mintingFeeRule []interface{}
	for _, _mintingFeeItem := range _mintingFee {
		_mintingFeeRule = append(_mintingFeeRule, _mintingFeeItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "MintingFeeSet", _mintingFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterMintingFeeSet)
				if err := _Poolcommitter.contract.UnpackLog(event, "MintingFeeSet", log); err != nil {
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

// ParseMintingFeeSet is a log parse operation binding the contract event 0x07aefbcc0d8cdf34e241279c78d93d1df9540b5a60d320437958df12ef59e463.
//
// Solidity: event MintingFeeSet(uint256 indexed _mintingFee)
func (_Poolcommitter *PoolcommitterFilterer) ParseMintingFeeSet(log types.Log) (*PoolcommitterMintingFeeSet, error) {
	event := new(PoolcommitterMintingFeeSet)
	if err := _Poolcommitter.contract.UnpackLog(event, "MintingFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Poolcommitter contract.
type PoolcommitterPausedIterator struct {
	Event *PoolcommitterPaused // Event containing the contract specifics and raw log

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
func (it *PoolcommitterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterPaused)
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
		it.Event = new(PoolcommitterPaused)
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
func (it *PoolcommitterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterPaused represents a Paused event raised by the Poolcommitter contract.
type PoolcommitterPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Poolcommitter *PoolcommitterFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolcommitterPausedIterator, error) {

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolcommitterPausedIterator{contract: _Poolcommitter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Poolcommitter *PoolcommitterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolcommitterPaused) (event.Subscription, error) {

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterPaused)
				if err := _Poolcommitter.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Poolcommitter *PoolcommitterFilterer) ParsePaused(log types.Log) (*PoolcommitterPaused, error) {
	event := new(PoolcommitterPaused)
	if err := _Poolcommitter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterPoolChangedIterator is returned from FilterPoolChanged and is used to iterate over the raw logs and unpacked data for PoolChanged events raised by the Poolcommitter contract.
type PoolcommitterPoolChangedIterator struct {
	Event *PoolcommitterPoolChanged // Event containing the contract specifics and raw log

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
func (it *PoolcommitterPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterPoolChanged)
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
		it.Event = new(PoolcommitterPoolChanged)
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
func (it *PoolcommitterPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterPoolChanged represents a PoolChanged event raised by the Poolcommitter contract.
type PoolcommitterPoolChanged struct {
	NewPool common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolChanged is a free log retrieval operation binding the contract event 0xa84ed99a050d1610bc3dddf2c220cd77842a3ab1a46664af49062a70417debdf.
//
// Solidity: event PoolChanged(address indexed newPool)
func (_Poolcommitter *PoolcommitterFilterer) FilterPoolChanged(opts *bind.FilterOpts, newPool []common.Address) (*PoolcommitterPoolChangedIterator, error) {

	var newPoolRule []interface{}
	for _, newPoolItem := range newPool {
		newPoolRule = append(newPoolRule, newPoolItem)
	}

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "PoolChanged", newPoolRule)
	if err != nil {
		return nil, err
	}
	return &PoolcommitterPoolChangedIterator{contract: _Poolcommitter.contract, event: "PoolChanged", logs: logs, sub: sub}, nil
}

// WatchPoolChanged is a free log subscription operation binding the contract event 0xa84ed99a050d1610bc3dddf2c220cd77842a3ab1a46664af49062a70417debdf.
//
// Solidity: event PoolChanged(address indexed newPool)
func (_Poolcommitter *PoolcommitterFilterer) WatchPoolChanged(opts *bind.WatchOpts, sink chan<- *PoolcommitterPoolChanged, newPool []common.Address) (event.Subscription, error) {

	var newPoolRule []interface{}
	for _, newPoolItem := range newPool {
		newPoolRule = append(newPoolRule, newPoolItem)
	}

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "PoolChanged", newPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterPoolChanged)
				if err := _Poolcommitter.contract.UnpackLog(event, "PoolChanged", log); err != nil {
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

// ParsePoolChanged is a log parse operation binding the contract event 0xa84ed99a050d1610bc3dddf2c220cd77842a3ab1a46664af49062a70417debdf.
//
// Solidity: event PoolChanged(address indexed newPool)
func (_Poolcommitter *PoolcommitterFilterer) ParsePoolChanged(log types.Log) (*PoolcommitterPoolChanged, error) {
	event := new(PoolcommitterPoolChanged)
	if err := _Poolcommitter.contract.UnpackLog(event, "PoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolcommitterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Poolcommitter contract.
type PoolcommitterUnpausedIterator struct {
	Event *PoolcommitterUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolcommitterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolcommitterUnpaused)
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
		it.Event = new(PoolcommitterUnpaused)
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
func (it *PoolcommitterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolcommitterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolcommitterUnpaused represents a Unpaused event raised by the Poolcommitter contract.
type PoolcommitterUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Poolcommitter *PoolcommitterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolcommitterUnpausedIterator, error) {

	logs, sub, err := _Poolcommitter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolcommitterUnpausedIterator{contract: _Poolcommitter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Poolcommitter *PoolcommitterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolcommitterUnpaused) (event.Subscription, error) {

	logs, sub, err := _Poolcommitter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolcommitterUnpaused)
				if err := _Poolcommitter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Poolcommitter *PoolcommitterFilterer) ParseUnpaused(log types.Log) (*PoolcommitterUnpaused, error) {
	event := new(PoolcommitterUnpaused)
	if err := _Poolcommitter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
