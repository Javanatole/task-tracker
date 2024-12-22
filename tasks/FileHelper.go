package tasks

import (
	"os"
)

func writeContentIntoFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	return err
}

func readContentFromFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	return string(content), err
}
