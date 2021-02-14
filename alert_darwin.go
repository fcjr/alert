package alert

import (
	"github.com/fcjr/alert/internal/nsalert"
)

// Message displays a basic alert with an "OK" button
func Message(title, message string) error {
	return nsalert.Message(title, message)
}

// Error displays a basic error alert with an "OK" button
func Error(title, message string) error {
	return nsalert.Error(title, message)
}

// Question displays an alert with two buttons
//
// returns true iff the the defaultButton was pressed
func Question(title, message, defaultButton, alternateButton string) (bool, error) {
	return nsalert.Question(title, message, defaultButton, alternateButton)
}
