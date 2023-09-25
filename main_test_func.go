package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestInputValidation(t *testing.T) {
	// Test cases for input validation in ReadInput function
	tests := []struct {
		input                string
		expectedX, expectedY int
		expectedErr          string
	}{
		{"1,2\n", 1, 2, ""},
		{"invalid_input\n", 0, 0, "Invalid input. Please enter x,y"},
		// Add more test cases for input validation
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, w, _ := os.Pipe()
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = r

			go func() {
				defer w.Close()
				fmt.Fprint(w, tt.input)
			}()

			x, y, err := ReadInput()

			if err != nil && tt.expectedErr != "" {
				if !strings.Contains(err.Error(), tt.expectedErr) {
					t.Errorf("Expected error containing '%s', got '%s'", tt.expectedErr, err)
				}
			} else if x != tt.expectedX || y != tt.expectedY {
				t.Errorf("Expected (%d, %d), got (%d, %d)", tt.expectedX, tt.expectedY, x, y)
			}
		})
	}
}

func TestMainFunction(t *testing.T) {
	// Test cases for main function
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"1,2\n", "All Moves: ... \nAll Count: ... \n**************************\n"},
		// Add more test cases for main function
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, w, _ := os.Pipe()
			oldStdout := os.Stdout
			defer func() { os.Stdout = oldStdout }()
			os.Stdout = w

			go func() {
				defer w.Close()
				fmt.Fprint(w, tt.expectedOutput)
			}()

			os.Stdin = strings.NewReader(tt.input)
			output := captureOutput(func() {
				main()
			})

			if !strings.Contains(output, tt.expectedOutput) {
				t.Errorf("Expected output:\n%s\n\nGot output:\n%s", tt.expectedOutput, output)
			}
		})
	}
}
