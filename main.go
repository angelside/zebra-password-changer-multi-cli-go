package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

//
// TODO: Generete log file and next run only run for the ip with error status ?
//

func main() {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	// CLI title
	fmt.Printf("=== %s =====================\n\n", "Zebra password changer")
	fmt.Printf("Env: %s \n", infoColor(appEnv))

	//
	// Password
	//

	// Read password from txt file
	password, err := readFileContent(passwordFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Printf("New password: %s \n", infoColor(password))

	// Validate password
	if err := validatePassword(password, passwordTotalChar); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//
	// ZPL code
	//

	// Prepare the password change zpl code
	zplCode := fmt.Sprintf("^XA^KP%s^JUS^XZ", password)
	//fmt.Printf("ZPL code: %s \n\n", zplCode)

	//
	// IP
	//

	// Read ip's from txt and convert to slice, so that it can be looped
	ipAddresses, err := getIpData(printerIpFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//
	// Prepare the table
	// https://github.com/rodaine/table
	//

	// Set table options
	headerFmt := color.New(color.FgCyan, color.Bold, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	// Set table header
	fmt.Print("\n")
	tbl := table.New("Ip", "Result")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	//
	// Loop
	//

	// Loop ip's
	for _, ip := range ipAddresses {
		//fmt.Println(ip)

		var resultStatus string

		// Validate the IP address
		err := validateIpAddress(ip)
		if err != nil {
			resultStatus = errorColor("INVALID IP")
		} else {
			// Change passwords
			if appEnv == "test" {
				resultStatus = infoColor("TEST")
			} else { // if is not test, Send Data
				// Send the password change request
				err := sendDataOverTcp(ip, "9100", zplCode, time.Second*1)
				if err == nil { // No error
					resultStatus = successColor("OK")
				} else {
					resultStatus = errorColor("NET ERROR")
				}
			}
		}

		// Add results to the table
		tbl.AddRow(ip, resultStatus)
	}

	// Print the table
	tbl.Print()

	// Exit
	fmt.Println("\n\nPress ENTER key for exit.")
	fmt.Scanln()
	os.Exit(0)
}
