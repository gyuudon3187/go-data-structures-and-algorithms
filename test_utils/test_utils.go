package testutils

import "testing"

func ValidateResult(t *testing.T, got, want interface{}) {
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
