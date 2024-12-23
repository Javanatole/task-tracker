package tasks

import (
	"os"
)

type FileHelper struct {
	filename string
}

func (fileHelper FileHelper) writeContentIntoFile(content string) error {
	err := os.WriteFile(fileHelper.filename, []byte(content), 0644)
	return err
}

func (fileHelper FileHelper) readContentFromFile() (string, error) {
	content, err := os.ReadFile(fileHelper.filename)
	return string(content), err
}
