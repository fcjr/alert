package alert

import (
	"github.com/fcjr/alert/internal/user32"
	"golang.org/x/sys/windows"
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

	// register hook to capture and set custom dialog button text
	var hook user32.HHOOK
	hook = user32.SetWindowsHookEx(user32.WH_CBT,
		(user32.HOOKPROC)(func(nCode int, wparam user32.WPARAM, lparam user32.LPARAM) user32.LRESULT {
			if nCode < 0 {
				return user32.CallNextHookEx(hook, nCode, wparam, lparam)
			}

			if nCode == user32.HCBT_ACTIVATE {
				// set custom button text
				user32.SetDlgItemText(wparam, user32.ID_BUT_YES, defaultButton)
				user32.SetDlgItemText(wparam, user32.ID_BUT_NO, alternateButton)
			}
			return user32.CallNextHookEx(hook, nCode, wparam, lparam)
		}), 0, (user32.DWORD)(windows.GetCurrentThreadId()))
	defer user32.UnhookWindowsHookEx(hook)

	var flags uint
	flags = user32.MB_ICONINFORMATION | user32.MB_YESNO | user32.MB_DEFBUTTON1
	press, err := user32.MessageBoxW(user32.NULL, title, message, flags)
	if err != nil {
		return false, err
	}
	return press == 6, nil
}
