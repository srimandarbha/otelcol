package uptime

import (
	"github.com/srimandarbha/otelcol/receivers/uptime/internal/metadata"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type Config struct {
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
	// ScraperControllerSettings to configure scraping interval (default: scrape every second)
	scraperhelper.ControllerConfig `mapstructure:",squash"`
}
