package main

import "testing"

func TestIsPositive(t *testing.T) {
	_, err := IsPositive(-21)
	if err == nil {
		t.Error("-21 is not a positive number")
	}
}
