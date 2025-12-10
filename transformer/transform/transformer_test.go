package transformer_test

import (
	transformer "gh/transform/transform"
	"testing"
)

func transformerTest_IsTrue(t *testing.T) {
	transformResult := transformer.Transform()
	if transformResult != true {
		t.Errorf("Expected true, got '%v'", transformResult)
	} else {
		t.Log("transformerTest_IsTrue passed!")
	}
}
