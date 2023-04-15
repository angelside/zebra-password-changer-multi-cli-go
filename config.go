package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// Config
var passwordTotalChar = 4
var passwordFile = "password.txt"
var printerIpFile = "printer_ip_list.txt"
var appEnv = GoDotEnvVariable("APP_ENV")

// Colors
var successColor = color.New(color.FgGreen).SprintFunc()
var infoColor = color.New(color.FgBlue).SprintFunc()
var errorColor = color.New(color.FgRed).SprintFunc()

// Use godot package to load/read the .env file and return the value of the key.
func GoDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		// We are not required to have an .env file
		//log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}