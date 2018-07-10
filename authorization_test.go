package main

import "testing"

func TestNewLimit(t *testing.T) {
	v := NewLimit(500.00, 100.00)
	if v != 400.00 {
		t.Error("Expected 400.00, got ", v)
	}
}
