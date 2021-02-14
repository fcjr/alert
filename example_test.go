package alert_test

import (
	"fmt"

	"github.com/fcjr/alert"
)

func ExampleMessage() {
	alert.Message("Example Message", "Example Message Text")
}
func ExampleError() {
	alert.Error("Example Error", "Example Error Text")
}

func ExampleQuestion() {
	resp := alert.Question("Example Question", "Yes or No?", "Yes", "No")
	fmt.Printf("Response: %t\n", resp)

	// Returns true iff 'Yes' is pressed
}
