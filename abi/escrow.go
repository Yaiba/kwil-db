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
)

// EscrowMetaData contains all meta data concerning the Escrow contract.
var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_escrowToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Return\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrowToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"returnDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowMetaData.ABI instead.
var EscrowABI = EscrowMetaData.ABI

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// EscrowToken is a free data retrieval call binding the contract method 0x2fe319da.
//
// Solidity: function escrowToken() view returns(address)
func (_Escrow *EscrowCaller) EscrowToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "escrowToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowToken is a free data retrieval call binding the contract method 0x2fe319da.
//
// Solidity: function escrowToken() view returns(address)
func (_Escrow *EscrowSession) EscrowToken() (common.Address, error) {
	return _Escrow.Contract.EscrowToken(&_Escrow.CallOpts)
}

// EscrowToken is a free data retrieval call binding the contract method 0x2fe319da.
//
// Solidity: function escrowToken() view returns(address)
func (_Escrow *EscrowCallerSession) EscrowToken() (common.Address, error) {
	return _Escrow.Contract.EscrowToken(&_Escrow.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0x901754d7.
//
// Solidity: function pools(address , address ) view returns(uint256)
func (_Escrow *EscrowCaller) Pools(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "pools", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pools is a free data retrieval call binding the contract method 0x901754d7.
//
// Solidity: function pools(address , address ) view returns(uint256)
func (_Escrow *EscrowSession) Pools(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Escrow.Contract.Pools(&_Escrow.CallOpts, arg0, arg1)
}

// Pools is a free data retrieval call binding the contract method 0x901754d7.
//
// Solidity: function pools(address , address ) view returns(uint256)
func (_Escrow *EscrowCallerSession) Pools(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Escrow.Contract.Pools(&_Escrow.CallOpts, arg0, arg1)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address validator, uint256 amt) payable returns()
func (_Escrow *EscrowTransactor) Deposit(opts *bind.TransactOpts, validator common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "deposit", validator, amt)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address validator, uint256 amt) payable returns()
func (_Escrow *EscrowSession) Deposit(validator common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts, validator, amt)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address validator, uint256 amt) payable returns()
func (_Escrow *EscrowTransactorSession) Deposit(validator common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts, validator, amt)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0x24d20ef4.
//
// Solidity: function returnDeposit(address recipient, uint256 amt, uint256 fee) returns()
func (_Escrow *EscrowTransactor) ReturnDeposit(opts *bind.TransactOpts, recipient common.Address, amt *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "returnDeposit", recipient, amt, fee)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0x24d20ef4.
//
// Solidity: function returnDeposit(address recipient, uint256 amt, uint256 fee) returns()
func (_Escrow *EscrowSession) ReturnDeposit(recipient common.Address, amt *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.ReturnDeposit(&_Escrow.TransactOpts, recipient, amt, fee)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0x24d20ef4.
//
// Solidity: function returnDeposit(address recipient, uint256 amt, uint256 fee) returns()
func (_Escrow *EscrowTransactorSession) ReturnDeposit(recipient common.Address, amt *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.ReturnDeposit(&_Escrow.TransactOpts, recipient, amt, fee)
}

// EscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Escrow contract.
type EscrowDepositedIterator struct {
	Event *EscrowDeposited // Event containing the contract specifics and raw log

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
func (it *EscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowDeposited)
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
		it.Event = new(EscrowDeposited)
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
func (it *EscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowDeposited represents a Deposited event raised by the Escrow contract.
type EscrowDeposited struct {
	Sender    common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address sender, address validator, uint256 amount)
func (_Escrow *EscrowFilterer) FilterDeposited(opts *bind.FilterOpts) (*EscrowDepositedIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &EscrowDepositedIterator{contract: _Escrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address sender, address validator, uint256 amount)
func (_Escrow *EscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EscrowDeposited) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowDeposited)
				if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address sender, address validator, uint256 amount)
func (_Escrow *EscrowFilterer) ParseDeposited(log types.Log) (*EscrowDeposited, error) {
	event := new(EscrowDeposited)
	if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowReturnIterator is returned from FilterReturn and is used to iterate over the raw logs and unpacked data for Return events raised by the Escrow contract.
type EscrowReturnIterator struct {
	Event *EscrowReturn // Event containing the contract specifics and raw log

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
func (it *EscrowReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowReturn)
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
		it.Event = new(EscrowReturn)
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
func (it *EscrowReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowReturn represents a Return event raised by the Escrow contract.
type EscrowReturn struct {
	Receiver  common.Address
	Validator common.Address
	Amount    *big.Int
	Fee       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReturn is a free log retrieval operation binding the contract event 0xadf247d5d167e5c6c708ae45d0beb8c3a670079b1a4097a05c8844964ff0a9ee.
//
// Solidity: event Return(address receiver, address validator, uint256 amount, uint256 fee)
func (_Escrow *EscrowFilterer) FilterReturn(opts *bind.FilterOpts) (*EscrowReturnIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Return")
	if err != nil {
		return nil, err
	}
	return &EscrowReturnIterator{contract: _Escrow.contract, event: "Return", logs: logs, sub: sub}, nil
}

// WatchReturn is a free log subscription operation binding the contract event 0xadf247d5d167e5c6c708ae45d0beb8c3a670079b1a4097a05c8844964ff0a9ee.
//
// Solidity: event Return(address receiver, address validator, uint256 amount, uint256 fee)
func (_Escrow *EscrowFilterer) WatchReturn(opts *bind.WatchOpts, sink chan<- *EscrowReturn) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Return")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowReturn)
				if err := _Escrow.contract.UnpackLog(event, "Return", log); err != nil {
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

// ParseReturn is a log parse operation binding the contract event 0xadf247d5d167e5c6c708ae45d0beb8c3a670079b1a4097a05c8844964ff0a9ee.
//
// Solidity: event Return(address receiver, address validator, uint256 amount, uint256 fee)
func (_Escrow *EscrowFilterer) ParseReturn(log types.Log) (*EscrowReturn, error) {
	event := new(EscrowReturn)
	if err := _Escrow.contract.UnpackLog(event, "Return", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
