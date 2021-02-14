package main

import (
	"fmt"

	"github.com/fcjr/prompt"
)

func main() {
	prompt.Message("Example Message", "Example Message Text")
	prompt.Error("Example Error", "Example Error Text")
	resp := prompt.Question("Example Question", "Yes or No?", "Yes", "No")
	fmt.Printf("Response: %t\n", resp)
}
