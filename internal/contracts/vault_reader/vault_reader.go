// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultreader

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

// VaultreaderMetaData contains all meta data concerning the Vaultreader contract.
var VaultreaderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_usdgAmount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"getVaultTokenInfoV3\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VaultreaderABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultreaderMetaData.ABI instead.
var VaultreaderABI = VaultreaderMetaData.ABI

// Vaultreader is an auto generated Go binding around an Ethereum contract.
type Vaultreader struct {
	VaultreaderCaller     // Read-only binding to the contract
	VaultreaderTransactor // Write-only binding to the contract
	VaultreaderFilterer   // Log filterer for contract events
}

// VaultreaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultreaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultreaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultreaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultreaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultreaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultreaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultreaderSession struct {
	Contract     *Vaultreader      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultreaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultreaderCallerSession struct {
	Contract *VaultreaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultreaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultreaderTransactorSession struct {
	Contract     *VaultreaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultreaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultreaderRaw struct {
	Contract *Vaultreader // Generic contract binding to access the raw methods on
}

// VaultreaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultreaderCallerRaw struct {
	Contract *VaultreaderCaller // Generic read-only contract binding to access the raw methods on
}

// VaultreaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultreaderTransactorRaw struct {
	Contract *VaultreaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultreader creates a new instance of Vaultreader, bound to a specific deployed contract.
func NewVaultreader(address common.Address, backend bind.ContractBackend) (*Vaultreader, error) {
	contract, err := bindVaultreader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vaultreader{VaultreaderCaller: VaultreaderCaller{contract: contract}, VaultreaderTransactor: VaultreaderTransactor{contract: contract}, VaultreaderFilterer: VaultreaderFilterer{contract: contract}}, nil
}

// NewVaultreaderCaller creates a new read-only instance of Vaultreader, bound to a specific deployed contract.
func NewVaultreaderCaller(address common.Address, caller bind.ContractCaller) (*VaultreaderCaller, error) {
	contract, err := bindVaultreader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultreaderCaller{contract: contract}, nil
}

// NewVaultreaderTransactor creates a new write-only instance of Vaultreader, bound to a specific deployed contract.
func NewVaultreaderTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultreaderTransactor, error) {
	contract, err := bindVaultreader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultreaderTransactor{contract: contract}, nil
}

// NewVaultreaderFilterer creates a new log filterer instance of Vaultreader, bound to a specific deployed contract.
func NewVaultreaderFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultreaderFilterer, error) {
	contract, err := bindVaultreader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultreaderFilterer{contract: contract}, nil
}

// bindVaultreader binds a generic wrapper to an already deployed contract.
func bindVaultreader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultreaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vaultreader *VaultreaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vaultreader.Contract.VaultreaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vaultreader *VaultreaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultreader.Contract.VaultreaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vaultreader *VaultreaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vaultreader.Contract.VaultreaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vaultreader *VaultreaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vaultreader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vaultreader *VaultreaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultreader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vaultreader *VaultreaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vaultreader.Contract.contract.Transact(opts, method, params...)
}

// GetVaultTokenInfoV3 is a free data retrieval call binding the contract method 0xf75e8101.
//
// Solidity: function getVaultTokenInfoV3(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Vaultreader *VaultreaderCaller) GetVaultTokenInfoV3(opts *bind.CallOpts, _vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Vaultreader.contract.Call(opts, &out, "getVaultTokenInfoV3", _vault, _positionManager, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfoV3 is a free data retrieval call binding the contract method 0xf75e8101.
//
// Solidity: function getVaultTokenInfoV3(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Vaultreader *VaultreaderSession) GetVaultTokenInfoV3(_vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Vaultreader.Contract.GetVaultTokenInfoV3(&_Vaultreader.CallOpts, _vault, _positionManager, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV3 is a free data retrieval call binding the contract method 0xf75e8101.
//
// Solidity: function getVaultTokenInfoV3(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Vaultreader *VaultreaderCallerSession) GetVaultTokenInfoV3(_vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Vaultreader.Contract.GetVaultTokenInfoV3(&_Vaultreader.CallOpts, _vault, _positionManager, _weth, _usdgAmount, _tokens)
}
