// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// INerifBridgeReceiverMetaData contains all meta data concerning the INerifBridgeReceiver contract.
var INerifBridgeReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nbReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INerifBridgeReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use INerifBridgeReceiverMetaData.ABI instead.
var INerifBridgeReceiverABI = INerifBridgeReceiverMetaData.ABI

// INerifBridgeReceiver is an auto generated Go binding around an Ethereum contract.
type INerifBridgeReceiver struct {
	INerifBridgeReceiverCaller     // Read-only binding to the contract
	INerifBridgeReceiverTransactor // Write-only binding to the contract
	INerifBridgeReceiverFilterer   // Log filterer for contract events
}

// INerifBridgeReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type INerifBridgeReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INerifBridgeReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INerifBridgeReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INerifBridgeReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INerifBridgeReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INerifBridgeReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INerifBridgeReceiverSession struct {
	Contract     *INerifBridgeReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// INerifBridgeReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INerifBridgeReceiverCallerSession struct {
	Contract *INerifBridgeReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// INerifBridgeReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INerifBridgeReceiverTransactorSession struct {
	Contract     *INerifBridgeReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// INerifBridgeReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type INerifBridgeReceiverRaw struct {
	Contract *INerifBridgeReceiver // Generic contract binding to access the raw methods on
}

// INerifBridgeReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INerifBridgeReceiverCallerRaw struct {
	Contract *INerifBridgeReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// INerifBridgeReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INerifBridgeReceiverTransactorRaw struct {
	Contract *INerifBridgeReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINerifBridgeReceiver creates a new instance of INerifBridgeReceiver, bound to a specific deployed contract.
func NewINerifBridgeReceiver(address common.Address, backend bind.ContractBackend) (*INerifBridgeReceiver, error) {
	contract, err := bindINerifBridgeReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INerifBridgeReceiver{INerifBridgeReceiverCaller: INerifBridgeReceiverCaller{contract: contract}, INerifBridgeReceiverTransactor: INerifBridgeReceiverTransactor{contract: contract}, INerifBridgeReceiverFilterer: INerifBridgeReceiverFilterer{contract: contract}}, nil
}

// NewINerifBridgeReceiverCaller creates a new read-only instance of INerifBridgeReceiver, bound to a specific deployed contract.
func NewINerifBridgeReceiverCaller(address common.Address, caller bind.ContractCaller) (*INerifBridgeReceiverCaller, error) {
	contract, err := bindINerifBridgeReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INerifBridgeReceiverCaller{contract: contract}, nil
}

// NewINerifBridgeReceiverTransactor creates a new write-only instance of INerifBridgeReceiver, bound to a specific deployed contract.
func NewINerifBridgeReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*INerifBridgeReceiverTransactor, error) {
	contract, err := bindINerifBridgeReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INerifBridgeReceiverTransactor{contract: contract}, nil
}

// NewINerifBridgeReceiverFilterer creates a new log filterer instance of INerifBridgeReceiver, bound to a specific deployed contract.
func NewINerifBridgeReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*INerifBridgeReceiverFilterer, error) {
	contract, err := bindINerifBridgeReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INerifBridgeReceiverFilterer{contract: contract}, nil
}

// bindINerifBridgeReceiver binds a generic wrapper to an already deployed contract.
func bindINerifBridgeReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INerifBridgeReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INerifBridgeReceiver *INerifBridgeReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INerifBridgeReceiver.Contract.INerifBridgeReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INerifBridgeReceiver *INerifBridgeReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INerifBridgeReceiver.Contract.INerifBridgeReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INerifBridgeReceiver *INerifBridgeReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INerifBridgeReceiver.Contract.INerifBridgeReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INerifBridgeReceiver *INerifBridgeReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INerifBridgeReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INerifBridgeReceiver *INerifBridgeReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INerifBridgeReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INerifBridgeReceiver *INerifBridgeReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INerifBridgeReceiver.Contract.contract.Transact(opts, method, params...)
}

// NbReceive is a paid mutator transaction binding the contract method 0x38e7d04a.
//
// Solidity: function nbReceive(uint256 chainId, address sender, bytes payload) returns()
func (_INerifBridgeReceiver *INerifBridgeReceiverTransactor) NbReceive(opts *bind.TransactOpts, chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _INerifBridgeReceiver.contract.Transact(opts, "nbReceive", chainId, sender, payload)
}

// NbReceive is a paid mutator transaction binding the contract method 0x38e7d04a.
//
// Solidity: function nbReceive(uint256 chainId, address sender, bytes payload) returns()
func (_INerifBridgeReceiver *INerifBridgeReceiverSession) NbReceive(chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _INerifBridgeReceiver.Contract.NbReceive(&_INerifBridgeReceiver.TransactOpts, chainId, sender, payload)
}

// NbReceive is a paid mutator transaction binding the contract method 0x38e7d04a.
//
// Solidity: function nbReceive(uint256 chainId, address sender, bytes payload) returns()
func (_INerifBridgeReceiver *INerifBridgeReceiverTransactorSession) NbReceive(chainId *big.Int, sender common.Address, payload []byte) (*types.Transaction, error) {
	return _INerifBridgeReceiver.Contract.NbReceive(&_INerifBridgeReceiver.TransactOpts, chainId, sender, payload)
}
