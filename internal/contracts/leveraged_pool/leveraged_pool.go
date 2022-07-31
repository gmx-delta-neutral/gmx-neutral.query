// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package leveragedpool

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

// ILeveragedPoolInitialization is an auto generated low-level Go binding around an user-defined struct.
type ILeveragedPoolInitialization struct {
	Owner                    common.Address
	Keeper                   common.Address
	OracleWrapper            common.Address
	SettlementEthOracle      common.Address
	LongToken                common.Address
	ShortToken               common.Address
	PoolCommitter            common.Address
	InvariantCheck           common.Address
	PoolName                 string
	FrontRunningInterval     uint32
	UpdateInterval           uint32
	LeverageAmount           uint16
	Fee                      *big.Int
	FeeAddress               common.Address
	SecondaryFeeAddress      common.Address
	SettlementToken          common.Address
	SecondaryFeeSplitPercent *big.Int
}

// LeveragedpoolMetaData contains all meta data concerning the Leveragedpool contract.
var LeveragedpoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"FeeAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"GovernanceAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"KeeperAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"}],\"name\":\"PoolBalancesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"longToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"shortToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settlementToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"poolName\",\"type\":\"string\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"shortBalanceChange\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"longBalanceChange\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shortFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"longFeeAmount\",\"type\":\"uint256\"}],\"name\":\"PoolRebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"startPrice\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"endPrice\",\"type\":\"int256\"}],\"name\":\"PriceChangeError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PrimaryFeesPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProvisionalGovernanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SecondaryFeeAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"secondaryFeeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SecondaryFeesPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"SettlementWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LONG_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHORT_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"burnTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrimaryFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimSecondaryFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"frontRunningInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLeverage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOraclePrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpkeepInformation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceTransferInProgress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_settlementEthOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_longToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_shortToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolCommitter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_invariantCheck\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_poolName\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"_frontRunningInterval\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_updateInterval\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_leverageAmount\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_secondaryFeeAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_settlementToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_secondaryFeeSplitPercent\",\"type\":\"uint256\"}],\"internalType\":\"structILeveragedPool.Initialization\",\"name\":\"initialization\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intervalPassed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantCheck\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPriceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leverageAmount\",\"outputs\":[{\"internalType\":\"bytes16\",\"name\":\"\",\"type\":\"bytes16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"longBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleWrapper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payKeeperFromBalances\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isLongToken\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"poolTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolTokens\",\"outputs\":[{\"internalType\":\"address[2]\",\"name\":\"\",\"type\":\"address[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_oldPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_newPrice\",\"type\":\"int256\"}],\"name\":\"poolUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"primaryFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"provisionalGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondaryFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondaryFeeSplitPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondaryFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_keeper\",\"type\":\"address\"}],\"name\":\"setKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_longBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_shortBalance\",\"type\":\"uint256\"}],\"name\":\"setNewPoolBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementEthOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settlementTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settlementTokenTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shortBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"updateFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"updateSecondaryFeeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LeveragedpoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LeveragedpoolMetaData.ABI instead.
var LeveragedpoolABI = LeveragedpoolMetaData.ABI

// Leveragedpool is an auto generated Go binding around an Ethereum contract.
type Leveragedpool struct {
	LeveragedpoolCaller     // Read-only binding to the contract
	LeveragedpoolTransactor // Write-only binding to the contract
	LeveragedpoolFilterer   // Log filterer for contract events
}

// LeveragedpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeveragedpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeveragedpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeveragedpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeveragedpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeveragedpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeveragedpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeveragedpoolSession struct {
	Contract     *Leveragedpool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeveragedpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeveragedpoolCallerSession struct {
	Contract *LeveragedpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LeveragedpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeveragedpoolTransactorSession struct {
	Contract     *LeveragedpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LeveragedpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeveragedpoolRaw struct {
	Contract *Leveragedpool // Generic contract binding to access the raw methods on
}

// LeveragedpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeveragedpoolCallerRaw struct {
	Contract *LeveragedpoolCaller // Generic read-only contract binding to access the raw methods on
}

// LeveragedpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeveragedpoolTransactorRaw struct {
	Contract *LeveragedpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeveragedpool creates a new instance of Leveragedpool, bound to a specific deployed contract.
func NewLeveragedpool(address common.Address, backend bind.ContractBackend) (*Leveragedpool, error) {
	contract, err := bindLeveragedpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Leveragedpool{LeveragedpoolCaller: LeveragedpoolCaller{contract: contract}, LeveragedpoolTransactor: LeveragedpoolTransactor{contract: contract}, LeveragedpoolFilterer: LeveragedpoolFilterer{contract: contract}}, nil
}

// NewLeveragedpoolCaller creates a new read-only instance of Leveragedpool, bound to a specific deployed contract.
func NewLeveragedpoolCaller(address common.Address, caller bind.ContractCaller) (*LeveragedpoolCaller, error) {
	contract, err := bindLeveragedpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolCaller{contract: contract}, nil
}

// NewLeveragedpoolTransactor creates a new write-only instance of Leveragedpool, bound to a specific deployed contract.
func NewLeveragedpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LeveragedpoolTransactor, error) {
	contract, err := bindLeveragedpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolTransactor{contract: contract}, nil
}

// NewLeveragedpoolFilterer creates a new log filterer instance of Leveragedpool, bound to a specific deployed contract.
func NewLeveragedpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LeveragedpoolFilterer, error) {
	contract, err := bindLeveragedpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolFilterer{contract: contract}, nil
}

// bindLeveragedpool binds a generic wrapper to an already deployed contract.
func bindLeveragedpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeveragedpoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leveragedpool *LeveragedpoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leveragedpool.Contract.LeveragedpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leveragedpool *LeveragedpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.Contract.LeveragedpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leveragedpool *LeveragedpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leveragedpool.Contract.LeveragedpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Leveragedpool *LeveragedpoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Leveragedpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Leveragedpool *LeveragedpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Leveragedpool *LeveragedpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Leveragedpool.Contract.contract.Transact(opts, method, params...)
}

// LONGINDEX is a free data retrieval call binding the contract method 0x062e0d69.
//
// Solidity: function LONG_INDEX() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) LONGINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "LONG_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LONGINDEX is a free data retrieval call binding the contract method 0x062e0d69.
//
// Solidity: function LONG_INDEX() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) LONGINDEX() (*big.Int, error) {
	return _Leveragedpool.Contract.LONGINDEX(&_Leveragedpool.CallOpts)
}

// LONGINDEX is a free data retrieval call binding the contract method 0x062e0d69.
//
// Solidity: function LONG_INDEX() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) LONGINDEX() (*big.Int, error) {
	return _Leveragedpool.Contract.LONGINDEX(&_Leveragedpool.CallOpts)
}

// SHORTINDEX is a free data retrieval call binding the contract method 0xc40f6aa9.
//
// Solidity: function SHORT_INDEX() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) SHORTINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "SHORT_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHORTINDEX is a free data retrieval call binding the contract method 0xc40f6aa9.
//
// Solidity: function SHORT_INDEX() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) SHORTINDEX() (*big.Int, error) {
	return _Leveragedpool.Contract.SHORTINDEX(&_Leveragedpool.CallOpts)
}

// SHORTINDEX is a free data retrieval call binding the contract method 0xc40f6aa9.
//
// Solidity: function SHORT_INDEX() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) SHORTINDEX() (*big.Int, error) {
	return _Leveragedpool.Contract.SHORTINDEX(&_Leveragedpool.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256, uint256)
func (_Leveragedpool *LeveragedpoolCaller) Balances(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "balances")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256, uint256)
func (_Leveragedpool *LeveragedpoolSession) Balances() (*big.Int, *big.Int, error) {
	return _Leveragedpool.Contract.Balances(&_Leveragedpool.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() view returns(uint256, uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) Balances() (*big.Int, *big.Int, error) {
	return _Leveragedpool.Contract.Balances(&_Leveragedpool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(bytes16)
func (_Leveragedpool *LeveragedpoolCaller) Fee(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(bytes16)
func (_Leveragedpool *LeveragedpoolSession) Fee() ([16]byte, error) {
	return _Leveragedpool.Contract.Fee(&_Leveragedpool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(bytes16)
func (_Leveragedpool *LeveragedpoolCallerSession) Fee() ([16]byte, error) {
	return _Leveragedpool.Contract.Fee(&_Leveragedpool.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) FeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "feeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) FeeAddress() (common.Address, error) {
	return _Leveragedpool.Contract.FeeAddress(&_Leveragedpool.CallOpts)
}

// FeeAddress is a free data retrieval call binding the contract method 0x41275358.
//
// Solidity: function feeAddress() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) FeeAddress() (common.Address, error) {
	return _Leveragedpool.Contract.FeeAddress(&_Leveragedpool.CallOpts)
}

// FrontRunningInterval is a free data retrieval call binding the contract method 0x7bfe789a.
//
// Solidity: function frontRunningInterval() view returns(uint32)
func (_Leveragedpool *LeveragedpoolCaller) FrontRunningInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "frontRunningInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FrontRunningInterval is a free data retrieval call binding the contract method 0x7bfe789a.
//
// Solidity: function frontRunningInterval() view returns(uint32)
func (_Leveragedpool *LeveragedpoolSession) FrontRunningInterval() (uint32, error) {
	return _Leveragedpool.Contract.FrontRunningInterval(&_Leveragedpool.CallOpts)
}

// FrontRunningInterval is a free data retrieval call binding the contract method 0x7bfe789a.
//
// Solidity: function frontRunningInterval() view returns(uint32)
func (_Leveragedpool *LeveragedpoolCallerSession) FrontRunningInterval() (uint32, error) {
	return _Leveragedpool.Contract.FrontRunningInterval(&_Leveragedpool.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) GetFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) GetFee() (*big.Int, error) {
	return _Leveragedpool.Contract.GetFee(&_Leveragedpool.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) GetFee() (*big.Int, error) {
	return _Leveragedpool.Contract.GetFee(&_Leveragedpool.CallOpts)
}

// GetLeverage is a free data retrieval call binding the contract method 0xbaba68de.
//
// Solidity: function getLeverage() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) GetLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "getLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLeverage is a free data retrieval call binding the contract method 0xbaba68de.
//
// Solidity: function getLeverage() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) GetLeverage() (*big.Int, error) {
	return _Leveragedpool.Contract.GetLeverage(&_Leveragedpool.CallOpts)
}

// GetLeverage is a free data retrieval call binding the contract method 0xbaba68de.
//
// Solidity: function getLeverage() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) GetLeverage() (*big.Int, error) {
	return _Leveragedpool.Contract.GetLeverage(&_Leveragedpool.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(int256)
func (_Leveragedpool *LeveragedpoolCaller) GetOraclePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "getOraclePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(int256)
func (_Leveragedpool *LeveragedpoolSession) GetOraclePrice() (*big.Int, error) {
	return _Leveragedpool.Contract.GetOraclePrice(&_Leveragedpool.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(int256)
func (_Leveragedpool *LeveragedpoolCallerSession) GetOraclePrice() (*big.Int, error) {
	return _Leveragedpool.Contract.GetOraclePrice(&_Leveragedpool.CallOpts)
}

// GetUpkeepInformation is a free data retrieval call binding the contract method 0x04883c27.
//
// Solidity: function getUpkeepInformation() view returns(int256, bytes, uint256, uint256)
func (_Leveragedpool *LeveragedpoolCaller) GetUpkeepInformation(opts *bind.CallOpts) (*big.Int, []byte, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "getUpkeepInformation")

	if err != nil {
		return *new(*big.Int), *new([]byte), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetUpkeepInformation is a free data retrieval call binding the contract method 0x04883c27.
//
// Solidity: function getUpkeepInformation() view returns(int256, bytes, uint256, uint256)
func (_Leveragedpool *LeveragedpoolSession) GetUpkeepInformation() (*big.Int, []byte, *big.Int, *big.Int, error) {
	return _Leveragedpool.Contract.GetUpkeepInformation(&_Leveragedpool.CallOpts)
}

// GetUpkeepInformation is a free data retrieval call binding the contract method 0x04883c27.
//
// Solidity: function getUpkeepInformation() view returns(int256, bytes, uint256, uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) GetUpkeepInformation() (*big.Int, []byte, *big.Int, *big.Int, error) {
	return _Leveragedpool.Contract.GetUpkeepInformation(&_Leveragedpool.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) Governance() (common.Address, error) {
	return _Leveragedpool.Contract.Governance(&_Leveragedpool.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) Governance() (common.Address, error) {
	return _Leveragedpool.Contract.Governance(&_Leveragedpool.CallOpts)
}

// GovernanceTransferInProgress is a free data retrieval call binding the contract method 0xf164e2a7.
//
// Solidity: function governanceTransferInProgress() view returns(bool)
func (_Leveragedpool *LeveragedpoolCaller) GovernanceTransferInProgress(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "governanceTransferInProgress")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GovernanceTransferInProgress is a free data retrieval call binding the contract method 0xf164e2a7.
//
// Solidity: function governanceTransferInProgress() view returns(bool)
func (_Leveragedpool *LeveragedpoolSession) GovernanceTransferInProgress() (bool, error) {
	return _Leveragedpool.Contract.GovernanceTransferInProgress(&_Leveragedpool.CallOpts)
}

// GovernanceTransferInProgress is a free data retrieval call binding the contract method 0xf164e2a7.
//
// Solidity: function governanceTransferInProgress() view returns(bool)
func (_Leveragedpool *LeveragedpoolCallerSession) GovernanceTransferInProgress() (bool, error) {
	return _Leveragedpool.Contract.GovernanceTransferInProgress(&_Leveragedpool.CallOpts)
}

// IntervalPassed is a free data retrieval call binding the contract method 0x7bbf1033.
//
// Solidity: function intervalPassed() view returns(bool)
func (_Leveragedpool *LeveragedpoolCaller) IntervalPassed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "intervalPassed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IntervalPassed is a free data retrieval call binding the contract method 0x7bbf1033.
//
// Solidity: function intervalPassed() view returns(bool)
func (_Leveragedpool *LeveragedpoolSession) IntervalPassed() (bool, error) {
	return _Leveragedpool.Contract.IntervalPassed(&_Leveragedpool.CallOpts)
}

// IntervalPassed is a free data retrieval call binding the contract method 0x7bbf1033.
//
// Solidity: function intervalPassed() view returns(bool)
func (_Leveragedpool *LeveragedpoolCallerSession) IntervalPassed() (bool, error) {
	return _Leveragedpool.Contract.IntervalPassed(&_Leveragedpool.CallOpts)
}

// InvariantCheck is a free data retrieval call binding the contract method 0x4509017e.
//
// Solidity: function invariantCheck() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) InvariantCheck(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "invariantCheck")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InvariantCheck is a free data retrieval call binding the contract method 0x4509017e.
//
// Solidity: function invariantCheck() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) InvariantCheck() (common.Address, error) {
	return _Leveragedpool.Contract.InvariantCheck(&_Leveragedpool.CallOpts)
}

// InvariantCheck is a free data retrieval call binding the contract method 0x4509017e.
//
// Solidity: function invariantCheck() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) InvariantCheck() (common.Address, error) {
	return _Leveragedpool.Contract.InvariantCheck(&_Leveragedpool.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) Keeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "keeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) Keeper() (common.Address, error) {
	return _Leveragedpool.Contract.Keeper(&_Leveragedpool.CallOpts)
}

// Keeper is a free data retrieval call binding the contract method 0xaced1661.
//
// Solidity: function keeper() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) Keeper() (common.Address, error) {
	return _Leveragedpool.Contract.Keeper(&_Leveragedpool.CallOpts)
}

// LastPriceTimestamp is a free data retrieval call binding the contract method 0x7de93f93.
//
// Solidity: function lastPriceTimestamp() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) LastPriceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "lastPriceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPriceTimestamp is a free data retrieval call binding the contract method 0x7de93f93.
//
// Solidity: function lastPriceTimestamp() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) LastPriceTimestamp() (*big.Int, error) {
	return _Leveragedpool.Contract.LastPriceTimestamp(&_Leveragedpool.CallOpts)
}

// LastPriceTimestamp is a free data retrieval call binding the contract method 0x7de93f93.
//
// Solidity: function lastPriceTimestamp() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) LastPriceTimestamp() (*big.Int, error) {
	return _Leveragedpool.Contract.LastPriceTimestamp(&_Leveragedpool.CallOpts)
}

// LeverageAmount is a free data retrieval call binding the contract method 0xab47c2b3.
//
// Solidity: function leverageAmount() view returns(bytes16)
func (_Leveragedpool *LeveragedpoolCaller) LeverageAmount(opts *bind.CallOpts) ([16]byte, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "leverageAmount")

	if err != nil {
		return *new([16]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([16]byte)).(*[16]byte)

	return out0, err

}

// LeverageAmount is a free data retrieval call binding the contract method 0xab47c2b3.
//
// Solidity: function leverageAmount() view returns(bytes16)
func (_Leveragedpool *LeveragedpoolSession) LeverageAmount() ([16]byte, error) {
	return _Leveragedpool.Contract.LeverageAmount(&_Leveragedpool.CallOpts)
}

// LeverageAmount is a free data retrieval call binding the contract method 0xab47c2b3.
//
// Solidity: function leverageAmount() view returns(bytes16)
func (_Leveragedpool *LeveragedpoolCallerSession) LeverageAmount() ([16]byte, error) {
	return _Leveragedpool.Contract.LeverageAmount(&_Leveragedpool.CallOpts)
}

// LongBalance is a free data retrieval call binding the contract method 0x7e71fc7d.
//
// Solidity: function longBalance() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) LongBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "longBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LongBalance is a free data retrieval call binding the contract method 0x7e71fc7d.
//
// Solidity: function longBalance() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) LongBalance() (*big.Int, error) {
	return _Leveragedpool.Contract.LongBalance(&_Leveragedpool.CallOpts)
}

// LongBalance is a free data retrieval call binding the contract method 0x7e71fc7d.
//
// Solidity: function longBalance() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) LongBalance() (*big.Int, error) {
	return _Leveragedpool.Contract.LongBalance(&_Leveragedpool.CallOpts)
}

// OracleWrapper is a free data retrieval call binding the contract method 0xb9ed8abf.
//
// Solidity: function oracleWrapper() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) OracleWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "oracleWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleWrapper is a free data retrieval call binding the contract method 0xb9ed8abf.
//
// Solidity: function oracleWrapper() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) OracleWrapper() (common.Address, error) {
	return _Leveragedpool.Contract.OracleWrapper(&_Leveragedpool.CallOpts)
}

// OracleWrapper is a free data retrieval call binding the contract method 0xb9ed8abf.
//
// Solidity: function oracleWrapper() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) OracleWrapper() (common.Address, error) {
	return _Leveragedpool.Contract.OracleWrapper(&_Leveragedpool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Leveragedpool *LeveragedpoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Leveragedpool *LeveragedpoolSession) Paused() (bool, error) {
	return _Leveragedpool.Contract.Paused(&_Leveragedpool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Leveragedpool *LeveragedpoolCallerSession) Paused() (bool, error) {
	return _Leveragedpool.Contract.Paused(&_Leveragedpool.CallOpts)
}

// PoolCommitter is a free data retrieval call binding the contract method 0xcd39f30f.
//
// Solidity: function poolCommitter() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) PoolCommitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "poolCommitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolCommitter is a free data retrieval call binding the contract method 0xcd39f30f.
//
// Solidity: function poolCommitter() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) PoolCommitter() (common.Address, error) {
	return _Leveragedpool.Contract.PoolCommitter(&_Leveragedpool.CallOpts)
}

// PoolCommitter is a free data retrieval call binding the contract method 0xcd39f30f.
//
// Solidity: function poolCommitter() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) PoolCommitter() (common.Address, error) {
	return _Leveragedpool.Contract.PoolCommitter(&_Leveragedpool.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_Leveragedpool *LeveragedpoolCaller) PoolName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "poolName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_Leveragedpool *LeveragedpoolSession) PoolName() (string, error) {
	return _Leveragedpool.Contract.PoolName(&_Leveragedpool.CallOpts)
}

// PoolName is a free data retrieval call binding the contract method 0xf3466dfa.
//
// Solidity: function poolName() view returns(string)
func (_Leveragedpool *LeveragedpoolCallerSession) PoolName() (string, error) {
	return _Leveragedpool.Contract.PoolName(&_Leveragedpool.CallOpts)
}

// PoolTokens is a free data retrieval call binding the contract method 0x6d3e313e.
//
// Solidity: function poolTokens() view returns(address[2])
func (_Leveragedpool *LeveragedpoolCaller) PoolTokens(opts *bind.CallOpts) ([2]common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "poolTokens")

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// PoolTokens is a free data retrieval call binding the contract method 0x6d3e313e.
//
// Solidity: function poolTokens() view returns(address[2])
func (_Leveragedpool *LeveragedpoolSession) PoolTokens() ([2]common.Address, error) {
	return _Leveragedpool.Contract.PoolTokens(&_Leveragedpool.CallOpts)
}

// PoolTokens is a free data retrieval call binding the contract method 0x6d3e313e.
//
// Solidity: function poolTokens() view returns(address[2])
func (_Leveragedpool *LeveragedpoolCallerSession) PoolTokens() ([2]common.Address, error) {
	return _Leveragedpool.Contract.PoolTokens(&_Leveragedpool.CallOpts)
}

// PrimaryFees is a free data retrieval call binding the contract method 0x89e59154.
//
// Solidity: function primaryFees() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) PrimaryFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "primaryFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrimaryFees is a free data retrieval call binding the contract method 0x89e59154.
//
// Solidity: function primaryFees() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) PrimaryFees() (*big.Int, error) {
	return _Leveragedpool.Contract.PrimaryFees(&_Leveragedpool.CallOpts)
}

// PrimaryFees is a free data retrieval call binding the contract method 0x89e59154.
//
// Solidity: function primaryFees() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) PrimaryFees() (*big.Int, error) {
	return _Leveragedpool.Contract.PrimaryFees(&_Leveragedpool.CallOpts)
}

// ProvisionalGovernance is a free data retrieval call binding the contract method 0x9fd55245.
//
// Solidity: function provisionalGovernance() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) ProvisionalGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "provisionalGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProvisionalGovernance is a free data retrieval call binding the contract method 0x9fd55245.
//
// Solidity: function provisionalGovernance() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) ProvisionalGovernance() (common.Address, error) {
	return _Leveragedpool.Contract.ProvisionalGovernance(&_Leveragedpool.CallOpts)
}

// ProvisionalGovernance is a free data retrieval call binding the contract method 0x9fd55245.
//
// Solidity: function provisionalGovernance() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) ProvisionalGovernance() (common.Address, error) {
	return _Leveragedpool.Contract.ProvisionalGovernance(&_Leveragedpool.CallOpts)
}

// SecondaryFeeAddress is a free data retrieval call binding the contract method 0xe3d346cd.
//
// Solidity: function secondaryFeeAddress() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) SecondaryFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "secondaryFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SecondaryFeeAddress is a free data retrieval call binding the contract method 0xe3d346cd.
//
// Solidity: function secondaryFeeAddress() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) SecondaryFeeAddress() (common.Address, error) {
	return _Leveragedpool.Contract.SecondaryFeeAddress(&_Leveragedpool.CallOpts)
}

// SecondaryFeeAddress is a free data retrieval call binding the contract method 0xe3d346cd.
//
// Solidity: function secondaryFeeAddress() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) SecondaryFeeAddress() (common.Address, error) {
	return _Leveragedpool.Contract.SecondaryFeeAddress(&_Leveragedpool.CallOpts)
}

// SecondaryFeeSplitPercent is a free data retrieval call binding the contract method 0x03850dbb.
//
// Solidity: function secondaryFeeSplitPercent() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) SecondaryFeeSplitPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "secondaryFeeSplitPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondaryFeeSplitPercent is a free data retrieval call binding the contract method 0x03850dbb.
//
// Solidity: function secondaryFeeSplitPercent() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) SecondaryFeeSplitPercent() (*big.Int, error) {
	return _Leveragedpool.Contract.SecondaryFeeSplitPercent(&_Leveragedpool.CallOpts)
}

// SecondaryFeeSplitPercent is a free data retrieval call binding the contract method 0x03850dbb.
//
// Solidity: function secondaryFeeSplitPercent() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) SecondaryFeeSplitPercent() (*big.Int, error) {
	return _Leveragedpool.Contract.SecondaryFeeSplitPercent(&_Leveragedpool.CallOpts)
}

// SecondaryFees is a free data retrieval call binding the contract method 0x46fdc99a.
//
// Solidity: function secondaryFees() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) SecondaryFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "secondaryFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondaryFees is a free data retrieval call binding the contract method 0x46fdc99a.
//
// Solidity: function secondaryFees() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) SecondaryFees() (*big.Int, error) {
	return _Leveragedpool.Contract.SecondaryFees(&_Leveragedpool.CallOpts)
}

// SecondaryFees is a free data retrieval call binding the contract method 0x46fdc99a.
//
// Solidity: function secondaryFees() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) SecondaryFees() (*big.Int, error) {
	return _Leveragedpool.Contract.SecondaryFees(&_Leveragedpool.CallOpts)
}

// SettlementEthOracle is a free data retrieval call binding the contract method 0x8226f396.
//
// Solidity: function settlementEthOracle() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) SettlementEthOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "settlementEthOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SettlementEthOracle is a free data retrieval call binding the contract method 0x8226f396.
//
// Solidity: function settlementEthOracle() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) SettlementEthOracle() (common.Address, error) {
	return _Leveragedpool.Contract.SettlementEthOracle(&_Leveragedpool.CallOpts)
}

// SettlementEthOracle is a free data retrieval call binding the contract method 0x8226f396.
//
// Solidity: function settlementEthOracle() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) SettlementEthOracle() (common.Address, error) {
	return _Leveragedpool.Contract.SettlementEthOracle(&_Leveragedpool.CallOpts)
}

// SettlementToken is a free data retrieval call binding the contract method 0x7b9e618d.
//
// Solidity: function settlementToken() view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) SettlementToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "settlementToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SettlementToken is a free data retrieval call binding the contract method 0x7b9e618d.
//
// Solidity: function settlementToken() view returns(address)
func (_Leveragedpool *LeveragedpoolSession) SettlementToken() (common.Address, error) {
	return _Leveragedpool.Contract.SettlementToken(&_Leveragedpool.CallOpts)
}

// SettlementToken is a free data retrieval call binding the contract method 0x7b9e618d.
//
// Solidity: function settlementToken() view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) SettlementToken() (common.Address, error) {
	return _Leveragedpool.Contract.SettlementToken(&_Leveragedpool.CallOpts)
}

// ShortBalance is a free data retrieval call binding the contract method 0xba8d5468.
//
// Solidity: function shortBalance() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCaller) ShortBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "shortBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShortBalance is a free data retrieval call binding the contract method 0xba8d5468.
//
// Solidity: function shortBalance() view returns(uint256)
func (_Leveragedpool *LeveragedpoolSession) ShortBalance() (*big.Int, error) {
	return _Leveragedpool.Contract.ShortBalance(&_Leveragedpool.CallOpts)
}

// ShortBalance is a free data retrieval call binding the contract method 0xba8d5468.
//
// Solidity: function shortBalance() view returns(uint256)
func (_Leveragedpool *LeveragedpoolCallerSession) ShortBalance() (*big.Int, error) {
	return _Leveragedpool.Contract.ShortBalance(&_Leveragedpool.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Leveragedpool *LeveragedpoolCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Leveragedpool *LeveragedpoolSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Leveragedpool.Contract.Tokens(&_Leveragedpool.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Leveragedpool *LeveragedpoolCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Leveragedpool.Contract.Tokens(&_Leveragedpool.CallOpts, arg0)
}

// UpdateInterval is a free data retrieval call binding the contract method 0xfd2c80ae.
//
// Solidity: function updateInterval() view returns(uint32)
func (_Leveragedpool *LeveragedpoolCaller) UpdateInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Leveragedpool.contract.Call(opts, &out, "updateInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// UpdateInterval is a free data retrieval call binding the contract method 0xfd2c80ae.
//
// Solidity: function updateInterval() view returns(uint32)
func (_Leveragedpool *LeveragedpoolSession) UpdateInterval() (uint32, error) {
	return _Leveragedpool.Contract.UpdateInterval(&_Leveragedpool.CallOpts)
}

// UpdateInterval is a free data retrieval call binding the contract method 0xfd2c80ae.
//
// Solidity: function updateInterval() view returns(uint32)
func (_Leveragedpool *LeveragedpoolCallerSession) UpdateInterval() (uint32, error) {
	return _Leveragedpool.Contract.UpdateInterval(&_Leveragedpool.CallOpts)
}

// BurnTokens is a paid mutator transaction binding the contract method 0x87f9ca5d.
//
// Solidity: function burnTokens(uint256 tokenType, uint256 amount, address burner) returns()
func (_Leveragedpool *LeveragedpoolTransactor) BurnTokens(opts *bind.TransactOpts, tokenType *big.Int, amount *big.Int, burner common.Address) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "burnTokens", tokenType, amount, burner)
}

// BurnTokens is a paid mutator transaction binding the contract method 0x87f9ca5d.
//
// Solidity: function burnTokens(uint256 tokenType, uint256 amount, address burner) returns()
func (_Leveragedpool *LeveragedpoolSession) BurnTokens(tokenType *big.Int, amount *big.Int, burner common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.BurnTokens(&_Leveragedpool.TransactOpts, tokenType, amount, burner)
}

// BurnTokens is a paid mutator transaction binding the contract method 0x87f9ca5d.
//
// Solidity: function burnTokens(uint256 tokenType, uint256 amount, address burner) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) BurnTokens(tokenType *big.Int, amount *big.Int, burner common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.BurnTokens(&_Leveragedpool.TransactOpts, tokenType, amount, burner)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_Leveragedpool *LeveragedpoolTransactor) ClaimGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "claimGovernance")
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_Leveragedpool *LeveragedpoolSession) ClaimGovernance() (*types.Transaction, error) {
	return _Leveragedpool.Contract.ClaimGovernance(&_Leveragedpool.TransactOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) ClaimGovernance() (*types.Transaction, error) {
	return _Leveragedpool.Contract.ClaimGovernance(&_Leveragedpool.TransactOpts)
}

// ClaimPrimaryFees is a paid mutator transaction binding the contract method 0x2fd7d643.
//
// Solidity: function claimPrimaryFees() returns()
func (_Leveragedpool *LeveragedpoolTransactor) ClaimPrimaryFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "claimPrimaryFees")
}

// ClaimPrimaryFees is a paid mutator transaction binding the contract method 0x2fd7d643.
//
// Solidity: function claimPrimaryFees() returns()
func (_Leveragedpool *LeveragedpoolSession) ClaimPrimaryFees() (*types.Transaction, error) {
	return _Leveragedpool.Contract.ClaimPrimaryFees(&_Leveragedpool.TransactOpts)
}

// ClaimPrimaryFees is a paid mutator transaction binding the contract method 0x2fd7d643.
//
// Solidity: function claimPrimaryFees() returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) ClaimPrimaryFees() (*types.Transaction, error) {
	return _Leveragedpool.Contract.ClaimPrimaryFees(&_Leveragedpool.TransactOpts)
}

// ClaimSecondaryFees is a paid mutator transaction binding the contract method 0xd11262a0.
//
// Solidity: function claimSecondaryFees() returns()
func (_Leveragedpool *LeveragedpoolTransactor) ClaimSecondaryFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "claimSecondaryFees")
}

// ClaimSecondaryFees is a paid mutator transaction binding the contract method 0xd11262a0.
//
// Solidity: function claimSecondaryFees() returns()
func (_Leveragedpool *LeveragedpoolSession) ClaimSecondaryFees() (*types.Transaction, error) {
	return _Leveragedpool.Contract.ClaimSecondaryFees(&_Leveragedpool.TransactOpts)
}

// ClaimSecondaryFees is a paid mutator transaction binding the contract method 0xd11262a0.
//
// Solidity: function claimSecondaryFees() returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) ClaimSecondaryFees() (*types.Transaction, error) {
	return _Leveragedpool.Contract.ClaimSecondaryFees(&_Leveragedpool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x129011d0.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,uint32,uint32,uint16,uint256,address,address,address,uint256) initialization) returns()
func (_Leveragedpool *LeveragedpoolTransactor) Initialize(opts *bind.TransactOpts, initialization ILeveragedPoolInitialization) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "initialize", initialization)
}

// Initialize is a paid mutator transaction binding the contract method 0x129011d0.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,uint32,uint32,uint16,uint256,address,address,address,uint256) initialization) returns()
func (_Leveragedpool *LeveragedpoolSession) Initialize(initialization ILeveragedPoolInitialization) (*types.Transaction, error) {
	return _Leveragedpool.Contract.Initialize(&_Leveragedpool.TransactOpts, initialization)
}

// Initialize is a paid mutator transaction binding the contract method 0x129011d0.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,uint32,uint32,uint16,uint256,address,address,address,uint256) initialization) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) Initialize(initialization ILeveragedPoolInitialization) (*types.Transaction, error) {
	return _Leveragedpool.Contract.Initialize(&_Leveragedpool.TransactOpts, initialization)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Leveragedpool *LeveragedpoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Leveragedpool *LeveragedpoolSession) Pause() (*types.Transaction, error) {
	return _Leveragedpool.Contract.Pause(&_Leveragedpool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) Pause() (*types.Transaction, error) {
	return _Leveragedpool.Contract.Pause(&_Leveragedpool.TransactOpts)
}

// PayKeeperFromBalances is a paid mutator transaction binding the contract method 0x6dc2b271.
//
// Solidity: function payKeeperFromBalances(address to, uint256 amount) returns(bool)
func (_Leveragedpool *LeveragedpoolTransactor) PayKeeperFromBalances(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "payKeeperFromBalances", to, amount)
}

// PayKeeperFromBalances is a paid mutator transaction binding the contract method 0x6dc2b271.
//
// Solidity: function payKeeperFromBalances(address to, uint256 amount) returns(bool)
func (_Leveragedpool *LeveragedpoolSession) PayKeeperFromBalances(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.PayKeeperFromBalances(&_Leveragedpool.TransactOpts, to, amount)
}

// PayKeeperFromBalances is a paid mutator transaction binding the contract method 0x6dc2b271.
//
// Solidity: function payKeeperFromBalances(address to, uint256 amount) returns(bool)
func (_Leveragedpool *LeveragedpoolTransactorSession) PayKeeperFromBalances(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.PayKeeperFromBalances(&_Leveragedpool.TransactOpts, to, amount)
}

// PoolTokenTransfer is a paid mutator transaction binding the contract method 0x110b1bed.
//
// Solidity: function poolTokenTransfer(bool isLongToken, address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolTransactor) PoolTokenTransfer(opts *bind.TransactOpts, isLongToken bool, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "poolTokenTransfer", isLongToken, to, amount)
}

// PoolTokenTransfer is a paid mutator transaction binding the contract method 0x110b1bed.
//
// Solidity: function poolTokenTransfer(bool isLongToken, address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolSession) PoolTokenTransfer(isLongToken bool, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.PoolTokenTransfer(&_Leveragedpool.TransactOpts, isLongToken, to, amount)
}

// PoolTokenTransfer is a paid mutator transaction binding the contract method 0x110b1bed.
//
// Solidity: function poolTokenTransfer(bool isLongToken, address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) PoolTokenTransfer(isLongToken bool, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.PoolTokenTransfer(&_Leveragedpool.TransactOpts, isLongToken, to, amount)
}

// PoolUpkeep is a paid mutator transaction binding the contract method 0xf633a740.
//
// Solidity: function poolUpkeep(int256 _oldPrice, int256 _newPrice) returns()
func (_Leveragedpool *LeveragedpoolTransactor) PoolUpkeep(opts *bind.TransactOpts, _oldPrice *big.Int, _newPrice *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "poolUpkeep", _oldPrice, _newPrice)
}

// PoolUpkeep is a paid mutator transaction binding the contract method 0xf633a740.
//
// Solidity: function poolUpkeep(int256 _oldPrice, int256 _newPrice) returns()
func (_Leveragedpool *LeveragedpoolSession) PoolUpkeep(_oldPrice *big.Int, _newPrice *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.PoolUpkeep(&_Leveragedpool.TransactOpts, _oldPrice, _newPrice)
}

// PoolUpkeep is a paid mutator transaction binding the contract method 0xf633a740.
//
// Solidity: function poolUpkeep(int256 _oldPrice, int256 _newPrice) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) PoolUpkeep(_oldPrice *big.Int, _newPrice *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.PoolUpkeep(&_Leveragedpool.TransactOpts, _oldPrice, _newPrice)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_Leveragedpool *LeveragedpoolTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "setKeeper", _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_Leveragedpool *LeveragedpoolSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SetKeeper(&_Leveragedpool.TransactOpts, _keeper)
}

// SetKeeper is a paid mutator transaction binding the contract method 0x748747e6.
//
// Solidity: function setKeeper(address _keeper) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) SetKeeper(_keeper common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SetKeeper(&_Leveragedpool.TransactOpts, _keeper)
}

// SetNewPoolBalances is a paid mutator transaction binding the contract method 0x3c3f8252.
//
// Solidity: function setNewPoolBalances(uint256 _longBalance, uint256 _shortBalance) returns()
func (_Leveragedpool *LeveragedpoolTransactor) SetNewPoolBalances(opts *bind.TransactOpts, _longBalance *big.Int, _shortBalance *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "setNewPoolBalances", _longBalance, _shortBalance)
}

// SetNewPoolBalances is a paid mutator transaction binding the contract method 0x3c3f8252.
//
// Solidity: function setNewPoolBalances(uint256 _longBalance, uint256 _shortBalance) returns()
func (_Leveragedpool *LeveragedpoolSession) SetNewPoolBalances(_longBalance *big.Int, _shortBalance *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SetNewPoolBalances(&_Leveragedpool.TransactOpts, _longBalance, _shortBalance)
}

// SetNewPoolBalances is a paid mutator transaction binding the contract method 0x3c3f8252.
//
// Solidity: function setNewPoolBalances(uint256 _longBalance, uint256 _shortBalance) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) SetNewPoolBalances(_longBalance *big.Int, _shortBalance *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SetNewPoolBalances(&_Leveragedpool.TransactOpts, _longBalance, _shortBalance)
}

// SettlementTokenTransfer is a paid mutator transaction binding the contract method 0xd0b06f09.
//
// Solidity: function settlementTokenTransfer(address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolTransactor) SettlementTokenTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "settlementTokenTransfer", to, amount)
}

// SettlementTokenTransfer is a paid mutator transaction binding the contract method 0xd0b06f09.
//
// Solidity: function settlementTokenTransfer(address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolSession) SettlementTokenTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SettlementTokenTransfer(&_Leveragedpool.TransactOpts, to, amount)
}

// SettlementTokenTransfer is a paid mutator transaction binding the contract method 0xd0b06f09.
//
// Solidity: function settlementTokenTransfer(address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) SettlementTokenTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SettlementTokenTransfer(&_Leveragedpool.TransactOpts, to, amount)
}

// SettlementTokenTransferFrom is a paid mutator transaction binding the contract method 0x0d20a1dc.
//
// Solidity: function settlementTokenTransferFrom(address from, address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolTransactor) SettlementTokenTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "settlementTokenTransferFrom", from, to, amount)
}

// SettlementTokenTransferFrom is a paid mutator transaction binding the contract method 0x0d20a1dc.
//
// Solidity: function settlementTokenTransferFrom(address from, address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolSession) SettlementTokenTransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SettlementTokenTransferFrom(&_Leveragedpool.TransactOpts, from, to, amount)
}

// SettlementTokenTransferFrom is a paid mutator transaction binding the contract method 0x0d20a1dc.
//
// Solidity: function settlementTokenTransferFrom(address from, address to, uint256 amount) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) SettlementTokenTransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Leveragedpool.Contract.SettlementTokenTransferFrom(&_Leveragedpool.TransactOpts, from, to, amount)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_Leveragedpool *LeveragedpoolTransactor) TransferGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "transferGovernance", _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_Leveragedpool *LeveragedpoolSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.TransferGovernance(&_Leveragedpool.TransactOpts, _governance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _governance) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) TransferGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.TransferGovernance(&_Leveragedpool.TransactOpts, _governance)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Leveragedpool *LeveragedpoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Leveragedpool *LeveragedpoolSession) Unpause() (*types.Transaction, error) {
	return _Leveragedpool.Contract.Unpause(&_Leveragedpool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _Leveragedpool.Contract.Unpause(&_Leveragedpool.TransactOpts)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address account) returns()
func (_Leveragedpool *LeveragedpoolTransactor) UpdateFeeAddress(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "updateFeeAddress", account)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address account) returns()
func (_Leveragedpool *LeveragedpoolSession) UpdateFeeAddress(account common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.UpdateFeeAddress(&_Leveragedpool.TransactOpts, account)
}

// UpdateFeeAddress is a paid mutator transaction binding the contract method 0xbbcaac38.
//
// Solidity: function updateFeeAddress(address account) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) UpdateFeeAddress(account common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.UpdateFeeAddress(&_Leveragedpool.TransactOpts, account)
}

// UpdateSecondaryFeeAddress is a paid mutator transaction binding the contract method 0x75a024a2.
//
// Solidity: function updateSecondaryFeeAddress(address account) returns()
func (_Leveragedpool *LeveragedpoolTransactor) UpdateSecondaryFeeAddress(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "updateSecondaryFeeAddress", account)
}

// UpdateSecondaryFeeAddress is a paid mutator transaction binding the contract method 0x75a024a2.
//
// Solidity: function updateSecondaryFeeAddress(address account) returns()
func (_Leveragedpool *LeveragedpoolSession) UpdateSecondaryFeeAddress(account common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.UpdateSecondaryFeeAddress(&_Leveragedpool.TransactOpts, account)
}

// UpdateSecondaryFeeAddress is a paid mutator transaction binding the contract method 0x75a024a2.
//
// Solidity: function updateSecondaryFeeAddress(address account) returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) UpdateSecondaryFeeAddress(account common.Address) (*types.Transaction, error) {
	return _Leveragedpool.Contract.UpdateSecondaryFeeAddress(&_Leveragedpool.TransactOpts, account)
}

// WithdrawSettlement is a paid mutator transaction binding the contract method 0x0c96871a.
//
// Solidity: function withdrawSettlement() returns()
func (_Leveragedpool *LeveragedpoolTransactor) WithdrawSettlement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Leveragedpool.contract.Transact(opts, "withdrawSettlement")
}

// WithdrawSettlement is a paid mutator transaction binding the contract method 0x0c96871a.
//
// Solidity: function withdrawSettlement() returns()
func (_Leveragedpool *LeveragedpoolSession) WithdrawSettlement() (*types.Transaction, error) {
	return _Leveragedpool.Contract.WithdrawSettlement(&_Leveragedpool.TransactOpts)
}

// WithdrawSettlement is a paid mutator transaction binding the contract method 0x0c96871a.
//
// Solidity: function withdrawSettlement() returns()
func (_Leveragedpool *LeveragedpoolTransactorSession) WithdrawSettlement() (*types.Transaction, error) {
	return _Leveragedpool.Contract.WithdrawSettlement(&_Leveragedpool.TransactOpts)
}

// LeveragedpoolFeeAddressUpdatedIterator is returned from FilterFeeAddressUpdated and is used to iterate over the raw logs and unpacked data for FeeAddressUpdated events raised by the Leveragedpool contract.
type LeveragedpoolFeeAddressUpdatedIterator struct {
	Event *LeveragedpoolFeeAddressUpdated // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolFeeAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolFeeAddressUpdated)
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
		it.Event = new(LeveragedpoolFeeAddressUpdated)
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
func (it *LeveragedpoolFeeAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolFeeAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolFeeAddressUpdated represents a FeeAddressUpdated event raised by the Leveragedpool contract.
type LeveragedpoolFeeAddressUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeAddressUpdated is a free log retrieval operation binding the contract event 0x11f35a22548bcd4c3788ab4a7e4fba427a2014f02e5d5e2da9af62212c03183f.
//
// Solidity: event FeeAddressUpdated(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) FilterFeeAddressUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*LeveragedpoolFeeAddressUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "FeeAddressUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolFeeAddressUpdatedIterator{contract: _Leveragedpool.contract, event: "FeeAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeAddressUpdated is a free log subscription operation binding the contract event 0x11f35a22548bcd4c3788ab4a7e4fba427a2014f02e5d5e2da9af62212c03183f.
//
// Solidity: event FeeAddressUpdated(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) WatchFeeAddressUpdated(opts *bind.WatchOpts, sink chan<- *LeveragedpoolFeeAddressUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "FeeAddressUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolFeeAddressUpdated)
				if err := _Leveragedpool.contract.UnpackLog(event, "FeeAddressUpdated", log); err != nil {
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

// ParseFeeAddressUpdated is a log parse operation binding the contract event 0x11f35a22548bcd4c3788ab4a7e4fba427a2014f02e5d5e2da9af62212c03183f.
//
// Solidity: event FeeAddressUpdated(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) ParseFeeAddressUpdated(log types.Log) (*LeveragedpoolFeeAddressUpdated, error) {
	event := new(LeveragedpoolFeeAddressUpdated)
	if err := _Leveragedpool.contract.UnpackLog(event, "FeeAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolGovernanceAddressChangedIterator is returned from FilterGovernanceAddressChanged and is used to iterate over the raw logs and unpacked data for GovernanceAddressChanged events raised by the Leveragedpool contract.
type LeveragedpoolGovernanceAddressChangedIterator struct {
	Event *LeveragedpoolGovernanceAddressChanged // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolGovernanceAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolGovernanceAddressChanged)
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
		it.Event = new(LeveragedpoolGovernanceAddressChanged)
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
func (it *LeveragedpoolGovernanceAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolGovernanceAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolGovernanceAddressChanged represents a GovernanceAddressChanged event raised by the Leveragedpool contract.
type LeveragedpoolGovernanceAddressChanged struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGovernanceAddressChanged is a free log retrieval operation binding the contract event 0x023588d3d1dcaad34e471c9e01b616b794156174bc693539c8fe15c0bfd5d826.
//
// Solidity: event GovernanceAddressChanged(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) FilterGovernanceAddressChanged(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*LeveragedpoolGovernanceAddressChangedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "GovernanceAddressChanged", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolGovernanceAddressChangedIterator{contract: _Leveragedpool.contract, event: "GovernanceAddressChanged", logs: logs, sub: sub}, nil
}

// WatchGovernanceAddressChanged is a free log subscription operation binding the contract event 0x023588d3d1dcaad34e471c9e01b616b794156174bc693539c8fe15c0bfd5d826.
//
// Solidity: event GovernanceAddressChanged(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) WatchGovernanceAddressChanged(opts *bind.WatchOpts, sink chan<- *LeveragedpoolGovernanceAddressChanged, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "GovernanceAddressChanged", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolGovernanceAddressChanged)
				if err := _Leveragedpool.contract.UnpackLog(event, "GovernanceAddressChanged", log); err != nil {
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

// ParseGovernanceAddressChanged is a log parse operation binding the contract event 0x023588d3d1dcaad34e471c9e01b616b794156174bc693539c8fe15c0bfd5d826.
//
// Solidity: event GovernanceAddressChanged(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) ParseGovernanceAddressChanged(log types.Log) (*LeveragedpoolGovernanceAddressChanged, error) {
	event := new(LeveragedpoolGovernanceAddressChanged)
	if err := _Leveragedpool.contract.UnpackLog(event, "GovernanceAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolKeeperAddressChangedIterator is returned from FilterKeeperAddressChanged and is used to iterate over the raw logs and unpacked data for KeeperAddressChanged events raised by the Leveragedpool contract.
type LeveragedpoolKeeperAddressChangedIterator struct {
	Event *LeveragedpoolKeeperAddressChanged // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolKeeperAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolKeeperAddressChanged)
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
		it.Event = new(LeveragedpoolKeeperAddressChanged)
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
func (it *LeveragedpoolKeeperAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolKeeperAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolKeeperAddressChanged represents a KeeperAddressChanged event raised by the Leveragedpool contract.
type LeveragedpoolKeeperAddressChanged struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterKeeperAddressChanged is a free log retrieval operation binding the contract event 0x402b3f9a8de3e388e7653c7a5892204fe18b579c8c23db19d6e00f1043ceb921.
//
// Solidity: event KeeperAddressChanged(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) FilterKeeperAddressChanged(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*LeveragedpoolKeeperAddressChangedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "KeeperAddressChanged", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolKeeperAddressChangedIterator{contract: _Leveragedpool.contract, event: "KeeperAddressChanged", logs: logs, sub: sub}, nil
}

// WatchKeeperAddressChanged is a free log subscription operation binding the contract event 0x402b3f9a8de3e388e7653c7a5892204fe18b579c8c23db19d6e00f1043ceb921.
//
// Solidity: event KeeperAddressChanged(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) WatchKeeperAddressChanged(opts *bind.WatchOpts, sink chan<- *LeveragedpoolKeeperAddressChanged, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "KeeperAddressChanged", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolKeeperAddressChanged)
				if err := _Leveragedpool.contract.UnpackLog(event, "KeeperAddressChanged", log); err != nil {
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

// ParseKeeperAddressChanged is a log parse operation binding the contract event 0x402b3f9a8de3e388e7653c7a5892204fe18b579c8c23db19d6e00f1043ceb921.
//
// Solidity: event KeeperAddressChanged(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) ParseKeeperAddressChanged(log types.Log) (*LeveragedpoolKeeperAddressChanged, error) {
	event := new(LeveragedpoolKeeperAddressChanged)
	if err := _Leveragedpool.contract.UnpackLog(event, "KeeperAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Leveragedpool contract.
type LeveragedpoolPausedIterator struct {
	Event *LeveragedpoolPaused // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolPaused)
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
		it.Event = new(LeveragedpoolPaused)
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
func (it *LeveragedpoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolPaused represents a Paused event raised by the Leveragedpool contract.
type LeveragedpoolPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Leveragedpool *LeveragedpoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LeveragedpoolPausedIterator, error) {

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolPausedIterator{contract: _Leveragedpool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Leveragedpool *LeveragedpoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LeveragedpoolPaused) (event.Subscription, error) {

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolPaused)
				if err := _Leveragedpool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Leveragedpool *LeveragedpoolFilterer) ParsePaused(log types.Log) (*LeveragedpoolPaused, error) {
	event := new(LeveragedpoolPaused)
	if err := _Leveragedpool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolPoolBalancesChangedIterator is returned from FilterPoolBalancesChanged and is used to iterate over the raw logs and unpacked data for PoolBalancesChanged events raised by the Leveragedpool contract.
type LeveragedpoolPoolBalancesChangedIterator struct {
	Event *LeveragedpoolPoolBalancesChanged // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolPoolBalancesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolPoolBalancesChanged)
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
		it.Event = new(LeveragedpoolPoolBalancesChanged)
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
func (it *LeveragedpoolPoolBalancesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolPoolBalancesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolPoolBalancesChanged represents a PoolBalancesChanged event raised by the Leveragedpool contract.
type LeveragedpoolPoolBalancesChanged struct {
	Long  *big.Int
	Short *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPoolBalancesChanged is a free log retrieval operation binding the contract event 0xb31885164844201997ff1f41bcad765029c634f5183880298a4c69ec4540d89f.
//
// Solidity: event PoolBalancesChanged(uint256 indexed long, uint256 indexed short)
func (_Leveragedpool *LeveragedpoolFilterer) FilterPoolBalancesChanged(opts *bind.FilterOpts, long []*big.Int, short []*big.Int) (*LeveragedpoolPoolBalancesChangedIterator, error) {

	var longRule []interface{}
	for _, longItem := range long {
		longRule = append(longRule, longItem)
	}
	var shortRule []interface{}
	for _, shortItem := range short {
		shortRule = append(shortRule, shortItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "PoolBalancesChanged", longRule, shortRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolPoolBalancesChangedIterator{contract: _Leveragedpool.contract, event: "PoolBalancesChanged", logs: logs, sub: sub}, nil
}

// WatchPoolBalancesChanged is a free log subscription operation binding the contract event 0xb31885164844201997ff1f41bcad765029c634f5183880298a4c69ec4540d89f.
//
// Solidity: event PoolBalancesChanged(uint256 indexed long, uint256 indexed short)
func (_Leveragedpool *LeveragedpoolFilterer) WatchPoolBalancesChanged(opts *bind.WatchOpts, sink chan<- *LeveragedpoolPoolBalancesChanged, long []*big.Int, short []*big.Int) (event.Subscription, error) {

	var longRule []interface{}
	for _, longItem := range long {
		longRule = append(longRule, longItem)
	}
	var shortRule []interface{}
	for _, shortItem := range short {
		shortRule = append(shortRule, shortItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "PoolBalancesChanged", longRule, shortRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolPoolBalancesChanged)
				if err := _Leveragedpool.contract.UnpackLog(event, "PoolBalancesChanged", log); err != nil {
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

// ParsePoolBalancesChanged is a log parse operation binding the contract event 0xb31885164844201997ff1f41bcad765029c634f5183880298a4c69ec4540d89f.
//
// Solidity: event PoolBalancesChanged(uint256 indexed long, uint256 indexed short)
func (_Leveragedpool *LeveragedpoolFilterer) ParsePoolBalancesChanged(log types.Log) (*LeveragedpoolPoolBalancesChanged, error) {
	event := new(LeveragedpoolPoolBalancesChanged)
	if err := _Leveragedpool.contract.UnpackLog(event, "PoolBalancesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolPoolInitializedIterator is returned from FilterPoolInitialized and is used to iterate over the raw logs and unpacked data for PoolInitialized events raised by the Leveragedpool contract.
type LeveragedpoolPoolInitializedIterator struct {
	Event *LeveragedpoolPoolInitialized // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolPoolInitialized)
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
		it.Event = new(LeveragedpoolPoolInitialized)
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
func (it *LeveragedpoolPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolPoolInitialized represents a PoolInitialized event raised by the Leveragedpool contract.
type LeveragedpoolPoolInitialized struct {
	LongToken       common.Address
	ShortToken      common.Address
	SettlementToken common.Address
	PoolName        string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0xbc8bd61e5dd69ff5cb8c389529c029503515ba368eaafbdaa406e072fdd36072.
//
// Solidity: event PoolInitialized(address indexed longToken, address indexed shortToken, address settlementToken, string poolName)
func (_Leveragedpool *LeveragedpoolFilterer) FilterPoolInitialized(opts *bind.FilterOpts, longToken []common.Address, shortToken []common.Address) (*LeveragedpoolPoolInitializedIterator, error) {

	var longTokenRule []interface{}
	for _, longTokenItem := range longToken {
		longTokenRule = append(longTokenRule, longTokenItem)
	}
	var shortTokenRule []interface{}
	for _, shortTokenItem := range shortToken {
		shortTokenRule = append(shortTokenRule, shortTokenItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "PoolInitialized", longTokenRule, shortTokenRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolPoolInitializedIterator{contract: _Leveragedpool.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0xbc8bd61e5dd69ff5cb8c389529c029503515ba368eaafbdaa406e072fdd36072.
//
// Solidity: event PoolInitialized(address indexed longToken, address indexed shortToken, address settlementToken, string poolName)
func (_Leveragedpool *LeveragedpoolFilterer) WatchPoolInitialized(opts *bind.WatchOpts, sink chan<- *LeveragedpoolPoolInitialized, longToken []common.Address, shortToken []common.Address) (event.Subscription, error) {

	var longTokenRule []interface{}
	for _, longTokenItem := range longToken {
		longTokenRule = append(longTokenRule, longTokenItem)
	}
	var shortTokenRule []interface{}
	for _, shortTokenItem := range shortToken {
		shortTokenRule = append(shortTokenRule, shortTokenItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "PoolInitialized", longTokenRule, shortTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolPoolInitialized)
				if err := _Leveragedpool.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
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

// ParsePoolInitialized is a log parse operation binding the contract event 0xbc8bd61e5dd69ff5cb8c389529c029503515ba368eaafbdaa406e072fdd36072.
//
// Solidity: event PoolInitialized(address indexed longToken, address indexed shortToken, address settlementToken, string poolName)
func (_Leveragedpool *LeveragedpoolFilterer) ParsePoolInitialized(log types.Log) (*LeveragedpoolPoolInitialized, error) {
	event := new(LeveragedpoolPoolInitialized)
	if err := _Leveragedpool.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolPoolRebalanceIterator is returned from FilterPoolRebalance and is used to iterate over the raw logs and unpacked data for PoolRebalance events raised by the Leveragedpool contract.
type LeveragedpoolPoolRebalanceIterator struct {
	Event *LeveragedpoolPoolRebalance // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolPoolRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolPoolRebalance)
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
		it.Event = new(LeveragedpoolPoolRebalance)
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
func (it *LeveragedpoolPoolRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolPoolRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolPoolRebalance represents a PoolRebalance event raised by the Leveragedpool contract.
type LeveragedpoolPoolRebalance struct {
	ShortBalanceChange *big.Int
	LongBalanceChange  *big.Int
	ShortFeeAmount     *big.Int
	LongFeeAmount      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPoolRebalance is a free log retrieval operation binding the contract event 0xf1e6677cbfc1dcb704bc91fc51ba87e318c4ed18c6a5571907e52a09a2c0f9aa.
//
// Solidity: event PoolRebalance(int256 shortBalanceChange, int256 longBalanceChange, uint256 shortFeeAmount, uint256 longFeeAmount)
func (_Leveragedpool *LeveragedpoolFilterer) FilterPoolRebalance(opts *bind.FilterOpts) (*LeveragedpoolPoolRebalanceIterator, error) {

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "PoolRebalance")
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolPoolRebalanceIterator{contract: _Leveragedpool.contract, event: "PoolRebalance", logs: logs, sub: sub}, nil
}

// WatchPoolRebalance is a free log subscription operation binding the contract event 0xf1e6677cbfc1dcb704bc91fc51ba87e318c4ed18c6a5571907e52a09a2c0f9aa.
//
// Solidity: event PoolRebalance(int256 shortBalanceChange, int256 longBalanceChange, uint256 shortFeeAmount, uint256 longFeeAmount)
func (_Leveragedpool *LeveragedpoolFilterer) WatchPoolRebalance(opts *bind.WatchOpts, sink chan<- *LeveragedpoolPoolRebalance) (event.Subscription, error) {

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "PoolRebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolPoolRebalance)
				if err := _Leveragedpool.contract.UnpackLog(event, "PoolRebalance", log); err != nil {
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

// ParsePoolRebalance is a log parse operation binding the contract event 0xf1e6677cbfc1dcb704bc91fc51ba87e318c4ed18c6a5571907e52a09a2c0f9aa.
//
// Solidity: event PoolRebalance(int256 shortBalanceChange, int256 longBalanceChange, uint256 shortFeeAmount, uint256 longFeeAmount)
func (_Leveragedpool *LeveragedpoolFilterer) ParsePoolRebalance(log types.Log) (*LeveragedpoolPoolRebalance, error) {
	event := new(LeveragedpoolPoolRebalance)
	if err := _Leveragedpool.contract.UnpackLog(event, "PoolRebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolPriceChangeErrorIterator is returned from FilterPriceChangeError and is used to iterate over the raw logs and unpacked data for PriceChangeError events raised by the Leveragedpool contract.
type LeveragedpoolPriceChangeErrorIterator struct {
	Event *LeveragedpoolPriceChangeError // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolPriceChangeErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolPriceChangeError)
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
		it.Event = new(LeveragedpoolPriceChangeError)
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
func (it *LeveragedpoolPriceChangeErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolPriceChangeErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolPriceChangeError represents a PriceChangeError event raised by the Leveragedpool contract.
type LeveragedpoolPriceChangeError struct {
	StartPrice *big.Int
	EndPrice   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceChangeError is a free log retrieval operation binding the contract event 0x27f70c863f1bd3e335bb4d6bf2a31075c861752f0b21a03b816a356035b1dec6.
//
// Solidity: event PriceChangeError(int256 indexed startPrice, int256 indexed endPrice)
func (_Leveragedpool *LeveragedpoolFilterer) FilterPriceChangeError(opts *bind.FilterOpts, startPrice []*big.Int, endPrice []*big.Int) (*LeveragedpoolPriceChangeErrorIterator, error) {

	var startPriceRule []interface{}
	for _, startPriceItem := range startPrice {
		startPriceRule = append(startPriceRule, startPriceItem)
	}
	var endPriceRule []interface{}
	for _, endPriceItem := range endPrice {
		endPriceRule = append(endPriceRule, endPriceItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "PriceChangeError", startPriceRule, endPriceRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolPriceChangeErrorIterator{contract: _Leveragedpool.contract, event: "PriceChangeError", logs: logs, sub: sub}, nil
}

// WatchPriceChangeError is a free log subscription operation binding the contract event 0x27f70c863f1bd3e335bb4d6bf2a31075c861752f0b21a03b816a356035b1dec6.
//
// Solidity: event PriceChangeError(int256 indexed startPrice, int256 indexed endPrice)
func (_Leveragedpool *LeveragedpoolFilterer) WatchPriceChangeError(opts *bind.WatchOpts, sink chan<- *LeveragedpoolPriceChangeError, startPrice []*big.Int, endPrice []*big.Int) (event.Subscription, error) {

	var startPriceRule []interface{}
	for _, startPriceItem := range startPrice {
		startPriceRule = append(startPriceRule, startPriceItem)
	}
	var endPriceRule []interface{}
	for _, endPriceItem := range endPrice {
		endPriceRule = append(endPriceRule, endPriceItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "PriceChangeError", startPriceRule, endPriceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolPriceChangeError)
				if err := _Leveragedpool.contract.UnpackLog(event, "PriceChangeError", log); err != nil {
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

// ParsePriceChangeError is a log parse operation binding the contract event 0x27f70c863f1bd3e335bb4d6bf2a31075c861752f0b21a03b816a356035b1dec6.
//
// Solidity: event PriceChangeError(int256 indexed startPrice, int256 indexed endPrice)
func (_Leveragedpool *LeveragedpoolFilterer) ParsePriceChangeError(log types.Log) (*LeveragedpoolPriceChangeError, error) {
	event := new(LeveragedpoolPriceChangeError)
	if err := _Leveragedpool.contract.UnpackLog(event, "PriceChangeError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolPrimaryFeesPaidIterator is returned from FilterPrimaryFeesPaid and is used to iterate over the raw logs and unpacked data for PrimaryFeesPaid events raised by the Leveragedpool contract.
type LeveragedpoolPrimaryFeesPaidIterator struct {
	Event *LeveragedpoolPrimaryFeesPaid // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolPrimaryFeesPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolPrimaryFeesPaid)
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
		it.Event = new(LeveragedpoolPrimaryFeesPaid)
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
func (it *LeveragedpoolPrimaryFeesPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolPrimaryFeesPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolPrimaryFeesPaid represents a PrimaryFeesPaid event raised by the Leveragedpool contract.
type LeveragedpoolPrimaryFeesPaid struct {
	FeeAddress common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPrimaryFeesPaid is a free log retrieval operation binding the contract event 0xe756bfac49549fbcb50e6d0f75a98c96a16261fbaf06f41dd35de05f2902fa4e.
//
// Solidity: event PrimaryFeesPaid(address indexed feeAddress, uint256 amount)
func (_Leveragedpool *LeveragedpoolFilterer) FilterPrimaryFeesPaid(opts *bind.FilterOpts, feeAddress []common.Address) (*LeveragedpoolPrimaryFeesPaidIterator, error) {

	var feeAddressRule []interface{}
	for _, feeAddressItem := range feeAddress {
		feeAddressRule = append(feeAddressRule, feeAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "PrimaryFeesPaid", feeAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolPrimaryFeesPaidIterator{contract: _Leveragedpool.contract, event: "PrimaryFeesPaid", logs: logs, sub: sub}, nil
}

// WatchPrimaryFeesPaid is a free log subscription operation binding the contract event 0xe756bfac49549fbcb50e6d0f75a98c96a16261fbaf06f41dd35de05f2902fa4e.
//
// Solidity: event PrimaryFeesPaid(address indexed feeAddress, uint256 amount)
func (_Leveragedpool *LeveragedpoolFilterer) WatchPrimaryFeesPaid(opts *bind.WatchOpts, sink chan<- *LeveragedpoolPrimaryFeesPaid, feeAddress []common.Address) (event.Subscription, error) {

	var feeAddressRule []interface{}
	for _, feeAddressItem := range feeAddress {
		feeAddressRule = append(feeAddressRule, feeAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "PrimaryFeesPaid", feeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolPrimaryFeesPaid)
				if err := _Leveragedpool.contract.UnpackLog(event, "PrimaryFeesPaid", log); err != nil {
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

// ParsePrimaryFeesPaid is a log parse operation binding the contract event 0xe756bfac49549fbcb50e6d0f75a98c96a16261fbaf06f41dd35de05f2902fa4e.
//
// Solidity: event PrimaryFeesPaid(address indexed feeAddress, uint256 amount)
func (_Leveragedpool *LeveragedpoolFilterer) ParsePrimaryFeesPaid(log types.Log) (*LeveragedpoolPrimaryFeesPaid, error) {
	event := new(LeveragedpoolPrimaryFeesPaid)
	if err := _Leveragedpool.contract.UnpackLog(event, "PrimaryFeesPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolProvisionalGovernanceChangedIterator is returned from FilterProvisionalGovernanceChanged and is used to iterate over the raw logs and unpacked data for ProvisionalGovernanceChanged events raised by the Leveragedpool contract.
type LeveragedpoolProvisionalGovernanceChangedIterator struct {
	Event *LeveragedpoolProvisionalGovernanceChanged // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolProvisionalGovernanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolProvisionalGovernanceChanged)
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
		it.Event = new(LeveragedpoolProvisionalGovernanceChanged)
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
func (it *LeveragedpoolProvisionalGovernanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolProvisionalGovernanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolProvisionalGovernanceChanged represents a ProvisionalGovernanceChanged event raised by the Leveragedpool contract.
type LeveragedpoolProvisionalGovernanceChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProvisionalGovernanceChanged is a free log retrieval operation binding the contract event 0x35681f4f23137fb58510a9854f1b6e95f90a2cf0b66d2fce4df74cc0becc82d5.
//
// Solidity: event ProvisionalGovernanceChanged(address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) FilterProvisionalGovernanceChanged(opts *bind.FilterOpts, newAddress []common.Address) (*LeveragedpoolProvisionalGovernanceChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "ProvisionalGovernanceChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolProvisionalGovernanceChangedIterator{contract: _Leveragedpool.contract, event: "ProvisionalGovernanceChanged", logs: logs, sub: sub}, nil
}

// WatchProvisionalGovernanceChanged is a free log subscription operation binding the contract event 0x35681f4f23137fb58510a9854f1b6e95f90a2cf0b66d2fce4df74cc0becc82d5.
//
// Solidity: event ProvisionalGovernanceChanged(address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) WatchProvisionalGovernanceChanged(opts *bind.WatchOpts, sink chan<- *LeveragedpoolProvisionalGovernanceChanged, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "ProvisionalGovernanceChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolProvisionalGovernanceChanged)
				if err := _Leveragedpool.contract.UnpackLog(event, "ProvisionalGovernanceChanged", log); err != nil {
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

// ParseProvisionalGovernanceChanged is a log parse operation binding the contract event 0x35681f4f23137fb58510a9854f1b6e95f90a2cf0b66d2fce4df74cc0becc82d5.
//
// Solidity: event ProvisionalGovernanceChanged(address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) ParseProvisionalGovernanceChanged(log types.Log) (*LeveragedpoolProvisionalGovernanceChanged, error) {
	event := new(LeveragedpoolProvisionalGovernanceChanged)
	if err := _Leveragedpool.contract.UnpackLog(event, "ProvisionalGovernanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolSecondaryFeeAddressUpdatedIterator is returned from FilterSecondaryFeeAddressUpdated and is used to iterate over the raw logs and unpacked data for SecondaryFeeAddressUpdated events raised by the Leveragedpool contract.
type LeveragedpoolSecondaryFeeAddressUpdatedIterator struct {
	Event *LeveragedpoolSecondaryFeeAddressUpdated // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolSecondaryFeeAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolSecondaryFeeAddressUpdated)
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
		it.Event = new(LeveragedpoolSecondaryFeeAddressUpdated)
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
func (it *LeveragedpoolSecondaryFeeAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolSecondaryFeeAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolSecondaryFeeAddressUpdated represents a SecondaryFeeAddressUpdated event raised by the Leveragedpool contract.
type LeveragedpoolSecondaryFeeAddressUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSecondaryFeeAddressUpdated is a free log retrieval operation binding the contract event 0x4a1fe1b44fae78218ab520acca2aaba5067e706a4084bccf913f291c8d6e0749.
//
// Solidity: event SecondaryFeeAddressUpdated(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) FilterSecondaryFeeAddressUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*LeveragedpoolSecondaryFeeAddressUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "SecondaryFeeAddressUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolSecondaryFeeAddressUpdatedIterator{contract: _Leveragedpool.contract, event: "SecondaryFeeAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchSecondaryFeeAddressUpdated is a free log subscription operation binding the contract event 0x4a1fe1b44fae78218ab520acca2aaba5067e706a4084bccf913f291c8d6e0749.
//
// Solidity: event SecondaryFeeAddressUpdated(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) WatchSecondaryFeeAddressUpdated(opts *bind.WatchOpts, sink chan<- *LeveragedpoolSecondaryFeeAddressUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "SecondaryFeeAddressUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolSecondaryFeeAddressUpdated)
				if err := _Leveragedpool.contract.UnpackLog(event, "SecondaryFeeAddressUpdated", log); err != nil {
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

// ParseSecondaryFeeAddressUpdated is a log parse operation binding the contract event 0x4a1fe1b44fae78218ab520acca2aaba5067e706a4084bccf913f291c8d6e0749.
//
// Solidity: event SecondaryFeeAddressUpdated(address indexed oldAddress, address indexed newAddress)
func (_Leveragedpool *LeveragedpoolFilterer) ParseSecondaryFeeAddressUpdated(log types.Log) (*LeveragedpoolSecondaryFeeAddressUpdated, error) {
	event := new(LeveragedpoolSecondaryFeeAddressUpdated)
	if err := _Leveragedpool.contract.UnpackLog(event, "SecondaryFeeAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolSecondaryFeesPaidIterator is returned from FilterSecondaryFeesPaid and is used to iterate over the raw logs and unpacked data for SecondaryFeesPaid events raised by the Leveragedpool contract.
type LeveragedpoolSecondaryFeesPaidIterator struct {
	Event *LeveragedpoolSecondaryFeesPaid // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolSecondaryFeesPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolSecondaryFeesPaid)
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
		it.Event = new(LeveragedpoolSecondaryFeesPaid)
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
func (it *LeveragedpoolSecondaryFeesPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolSecondaryFeesPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolSecondaryFeesPaid represents a SecondaryFeesPaid event raised by the Leveragedpool contract.
type LeveragedpoolSecondaryFeesPaid struct {
	SecondaryFeeAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSecondaryFeesPaid is a free log retrieval operation binding the contract event 0x2f76018fcfcd8aaaa8299da890bcca2dfebfc638250a89e5785b34bdc68b6363.
//
// Solidity: event SecondaryFeesPaid(address indexed secondaryFeeAddress, uint256 amount)
func (_Leveragedpool *LeveragedpoolFilterer) FilterSecondaryFeesPaid(opts *bind.FilterOpts, secondaryFeeAddress []common.Address) (*LeveragedpoolSecondaryFeesPaidIterator, error) {

	var secondaryFeeAddressRule []interface{}
	for _, secondaryFeeAddressItem := range secondaryFeeAddress {
		secondaryFeeAddressRule = append(secondaryFeeAddressRule, secondaryFeeAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "SecondaryFeesPaid", secondaryFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolSecondaryFeesPaidIterator{contract: _Leveragedpool.contract, event: "SecondaryFeesPaid", logs: logs, sub: sub}, nil
}

// WatchSecondaryFeesPaid is a free log subscription operation binding the contract event 0x2f76018fcfcd8aaaa8299da890bcca2dfebfc638250a89e5785b34bdc68b6363.
//
// Solidity: event SecondaryFeesPaid(address indexed secondaryFeeAddress, uint256 amount)
func (_Leveragedpool *LeveragedpoolFilterer) WatchSecondaryFeesPaid(opts *bind.WatchOpts, sink chan<- *LeveragedpoolSecondaryFeesPaid, secondaryFeeAddress []common.Address) (event.Subscription, error) {

	var secondaryFeeAddressRule []interface{}
	for _, secondaryFeeAddressItem := range secondaryFeeAddress {
		secondaryFeeAddressRule = append(secondaryFeeAddressRule, secondaryFeeAddressItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "SecondaryFeesPaid", secondaryFeeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolSecondaryFeesPaid)
				if err := _Leveragedpool.contract.UnpackLog(event, "SecondaryFeesPaid", log); err != nil {
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

// ParseSecondaryFeesPaid is a log parse operation binding the contract event 0x2f76018fcfcd8aaaa8299da890bcca2dfebfc638250a89e5785b34bdc68b6363.
//
// Solidity: event SecondaryFeesPaid(address indexed secondaryFeeAddress, uint256 amount)
func (_Leveragedpool *LeveragedpoolFilterer) ParseSecondaryFeesPaid(log types.Log) (*LeveragedpoolSecondaryFeesPaid, error) {
	event := new(LeveragedpoolSecondaryFeesPaid)
	if err := _Leveragedpool.contract.UnpackLog(event, "SecondaryFeesPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolSettlementWithdrawnIterator is returned from FilterSettlementWithdrawn and is used to iterate over the raw logs and unpacked data for SettlementWithdrawn events raised by the Leveragedpool contract.
type LeveragedpoolSettlementWithdrawnIterator struct {
	Event *LeveragedpoolSettlementWithdrawn // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolSettlementWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolSettlementWithdrawn)
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
		it.Event = new(LeveragedpoolSettlementWithdrawn)
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
func (it *LeveragedpoolSettlementWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolSettlementWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolSettlementWithdrawn represents a SettlementWithdrawn event raised by the Leveragedpool contract.
type LeveragedpoolSettlementWithdrawn struct {
	To       common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSettlementWithdrawn is a free log retrieval operation binding the contract event 0x3304a35fcccaae1037a5405395dfc497e010ed92b42060cc21ac54cfb0e3e850.
//
// Solidity: event SettlementWithdrawn(address indexed to, uint256 indexed quantity)
func (_Leveragedpool *LeveragedpoolFilterer) FilterSettlementWithdrawn(opts *bind.FilterOpts, to []common.Address, quantity []*big.Int) (*LeveragedpoolSettlementWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var quantityRule []interface{}
	for _, quantityItem := range quantity {
		quantityRule = append(quantityRule, quantityItem)
	}

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "SettlementWithdrawn", toRule, quantityRule)
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolSettlementWithdrawnIterator{contract: _Leveragedpool.contract, event: "SettlementWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSettlementWithdrawn is a free log subscription operation binding the contract event 0x3304a35fcccaae1037a5405395dfc497e010ed92b42060cc21ac54cfb0e3e850.
//
// Solidity: event SettlementWithdrawn(address indexed to, uint256 indexed quantity)
func (_Leveragedpool *LeveragedpoolFilterer) WatchSettlementWithdrawn(opts *bind.WatchOpts, sink chan<- *LeveragedpoolSettlementWithdrawn, to []common.Address, quantity []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var quantityRule []interface{}
	for _, quantityItem := range quantity {
		quantityRule = append(quantityRule, quantityItem)
	}

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "SettlementWithdrawn", toRule, quantityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolSettlementWithdrawn)
				if err := _Leveragedpool.contract.UnpackLog(event, "SettlementWithdrawn", log); err != nil {
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

// ParseSettlementWithdrawn is a log parse operation binding the contract event 0x3304a35fcccaae1037a5405395dfc497e010ed92b42060cc21ac54cfb0e3e850.
//
// Solidity: event SettlementWithdrawn(address indexed to, uint256 indexed quantity)
func (_Leveragedpool *LeveragedpoolFilterer) ParseSettlementWithdrawn(log types.Log) (*LeveragedpoolSettlementWithdrawn, error) {
	event := new(LeveragedpoolSettlementWithdrawn)
	if err := _Leveragedpool.contract.UnpackLog(event, "SettlementWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeveragedpoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Leveragedpool contract.
type LeveragedpoolUnpausedIterator struct {
	Event *LeveragedpoolUnpaused // Event containing the contract specifics and raw log

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
func (it *LeveragedpoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeveragedpoolUnpaused)
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
		it.Event = new(LeveragedpoolUnpaused)
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
func (it *LeveragedpoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeveragedpoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeveragedpoolUnpaused represents a Unpaused event raised by the Leveragedpool contract.
type LeveragedpoolUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Leveragedpool *LeveragedpoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LeveragedpoolUnpausedIterator, error) {

	logs, sub, err := _Leveragedpool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LeveragedpoolUnpausedIterator{contract: _Leveragedpool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Leveragedpool *LeveragedpoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LeveragedpoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Leveragedpool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeveragedpoolUnpaused)
				if err := _Leveragedpool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Leveragedpool *LeveragedpoolFilterer) ParseUnpaused(log types.Log) (*LeveragedpoolUnpaused, error) {
	event := new(LeveragedpoolUnpaused)
	if err := _Leveragedpool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
