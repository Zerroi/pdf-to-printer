package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	// 创建临时文件
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	err := os.WriteFile(tmpFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"existing file", tmpFile, true},
		{"non-existing file", filepath.Join(tmpDir, "nonexistent.txt"), false},
		{"directory", tmpDir, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FileExists(tt.path)
			if result != tt.expected {
				t.Errorf("FileExists(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestIsPDFFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"pdf file", "document.pdf", true},
		{"PDF file uppercase", "document.PDF", true},
		{"pdf file mixed case", "document.Pdf", true},
		{"txt file", "document.txt", false},
		{"doc file", "document.doc", false},
		{"no extension", "document", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPDFFile(tt.path)
			if result != tt.expected {
				t.Errorf("IsPDFFile(%q) = %v, want %v", tt.path, result, tt.expected)
			}
		})
	}
}
