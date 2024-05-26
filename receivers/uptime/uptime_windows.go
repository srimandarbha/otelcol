//go:build windows
// +build windows

package uptime

import (
	"time"

	"golang.org/x/sys/windows"

	"go.uber.org/zap"
)

type upTime struct {
	minutes time.Duration
}

type upTimeReader struct {
	logger *zap.Logger
}

func newUpTimeReader(logger *zap.Logger) *upTimeReader {
	return &upTimeReader{
		logger: logger,
	}
}

func (u *upTimeReader) getUptime() (*upTime, error) {
	var upTime upTime
	modkernel32 := windows.NewLazySystemDLL("kernel32.dll")
	procGetTickCount64 := modkernel32.NewProc("GetTickCount64")

	ret, _, err := procGetTickCount64.Call()
	if ret == 0 {
		return &upTime, err
	}

	upTime.minutes = time.Duration(ret) * time.Millisecond
	return &upTime, nil
}
