package main

import (
	"fmt"
	"log"

	"github.com/selimserbes/go-openshowvar/openshowvar"
)

func main() {
	// Create a new OpenShowVar instance with the IP address and port of the TCP server
	osv := openshowvar.NewOpenShowVar("10.145.173.160", 7000)

	// Connect to the TCP server
	err := osv.Connect()
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer osv.Disconnect()

	// Read a variable value
	varName := "COUNT"
	initialValue, err := osv.Read(varName)
	if err != nil {
		log.Fatalf("Failed to read variable: %v", err)
	}
	fmt.Printf("Initial value of %s: %s\n", varName, initialValue)

	// Write a new value to the variable
	newValue := "1"
	_, err = osv.Write(varName, newValue)
	if err != nil {
		log.Fatalf("Failed to write variable: %v", err)
	}
	fmt.Printf("Written new value to %s: %s\n", varName, newValue)

	// Read the variable value again to verify the change
	updatedValue, err := osv.Read(varName)
	if err != nil {
		log.Fatalf("Failed to read variable: %v", err)
	}
	fmt.Printf("Updated value of %s: %s\n", varName, updatedValue)
}
