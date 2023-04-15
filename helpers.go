package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Read given file content
func readFileContent(file string) (string, error) {
	// Read the file into a []byte slice
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("\n%s File does not exist: %s", errorColor("[ERROR]"), infoColor(file))
	}

	return strings.TrimSpace(string(data)), nil
}

/*
Split lines from a string

Takes in a string and returns a slice of strings, where each string is a line from the input string.
*/
func splitLines(s string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(s))

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

// Read ip's from txt and convert to slice, so we can loop
func getIpData() ([]string, error) {
	ipAddresses := []string{}

	// If env is "test" use fake ip address data
	if appEnv == "test" {
		ipAddresses = []string{"127.0.0.1", "127.0.0.2", "127.0.0."}
	} else {
		// Read ip's from txt file
		ipData, err := readFileContent(printerIpFile)
		if err != nil {
			return []string{}, err
		}

		// Convert string to slice, so we can loop
		ipAddresses = splitLines(ipData)
	}

	return ipAddresses, nil
}
