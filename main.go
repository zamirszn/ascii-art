package main

import (
	"fmt"
	"os"

	"ascii-art/ascii"
)

func main() {
	argCount := len(os.Args)

	// Check if we have at least the input string
	if argCount < 2 || argCount > 3 {
		fmt.Println("Usage: go run . [STRING] [banners]")
		os.Exit(1)
	}

	input := os.Args[1]

	// Default banner is "standard"
	banners := "standard"

	// If the user provided a second argument, override the default
	if argCount == 3 {
		banners = os.Args[2]
	}

	result, err := ascii.Render(input, banners)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Print(result)
}
