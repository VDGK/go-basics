package files

import (
	"fmt"
	"os"
	"strings"
)

func ReadFromFile(filepath string) ([]byte, error) {
	ext, _ := CheckFileExtension(filepath)
	if !ext {
		return nil, fmt.Errorf("%s is not a valid file extension", filepath)
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s %w", filepath, err)
	}
	defer file.Close()
	data, err := os.ReadFile(filepath)
	if err != nil {
		return data, fmt.Errorf("failed to read file: %s %w", filepath, err)
	}
	return data, nil
}

func CheckFileExtension(filepath string) (bool, error) {
	if !strings.Contains(filepath, ".json") {
		return false, fmt.Errorf("%s is not a valid json file", filepath)
	}
	return true, nil
}
