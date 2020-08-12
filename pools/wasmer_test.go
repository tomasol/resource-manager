package pools

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWasmerInvokeJsIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	wasmer, err := NewWasmerUsingEnvVars()
	if err != nil {
		t.Fatalf("Unable create wasmer - %s", err)
	}
	script := `
function invoke() {
	log('log');
	return {"result": userInput['name']};
}
`
	userInput := make(map[string]interface{})
	userInput["name"] = "Alice"
	actual, stderr, err := wasmer.invokeJs(script, userInput)
	if err != nil {
		t.Fatalf("Unable run - %s", err)
	}
	expected := make(map[string]interface{})
	if err := json.Unmarshal([]byte(`{"result": "Alice"}`), &expected); err != nil {
		t.Fatalf("Cannot unmarshal expected response - %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Unexpected evaluation response: %v, should be %v", actual, expected)
	}
	if stderr != "log\n" {
		t.Fatalf("Unexpected logging result: \"%s\"", stderr)
	}
}

func TestPrefixLines(t *testing.T) {
	actual := prefixLines("aa\nbb", "  ")
	expected := "  aa\n  bb\n"
	if actual != expected {
		t.Fatalf("Got \"%s\", expected \"%s\"", actual, expected)
	}
}

func TestWasmerInvokePyIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	wasmer, err := NewWasmerUsingEnvVars()
	if err != nil {
		t.Fatalf("Unable create wasmer - %s", err)
	}
	script := `
log('log')
return {"result": "Alice"}
`
	userInput := make(map[string]interface{})
	userInput["name"] = "Alice"
	actual, stderr, err := wasmer.invokePy(script, userInput)
	if err != nil {
		t.Fatalf("Unable run - %s", err)
	}
	expected := make(map[string]interface{})
	if err := json.Unmarshal([]byte(`{"result": "Alice"}`), &expected); err != nil {
		t.Fatalf("Cannot unmarshal expected response - %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Unexpected evaluation response: %v, should be %v", actual, expected)
	}
	if stderr != "log\n" {
		t.Fatalf("Unexpected logging result: \"%s\"", stderr)
	}
}
