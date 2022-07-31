// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolstatehelper

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

// IPoolCommitterTotalCommitment is an auto generated low-level Go binding around an user-defined struct.
type IPoolCommitterTotalCommitment struct {
	LongMintSettlement          *big.Int
	LongBurnPoolTokens          *big.Int
	ShortMintSettlement         *big.Int
	ShortBurnPoolTokens         *big.Int
	ShortBurnLongMintPoolTokens *big.Int
	LongBurnShortMintPoolTokens *big.Int
}

// IPoolStateHelperExpectedPoolState is an auto generated low-level Go binding around an user-defined struct.
type IPoolStateHelperExpectedPoolState struct {
	CumulativePendingMintSettlement *big.Int
	RemainingPendingShortBurnTokens *big.Int
	RemainingPendingLongBurnTokens  *big.Int
	LongSupply                      *big.Int
	LongBalance                     *big.Int
	ShortSupply                     *big.Int
	ShortBalance                    *big.Int
	OraclePrice                     *big.Int
}

// IPoolStateHelperPoolInfo is an auto generated low-level Go binding around an user-defined struct.
type IPoolStateHelperPoolInfo struct {
	Long  IPoolStateHelperSideInfo
	Short IPoolStateHelperSideInfo
}

// IPoolStateHelperSideInfo is an auto generated low-level Go binding around an user-defined struct.
type IPoolStateHelperSideInfo struct {
	Supply                *big.Int
	SettlementBalance     *big.Int
	PendingBurnPoolTokens *big.Int
}

// PoolstatehelperMetaData contains all meta data concerning the Poolstatehelper contract.
var PoolstatehelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"INVALID_PERIOD\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractILeveragedPool2\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"fullCommitPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolCommitter2\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"periods\",\"type\":\"uint256\"}],\"name\":\"getCommitQueue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"longMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBurnPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBurnPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBurnLongMintPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBurnShortMintPoolTokens\",\"type\":\"uint256\"}],\"internalType\":\"structIPoolCommitter.TotalCommitment[]\",\"name\":\"commitQueue\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILeveragedPool2\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"periods\",\"type\":\"uint256\"}],\"name\":\"getExpectedState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cumulativePendingMintSettlement\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingPendingShortBurnTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingPendingLongBurnTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortBalance\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"oraclePrice\",\"type\":\"int256\"}],\"internalType\":\"structIPoolStateHelper.ExpectedPoolState\",\"name\":\"finalExpectedPoolState\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILeveragedPool2\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIPoolCommitter2\",\"name\":\"committer\",\"type\":\"address\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingBurnPoolTokens\",\"type\":\"uint256\"}],\"internalType\":\"structIPoolStateHelper.SideInfo\",\"name\":\"long\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"settlementBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingBurnPoolTokens\",\"type\":\"uint256\"}],\"internalType\":\"structIPoolStateHelper.SideInfo\",\"name\":\"short\",\"type\":\"tuple\"}],\"internalType\":\"structIPoolStateHelper.PoolInfo\",\"name\":\"poolInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolstatehelperABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolstatehelperMetaData.ABI instead.
var PoolstatehelperABI = PoolstatehelperMetaData.ABI

// Poolstatehelper is an auto generated Go binding around an Ethereum contract.
type Poolstatehelper struct {
	PoolstatehelperCaller     // Read-only binding to the contract
	PoolstatehelperTransactor // Write-only binding to the contract
	PoolstatehelperFilterer   // Log filterer for contract events
}

// PoolstatehelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolstatehelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolstatehelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolstatehelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolstatehelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolstatehelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolstatehelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolstatehelperSession struct {
	Contract     *Poolstatehelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolstatehelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolstatehelperCallerSession struct {
	Contract *PoolstatehelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PoolstatehelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolstatehelperTransactorSession struct {
	Contract     *PoolstatehelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PoolstatehelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolstatehelperRaw struct {
	Contract *Poolstatehelper // Generic contract binding to access the raw methods on
}

// PoolstatehelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolstatehelperCallerRaw struct {
	Contract *PoolstatehelperCaller // Generic read-only contract binding to access the raw methods on
}

// PoolstatehelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolstatehelperTransactorRaw struct {
	Contract *PoolstatehelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolstatehelper creates a new instance of Poolstatehelper, bound to a specific deployed contract.
func NewPoolstatehelper(address common.Address, backend bind.ContractBackend) (*Poolstatehelper, error) {
	contract, err := bindPoolstatehelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolstatehelper{PoolstatehelperCaller: PoolstatehelperCaller{contract: contract}, PoolstatehelperTransactor: PoolstatehelperTransactor{contract: contract}, PoolstatehelperFilterer: PoolstatehelperFilterer{contract: contract}}, nil
}

// NewPoolstatehelperCaller creates a new read-only instance of Poolstatehelper, bound to a specific deployed contract.
func NewPoolstatehelperCaller(address common.Address, caller bind.ContractCaller) (*PoolstatehelperCaller, error) {
	contract, err := bindPoolstatehelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolstatehelperCaller{contract: contract}, nil
}

// NewPoolstatehelperTransactor creates a new write-only instance of Poolstatehelper, bound to a specific deployed contract.
func NewPoolstatehelperTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolstatehelperTransactor, error) {
	contract, err := bindPoolstatehelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolstatehelperTransactor{contract: contract}, nil
}

// NewPoolstatehelperFilterer creates a new log filterer instance of Poolstatehelper, bound to a specific deployed contract.
func NewPoolstatehelperFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolstatehelperFilterer, error) {
	contract, err := bindPoolstatehelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolstatehelperFilterer{contract: contract}, nil
}

// bindPoolstatehelper binds a generic wrapper to an already deployed contract.
func bindPoolstatehelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolstatehelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolstatehelper *PoolstatehelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolstatehelper.Contract.PoolstatehelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolstatehelper *PoolstatehelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolstatehelper.Contract.PoolstatehelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolstatehelper *PoolstatehelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolstatehelper.Contract.PoolstatehelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolstatehelper *PoolstatehelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolstatehelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolstatehelper *PoolstatehelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolstatehelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolstatehelper *PoolstatehelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolstatehelper.Contract.contract.Transact(opts, method, params...)
}

// FullCommitPeriod is a free data retrieval call binding the contract method 0xb3b7d503.
//
// Solidity: function fullCommitPeriod(address pool) view returns(uint256)
func (_Poolstatehelper *PoolstatehelperCaller) FullCommitPeriod(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolstatehelper.contract.Call(opts, &out, "fullCommitPeriod", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullCommitPeriod is a free data retrieval call binding the contract method 0xb3b7d503.
//
// Solidity: function fullCommitPeriod(address pool) view returns(uint256)
func (_Poolstatehelper *PoolstatehelperSession) FullCommitPeriod(pool common.Address) (*big.Int, error) {
	return _Poolstatehelper.Contract.FullCommitPeriod(&_Poolstatehelper.CallOpts, pool)
}

// FullCommitPeriod is a free data retrieval call binding the contract method 0xb3b7d503.
//
// Solidity: function fullCommitPeriod(address pool) view returns(uint256)
func (_Poolstatehelper *PoolstatehelperCallerSession) FullCommitPeriod(pool common.Address) (*big.Int, error) {
	return _Poolstatehelper.Contract.FullCommitPeriod(&_Poolstatehelper.CallOpts, pool)
}

// GetCommitQueue is a free data retrieval call binding the contract method 0xd69fa4c6.
//
// Solidity: function getCommitQueue(address committer, uint256 periods) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[] commitQueue)
func (_Poolstatehelper *PoolstatehelperCaller) GetCommitQueue(opts *bind.CallOpts, committer common.Address, periods *big.Int) ([]IPoolCommitterTotalCommitment, error) {
	var out []interface{}
	err := _Poolstatehelper.contract.Call(opts, &out, "getCommitQueue", committer, periods)

	if err != nil {
		return *new([]IPoolCommitterTotalCommitment), err
	}

	out0 := *abi.ConvertType(out[0], new([]IPoolCommitterTotalCommitment)).(*[]IPoolCommitterTotalCommitment)

	return out0, err

}

// GetCommitQueue is a free data retrieval call binding the contract method 0xd69fa4c6.
//
// Solidity: function getCommitQueue(address committer, uint256 periods) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[] commitQueue)
func (_Poolstatehelper *PoolstatehelperSession) GetCommitQueue(committer common.Address, periods *big.Int) ([]IPoolCommitterTotalCommitment, error) {
	return _Poolstatehelper.Contract.GetCommitQueue(&_Poolstatehelper.CallOpts, committer, periods)
}

// GetCommitQueue is a free data retrieval call binding the contract method 0xd69fa4c6.
//
// Solidity: function getCommitQueue(address committer, uint256 periods) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[] commitQueue)
func (_Poolstatehelper *PoolstatehelperCallerSession) GetCommitQueue(committer common.Address, periods *big.Int) ([]IPoolCommitterTotalCommitment, error) {
	return _Poolstatehelper.Contract.GetCommitQueue(&_Poolstatehelper.CallOpts, committer, periods)
}

// GetExpectedState is a free data retrieval call binding the contract method 0x0552cd18.
//
// Solidity: function getExpectedState(address pool, uint256 periods) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256) finalExpectedPoolState)
func (_Poolstatehelper *PoolstatehelperCaller) GetExpectedState(opts *bind.CallOpts, pool common.Address, periods *big.Int) (IPoolStateHelperExpectedPoolState, error) {
	var out []interface{}
	err := _Poolstatehelper.contract.Call(opts, &out, "getExpectedState", pool, periods)

	if err != nil {
		return *new(IPoolStateHelperExpectedPoolState), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolStateHelperExpectedPoolState)).(*IPoolStateHelperExpectedPoolState)

	return out0, err

}

// GetExpectedState is a free data retrieval call binding the contract method 0x0552cd18.
//
// Solidity: function getExpectedState(address pool, uint256 periods) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256) finalExpectedPoolState)
func (_Poolstatehelper *PoolstatehelperSession) GetExpectedState(pool common.Address, periods *big.Int) (IPoolStateHelperExpectedPoolState, error) {
	return _Poolstatehelper.Contract.GetExpectedState(&_Poolstatehelper.CallOpts, pool, periods)
}

// GetExpectedState is a free data retrieval call binding the contract method 0x0552cd18.
//
// Solidity: function getExpectedState(address pool, uint256 periods) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,int256) finalExpectedPoolState)
func (_Poolstatehelper *PoolstatehelperCallerSession) GetExpectedState(pool common.Address, periods *big.Int) (IPoolStateHelperExpectedPoolState, error) {
	return _Poolstatehelper.Contract.GetExpectedState(&_Poolstatehelper.CallOpts, pool, periods)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x486e51f4.
//
// Solidity: function getPoolInfo(address pool, address committer) view returns(((uint256,uint256,uint256),(uint256,uint256,uint256)) poolInfo)
func (_Poolstatehelper *PoolstatehelperCaller) GetPoolInfo(opts *bind.CallOpts, pool common.Address, committer common.Address) (IPoolStateHelperPoolInfo, error) {
	var out []interface{}
	err := _Poolstatehelper.contract.Call(opts, &out, "getPoolInfo", pool, committer)

	if err != nil {
		return *new(IPoolStateHelperPoolInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolStateHelperPoolInfo)).(*IPoolStateHelperPoolInfo)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x486e51f4.
//
// Solidity: function getPoolInfo(address pool, address committer) view returns(((uint256,uint256,uint256),(uint256,uint256,uint256)) poolInfo)
func (_Poolstatehelper *PoolstatehelperSession) GetPoolInfo(pool common.Address, committer common.Address) (IPoolStateHelperPoolInfo, error) {
	return _Poolstatehelper.Contract.GetPoolInfo(&_Poolstatehelper.CallOpts, pool, committer)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x486e51f4.
//
// Solidity: function getPoolInfo(address pool, address committer) view returns(((uint256,uint256,uint256),(uint256,uint256,uint256)) poolInfo)
func (_Poolstatehelper *PoolstatehelperCallerSession) GetPoolInfo(pool common.Address, committer common.Address) (IPoolStateHelperPoolInfo, error) {
	return _Poolstatehelper.Contract.GetPoolInfo(&_Poolstatehelper.CallOpts, pool, committer)
}
