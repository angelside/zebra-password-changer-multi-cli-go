package main

import (
	"fmt"
	"net"
	"time"
)

func sendDataOverTcp(ip, port string, data string, timeout time.Duration) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		return err // Error during address resolution
	}

	conn, err := net.DialTimeout("tcp4", tcpAddr.String(), timeout)
	if err == nil {
		defer conn.Close()

		payloadBytes := []byte(fmt.Sprintf("%s\r\n\r\n", data))
		_, err = conn.Write(payloadBytes)

		return err // Error during net.DialTimeout
	}

	return err
}
