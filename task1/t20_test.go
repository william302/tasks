package task1

import "testing"

func TestIsValid(t *testing.T) {
	b := isValid("()")
	if !b {
		t.Errorf("Expected true, got false")
	}

}
