package transformer_test

import (
	transformer "gh/transform/transformer"
	"testing"
)

func transformerTest_IsTrue(t *testing.T) {
	transformResult := transformer.Transform()
	if transformResult != true {
		t.Errorf("Expected true, got '%v'", daGoResult)
	} else {
		t.Log("we passed!")
	}
}
