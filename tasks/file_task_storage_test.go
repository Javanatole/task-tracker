package tasks

import (
	"os"
	"testing"
)

// TestFileTaskStorage_Read tests the Read method of FileTaskStorage
func TestFileTaskStorage_Read(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("can't remove the temp file")
		}
	}(tempFile.Name())

	// Write test data to the temporary file
	expectedContent := "Hello, FileTaskStorage!"
	if _, err := tempFile.Write([]byte(expectedContent)); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	err = tempFile.Close()
	if err != nil {
		t.Fatalf("Can't close temp file")
	}

	// Create FileTaskStorage with the temp file
	storage := FileTaskStorage{Filename: tempFile.Name()}
	// Call the Read method
	actualContent, err := storage.Read()
	if err != nil {
		t.Fatalf("unexpected error during Read: %v", err)
	}

	// Assert the content matches
	if actualContent != expectedContent {
		t.Errorf("expected %q, got %q", expectedContent, actualContent)
	}
}

func TestFileTaskStorage_Write(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Can't remove the temp file")
		}
	}(tempFile.Name())

	storage := FileTaskStorage{Filename: tempFile.Name()}

	// Define content to write
	contentToWrite := "Testing Write Method!"

	if err := storage.Write(contentToWrite); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// Verify the file's content
	actualContent, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	if string(actualContent) != contentToWrite {
		t.Errorf("expected %q, got %q", contentToWrite, actualContent)
	}
}
