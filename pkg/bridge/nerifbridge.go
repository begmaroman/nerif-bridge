// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nerifbridge

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

// NerifBridgeMetaData contains all meta data concerning the NerifBridge contract.
var NerifBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Send\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"addSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_senders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"rec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"removeSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"senders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NerifBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use NerifBridgeMetaData.ABI instead.
var NerifBridgeABI = NerifBridgeMetaData.ABI

// NerifBridge is an auto generated Go binding around an Ethereum contract.
type NerifBridge struct {
	NerifBridgeCaller     // Read-only binding to the contract
	NerifBridgeTransactor // Write-only binding to the contract
	NerifBridgeFilterer   // Log filterer for contract events
}

// NerifBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NerifBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NerifBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NerifBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NerifBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NerifBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NerifBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NerifBridgeSession struct {
	Contract     *NerifBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NerifBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NerifBridgeCallerSession struct {
	Contract *NerifBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NerifBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NerifBridgeTransactorSession struct {
	Contract     *NerifBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NerifBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NerifBridgeRaw struct {
	Contract *NerifBridge // Generic contract binding to access the raw methods on
}

// NerifBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NerifBridgeCallerRaw struct {
	Contract *NerifBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// NerifBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NerifBridgeTransactorRaw struct {
	Contract *NerifBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNerifBridge creates a new instance of NerifBridge, bound to a specific deployed contract.
func NewNerifBridge(address common.Address, backend bind.ContractBackend) (*NerifBridge, error) {
	contract, err := bindNerifBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NerifBridge{NerifBridgeCaller: NerifBridgeCaller{contract: contract}, NerifBridgeTransactor: NerifBridgeTransactor{contract: contract}, NerifBridgeFilterer: NerifBridgeFilterer{contract: contract}}, nil
}

// NewNerifBridgeCaller creates a new read-only instance of NerifBridge, bound to a specific deployed contract.
func NewNerifBridgeCaller(address common.Address, caller bind.ContractCaller) (*NerifBridgeCaller, error) {
	contract, err := bindNerifBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NerifBridgeCaller{contract: contract}, nil
}

// NewNerifBridgeTransactor creates a new write-only instance of NerifBridge, bound to a specific deployed contract.
func NewNerifBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*NerifBridgeTransactor, error) {
	contract, err := bindNerifBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NerifBridgeTransactor{contract: contract}, nil
}

// NewNerifBridgeFilterer creates a new log filterer instance of NerifBridge, bound to a specific deployed contract.
func NewNerifBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*NerifBridgeFilterer, error) {
	contract, err := bindNerifBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NerifBridgeFilterer{contract: contract}, nil
}

// bindNerifBridge binds a generic wrapper to an already deployed contract.
func bindNerifBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NerifBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NerifBridge *NerifBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NerifBridge.Contract.NerifBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NerifBridge *NerifBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NerifBridge.Contract.NerifBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NerifBridge *NerifBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NerifBridge.Contract.NerifBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NerifBridge *NerifBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NerifBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NerifBridge *NerifBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NerifBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NerifBridge *NerifBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NerifBridge.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NerifBridge *NerifBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NerifBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NerifBridge *NerifBridgeSession) Owner() (common.Address, error) {
	return _NerifBridge.Contract.Owner(&_NerifBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NerifBridge *NerifBridgeCallerSession) Owner() (common.Address, error) {
	return _NerifBridge.Contract.Owner(&_NerifBridge.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_NerifBridge *NerifBridgeCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NerifBridge.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_NerifBridge *NerifBridgeSession) Registry() (common.Address, error) {
	return _NerifBridge.Contract.Registry(&_NerifBridge.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_NerifBridge *NerifBridgeCallerSession) Registry() (common.Address, error) {
	return _NerifBridge.Contract.Registry(&_NerifBridge.CallOpts)
}

// Senders is a free data retrieval call binding the contract method 0x982fb9d8.
//
// Solidity: function senders(address ) view returns(bool)
func (_NerifBridge *NerifBridgeCaller) Senders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NerifBridge.contract.Call(opts, &out, "senders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Senders is a free data retrieval call binding the contract method 0x982fb9d8.
//
// Solidity: function senders(address ) view returns(bool)
func (_NerifBridge *NerifBridgeSession) Senders(arg0 common.Address) (bool, error) {
	return _NerifBridge.Contract.Senders(&_NerifBridge.CallOpts, arg0)
}

// Senders is a free data retrieval call binding the contract method 0x982fb9d8.
//
// Solidity: function senders(address ) view returns(bool)
func (_NerifBridge *NerifBridgeCallerSession) Senders(arg0 common.Address) (bool, error) {
	return _NerifBridge.Contract.Senders(&_NerifBridge.CallOpts, arg0)
}

// AddSender is a paid mutator transaction binding the contract method 0xb697f531.
//
// Solidity: function addSender(address sender) returns()
func (_NerifBridge *NerifBridgeTransactor) AddSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "addSender", sender)
}

// AddSender is a paid mutator transaction binding the contract method 0xb697f531.
//
// Solidity: function addSender(address sender) returns()
func (_NerifBridge *NerifBridgeSession) AddSender(sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.AddSender(&_NerifBridge.TransactOpts, sender)
}

// AddSender is a paid mutator transaction binding the contract method 0xb697f531.
//
// Solidity: function addSender(address sender) returns()
func (_NerifBridge *NerifBridgeTransactorSession) AddSender(sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.AddSender(&_NerifBridge.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _senders, address _registry) returns()
func (_NerifBridge *NerifBridgeTransactor) Initialize(opts *bind.TransactOpts, _senders []common.Address, _registry common.Address) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "initialize", _senders, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _senders, address _registry) returns()
func (_NerifBridge *NerifBridgeSession) Initialize(_senders []common.Address, _registry common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.Initialize(&_NerifBridge.TransactOpts, _senders, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _senders, address _registry) returns()
func (_NerifBridge *NerifBridgeTransactorSession) Initialize(_senders []common.Address, _registry common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.Initialize(&_NerifBridge.TransactOpts, _senders, _registry)
}

// Rec is a paid mutator transaction binding the contract method 0x83a622de.
//
// Solidity: function rec(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender) returns()
func (_NerifBridge *NerifBridgeTransactor) Rec(opts *bind.TransactOpts, chainId *big.Int, target common.Address, payload []byte, gasAmount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "rec", chainId, target, payload, gasAmount, sender)
}

// Rec is a paid mutator transaction binding the contract method 0x83a622de.
//
// Solidity: function rec(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender) returns()
func (_NerifBridge *NerifBridgeSession) Rec(chainId *big.Int, target common.Address, payload []byte, gasAmount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.Rec(&_NerifBridge.TransactOpts, chainId, target, payload, gasAmount, sender)
}

// Rec is a paid mutator transaction binding the contract method 0x83a622de.
//
// Solidity: function rec(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender) returns()
func (_NerifBridge *NerifBridgeTransactorSession) Rec(chainId *big.Int, target common.Address, payload []byte, gasAmount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.Rec(&_NerifBridge.TransactOpts, chainId, target, payload, gasAmount, sender)
}

// RemoveSender is a paid mutator transaction binding the contract method 0xb2f87643.
//
// Solidity: function removeSender(address sender) returns()
func (_NerifBridge *NerifBridgeTransactor) RemoveSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "removeSender", sender)
}

// RemoveSender is a paid mutator transaction binding the contract method 0xb2f87643.
//
// Solidity: function removeSender(address sender) returns()
func (_NerifBridge *NerifBridgeSession) RemoveSender(sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.RemoveSender(&_NerifBridge.TransactOpts, sender)
}

// RemoveSender is a paid mutator transaction binding the contract method 0xb2f87643.
//
// Solidity: function removeSender(address sender) returns()
func (_NerifBridge *NerifBridgeTransactorSession) RemoveSender(sender common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.RemoveSender(&_NerifBridge.TransactOpts, sender)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NerifBridge *NerifBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NerifBridge *NerifBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NerifBridge.Contract.RenounceOwnership(&_NerifBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NerifBridge *NerifBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NerifBridge.Contract.RenounceOwnership(&_NerifBridge.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x720a995b.
//
// Solidity: function send(uint256 chainId, address target, bytes payload, uint256 gasAmount) returns()
func (_NerifBridge *NerifBridgeTransactor) Send(opts *bind.TransactOpts, chainId *big.Int, target common.Address, payload []byte, gasAmount *big.Int) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "send", chainId, target, payload, gasAmount)
}

// Send is a paid mutator transaction binding the contract method 0x720a995b.
//
// Solidity: function send(uint256 chainId, address target, bytes payload, uint256 gasAmount) returns()
func (_NerifBridge *NerifBridgeSession) Send(chainId *big.Int, target common.Address, payload []byte, gasAmount *big.Int) (*types.Transaction, error) {
	return _NerifBridge.Contract.Send(&_NerifBridge.TransactOpts, chainId, target, payload, gasAmount)
}

// Send is a paid mutator transaction binding the contract method 0x720a995b.
//
// Solidity: function send(uint256 chainId, address target, bytes payload, uint256 gasAmount) returns()
func (_NerifBridge *NerifBridgeTransactorSession) Send(chainId *big.Int, target common.Address, payload []byte, gasAmount *big.Int) (*types.Transaction, error) {
	return _NerifBridge.Contract.Send(&_NerifBridge.TransactOpts, chainId, target, payload, gasAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NerifBridge *NerifBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NerifBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NerifBridge *NerifBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.TransferOwnership(&_NerifBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NerifBridge *NerifBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NerifBridge.Contract.TransferOwnership(&_NerifBridge.TransactOpts, newOwner)
}

// NerifBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NerifBridge contract.
type NerifBridgeInitializedIterator struct {
	Event *NerifBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *NerifBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifBridgeInitialized)
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
		it.Event = new(NerifBridgeInitialized)
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
func (it *NerifBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifBridgeInitialized represents a Initialized event raised by the NerifBridge contract.
type NerifBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NerifBridge *NerifBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NerifBridgeInitializedIterator, error) {

	logs, sub, err := _NerifBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NerifBridgeInitializedIterator{contract: _NerifBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NerifBridge *NerifBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NerifBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _NerifBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifBridgeInitialized)
				if err := _NerifBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NerifBridge *NerifBridgeFilterer) ParseInitialized(log types.Log) (*NerifBridgeInitialized, error) {
	event := new(NerifBridgeInitialized)
	if err := _NerifBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NerifBridge contract.
type NerifBridgeOwnershipTransferredIterator struct {
	Event *NerifBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NerifBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifBridgeOwnershipTransferred)
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
		it.Event = new(NerifBridgeOwnershipTransferred)
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
func (it *NerifBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the NerifBridge contract.
type NerifBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NerifBridge *NerifBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NerifBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NerifBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NerifBridgeOwnershipTransferredIterator{contract: _NerifBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NerifBridge *NerifBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NerifBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NerifBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifBridgeOwnershipTransferred)
				if err := _NerifBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NerifBridge *NerifBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*NerifBridgeOwnershipTransferred, error) {
	event := new(NerifBridgeOwnershipTransferred)
	if err := _NerifBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifBridgeReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the NerifBridge contract.
type NerifBridgeReceiveIterator struct {
	Event *NerifBridgeReceive // Event containing the contract specifics and raw log

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
func (it *NerifBridgeReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifBridgeReceive)
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
		it.Event = new(NerifBridgeReceive)
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
func (it *NerifBridgeReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifBridgeReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifBridgeReceive represents a Receive event raised by the NerifBridge contract.
type NerifBridgeReceive struct {
	Target    common.Address
	Payload   []byte
	GasAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0x29e7ed49362e97d1632e595994a7586fdf3953cc094a3ea2c4f8a456ea186ef4.
//
// Solidity: event Receive(address target, bytes payload, uint256 gasAmount)
func (_NerifBridge *NerifBridgeFilterer) FilterReceive(opts *bind.FilterOpts) (*NerifBridgeReceiveIterator, error) {

	logs, sub, err := _NerifBridge.contract.FilterLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return &NerifBridgeReceiveIterator{contract: _NerifBridge.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0x29e7ed49362e97d1632e595994a7586fdf3953cc094a3ea2c4f8a456ea186ef4.
//
// Solidity: event Receive(address target, bytes payload, uint256 gasAmount)
func (_NerifBridge *NerifBridgeFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *NerifBridgeReceive) (event.Subscription, error) {

	logs, sub, err := _NerifBridge.contract.WatchLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifBridgeReceive)
				if err := _NerifBridge.contract.UnpackLog(event, "Receive", log); err != nil {
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

// ParseReceive is a log parse operation binding the contract event 0x29e7ed49362e97d1632e595994a7586fdf3953cc094a3ea2c4f8a456ea186ef4.
//
// Solidity: event Receive(address target, bytes payload, uint256 gasAmount)
func (_NerifBridge *NerifBridgeFilterer) ParseReceive(log types.Log) (*NerifBridgeReceive, error) {
	event := new(NerifBridgeReceive)
	if err := _NerifBridge.contract.UnpackLog(event, "Receive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifBridgeSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the NerifBridge contract.
type NerifBridgeSendIterator struct {
	Event *NerifBridgeSend // Event containing the contract specifics and raw log

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
func (it *NerifBridgeSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifBridgeSend)
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
		it.Event = new(NerifBridgeSend)
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
func (it *NerifBridgeSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifBridgeSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifBridgeSend represents a Send event raised by the NerifBridge contract.
type NerifBridgeSend struct {
	ChainId   *big.Int
	Target    common.Address
	Payload   []byte
	GasAmount *big.Int
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0xc97d87de2251935700defa2c3bee343060e3a803d01b58399698b836deef2829.
//
// Solidity: event Send(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender)
func (_NerifBridge *NerifBridgeFilterer) FilterSend(opts *bind.FilterOpts) (*NerifBridgeSendIterator, error) {

	logs, sub, err := _NerifBridge.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &NerifBridgeSendIterator{contract: _NerifBridge.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0xc97d87de2251935700defa2c3bee343060e3a803d01b58399698b836deef2829.
//
// Solidity: event Send(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender)
func (_NerifBridge *NerifBridgeFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *NerifBridgeSend) (event.Subscription, error) {

	logs, sub, err := _NerifBridge.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifBridgeSend)
				if err := _NerifBridge.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0xc97d87de2251935700defa2c3bee343060e3a803d01b58399698b836deef2829.
//
// Solidity: event Send(uint256 chainId, address target, bytes payload, uint256 gasAmount, address sender)
func (_NerifBridge *NerifBridgeFilterer) ParseSend(log types.Log) (*NerifBridgeSend, error) {
	event := new(NerifBridgeSend)
	if err := _NerifBridge.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
