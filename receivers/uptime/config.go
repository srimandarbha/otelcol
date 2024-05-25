package uptime

import (
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type Config struct {
	// ScraperControllerSettings to configure scraping interval (default: scrape every second)
	scraperhelper.ScraperControllerOption `mapstructure:",squash"`
}
