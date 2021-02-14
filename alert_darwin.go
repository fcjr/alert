package alert

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit

#import "alert.h"
*/
import "C"
import (
	"unsafe"
)

// Message displays a basic alert with an "OK" button
func Message(title, message string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	C.message(cTitle, cMessage)
}

// Error displays a basic error alert with an "OK" button
func Error(title, message string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	C.error(cTitle, cMessage)
}

// Question displays an alert with two buttons
//
// returns true iff the the defaultButton was pressed
func Question(title, message, defaultButton, alternateButton string) bool {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	cDefaultButton := C.CString(defaultButton)
	defer C.free(unsafe.Pointer(cDefaultButton))
	cAlternateButton := C.CString(alternateButton)
	defer C.free(unsafe.Pointer(cAlternateButton))
	ret := C.question(cTitle, cMessage, cDefaultButton, cAlternateButton)
	return bool(ret)
}
