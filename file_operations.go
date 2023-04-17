package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
data, err := FileToSlice("data.txt")
if err != nil {
	fmt.Println(err)
}
*/

// Main package function
func FileToSlice(file string) ([]string, error) {
	// Read data from txt file
	data, err := ReadFileContent(file)
	if err != nil {
		return []string{}, err
	}

	return splitLines(data), nil
}

// Read given file content
// I am exporting this because can be usefull for reading file with only one line
func ReadFileContent(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
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
