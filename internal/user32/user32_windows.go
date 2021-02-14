package user32

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const ERROR_SUCESS = 0
const NULL = 0

var (
	user32            = windows.NewLazyDLL("user32.dll")
	user32MessageBoxW = user32.NewProc("MessageBoxW")
)

const (
	// MessageBox Types

	// The message box contains three push buttons: Abort, Retry, and Ignore.
	MB_ABORTRETRYIGNORE = 0x00000002
	// The message box contains three push buttons: Cancel, Try Again, Continue. Use this message box type instead of MB_ABORTRETRYIGNORE.
	MB_CANCELTRYCONTINUE = 0x00000006
	// Adds a Help button to the message box. When the user clicks the Help button or presses F1, the system sends a WM_HELP message to the owner.
	MB_HELP = 0x00004000
	// The message box contains one push button: OK. This is the default.
	MB_OK = 0x00000000
	// The message box contains two push buttons: OK and Cancel.
	MB_OKCANCEL = 0x00000001
	// The message box contains two push buttons: Retry and Cancel.
	MB_RETRYCANCEL = 0x00000005
	// The message box contains two push buttons: Yes and No.
	MB_YESNO = 0x00000004
	// The message box contains three push buttons: Yes, No, and Cancel.
	MB_YESNOCANCEL = 0x00000003

	// MessageBox Icons

	// An exclamation-point icon appears in the message box.
	MB_ICONEXCLAMATION = 0x00000030
	// An exclamation-point icon appears in the message box.
	MB_ICONWARNING = 0x00000030
	// An icon consisting of a lowercase letter i in a circle appears in the message box.
	MB_ICONINFORMATION = 0x00000040
	// An icon consisting of a lowercase letter i in a circle appears in the message box.
	MB_ICONASTERISK = 0x00000040
	// A question-mark icon appears in the message box. The question-mark message icon is no longer recommended because it does not clearly represent a specific type of message and because the phrasing of a message as a question could apply to any message type. In addition, users can confuse the message symbol question mark with Help information. Therefore, do not use this question mark message symbol in your message boxes. The system continues to support its inclusion only for backward compatibility.
	MB_ICONQUESTION = 0x00000020
	// A stop-sign icon appears in the message box.
	MB_ICONSTOP = 0x00000010
	// A stop-sign icon appears in the message box.
	MB_ICONERROR = 0x00000010
	// A stop-sign icon appears in the message box.
	MB_ICONHAND = 0x00000010

	// MessageBox Default Button Settings
	// MB_DEFBUTTON1 is the default unless MB_DEFBUTTON2, MB_DEFBUTTON3, or MB_DEFBUTTON4 is specified.

	// The first button is the default button.
	MB_DEFBUTTON1 = 0x00000000
	// The second button is the default button.
	MB_DEFBUTTON2 = 0x00000100
	// The third button is the default button.
	MB_DEFBUTTON3 = 0x00000200
	// The fourth button is the default button.
	MB_DEFBUTTON4 = 0x00000300
)

func MessageBoxW(hwnd uintptr, text, caption string, flags uint) (int, error) {
	if err := user32MessageBoxW.Find(); err != nil {
		return 0, err
	}

	txtPntr, err := windows.UTF16PtrFromString(text)
	if err != nil {
		return 0, err
	}

	captPntr, err := windows.UTF16PtrFromString(caption)
	if err != nil {
		return 0, err
	}

	r1, _, _ := user32MessageBoxW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(txtPntr)),
		uintptr(unsafe.Pointer(captPntr)),
		uintptr(flags),
	)
	return int(r1), nil
}
