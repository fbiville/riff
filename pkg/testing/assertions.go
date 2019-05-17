package testing

import "testing"

func ExpectNonNil(t *testing.T, object interface{}) {
	if object == nil {
		t.Error("Expected object not to be nil")
	}
}

func ExpectNonEmptyString(t *testing.T, s string, errorMsg string) {
	if s == "" {
		t.Error(errorMsg)
	}
}


