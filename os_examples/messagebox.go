package main

import (
	"sys/windows"
	"unsafe"
)

func main() {
	var (
		user32dll = windows.NewLazySystemDLL("user32.dll") // load user32.dll.

		messagebox = user32dll.NewProc("MessageBoxW") //our function is messageboxw or messagebox.
	)
	// _ mean is error.
	// msgText and _ (error) init.
	msgText, _ := windows.UTF16PtrFromString("Go is perfect!") //convert from string to UTF16.
	// msgCaption and _ (error) init.
	msgCaption, _ := windows.UTF16PtrFromString("HELLOOO!!!") //convert from string to UTF16.

	// Call the function.
	messagebox.Call( // 3 parameter, (r1 uintptr, r2 uintptr, lastErr error)
		0,
		uintptr(unsafe.Pointer(msgText)),
		//byte is an alias for uint8 and is equivalent to uint8 in all ways.
		uintptr(unsafe.Pointer(msgCaption)),
		0,
	)
}
