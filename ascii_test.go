package main

import (
	"bytes"
	"strings"
	"testing"
)

// TestInitWithValidInput tests the Init function with valid input
func TestInitWithValidInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "single character",
			input:     "A",
			expectErr: false,
		},
		{
			name:      "multiple characters",
			input:     "Hello",
			expectErr: false,
		},
		{
			name:      "with newline escape sequence",
			input:     "Hello\\nWorld",
			expectErr: false,
		},
		{
			name:      "special characters",
			input:     "Test!",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := Init(tt.input, &buf)

			if (err != nil) != tt.expectErr {
				t.Errorf("Init() error = %v, expectErr %v", err, tt.expectErr)
			}

			// Output should not be empty for valid input
			if err == nil && tt.input != "" && buf.String() == "" {
				t.Error("Expected non-empty output for valid input")
			}
		})
	}
}

// TestInitWithEmptyInput tests the Init function with empty input
func TestInitWithEmptyInput(t *testing.T) {
	var buf bytes.Buffer
	err := Init("", &buf)

	if err != nil {
		t.Errorf("Init() with empty input should not error, got %v", err)
	}

	// Empty input should produce no output
	if buf.String() != "" {
		t.Error("Expected empty output for empty input")
	}
}


// TestFileStructureInitialization tests that File struct is properly initialized
func TestFileStructureInitialization(t *testing.T) {
	var buf bytes.Buffer
	err := Init("X", &buf)

	if err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	output := buf.String()

	// Verify output is not empty
	if output == "" {
		t.Error("Expected non-empty output after Init()")
	}
}

// TestEmbeddedFileContent tests that embedded file content is loaded
func TestEmbeddedFileContent(t *testing.T) {
	if len(embedFileContent) == 0 {
		t.Error("Embedded file content is empty or missing")
	}

	// Embedded content should contain multiple lines
	lines := strings.Split(embedFileContent, "\n")
	if len(lines) < 100 {
		t.Errorf("Expected substantial embedded content, got %d lines", len(lines))
	}
}


// TestOutputConsistency tests that the same input produces the same output
func TestOutputConsistency(t *testing.T) {
	var buf1, buf2 bytes.Buffer

	err1 := Init("Test", &buf1)
	err2 := Init("Test", &buf2)

	if err1 != nil || err2 != nil {
		t.Fatalf("Init() failed: err1=%v, err2=%v", err1, err2)
	}

	if buf1.String() != buf2.String() {
		t.Error("Expected consistent output for same input")
	}
}