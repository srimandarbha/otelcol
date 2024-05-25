//go:build windows
// +build windows

package uptime

import (
	"time"

	"golang.org/x/sys/windows"
)

func getUptime() (time.Duration, error) {
	modkernel32 := windows.NewLazySystemDLL("kernel32.dll")
	procGetTickCount64 := modkernel32.NewProc("GetTickCount64")

	ret, _, err := procGetTickCount64.Call()
	if ret == 0 {
		return 0, err
	}
	return time.Duration(ret) * time.Millisecond, nil
}
