package main

import (
	"log"
	"os"
)

func main() {
	// Create or open a log file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()

	// Set up logger to write to both file and console
	log.SetOutput(logFile)
	
	// Log some messages
	log.Println("Application started")
	log.Println("This is an info message")
	log.Println("Processing data...")
	log.Println("Application completed successfully")

	// Also print to console
	println("Logs have been written to app.log")
}
