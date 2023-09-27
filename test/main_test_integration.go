package main

import (
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	// Simulate user input
	input := "1,2\n"
	expectedX, expectedY := 1, 2

	// Replace os.Stdin with a custom reader
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r
	w.WriteString(input)
	w.Close()

	x, y, err := ReadInput()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if x != expectedX || y != expectedY {
		t.Errorf("Expected (%d, %d), got (%d, %d)", expectedX, expectedY, x, y)
	}
}
