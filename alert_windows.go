package alert

import (
	"fmt"

	"github.com/fcjr/alert/internal/user32"
)

// Info displays a basic alert with an "OK" button
func Info(title, message string) error {
	var flags uint
	flags = user32.MB_ICONINFORMATION | user32.MB_OK
	_, err := user32.MessageBoxW(user32.NULL, title, message, flags)
	if err != nil {
		return err
	}
	return nil
}

// Error displays a basic error alert with an "OK" button
func Error(title, message string) error {
	var flags uint
	flags = user32.MB_ICONERROR | user32.MB_OK
	_, err := user32.MessageBoxW(user32.NULL, title, message, flags)
	if err != nil {
		return err
	}
	return nil
}

// Question displays an alert with two buttons
//
// returns true iff the the defaultButton was pressed
func Question(title, message, defaultButton, alternateButton string) (bool, error) {
	return false, fmt.Errorf("unimplemented")
}
