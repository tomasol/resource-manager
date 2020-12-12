package pools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
)

func readQuickJs() ([]byte, error) {
	qjsFile, ok := os.LookupEnv("WASMER_JS")
	if !ok {
		return nil, errors.New("Environment variable WASMER_JS must be set")
	}
	return ioutil.ReadFile(qjsFile)
}

func createQuickJSInstance(args []string) *wasmtime.Instance {
	wasmBytes, err := readQuickJs()
	check(err)

	wasmtimeConfig := wasmtime.NewConfig()
	wasmtimeConfig.CacheConfigLoadDefault()
	engine := wasmtime.NewEngineWithConfig(wasmtimeConfig)

	store := wasmtime.NewStore(engine)

	wasiConfig := wasmtime.NewWasiConfig()

	wasiConfig.SetArgv(args)
	// 'wasi_snapshot_preview1' errors with: unknown import: `wasi_unstable::fd_prestat_get` has not been defined
	wasiInstance, err := wasmtime.NewWasiInstance(store, wasiConfig, "wasi_unstable")
	check(err)

	module, err := wasmtime.NewModule(engine, wasmBytes)
	check(err)

	linker := wasmtime.NewLinker(store)
	err = linker.DefineWasi(wasiInstance)
	check(err)

	instance, err := linker.Instantiate(module)
	check(err)
	return instance
}

func TestWasmtime(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	instance := createQuickJSInstance([]string{"--std", "-e", "2+2"})
	run := instance.GetExport("_start").Func()
	result, err := run.Call()
	check(err)
	fmt.Printf("DONE: %T %v", result, result)
}

func BenchmarkWasmtimeQuickJS(b *testing.B) {
	wasmBytes, err := readQuickJs()
	check(err)
	wasmtimeConfig := wasmtime.NewConfig()
	wasmtimeConfig.CacheConfigLoadDefault()
	engine := wasmtime.NewEngineWithConfig(wasmtimeConfig)
	store := wasmtime.NewStore(engine)

	for n := 0; n < b.N; n++ {

		module, err := wasmtime.NewModule(engine, wasmBytes)
		check(err)
		linker := wasmtime.NewLinker(store)

		wasiConfig := wasmtime.NewWasiConfig()
		args := []string{"--std", "-e", "2+2"}
		wasiConfig.SetArgv(args)
		wasiInstance, err := wasmtime.NewWasiInstance(store, wasiConfig, "wasi_unstable")
		check(err)
		err = linker.DefineWasi(wasiInstance)
		check(err)
		instance, err := linker.Instantiate(module)
		run := instance.GetExport("_start").Func()
		_, err = run.Call()
		check(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
