//go:build linux
// +build linux

package uptime

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getUptime() (time.Duration, error) {
	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return 0, fmt.Errorf("unexpected format of /proc/uptime")
	}

	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}

	return time.Duration(seconds) * time.Second, nil
}
