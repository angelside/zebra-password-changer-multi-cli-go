package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// Config
var passwordFile = "password.txt"
var printerIpFile = "printer_ip_list.txt"
var errorFile = "errors.txt"
var appEnv = GoDotEnvVariable("APP_ENV")

/*
var appEnv = func() string {
    value := GoDotEnvVariable("APP_ENV")
    if value == "" {
        return "prod"
    }
    return value
}()
*/

// Colors
var successColor = color.New(color.FgGreen).SprintFunc()
var infoColor = color.New(color.FgBlue).SprintFunc()
var errorColor = color.New(color.FgRed).SprintFunc()

// Use godot package to load/read the .env file and return the value of the key.
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		// We are not required to have a .env file
		//log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
