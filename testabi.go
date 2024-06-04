// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// WriteY is an auto generated low-level Go binding around an user-defined struct.
type WriteY struct {
	Y1 []byte
	Y2 []byte
	Y3 []byte
	Y4 []byte
	Y5 []byte
}

// WritepiFsCk is an auto generated low-level Go binding around an user-defined struct.
type WritepiFsCk struct {
	Ck1 *big.Int
	Ck2 *big.Int
	Ck3 *big.Int
}

// WritepiFsG1 is an auto generated low-level Go binding around an user-defined struct.
type WritepiFsG1 struct {
	Pi1      []byte
	Pi2      []byte
	Pi2RmNeg []byte
}

// WritepiFsG2 is an auto generated low-level Go binding around an user-defined struct.
type WritepiFsG2 struct {
	G2  []byte
	XA  []byte
	YA1 []byte
	YA2 []byte
}

// WritepiOkCk is an auto generated low-level Go binding around an user-defined struct.
type WritepiOkCk struct {
	Ck1 *big.Int
	Ck2 *big.Int
	Ck3 *big.Int
	Ck4 *big.Int
}

// WritepiOkTt is an auto generated low-level Go binding around an user-defined struct.
type WritepiOkTt struct {
	E1  []byte
	Ee1 []byte
	E2  []byte
	Ee2 []byte
	Tl  []byte
	Ttl []byte
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"x\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"y\",\"type\":\"bytes\"}],\"name\":\"G1_ADD\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"x\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"G1_MUL\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"x1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"x2\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"rM\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"y1\",\"type\":\"bytes\"}],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"p1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c2\",\"type\":\"uint256\"}],\"name\":\"verify_pi_fj\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pi_1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pi_2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pi_2_rm_neg\",\"type\":\"bytes\"}],\"internalType\":\"structwrite.pi_fs_g1\",\"name\":\"g1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"g2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"X_A\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Y_A_1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Y_A_2\",\"type\":\"bytes\"}],\"internalType\":\"structwrite.pi_fs_g2\",\"name\":\"g2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ck1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ck2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ck3\",\"type\":\"uint256\"}],\"internalType\":\"structwrite.pi_fs_ck\",\"name\":\"ck\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"m_f\",\"type\":\"bytes\"}],\"name\":\"verify_pi_fs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"Y1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Y2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Y3\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Y4\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"Y5\",\"type\":\"bytes\"}],\"internalType\":\"structwrite.Y\",\"name\":\"y\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"e1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ee1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"e2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ee2\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"tl\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"ttl\",\"type\":\"bytes\"}],\"internalType\":\"structwrite.pi_ok_tt\",\"name\":\"tt\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"m_ok\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ck1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ck2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ck3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ck4\",\"type\":\"uint256\"}],\"internalType\":\"structwrite.pi_ok_ck\",\"name\":\"ck\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"g1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pks\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pkt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"gc\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"h3_idf\",\"type\":\"bytes\"}],\"name\":\"verify_pi_ok\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// G1ADD is a free data retrieval call binding the contract method 0xe9d7fda7.
//
// Solidity: function G1_ADD(bytes x, bytes y) view returns(bytes)
func (_Test *TestCaller) G1ADD(opts *bind.CallOpts, x []byte, y []byte) ([]byte, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "G1_ADD", x, y)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// G1ADD is a free data retrieval call binding the contract method 0xe9d7fda7.
//
// Solidity: function G1_ADD(bytes x, bytes y) view returns(bytes)
func (_Test *TestSession) G1ADD(x []byte, y []byte) ([]byte, error) {
	return _Test.Contract.G1ADD(&_Test.CallOpts, x, y)
}

// G1ADD is a free data retrieval call binding the contract method 0xe9d7fda7.
//
// Solidity: function G1_ADD(bytes x, bytes y) view returns(bytes)
func (_Test *TestCallerSession) G1ADD(x []byte, y []byte) ([]byte, error) {
	return _Test.Contract.G1ADD(&_Test.CallOpts, x, y)
}

// G1MUL is a free data retrieval call binding the contract method 0x0b59ae88.
//
// Solidity: function G1_MUL(bytes x, uint256 k) view returns(bytes)
func (_Test *TestCaller) G1MUL(opts *bind.CallOpts, x []byte, k *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "G1_MUL", x, k)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// G1MUL is a free data retrieval call binding the contract method 0x0b59ae88.
//
// Solidity: function G1_MUL(bytes x, uint256 k) view returns(bytes)
func (_Test *TestSession) G1MUL(x []byte, k *big.Int) ([]byte, error) {
	return _Test.Contract.G1MUL(&_Test.CallOpts, x, k)
}

// G1MUL is a free data retrieval call binding the contract method 0x0b59ae88.
//
// Solidity: function G1_MUL(bytes x, uint256 k) view returns(bytes)
func (_Test *TestCallerSession) G1MUL(x []byte, k *big.Int) ([]byte, error) {
	return _Test.Contract.G1MUL(&_Test.CallOpts, x, k)
}

// GetCode is a free data retrieval call binding the contract method 0xea879634.
//
// Solidity: function getCode() view returns(uint256)
func (_Test *TestCaller) GetCode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "getCode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCode is a free data retrieval call binding the contract method 0xea879634.
//
// Solidity: function getCode() view returns(uint256)
func (_Test *TestSession) GetCode() (*big.Int, error) {
	return _Test.Contract.GetCode(&_Test.CallOpts)
}

// GetCode is a free data retrieval call binding the contract method 0xea879634.
//
// Solidity: function getCode() view returns(uint256)
func (_Test *TestCallerSession) GetCode() (*big.Int, error) {
	return _Test.Contract.GetCode(&_Test.CallOpts)
}

// Test is a paid mutator transaction binding the contract method 0xae8cfb9b.
//
// Solidity: function test(bytes x1, bytes x2, uint256 rM, bytes y1) returns()
func (_Test *TestTransactor) Test(opts *bind.TransactOpts, x1 []byte, x2 []byte, rM *big.Int, y1 []byte) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "test", x1, x2, rM, y1)
}

// Test is a paid mutator transaction binding the contract method 0xae8cfb9b.
//
// Solidity: function test(bytes x1, bytes x2, uint256 rM, bytes y1) returns()
func (_Test *TestSession) Test(x1 []byte, x2 []byte, rM *big.Int, y1 []byte) (*types.Transaction, error) {
	return _Test.Contract.Test(&_Test.TransactOpts, x1, x2, rM, y1)
}

// Test is a paid mutator transaction binding the contract method 0xae8cfb9b.
//
// Solidity: function test(bytes x1, bytes x2, uint256 rM, bytes y1) returns()
func (_Test *TestTransactorSession) Test(x1 []byte, x2 []byte, rM *big.Int, y1 []byte) (*types.Transaction, error) {
	return _Test.Contract.Test(&_Test.TransactOpts, x1, x2, rM, y1)
}

// VerifyPiFj is a paid mutator transaction binding the contract method 0x3cdb26c8.
//
// Solidity: function verify_pi_fj(uint256 p1, uint256 p2, uint256 p3, uint256 c1, uint256 c2) returns()
func (_Test *TestTransactor) VerifyPiFj(opts *bind.TransactOpts, p1 *big.Int, p2 *big.Int, p3 *big.Int, c1 *big.Int, c2 *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "verify_pi_fj", p1, p2, p3, c1, c2)
}

// VerifyPiFj is a paid mutator transaction binding the contract method 0x3cdb26c8.
//
// Solidity: function verify_pi_fj(uint256 p1, uint256 p2, uint256 p3, uint256 c1, uint256 c2) returns()
func (_Test *TestSession) VerifyPiFj(p1 *big.Int, p2 *big.Int, p3 *big.Int, c1 *big.Int, c2 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.VerifyPiFj(&_Test.TransactOpts, p1, p2, p3, c1, c2)
}

// VerifyPiFj is a paid mutator transaction binding the contract method 0x3cdb26c8.
//
// Solidity: function verify_pi_fj(uint256 p1, uint256 p2, uint256 p3, uint256 c1, uint256 c2) returns()
func (_Test *TestTransactorSession) VerifyPiFj(p1 *big.Int, p2 *big.Int, p3 *big.Int, c1 *big.Int, c2 *big.Int) (*types.Transaction, error) {
	return _Test.Contract.VerifyPiFj(&_Test.TransactOpts, p1, p2, p3, c1, c2)
}

// VerifyPiFs is a paid mutator transaction binding the contract method 0xd365be03.
//
// Solidity: function verify_pi_fs((bytes,bytes,bytes) g1, (bytes,bytes,bytes,bytes) g2, (uint256,uint256,uint256) ck, bytes m_f) returns()
func (_Test *TestTransactor) VerifyPiFs(opts *bind.TransactOpts, g1 WritepiFsG1, g2 WritepiFsG2, ck WritepiFsCk, m_f []byte) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "verify_pi_fs", g1, g2, ck, m_f)
}

// VerifyPiFs is a paid mutator transaction binding the contract method 0xd365be03.
//
// Solidity: function verify_pi_fs((bytes,bytes,bytes) g1, (bytes,bytes,bytes,bytes) g2, (uint256,uint256,uint256) ck, bytes m_f) returns()
func (_Test *TestSession) VerifyPiFs(g1 WritepiFsG1, g2 WritepiFsG2, ck WritepiFsCk, m_f []byte) (*types.Transaction, error) {
	return _Test.Contract.VerifyPiFs(&_Test.TransactOpts, g1, g2, ck, m_f)
}

// VerifyPiFs is a paid mutator transaction binding the contract method 0xd365be03.
//
// Solidity: function verify_pi_fs((bytes,bytes,bytes) g1, (bytes,bytes,bytes,bytes) g2, (uint256,uint256,uint256) ck, bytes m_f) returns()
func (_Test *TestTransactorSession) VerifyPiFs(g1 WritepiFsG1, g2 WritepiFsG2, ck WritepiFsCk, m_f []byte) (*types.Transaction, error) {
	return _Test.Contract.VerifyPiFs(&_Test.TransactOpts, g1, g2, ck, m_f)
}

// VerifyPiOk is a paid mutator transaction binding the contract method 0x41098092.
//
// Solidity: function verify_pi_ok((bytes,bytes,bytes,bytes,bytes) y, (bytes,bytes,bytes,bytes,bytes,bytes) tt, bytes m_ok, (uint256,uint256,uint256,uint256) ck, bytes g1, bytes pks, bytes pkt, bytes gc, bytes h3_idf) returns()
func (_Test *TestTransactor) VerifyPiOk(opts *bind.TransactOpts, y WriteY, tt WritepiOkTt, m_ok []byte, ck WritepiOkCk, g1 []byte, pks []byte, pkt []byte, gc []byte, h3_idf []byte) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "verify_pi_ok", y, tt, m_ok, ck, g1, pks, pkt, gc, h3_idf)
}

// VerifyPiOk is a paid mutator transaction binding the contract method 0x41098092.
//
// Solidity: function verify_pi_ok((bytes,bytes,bytes,bytes,bytes) y, (bytes,bytes,bytes,bytes,bytes,bytes) tt, bytes m_ok, (uint256,uint256,uint256,uint256) ck, bytes g1, bytes pks, bytes pkt, bytes gc, bytes h3_idf) returns()
func (_Test *TestSession) VerifyPiOk(y WriteY, tt WritepiOkTt, m_ok []byte, ck WritepiOkCk, g1 []byte, pks []byte, pkt []byte, gc []byte, h3_idf []byte) (*types.Transaction, error) {
	return _Test.Contract.VerifyPiOk(&_Test.TransactOpts, y, tt, m_ok, ck, g1, pks, pkt, gc, h3_idf)
}

// VerifyPiOk is a paid mutator transaction binding the contract method 0x41098092.
//
// Solidity: function verify_pi_ok((bytes,bytes,bytes,bytes,bytes) y, (bytes,bytes,bytes,bytes,bytes,bytes) tt, bytes m_ok, (uint256,uint256,uint256,uint256) ck, bytes g1, bytes pks, bytes pkt, bytes gc, bytes h3_idf) returns()
func (_Test *TestTransactorSession) VerifyPiOk(y WriteY, tt WritepiOkTt, m_ok []byte, ck WritepiOkCk, g1 []byte, pks []byte, pkt []byte, gc []byte, h3_idf []byte) (*types.Transaction, error) {
	return _Test.Contract.VerifyPiOk(&_Test.TransactOpts, y, tt, m_ok, ck, g1, pks, pkt, gc, h3_idf)
}
