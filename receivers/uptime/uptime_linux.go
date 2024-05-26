//go:build linux
// +build linux

package uptime

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"

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
	data, err := ioutil.ReadFile("/proc/uptime")
	var upTime upTime

	if err != nil {
		return &upTime, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return &upTime, err
	}

	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return &upTime, err
	}

	upTime.minutes = time.Duration(seconds) * time.Second

	return &upTime, nil
}
