package main

import (
	"log"
	"os"
)

func main() {
	csvFile, err := os.OpenFile("example.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer csvFile.Close()
}
