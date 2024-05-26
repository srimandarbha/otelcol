package uptime_test

import (
	"context"

	"github.com/srimandarbha/otelcol/receivers/uptime/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type scraper struct {
	logger         *zap.Logger              // Logger to log events
	metricsBuilder *metadata.MetricsBuilder // MetricsBuilder to build metrics
	reader         *upTimeReader            // vmStatReader to read vmstat output
}

func newScraper(metricsBuilder *metadata.MetricsBuilder, logger *zap.Logger) *scraper {
	return &scraper{
		logger:         logger,
		metricsBuilder: metricsBuilder,
		reader:         newUptimeReader(logger),
	}
}

func (s *scraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	s.logger.Info("Scraping uptime minutes")

}
