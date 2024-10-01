package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://vyx2d9kw6h72y4h8wox21jknnet5h05p.oastify.com/gotest")
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	
	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)
}
