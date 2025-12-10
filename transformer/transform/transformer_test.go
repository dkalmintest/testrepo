package transformer_test

import (
	transformer "gh/transform/transformer"
	"testing"
)

func transformerTest_IsTrue(t *testing.T) {
	var transformResult := transformer.Transform()
	if transformResult != true {
		t.Errorf("Expected true, got '%v'", transformResult)
	} else {
		t.Log("we passed!")
	}
}
