package uptime

import (
	"os/exec"

	"go.uber.org/zap"
)

type attributes struct {
	host string
}

type attributeReader struct {
	logger *zap.Logger
}

func newAttributeReader(logger *zap.Logger) *attributeReader {
	return &attributeReader{
		logger: logger,
	}
}

func (a *attributeReader) getAttributes() *attributes {
	var h string

	cmd := exec.Command("hostname")
	bytes, err := cmd.Output()
	if err != nil {
		a.logger.Error("failed to execute hostname", zap.Error(err))
		h = "unknown"
	}
	h = string(bytes)
	return &attributes{
		host: h,
	}
}
