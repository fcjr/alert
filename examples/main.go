package main

import (
	"fmt"

	"github.com/fcjr/prompt"
)

func main() {
	prompt.Simple("Example Prompt", "Example Message")
	prompt.Error("Example Error", "Example Error Message")
	resp := prompt.Question("Example Question", "Yes or No?", "Yes", "No")
	fmt.Printf("Response: %t\n", resp)
}
