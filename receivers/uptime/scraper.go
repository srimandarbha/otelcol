package uptime

import (
	"context"
	"time"

	"github.com/srimandarbha/otelcol/receivers/uptime/internal/metadata"
	"go.opentelemetry.io/collector/pdata/pcommon"
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
		reader:         newUpTimeReader(logger),
	}
}

func (s *scraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	s.logger.Info("Scraping uptime minutes")
	upMin, err := s.reader.getUptime()
	if err != nil {
		return pmetric.Metrics{}, err
	}
	attr := newAttributeReader(s.logger).getAttributes()
	s.recordVmStats(upMin, attr)
	return s.metricsBuilder.Emit(), nil
}

func (s *scraper) recordVmStats(stat *upTime, attr *attributes) {
	now := pcommon.NewTimestampFromTime(time.Now())
	s.metricsBuilder.RecordUptimeSecondsDataPoint(now, int64(stat.minutes), attr.host)
}
