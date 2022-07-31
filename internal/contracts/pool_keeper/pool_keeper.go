// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolkeeper

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

// PoolkeeperMetaData contains all meta data concerning the Poolkeeper contract.
var PoolkeeperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"name\":\"FactoryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"GasPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"KeeperPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedReward\",\"type\":\"uint256\"}],\"name\":\"KeeperPaymentError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldKeeperRewards\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newKeeperRewards\",\"type\":\"address\"}],\"name\":\"KeeperRewardsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"firstPrice\",\"type\":\"int256\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"PoolUpkeepError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"startPrice\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"endPrice\",\"type\":\"int256\"}],\"name\":\"UpkeepSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_pools\",\"type\":\"address[]\"}],\"name\":\"checkUpkeepMultiplePools\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executionPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIPoolFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"isUpkeepRequiredSinglePool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeperRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"newPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"pools\",\"type\":\"address[]\"}],\"name\":\"performUpkeepMultiplePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pools\",\"type\":\"bytes\"}],\"name\":\"performUpkeepMultiplePoolsPacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"performUpkeepSinglePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeperRewards\",\"type\":\"address\"}],\"name\":\"setKeeperRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolkeeperABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolkeeperMetaData.ABI instead.
var PoolkeeperABI = PoolkeeperMetaData.ABI

// Poolkeeper is an auto generated Go binding around an Ethereum contract.
type Poolkeeper struct {
	PoolkeeperCaller     // Read-only binding to the contract
	PoolkeeperTransactor // Write-only binding to the contract
	PoolkeeperFilterer   // Log filterer for contract events
}

// PoolkeeperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolkeeperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolkeeperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolkeeperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolkeeperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolkeeperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolkeeperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolkeeperSession struct {
	Contract     *Poolkeeper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolkeeperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolkeeperCallerSession struct {
	Contract *PoolkeeperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PoolkeeperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolkeeperTransactorSession struct {
	Contract     *PoolkeeperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PoolkeeperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolkeeperRaw struct {
	Contract *Poolkeeper // Generic contract binding to access the raw methods on
}

// PoolkeeperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolkeeperCallerRaw struct {
	Contract *PoolkeeperCaller // Generic read-only contract binding to access the raw methods on
}

// PoolkeeperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolkeeperTransactorRaw struct {
	Contract *PoolkeeperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolkeeper creates a new instance of Poolkeeper, bound to a specific deployed contract.
func NewPoolkeeper(address common.Address, backend bind.ContractBackend) (*Poolkeeper, error) {
	contract, err := bindPoolkeeper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolkeeper{PoolkeeperCaller: PoolkeeperCaller{contract: contract}, PoolkeeperTransactor: PoolkeeperTransactor{contract: contract}, PoolkeeperFilterer: PoolkeeperFilterer{contract: contract}}, nil
}

// NewPoolkeeperCaller creates a new read-only instance of Poolkeeper, bound to a specific deployed contract.
func NewPoolkeeperCaller(address common.Address, caller bind.ContractCaller) (*PoolkeeperCaller, error) {
	contract, err := bindPoolkeeper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperCaller{contract: contract}, nil
}

// NewPoolkeeperTransactor creates a new write-only instance of Poolkeeper, bound to a specific deployed contract.
func NewPoolkeeperTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolkeeperTransactor, error) {
	contract, err := bindPoolkeeper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperTransactor{contract: contract}, nil
}

// NewPoolkeeperFilterer creates a new log filterer instance of Poolkeeper, bound to a specific deployed contract.
func NewPoolkeeperFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolkeeperFilterer, error) {
	contract, err := bindPoolkeeper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperFilterer{contract: contract}, nil
}

// bindPoolkeeper binds a generic wrapper to an already deployed contract.
func bindPoolkeeper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolkeeperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolkeeper *PoolkeeperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolkeeper.Contract.PoolkeeperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolkeeper *PoolkeeperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PoolkeeperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolkeeper *PoolkeeperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PoolkeeperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolkeeper *PoolkeeperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolkeeper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolkeeper *PoolkeeperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolkeeper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolkeeper *PoolkeeperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolkeeper.Contract.contract.Transact(opts, method, params...)
}

// CheckUpkeepMultiplePools is a free data retrieval call binding the contract method 0xce1dcdfc.
//
// Solidity: function checkUpkeepMultiplePools(address[] _pools) view returns(bool)
func (_Poolkeeper *PoolkeeperCaller) CheckUpkeepMultiplePools(opts *bind.CallOpts, _pools []common.Address) (bool, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "checkUpkeepMultiplePools", _pools)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckUpkeepMultiplePools is a free data retrieval call binding the contract method 0xce1dcdfc.
//
// Solidity: function checkUpkeepMultiplePools(address[] _pools) view returns(bool)
func (_Poolkeeper *PoolkeeperSession) CheckUpkeepMultiplePools(_pools []common.Address) (bool, error) {
	return _Poolkeeper.Contract.CheckUpkeepMultiplePools(&_Poolkeeper.CallOpts, _pools)
}

// CheckUpkeepMultiplePools is a free data retrieval call binding the contract method 0xce1dcdfc.
//
// Solidity: function checkUpkeepMultiplePools(address[] _pools) view returns(bool)
func (_Poolkeeper *PoolkeeperCallerSession) CheckUpkeepMultiplePools(_pools []common.Address) (bool, error) {
	return _Poolkeeper.Contract.CheckUpkeepMultiplePools(&_Poolkeeper.CallOpts, _pools)
}

// ExecutionPrice is a free data retrieval call binding the contract method 0xcdd9e137.
//
// Solidity: function executionPrice(address ) view returns(int256)
func (_Poolkeeper *PoolkeeperCaller) ExecutionPrice(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "executionPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionPrice is a free data retrieval call binding the contract method 0xcdd9e137.
//
// Solidity: function executionPrice(address ) view returns(int256)
func (_Poolkeeper *PoolkeeperSession) ExecutionPrice(arg0 common.Address) (*big.Int, error) {
	return _Poolkeeper.Contract.ExecutionPrice(&_Poolkeeper.CallOpts, arg0)
}

// ExecutionPrice is a free data retrieval call binding the contract method 0xcdd9e137.
//
// Solidity: function executionPrice(address ) view returns(int256)
func (_Poolkeeper *PoolkeeperCallerSession) ExecutionPrice(arg0 common.Address) (*big.Int, error) {
	return _Poolkeeper.Contract.ExecutionPrice(&_Poolkeeper.CallOpts, arg0)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Poolkeeper *PoolkeeperCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Poolkeeper *PoolkeeperSession) Factory() (common.Address, error) {
	return _Poolkeeper.Contract.Factory(&_Poolkeeper.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Poolkeeper *PoolkeeperCallerSession) Factory() (common.Address, error) {
	return _Poolkeeper.Contract.Factory(&_Poolkeeper.CallOpts)
}

// GasPrice is a free data retrieval call binding the contract method 0xfe173b97.
//
// Solidity: function gasPrice() view returns(uint256)
func (_Poolkeeper *PoolkeeperCaller) GasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "gasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPrice is a free data retrieval call binding the contract method 0xfe173b97.
//
// Solidity: function gasPrice() view returns(uint256)
func (_Poolkeeper *PoolkeeperSession) GasPrice() (*big.Int, error) {
	return _Poolkeeper.Contract.GasPrice(&_Poolkeeper.CallOpts)
}

// GasPrice is a free data retrieval call binding the contract method 0xfe173b97.
//
// Solidity: function gasPrice() view returns(uint256)
func (_Poolkeeper *PoolkeeperCallerSession) GasPrice() (*big.Int, error) {
	return _Poolkeeper.Contract.GasPrice(&_Poolkeeper.CallOpts)
}

// IsUpkeepRequiredSinglePool is a free data retrieval call binding the contract method 0x33261111.
//
// Solidity: function isUpkeepRequiredSinglePool(address _pool) view returns(bool)
func (_Poolkeeper *PoolkeeperCaller) IsUpkeepRequiredSinglePool(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "isUpkeepRequiredSinglePool", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpkeepRequiredSinglePool is a free data retrieval call binding the contract method 0x33261111.
//
// Solidity: function isUpkeepRequiredSinglePool(address _pool) view returns(bool)
func (_Poolkeeper *PoolkeeperSession) IsUpkeepRequiredSinglePool(_pool common.Address) (bool, error) {
	return _Poolkeeper.Contract.IsUpkeepRequiredSinglePool(&_Poolkeeper.CallOpts, _pool)
}

// IsUpkeepRequiredSinglePool is a free data retrieval call binding the contract method 0x33261111.
//
// Solidity: function isUpkeepRequiredSinglePool(address _pool) view returns(bool)
func (_Poolkeeper *PoolkeeperCallerSession) IsUpkeepRequiredSinglePool(_pool common.Address) (bool, error) {
	return _Poolkeeper.Contract.IsUpkeepRequiredSinglePool(&_Poolkeeper.CallOpts, _pool)
}

// KeeperRewards is a free data retrieval call binding the contract method 0x857ad312.
//
// Solidity: function keeperRewards() view returns(address)
func (_Poolkeeper *PoolkeeperCaller) KeeperRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "keeperRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeeperRewards is a free data retrieval call binding the contract method 0x857ad312.
//
// Solidity: function keeperRewards() view returns(address)
func (_Poolkeeper *PoolkeeperSession) KeeperRewards() (common.Address, error) {
	return _Poolkeeper.Contract.KeeperRewards(&_Poolkeeper.CallOpts)
}

// KeeperRewards is a free data retrieval call binding the contract method 0x857ad312.
//
// Solidity: function keeperRewards() view returns(address)
func (_Poolkeeper *PoolkeeperCallerSession) KeeperRewards() (common.Address, error) {
	return _Poolkeeper.Contract.KeeperRewards(&_Poolkeeper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Poolkeeper *PoolkeeperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolkeeper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Poolkeeper *PoolkeeperSession) Owner() (common.Address, error) {
	return _Poolkeeper.Contract.Owner(&_Poolkeeper.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Poolkeeper *PoolkeeperCallerSession) Owner() (common.Address, error) {
	return _Poolkeeper.Contract.Owner(&_Poolkeeper.CallOpts)
}

// NewPool is a paid mutator transaction binding the contract method 0x3c5c7fb3.
//
// Solidity: function newPool(address _poolAddress) returns()
func (_Poolkeeper *PoolkeeperTransactor) NewPool(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "newPool", _poolAddress)
}

// NewPool is a paid mutator transaction binding the contract method 0x3c5c7fb3.
//
// Solidity: function newPool(address _poolAddress) returns()
func (_Poolkeeper *PoolkeeperSession) NewPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.NewPool(&_Poolkeeper.TransactOpts, _poolAddress)
}

// NewPool is a paid mutator transaction binding the contract method 0x3c5c7fb3.
//
// Solidity: function newPool(address _poolAddress) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) NewPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.NewPool(&_Poolkeeper.TransactOpts, _poolAddress)
}

// PerformUpkeepMultiplePools is a paid mutator transaction binding the contract method 0xfcab9e98.
//
// Solidity: function performUpkeepMultiplePools(address[] pools) returns()
func (_Poolkeeper *PoolkeeperTransactor) PerformUpkeepMultiplePools(opts *bind.TransactOpts, pools []common.Address) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "performUpkeepMultiplePools", pools)
}

// PerformUpkeepMultiplePools is a paid mutator transaction binding the contract method 0xfcab9e98.
//
// Solidity: function performUpkeepMultiplePools(address[] pools) returns()
func (_Poolkeeper *PoolkeeperSession) PerformUpkeepMultiplePools(pools []common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PerformUpkeepMultiplePools(&_Poolkeeper.TransactOpts, pools)
}

// PerformUpkeepMultiplePools is a paid mutator transaction binding the contract method 0xfcab9e98.
//
// Solidity: function performUpkeepMultiplePools(address[] pools) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) PerformUpkeepMultiplePools(pools []common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PerformUpkeepMultiplePools(&_Poolkeeper.TransactOpts, pools)
}

// PerformUpkeepMultiplePoolsPacked is a paid mutator transaction binding the contract method 0x2027574e.
//
// Solidity: function performUpkeepMultiplePoolsPacked(bytes pools) returns()
func (_Poolkeeper *PoolkeeperTransactor) PerformUpkeepMultiplePoolsPacked(opts *bind.TransactOpts, pools []byte) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "performUpkeepMultiplePoolsPacked", pools)
}

// PerformUpkeepMultiplePoolsPacked is a paid mutator transaction binding the contract method 0x2027574e.
//
// Solidity: function performUpkeepMultiplePoolsPacked(bytes pools) returns()
func (_Poolkeeper *PoolkeeperSession) PerformUpkeepMultiplePoolsPacked(pools []byte) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PerformUpkeepMultiplePoolsPacked(&_Poolkeeper.TransactOpts, pools)
}

// PerformUpkeepMultiplePoolsPacked is a paid mutator transaction binding the contract method 0x2027574e.
//
// Solidity: function performUpkeepMultiplePoolsPacked(bytes pools) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) PerformUpkeepMultiplePoolsPacked(pools []byte) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PerformUpkeepMultiplePoolsPacked(&_Poolkeeper.TransactOpts, pools)
}

// PerformUpkeepSinglePool is a paid mutator transaction binding the contract method 0x91a60782.
//
// Solidity: function performUpkeepSinglePool(address _pool) returns()
func (_Poolkeeper *PoolkeeperTransactor) PerformUpkeepSinglePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "performUpkeepSinglePool", _pool)
}

// PerformUpkeepSinglePool is a paid mutator transaction binding the contract method 0x91a60782.
//
// Solidity: function performUpkeepSinglePool(address _pool) returns()
func (_Poolkeeper *PoolkeeperSession) PerformUpkeepSinglePool(_pool common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PerformUpkeepSinglePool(&_Poolkeeper.TransactOpts, _pool)
}

// PerformUpkeepSinglePool is a paid mutator transaction binding the contract method 0x91a60782.
//
// Solidity: function performUpkeepSinglePool(address _pool) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) PerformUpkeepSinglePool(_pool common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.PerformUpkeepSinglePool(&_Poolkeeper.TransactOpts, _pool)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Poolkeeper *PoolkeeperTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Poolkeeper *PoolkeeperSession) RenounceOwnership() (*types.Transaction, error) {
	return _Poolkeeper.Contract.RenounceOwnership(&_Poolkeeper.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Poolkeeper *PoolkeeperTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Poolkeeper.Contract.RenounceOwnership(&_Poolkeeper.TransactOpts)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xbf1fe420.
//
// Solidity: function setGasPrice(uint256 _price) returns()
func (_Poolkeeper *PoolkeeperTransactor) SetGasPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "setGasPrice", _price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xbf1fe420.
//
// Solidity: function setGasPrice(uint256 _price) returns()
func (_Poolkeeper *PoolkeeperSession) SetGasPrice(_price *big.Int) (*types.Transaction, error) {
	return _Poolkeeper.Contract.SetGasPrice(&_Poolkeeper.TransactOpts, _price)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0xbf1fe420.
//
// Solidity: function setGasPrice(uint256 _price) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) SetGasPrice(_price *big.Int) (*types.Transaction, error) {
	return _Poolkeeper.Contract.SetGasPrice(&_Poolkeeper.TransactOpts, _price)
}

// SetKeeperRewards is a paid mutator transaction binding the contract method 0xe98ccdc4.
//
// Solidity: function setKeeperRewards(address _keeperRewards) returns()
func (_Poolkeeper *PoolkeeperTransactor) SetKeeperRewards(opts *bind.TransactOpts, _keeperRewards common.Address) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "setKeeperRewards", _keeperRewards)
}

// SetKeeperRewards is a paid mutator transaction binding the contract method 0xe98ccdc4.
//
// Solidity: function setKeeperRewards(address _keeperRewards) returns()
func (_Poolkeeper *PoolkeeperSession) SetKeeperRewards(_keeperRewards common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.SetKeeperRewards(&_Poolkeeper.TransactOpts, _keeperRewards)
}

// SetKeeperRewards is a paid mutator transaction binding the contract method 0xe98ccdc4.
//
// Solidity: function setKeeperRewards(address _keeperRewards) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) SetKeeperRewards(_keeperRewards common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.SetKeeperRewards(&_Poolkeeper.TransactOpts, _keeperRewards)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Poolkeeper *PoolkeeperTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Poolkeeper.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Poolkeeper *PoolkeeperSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.TransferOwnership(&_Poolkeeper.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Poolkeeper *PoolkeeperTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Poolkeeper.Contract.TransferOwnership(&_Poolkeeper.TransactOpts, newOwner)
}

// PoolkeeperFactoryChangedIterator is returned from FilterFactoryChanged and is used to iterate over the raw logs and unpacked data for FactoryChanged events raised by the Poolkeeper contract.
type PoolkeeperFactoryChangedIterator struct {
	Event *PoolkeeperFactoryChanged // Event containing the contract specifics and raw log

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
func (it *PoolkeeperFactoryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperFactoryChanged)
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
		it.Event = new(PoolkeeperFactoryChanged)
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
func (it *PoolkeeperFactoryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperFactoryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperFactoryChanged represents a FactoryChanged event raised by the Poolkeeper contract.
type PoolkeeperFactoryChanged struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFactoryChanged is a free log retrieval operation binding the contract event 0x7ead54cd1fd7e5ddea18605dec62ecb6a27d1f68bee1054cbf93b7abb4e1dc6b.
//
// Solidity: event FactoryChanged(address indexed factory)
func (_Poolkeeper *PoolkeeperFilterer) FilterFactoryChanged(opts *bind.FilterOpts, factory []common.Address) (*PoolkeeperFactoryChangedIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "FactoryChanged", factoryRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperFactoryChangedIterator{contract: _Poolkeeper.contract, event: "FactoryChanged", logs: logs, sub: sub}, nil
}

// WatchFactoryChanged is a free log subscription operation binding the contract event 0x7ead54cd1fd7e5ddea18605dec62ecb6a27d1f68bee1054cbf93b7abb4e1dc6b.
//
// Solidity: event FactoryChanged(address indexed factory)
func (_Poolkeeper *PoolkeeperFilterer) WatchFactoryChanged(opts *bind.WatchOpts, sink chan<- *PoolkeeperFactoryChanged, factory []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "FactoryChanged", factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperFactoryChanged)
				if err := _Poolkeeper.contract.UnpackLog(event, "FactoryChanged", log); err != nil {
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

// ParseFactoryChanged is a log parse operation binding the contract event 0x7ead54cd1fd7e5ddea18605dec62ecb6a27d1f68bee1054cbf93b7abb4e1dc6b.
//
// Solidity: event FactoryChanged(address indexed factory)
func (_Poolkeeper *PoolkeeperFilterer) ParseFactoryChanged(log types.Log) (*PoolkeeperFactoryChanged, error) {
	event := new(PoolkeeperFactoryChanged)
	if err := _Poolkeeper.contract.UnpackLog(event, "FactoryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperGasPriceChangedIterator is returned from FilterGasPriceChanged and is used to iterate over the raw logs and unpacked data for GasPriceChanged events raised by the Poolkeeper contract.
type PoolkeeperGasPriceChangedIterator struct {
	Event *PoolkeeperGasPriceChanged // Event containing the contract specifics and raw log

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
func (it *PoolkeeperGasPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperGasPriceChanged)
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
		it.Event = new(PoolkeeperGasPriceChanged)
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
func (it *PoolkeeperGasPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperGasPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperGasPriceChanged represents a GasPriceChanged event raised by the Poolkeeper contract.
type PoolkeeperGasPriceChanged struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGasPriceChanged is a free log retrieval operation binding the contract event 0x52264b89e0fceafb26e79fd49ef8a366eb6297483bf4035b027f0c99a7ad512e.
//
// Solidity: event GasPriceChanged(uint256 indexed price)
func (_Poolkeeper *PoolkeeperFilterer) FilterGasPriceChanged(opts *bind.FilterOpts, price []*big.Int) (*PoolkeeperGasPriceChangedIterator, error) {

	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "GasPriceChanged", priceRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperGasPriceChangedIterator{contract: _Poolkeeper.contract, event: "GasPriceChanged", logs: logs, sub: sub}, nil
}

// WatchGasPriceChanged is a free log subscription operation binding the contract event 0x52264b89e0fceafb26e79fd49ef8a366eb6297483bf4035b027f0c99a7ad512e.
//
// Solidity: event GasPriceChanged(uint256 indexed price)
func (_Poolkeeper *PoolkeeperFilterer) WatchGasPriceChanged(opts *bind.WatchOpts, sink chan<- *PoolkeeperGasPriceChanged, price []*big.Int) (event.Subscription, error) {

	var priceRule []interface{}
	for _, priceItem := range price {
		priceRule = append(priceRule, priceItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "GasPriceChanged", priceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperGasPriceChanged)
				if err := _Poolkeeper.contract.UnpackLog(event, "GasPriceChanged", log); err != nil {
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

// ParseGasPriceChanged is a log parse operation binding the contract event 0x52264b89e0fceafb26e79fd49ef8a366eb6297483bf4035b027f0c99a7ad512e.
//
// Solidity: event GasPriceChanged(uint256 indexed price)
func (_Poolkeeper *PoolkeeperFilterer) ParseGasPriceChanged(log types.Log) (*PoolkeeperGasPriceChanged, error) {
	event := new(PoolkeeperGasPriceChanged)
	if err := _Poolkeeper.contract.UnpackLog(event, "GasPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperKeeperPaidIterator is returned from FilterKeeperPaid and is used to iterate over the raw logs and unpacked data for KeeperPaid events raised by the Poolkeeper contract.
type PoolkeeperKeeperPaidIterator struct {
	Event *PoolkeeperKeeperPaid // Event containing the contract specifics and raw log

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
func (it *PoolkeeperKeeperPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperKeeperPaid)
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
		it.Event = new(PoolkeeperKeeperPaid)
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
func (it *PoolkeeperKeeperPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperKeeperPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperKeeperPaid represents a KeeperPaid event raised by the Poolkeeper contract.
type PoolkeeperKeeperPaid struct {
	Pool   common.Address
	Keeper common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeeperPaid is a free log retrieval operation binding the contract event 0xf03012dd4aab142682aa01439f94e975e2ec77ef558352323bc05f382f74b082.
//
// Solidity: event KeeperPaid(address indexed _pool, address indexed keeper, uint256 reward)
func (_Poolkeeper *PoolkeeperFilterer) FilterKeeperPaid(opts *bind.FilterOpts, _pool []common.Address, keeper []common.Address) (*PoolkeeperKeeperPaidIterator, error) {

	var _poolRule []interface{}
	for _, _poolItem := range _pool {
		_poolRule = append(_poolRule, _poolItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "KeeperPaid", _poolRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperKeeperPaidIterator{contract: _Poolkeeper.contract, event: "KeeperPaid", logs: logs, sub: sub}, nil
}

// WatchKeeperPaid is a free log subscription operation binding the contract event 0xf03012dd4aab142682aa01439f94e975e2ec77ef558352323bc05f382f74b082.
//
// Solidity: event KeeperPaid(address indexed _pool, address indexed keeper, uint256 reward)
func (_Poolkeeper *PoolkeeperFilterer) WatchKeeperPaid(opts *bind.WatchOpts, sink chan<- *PoolkeeperKeeperPaid, _pool []common.Address, keeper []common.Address) (event.Subscription, error) {

	var _poolRule []interface{}
	for _, _poolItem := range _pool {
		_poolRule = append(_poolRule, _poolItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "KeeperPaid", _poolRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperKeeperPaid)
				if err := _Poolkeeper.contract.UnpackLog(event, "KeeperPaid", log); err != nil {
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

// ParseKeeperPaid is a log parse operation binding the contract event 0xf03012dd4aab142682aa01439f94e975e2ec77ef558352323bc05f382f74b082.
//
// Solidity: event KeeperPaid(address indexed _pool, address indexed keeper, uint256 reward)
func (_Poolkeeper *PoolkeeperFilterer) ParseKeeperPaid(log types.Log) (*PoolkeeperKeeperPaid, error) {
	event := new(PoolkeeperKeeperPaid)
	if err := _Poolkeeper.contract.UnpackLog(event, "KeeperPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperKeeperPaymentErrorIterator is returned from FilterKeeperPaymentError and is used to iterate over the raw logs and unpacked data for KeeperPaymentError events raised by the Poolkeeper contract.
type PoolkeeperKeeperPaymentErrorIterator struct {
	Event *PoolkeeperKeeperPaymentError // Event containing the contract specifics and raw log

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
func (it *PoolkeeperKeeperPaymentErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperKeeperPaymentError)
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
		it.Event = new(PoolkeeperKeeperPaymentError)
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
func (it *PoolkeeperKeeperPaymentErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperKeeperPaymentErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperKeeperPaymentError represents a KeeperPaymentError event raised by the Poolkeeper contract.
type PoolkeeperKeeperPaymentError struct {
	Pool           common.Address
	Keeper         common.Address
	ExpectedReward *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeeperPaymentError is a free log retrieval operation binding the contract event 0x08b35d4c86707ba44cdcc5efaa79b8a307e6c7f789f1eb86ca1ff1e9bde05697.
//
// Solidity: event KeeperPaymentError(address indexed _pool, address indexed keeper, uint256 expectedReward)
func (_Poolkeeper *PoolkeeperFilterer) FilterKeeperPaymentError(opts *bind.FilterOpts, _pool []common.Address, keeper []common.Address) (*PoolkeeperKeeperPaymentErrorIterator, error) {

	var _poolRule []interface{}
	for _, _poolItem := range _pool {
		_poolRule = append(_poolRule, _poolItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "KeeperPaymentError", _poolRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperKeeperPaymentErrorIterator{contract: _Poolkeeper.contract, event: "KeeperPaymentError", logs: logs, sub: sub}, nil
}

// WatchKeeperPaymentError is a free log subscription operation binding the contract event 0x08b35d4c86707ba44cdcc5efaa79b8a307e6c7f789f1eb86ca1ff1e9bde05697.
//
// Solidity: event KeeperPaymentError(address indexed _pool, address indexed keeper, uint256 expectedReward)
func (_Poolkeeper *PoolkeeperFilterer) WatchKeeperPaymentError(opts *bind.WatchOpts, sink chan<- *PoolkeeperKeeperPaymentError, _pool []common.Address, keeper []common.Address) (event.Subscription, error) {

	var _poolRule []interface{}
	for _, _poolItem := range _pool {
		_poolRule = append(_poolRule, _poolItem)
	}
	var keeperRule []interface{}
	for _, keeperItem := range keeper {
		keeperRule = append(keeperRule, keeperItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "KeeperPaymentError", _poolRule, keeperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperKeeperPaymentError)
				if err := _Poolkeeper.contract.UnpackLog(event, "KeeperPaymentError", log); err != nil {
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

// ParseKeeperPaymentError is a log parse operation binding the contract event 0x08b35d4c86707ba44cdcc5efaa79b8a307e6c7f789f1eb86ca1ff1e9bde05697.
//
// Solidity: event KeeperPaymentError(address indexed _pool, address indexed keeper, uint256 expectedReward)
func (_Poolkeeper *PoolkeeperFilterer) ParseKeeperPaymentError(log types.Log) (*PoolkeeperKeeperPaymentError, error) {
	event := new(PoolkeeperKeeperPaymentError)
	if err := _Poolkeeper.contract.UnpackLog(event, "KeeperPaymentError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperKeeperRewardsSetIterator is returned from FilterKeeperRewardsSet and is used to iterate over the raw logs and unpacked data for KeeperRewardsSet events raised by the Poolkeeper contract.
type PoolkeeperKeeperRewardsSetIterator struct {
	Event *PoolkeeperKeeperRewardsSet // Event containing the contract specifics and raw log

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
func (it *PoolkeeperKeeperRewardsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperKeeperRewardsSet)
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
		it.Event = new(PoolkeeperKeeperRewardsSet)
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
func (it *PoolkeeperKeeperRewardsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperKeeperRewardsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperKeeperRewardsSet represents a KeeperRewardsSet event raised by the Poolkeeper contract.
type PoolkeeperKeeperRewardsSet struct {
	OldKeeperRewards common.Address
	NewKeeperRewards common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterKeeperRewardsSet is a free log retrieval operation binding the contract event 0xed60217a0f5e211a1105c7e6fdc9df5753b136c553ae2974a15e8be0cc02ac73.
//
// Solidity: event KeeperRewardsSet(address indexed oldKeeperRewards, address indexed newKeeperRewards)
func (_Poolkeeper *PoolkeeperFilterer) FilterKeeperRewardsSet(opts *bind.FilterOpts, oldKeeperRewards []common.Address, newKeeperRewards []common.Address) (*PoolkeeperKeeperRewardsSetIterator, error) {

	var oldKeeperRewardsRule []interface{}
	for _, oldKeeperRewardsItem := range oldKeeperRewards {
		oldKeeperRewardsRule = append(oldKeeperRewardsRule, oldKeeperRewardsItem)
	}
	var newKeeperRewardsRule []interface{}
	for _, newKeeperRewardsItem := range newKeeperRewards {
		newKeeperRewardsRule = append(newKeeperRewardsRule, newKeeperRewardsItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "KeeperRewardsSet", oldKeeperRewardsRule, newKeeperRewardsRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperKeeperRewardsSetIterator{contract: _Poolkeeper.contract, event: "KeeperRewardsSet", logs: logs, sub: sub}, nil
}

// WatchKeeperRewardsSet is a free log subscription operation binding the contract event 0xed60217a0f5e211a1105c7e6fdc9df5753b136c553ae2974a15e8be0cc02ac73.
//
// Solidity: event KeeperRewardsSet(address indexed oldKeeperRewards, address indexed newKeeperRewards)
func (_Poolkeeper *PoolkeeperFilterer) WatchKeeperRewardsSet(opts *bind.WatchOpts, sink chan<- *PoolkeeperKeeperRewardsSet, oldKeeperRewards []common.Address, newKeeperRewards []common.Address) (event.Subscription, error) {

	var oldKeeperRewardsRule []interface{}
	for _, oldKeeperRewardsItem := range oldKeeperRewards {
		oldKeeperRewardsRule = append(oldKeeperRewardsRule, oldKeeperRewardsItem)
	}
	var newKeeperRewardsRule []interface{}
	for _, newKeeperRewardsItem := range newKeeperRewards {
		newKeeperRewardsRule = append(newKeeperRewardsRule, newKeeperRewardsItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "KeeperRewardsSet", oldKeeperRewardsRule, newKeeperRewardsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperKeeperRewardsSet)
				if err := _Poolkeeper.contract.UnpackLog(event, "KeeperRewardsSet", log); err != nil {
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

// ParseKeeperRewardsSet is a log parse operation binding the contract event 0xed60217a0f5e211a1105c7e6fdc9df5753b136c553ae2974a15e8be0cc02ac73.
//
// Solidity: event KeeperRewardsSet(address indexed oldKeeperRewards, address indexed newKeeperRewards)
func (_Poolkeeper *PoolkeeperFilterer) ParseKeeperRewardsSet(log types.Log) (*PoolkeeperKeeperRewardsSet, error) {
	event := new(PoolkeeperKeeperRewardsSet)
	if err := _Poolkeeper.contract.UnpackLog(event, "KeeperRewardsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Poolkeeper contract.
type PoolkeeperOwnershipTransferredIterator struct {
	Event *PoolkeeperOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolkeeperOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperOwnershipTransferred)
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
		it.Event = new(PoolkeeperOwnershipTransferred)
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
func (it *PoolkeeperOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperOwnershipTransferred represents a OwnershipTransferred event raised by the Poolkeeper contract.
type PoolkeeperOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Poolkeeper *PoolkeeperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolkeeperOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperOwnershipTransferredIterator{contract: _Poolkeeper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Poolkeeper *PoolkeeperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolkeeperOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperOwnershipTransferred)
				if err := _Poolkeeper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Poolkeeper *PoolkeeperFilterer) ParseOwnershipTransferred(log types.Log) (*PoolkeeperOwnershipTransferred, error) {
	event := new(PoolkeeperOwnershipTransferred)
	if err := _Poolkeeper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the Poolkeeper contract.
type PoolkeeperPoolAddedIterator struct {
	Event *PoolkeeperPoolAdded // Event containing the contract specifics and raw log

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
func (it *PoolkeeperPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperPoolAdded)
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
		it.Event = new(PoolkeeperPoolAdded)
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
func (it *PoolkeeperPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperPoolAdded represents a PoolAdded event raised by the Poolkeeper contract.
type PoolkeeperPoolAdded struct {
	PoolAddress common.Address
	FirstPrice  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0xe950cb32f7a787627d713d7365623fef2d762ef8fcf813b865a04350852deb9b.
//
// Solidity: event PoolAdded(address indexed poolAddress, int256 indexed firstPrice)
func (_Poolkeeper *PoolkeeperFilterer) FilterPoolAdded(opts *bind.FilterOpts, poolAddress []common.Address, firstPrice []*big.Int) (*PoolkeeperPoolAddedIterator, error) {

	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var firstPriceRule []interface{}
	for _, firstPriceItem := range firstPrice {
		firstPriceRule = append(firstPriceRule, firstPriceItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "PoolAdded", poolAddressRule, firstPriceRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperPoolAddedIterator{contract: _Poolkeeper.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0xe950cb32f7a787627d713d7365623fef2d762ef8fcf813b865a04350852deb9b.
//
// Solidity: event PoolAdded(address indexed poolAddress, int256 indexed firstPrice)
func (_Poolkeeper *PoolkeeperFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *PoolkeeperPoolAdded, poolAddress []common.Address, firstPrice []*big.Int) (event.Subscription, error) {

	var poolAddressRule []interface{}
	for _, poolAddressItem := range poolAddress {
		poolAddressRule = append(poolAddressRule, poolAddressItem)
	}
	var firstPriceRule []interface{}
	for _, firstPriceItem := range firstPrice {
		firstPriceRule = append(firstPriceRule, firstPriceItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "PoolAdded", poolAddressRule, firstPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperPoolAdded)
				if err := _Poolkeeper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0xe950cb32f7a787627d713d7365623fef2d762ef8fcf813b865a04350852deb9b.
//
// Solidity: event PoolAdded(address indexed poolAddress, int256 indexed firstPrice)
func (_Poolkeeper *PoolkeeperFilterer) ParsePoolAdded(log types.Log) (*PoolkeeperPoolAdded, error) {
	event := new(PoolkeeperPoolAdded)
	if err := _Poolkeeper.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperPoolUpkeepErrorIterator is returned from FilterPoolUpkeepError and is used to iterate over the raw logs and unpacked data for PoolUpkeepError events raised by the Poolkeeper contract.
type PoolkeeperPoolUpkeepErrorIterator struct {
	Event *PoolkeeperPoolUpkeepError // Event containing the contract specifics and raw log

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
func (it *PoolkeeperPoolUpkeepErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperPoolUpkeepError)
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
		it.Event = new(PoolkeeperPoolUpkeepError)
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
func (it *PoolkeeperPoolUpkeepErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperPoolUpkeepErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperPoolUpkeepError represents a PoolUpkeepError event raised by the Poolkeeper contract.
type PoolkeeperPoolUpkeepError struct {
	Pool   common.Address
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolUpkeepError is a free log retrieval operation binding the contract event 0x8a47468bf96478429d382ed6385d9c77fba8b89540cd88199b1b51d5723b11fa.
//
// Solidity: event PoolUpkeepError(address indexed pool, string reason)
func (_Poolkeeper *PoolkeeperFilterer) FilterPoolUpkeepError(opts *bind.FilterOpts, pool []common.Address) (*PoolkeeperPoolUpkeepErrorIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "PoolUpkeepError", poolRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperPoolUpkeepErrorIterator{contract: _Poolkeeper.contract, event: "PoolUpkeepError", logs: logs, sub: sub}, nil
}

// WatchPoolUpkeepError is a free log subscription operation binding the contract event 0x8a47468bf96478429d382ed6385d9c77fba8b89540cd88199b1b51d5723b11fa.
//
// Solidity: event PoolUpkeepError(address indexed pool, string reason)
func (_Poolkeeper *PoolkeeperFilterer) WatchPoolUpkeepError(opts *bind.WatchOpts, sink chan<- *PoolkeeperPoolUpkeepError, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "PoolUpkeepError", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperPoolUpkeepError)
				if err := _Poolkeeper.contract.UnpackLog(event, "PoolUpkeepError", log); err != nil {
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

// ParsePoolUpkeepError is a log parse operation binding the contract event 0x8a47468bf96478429d382ed6385d9c77fba8b89540cd88199b1b51d5723b11fa.
//
// Solidity: event PoolUpkeepError(address indexed pool, string reason)
func (_Poolkeeper *PoolkeeperFilterer) ParsePoolUpkeepError(log types.Log) (*PoolkeeperPoolUpkeepError, error) {
	event := new(PoolkeeperPoolUpkeepError)
	if err := _Poolkeeper.contract.UnpackLog(event, "PoolUpkeepError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolkeeperUpkeepSuccessfulIterator is returned from FilterUpkeepSuccessful and is used to iterate over the raw logs and unpacked data for UpkeepSuccessful events raised by the Poolkeeper contract.
type PoolkeeperUpkeepSuccessfulIterator struct {
	Event *PoolkeeperUpkeepSuccessful // Event containing the contract specifics and raw log

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
func (it *PoolkeeperUpkeepSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolkeeperUpkeepSuccessful)
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
		it.Event = new(PoolkeeperUpkeepSuccessful)
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
func (it *PoolkeeperUpkeepSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolkeeperUpkeepSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolkeeperUpkeepSuccessful represents a UpkeepSuccessful event raised by the Poolkeeper contract.
type PoolkeeperUpkeepSuccessful struct {
	Pool       common.Address
	Data       []byte
	StartPrice *big.Int
	EndPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpkeepSuccessful is a free log retrieval operation binding the contract event 0x0aa8d2838ad6d4e2a5d7ca0826dfdbc36d922adf8fe5c6d25aee2724b07de074.
//
// Solidity: event UpkeepSuccessful(address indexed pool, bytes data, int256 indexed startPrice, int256 indexed endPrice)
func (_Poolkeeper *PoolkeeperFilterer) FilterUpkeepSuccessful(opts *bind.FilterOpts, pool []common.Address, startPrice []*big.Int, endPrice []*big.Int) (*PoolkeeperUpkeepSuccessfulIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	var startPriceRule []interface{}
	for _, startPriceItem := range startPrice {
		startPriceRule = append(startPriceRule, startPriceItem)
	}
	var endPriceRule []interface{}
	for _, endPriceItem := range endPrice {
		endPriceRule = append(endPriceRule, endPriceItem)
	}

	logs, sub, err := _Poolkeeper.contract.FilterLogs(opts, "UpkeepSuccessful", poolRule, startPriceRule, endPriceRule)
	if err != nil {
		return nil, err
	}
	return &PoolkeeperUpkeepSuccessfulIterator{contract: _Poolkeeper.contract, event: "UpkeepSuccessful", logs: logs, sub: sub}, nil
}

// WatchUpkeepSuccessful is a free log subscription operation binding the contract event 0x0aa8d2838ad6d4e2a5d7ca0826dfdbc36d922adf8fe5c6d25aee2724b07de074.
//
// Solidity: event UpkeepSuccessful(address indexed pool, bytes data, int256 indexed startPrice, int256 indexed endPrice)
func (_Poolkeeper *PoolkeeperFilterer) WatchUpkeepSuccessful(opts *bind.WatchOpts, sink chan<- *PoolkeeperUpkeepSuccessful, pool []common.Address, startPrice []*big.Int, endPrice []*big.Int) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	var startPriceRule []interface{}
	for _, startPriceItem := range startPrice {
		startPriceRule = append(startPriceRule, startPriceItem)
	}
	var endPriceRule []interface{}
	for _, endPriceItem := range endPrice {
		endPriceRule = append(endPriceRule, endPriceItem)
	}

	logs, sub, err := _Poolkeeper.contract.WatchLogs(opts, "UpkeepSuccessful", poolRule, startPriceRule, endPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolkeeperUpkeepSuccessful)
				if err := _Poolkeeper.contract.UnpackLog(event, "UpkeepSuccessful", log); err != nil {
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

// ParseUpkeepSuccessful is a log parse operation binding the contract event 0x0aa8d2838ad6d4e2a5d7ca0826dfdbc36d922adf8fe5c6d25aee2724b07de074.
//
// Solidity: event UpkeepSuccessful(address indexed pool, bytes data, int256 indexed startPrice, int256 indexed endPrice)
func (_Poolkeeper *PoolkeeperFilterer) ParseUpkeepSuccessful(log types.Log) (*PoolkeeperUpkeepSuccessful, error) {
	event := new(PoolkeeperUpkeepSuccessful)
	if err := _Poolkeeper.contract.UnpackLog(event, "UpkeepSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
