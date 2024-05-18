package main

import (
	"fmt"
	"time"
)

func formatUptime(uptime time.Duration) string {
	days := uptime / (24 * time.Hour)
	uptime -= days * 24 * time.Hour
	hours := uptime / time.Hour
	uptime -= hours * time.Hour
	minutes := uptime / time.Minute
	uptime -= minutes * time.Minute
	seconds := uptime / time.Second

	return fmt.Sprintf("%d days, %d hours, %d minutes, %d seconds", days, hours, minutes, seconds)
}

func main() {
	uptime, err := getUptime()
	if err != nil {
		fmt.Println("Error getting uptime:", err)
		return
	}

	fmt.Printf("System Uptime: %s\n", formatUptime(uptime))
}
