package storage

import (
	"fmt"
	"os"
)

func WriteToFile(filepath string, data []byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("can't create file %s", filepath)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("can't write file %s", filepath)
	}
	fmt.Println("Write content to file successfully")
	return nil
}

func ReadFromFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("can't open file %s", filepath)
	}
	defer file.Close()
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("can't read file %s", filepath)
	}
	return data, nil
}
