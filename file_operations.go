package main

import (
	"fmt"
	"os"
	"strings"
)

func fileToSlice(file string) ([]string, error) {
	// Read data from a file
	data, err := readFileContent(file)
	if err != nil {
		return []string{}, err
	}

	return splitLines(data), nil
}

// Read the given file content
func readFileContent(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("File does not exist: %s", file)
	}

	return strings.TrimSpace(string(data)), nil
}

// Convert string to slice, so we can loop
// TODO: Test this in Linux
func splitLines(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		// Linux: "\n", Windows: "\r\n"
		return r == '\r' || r == '\n'
	})
}

// ------------------------------------------------------------------------------------------------------
// Error file
//

func createErrorFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func writeIPToErrorFile(ip, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	data := []byte(ip + "\n")
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
