package Counter

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to simulate file reading from a string
func mockFile(content string) *os.File {
	file, _ := os.CreateTemp("", "mock")
	file.WriteString(content)
	file.Seek(0, io.SeekStart)
	return file
}

// TestNewCounter checks that the factory function returns the correct Counter type
func TestNewCounter(t *testing.T) {
	tests := []struct {
		language string
		expected string
	}{
		{"go", "Counter.GoCounter"},
		{"c", "Counter.CCounter"},
		{"python", "Counter.PythonCounter"},
	}

	for _, test := range tests {
		counter, err := NewCounter(test.language)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		actualType := strings.TrimPrefix(strings.TrimPrefix(fmt.Sprintf("%T", counter), "*"), "main.")
		if actualType != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, actualType)
		}
	}
}

// TestGoCounter tests the Go language counter
func TestGoCounter(t *testing.T) {
	content := `
		// This is a Go file
		import "fmt"
		
		var greeting = "Hello, World!"
		fmt.Println(greeting)
	`
	file := mockFile(content)
	defer os.Remove(file.Name())

	counter := &GoCounter{}
	result, err := counter.CountLines(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.CommentLines != 1 {
		t.Errorf("Expected 1 comment line, got %d", result.CommentLines)
	}
	if result.ImportLines != 1 {
		t.Errorf("Expected 1 import line, got %d", result.ImportLines)
	}
	if result.VariableLines != 1 {
		t.Errorf("Expected 1 variable line, got %d", result.VariableLines)
	}
}

// TestCCounter tests the C language counter
func TestCCounter(t *testing.T) {
	content := `
		// This is a C file
		#include <stdio.h>
		
		int main() {
			char greeting[] = "Hello, C!";
			printf("%s\n", greeting);
			return 0;
		}
	`
	file := mockFile(content)
	defer os.Remove(file.Name())

	counter := &CCounter{}
	result, err := counter.CountLines(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.CommentLines != 1 {
		t.Errorf("Expected 1 comment line, got %d", result.CommentLines)
	}
	if result.ImportLines != 1 {
		t.Errorf("Expected 1 import line, got %d", result.ImportLines)
	}
}

// TestPythonCounter tests the Python language counter
func TestPythonCounter(t *testing.T) {
	content := `
		# This is a Python file
		import os
		greeting = "Hello, Python!"
		print(greeting)
	`
	file := mockFile(content)
	defer os.Remove(file.Name())

	counter := &PythonCounter{}
	result, err := counter.CountLines(file)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.CommentLines != 1 {
		t.Errorf("Expected 1 comment line, got %d", result.CommentLines)
	}
	if result.ImportLines != 1 {
		t.Errorf("Expected 1 import line, got %d", result.ImportLines)
	}
	if result.VariableLines != 1 {
		t.Errorf("Expected 1 variable line, got %d", result.VariableLines)
	}
}
