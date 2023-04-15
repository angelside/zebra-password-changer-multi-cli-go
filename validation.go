package main

import (
	"fmt"
	"net"
	"strings"
)

// Ip address validation
func validateIpAddress(ipAddress string) error {
	if net.ParseIP(ipAddress) == nil {
		return fmt.Errorf("invalid IP address: %s", ipAddress)
	}

	return nil
}

// Password validation / x digit number
func validatePassword(password string, totalChar int) error {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	b := strings.IndexFunc(password, isNotDigit) == -1

	// Invalid password
	if len(password) != totalChar || !b {
		return fmt.Errorf("%s Password is invalid! Please use %d-digit number.", errorColor("[ERROR]"), passwordTotalChar)
	}

	return nil
}
