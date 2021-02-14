package main

import (
	"fmt"

	"github.com/fcjr/alert"
)

func main() {
	alert.Message("Example Message", "Example Message Text")
	alert.Error("Example Error", "Example Error Text")
	resp := alert.Question("Example Question", "Yes or No?", "Yes", "No")
	fmt.Printf("Response: %t\n", resp)
}
