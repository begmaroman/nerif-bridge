// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

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

// TestReceiverMetaData contains all meta data concerning the TestReceiver contract.
var TestReceiverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"MsgReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nbReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use TestReceiverMetaData.ABI instead.
var TestReceiverABI = TestReceiverMetaData.ABI

// TestReceiver is an auto generated Go binding around an Ethereum contract.
type TestReceiver struct {
	TestReceiverCaller     // Read-only binding to the contract
	TestReceiverTransactor // Write-only binding to the contract
	TestReceiverFilterer   // Log filterer for contract events
}

// TestReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestReceiverSession struct {
	Contract     *TestReceiver     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestReceiverCallerSession struct {
	Contract *TestReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestReceiverTransactorSession struct {
	Contract     *TestReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestReceiverRaw struct {
	Contract *TestReceiver // Generic contract binding to access the raw methods on
}

// TestReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestReceiverCallerRaw struct {
	Contract *TestReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// TestReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestReceiverTransactorRaw struct {
	Contract *TestReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestReceiver creates a new instance of TestReceiver, bound to a specific deployed contract.
func NewTestReceiver(address common.Address, backend bind.ContractBackend) (*TestReceiver, error) {
	contract, err := bindTestReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestReceiver{TestReceiverCaller: TestReceiverCaller{contract: contract}, TestReceiverTransactor: TestReceiverTransactor{contract: contract}, TestReceiverFilterer: TestReceiverFilterer{contract: contract}}, nil
}

// NewTestReceiverCaller creates a new read-only instance of TestReceiver, bound to a specific deployed contract.
func NewTestReceiverCaller(address common.Address, caller bind.ContractCaller) (*TestReceiverCaller, error) {
	contract, err := bindTestReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestReceiverCaller{contract: contract}, nil
}

// NewTestReceiverTransactor creates a new write-only instance of TestReceiver, bound to a specific deployed contract.
func NewTestReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*TestReceiverTransactor, error) {
	contract, err := bindTestReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestReceiverTransactor{contract: contract}, nil
}

// NewTestReceiverFilterer creates a new log filterer instance of TestReceiver, bound to a specific deployed contract.
func NewTestReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*TestReceiverFilterer, error) {
	contract, err := bindTestReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestReceiverFilterer{contract: contract}, nil
}

// bindTestReceiver binds a generic wrapper to an already deployed contract.
func bindTestReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestReceiver *TestReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestReceiver.Contract.TestReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestReceiver *TestReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestReceiver.Contract.TestReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestReceiver *TestReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestReceiver.Contract.TestReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestReceiver *TestReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestReceiver *TestReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestReceiver *TestReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestReceiver.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_TestReceiver *TestReceiverCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestReceiver.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_TestReceiver *TestReceiverSession) Bridge() (common.Address, error) {
	return _TestReceiver.Contract.Bridge(&_TestReceiver.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_TestReceiver *TestReceiverCallerSession) Bridge() (common.Address, error) {
	return _TestReceiver.Contract.Bridge(&_TestReceiver.CallOpts)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 chainId, address sender, bytes payload)
func (_TestReceiver *TestReceiverCaller) Messages(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainId *big.Int
	Sender  common.Address
	Payload []byte
}, error) {
	var out []interface{}
	err := _TestReceiver.contract.Call(opts, &out, "messages", arg0)

	outstruct := new(struct {
		ChainId *big.Int
		Sender  common.Address
		Payload []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Payload = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 chainId, address sender, bytes payload)
func (_TestReceiver *TestReceiverSession) Messages(arg0 *big.Int) (struct {
	ChainId *big.Int
	Sender  common.Address
	Payload []byte
}, error) {
	return _TestReceiver.Contract.Messages(&_TestReceiver.CallOpts, arg0)
}

// Messages is a free data retrieval call binding the contract method 0x0d80fefd.
//
// Solidity: function messages(uint256 ) view returns(uint256 chainId, address sender, bytes payload)
func (_TestReceiver *TestReceiverCallerSession) Messages(arg0 *big.Int) (struct {
	ChainId *big.Int
	Sender  common.Address
	Payload []byte
}, error) {
	return _TestReceiver.Contract.Messages(&_TestReceiver.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestReceiver *TestReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestReceiver *TestReceiverSession) Owner() (common.Address, error) {
	return _TestReceiver.Contract.Owner(&_TestReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestReceiver *TestReceiverCallerSession) Owner() (common.Address, error) {
	return _TestReceiver.Contract.Owner(&_TestReceiver.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_TestReceiver *TestReceiverTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _TestReceiver.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_TestReceiver *TestReceiverSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _TestReceiver.Contract.Initialize(&_TestReceiver.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_TestReceiver *TestReceiverTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _TestReceiver.Contract.Initialize(&_TestReceiver.TransactOpts, _bridge)
}

// NbReceive is a paid mutator transaction binding the contract method 0x38e7d04a.
//
// Solidity: function nbReceive(uint256 chainId, address sender, bytes payload) returns()
func (_TestReceiver *TestReceiverTransactor) NbReceive(opts *bind.TransactOpts, chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _TestReceiver.contract.Transact(opts, "nbReceive", chainId, sender, payload)
}

// NbReceive is a paid mutator transaction binding the contract method 0x38e7d04a.
//
// Solidity: function nbReceive(uint256 chainId, address sender, bytes payload) returns()
func (_TestReceiver *TestReceiverSession) NbReceive(chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _TestReceiver.Contract.NbReceive(&_TestReceiver.TransactOpts, chainId, sender, payload)
}

// NbReceive is a paid mutator transaction binding the contract method 0x38e7d04a.
//
// Solidity: function nbReceive(uint256 chainId, address sender, bytes payload) returns()
func (_TestReceiver *TestReceiverTransactorSession) NbReceive(chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _TestReceiver.Contract.NbReceive(&_TestReceiver.TransactOpts, chainId, sender, payload)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestReceiver *TestReceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestReceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestReceiver *TestReceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestReceiver.Contract.RenounceOwnership(&_TestReceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TestReceiver *TestReceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TestReceiver.Contract.RenounceOwnership(&_TestReceiver.TransactOpts)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_TestReceiver *TestReceiverTransactor) SetBridge(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _TestReceiver.contract.Transact(opts, "setBridge", _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_TestReceiver *TestReceiverSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _TestReceiver.Contract.SetBridge(&_TestReceiver.TransactOpts, _bridge)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address _bridge) returns()
func (_TestReceiver *TestReceiverTransactorSession) SetBridge(_bridge common.Address) (*types.Transaction, error) {
	return _TestReceiver.Contract.SetBridge(&_TestReceiver.TransactOpts, _bridge)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestReceiver *TestReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestReceiver *TestReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestReceiver.Contract.TransferOwnership(&_TestReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestReceiver *TestReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestReceiver.Contract.TransferOwnership(&_TestReceiver.TransactOpts, newOwner)
}

// TestReceiverInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TestReceiver contract.
type TestReceiverInitializedIterator struct {
	Event *TestReceiverInitialized // Event containing the contract specifics and raw log

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
func (it *TestReceiverInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestReceiverInitialized)
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
		it.Event = new(TestReceiverInitialized)
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
func (it *TestReceiverInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestReceiverInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestReceiverInitialized represents a Initialized event raised by the TestReceiver contract.
type TestReceiverInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TestReceiver *TestReceiverFilterer) FilterInitialized(opts *bind.FilterOpts) (*TestReceiverInitializedIterator, error) {

	logs, sub, err := _TestReceiver.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TestReceiverInitializedIterator{contract: _TestReceiver.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TestReceiver *TestReceiverFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TestReceiverInitialized) (event.Subscription, error) {

	logs, sub, err := _TestReceiver.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestReceiverInitialized)
				if err := _TestReceiver.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TestReceiver *TestReceiverFilterer) ParseInitialized(log types.Log) (*TestReceiverInitialized, error) {
	event := new(TestReceiverInitialized)
	if err := _TestReceiver.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestReceiverMsgReceivedIterator is returned from FilterMsgReceived and is used to iterate over the raw logs and unpacked data for MsgReceived events raised by the TestReceiver contract.
type TestReceiverMsgReceivedIterator struct {
	Event *TestReceiverMsgReceived // Event containing the contract specifics and raw log

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
func (it *TestReceiverMsgReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestReceiverMsgReceived)
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
		it.Event = new(TestReceiverMsgReceived)
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
func (it *TestReceiverMsgReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestReceiverMsgReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestReceiverMsgReceived represents a MsgReceived event raised by the TestReceiver contract.
type TestReceiverMsgReceived struct {
	ChainId *big.Int
	Sender  common.Address
	Payload []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMsgReceived is a free log retrieval operation binding the contract event 0xf53a6540776dc100ae0749185c1ae0b58423f148ecde0078f16869565af26419.
//
// Solidity: event MsgReceived(uint256 chainId, address sender, bytes payload)
func (_TestReceiver *TestReceiverFilterer) FilterMsgReceived(opts *bind.FilterOpts) (*TestReceiverMsgReceivedIterator, error) {

	logs, sub, err := _TestReceiver.contract.FilterLogs(opts, "MsgReceived")
	if err != nil {
		return nil, err
	}
	return &TestReceiverMsgReceivedIterator{contract: _TestReceiver.contract, event: "MsgReceived", logs: logs, sub: sub}, nil
}

// WatchMsgReceived is a free log subscription operation binding the contract event 0xf53a6540776dc100ae0749185c1ae0b58423f148ecde0078f16869565af26419.
//
// Solidity: event MsgReceived(uint256 chainId, address sender, bytes payload)
func (_TestReceiver *TestReceiverFilterer) WatchMsgReceived(opts *bind.WatchOpts, sink chan<- *TestReceiverMsgReceived) (event.Subscription, error) {

	logs, sub, err := _TestReceiver.contract.WatchLogs(opts, "MsgReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestReceiverMsgReceived)
				if err := _TestReceiver.contract.UnpackLog(event, "MsgReceived", log); err != nil {
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

// ParseMsgReceived is a log parse operation binding the contract event 0xf53a6540776dc100ae0749185c1ae0b58423f148ecde0078f16869565af26419.
//
// Solidity: event MsgReceived(uint256 chainId, address sender, bytes payload)
func (_TestReceiver *TestReceiverFilterer) ParseMsgReceived(log types.Log) (*TestReceiverMsgReceived, error) {
	event := new(TestReceiverMsgReceived)
	if err := _TestReceiver.contract.UnpackLog(event, "MsgReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestReceiver contract.
type TestReceiverOwnershipTransferredIterator struct {
	Event *TestReceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestReceiverOwnershipTransferred)
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
		it.Event = new(TestReceiverOwnershipTransferred)
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
func (it *TestReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the TestReceiver contract.
type TestReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestReceiver *TestReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestReceiverOwnershipTransferredIterator{contract: _TestReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestReceiver *TestReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestReceiverOwnershipTransferred)
				if err := _TestReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestReceiver *TestReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*TestReceiverOwnershipTransferred, error) {
	event := new(TestReceiverOwnershipTransferred)
	if err := _TestReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
