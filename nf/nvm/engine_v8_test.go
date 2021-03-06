// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package nvm

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/nebulasio/go-nebulas/consensus/dpos"

	"encoding/json"

	"github.com/nebulasio/go-nebulas/core"
	"github.com/nebulasio/go-nebulas/core/state"
	"github.com/nebulasio/go-nebulas/crypto"
	"github.com/nebulasio/go-nebulas/crypto/keystore"
	"github.com/nebulasio/go-nebulas/crypto/keystore/secp256k1"
	"github.com/nebulasio/go-nebulas/storage"
	"github.com/nebulasio/go-nebulas/util"
	"github.com/nebulasio/go-nebulas/util/byteutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newUint128FromIntWrapper(a int64) *util.Uint128 {
	b, _ := util.NewUint128FromInt(a)
	return b
}

type testBlock struct {
}

// Coinbase mock
func (block *testBlock) Coinbase() *core.Address {
	addr, _ := core.AddressParse("n1FkntVUMPAsESuCAAPK711omQk19JotBjM")
	return addr
}

// Hash mock
func (block *testBlock) Hash() byteutils.Hash {
	return []byte("59fc526072b09af8a8ca9732dae17132c4e9127e43cf2232")
}

// Height mock
func (block *testBlock) Height() uint64 {
	return 1
}

// GetTransaction mock
func (block *testBlock) GetTransaction(hash byteutils.Hash) (*core.Transaction, error) {
	return nil, nil
}

// RecordEvent mock
func (block *testBlock) RecordEvent(txHash byteutils.Hash, topic, data string) error {
	return nil
}

func (block *testBlock) Timestamp() int64 {
	return int64(0)
}

func mockBlock() Block {
	block := &testBlock{}
	return block
}

func mockTransaction() *core.Transaction {
	return mockNormalTransaction("n1FkntVUMPAsESuCAAPK711omQk19JotBjM", "n1JNHZJEUvfBYfjDRD14Q73FX62nJAzXkMR", "0")
}

const ContractName = "contract.js"

func mockNormalTransaction(from, to, value string) *core.Transaction {

	fromAddr, _ := core.AddressParse(from)
	toAddr, _ := core.AddressParse(to)
	payload, _ := core.NewBinaryPayload(nil).ToBytes()
	gasPrice, _ := util.NewUint128FromString("1000000")
	gasLimit, _ := util.NewUint128FromString("2000000")
	v, _ := util.NewUint128FromString(value)
	tx, _ := core.NewTransaction(1, fromAddr, toAddr, v, 1, core.TxPayloadBinaryType, payload, gasPrice, gasLimit)

	priv1 := secp256k1.GeneratePrivateKey()
	signature, _ := crypto.NewSignature(keystore.SECP256K1)
	signature.InitSign(priv1)
	tx.Sign(signature)
	return tx
}

func TestRunScriptSource(t *testing.T) {
	tests := []struct {
		filepath       string
		expectedErr    error
		expectedResult string
	}{
		{"test/test_require.js", nil, "\"\""},
		{"test/test_console.js", nil, "\"\""},
		{"test/test_storage_handlers.js", nil, "\"\""},
		{"test/test_storage_class.js", nil, "\"\""},
		{"test/test_storage.js", nil, "\"\""},
		{"test/test_eval.js", core.ErrExecutionFailed, "EvalError: Code generation from strings disallowed for this context"},
		{"test/test_date.js", core.ErrExecutionFailed, "TypeError: Date.now is not a function"},
		{"test/test_bignumber_random.js", core.ErrExecutionFailed, "Error: BigNumber.random is not allowed in nvm."},
		{"test/test_random.js", core.ErrExecutionFailed, "Error: Math.random func is not allowed in nvm."},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(900000, 10000000)
			result, err := engine.RunScriptSource(string(data), 0)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedResult, result)
			engine.Dispose()
		})
	}
}

func TestRunScriptSourceInModule(t *testing.T) {
	tests := []struct {
		filepath    string
		sourceType  string
		expectedErr error
	}{
		{"./test/test_require.js", "js", nil},
		{"./test/test_setTimeout.js", "js", core.ErrExecutionFailed},
		{"./test/test_console.js", "js", nil},
		{"./test/test_storage_handlers.js", "js", nil},
		{"./test/test_storage_class.js", "js", nil},
		{"./test/test_storage.js", "js", nil},
		{"./test/test_ERC20.js", "js", nil},
		{"./test/test_eval.js", "js", core.ErrExecutionFailed},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(100000, 10000000)
			engine.AddModule(ContractName, string(data), 0)
			runnableSource := fmt.Sprintf("require(\"%s\");", ContractName)
			_, err = engine.RunScriptSource(runnableSource, 0)

			assert.Equal(t, tt.expectedErr, err)
			engine.Dispose()
		})
	}
}

func TestRunScriptSourceWithLimits(t *testing.T) {
	tests := []struct {
		name                          string
		filepath                      string
		limitsOfExecutionInstructions uint64
		limitsOfTotalMemorySize       uint64
		expectedErr                   error
	}{
		{"1", "test/test_oom_1.js", 100000, 0, ErrInsufficientGas},
		{"2", "test/test_oom_1.js", 0, 500000, ErrExceedMemoryLimits},
		{"3", "test/test_oom_1.js", 1000000, 50000000, ErrInsufficientGas},
		{"4", "test/test_oom_1.js", 5000000, 70000, ErrExceedMemoryLimits},

		{"5", "test/test_oom_2.js", 100000, 0, ErrInsufficientGas},
		{"6", "test/test_oom_2.js", 0, 80000, ErrExceedMemoryLimits},
		{"7", "test/test_oom_2.js", 10000000, 10000000, ErrInsufficientGas},
		{"8", "test/test_oom_2.js", 10000000, 70000, ErrExceedMemoryLimits},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(100000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			// direct run.
			(func() {
				engine := NewV8Engine(ctx)
				engine.SetExecutionLimits(tt.limitsOfExecutionInstructions, tt.limitsOfTotalMemorySize)
				source, _, _ := engine.InjectTracingInstructions(string(data))
				_, err = engine.RunScriptSource(source, 0)
				assert.Equal(t, tt.expectedErr, err)
				engine.Dispose()
			})()

			// modularized run.
			(func() {
				moduleID := fmt.Sprintf("%s", ContractName)
				runnableSource := fmt.Sprintf("require(\"%s\");", moduleID)

				engine := NewV8Engine(ctx)
				engine.SetExecutionLimits(tt.limitsOfExecutionInstructions, tt.limitsOfTotalMemorySize)
				engine.AddModule(ContractName, string(data), 0)
				_, err = engine.RunScriptSource(runnableSource, 0)
				assert.Equal(t, tt.expectedErr, err)
				engine.Dispose()
			})()
		})
	}
}

func TestRunScriptSourceTimeout(t *testing.T) {
	tests := []struct {
		filepath string
	}{
		{"test/test_infinite_loop.js"},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)

			// owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			// assert.Nil(t, err)

			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			// direct run.
			(func() {
				engine := NewV8Engine(ctx)
				_, err = engine.RunScriptSource(string(data), 0)
				assert.Equal(t, ErrExecutionTimeout, err)
				engine.Dispose()
			})()

			// modularized run.
			(func() {
				moduleID := fmt.Sprintf("%s", ContractName)
				runnableSource := fmt.Sprintf("require(\"%s\");", moduleID)

				engine := NewV8Engine(ctx)
				engine.AddModule(moduleID, string(data), 0)
				_, err = engine.RunScriptSource(runnableSource, 0)
				assert.Equal(t, ErrExecutionTimeout, err)
				engine.Dispose()
			})()
		})
	}
}

func TestDeployAndInitAndCall(t *testing.T) {
	tests := []struct {
		name         string
		contractPath string
		sourceType   string
		initArgs     string
		verifyArgs   string
	}{
		{"deploy sample_contract.js", "./test/sample_contract.js", "js", "[\"TEST001\", 123,[{\"name\":\"robin\",\"count\":2},{\"name\":\"roy\",\"count\":3},{\"name\":\"leon\",\"count\":4}]]", "[\"TEST001\", 123,[{\"name\":\"robin\",\"count\":2},{\"name\":\"roy\",\"count\":3},{\"name\":\"leon\",\"count\":4}]]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contractPath)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)

			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, tt.initArgs)
			assert.Nil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "dump", "")
			assert.Nil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "verify", tt.verifyArgs)
			assert.Nil(t, err)
			engine.Dispose()

			// force error.
			mem, _ = storage.NewMemoryStorage()
			context, _ = state.NewWorldState(dpos.NewDpos(), mem)
			owner, err = context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			contract, err = context.CreateContractAccount([]byte("account2"), nil)
			assert.Nil(t, err)

			ctx, err = NewContext(mockBlock(), mockTransaction(), contract, context)
			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "verify", tt.verifyArgs)
			assert.NotNil(t, err)
			engine.Dispose()
		})
	}
}

func TestERC20(t *testing.T) {
	tests := []struct {
		name         string
		contractPath string
		sourceType   string
		initArgs     string
		totalSupply  string
	}{
		{"deploy ERC20.js", "./test/ERC20.js", "js", "[\"TEST001\", \"TEST\", 1000000000]", "1000000000"},
	}

	// TODO: Addd more test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contractPath)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)

			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, tt.initArgs)
			assert.Nil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "totalSupply", "[]")
			assert.Nil(t, err)
			engine.Dispose()

		})
	}
}

func TestContracts(t *testing.T) {
	type fields struct {
		function string
		args     string
	}
	tests := []struct {
		contract   string
		sourceType string
		initArgs   string
		calls      []fields
	}{
		{
			"./test/contract_rectangle.js",
			"js",
			"[\"1024\", \"768\"]",
			[]fields{
				{"calcArea", "[]"},
				{"verify", "[\"786432\"]"},
			},
		},
		{
			"./test/contract_rectangle.js",
			"js",
			"[\"999\", \"123\"]",
			[]fields{
				{"calcArea", "[]"},
				{"verify", "[\"122877\"]"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.contract, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contract)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))
			contract, err := context.CreateContractAccount([]byte("account2"), nil)
			assert.Nil(t, err)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			// deploy and init.
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(1000, 10000000)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, tt.initArgs)
			assert.Nil(t, err)
			engine.Dispose()

			// call.
			for _, fields := range tt.calls {
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(1000, 10000000)
				_, err = engine.Call(string(data), tt.sourceType, fields.function, fields.args)
				assert.Nil(t, err)
				engine.Dispose()
			}
		})
	}
}

func TestFunctionNameCheck(t *testing.T) {
	tests := []struct {
		function    string
		expectedErr error
		args        string
	}{
		{"$dump", nil, ""},
		{"dump", nil, ""},
		{"dump_1", nil, ""},
		{"init", ErrDisallowCallPrivateFunction, ""},
		{"Init", ErrDisallowCallPrivateFunction, ""},
		{"9dump", ErrDisallowCallNotStandardFunction, ""},
		{"_dump", ErrDisallowCallNotStandardFunction, ""},
	}

	for _, tt := range tests {
		t.Run(tt.function, func(t *testing.T) {
			data, err := ioutil.ReadFile("test/sample_contract.js")
			sourceType := "js"
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(1000, 10000000)
			_, err = engine.Call(string(data), sourceType, tt.function, tt.args)
			assert.Equal(t, tt.expectedErr, err)
			engine.Dispose()
		})
	}
}

func TestMultiEngine(t *testing.T) {
	mem, _ := storage.NewMemoryStorage()
	context, _ := state.NewWorldState(dpos.NewDpos(), mem)
	owner, err := context.GetOrCreateUserAccount([]byte("account1"))
	assert.Nil(t, err)
	owner.AddBalance(newUint128FromIntWrapper(1000000))
	contract, _ := context.CreateContractAccount([]byte("account2"), nil)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(1000, 10000000)
			defer engine.Dispose()

			_, err = engine.RunScriptSource("console.log('running.');", 0)
			assert.Nil(t, err)
		}()
	}
	wg.Wait()
}

func TestInstructionCounterTestSuite(t *testing.T) {
	tests := []struct {
		filepath                                string
		strictDisallowUsageOfInstructionCounter int
		expectedErr                             error
		expectedResult                          string
	}{
		{"./test/instruction_counter_tests/redefine1.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine2.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine3.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine4.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine5.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine6.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine7.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/function.js", 1, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine1.js", 0, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine2.js", 0, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine3.js", 0, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine4.js", 0, core.ErrExecutionFailed, "Error: still not break the jail of _instruction_counter."},
		{"./test/instruction_counter_tests/redefine5.js", 0, ErrInjectTracingInstructionFailed, ""},
		{"./test/instruction_counter_tests/redefine6.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/redefine7.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/function.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/if.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/switch.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/for.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/with.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/while.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/throw.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/switch.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/condition_operator.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/storage_usage.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/event_usage.js", 0, nil, "\"\""},
		{"./test/instruction_counter_tests/blockchain_usage.js", 0, nil, "\"\""},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000000))
			contract, err := context.CreateContractAccount([]byte("account2"), nil)
			assert.Nil(t, err)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			moduleID := ContractName
			runnableSource := fmt.Sprintf("var x = require(\"%s\");", moduleID)

			engine := NewV8Engine(ctx)
			engine.strictDisallowUsageOfInstructionCounter = tt.strictDisallowUsageOfInstructionCounter
			engine.enableLimits = true
			err = engine.AddModule(moduleID, string(data), 0)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				result, err := engine.RunScriptSource(runnableSource, 0)
				assert.Equal(t, tt.expectedErr, err)
				assert.Equal(t, tt.expectedResult, result)
			}
			engine.Dispose()
		})
	}
}

func TestTypeScriptExecution(t *testing.T) {
	tests := []struct {
		filepath    string
		expectedErr error
	}{
		{"./test/test_greeter.ts", nil},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000000))
			contract, err := context.CreateContractAccount([]byte("account2"), nil)
			assert.Nil(t, err)
			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

			moduleID := ContractName
			runnableSource := fmt.Sprintf("require(\"%s\");", moduleID)

			engine := NewV8Engine(ctx)
			defer engine.Dispose()

			engine.enableLimits = true
			jsSource, _, err := engine.TranspileTypeScript(string(data))
			if err != nil {
				assert.Equal(t, tt.expectedErr, err)
				return
			}

			err = engine.AddModule(moduleID, string(jsSource), 0)
			if err != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				_, err := engine.RunScriptSource(runnableSource, 0)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func DeprecatedTestRunMozillaJSTestSuite(t *testing.T) {
	mem, _ := storage.NewMemoryStorage()
	context, _ := state.NewWorldState(dpos.NewDpos(), mem)
	owner, err := context.GetOrCreateUserAccount([]byte("account1"))
	assert.Nil(t, err)
	owner.AddBalance(newUint128FromIntWrapper(1000000000))

	contract, err := context.CreateContractAccount([]byte("account2"), nil)
	assert.Nil(t, err)
	ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

	var runTest func(dir string, shelljs string)
	runTest = func(dir string, shelljs string) {
		files, err := ioutil.ReadDir(dir)
		require.Nil(t, err)

		cwdShelljs := fmt.Sprintf("%s/shell.js", dir)
		if _, err := os.Stat(cwdShelljs); !os.IsNotExist(err) {
			shelljs = fmt.Sprintf("%s;%s", shelljs, cwdShelljs)
		}

		for _, file := range files {
			filepath := fmt.Sprintf("%s/%s", dir, file.Name())
			fi, err := os.Stat(filepath)
			require.Nil(t, err)

			if fi.IsDir() {
				runTest(filepath, shelljs)
				continue
			}

			if !strings.HasSuffix(file.Name(), ".js") {
				continue
			}
			if strings.Compare(file.Name(), "browser.js") == 0 || strings.Compare(file.Name(), "shell.js") == 0 || strings.HasPrefix(file.Name(), "toLocale") {
				continue
			}

			buf := bytes.NewBufferString("this.print = console.log;var native_eval = eval;eval = function (s) { try {  return native_eval(s); } catch (e) { return \"error\"; }};")

			jsfiles := fmt.Sprintf("%s;%s;%s", shelljs, "test/mozilla_js_tests_loader.js", filepath)

			for _, v := range strings.Split(jsfiles, ";") {
				if len(v) == 0 {
					continue
				}

				fi, err := os.Stat(v)
				require.Nil(t, err)
				f, err := os.Open(v)
				require.Nil(t, err)
				reader := bufio.NewReader(f)
				buf.Grow(int(fi.Size()))
				buf.ReadFrom(reader)
			}
			// execute.
			engine := NewV8Engine(ctx)
			engine.SetTestingFlag(true)
			engine.enableLimits = true
			_, err = engine.RunScriptSource(buf.String(), 0)
			//t.Logf("ret:%v, err:%v", ret, err)
			assert.Nil(t, err)
		}
	}

	runTest("test/mozilla_js_tests", "")
}

func TestBlockChain(t *testing.T) {
	tests := []struct {
		filepath    string
		expectedErr error
	}{
		{"test/test_blockchain.js", nil},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("n1FkntVUMPAsESuCAAPK711omQk19JotBjM"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000000))
			contract, err := context.CreateContractAccount([]byte("n1JNHZJEUvfBYfjDRD14Q73FX62nJAzXkMR"), nil)
			assert.Nil(t, err)

			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(100000, 10000000)
			_, err = engine.RunScriptSource(string(data), 0)
			assert.Equal(t, tt.expectedErr, err)
			engine.Dispose()
		})
	}
}

func TestBankVaultContract(t *testing.T) {
	type TakeoutTest struct {
		args          string
		expectedErr   error
		beforeBalance string
		afterBalance  string
	}

	tests := []struct {
		name         string
		contractPath string
		sourceType   string
		saveValue    string
		saveArgs     string
		takeoutTests []TakeoutTest
	}{
		{"deploy bank_vault_contract.js", "./test/bank_vault_contract.js", "js", "5", "[0]",
			[]TakeoutTest{
				{"[1]", nil, "5", "4"},
				{"[5]", core.ErrExecutionFailed, "4", "4"},
				{"[4]", nil, "4", "0"},
				{"[1]", core.ErrExecutionFailed, "0", "0"},
			},
		},
		{"deploy bank_vault_contract.ts", "./test/bank_vault_contract.ts", "ts", "5", "[0]",
			[]TakeoutTest{
				{"[1]", nil, "5", "4"},
				{"[5]", core.ErrExecutionFailed, "4", "4"},
				{"[4]", nil, "4", "0"},
				{"[1]", core.ErrExecutionFailed, "0", "0"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contractPath)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))

			// prepare the contract.
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			contract.AddBalance(newUint128FromIntWrapper(5))

			// parepare env, block & transactions.
			tx := mockNormalTransaction("n1FkntVUMPAsESuCAAPK711omQk19JotBjM", "n1JNHZJEUvfBYfjDRD14Q73FX62nJAzXkMR", tt.saveValue)
			ctx, err := NewContext(mockBlock(), tx, contract, context)

			// execute.
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, "")
			assert.Nil(t, err)
			engine.Dispose()

			// call save.
			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			_, err = engine.Call(string(data), tt.sourceType, "save", tt.saveArgs)
			assert.Nil(t, err)
			engine.Dispose()

			var (
				bal struct {
					Balance string `json:"balance"`
				}
			)

			// call takeout.
			for _, tot := range tt.takeoutTests {
				// call balanceOf.
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				balance, err := engine.Call(string(data), tt.sourceType, "balanceOf", "")
				assert.Nil(t, err)
				bal.Balance = ""
				err = json.Unmarshal([]byte(balance), &bal)
				assert.Nil(t, err)
				assert.Equal(t, tot.beforeBalance, bal.Balance)
				engine.Dispose()

				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				_, err = engine.Call(string(data), tt.sourceType, "takeout", tot.args)
				assert.Equal(t, err, tot.expectedErr)
				engine.Dispose()

				// call balanceOf.
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				balance, err = engine.Call(string(data), tt.sourceType, "balanceOf", "")
				assert.Nil(t, err)
				bal.Balance = ""
				err = json.Unmarshal([]byte(balance), &bal)
				assert.Nil(t, err)
				assert.Equal(t, tot.afterBalance, bal.Balance)
				engine.Dispose()
			}
		})
	}
}

func TestEvent(t *testing.T) {
	tests := []struct {
		filepath string
	}{
		{"test/test_event.js"},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.filepath)
			assert.Nil(t, err, "filepath read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("n1FkntVUMPAsESuCAAPK711omQk19JotBjM"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(1000000000))
			contract, _ := context.CreateContractAccount([]byte("n1JNHZJEUvfBYfjDRD14Q73FX62nJAzXkMR"), nil)

			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(100000, 10000000)
			_, err = engine.RunScriptSource(string(data), 0)
			engine.Dispose()
		})
	}
}

func TestNRC20Contract(t *testing.T) {
	type TransferTest struct {
		to     string
		result bool
		value  string
	}

	tests := []struct {
		test          string
		contractPath  string
		sourceType    string
		name          string
		symbol        string
		decimals      int
		totalSupply   string
		from          string
		transferTests []TransferTest
	}{
		{"nrc20", "./test/NRC20.js", "js", "StandardToken标准代币", "ST", 18, "1000000000",
			"n1FkntVUMPAsESuCAAPK711omQk19JotBjM",
			[]TransferTest{
				{"n1FkntVUMPAsESuCAAPK711omQk19JotBjM", true, "5"},
				{"n1JNHZJEUvfBYfjDRD14Q73FX62nJAzXkMR", true, "10"},
				{"n1Kjom3J4KPsHKKzZ2xtt8Lc9W5pRDjeLcW", true, "15"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contractPath)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte(tt.from))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))

			// prepare the contract.
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)
			contract.AddBalance(newUint128FromIntWrapper(5))

			// parepare env, block & transactions.
			tx := mockNormalTransaction(tt.from, "n1TV3sU6jyzR4rJ1D7jCAmtVGSntJagXZHC", "0")
			ctx, err := NewContext(mockBlock(), tx, contract, context)

			// execute.
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			args := fmt.Sprintf("[\"%s\", \"%s\", %d, \"%s\"]", tt.name, tt.symbol, tt.decimals, tt.totalSupply)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, args)
			assert.Nil(t, err)
			engine.Dispose()

			// call name.
			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			name, err := engine.Call(string(data), tt.sourceType, "name", "")
			assert.Nil(t, err)
			var nameStr string
			err = json.Unmarshal([]byte(name), &nameStr)
			assert.Nil(t, err)
			assert.Equal(t, tt.name, nameStr)
			engine.Dispose()

			// call symbol.
			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			symbol, err := engine.Call(string(data), tt.sourceType, "symbol", "")
			assert.Nil(t, err)
			var symbolStr string
			err = json.Unmarshal([]byte(symbol), &symbolStr)
			assert.Nil(t, err)
			assert.Equal(t, tt.symbol, symbolStr)
			assert.Nil(t, err)
			engine.Dispose()

			// call decimals.
			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			decimals, err := engine.Call(string(data), tt.sourceType, "decimals", "")
			assert.Nil(t, err)
			var decimalsInt int
			err = json.Unmarshal([]byte(decimals), &decimalsInt)
			assert.Nil(t, err)
			assert.Equal(t, tt.decimals, decimalsInt)
			assert.Nil(t, err)
			engine.Dispose()

			// call totalSupply.
			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			totalSupply, err := engine.Call(string(data), tt.sourceType, "totalSupply", "")
			assert.Nil(t, err)
			var totalSupplyStr string
			err = json.Unmarshal([]byte(totalSupply), &totalSupplyStr)
			assert.Nil(t, err)
			expect, _ := big.NewInt(0).SetString(tt.totalSupply, 10)
			expect = expect.Mul(expect, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(tt.decimals)), nil))
			assert.Equal(t, expect.String(), totalSupplyStr)
			assert.Nil(t, err)
			engine.Dispose()

			// call takeout.
			for _, tot := range tt.transferTests {
				// call balanceOf.
				ctx.tx = mockNormalTransaction(tt.from, "n1TV3sU6jyzR4rJ1D7jCAmtVGSntJagXZHC", "0")
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				balArgs := fmt.Sprintf("[\"%s\"]", tt.from)
				_, err := engine.Call(string(data), tt.sourceType, "balanceOf", balArgs)
				assert.Nil(t, err)
				engine.Dispose()

				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				transferArgs := fmt.Sprintf("[\"%s\", \"%s\"]", tot.to, tot.value)
				result, err := engine.Call(string(data), tt.sourceType, "transfer", transferArgs)
				assert.Nil(t, err)
				assert.Equal(t, "\"\"", result)
				engine.Dispose()

				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				approveArgs := fmt.Sprintf("[\"%s\", \"0\", \"%s\"]", tot.to, tot.value)
				result, err = engine.Call(string(data), tt.sourceType, "approve", approveArgs)
				assert.Nil(t, err)
				assert.Equal(t, "\"\"", result)
				engine.Dispose()

				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				allowanceArgs := fmt.Sprintf("[\"%s\", \"%s\"]", tt.from, tot.to)
				amount, err := engine.Call(string(data), tt.sourceType, "allowance", allowanceArgs)
				assert.Nil(t, err)
				var amountStr string
				err = json.Unmarshal([]byte(amount), &amountStr)
				assert.Nil(t, err)
				assert.Equal(t, tot.value, amountStr)
				engine.Dispose()

				ctx.tx = mockNormalTransaction(tot.to, "n1TV3sU6jyzR4rJ1D7jCAmtVGSntJagXZHC", "0")
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				transferFromArgs := fmt.Sprintf("[\"%s\", \"%s\", \"%s\"]", tt.from, tot.to, tot.value)
				result, err = engine.Call(string(data), tt.sourceType, "transferFrom", transferFromArgs)
				assert.Nil(t, err)
				assert.Equal(t, "\"\"", result)
				engine.Dispose()

				ctx.tx = mockNormalTransaction(tot.to, "n1TV3sU6jyzR4rJ1D7jCAmtVGSntJagXZHC", "0")
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 100000000)
				transferFromArgs = fmt.Sprintf("[\"%s\", \"%s\", \"%s\"]", tt.from, tot.to, tot.value)
				_, err = engine.Call(string(data), tt.sourceType, "transferFrom", transferFromArgs)
				assert.NotNil(t, err)
				engine.Dispose()
			}
		})
	}
}

func TestNRC20ContractMultitimes(t *testing.T) {
	for i := 0; i < 5; i++ {
		TestNRC20Contract(t)
	}
}

func TestNebulasContract(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		function string
		args     string
		err      error
	}{
		{"1", "0", "unpayable", "", nil},
		{"2", "0", "unpayable", "[1]", nil},
		{"3", "1", "unpayable", "", nil},
		{"4", "0", "payable", "", core.ErrExecutionFailed},
		{"5", "1", "payable", "", nil},
		{"6", "1", "payable", "[1]", nil},
		{"7", "0", "contract1", "[1]", nil},
		{"8", "1", "contract1", "[1]", nil},
		{"9", "0", "contract2", "[1]", core.ErrExecutionFailed},
		{"10", "1", "contract2", "[1]", core.ErrExecutionFailed},
		{"11", "0", "contract3", "[1]", core.ErrExecutionFailed},
		{"12", "1", "contract3", "[1]", nil},
		{"13", "0", "contract4", "[1]", core.ErrExecutionFailed},
		{"14", "1", "contract4", "[1]", core.ErrExecutionFailed},
	}

	mem, _ := storage.NewMemoryStorage()
	context, _ := state.NewWorldState(dpos.NewDpos(), mem)

	addr, _ := core.NewAddressFromPublicKey([]byte{
		2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7,
		2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7,
		2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 2, 3, 5, 7, 1, 2, 4, 5, 3})
	owner, err := context.GetOrCreateUserAccount(addr.Bytes())
	assert.Nil(t, err)
	owner.AddBalance(newUint128FromIntWrapper(1000000000))

	addr, _ = core.NewContractAddressFromData([]byte{1, 2, 3, 5, 7}, []byte{1, 2, 3, 5, 7})
	contract, _ := context.CreateContractAccount(addr.Bytes(), nil)

	ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)

	data, err := ioutil.ReadFile("test/mixin.js")
	assert.Nil(t, err, "filepath read error")
	sourceType := "js"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx.tx = mockNormalTransaction("n1FkntVUMPAsESuCAAPK711omQk19JotBjM", "n1JNHZJEUvfBYfjDRD14Q73FX62nJAzXkMR", tt.value)
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 100000000)
			_, err := engine.Call(string(data), sourceType, tt.function, tt.args)
			assert.Equal(t, tt.err, err)
			engine.Dispose()
		})
	}
}
func TestTransferValueFromContracts(t *testing.T) {
	type fields struct {
		function string
		args     string
	}
	tests := []struct {
		contract   string
		sourceType string
		initArgs   string
		calls      []fields
		value      string
		success    bool
	}{
		{
			"./test/transfer_value_from_contract.js",
			"js",
			"",
			[]fields{
				{"transfer", "[\"n1SAeQRVn33bamxN4ehWUT7JGdxipwn8b17\"]"},
			},
			"100",
			true,
		},
		{
			"./test/transfer_value_from_contract.js",
			"js",
			"",
			[]fields{
				{"transfer", "[\"n1SAeQRVn33bamxN4ehWUT7JGdxipwn8b17\"]"},
			},
			"101",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.contract, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contract)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))
			contract, err := context.CreateContractAccount([]byte("account2"), nil)
			assert.Nil(t, err)

			contract.AddBalance(newUint128FromIntWrapper(100))
			mockTx := mockNormalTransaction("n1FkntVUMPAsESuCAAPK711omQk19JotBjM", "n1FkntVUMPAsESuCAAPK711omQk19JotBjM", tt.value)
			ctx, err := NewContext(mockBlock(), mockTx, contract, context)

			// deploy and init.
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(1000, 10000000)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, tt.initArgs)
			assert.Nil(t, err)
			engine.Dispose()

			// call.
			for _, fields := range tt.calls {
				engine = NewV8Engine(ctx)
				engine.SetExecutionLimits(10000, 10000000)
				result, err := engine.Call(string(data), tt.sourceType, fields.function, fields.args)
				if tt.success {
					assert.Equal(t, result, "\""+fmt.Sprint(tt.value)+"\"")
					assert.Nil(t, err)
				} else {
					assert.NotNil(t, err)
				}
				engine.Dispose()
			}
		})
	}
}

func TestRequireModule(t *testing.T) {
	tests := []struct {
		name         string
		contractPath string
		sourceType   string
		initArgs     string
	}{
		{"deploy test_require_module.js", "./test/test_require_module.js", "js", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := ioutil.ReadFile(tt.contractPath)
			assert.Nil(t, err, "contract path read error")

			mem, _ := storage.NewMemoryStorage()
			context, _ := state.NewWorldState(dpos.NewDpos(), mem)
			owner, err := context.GetOrCreateUserAccount([]byte("account1"))
			assert.Nil(t, err)
			owner.AddBalance(newUint128FromIntWrapper(10000000))
			contract, _ := context.CreateContractAccount([]byte("account2"), nil)

			ctx, err := NewContext(mockBlock(), mockTransaction(), contract, context)
			engine := NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.DeployAndInit(string(data), tt.sourceType, tt.initArgs)
			assert.Nil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "requireNULL", "")
			assert.NotNil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "requireNotExistPath", "")
			assert.NotNil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "requireCurPath", "")
			assert.NotNil(t, err)
			engine.Dispose()

			engine = NewV8Engine(ctx)
			engine.SetExecutionLimits(10000, 10000000)
			_, err = engine.Call(string(data), tt.sourceType, "requireNotExistFile", "")
			assert.NotNil(t, err)
			engine.Dispose()
		})
	}
}
