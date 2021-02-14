package alert_test

import (
	"fmt"

	"github.com/fcjr/alert"
)

func ExampleMessage() {
	_ = alert.Message("Example Message", "Example Message Text")
}
func ExampleError() {
	_ = alert.Error("Example Error", "Example Error Text")
}

func ExampleQuestion() {
	resp, _ := alert.Question("Example Question", "Yes or No?", "Yes", "No")
	fmt.Printf("Response: %t\n", resp)

	// Returns true iff 'Yes' is pressed
}
