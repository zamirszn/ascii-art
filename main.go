package main

import (
	"fmt"
	"os"

	"ascii-art/ascii"
)

func main() {
	// go run . "Hello" => os.Args = ["./main", "Hello"]
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [banners]")
		os.Exit(1)
	}

	input := os.Args[1]
	banners := os.Args[2]

	// Default banner is "standard"
	result, err := ascii.Render(input, banners)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Print(result)
}