package main

import (
	"testing"
)

func TestCheckCard(t *testing.T) {
	_, check := CheckCard(false)
	if check == true {
		t.Error("Expected false, got ", check)
	}
}

func TestCheckIsWhitelisted(t *testing.T) {
	check := IsWhitelisted(false)
	if check == true {
		t.Error("Expected false, got ", check)
	}
}

func TestCheckLimit(t *testing.T) {
	_, check := CheckLimit(500, 600)
	if check == true {
		t.Error("Expected false, got ", check)
	}
}
