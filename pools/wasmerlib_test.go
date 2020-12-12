package pools

import (
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/net-auto/resourceManager/ent/runtime"
	"github.com/pkg/errors"
	wasmer "github.com/tomasol/go-ext-wasm/wasmer"
)

func readQuickJs() ([]byte, error) {
	qjsFile, ok := os.LookupEnv("WASMER_JS")
	if !ok {
		return nil, errors.New("Environment variable WASMER_JS must be set")
	}
	return ioutil.ReadFile(qjsFile)
}

func TestQuickJS(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	wasmBytes, err := readQuickJs()
	if err != nil {
		t.Fatalf("Cannot load wasm file - %v", err)
	}
	module, err := wasmer.Compile(wasmBytes)
	if err != nil {
		t.Fatalf("Cannot create module - %v", err)
	}
	defer module.Close()
	importObject := wasmer.NewWasiImportObjectForVersion(wasmer.Latest, []string{"--std", "-e", "2+2"}, []string{}, []string{}, []wasmer.MapDirEntry{})
	// importObject := wasmer.NewDefaultWasiImportObject()

	instance, err := module.InstantiateWithImportObject(importObject)
	if err != nil {
		t.Fatalf("Cannot create instance - %v", err)
	}
	defer instance.Close()

	// Run the module
	instance.Exports["_start"]()

}
